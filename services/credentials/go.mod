module github.com/quantumworld-dpdns-io/escort-compliance-crm/services/credentials

go 1.22

require (
	github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared v0.0.0
	github.com/gin-gonic/gin v1.9.1
	github.com/jackc/pgx/v5 v5.5.5
	github.com/rs/zerolog v1.32.0
)

replace github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared => ../shared
