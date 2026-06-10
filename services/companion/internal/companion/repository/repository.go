package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

type Companion struct {
	ID                 string                 `json:"id"`
	UserID             string                 `json:"user_id"`
	StageName          string                 `json:"stage_name"`
	Bio                string                 `json:"bio"`
	ServicesOffered    []string               `json:"services_offered"`
	HourlyRate         float64                `json:"hourly_rate"`
	Currency           string                 `json:"currency"`
	LocationCity       string                 `json:"location_city"`
	LocationState      string                 `json:"location_state"`
	LocationCountry    string                 `json:"location_country"`
	VerificationStatus string                 `json:"verification_status"`
	IsActive           bool                   `json:"is_active"`
	Rating             float64                `json:"rating"`
	TotalReviews       int                    `json:"total_reviews"`
	Metadata           map[string]interface{} `json:"metadata"`
	CreatedAt          string                 `json:"created_at"`
	UpdatedAt          string                 `json:"updated_at"`
}

type Repository struct {
	db *database.Postgres
}

func New(db *database.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, c *Companion) (*Companion, error) {
	result := &Companion{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO companions (user_id, stage_name, bio, services_offered, hourly_rate, currency,
		 location_city, location_state, location_country, verification_status)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		 RETURNING id, user_id, stage_name, bio, services_offered, hourly_rate, currency,
		 location_city, location_state, location_country, verification_status, is_active,
		 rating, total_reviews, created_at, updated_at`,
		c.UserID, c.StageName, c.Bio, c.ServicesOffered, c.HourlyRate, c.Currency,
		c.LocationCity, c.LocationState, c.LocationCountry, c.VerificationStatus,
	).Scan(&result.ID, &result.UserID, &result.StageName, &result.Bio, &result.ServicesOffered,
		&result.HourlyRate, &result.Currency, &result.LocationCity, &result.LocationState,
		&result.LocationCountry, &result.VerificationStatus, &result.IsActive,
		&result.Rating, &result.TotalReviews, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Companion, error) {
	c := &Companion{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, user_id, stage_name, bio, services_offered, hourly_rate, currency,
		 location_city, location_state, location_country, verification_status, is_active,
		 rating, total_reviews, created_at, updated_at
		 FROM companions WHERE id = $1`, id,
	).Scan(&c.ID, &c.UserID, &c.StageName, &c.Bio, &c.ServicesOffered, &c.HourlyRate,
		&c.Currency, &c.LocationCity, &c.LocationState, &c.LocationCountry,
		&c.VerificationStatus, &c.IsActive, &c.Rating, &c.TotalReviews, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]*Companion, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, user_id, stage_name, bio, services_offered, hourly_rate, currency,
		 location_city, location_state, location_country, verification_status, is_active,
		 rating, total_reviews, created_at, updated_at
		 FROM companions WHERE is_active = true ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companions []*Companion
	for rows.Next() {
		c := &Companion{}
		err := rows.Scan(&c.ID, &c.UserID, &c.StageName, &c.Bio, &c.ServicesOffered, &c.HourlyRate,
			&c.Currency, &c.LocationCity, &c.LocationState, &c.LocationCountry,
			&c.VerificationStatus, &c.IsActive, &c.Rating, &c.TotalReviews, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		companions = append(companions, c)
	}
	return companions, nil
}

func (r *Repository) Update(ctx context.Context, id string, c *Companion) (*Companion, error) {
	result := &Companion{}
	err := r.db.Pool.QueryRow(ctx,
		`UPDATE companions SET stage_name=$2, bio=$3, services_offered=$4, hourly_rate=$5,
		 location_city=$6, location_state=$7, location_country=$8, updated_at=NOW()
		 WHERE id=$1
		 RETURNING id, user_id, stage_name, bio, services_offered, hourly_rate, currency,
		 location_city, location_state, location_country, verification_status, is_active,
		 rating, total_reviews, created_at, updated_at`,
		id, c.StageName, c.Bio, c.ServicesOffered, c.HourlyRate,
		c.LocationCity, c.LocationState, c.LocationCountry,
	).Scan(&result.ID, &result.UserID, &result.StageName, &result.Bio, &result.ServicesOffered,
		&result.HourlyRate, &result.Currency, &result.LocationCity, &result.LocationState,
		&result.LocationCountry, &result.VerificationStatus, &result.IsActive,
		&result.Rating, &result.TotalReviews, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx, `DELETE FROM companions WHERE id = $1`, id)
	return err
}

func (r *Repository) Search(ctx context.Context, query string, limit, offset int) ([]*Companion, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, user_id, stage_name, bio, services_offered, hourly_rate, currency,
		 location_city, location_state, location_country, verification_status, is_active,
		 rating, total_reviews, created_at, updated_at
		 FROM companions WHERE is_active = true
		 AND (stage_name ILIKE $1 OR bio ILIKE $1 OR location_city ILIKE $1)
		 ORDER BY rating DESC LIMIT $2 OFFSET $3`,
		"%"+query+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companions []*Companion
	for rows.Next() {
		c := &Companion{}
		err := rows.Scan(&c.ID, &c.UserID, &c.StageName, &c.Bio, &c.ServicesOffered, &c.HourlyRate,
			&c.Currency, &c.LocationCity, &c.LocationState, &c.LocationCountry,
			&c.VerificationStatus, &c.IsActive, &c.Rating, &c.TotalReviews, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		companions = append(companions, c)
	}
	return companions, nil
}

func (r *Repository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	return r.db.WithTx(ctx, fn)
}
