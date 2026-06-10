package service

import (
	"context"
	"errors"
	"time"

	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/booking/internal/booking/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

type CreateBookingRequest struct {
	CompanionID     string  `json:"companion_id"`
	ScheduledAt     string  `json:"scheduled_at"`
	DurationMinutes int     `json:"duration_minutes"`
	LocationAddress string  `json:"location_address"`
	LocationLat     float64 `json:"location_lat"`
	LocationLng     float64 `json:"location_lng"`
	TotalAmount     float64 `json:"total_amount"`
	Currency        string  `json:"currency"`
	Notes           string  `json:"notes"`
}

func (s *Service) Create(ctx context.Context, clientID string, req *CreateBookingRequest) (*repository.Booking, error) {
	scheduledAt, err := time.Parse(time.RFC3339, req.ScheduledAt)
	if err != nil {
		return nil, errors.New("invalid scheduled_at format")
	}

	if scheduledAt.Before(time.Now()) {
		return nil, errors.New("cannot book in the past")
	}

	if req.DurationMinutes <= 0 {
		req.DurationMinutes = 60
	}

	hasConflict, err := s.repo.CheckConflicts(ctx, req.CompanionID, scheduledAt, req.DurationMinutes)
	if err != nil {
		return nil, err
	}
	if hasConflict {
		return nil, errors.New("companion has a conflicting booking")
	}

	if req.Currency == "" {
		req.Currency = "USD"
	}

	b := &repository.Booking{
		ClientID:        clientID,
		CompanionID:     req.CompanionID,
		ScheduledAt:     req.ScheduledAt,
		DurationMinutes: req.DurationMinutes,
		LocationAddress: req.LocationAddress,
		LocationLat:     req.LocationLat,
		LocationLng:     req.LocationLng,
		TotalAmount:     req.TotalAmount,
		Currency:        req.Currency,
		Notes:           req.Notes,
	}

	return s.repo.Create(ctx, b)
}

func (s *Service) GetByID(ctx context.Context, id string) (*repository.Booking, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) ListByClient(ctx context.Context, clientID string, limit, offset int) ([]*repository.Booking, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.repo.ListByClient(ctx, clientID, limit, offset)
}

func (s *Service) ListByCompanion(ctx context.Context, companionID string, limit, offset int) ([]*repository.Booking, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.repo.ListByCompanion(ctx, companionID, limit, offset)
}

func (s *Service) Confirm(ctx context.Context, id string) error {
	b, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("booking not found")
	}
	if b.Status != "pending" {
		return errors.New("can only confirm pending bookings")
	}
	return s.repo.UpdateStatus(ctx, id, "confirmed")
}

func (s *Service) Complete(ctx context.Context, id string) error {
	b, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("booking not found")
	}
	if b.Status != "confirmed" {
		return errors.New("can only complete confirmed bookings")
	}
	return s.repo.Complete(ctx, id)
}

func (s *Service) Cancel(ctx context.Context, id, reason string) error {
	b, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("booking not found")
	}
	if b.Status == "completed" || b.Status == "cancelled" {
		return errors.New("cannot cancel completed or already cancelled bookings")
	}
	return s.repo.Cancel(ctx, id, reason)
}
