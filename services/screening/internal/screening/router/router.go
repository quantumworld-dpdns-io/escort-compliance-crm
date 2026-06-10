package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/screening/internal/screening/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/screening")
	{
		api.POST("", h.Initiate)
		api.GET("/:id", h.GetByID)
		api.GET("/user/:user_id", h.ListByUser)
		api.POST("/:id/process", h.Process)
	}
}
