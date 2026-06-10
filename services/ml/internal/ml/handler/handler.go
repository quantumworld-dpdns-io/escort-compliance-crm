package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/ml/internal/ml/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Predict(c *gin.Context) {
	var req struct {
		Model    string             `json:"model"`
		Features map[string]float64 `json:"features"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prediction, err := h.svc.Predict(c.Request.Context(), req.Model, req.Features)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if prediction == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "model not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"prediction": prediction})
}

func (h *Handler) ListModels(c *gin.Context) {
	models := h.svc.ListModels(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{"models": models})
}

func (h *Handler) GetModel(c *gin.Context) {
	name := c.Param("name")
	model := h.svc.GetModel(c.Request.Context(), name)
	if model == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "model not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"model": model})
}

func (h *Handler) Retrain(c *gin.Context) {
	name := c.Param("name")
	model, err := h.svc.RetrainModel(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "model not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"model": model})
}

func (h *Handler) RAGQuery(c *gin.Context) {
	var req struct {
		Query  string `json:"query"`
		TopK   int    `json:"top_k"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.TopK <= 0 {
		req.TopK = 5
	}
	c.JSON(http.StatusOK, gin.H{
		"query": req.Query,
		"results": []map[string]interface{}{
			{"chunk": "Sample RAG result 1", "score": 0.95, "source": "compliance-doc-1"},
			{"chunk": "Sample RAG result 2", "score": 0.87, "source": "regulation-doc-3"},
			{"chunk": "Sample RAG result 3", "score": 0.82, "source": "policy-doc-2"},
		},
		"total": 3,
	})
}

func (h *Handler) AgentExecute(c *gin.Context) {
	var req struct {
		Agent  string                 `json:"agent"`
		Action string                 `json:"action"`
		Params map[string]interface{} `json:"params"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"agent":  req.Agent,
		"action": req.Action,
		"status": "completed",
		"result": map[string]interface{}{
			"message": "Agent execution simulated",
			"output":  "Task completed successfully",
		},
	})
}
