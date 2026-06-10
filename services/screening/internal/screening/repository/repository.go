package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

type ScreeningRecord struct {
	ID                string                 `json:"id"`
	UserID            string                 `json:"user_id"`
	CompanionID       string                 `json:"companion_id"`
	ScreeningType     string                 `json:"screening_type"`
	Status            string                 `json:"status"`
	Provider          string                 `json:"provider"`
	ProviderReference string                 `json:"provider_reference"`
	Result            map[string]interface{} `json:"result"`
	RiskFlags         []string               `json:"risk_flags"`
	VerifiedAt        string                 `json:"verified_at"`
	ExpiresAt         string                 `json:"expires_at"`
	Metadata          map[string]interface{} `json:"metadata"`
	CreatedAt         string                 `json:"created_at"`
	UpdatedAt         string                 `json:"updated_at"`
}

type Repository struct {
	db *database.Postgres
}

func New(db *database.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, s *ScreeningRecord) (*ScreeningRecord, error) {
	result := &ScreeningRecord{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO screening_records (user_id, companion_id, screening_type, status, provider)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, user_id, companion_id, screening_type, status, provider, created_at, updated_at`,
		s.UserID, s.CompanionID, s.ScreeningType, s.Status, s.Provider,
	).Scan(&result.ID, &result.UserID, &result.CompanionID, &result.ScreeningType,
		&result.Status, &result.Provider, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*ScreeningRecord, error) {
	s := &ScreeningRecord{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, user_id, companion_id, screening_type, status, provider, provider_reference,
		 risk_flags, verified_at, expires_at, created_at, updated_at
		 FROM screening_records WHERE id = $1`, id,
	).Scan(&s.ID, &s.UserID, &s.CompanionID, &s.ScreeningType, &s.Status,
		&s.Provider, &s.ProviderReference, &s.RiskFlags, &s.VerifiedAt, &s.ExpiresAt,
		&s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) ListByUser(ctx context.Context, userID string) ([]*ScreeningRecord, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, user_id, companion_id, screening_type, status, provider, risk_flags, created_at
		 FROM screening_records WHERE user_id = $1 ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanRecords(rows)
}

func (r *Repository) ListByCompanion(ctx context.Context, companionID string) ([]*ScreeningRecord, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, user_id, companion_id, screening_type, status, provider, risk_flags, created_at
		 FROM screening_records WHERE companion_id = $1 ORDER BY created_at DESC`, companionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanRecords(rows)
}

func (r *Repository) UpdateStatus(ctx context.Context, id, status, providerRef string, result map[string]interface{}, riskFlags []string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE screening_records SET status = $2, provider_reference = $3, result = $4, risk_flags = $5,
		 verified_at = CASE WHEN $2 = 'verified' THEN NOW() ELSE verified_at END, updated_at = NOW()
		 WHERE id = $1`, id, status, providerRef, result, riskFlags)
	return err
}

func (r *Repository) scanRecords(rows pgx.Rows) ([]*ScreeningRecord, error) {
	var records []*ScreeningRecord
	for rows.Next() {
		s := &ScreeningRecord{}
		err := rows.Scan(&s.ID, &s.UserID, &s.CompanionID, &s.ScreeningType, &s.Status,
			&s.Provider, &s.RiskFlags, &s.CreatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, s)
	}
	return records, nil
}

func (r *Repository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	return r.db.WithTx(ctx, fn)
}
