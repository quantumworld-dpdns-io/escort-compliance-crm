package gateway

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
)

func setupRouter(cfg *config.Config, log zerolog.Logger) *mux.Router {
	r := mux.NewRouter()

	// Middleware
	r.Use(requestIDMiddleware)
	r.Use(loggingMiddleware(log))
	r.Use(corsMiddleware())
	r.Use(rateLimitMiddleware())
	r.Use(metricsMiddleware)

	// Health checks
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/ready", readyHandler).Methods("GET")

	// API v1
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Auth routes (proxy to auth service)
	v1.HandleFunc("/auth/login", proxyHandler(cfg, "auth", "/auth/login")).Methods("POST")
	v1.HandleFunc("/auth/register", proxyHandler(cfg, "auth", "/auth/register")).Methods("POST")
	v1.HandleFunc("/auth/logout", proxyHandler(cfg, "auth", "/auth/logout")).Methods("POST")
	v1.HandleFunc("/auth/refresh", proxyHandler(cfg, "auth", "/auth/refresh")).Methods("POST")
	v1.HandleFunc("/auth/mfa/setup", proxyHandler(cfg, "auth", "/auth/mfa/setup")).Methods("POST")
	v1.HandleFunc("/auth/mfa/verify", proxyHandler(cfg, "auth", "/auth/mfa/verify")).Methods("POST")

	// Companion routes
	v1.HandleFunc("/companions", proxyHandler(cfg, "companion", "/companions")).Methods("GET", "POST")
	v1.HandleFunc("/companions/{id}", proxyHandler(cfg, "companion", "/companions/{id}")).Methods("GET", "PUT", "DELETE")
	v1.HandleFunc("/companions/search", proxyHandler(cfg, "companion", "/companions/search")).Methods("POST")
	v1.HandleFunc("/companions/{id}/availability", proxyHandler(cfg, "companion", "/companions/{id}/availability")).Methods("GET", "PUT")

	// Booking routes
	v1.HandleFunc("/bookings", proxyHandler(cfg, "booking", "/bookings")).Methods("GET", "POST")
	v1.HandleFunc("/bookings/{id}", proxyHandler(cfg, "booking", "/bookings/{id}")).Methods("GET", "PUT", "DELETE")
	v1.HandleFunc("/bookings/{id}/confirm", proxyHandler(cfg, "booking", "/bookings/{id}/confirm")).Methods("POST")
	v1.HandleFunc("/bookings/{id}/cancel", proxyHandler(cfg, "booking", "/bookings/{id}/cancel")).Methods("POST")
	v1.HandleFunc("/bookings/{id}/complete", proxyHandler(cfg, "booking", "/bookings/{id}/complete")).Methods("POST")

	// Compliance routes
	v1.HandleFunc("/compliance/jurisdictions", proxyHandler(cfg, "compliance", "/jurisdictions")).Methods("GET")
	v1.HandleFunc("/compliance/check", proxyHandler(cfg, "compliance", "/check")).Methods("POST")
	v1.HandleFunc("/compliance/alerts", proxyHandler(cfg, "compliance", "/alerts")).Methods("GET")

	// Screening routes
	v1.HandleFunc("/screening", proxyHandler(cfg, "screening", "/screening")).Methods("GET", "POST")
	v1.HandleFunc("/screening/{id}", proxyHandler(cfg, "screening", "/screening/{id}")).Methods("GET")
	v1.HandleFunc("/screening/{id}/status", proxyHandler(cfg, "screening", "/screening/{id}/status")).Methods("GET")

	// Credential routes
	v1.HandleFunc("/credentials", proxyHandler(cfg, "credentials", "/credentials")).Methods("GET", "POST")
	v1.HandleFunc("/credentials/{id}", proxyHandler(cfg, "credentials", "/credentials/{id}")).Methods("GET")
	v1.HandleFunc("/credentials/{id}/verify", proxyHandler(cfg, "credentials", "/credentials/{id}/verify")).Methods("POST")
	v1.HandleFunc("/credentials/{id}/revoke", proxyHandler(cfg, "credentials", "/credentials/{id}/revoke")).Methods("POST")

	// Quantum routes
	v1.HandleFunc("/quantum/pqc/keygen", proxyHandler(cfg, "quantum", "/pqc/keygen")).Methods("POST")
	v1.HandleFunc("/quantum/pqc/encrypt", proxyHandler(cfg, "quantum", "/pqc/encrypt")).Methods("POST")
	v1.HandleFunc("/quantum/pqc/decrypt", proxyHandler(cfg, "quantum", "/pqc/decrypt")).Methods("POST")
	v1.HandleFunc("/quantum/pqc/sign", proxyHandler(cfg, "quantum", "/pqc/sign")).Methods("POST")
	v1.HandleFunc("/quantum/pqc/verify", proxyHandler(cfg, "quantum", "/pqc/verify")).Methods("POST")
	v1.HandleFunc("/quantum/optimize", proxyHandler(cfg, "quantum", "/optimize")).Methods("POST")
	v1.HandleFunc("/quantum/ml/predict", proxyHandler(cfg, "quantum", "/ml/predict")).Methods("POST")
	v1.HandleFunc("/quantum/qrng/random", proxyHandler(cfg, "quantum", "/qrng/random")).Methods("GET")

	// WebSocket endpoint for realtime
	r.HandleFunc("/ws", wsHandler).Methods("GET")

	return r
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy","service":"gateway","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ready","service":"gateway"}`))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// WebSocket upgrade handled by realtime service
	http.Error(w, "WebSocket endpoint", http.StatusSwitchingProtocols)
}
