package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/sse"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// heartbeatInterval keeps the SSE connection alive through idle-timeout proxies.
const heartbeatInterval = 25 * time.Second

// TokenStreamHandler streams curated token lifecycle events to authenticated clients via SSE.
type TokenStreamHandler struct {
	hub *sse.Hub
	log logger.Logger
}

// NewTokenStreamHandler creates a new TokenStreamHandler.
func NewTokenStreamHandler(hub *sse.Hub, log logger.Logger) *TokenStreamHandler {
	return &TokenStreamHandler{hub: hub, log: log}
}

// Stream opens a Server-Sent Events stream of token lifecycle updates.
// @Summary Stream token updates (SSE)
// @Description Server-Sent Events stream of curated token lifecycle events. Each message is
// @Description `event: token` with a JSON `data` payload `{type,address,status,...}`. Authenticated
// @Description via the session cookie — connect with `new EventSource(url, { withCredentials: true })`.
// @Tags tokens
// @Produce text/event-stream
// @Success 200 {string} string "event stream"
// @Failure 401 {object} map[string]string
// @Router /api/tokens/stream [get]
func (h *TokenStreamHandler) Stream(c *gin.Context) {
	if _, ok := middleware.GetAuthUser(c); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

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
			c.SSEvent("token", json.RawMessage(msg))
			return true
		case <-ticker.C:
			// SSE comment line — ignored by EventSource but keeps the connection warm.
			_, _ = fmt.Fprint(w, ": ping\n\n")
			return true
		}
	})
}
