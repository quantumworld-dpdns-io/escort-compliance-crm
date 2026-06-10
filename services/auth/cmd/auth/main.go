package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/auth/internal/auth/handler"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/auth/internal/auth/repository"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/auth/internal/auth/service"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/auth/internal/auth/router"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/config"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/cache"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(database.PostgresConfig{
		URL: cfg.DatabaseURL,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	redisCache, err := cache.New(cache.CacheConfig{
		URL: cfg.RedisURL,
	})
	if err != nil {
		log.Printf("Warning: Redis unavailable, running without cache: %v", err)
	} else {
		defer redisCache.Close()
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = cfg.JWTSecret
	}
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-in-production"
	}

	repo := repository.New(db)
	svc := service.New(repo, redisCache, jwtSecret)
	h := handler.New(svc)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
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
		port = "8081"
	}

	fmt.Printf("Auth service starting on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
