package config

import "os"

type Config struct {
	Port          string
	DatabaseURL   string
	RedisURL      string
	JWTSecret     string
	RabbitMQURL   string
	KafkaBrokers  string
	ServiceName   string
	Environment   string
	LogLevel      string
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", ""),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/escort_crm?sslmode=disable"),
		RedisURL:     getEnv("REDIS_URL", "localhost:6379"),
		JWTSecret:    getEnv("JWT_SECRET", "dev-secret-change-in-production"),
		RabbitMQURL:  getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
		ServiceName:  getEnv("SERVICE_NAME", "escort-crm"),
		Environment:  getEnv("ENVIRONMENT", "development"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
