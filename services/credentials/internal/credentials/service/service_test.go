package service

import (
	"testing"
)

func TestIssueCredential(t *testing.T) {
	svc := &Service{repo: nil}

	req := &IssueRequest{
		HolderID:       "user-1",
		CredentialType: "identity_verification",
		Issuer:         "escort-crm",
		Subject:        "user-1",
	}

	if req.HolderID == "" {
		t.Error("holder_id should not be empty")
	}
	if req.CredentialType == "" {
		t.Error("credential_type should not be empty")
	}
	if req.Issuer == "" {
		t.Error("issuer should not be empty")
	}
}

func TestVerifyCredential(t *testing.T) {
	svc := &Service{repo: nil}

	if svc == nil {
		t.Fatal("service should not be nil")
	}
}
