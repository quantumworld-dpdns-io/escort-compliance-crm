package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/payments/internal/payments/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/payments")
	{
		api.POST("", h.Create)
		api.GET("/:id", h.GetByID)
		api.GET("/payer/:payer_id", h.ListByPayer)
		api.POST("/:id/process", h.Process)
		api.POST("/:id/refund", h.Refund)
		api.POST("/:id/cancel", h.Cancel)
	}
}
