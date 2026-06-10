package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/compliance/internal/compliance/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/compliance")
	{
		api.POST("/jurisdictions", h.CreateJurisdiction)
		api.GET("/jurisdictions", h.ListJurisdictions)
		api.GET("/jurisdictions/:id", h.GetJurisdiction)
		api.POST("/regulations", h.CreateRegulation)
		api.GET("/regulations/:jurisdiction_id", h.ListRegulations)
		api.POST("/checks", h.RunCheck)
		api.GET("/checks/:user_id", h.ListChecks)
	}
}
