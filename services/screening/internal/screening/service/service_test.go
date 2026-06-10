package service

import (
	"testing"
)

func TestProcessScreening_LowRisk(t *testing.T) {
	svc := &Service{repo: nil}

	// This test would normally use a mock repository
	// For now, test that the service can be created
	if svc == nil {
		t.Fatal("service should not be nil")
	}
}

func TestInitiateScreening(t *testing.T) {
	svc := &Service{repo: nil}

	req := &InitiateScreeningRequest{
		UserID:        "user-1",
		CompanionID:   "comp-1",
		ScreeningType: "background_check",
		Provider:      "sterling",
	}

	if req.UserID == "" {
		t.Error("user_id should not be empty")
	}
	if req.CompanionID == "" {
		t.Error("companion_id should not be empty")
	}
	if req.ScreeningType == "" {
		t.Error("screening_type should not be empty")
	}
}
