package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// TokenPermissionHandler exposes what the authenticated user can do on a token contract.
type TokenPermissionHandler struct {
	svc        core.TokenPermissionService
	walletRepo core.UserWalletRepository
	log        logger.Logger
}

// NewTokenPermissionHandler creates a new TokenPermissionHandler.
func NewTokenPermissionHandler(
	svc core.TokenPermissionService,
	walletRepo core.UserWalletRepository,
	log logger.Logger,
) *TokenPermissionHandler {
	return &TokenPermissionHandler{svc: svc, walletRepo: walletRepo, log: log}
}

// Get returns the functions the authenticated user's wallet can call on the given token.
// @Summary Get token permissions for the current user
// @Description Returns, for the authenticated user's custody wallet, which Access Manager-gated
// @Description functions (mint/burn/...) it can call on the token, plus convenience flags.
// @Tags tokens
// @Produce json
// @Param address path string true "Token contract address (0x...)"
// @Success 200 {object} core.TokenPermissions
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tokens/{address}/permissions [get]
func (h *TokenPermissionHandler) Get(c *gin.Context) {
	jwtUser, ok := middleware.GetAuthUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	userID, err := uuid.Parse(jwtUser.ID)
	if err != nil {
		h.log.Error("Invalid user ID in JWT", "id", jwtUser.ID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user session"})
		return
	}

	wallet, err := h.walletRepo.GetSignerWalletForChain(c.Request.Context(), userID, domain.WalletChainPrivate)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no custody wallet for user"})
			return
		}
		h.log.Error("Failed to fetch wallet for user", "userID", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch custody wallet"})
		return
	}

	perms, err := h.svc.GetTokenPermissions(c.Request.Context(), c.Param("address"), wallet.RaylsAddress)
	if err != nil {
		h.log.Error("Failed to compute token permissions", "address", c.Param("address"), "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to compute token permissions"})
		return
	}

	c.JSON(http.StatusOK, perms)
}
