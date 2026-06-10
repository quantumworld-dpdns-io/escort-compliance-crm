module github.com/quantumworld-dpdns-io/escort-compliance-crm/services/auth

go 1.22

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared v0.0.0
	golang.org/x/crypto v0.21.0
)

replace github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared => ../shared
