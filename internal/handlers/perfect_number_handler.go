package handlers

import (
	"net/http"
	"time"

	"perfect-numbers-api/internal/models"
	"perfect-numbers-api/internal/services"

	"github.com/gin-gonic/gin"
)

type PerfectNumberHandler struct {
	service services.PerfectNumberService
}

func NewPerfectNumberHandler(service services.PerfectNumberService) *PerfectNumberHandler {
	return &PerfectNumberHandler{
		service: service,
	}
}

func (h *PerfectNumberHandler) FindPerfectNumbers(c *gin.Context) {
	var req models.PerfectNumberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:     "JSON inválido: " + err.Error(),
			Code:      "INVALID_JSON",
			Timestamp: time.Now(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:     err.Error(),
			Code:      "VALIDATION_ERROR",
			Timestamp: time.Now(),
		})
		return
	}

	response := h.service.FindPerfectNumbers(req.Start, req.End)
	c.JSON(http.StatusOK, response)
}

func (h *PerfectNumberHandler) Health(c *gin.Context) {
	response := models.HealthResponse{
		Status:    "OK",
		Message:   "Perfect Numbers API está funcionando",
		Version:   "2.0.0",
		Timestamp: time.Now(),
		Uptime:    "N/A",
	}
	c.JSON(http.StatusOK, response)
}

func (h *PerfectNumberHandler) APIInfo(c *gin.Context) {
	response := models.APIInfoResponse{
		Name:        "Perfect Numbers API",
		Version:     "2.0.0",
		Description: "API REST para encontrar números perfeitos em um range especificado",
		Endpoints: map[string]string{
			"POST /api/v1/perfect-numbers": "Encontrar números perfeitos",
			"GET /api/v1/health":           "Health check",
			"GET /api/v1/info":             "Informações da API",
		},
		Timestamp: time.Now(),
	}
	c.JSON(http.StatusOK, response)
}
