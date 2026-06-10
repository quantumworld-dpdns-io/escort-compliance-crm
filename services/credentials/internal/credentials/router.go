package credentials

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
	"github.com/rs/zerolog"
)

func setupCredentialsRouter(cfg *config.Config, log zerolog.Logger) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy","service":"credentials"}`))
	}).Methods("GET")

	v1 := r.PathPrefix("/credentials").Subrouter()
	v1.HandleFunc("", listCredentialsHandler(log)).Methods("GET")
	v1.HandleFunc("", issueCredentialHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}", getCredentialHandler(log)).Methods("GET")
	v1.HandleFunc("/{id}/verify", verifyCredentialHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}/revoke", revokeCredentialHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}/disclose", selectiveDiscloseHandler(log)).Methods("POST")

	return r
}

type Credential struct {
	ID           string                 `json:"id"`
	Type         string                 `json:"type"`
	Issuer       string                 `json:"issuer"`
	Subject      string                 `json:"subject"`
	Claims       map[string]interface{} `json:"claims"`
	DisclosedClaims []string           `json:"disclosed_claims,omitempty"`
	Proof        string                 `json:"proof"`
	Status       string                 `json:"status"`
	IssuedAt     string                 `json:"issued_at"`
	ExpiresAt    string                 `json:"expires_at,omitempty"`
}

func listCredentialsHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Credential{})
	}
}

func issueCredentialHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Credential{ID: "new-cred", Status: "issued"})
	}
}

func getCredentialHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Credential{})
	}
}

func verifyCredentialHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"valid": true})
	}
}

func revokeCredentialHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "revoked"})
	}
}

func selectiveDiscloseHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"disclosed_claims": []string{},
			"proof":            "zkp-proof-data",
		})
	}
}
