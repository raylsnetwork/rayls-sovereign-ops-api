package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/sse"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// WalletBalanceStreamHandler streams curated wallet-balance events to authenticated
// clients via SSE, filtered to the caller's own wallet.
type WalletBalanceStreamHandler struct {
	hub        *sse.Hub
	walletRepo core.UserWalletRepository
	log        logger.Logger
}

func NewWalletBalanceStreamHandler(
	hub *sse.Hub,
	walletRepo core.UserWalletRepository,
	log logger.Logger,
) *WalletBalanceStreamHandler {
	return &WalletBalanceStreamHandler{hub: hub, walletRepo: walletRepo, log: log}
}

// streamedBalanceEvent is the subset of the WalletBalanceEvent payload used for filtering.
type streamedBalanceEvent struct {
	WalletAddress string `json:"walletAddress"`
}

// Stream opens a Server-Sent Events stream of wallet balance updates scoped to the
// authenticated user's wallet.
// @Summary Stream wallet balance updates (SSE)
// @Description Server-Sent Events stream of curated wallet balance events. Each message is
// @Description `event: wallet_balance` with a JSON `data` payload `{type,walletAddress,tokenAddress,balance,blockNumber}`.
// @Description Authenticated via the session cookie — connect with `new EventSource(url, { withCredentials: true })`.
// @Tags wallets
// @Produce text/event-stream
// @Success 200 {string} string "event stream"
// @Failure 401 {object} map[string]string
// @Router /api/wallets/balances/stream [get]
func (h *WalletBalanceStreamHandler) Stream(c *gin.Context) {
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

	wallet, err := h.walletRepo.FindByUserID(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found for user"})
			return
		}
		h.log.Error("Failed to load wallet for SSE filter", "userID", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load wallet"})
		return
	}

	allowed := domain.NormalizeAddress(wallet.RaylsAddress)

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")

	ch := h.hub.Register()
	defer h.hub.Unregister(ch)

	ticker := time.NewTicker(heartbeatInterval)
	defer ticker.Stop()

	ctx := c.Request.Context()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-ctx.Done():
			return false
		case msg, ok := <-ch:
			if !ok {
				return false
			}
			var evt streamedBalanceEvent
			if unmarshalErr := json.Unmarshal(msg, &evt); unmarshalErr != nil {
				h.log.Warn("Skipping unparseable wallet balance SSE event", "error", unmarshalErr)
				return true
			}
			if !strings.EqualFold(evt.WalletAddress, allowed) {
				return true
			}
			c.SSEvent("wallet_balance", json.RawMessage(msg))
			return true
		case <-ticker.C:
			_, _ = fmt.Fprint(w, ": ping\n\n")
			return true
		}
	})
}
