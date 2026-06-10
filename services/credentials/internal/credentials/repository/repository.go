package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

type Credential struct {
	ID               string                 `json:"id"`
	HolderID         string                 `json:"holder_id"`
	CredentialType   string                 `json:"credential_type"`
	Issuer           string                 `json:"issuer"`
	Subject          string                 `json:"subject"`
	IssuedAt         string                 `json:"issued_at"`
	ExpiresAt        string                 `json:"expires_at"`
	RevokedAt        string                 `json:"revoked_at"`
	Status           string                 `json:"status"`
	DocumentURL      string                 `json:"document_url"`
	VerificationData map[string]interface{} `json:"verification_data"`
	Metadata         map[string]interface{} `json:"metadata"`
	CreatedAt        string                 `json:"created_at"`
	UpdatedAt        string                 `json:"updated_at"`
}

type Repository struct {
	db *database.Postgres
}

func New(db *database.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, cred *Credential) (*Credential, error) {
	result := &Credential{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO credentials (holder_id, credential_type, issuer, subject, issued_at, expires_at, status, document_url)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id, holder_id, credential_type, issuer, subject, issued_at, expires_at, status, document_url, created_at, updated_at`,
		cred.HolderID, cred.CredentialType, cred.Issuer, cred.Subject, cred.IssuedAt, cred.ExpiresAt, "active", cred.DocumentURL,
	).Scan(&result.ID, &result.HolderID, &result.CredentialType, &result.Issuer, &result.Subject,
		&result.IssuedAt, &result.ExpiresAt, &result.Status, &result.DocumentURL, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Credential, error) {
	c := &Credential{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, holder_id, credential_type, issuer, subject, issued_at, expires_at, revoked_at,
		 status, document_url, created_at, updated_at
		 FROM credentials WHERE id = $1`, id,
	).Scan(&c.ID, &c.HolderID, &c.CredentialType, &c.Issuer, &c.Subject, &c.IssuedAt,
		&c.ExpiresAt, &c.RevokedAt, &c.Status, &c.DocumentURL, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repository) ListByHolder(ctx context.Context, holderID string) ([]*Credential, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, holder_id, credential_type, issuer, subject, issued_at, expires_at, status, document_url, created_at
		 FROM credentials WHERE holder_id = $1 ORDER BY created_at DESC`, holderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.scanCredentials(rows)
}

func (r *Repository) Revoke(ctx context.Context, id string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE credentials SET status = 'revoked', revoked_at = NOW(), updated_at = NOW() WHERE id = $1`, id)
	return err
}

func (r *Repository) Verify(ctx context.Context, id string) (bool, error) {
	var status string
	err := r.db.Pool.QueryRow(ctx,
		`SELECT status FROM credentials WHERE id = $1`, id).Scan(&status)
	if err != nil {
		return false, err
	}
	return status == "active", nil
}

func (r *Repository) scanCredentials(rows pgx.Rows) ([]*Credential, error) {
	var credentials []*Credential
	for rows.Next() {
		c := &Credential{}
		err := rows.Scan(&c.ID, &c.HolderID, &c.CredentialType, &c.Issuer, &c.Subject,
			&c.IssuedAt, &c.ExpiresAt, &c.Status, &c.DocumentURL, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		credentials = append(credentials, c)
	}
	return credentials, nil
}

func (r *Repository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	return r.db.WithTx(ctx, fn)
}
