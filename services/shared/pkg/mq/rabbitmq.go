package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

type RabbitMQConfig struct {
	URL string
}

func NewRabbitMQ(cfg RabbitMQConfig) (*RabbitMQ, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("connect to rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("open channel: %w", err)
	}

	return &RabbitMQ{conn: conn, channel: ch}, nil
}

type Message struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func (r *RabbitMQ) Publish(ctx context.Context, exchange, routingKey string, msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return r.channel.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Body:         data,
		},
	)
}

func (r *RabbitMQ) Consume(ctx context.Context, queue string, handler func(Message) error) error {
	msgs, err := r.channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-msgs:
			if !ok {
				return fmt.Errorf("channel closed")
			}

			var m Message
			if err := json.Unmarshal(msg.Body, &m); err != nil {
				msg.Nack(false, false)
				continue
			}

			if err := handler(m); err != nil {
				msg.Nack(false, true)
				continue
			}
			msg.Ack(false)
		}
	}
}

func (r *RabbitMQ) DeclareQueue(name string) error {
	_, err := r.channel.QueueDeclare(name, true, false, false, false, nil)
	return err
}

func (r *RabbitMQ) Close() error {
	r.channel.Close()
	return r.conn.Close()
}
