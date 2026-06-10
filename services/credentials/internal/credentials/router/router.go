package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/credentials/internal/credentials/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/credentials")
	{
		api.POST("", h.Issue)
		api.GET("/:id", h.GetByID)
		api.GET("/holder/:holder_id", h.ListByHolder)
		api.GET("/:id/verify", h.Verify)
		api.POST("/:id/revoke", h.Revoke)
	}
}
