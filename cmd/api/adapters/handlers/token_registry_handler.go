package handlers

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/utils"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// TokenRegistryHandler handles registration of already-deployed token contracts into the on-chain
// TokenRegistry catalog.
type TokenRegistryHandler struct {
	svc core.TokenRegistryService
	log logger.Logger
}

func NewTokenRegistryHandler(svc core.TokenRegistryService, log logger.Logger) *TokenRegistryHandler {
	return &TokenRegistryHandler{svc: svc, log: log}
}

// setTokenStatusRequest is the PATCH body. The token address comes from the path, never the body.
// Status is the string label ("authorized" / "unauthorized"), matching the "layer" and "target"
// discriminators on the freeze/submit endpoints rather than leaking the on-chain enum number.
type setTokenStatusRequest struct {
	Status string `json:"status"`
}

// freezeRequest is the POST body for the freeze/unfreeze endpoints. The token address comes from the
// path, never the body; only the target layer is supplied here.
type freezeRequest struct {
	Layer string `json:"layer"`
}

// submitRequest is the POST body for the submit endpoint. The token address comes from the path,
// never the body; only the target layer ("hub" or "public_chain") is supplied here.
type submitRequest struct {
	Target string `json:"target"`
}

// registeredTokenResponse mirrors the legacy rayls-privacy-backend domain.Token, so existing clients
// need no change: numeric standard/status, no additive label fields.
type registeredTokenResponse struct {
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	TokenAddress string    `json:"address"`
	URI          string    `json:"uri"`
	Standard     uint8     `json:"standard"`
	Status       uint8     `json:"status"`
	LastUpdated  time.Time `json:"updated_at"`
}

func toRegisteredTokenResponse(t *core.RegisteredToken) registeredTokenResponse {
	return registeredTokenResponse{
		Name:         t.Name,
		Symbol:       t.Symbol,
		TokenAddress: domain.ChecksumAddress(t.TokenAddress),
		URI:          t.URI,
		Standard:     uint8(t.Standard),
		Status:       uint8(t.Status),
		LastUpdated:  t.LastUpdated,
	}
}

// toRegisteredTokenResponses maps a slice of catalog entries to their response form. The slice is
// initialized so an empty catalog serializes as [] rather than null.
func toRegisteredTokenResponses(tokens []core.RegisteredToken) []registeredTokenResponse {
	items := make([]registeredTokenResponse, 0, len(tokens))
	for i := range tokens {
		items = append(items, toRegisteredTokenResponse(&tokens[i]))
	}
	return items
}

// Register adds an already-deployed token contract to the TokenRegistry catalog.
// @Summary Register a token
// @Description Adds an already-deployed token contract (at the path :address) to the on-chain
// @Description TokenRegistry catalog. Registration is address-only — the contract reads the token's
// @Description name/symbol/standard/supply on-chain, so no request body is needed. The contract must
// @Description have deployed code (EOAs and empty addresses are rejected). The on-chain write is
// @Description operator-signed; the token starts in WAITING_APPROVAL status (privacyNodeStatus 1).
// @Description This does not deploy a token — that remains POST /api/tokens.
// @Tags tokens
// @Produce json
// @Param address path string true "Deployed token contract address"
// @Success 201 {object} registeredTokenResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/tokens/{address}/register [post]
func (h *TokenRegistryHandler) Register(c *gin.Context) {
	if _, ok := middleware.GetAuthUser(c); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	address := c.Param("address")
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "address must be a valid hex address"})
		return
	}

	in := core.RegisterTokenInput{
		TokenAddress: domain.NormalizeAddress(address),
	}

	token, err := h.svc.Register(c.Request.Context(), in)
	if err != nil {
		// A revert means the contract rejected the registration (e.g. the token or symbol is
		// already registered) — a client-correctable condition, not a server fault. Surface it as
		// 422, consistent with the mint/burn and deploy handlers.
		if RespondIfReverted(c, h.log, err) {
			return
		}
		HandleError(c, h.log, err)
		return
	}

	c.JSON(http.StatusCreated, toRegisteredTokenResponse(token))
}

// List returns every token in the TokenRegistry catalog.
// @Summary List registered tokens
// @Description Returns all tokens in the on-chain TokenRegistry catalog. This is the registry
// @Description catalog (read from-chain), distinct from GET /api/tokens, which lists indexer-discovered
// @Description tokens. No query parameters are accepted.
// @Tags tokens
// @Produce json
// @Success 200 {array} registeredTokenResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tokens/registry [get]
func (h *TokenRegistryHandler) List(c *gin.Context) {
	if _, ok := middleware.GetAuthUser(c); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	tokens, err := h.svc.List(c.Request.Context())
	if err != nil {
		HandleError(c, h.log, err)
		return
	}

	c.JSON(http.StatusOK, toRegisteredTokenResponses(tokens))
}

// ListPending returns the pending (WAITING_APPROVAL) tokens in the TokenRegistry catalog.
// @Summary List pending registered tokens
// @Description Returns only the pending tokens in the on-chain TokenRegistry catalog — those in the
// @Description WAITING_APPROVAL status (privacyNodeStatus 1), i.e. registered but not yet approved by
// @Description an operator. No query parameters are accepted.
// @Tags tokens
// @Produce json
// @Success 200 {array} registeredTokenResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tokens/registry/pending [get]
func (h *TokenRegistryHandler) ListPending(c *gin.Context) {
	if _, ok := middleware.GetAuthUser(c); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	tokens, err := h.svc.ListByStatus(c.Request.Context(), domain.PrivacyNodeStatusWaitingApproval)
	if err != nil {
		HandleError(c, h.log, err)
		return
	}

	c.JSON(http.StatusOK, toRegisteredTokenResponses(tokens))
}

// SetStatus moves a registered token through the TokenRegistry privacyNodeStatus lifecycle (operator-only).
// @Summary Set a registered token's privacy-node status (operator)
// @Description Moves a registered token through the PN lifecycle by submitting an operator-signed
// @Description updatePrivacyNodeStatus to the TokenRegistry. The token address is taken from the path
// @Description :address, never the body. Allowed statuses are "authorized" and "unauthorized";
// @Description "undefined", "waiting_approval" (the initial state set by registration) and
// @Description "frozen" are rejected. Freezing/unfreezing has dedicated contract methods
// @Description (freezeOnPrivacyNode/unfreezeOnPrivacyNode) exposed via the freeze/unfreeze endpoints.
// @Description Requires the PRIVACY_NODE_OPERATOR role. Returns 200 with no body.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Registered token contract address"
// @Param request body setTokenStatusRequest true "New privacyNodeStatus (\"authorized\" or \"unauthorized\")"
// @Success 200 {object} nil
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/v1/admin/tokens/{address}/status [patch]
func (h *TokenRegistryHandler) SetStatus(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "address must be a valid hex address"})
		return
	}

	var req setTokenStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request body", Hint: err.Error()})
		return
	}
	// Freezing/unfreezing is intentionally excluded here — it has dedicated contract methods
	// (freezeOnPrivacyNode / unfreezeOnPrivacyNode / freezeOnPublicChain / unfreezeOnPublicChain),
	// exposed via the POST /api/v1/admin/tokens/{address}/freeze and .../unfreeze endpoints, rather
	// than routing FROZEN through updatePrivacyNodeStatus.
	status, ok := domain.ParseSettableStatus(req.Status)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error: `status must be "authorized" or "unauthorized"`,
			Hint:  `"waiting_approval" is set by registration; "frozen" has dedicated freeze/unfreeze endpoints`,
		})
		return
	}

	_, err := h.svc.SetStatus(
		c.Request.Context(),
		domain.NormalizeAddress(address),
		status,
	)
	if err != nil {
		// A revert means the contract rejected the status change (e.g. the token is not registered) —
		// a client-correctable condition, surfaced as 422 consistent with the register handler.
		if RespondIfReverted(c, h.log, err) {
			return
		}
		HandleError(c, h.log, err)
		return
	}

	c.Status(http.StatusOK)
}

// Freeze freezes a registered token at the given layer (operator-only).
// @Summary Freeze a registered token at a layer (operator)
// @Description Freezes a registered token by submitting an operator-signed freeze to the TokenRegistry.
// @Description The token address is taken from the path :address, never the body. The layer is one of
// @Description "privacy_node" (freezeOnPrivacyNode — blocks all operations) or "public_chain"
// @Description (freezeOnPublicChain — blocks public chain operations only). The Hub (PNH) layer is not
// @Description supported here — hub freezing is a cross-chain operation owned by the Private Hub.
// @Description Requires the PRIVACY_NODE_OPERATOR role. Returns 200 with no body.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Registered token contract address"
// @Param request body freezeRequest true "Target freeze layer (privacy_node or public_chain)"
// @Success 200 {object} nil
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/v1/admin/tokens/{address}/freeze [post]
func (h *TokenRegistryHandler) Freeze(c *gin.Context) {
	h.handleFreeze(c, true)
}

// Unfreeze unfreezes a registered token at the given layer (operator-only).
// @Summary Unfreeze a registered token at a layer (operator)
// @Description Unfreezes a registered token by submitting an operator-signed unfreeze to the
// @Description TokenRegistry. The token address is taken from the path :address, never the body. The
// @Description layer is one of "privacy_node" (unfreezeOnPrivacyNode) or "public_chain"
// @Description (unfreezeOnPublicChain). The Hub (PNH) layer is not supported here. Requires the
// @Description PRIVACY_NODE_OPERATOR role. Returns 200 with no body.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Registered token contract address"
// @Param request body freezeRequest true "Target freeze layer (privacy_node or public_chain)"
// @Success 200 {object} nil
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/v1/admin/tokens/{address}/unfreeze [post]
func (h *TokenRegistryHandler) Unfreeze(c *gin.Context) {
	h.handleFreeze(c, false)
}

// handleFreeze is the shared body of Freeze/Unfreeze: it validates the path address and request
// layer, then submits the operator-signed freeze/unfreeze. The layer must be a supported,
// address-based layer (privacy_node or public_chain); the Hub (PNH) layer is rejected as 400.
func (h *TokenRegistryHandler) handleFreeze(c *gin.Context, frozen bool) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "address must be a valid hex address"})
		return
	}

	var req freezeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request body", Hint: err.Error()})
		return
	}

	layer, ok := domain.ParseFreezeLayer(req.Layer)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error: `layer must be "privacy_node" or "public_chain"; the hub layer is not supported`,
		})
		return
	}

	normalized := domain.NormalizeAddress(address)
	var err error
	if frozen {
		_, err = h.svc.Freeze(c.Request.Context(), normalized, layer)
	} else {
		_, err = h.svc.Unfreeze(c.Request.Context(), normalized, layer)
	}
	if err != nil {
		// A revert means the contract rejected the freeze (e.g. the token is not registered, or the
		// operator wallet lacks the on-chain authorization for this layer) — a client-correctable
		// condition, surfaced as 422 consistent with the SetStatus handler.
		if RespondIfReverted(c, h.log, err) {
			return
		}
		HandleError(c, h.log, err)
		return
	}

	c.Status(http.StatusOK)
}

// Submit submits an AUTHORIZED token to the Hub or the Public Chain (operator-only).
// @Summary Submit a registered token to the Hub or Public Chain (operator)
// @Description Submits an AUTHORIZED token to another layer by submitting an operator-signed
// @Description submitToHub / submitToPublicChain to the TokenRegistry. The token address is taken from
// @Description the path :address, never the body. The target is one of "hub" (submitToHub — sends
// @Description addToken() to the Private Hub and moves hubStatus to WAITING_APPROVAL) or "public_chain"
// @Description (submitToPublicChain — moves publicChainStatus to PENDING_DEPLOYMENT). Submitting only
// @Description initiates the flow; activation on the Hub and deployment on the Public Chain complete
// @Description later via cross-chain PNH / relayer callbacks. Requires privacyNodeStatus == AUTHORIZED
// @Description (enforced on-chain; a token not yet authorized reverts as 422) and the
// @Description PRIVACY_NODE_OPERATOR role. Returns 200 with no body.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Registered token contract address"
// @Param request body submitRequest true "Target layer (hub or public_chain)"
// @Success 200 {object} nil
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/v1/admin/tokens/{address}/submit [post]
func (h *TokenRegistryHandler) Submit(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "address must be a valid hex address"})
		return
	}

	var req submitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request body", Hint: err.Error()})
		return
	}

	target, ok := domain.ParseSubmitTarget(req.Target)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error: `target must be "hub" or "public_chain"`,
		})
		return
	}

	if _, err := h.svc.Submit(c.Request.Context(), domain.NormalizeAddress(address), target); err != nil {
		// A revert means the contract rejected the submit (e.g. the token is not registered / not yet
		// AUTHORIZED on the PN, or the operator wallet lacks authorization) — a client-correctable
		// condition, surfaced as 422 consistent with the SetStatus handler.
		if RespondIfReverted(c, h.log, err) {
			return
		}
		HandleError(c, h.log, err)
		return
	}

	c.Status(http.StatusOK)
}
