package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

type Booking struct {
	ID                string  `json:"id"`
	ClientID          string  `json:"client_id"`
	CompanionID       string  `json:"companion_id"`
	Status            string  `json:"status"`
	ScheduledAt       string  `json:"scheduled_at"`
	DurationMinutes   int     `json:"duration_minutes"`
	LocationAddress   string  `json:"location_address"`
	LocationLat       float64 `json:"location_lat"`
	LocationLng       float64 `json:"location_lng"`
	TotalAmount       float64 `json:"total_amount"`
	Currency          string  `json:"currency"`
	PaymentStatus     string  `json:"payment_status"`
	Notes             string  `json:"notes"`
	CancellationReason string `json:"cancellation_reason"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type Repository struct {
	db *database.Postgres
}

func New(db *database.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, b *Booking) (*Booking, error) {
	result := &Booking{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO bookings (client_id, companion_id, status, scheduled_at, duration_minutes,
		 location_address, location_lat, location_lng, total_amount, currency, notes)
		 VALUES ($1, $2, 'pending', $3, $4, $5, $6, $7, $8, $9, $10)
		 RETURNING id, client_id, companion_id, status, scheduled_at, duration_minutes,
		 location_address, location_lat, location_lng, total_amount, currency, payment_status, created_at, updated_at`,
		b.ClientID, b.CompanionID, b.ScheduledAt, b.DurationMinutes,
		b.LocationAddress, b.LocationLat, b.LocationLng, b.TotalAmount, b.Currency, b.Notes,
	).Scan(&result.ID, &result.ClientID, &result.CompanionID, &result.Status, &result.ScheduledAt,
		&result.DurationMinutes, &result.LocationAddress, &result.LocationLat, &result.LocationLng,
		&result.TotalAmount, &result.Currency, &result.PaymentStatus, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Booking, error) {
	b := &Booking{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, client_id, companion_id, status, scheduled_at, duration_minutes,
		 location_address, location_lat, location_lng, total_amount, currency, payment_status,
		 notes, cancellation_reason, created_at, updated_at
		 FROM bookings WHERE id = $1`, id,
	).Scan(&b.ID, &b.ClientID, &b.CompanionID, &b.Status, &b.ScheduledAt,
		&b.DurationMinutes, &b.LocationAddress, &b.LocationLat, &b.LocationLng,
		&b.TotalAmount, &b.Currency, &b.PaymentStatus, &b.Notes, &b.CancellationReason,
		&b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *Repository) ListByClient(ctx context.Context, clientID string, limit, offset int) ([]*Booking, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, client_id, companion_id, status, scheduled_at, duration_minutes,
		 location_address, location_lat, location_lng, total_amount, currency, payment_status,
		 notes, cancellation_reason, created_at, updated_at
		 FROM bookings WHERE client_id = $1 ORDER BY scheduled_at DESC LIMIT $2 OFFSET $3`,
		clientID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanBookings(rows)
}

func (r *Repository) ListByCompanion(ctx context.Context, companionID string, limit, offset int) ([]*Booking, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, client_id, companion_id, status, scheduled_at, duration_minutes,
		 location_address, location_lat, location_lng, total_amount, currency, payment_status,
		 notes, cancellation_reason, created_at, updated_at
		 FROM bookings WHERE companion_id = $1 ORDER BY scheduled_at DESC LIMIT $2 OFFSET $3`,
		companionID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanBookings(rows)
}

func (r *Repository) UpdateStatus(ctx context.Context, id, status string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE bookings SET status = $2, updated_at = NOW() WHERE id = $1`, id, status)
	return err
}

func (r *Repository) Complete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE bookings SET status = 'completed', completed_at = NOW(), updated_at = NOW() WHERE id = $1`, id)
	return err
}

func (r *Repository) Cancel(ctx context.Context, id, reason string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE bookings SET status = 'cancelled', cancellation_reason = $2, cancelled_at = NOW(), updated_at = NOW() WHERE id = $1`,
		id, reason)
	return err
}

func (r *Repository) scanBookings(rows pgx.Rows) ([]*Booking, error) {
	var bookings []*Booking
	for rows.Next() {
		b := &Booking{}
		err := rows.Scan(&b.ID, &b.ClientID, &b.CompanionID, &b.Status, &b.ScheduledAt,
			&b.DurationMinutes, &b.LocationAddress, &b.LocationLat, &b.LocationLng,
			&b.TotalAmount, &b.Currency, &b.PaymentStatus, &b.Notes, &b.CancellationReason,
			&b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *Repository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	return r.db.WithTx(ctx, fn)
}

func (r *Repository) CheckConflicts(ctx context.Context, companionID string, scheduledAt time.Time, durationMinutes int) (bool, error) {
	var count int
	err := r.db.Pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM bookings
		 WHERE companion_id = $1 AND status NOT IN ('cancelled', 'failed')
		 AND scheduled_at < $2 + ($3 || ' minutes')::INTERVAL
		 AND scheduled_at + (duration_minutes || ' minutes')::INTERVAL > $2`,
		companionID, scheduledAt, durationMinutes).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
