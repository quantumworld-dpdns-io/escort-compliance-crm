package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/ml/internal/ml/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/ml")
	{
		api.POST("/predict", h.Predict)
		api.GET("/models", h.ListModels)
		api.GET("/models/:name", h.GetModel)
		api.POST("/models/:name/retrain", h.Retrain)
		api.POST("/rag/query", h.RAGQuery)
		api.POST("/agents/execute", h.AgentExecute)
	}
}
