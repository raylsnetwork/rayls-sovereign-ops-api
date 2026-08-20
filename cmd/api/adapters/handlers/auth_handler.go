package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/auth"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// AuthHandler handles token refresh and logout endpoints.
type AuthHandler struct {
	authSvc       core.AuthService
	tokenWrapper  *auth.TokenWrapper
	blacklistRepo core.TokenBlacklistRepository
	walletRepo    core.UserWalletRepository
	userRepo      core.UserRepository
	ramClient     core.RaylsAccessManagerClient
	log           logger.Logger
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(
	authSvc core.AuthService,
	tokenWrapper *auth.TokenWrapper,
	blacklistRepo core.TokenBlacklistRepository,
	walletRepo core.UserWalletRepository,
	userRepo core.UserRepository,
	ramClient core.RaylsAccessManagerClient,
	log logger.Logger,
) *AuthHandler {
	return &AuthHandler{
		authSvc:       authSvc,
		tokenWrapper:  tokenWrapper,
		blacklistRepo: blacklistRepo,
		walletRepo:    walletRepo,
		userRepo:      userRepo,
		ramClient:     ramClient,
		log:           log,
	}
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Refresh validates a refresh token and issues a new JWT + rotated refresh token.
// @Summary Refresh token
// @Description Validates the refresh token, checks user eligibility, rotates the token pair.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body refreshRequest true "Refresh token"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req refreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "INVALID_REFRESH_TOKEN",
			"message": "Your session has expired. Please log in again.",
		})
		return
	}

	userID, jti, err := h.tokenWrapper.ParseRefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "INVALID_REFRESH_TOKEN",
			"message": "Your session has expired. Please log in again.",
		})
		return
	}

	revoked, err := h.blacklistRepo.ExistsByJTI(c.Request.Context(), jti)
	if err != nil {
		h.log.Error("Failed to check token blacklist", "jti", jti, "error", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "An unexpected error occurred."},
		)
		return
	}
	if revoked {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "TOKEN_REVOKED",
			"message": "Your session was terminated. Please log in again.",
		})
		return
	}

	parsedID, parseErr := uuid.Parse(userID)
	if parseErr != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"code": "INVALID_REFRESH_TOKEN", "message": "Your session has expired. Please log in again."},
		)
		return
	}

	user, findErr := h.userRepo.FindByID(c.Request.Context(), parsedID)
	if findErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "INVALID_USER_STATUS",
			"message": "Your account has been suspended. Contact your operator.",
		})
		return
	}

	if !user.IsActive || user.Status != domain.UserStatusRoleAssigned {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "INVALID_USER_STATUS",
			"message": "Your account has been suspended. Contact your operator.",
		})
		return
	}

	// The shared identity service has no chain and no HSM wallets: roles live in each
	// chain's AccessManager, so a cross-chain token must not assert any. Re-issue an
	// identity-only token and let the per-chain ops-api authorize each request against
	// its own am_* tables. Mirrors the chain-less branch of applyLoginDecisionTree.
	if h.ramClient == nil || h.walletRepo == nil {
		h.issueRefreshed(c, user, nil, jti, userID)
		return
	}

	wallet, walletErr := h.walletRepo.FindByUserID(c.Request.Context(), user.ID)
	if walletErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "NO_ACTIVE_ROLES",
			"message": "Your access has been revoked. Contact your operator.",
		})
		return
	}

	roles, rolesErr := h.ramClient.GetRoles(c.Request.Context(), wallet.RaylsAddress)
	if rolesErr != nil {
		h.log.Error("GetRoles RPC failed at refresh", "userID", user.ID, "error", rolesErr)
		c.Header("Retry-After", "30")
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    "SERVICE_UNAVAILABLE",
			"message": "Service temporarily unavailable. Please try again shortly.",
		})
		return
	}
	if len(roles) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "NO_ACTIVE_ROLES",
			"message": "Your access has been revoked. Contact your operator.",
		})
		return
	}

	h.issueRefreshed(c, user, roles, jti, userID)
}

// issueRefreshed rotates the refresh token and re-issues the session JWT. roles may be
// nil — the identity service mints identity-only tokens (see the chain-less branch above).
func (h *AuthHandler) issueRefreshed(
	c *gin.Context,
	user *domain.User,
	roles []string,
	jti, userID string,
) {
	// Blacklist the old token before issuing a new one (rotation).
	_ = h.blacklistRepo.Create(c.Request.Context(), &domain.TokenBlacklist{
		JTI:       jti,
		ExpiresAt: time.Now().UTC().Add(25 * time.Hour),
	})

	newRefreshToken, _, issueErr := h.tokenWrapper.CreateRefreshTokenString(userID)
	if issueErr != nil {
		h.log.Error("Failed to create refresh token", "error", issueErr)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "An unexpected error occurred."},
		)
		return
	}

	if tokenErr := h.tokenWrapper.IssueToken(c.Request.Context(), c.Writer, user, "refresh", roles); tokenErr != nil {
		h.log.Error("Failed to issue JWT", "error", tokenErr)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "An unexpected error occurred."},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"refresh_token": newRefreshToken})
}

// Logout blacklists the refresh token and clears the JWT cookie.
// @Summary Logout
// @Description Revokes the refresh token and clears the auth cookie.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body refreshRequest true "Refresh token"
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req refreshRequest
	if err := c.ShouldBindJSON(&req); err == nil && req.RefreshToken != "" {
		_, jti, parseErr := h.tokenWrapper.ParseRefreshToken(req.RefreshToken)
		if parseErr == nil {
			_ = h.blacklistRepo.Create(c.Request.Context(), &domain.TokenBlacklist{
				JTI:       jti,
				ExpiresAt: time.Now().UTC().Add(25 * time.Hour),
			})
		}
	}

	h.tokenWrapper.ClearCookie(c.Writer)
	c.JSON(http.StatusOK, gin.H{"status": "logged out"})
}
