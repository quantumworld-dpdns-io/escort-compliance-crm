package companion

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
	"github.com/rs/zerolog"
)

func setupCompanionRouter(cfg *config.Config, log zerolog.Logger) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy","service":"companion"}`))
	}).Methods("GET")

	v1 := r.PathPrefix("/companions").Subrouter()
	v1.HandleFunc("", listCompanionsHandler(log)).Methods("GET")
	v1.HandleFunc("", createCompanionHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}", getCompanionHandler(log)).Methods("GET")
	v1.HandleFunc("/{id}", updateCompanionHandler(log)).Methods("PUT")
	v1.HandleFunc("/{id}", deleteCompanionHandler(log)).Methods("DELETE")
	v1.HandleFunc("/search", searchCompanionsHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}/availability", getAvailabilityHandler(log)).Methods("GET")
	v1.HandleFunc("/{id}/availability", updateAvailabilityHandler(log)).Methods("PUT")

	return r
}

type CompanionProfile struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio,omitempty"`
	AvatarURL   string    `json:"avatar_url,omitempty"`
	Jurisdiction string  `json:"jurisdiction"`
	Languages   []string  `json:"languages,omitempty"`
	Rate        float64   `json:"rate"`
	RateUnit    string    `json:"rate_unit"`
	Services    []string  `json:"services,omitempty"`
	Verified    bool      `json:"verified"`
	Active      bool      `json:"active"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

func listCompanionsHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]CompanionProfile{})
	}
}

func createCompanionHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(CompanionProfile{ID: "new-id"})
	}
}

func getCompanionHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CompanionProfile{})
	}
}

func updateCompanionHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CompanionProfile{})
	}
}

func deleteCompanionHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

func searchCompanionsHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]CompanionProfile{})
	}
}

func getAvailabilityHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"slots": []string{}})
	}
}

func updateAvailabilityHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
	}
}
