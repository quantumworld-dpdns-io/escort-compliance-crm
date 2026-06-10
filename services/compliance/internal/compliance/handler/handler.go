package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/compliance/internal/compliance/repository"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/compliance/internal/compliance/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateJurisdiction(c *gin.Context) {
	var req repository.Jurisdiction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.svc.CreateJurisdiction(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"jurisdiction": result})
}

func (h *Handler) GetJurisdiction(c *gin.Context) {
	id := c.Param("id")
	j, err := h.svc.GetJurisdiction(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "jurisdiction not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jurisdiction": j})
}

func (h *Handler) ListJurisdictions(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	jurisdictions, err := h.svc.ListJurisdictions(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jurisdictions": jurisdictions})
}

func (h *Handler) CreateRegulation(c *gin.Context) {
	var req repository.Regulation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.svc.CreateRegulation(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"regulation": result})
}

func (h *Handler) ListRegulations(c *gin.Context) {
	jurisdictionID := c.Param("jurisdiction_id")
	regulations, err := h.svc.ListRegulations(c.Request.Context(), jurisdictionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"regulations": regulations})
}

func (h *Handler) RunCheck(c *gin.Context) {
	var req struct {
		UserID         string `json:"user_id"`
		JurisdictionID string `json:"jurisdiction_id"`
		CheckType      string `json:"check_type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	check, err := h.svc.RunComplianceCheck(c.Request.Context(), req.UserID, req.JurisdictionID, req.CheckType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"check": check})
}

func (h *Handler) ListChecks(c *gin.Context) {
	userID := c.Param("user_id")
	checks, err := h.svc.ListChecksByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"checks": checks})
}
