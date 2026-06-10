package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

type Payment struct {
	ID             string  `json:"id"`
	BookingID      string  `json:"booking_id"`
	PayerID        string  `json:"payer_id"`
	PayeeID        string  `json:"payee_id"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	Method         string  `json:"method"`
	Provider       string  `json:"provider"`
	ProviderRef    string  `json:"provider_ref"`
	Status         string  `json:"status"`
	FailureReason  string  `json:"failure_reason"`
	RefundAmount   float64 `json:"refund_amount"`
	RefundedAt     string  `json:"refunded_at"`
	Metadata       map[string]interface{} `json:"metadata"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type Repository struct {
	db *database.Postgres
}

func New(db *database.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, p *Payment) (*Payment, error) {
	result := &Payment{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO payments (booking_id, payer_id, payee_id, amount, currency, method, provider, status)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, 'pending')
		 RETURNING id, booking_id, payer_id, payee_id, amount, currency, method, provider, status, created_at, updated_at`,
		p.BookingID, p.PayerID, p.PayeeID, p.Amount, p.Currency, p.Method, p.Provider,
	).Scan(&result.ID, &result.BookingID, &result.PayerID, &result.PayeeID, &result.Amount,
		&result.Currency, &result.Method, &result.Provider, &result.Status, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Payment, error) {
	p := &Payment{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, booking_id, payer_id, payee_id, amount, currency, method, provider, provider_ref,
		 status, failure_reason, refund_amount, refunded_at, created_at, updated_at
		 FROM payments WHERE id = $1`, id,
	).Scan(&p.ID, &p.BookingID, &p.PayerID, &p.PayeeID, &p.Amount, &p.Currency, &p.Method,
		&p.Provider, &p.ProviderRef, &p.Status, &p.FailureReason, &p.RefundAmount, &p.RefundedAt,
		&p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *Repository) ListByPayer(ctx context.Context, payerID string, limit, offset int) ([]*Payment, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, booking_id, payer_id, payee_id, amount, currency, method, provider, status, created_at
		 FROM payments WHERE payer_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		payerID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanPayments(rows)
}

func (r *Repository) UpdateStatus(ctx context.Context, id, status, providerRef string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE payments SET status = $2, provider_ref = $3, updated_at = NOW() WHERE id = $1`,
		id, status, providerRef)
	return err
}

func (r *Repository) ProcessRefund(ctx context.Context, id string, amount float64) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE payments SET status = 'refunded', refund_amount = $2, refunded_at = NOW(), updated_at = NOW() WHERE id = $1`,
		id, amount)
	return err
}

func (r *Repository) scanPayments(rows pgx.Rows) ([]*Payment, error) {
	var payments []*Payment
	for rows.Next() {
		p := &Payment{}
		err := rows.Scan(&p.ID, &p.BookingID, &p.PayerID, &p.PayeeID, &p.Amount,
			&p.Currency, &p.Method, &p.Provider, &p.Status, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}

func (r *Repository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	return r.db.WithTx(ctx, fn)
}
