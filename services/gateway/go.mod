module github.com/quantumworld-dpdns-io/escort-compliance-crm/services/gateway

go 1.22

require (
	github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared v0.0.0
	github.com/gin-gonic/gin v1.9.1
	github.com/gorilla/mux v1.8.1
	github.com/rs/zerolog v1.32.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.49.0
	go.opentelemetry.io/otel v1.24.0
)

replace github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared => ../shared
