module github.com/quantumworld-dpdns-io/escort-compliance-crm/services/quantum

go 1.22

require (
	github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared v0.0.0
	github.com/gin-gonic/gin v1.9.1
	github.com/rs/zerolog v1.32.0
)

replace github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared => ../shared
