package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/booking/internal/booking/handler"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/booking/internal/booking/repository"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/booking/internal/booking/service"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/booking/internal/booking/router"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

func main() {
	cfg := config.Load()
	db, err := database.NewPostgres(database.PostgresConfig{URL: cfg.DatabaseURL})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := repository.New(db)
	svc := service.New(repo)
	h := handler.New(svc)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.GET("/readyz", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(503, gin.H{"status": "not ready"})
			return
		}
		c.JSON(200, gin.H{"status": "ready"})
	})
	router.RegisterRoutes(r, h)

	port := cfg.Port
	if port == "" {
		port = "8083"
	}
	fmt.Printf("Booking service starting on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
