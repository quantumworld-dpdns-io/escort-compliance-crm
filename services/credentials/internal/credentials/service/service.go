package service

import (
	"context"
	"time"

	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/credentials/internal/credentials/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

type IssueRequest struct {
	HolderID       string `json:"holder_id"`
	CredentialType string `json:"credential_type"`
	Issuer         string `json:"issuer"`
	Subject        string `json:"subject"`
	ExpiresAt      string `json:"expires_at"`
	DocumentURL    string `json:"document_url"`
}

func (s *Service) Issue(ctx context.Context, req *IssueRequest) (*repository.Credential, error) {
	cred := &repository.Credential{
		HolderID:       req.HolderID,
		CredentialType: req.CredentialType,
		Issuer:         req.Issuer,
		Subject:        req.Subject,
		IssuedAt:       time.Now().Format(time.RFC3339),
		ExpiresAt:      req.ExpiresAt,
		Status:         "active",
		DocumentURL:    req.DocumentURL,
	}
	return s.repo.Create(ctx, cred)
}

func (s *Service) GetByID(ctx context.Context, id string) (*repository.Credential, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) ListByHolder(ctx context.Context, holderID string) ([]*repository.Credential, error) {
	return s.repo.ListByHolder(ctx, holderID)
}

func (s *Service) Verify(ctx context.Context, id string) (bool, error) {
	return s.repo.Verify(ctx, id)
}

func (s *Service) Revoke(ctx context.Context, id string) error {
	return s.repo.Revoke(ctx, id)
}
