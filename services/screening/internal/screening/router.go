package screening

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
	"github.com/rs/zerolog"
)

func setupScreeningRouter(cfg *config.Config, log zerolog.Logger) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy","service":"screening"}`))
	}).Methods("GET")

	v1 := r.PathPrefix("/screening").Subrouter()
	v1.HandleFunc("", listScreeningsHandler(log)).Methods("GET")
	v1.HandleFunc("", createScreeningHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}", getScreeningHandler(log)).Methods("GET")
	v1.HandleFunc("/{id}/status", getScreeningStatusHandler(log)).Methods("GET")

	return r
}

type Screening struct {
	ID        string `json:"id"`
	EntityID  string `json:"entity_id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Result    string `json:"result,omitempty"`
	Score     float64 `json:"score,omitempty"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at,omitempty"`
}

func listScreeningsHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Screening{})
	}
}

func createScreeningHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Screening{ID: "new-screening", Status: "initiated"})
	}
}

func getScreeningHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Screening{})
	}
}

func getScreeningStatusHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "completed"})
	}
}
