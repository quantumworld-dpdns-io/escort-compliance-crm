package service

import (
	"context"
	"fmt"

	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/payments/internal/payments/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

type CreatePaymentRequest struct {
	BookingID string  `json:"booking_id"`
	PayerID   string  `json:"payer_id"`
	PayeeID   string  `json:"payee_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Method    string  `json:"method"`
	Provider  string  `json:"provider"`
}

func (s *Service) Create(ctx context.Context, req *CreatePaymentRequest) (*repository.Payment, error) {
	if req.Currency == "" {
		req.Currency = "USD"
	}
	if req.Method == "" {
		req.Method = "card"
	}
	p := &repository.Payment{
		BookingID: req.BookingID,
		PayerID:   req.PayerID,
		PayeeID:   req.PayeeID,
		Amount:    req.Amount,
		Currency:  req.Currency,
		Method:    req.Method,
		Provider:  req.Provider,
	}
	return s.repo.Create(ctx, p)
}

func (s *Service) GetByID(ctx context.Context, id string) (*repository.Payment, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) ListByPayer(ctx context.Context, payerID string, limit, offset int) ([]*repository.Payment, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.repo.ListByPayer(ctx, payerID, limit, offset)
}

func (s *Service) ProcessPayment(ctx context.Context, id string) (*repository.Payment, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("payment not found")
	}
	if p.Status != "pending" {
		return nil, fmt.Errorf("payment already processed: %s", p.Status)
	}

	providerRef := fmt.Sprintf("PAY-%s", id[:8])
	err = s.repo.UpdateStatus(ctx, id, "completed", providerRef)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, id)
}

func (s *Service) Refund(ctx context.Context, id string, amount float64) error {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("payment not found")
	}
	if p.Status != "completed" {
		return fmt.Errorf("can only refund completed payments")
	}
	if amount > p.Amount {
		return fmt.Errorf("refund amount exceeds payment amount")
	}
	return s.repo.ProcessRefund(ctx, id, amount)
}

func (s *Service) Cancel(ctx context.Context, id string) error {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("payment not found")
	}
	if p.Status != "pending" {
		return fmt.Errorf("can only cancel pending payments")
	}
	return s.repo.UpdateStatus(ctx, id, "cancelled", "")
}
