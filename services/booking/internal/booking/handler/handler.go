package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/booking/internal/booking/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(c *gin.Context) {
	var req service.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clientID, _ := c.Get("user_id")
	booking, err := h.svc.Create(c.Request.Context(), clientID.(string), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"booking": booking})
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	booking, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "booking not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"booking": booking})
}

func (h *Handler) ListByClient(c *gin.Context) {
	clientID := c.Param("client_id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	bookings, err := h.svc.ListByClient(c.Request.Context(), clientID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func (h *Handler) ListByCompanion(c *gin.Context) {
	companionID := c.Param("companion_id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	bookings, err := h.svc.ListByCompanion(c.Request.Context(), companionID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func (h *Handler) Confirm(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Confirm(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "booking confirmed"})
}

func (h *Handler) Complete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Complete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "booking completed"})
}

func (h *Handler) Cancel(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Reason string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		req.Reason = "no reason provided"
	}
	if err := h.svc.Cancel(c.Request.Context(), id, req.Reason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "booking cancelled"})
}
