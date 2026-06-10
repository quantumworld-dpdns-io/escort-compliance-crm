package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/auth/internal/auth/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
		api.POST("/refresh", h.Refresh)
		api.POST("/logout", authMiddleware(), h.Logout)
		api.GET("/me", authMiddleware(), h.Me)
		api.POST("/mfa/setup", authMiddleware(), h.MFASetup)
		api.POST("/mfa/verify", authMiddleware(), h.MFAVerify)
		api.POST("/change-password", authMiddleware(), h.ChangePassword)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
		c.Set("access_token", token)
		c.Next()
	}
}
