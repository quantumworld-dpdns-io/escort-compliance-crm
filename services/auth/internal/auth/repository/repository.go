package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

type User struct {
	ID                string     `json:"id"`
	Email             string     `json:"email"`
	PasswordHash      string     `json:"-"`
	Name              string     `json:"name"`
	Role              string     `json:"role"`
	EmailVerified     bool       `json:"email_verified"`
	MFAEnabled        bool       `json:"mfa_enabled"`
	MFASecret         string     `json:"-"`
	FailedLoginAttempts int      `json:"-"`
	LockedUntil       *time.Time `json:"-"`
	LastLoginAt       *time.Time `json:"last_login_at"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type Repository struct {
	db *database.Postgres
}

func New(db *database.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, email, passwordHash, name, role string) (*User, error) {
	user := &User{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO users (email, password_hash, name, role)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, email, name, role, email_verified, mfa_enabled, created_at, updated_at`,
		email, passwordHash, name, role,
	).Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.EmailVerified, &user.MFAEnabled, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, email, password_hash, name, role, email_verified, mfa_enabled,
		        failed_login_attempts, locked_until, last_login_at, created_at, updated_at
		 FROM users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Role,
		&user.EmailVerified, &user.MFAEnabled, &user.FailedLoginAttempts,
		&user.LockedUntil, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*User, error) {
	user := &User{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, email, name, role, email_verified, mfa_enabled, last_login_at, created_at, updated_at
		 FROM users WHERE id = $1`, id,
	).Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.EmailVerified,
		&user.MFAEnabled, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateLastLogin(ctx context.Context, userID string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE users SET last_login_at = NOW(), failed_login_attempts = 0, updated_at = NOW() WHERE id = $1`, userID)
	return err
}

func (r *Repository) IncrementFailedLogin(ctx context.Context, email string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE users SET failed_login_attempts = failed_login_attempts + 1,
		 locked_until = CASE WHEN failed_login_attempts >= 4 THEN NOW() + INTERVAL '15 minutes' ELSE locked_until END,
		 updated_at = NOW() WHERE email = $1`, email)
	return err
}

func (r *Repository) CreateSession(ctx context.Context, userID, tokenHash, ip, userAgent string, expiresAt time.Time) error {
	_, err := r.db.Pool.Exec(ctx,
		`INSERT INTO sessions (user_id, token_hash, ip_address, user_agent, expires_at)
		 VALUES ($1, $2, $3, $4, $5)`, userID, tokenHash, ip, userAgent, expiresAt)
	return err
}

func (r *Repository) RevokeSession(ctx context.Context, tokenHash string) error {
	_, err := r.db.Pool.Exec(ctx,
		`UPDATE sessions SET revoked_at = NOW() WHERE token_hash = $1`, tokenHash)
	return err
}

func (r *Repository) CreateAuditLog(ctx context.Context, userID, action, resource, resourceID, ip, userAgent string, details map[string]interface{}) error {
	_, err := r.db.Pool.Exec(ctx,
		`INSERT INTO audit_log (user_id, action, resource, resource_id, ip_address, user_agent, details)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`, userID, action, resource, resourceID, ip, userAgent, details)
	return err
}

func (r *Repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.Pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`, email).Scan(&exists)
	return exists, err
}

func (r *Repository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	return r.db.WithTx(ctx, fn)
}
