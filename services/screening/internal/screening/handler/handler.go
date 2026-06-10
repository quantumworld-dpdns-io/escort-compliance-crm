package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/screening/internal/screening/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Initiate(c *gin.Context) {
	var req service.InitiateScreeningRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record, err := h.svc.Initiate(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"screening": record})
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	record, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "screening record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"screening": record})
}

func (h *Handler) ListByUser(c *gin.Context) {
	userID := c.Param("user_id")
	records, err := h.svc.ListByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"screenings": records})
}

func (h *Handler) Process(c *gin.Context) {
	id := c.Param("id")
	record, err := h.svc.ProcessScreening(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"screening": record})
}
