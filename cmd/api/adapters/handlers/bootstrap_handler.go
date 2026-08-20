package handlers

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// BootstrapHandler handles the one-time admin bootstrap endpoint.
type BootstrapHandler struct {
	bootstrapSvc core.BootstrapService
	// token, when set, must be presented as `Authorization: Bearer <token>`. Bootstrap
	// creates the FIRST admin and is otherwise unauthenticated, so without this whoever
	// reaches a fresh deployment first owns it. Empty = no check (dev, or an endpoint
	// already unreachable from outside).
	token string
	log   logger.Logger
}

// NewBootstrapHandler creates a new BootstrapHandler. token guards the endpoint; pass ""
// to leave it unauthenticated.
func NewBootstrapHandler(bootstrapSvc core.BootstrapService, token string, log logger.Logger) *BootstrapHandler {
	return &BootstrapHandler{
		bootstrapSvc: bootstrapSvc,
		token:        token,
		log:          log,
	}
}

type bootstrapRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// Bootstrap provisions the initial system administrator.
// @Summary Bootstrap admin
// @Description One-time endpoint that creates the initial admin user with an HSM wallet. Returns the wallet address for on-chain role assignment. Returns 409 if any user already exists.
// @Tags admin
// @Accept json
// @Produce json
// @Param body body bootstrapRequest true "Admin email"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/bootstrap [post]
func (h *BootstrapHandler) Bootstrap(c *gin.Context) {
	if h.token != "" {
		presented := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		// Constant-time: a timing-comparable check would leak the token byte by byte.
		if subtle.ConstantTimeCompare([]byte(presented), []byte(h.token)) != 1 {
			h.log.Warn("Bootstrap rejected: bad or missing token", "remoteAddr", c.ClientIP())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid bootstrap token"})
			return
		}
	}

	var req bootstrapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required and must be valid"})
		return
	}

	address, err := h.bootstrapSvc.Bootstrap(c.Request.Context(), req.Email)
	if err != nil {
		HandleError(c, h.log, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"address": address})
}
