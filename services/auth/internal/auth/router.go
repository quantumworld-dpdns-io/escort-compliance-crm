package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
	"github.com/rs/zerolog"
)

func setupAuthRouter(cfg *config.Config, log zerolog.Logger) *mux.Router {
	r := mux.NewRouter()
	r.Use(corsMiddleware)
	r.Use(requestIDMiddleware)

	// Health
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// Auth endpoints
	v1 := r.PathPrefix("/auth").Subrouter()
	v1.HandleFunc("/login", loginHandler(cfg, log)).Methods("POST")
	v1.HandleFunc("/register", registerHandler(cfg, log)).Methods("POST")
	v1.HandleFunc("/logout", logoutHandler(log)).Methods("POST")
	v1.HandleFunc("/refresh", refreshHandler(cfg, log)).Methods("POST")
	v1.HandleFunc("/mfa/setup", mfaSetupHandler(log)).Methods("POST")
	v1.HandleFunc("/mfa/verify", mfaVerifyHandler(log)).Methods("POST")
	v1.HandleFunc("/password/reset", passwordResetHandler(log)).Methods("POST")
	v1.HandleFunc("/password/change", passwordChangeHandler(log)).Methods("POST")

	return r
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"healthy","service":"auth"}`))
}

func loginHandler(cfg *config.Config, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Login logic: validate credentials, generate JWT, return tokens
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"login endpoint","token":"placeholder"}`))
	}
}

func registerHandler(cfg *config.Config, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"register endpoint"}`))
	}
}

func logoutHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"logged out"}`))
	}
}

func refreshHandler(cfg *config.Config, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token":"new_access_token","refresh":"new_refresh_token"}`))
	}
}

func mfaSetupHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"secret":"JBSWY3DPEHPK3PXP","qr_code":"otpauth://..."}`))
	}
}

func mfaVerifyHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"verified":true}`))
	}
}

func passwordResetHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"reset email sent"}`))
	}
}

func passwordChangeHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"password changed"}`))
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
