package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/utils"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// TokenDeployHandler deploys a protocol token via RNContractFactory, signed by the
// authenticated user's custody wallet.
type TokenDeployHandler struct {
	svc        core.TokenDeployService
	walletRepo core.UserWalletRepository
	tokenRepo  core.TokenRepository
	// registry auto-registers deployed tokens; nil when the TokenRegistry is not wired.
	registry core.TokenRegistryService
	log      logger.Logger
}

// NewTokenDeployHandler creates a new TokenDeployHandler.
func NewTokenDeployHandler(
	svc core.TokenDeployService,
	walletRepo core.UserWalletRepository,
	tokenRepo core.TokenRepository,
	log logger.Logger,
) *TokenDeployHandler {
	return &TokenDeployHandler{svc: svc, walletRepo: walletRepo, tokenRepo: tokenRepo, log: log}
}

// SetTokenRegistry injects the registry after construction (built later in the container).
func (h *TokenDeployHandler) SetTokenRegistry(registry core.TokenRegistryService) {
	h.registry = registry
}

type deployTokenRequest struct {
	Standard string `json:"standard" binding:"required"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	URI      string `json:"uri"`
	Decimals uint8  `json:"decimals"`
}

type deployTokenResponse struct {
	DeployedAddress string `json:"deployedAddress"`
	TxHash          string `json:"txHash"`
	Standard        string `json:"standard"`
}

// Deploy deploys a single protocol token through the factory.
// @Summary Deploy a protocol token
// @Description Deploys a token via RNContractFactory, signed with the authenticated user's custody wallet.
// @Tags tokens
// @Accept json
// @Produce json
// @Param request body deployTokenRequest true "Token to deploy"
// @Success 201 {object} deployTokenResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Failure 502 {object} map[string]string
// @Router /api/tokens [post]
func (h *TokenDeployHandler) Deploy(c *gin.Context) {
	wallet, spec, reqStandard, ok := h.resolveDeployContext(c)
	if !ok {
		return
	}

	deployedAddr, txHash, err := h.svc.Deploy(c.Request.Context(), wallet.RaylsAddress, spec)
	if err != nil {
		h.log.Error("Token deploy failed", "standard", reqStandard, "error", err)
		if errors.Is(err, core.ErrTxReverted) {
			c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Error: "token deploy reverted on-chain",
				Hint:  "the custody wallet may lack the FACTORY_DEPLOYER role required to deploy via the factory",
			})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to deploy token on-chain"})
		return
	}

	// Canonicalize so the response, the DB row, and the SSE events all use the same address form.
	deployedAddr = domain.NormalizeAddress(deployedAddr)

	if regErr := h.registerAndAuthorize(c.Request.Context(), deployedAddr); regErr != nil {
		h.log.Error("Token register/authorize failed after deploy", "address", deployedAddr, "error", regErr)
		h.recordToken(c, spec, deployedAddr) // still track it so it is not lost
		c.JSON(http.StatusBadGateway, utils.ErrorResponse{
			Error: "token deployed but on-chain registration failed",
			Hint:  "deployed at " + deployedAddr + "; it must be registered and authorized before mint/transfer",
		})
		return
	}

	h.recordToken(c, spec, deployedAddr)

	c.JSON(http.StatusCreated, deployTokenResponse{
		DeployedAddress: deployedAddr,
		TxHash:          txHash,
		Standard:        reqStandard,
	})
}

// registerAndAuthorize runs the two operator steps the protocol requires before a token can
// mint/transfer (register -> WAITING_APPROVAL, then authorize -> AUTHORIZED). No-op when unwired.
func (h *TokenDeployHandler) registerAndAuthorize(ctx context.Context, tokenAddress string) error {
	if h.registry == nil {
		return nil
	}
	if _, err := h.registry.Register(ctx, core.RegisterTokenInput{TokenAddress: tokenAddress}); err != nil {
		return fmt.Errorf("register token: %w", err)
	}
	if _, err := h.registry.SetStatus(ctx, tokenAddress, domain.PrivacyNodeStatusAuthorized); err != nil {
		return fmt.Errorf("authorize token: %w", err)
	}
	return nil
}

// Estimate returns the real on-chain gas estimate for deploying a token, without executing it.
// @Summary Estimate token deploy gas
// @Description Returns the real gas estimate (eth_estimateGas + gas price) for deploying a token via RNContractFactory, signed by the authenticated user's custody wallet. Does not deploy.
// @Tags tokens
// @Accept json
// @Produce json
// @Param request body deployTokenRequest true "Token to estimate"
// @Success 200 {object} core.TokenDeployEstimate
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 502 {object} map[string]string
// @Router /api/tokens/estimate [post]
func (h *TokenDeployHandler) Estimate(c *gin.Context) {
	wallet, spec, reqStandard, ok := h.resolveDeployContext(c)
	if !ok {
		return
	}

	estimate, err := h.svc.EstimateDeploy(c.Request.Context(), wallet.RaylsAddress, spec)
	if err != nil {
		h.log.Error("Token deploy estimate failed", "standard", reqStandard, "error", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to estimate token deploy gas"})
		return
	}

	c.JSON(http.StatusOK, estimate)
}

// resolveDeployContext runs the shared auth + request-parsing pipeline for the deploy and
// estimate endpoints: it authenticates the user, binds and validates the request, and loads
// the caller's custody wallet. On any failure it writes the error response and returns ok=false.
// The returned reqStandard is the client's original standard label, echoed back verbatim.
func (h *TokenDeployHandler) resolveDeployContext(
	c *gin.Context,
) (wallet *domain.UserWallet, spec core.TokenDeploySpec, reqStandard string, ok bool) {
	jwtUser, ok := middleware.GetAuthUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return nil, core.TokenDeploySpec{}, "", false
	}

	userID, err := uuid.Parse(jwtUser.ID)
	if err != nil {
		h.log.Error("Invalid user ID in JWT", "id", jwtUser.ID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user session"})
		return nil, core.TokenDeploySpec{}, "", false
	}

	var req deployTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request body", Hint: err.Error()})
		return nil, core.TokenDeploySpec{}, "", false
	}

	standard, ok := parseErcStandard(req.Standard)
	if !ok {
		c.JSON(
			http.StatusBadRequest,
			utils.NewInvalidEnumValueError("standard", req.Standard, strings.Join(supportedStandards(), ", ")),
		)
		return nil, core.TokenDeploySpec{}, "", false
	}

	spec = core.TokenDeploySpec{
		ErcStandard: standard,
		Name:        req.Name,
		Symbol:      req.Symbol,
		URI:         req.URI,
		Decimals:    req.Decimals,
	}
	if hint := validateSpec(spec); hint != "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid token parameters", Hint: hint})
		return nil, core.TokenDeploySpec{}, "", false
	}

	wallet, err = h.walletRepo.FindByUserID(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no custody wallet for user"})
			return nil, core.TokenDeploySpec{}, "", false
		}
		h.log.Error("Failed to fetch wallet for user", "userID", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch custody wallet"})
		return nil, core.TokenDeploySpec{}, "", false
	}

	return wallet, spec, req.Standard, true
}

// recordToken best-effort persists the freshly deployed token so it shows up in GET /api/tokens.
// A persistence failure does not fail the request — the on-chain deploy already succeeded.
func (h *TokenDeployHandler) recordToken(c *gin.Context, spec core.TokenDeploySpec, address string) {
	token := &domain.Token{
		Name:   spec.Name,
		Symbol: spec.Symbol,
		// ResourceID left nil — not assigned at deploy time (deploy passes bytes32(0)).
		MetadataURL: spec.URI,
		ErcStandard: spec.ErcStandard,
		Decimals:    spec.Decimals,
		// IssuerID records the chain (PN) the token was deployed on — the deployer's current chain id.
		IssuerID: h.svc.ChainID(),
		// Internal: tokens deployed via the API stay internal (not promoted to Active by the indexer).
		Status:          domain.TokenStatusInternal,
		ContractAddress: address,
	}
	if err := h.tokenRepo.Upsert(c.Request.Context(), token); err != nil {
		h.log.Warn("Deployed token persisted failed (deploy succeeded on-chain)", "address", address, "error", err)
	}
}

// parseErcStandard maps a request standard label to the domain ErcStandard.
func parseErcStandard(s string) (domain.ErcStandard, bool) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "ERC20", "ERC-20":
		return domain.ErcStandardERC20, true
	case "ENYGMA":
		return domain.ErcStandardEnygma, true
	case "ERC721", "ERC-721":
		return domain.ErcStandardERC721, true
	case "ERC1155", "ERC-1155":
		return domain.ErcStandardERC1155, true
	case "ERC721_DVP", "ERC721DVP":
		return domain.ErcStandardZkDvpERC721, true
	case "ERC1155_DVP", "ERC1155DVP":
		return domain.ErcStandardZkDvpERC1155, true
	case "STABLECOIN", "STABLE_COIN", "RAYLS_STABLECOIN":
		return domain.ErcStandardStableCoin, true
	default:
		return 0, false
	}
}

func supportedStandards() []string {
	return []string{"ERC20", "ENYGMA", "ERC721", "ERC1155", "ERC721_DVP", "ERC1155_DVP", "STABLECOIN"}
}

// validateSpec returns a non-empty hint describing the first missing required field for the
// token's standard, or "" when the spec is valid. ResourceID presence is enforced by binding.
func validateSpec(spec core.TokenDeploySpec) string {
	switch spec.ErcStandard {
	case domain.ErcStandardERC20, domain.ErcStandardEnygma, domain.ErcStandardStableCoin:
		if spec.Name == "" || spec.Symbol == "" {
			return "name and symbol are required for ERC20/Enygma/StableCoin"
		}
	case domain.ErcStandardERC721, domain.ErcStandardZkDvpERC721:
		if spec.URI == "" || spec.Name == "" || spec.Symbol == "" {
			return "uri, name and symbol are required for ERC721 standards"
		}
	case domain.ErcStandardERC1155, domain.ErcStandardZkDvpERC1155:
		if spec.URI == "" || spec.Name == "" {
			return "uri and name are required for ERC1155 standards"
		}
	}
	return ""
}
