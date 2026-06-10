package service

import (
	"testing"
)

func TestCalculateRisk_IllegalJurisdiction(t *testing.T) {
	svc := &Service{repo: nil}

	jurisdiction := &Jurisdiction{
		Name:        "Test Jurisdiction",
		LegalStatus: "illegal",
	}

	regulations := []*Regulation{
		{Severity: "critical", Category: "safety"},
		{Severity: "high", Category: "labor"},
	}

	score, findings, recommendations := svc.calculateRisk(jurisdiction, regulations, "comprehensive")

	if score < 60 {
		t.Errorf("expected risk score >= 60 for illegal jurisdiction, got %f", score)
	}
	if len(findings) == 0 {
		t.Error("expected findings for illegal jurisdiction")
	}
	if len(recommendations) == 0 {
		t.Error("expected recommendations for illegal jurisdiction")
	}
}

func TestCalculateRisk_LegalJurisdiction(t *testing.T) {
	svc := &Service{repo: nil}

	jurisdiction := &Jurisdiction{
		Name:        "Legal Jurisdiction",
		LegalStatus: "legal",
	}

	regulations := []*Regulation{}

	score, _, _ := svc.calculateRisk(jurisdiction, regulations, "basic")

	if score > 25 {
		t.Errorf("expected low risk score for legal jurisdiction, got %f", score)
	}
}

func TestCalculateRisk_RestrictedJurisdiction(t *testing.T) {
	svc := &Service{repo: nil}

	jurisdiction := &Jurisdiction{
		Name:        "Restricted Jurisdiction",
		LegalStatus: "restricted",
	}

	regulations := []*Regulation{}

	score, findings, _ := svc.calculateRisk(jurisdiction, regulations, "basic")

	if score < 25 {
		t.Errorf("expected medium risk score for restricted jurisdiction, got %f", score)
	}
	if len(findings) == 0 {
		t.Error("expected findings for restricted jurisdiction")
	}
}

func TestCalculateRisk_CriticalRegulations(t *testing.T) {
	svc := &Service{repo: nil}

	jurisdiction := &Jurisdiction{
		Name:        "Moderate Jurisdiction",
		LegalStatus: "regulated",
	}

	regulations := []*Regulation{
		{Severity: "critical", Category: "safety"},
		{Severity: "critical", Category: "health"},
		{Severity: "critical", Category: "labor"},
		{Severity: "critical", Category: "tax"},
		{Severity: "critical", Category: "reporting"},
	}

	score, findings, _ := svc.calculateRisk(jurisdiction, regulations, "comprehensive")

	if score < 50 {
		t.Errorf("expected high risk score with 5 critical regulations, got %f", score)
	}
	if len(findings) < 2 {
		t.Errorf("expected at least 2 findings, got %d", len(findings))
	}
}

func TestCalculateRisk_ScoreCap(t *testing.T) {
	svc := &Service{repo: nil}

	jurisdiction := &Jurisdiction{
		Name:        "Max Risk",
		LegalStatus: "criminalized",
	}

	regulations := []*Regulation{
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
		{Severity: "critical"},
	}

	score, _, _ := svc.calculateRisk(jurisdiction, regulations, "full")

	if score > 100 {
		t.Errorf("risk score should not exceed 100, got %f", score)
	}
}
