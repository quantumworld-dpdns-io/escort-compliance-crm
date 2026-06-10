package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/payments/internal/payments/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(c *gin.Context) {
	var req service.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment, err := h.svc.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"payment": payment})
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	payment, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func (h *Handler) ListByPayer(c *gin.Context) {
	payerID := c.Param("payer_id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	payments, err := h.svc.ListByPayer(c.Request.Context(), payerID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payments": payments})
}

func (h *Handler) Process(c *gin.Context) {
	id := c.Param("id")
	payment, err := h.svc.ProcessPayment(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func (h *Handler) Refund(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Refund(c.Request.Context(), id, req.Amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "refund processed"})
}

func (h *Handler) Cancel(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Cancel(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "payment cancelled"})
}
