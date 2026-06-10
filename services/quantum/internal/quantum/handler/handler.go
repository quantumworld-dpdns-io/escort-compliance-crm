package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/quantum/internal/quantum/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GeneratePQCKeyPair(c *gin.Context) {
	var req struct {
		Algorithm string `json:"algorithm"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		req.Algorithm = "kyber768"
	}
	keyPair, err := h.svc.GeneratePQCKeyPair(c.Request.Context(), req.Algorithm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"keypair": keyPair})
}

func (h *Handler) PQCEncrypt(c *gin.Context) {
	var req struct {
		Plaintext string `json:"plaintext"`
		Algorithm string `json:"algorithm"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enc, err := h.svc.PQCEncrypt(c.Request.Context(), req.Plaintext, req.Algorithm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"encryption": enc})
}

func (h *Handler) PQCSign(c *gin.Context) {
	var req struct {
		Message   string `json:"message"`
		Algorithm string `json:"algorithm"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sig, err := h.svc.PQCSign(c.Request.Context(), req.Message, req.Algorithm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"signature": sig, "algorithm": req.Algorithm})
}

func (h *Handler) VerifyPQC(c *gin.Context) {
	var req struct {
		Message   string `json:"message"`
		Signature string `json:"signature"`
		Algorithm string `json:"algorithm"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	valid, err := h.svc.VerifyPQC(c.Request.Context(), req.Message, req.Signature, req.Algorithm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valid": valid})
}

func (h *Handler) QAOAOptimize(c *gin.Context) {
	var req struct {
		ProblemType string                 `json:"problem_type"`
		Params      map[string]interface{} `json:"params"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.svc.QAOAOptimize(c.Request.Context(), req.ProblemType, req.Params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"optimization": result})
}

func (h *Handler) VQEOptimize(c *gin.Context) {
	var req struct {
		ProblemType string                 `json:"problem_type"`
		Params      map[string]interface{} `json:"params"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.svc.VQEOptimize(c.Request.Context(), req.ProblemType, req.Params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"optimization": result})
}

func (h *Handler) GenerateQRNG(c *gin.Context) {
	var req struct {
		Count int `json:"count"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		req.Count = 256
	}
	result, err := h.svc.GenerateQRNG(c.Request.Context(), req.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"qrng": result})
}

func (h *Handler) QuantumPredict(c *gin.Context) {
	var req struct {
		Model    string             `json:"model"`
		Features map[string]float64 `json:"features"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.svc.QuantumPredict(c.Request.Context(), req.Model, req.Features)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"prediction": result})
}

func (h *Handler) QKDExchange(c *gin.Context) {
	result, err := h.svc.QKDExchange(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"qkd": result})
}
