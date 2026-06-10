package service

import (
	"context"
	"math"
	"strings"

	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/compliance/internal/compliance/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateJurisdiction(ctx context.Context, j *repository.Jurisdiction) (*repository.Jurisdiction, error) {
	return s.repo.CreateJurisdiction(ctx, j)
}

func (s *Service) GetJurisdiction(ctx context.Context, id string) (*repository.Jurisdiction, error) {
	return s.repo.FindJurisdictionByID(ctx, id)
}

func (s *Service) ListJurisdictions(ctx context.Context, limit, offset int) ([]*repository.Jurisdiction, error) {
	if limit <= 0 {
		limit = 20
	}
	return s.repo.ListJurisdictions(ctx, limit, offset)
}

func (s *Service) CreateRegulation(ctx context.Context, reg *repository.Regulation) (*repository.Regulation, error) {
	return s.repo.CreateRegulation(ctx, reg)
}

func (s *Service) ListRegulations(ctx context.Context, jurisdictionID string) ([]*repository.Regulation, error) {
	return s.repo.ListRegulationsByJurisdiction(ctx, jurisdictionID)
}

func (s *Service) RunComplianceCheck(ctx context.Context, userID, jurisdictionID, checkType string) (*repository.ComplianceCheck, error) {
	jurisdiction, err := s.repo.FindJurisdictionByID(ctx, jurisdictionID)
	if err != nil {
		return nil, err
	}

	regulations, err := s.repo.ListRegulationsByJurisdiction(ctx, jurisdictionID)
	if err != nil {
		return nil, err
	}

	riskScore, findings, recommendations := s.calculateRisk(jurisdiction, regulations, checkType)

	riskLevel := "low"
	if riskScore >= 75 {
		riskLevel = "critical"
	} else if riskScore >= 50 {
		riskLevel = "high"
	} else if riskScore >= 25 {
		riskLevel = "medium"
	}

	check := &repository.ComplianceCheck{
		UserID:          userID,
		JurisdictionID:  jurisdictionID,
		CheckType:       checkType,
		Status:          "completed",
		RiskScore:       riskScore,
		RiskLevel:       riskLevel,
		Findings:        findings,
		Recommendations: recommendations,
	}

	return s.repo.CreateComplianceCheck(ctx, check)
}

func (s *Service) ListChecksByUser(ctx context.Context, userID string) ([]*repository.ComplianceCheck, error) {
	return s.repo.ListComplianceChecksByUser(ctx, userID)
}

func (s *Service) calculateRisk(jurisdiction *repository.Jurisdiction, regulations []*repository.Regulation, checkType string) (float64, []interface{}, []interface{}) {
	var findings []interface{}
	var recommendations []interface{}

	score := 0.0

	switch strings.ToLower(jurisdiction.LegalStatus) {
	case "illegal", "criminalized":
		score += 60
		findings = append(findings, map[string]string{
			"type": "legal_status", "severity": "critical",
			"detail": "Activity is illegal/criminalized in this jurisdiction",
		})
		recommendations = append(recommendations, map[string]string{
			"type": "legal", "priority": "critical",
			"detail": "Do not operate in this jurisdiction without legal counsel",
		})
	case "restricted", "regulated":
		score += 30
		findings = append(findings, map[string]string{
			"type": "legal_status", "severity": "high",
			"detail": "Activity is restricted/regulated in this jurisdiction",
		})
		recommendations = append(recommendations, map[string]string{
			"type": "legal", "priority": "high",
			"detail": "Ensure full compliance with local regulations",
		})
	case "legal", "decriminalized":
		score += 5
	}

	criticalRegs := 0
	for _, reg := range regulations {
		if reg.Severity == "critical" || reg.Severity == "high" {
			criticalRegs++
		}
	}
	if criticalRegs > 0 {
		score += math.Min(float64(criticalRegs)*10, 30)
		findings = append(findings, map[string]interface{}{
			"type": "regulations", "severity": "high",
			"detail": "Found critical/high severity regulations", "count": criticalRegs,
		})
	}

	if score > 100 {
		score = 100
	}

	return score, findings, recommendations
}
