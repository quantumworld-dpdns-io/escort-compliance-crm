package gateway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
)

func proxyHandler(cfg *config.Config, service, targetPath string) http.HandlerFunc {
	serviceURLs := map[string]string{
		"auth":        "http://localhost:8081",
		"companion":   "http://localhost:8082",
		"booking":     "http://localhost:8083",
		"compliance":  "http://localhost:8084",
		"screening":   "http://localhost:8085",
		"credentials": "http://localhost:8086",
		"payments":    "http://localhost:8087",
		"messaging":   "http://localhost:8088",
		"realtime":    "http://localhost:8089",
		"quantum":     "http://localhost:8090",
		"ml":          "http://localhost:8091",
		"data":        "http://localhost:8092",
	}

	serviceURL, ok := serviceURLs[service]
	if !ok {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, fmt.Sprintf(`{"error":"unknown service: %s"}`, service), http.StatusBadGateway)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		target, _ := url.Parse(serviceURL)
		proxy := httputil.NewSingleHostReverseProxy(target)

		// Rewrite path
		vars := mux.Vars(r)
		path := targetPath
		for k, v := range vars {
			path = fmt.Sprintf("%s/%s", path, v)
		}
		r.URL.Path = path

		// Forward headers
		r.Header.Set("X-Forwarded-For", r.RemoteAddr)
		r.Header.Set("X-Service-Name", "gateway")

		proxy.ServeHTTP(w, r)
	}
}
