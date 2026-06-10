package service

import (
	"context"
	"errors"

	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/companion/internal/companion/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

type CreateCompanionRequest struct {
	StageName       string   `json:"stage_name"`
	Bio             string   `json:"bio"`
	ServicesOffered []string `json:"services_offered"`
	HourlyRate      float64  `json:"hourly_rate"`
	Currency        string   `json:"currency"`
	LocationCity    string   `json:"location_city"`
	LocationState   string   `json:"location_state"`
	LocationCountry string   `json:"location_country"`
}

func (s *Service) Create(ctx context.Context, userID string, req *CreateCompanionRequest) (*repository.Companion, error) {
	c := &repository.Companion{
		UserID:             userID,
		StageName:          req.StageName,
		Bio:                req.Bio,
		ServicesOffered:    req.ServicesOffered,
		HourlyRate:         req.HourlyRate,
		Currency:           req.Currency,
		LocationCity:       req.LocationCity,
		LocationState:      req.LocationState,
		LocationCountry:    req.LocationCountry,
		VerificationStatus: "pending",
	}
	if c.Currency == "" {
		c.Currency = "USD"
	}
	return s.repo.Create(ctx, c)
}

func (s *Service) GetByID(ctx context.Context, id string) (*repository.Companion, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) List(ctx context.Context, limit, offset int) ([]*repository.Companion, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) Update(ctx context.Context, id string, req *CreateCompanionRequest) (*repository.Companion, error) {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("companion not found")
	}

	if req.StageName != "" {
		existing.StageName = req.StageName
	}
	if req.Bio != "" {
		existing.Bio = req.Bio
	}
	if len(req.ServicesOffered) > 0 {
		existing.ServicesOffered = req.ServicesOffered
	}
	if req.HourlyRate > 0 {
		existing.HourlyRate = req.HourlyRate
	}
	if req.LocationCity != "" {
		existing.LocationCity = req.LocationCity
	}
	if req.LocationState != "" {
		existing.LocationState = req.LocationState
	}
	if req.LocationCountry != "" {
		existing.LocationCountry = req.LocationCountry
	}

	return s.repo.Update(ctx, id, existing)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Search(ctx context.Context, query string, limit, offset int) ([]*repository.Companion, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.repo.Search(ctx, query, limit, offset)
}
