package service

import (
	"context"
	"math/rand"

	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/screening/internal/screening/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

type InitiateScreeningRequest struct {
	UserID         string `json:"user_id"`
	CompanionID    string `json:"companion_id"`
	ScreeningType  string `json:"screening_type"`
	Provider       string `json:"provider"`
}

func (s *Service) Initiate(ctx context.Context, req *InitiateScreeningRequest) (*repository.ScreeningRecord, error) {
	record := &repository.ScreeningRecord{
		UserID:        req.UserID,
		CompanionID:   req.CompanionID,
		ScreeningType: req.ScreeningType,
		Status:        "pending",
		Provider:      req.Provider,
	}
	return s.repo.Create(ctx, record)
}

func (s *Service) GetByID(ctx context.Context, id string) (*repository.ScreeningRecord, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) ListByUser(ctx context.Context, userID string) ([]*repository.ScreeningRecord, error) {
	return s.repo.ListByUser(ctx, userID)
}

func (s *Service) ListByCompanion(ctx context.Context, companionID string) ([]*repository.ScreeningRecord, error) {
	return s.repo.ListByCompanion(ctx, companionID)
}

func (s *Service) ProcessScreening(ctx context.Context, id string) (*repository.ScreeningRecord, error) {
	record, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	riskFlags := []string{}
	result := map[string]interface{}{}

	score := rand.Float64() * 100

	if score > 80 {
		riskFlags = append(riskFlags, "high_risk")
		result["recommendation"] = "manual_review"
	} else if score > 50 {
		riskFlags = append(riskFlags, "medium_risk")
		result["recommendation"] = "additional_verification"
	} else {
		result["recommendation"] = "approved"
	}

	result["score"] = score
	result["provider"] = record.Provider

	status := "verified"
	if len(riskFlags) > 0 {
		status = "flagged"
	}

	err = s.repo.UpdateStatus(ctx, id, status, "PROV-"+id[:8], result, riskFlags)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, id)
}
