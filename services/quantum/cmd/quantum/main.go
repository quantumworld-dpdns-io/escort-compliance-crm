package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/quantum/internal/quantum/handler"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/quantum/internal/quantum/service"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/quantum/internal/quantum/router"
)

func main() {
	svc := service.New()
	h := handler.New(svc)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.GET("/readyz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ready"}) })
	router.RegisterRoutes(r, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8091"
	}
	fmt.Printf("Quantum service starting on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	_ = http.StatusOK
}
