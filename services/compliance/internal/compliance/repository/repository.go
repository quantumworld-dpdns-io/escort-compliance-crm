package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

type Jurisdiction struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	State       string  `json:"state"`
	City        string  `json:"city"`
	LegalStatus string  `json:"legal_status"`
	LegalNotes  string  `json:"legal_notes"`
	Metadata    map[string]interface{} `json:"metadata"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type Regulation struct {
	ID              string  `json:"id"`
	JurisdictionID  string  `json:"jurisdiction_id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Severity        string  `json:"severity"`
	EffectiveDate   string  `json:"effective_date"`
	ExpiryDate      string  `json:"expiry_date"`
	TextURL         string  `json:"text_url"`
	Metadata        map[string]interface{} `json:"metadata"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type ComplianceCheck struct {
	ID              string  `json:"id"`
	UserID          string  `json:"user_id"`
	JurisdictionID  string  `json:"jurisdiction_id"`
	CheckType       string  `json:"check_type"`
	Status          string  `json:"status"`
	RiskScore       float64 `json:"risk_score"`
	RiskLevel       string  `json:"risk_level"`
	Findings        []interface{} `json:"findings"`
	Recommendations []interface{} `json:"recommendations"`
	CheckedAt       string  `json:"checked_at"`
	ExpiresAt       string  `json:"expires_at"`
	Metadata        map[string]interface{} `json:"metadata"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type Repository struct {
	db *database.Postgres
}

func New(db *database.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateJurisdiction(ctx context.Context, j *Jurisdiction) (*Jurisdiction, error) {
	result := &Jurisdiction{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO jurisdictions (name, country, state, city, legal_status, legal_notes)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, name, country, state, city, legal_status, legal_notes, created_at, updated_at`,
		j.Name, j.Country, j.State, j.City, j.LegalStatus, j.LegalNotes,
	).Scan(&result.ID, &result.Name, &result.Country, &result.State, &result.City,
		&result.LegalStatus, &result.LegalNotes, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindJurisdictionByID(ctx context.Context, id string) (*Jurisdiction, error) {
	j := &Jurisdiction{}
	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, name, country, state, city, legal_status, legal_notes, created_at, updated_at
		 FROM jurisdictions WHERE id = $1`, id,
	).Scan(&j.ID, &j.Name, &j.Country, &j.State, &j.City, &j.LegalStatus, &j.LegalNotes, &j.CreatedAt, &j.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (r *Repository) ListJurisdictions(ctx context.Context, limit, offset int) ([]*Jurisdiction, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, name, country, state, city, legal_status, legal_notes, created_at, updated_at
		 FROM jurisdictions ORDER BY name LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jurisdictions []*Jurisdiction
	for rows.Next() {
		j := &Jurisdiction{}
		err := rows.Scan(&j.ID, &j.Name, &j.Country, &j.State, &j.City, &j.LegalStatus, &j.LegalNotes, &j.CreatedAt, &j.UpdatedAt)
		if err != nil {
			return nil, err
		}
		jurisdictions = append(jurisdictions, j)
	}
	return jurisdictions, nil
}

func (r *Repository) CreateRegulation(ctx context.Context, reg *Regulation) (*Regulation, error) {
	result := &Regulation{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO regulations (jurisdiction_id, name, description, category, severity, effective_date, expiry_date, text_url)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id, jurisdiction_id, name, description, category, severity, effective_date, expiry_date, text_url, created_at, updated_at`,
		reg.JurisdictionID, reg.Name, reg.Description, reg.Category, reg.Severity, reg.EffectiveDate, reg.ExpiryDate, reg.TextURL,
	).Scan(&result.ID, &result.JurisdictionID, &result.Name, &result.Description, &result.Category,
		&result.Severity, &result.EffectiveDate, &result.ExpiryDate, &result.TextURL, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) ListRegulationsByJurisdiction(ctx context.Context, jurisdictionID string) ([]*Regulation, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, jurisdiction_id, name, description, category, severity, effective_date, expiry_date, text_url, created_at, updated_at
		 FROM regulations WHERE jurisdiction_id = $1 ORDER BY severity DESC, name`, jurisdictionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var regulations []*Regulation
	for rows.Next() {
		reg := &Regulation{}
		err := rows.Scan(&reg.ID, &reg.JurisdictionID, &reg.Name, &reg.Description, &reg.Category,
			&reg.Severity, &reg.EffectiveDate, &reg.ExpiryDate, &reg.TextURL, &reg.CreatedAt, &reg.UpdatedAt)
		if err != nil {
			return nil, err
		}
		regulations = append(regulations, reg)
	}
	return regulations, nil
}

func (r *Repository) CreateComplianceCheck(ctx context.Context, check *ComplianceCheck) (*ComplianceCheck, error) {
	result := &ComplianceCheck{}
	err := r.db.Pool.QueryRow(ctx,
		`INSERT INTO compliance_checks (user_id, jurisdiction_id, check_type, status, risk_score, risk_level, findings, recommendations)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id, user_id, jurisdiction_id, check_type, status, risk_score, risk_level, created_at, updated_at`,
		check.UserID, check.JurisdictionID, check.CheckType, check.Status, check.RiskScore, check.RiskLevel,
		check.Findings, check.Recommendations,
	).Scan(&result.ID, &result.UserID, &result.JurisdictionID, &result.CheckType, &result.Status,
		&result.RiskScore, &result.RiskLevel, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) ListComplianceChecksByUser(ctx context.Context, userID string) ([]*ComplianceCheck, error) {
	rows, err := r.db.Pool.Query(ctx,
		`SELECT id, user_id, jurisdiction_id, check_type, status, risk_score, risk_level, created_at, updated_at
		 FROM compliance_checks WHERE user_id = $1 ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var checks []*ComplianceCheck
	for rows.Next() {
		check := &ComplianceCheck{}
		err := rows.Scan(&check.ID, &check.UserID, &check.JurisdictionID, &check.CheckType, &check.Status,
			&check.RiskScore, &check.RiskLevel, &check.CreatedAt, &check.UpdatedAt)
		if err != nil {
			return nil, err
		}
		checks = append(checks, check)
	}
	return checks, nil
}

func (r *Repository) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	return r.db.WithTx(ctx, fn)
}
