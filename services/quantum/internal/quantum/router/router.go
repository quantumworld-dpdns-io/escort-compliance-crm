package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/quantum/internal/quantum/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handler) {
	api := r.Group("/api/v1/quantum")
	{
		api.POST("/pqc/keygen", h.GeneratePQCKeyPair)
		api.POST("/pqc/encrypt", h.PQCEncrypt)
		api.POST("/pqc/sign", h.PQCSign)
		api.POST("/pqc/verify", h.VerifyPQC)
		api.POST("/optimization/qaoa", h.QAOAOptimize)
		api.POST("/optimization/vqe", h.VQEOptimize)
		api.POST("/qrng", h.GenerateQRNG)
		api.POST("/ml/predict", h.QuantumPredict)
		api.POST("/qkd/exchange", h.QKDExchange)
	}
}
