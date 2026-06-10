package booking

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
	"github.com/rs/zerolog"
)

func setupBookingRouter(cfg *config.Config, log zerolog.Logger) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy","service":"booking"}`))
	}).Methods("GET")

	v1 := r.PathPrefix("/bookings").Subrouter()
	v1.HandleFunc("", listBookingsHandler(log)).Methods("GET")
	v1.HandleFunc("", createBookingHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}", getBookingHandler(log)).Methods("GET")
	v1.HandleFunc("/{id}", updateBookingHandler(log)).Methods("PUT")
	v1.HandleFunc("/{id}", deleteBookingHandler(log)).Methods("DELETE")
	v1.HandleFunc("/{id}/confirm", confirmBookingHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}/cancel", cancelBookingHandler(log)).Methods("POST")
	v1.HandleFunc("/{id}/complete", completeBookingHandler(log)).Methods("POST")

	return r
}

type Booking struct {
	ID            string  `json:"id"`
	ClientID      string  `json:"client_id"`
	CompanionID   string  `json:"companion_id"`
	Status        string  `json:"status"`
	StartTime     string  `json:"start_time"`
	EndTime       string  `json:"end_time"`
	Jurisdiction  string  `json:"jurisdiction"`
	Location      string  `json:"location,omitempty"`
	Rate          float64 `json:"rate"`
	Currency      string  `json:"currency"`
	Notes         string  `json:"notes,omitempty"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func listBookingsHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Booking{})
	}
}

func createBookingHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Booking{ID: "new-booking", Status: "pending"})
	}
}

func getBookingHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Booking{})
	}
}

func updateBookingHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Booking{})
	}
}

func deleteBookingHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

func confirmBookingHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Booking{Status: "confirmed"})
	}
}

func cancelBookingHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Booking{Status: "cancelled"})
	}
}

func completeBookingHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Booking{Status: "completed"})
	}
}
