package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	log logger.Logger
}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler(log logger.Logger) *HealthHandler {
	return &HealthHandler{log: log}
}

// HealthCheck returns a simple health status
// @Summary Health check
// @Description Returns OK if the service is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
