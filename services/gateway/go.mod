module github.com/quantumworld-dpdns-io/escort-compliance-crm/services/gateway

go 1.22

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared v0.0.0
)

replace github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared => ../shared
