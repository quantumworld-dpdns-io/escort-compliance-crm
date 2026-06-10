package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/companion/internal/companion/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/companions")
	{
		api.POST("", h.Create)
		api.GET("", h.List)
		api.GET("/search", h.Search)
		api.GET("/:id", h.GetByID)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}
