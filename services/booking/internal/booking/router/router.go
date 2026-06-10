package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/booking/internal/booking/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/bookings")
	{
		api.POST("", h.Create)
		api.GET("/:id", h.GetByID)
		api.GET("/client/:client_id", h.ListByClient)
		api.GET("/companion/:companion_id", h.ListByCompanion)
		api.POST("/:id/confirm", h.Confirm)
		api.POST("/:id/complete", h.Complete)
		api.POST("/:id/cancel", h.Cancel)
	}
}
