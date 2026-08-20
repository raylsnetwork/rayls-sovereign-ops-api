package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// UserHandler handles user profile endpoints.
type UserHandler struct {
	userRepo   core.UserRepository
	walletRepo core.UserWalletRepository
	log        logger.Logger
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(
	userRepo core.UserRepository,
	walletRepo core.UserWalletRepository,
	log logger.Logger,
) *UserHandler {
	return &UserHandler{userRepo: userRepo, walletRepo: walletRepo, log: log}
}

type meResponse struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Picture       string   `json:"picture,omitempty"`
	Roles         []string `json:"roles"`
	WalletAddress string   `json:"wallet_address,omitempty"`
}

// Me returns the authenticated user's profile.
// @Summary Get current user
// @Description Returns profile information for the authenticated PRIVACY_NODE_OPERATOR.
// @Tags users
// @Produce json
// @Success 200 {object} meResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/me [get]
func (h *UserHandler) Me(c *gin.Context) {
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

	user, err := h.userRepo.FindByID(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}
		h.log.Error("Failed to fetch user", "userID", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return
	}

	var walletAddress string
	wallet, walletErr := h.walletRepo.FindByUserID(c.Request.Context(), userID)
	if walletErr != nil && !errors.Is(walletErr, core.ErrRecordNotFound) {
		h.log.Warn("Failed to fetch wallet for user", "userID", userID, "error", walletErr)
	} else if wallet != nil {
		walletAddress = wallet.RaylsAddress
	}

	raw, _ := jwtUser.Attributes["roles"].([]any)
	roles := make([]string, 0, len(raw))
	for _, r := range raw {
		if s, ok := r.(string); ok {
			roles = append(roles, s)
		}
	}

	c.JSON(http.StatusOK, meResponse{
		ID:            user.ID.String(),
		Name:          user.Name,
		Email:         user.Email,
		Picture:       jwtUser.Picture,
		Roles:         roles,
		WalletAddress: walletAddress,
	})
}
