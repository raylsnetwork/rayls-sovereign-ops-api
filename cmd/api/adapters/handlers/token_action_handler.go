package handlers

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/utils"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// TokenActionHandler executes mint/burn on a token, signed by the authenticated user's custody
// wallet, after verifying the user is allowed (via the Access Manager).
type TokenActionHandler struct {
	svc         core.TokenActionService
	teleportSvc core.TeleportService
	perms       core.TokenPermissionService
	tokenRepo   core.TokenRepository
	walletRepo  core.UserWalletRepository
	log         logger.Logger
}

func NewTokenActionHandler(
	svc core.TokenActionService,
	teleportSvc core.TeleportService,
	perms core.TokenPermissionService,
	tokenRepo core.TokenRepository,
	walletRepo core.UserWalletRepository,
	log logger.Logger,
) *TokenActionHandler {
	return &TokenActionHandler{
		svc:         svc,
		teleportSvc: teleportSvc,
		perms:       perms,
		tokenRepo:   tokenRepo,
		walletRepo:  walletRepo,
		log:         log,
	}
}

type mintRequest struct {
	To      string `json:"to"`
	Amount  string `json:"amount"`
	TokenID string `json:"tokenId"`
	Data    string `json:"data"`
}

// pauseRequest is the body of POST /api/tokens/:address/pause. Paused is a pointer so an
// omitted field is distinguishable from an explicit false — the handler rejects the former
// rather than silently treating a malformed body as "unpause".
type pauseRequest struct {
	Paused *bool `json:"paused"`
}

type burnRequest struct {
	From    string `json:"from"`
	Amount  string `json:"amount"`
	TokenID string `json:"tokenId"`
}

type teleportRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   string `json:"amount"`
	TokenID  string `json:"tokenId"`
	Data     string `json:"data"`
	Standard uint8  `json:"standard"`
}

type txResponse struct {
	TxHash string `json:"txHash"`
}

// teleportResponse is the teleport-only success body. It keeps the snake_case tx_hash field name of
// the legacy backend's tokenLock response so existing e2e clients need no response change, while
// mint/burn keep the camelCase txResponse used by the rest of the token API.
type teleportResponse struct {
	TxHash string `json:"tx_hash"`
}

// Mint mints tokens on the given contract.
// @Summary Mint tokens
// @Description Mints tokens on the token contract, signed with the authenticated user's custody wallet.
// @Description Requires the user's wallet to hold the mint permission (Access Manager). Body fields depend
// @Description on the token standard: fungible uses {to,amount}; ERC721 uses {to,tokenId}; ERC1155 uses {to,tokenId,amount,data}.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Token contract address"
// @Param request body mintRequest true "Mint parameters"
// @Success 200 {object} txResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Router /api/tokens/{address}/mint [post]
func (h *TokenActionHandler) Mint(c *gin.Context) {
	signer, token, ok := h.authorize(c, func(p *core.TokenPermissions) bool { return p.CanMint }, "mint")
	if !ok {
		return
	}

	var req mintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request body", Hint: err.Error()})
		return
	}

	in, hint := buildMintInput(token, req)
	if hint != "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid mint parameters", Hint: hint})
		return
	}

	txHash, err := h.svc.Mint(c.Request.Context(), signer, token.ContractAddress, token.ErcStandard, in)
	h.respond(c, txHash, err, "mint")
}

// Burn burns tokens on the given contract.
// @Summary Burn tokens
// @Description Burns tokens on the token contract, signed with the authenticated user's custody wallet.
// @Description Requires the burn permission. Body fields depend on standard: fungible uses {from,amount};
// @Description ERC721 uses {tokenId}; ERC1155 uses {from,tokenId,amount}.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Token contract address"
// @Param request body burnRequest true "Burn parameters"
// @Success 200 {object} txResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Router /api/tokens/{address}/burn [post]
func (h *TokenActionHandler) Burn(c *gin.Context) {
	signer, token, ok := h.authorize(c, func(p *core.TokenPermissions) bool { return p.CanBurn }, "burn")
	if !ok {
		return
	}

	var req burnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request body", Hint: err.Error()})
		return
	}

	in, hint := buildBurnInput(token, req)
	if hint != "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid burn parameters", Hint: hint})
		return
	}

	txHash, err := h.svc.Burn(c.Request.Context(), signer, token.ContractAddress, token.ErcStandard, in)
	h.respond(c, txHash, err, "burn")
}

// Pause halts or resumes all transfers, mints and burns on a stablecoin.
// @Summary Pause or unpause a stablecoin
// @Description Calls pause() or unpause() on the token contract, signed with the authenticated
// @Description user's custody wallet. Stablecoin only (RAYLS_STABLECOIN): other standards have no
// @Description pause function. Authorization is NOT an Access Manager role — the contract accepts
// @Description only its own `pauser` address, so the caller's wallet must equal it.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Token contract address"
// @Param request body pauseRequest true "Desired paused state"
// @Success 200 {object} txResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 409 {object} utils.ErrorResponse
// @Failure 422 {object} utils.ErrorResponse
// @Router /api/tokens/{address}/pause [post]
func (h *TokenActionHandler) Pause(c *gin.Context) {
	signer, token, ok := h.resolveSignerAndToken(c)
	if !ok {
		return
	}

	if token.ErcStandard != domain.ErcStandardStableCoin {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error: "pause is not supported for this token standard",
			Hint:  "only RAYLS_STABLECOIN exposes pause()/unpause()",
		})
		return
	}

	var req pauseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request body", Hint: err.Error()})
		return
	}
	if req.Paused == nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Error: "paused is required",
			Hint:  `send {"paused":true} to pause or {"paused":false} to resume`,
		})
		return
	}
	want := *req.Paused

	// onlyPauser is a msg.sender equality check against the contract's `pauser`, so authorize
	// by reading it rather than consulting the Access Manager. Doing this here turns what would
	// be an opaque on-chain revert into a 403 that names the reason.
	pauser, err := h.svc.Pauser(c.Request.Context(), token.ContractAddress)
	if err != nil {
		h.log.Error("Failed to read pauser", "address", token.ContractAddress, "error", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to read the token's pauser"})
		return
	}
	if !strings.EqualFold(pauser, signer) {
		c.JSON(http.StatusForbidden, gin.H{"error": "wallet is not the pauser for this token"})
		return
	}

	// Already in the requested state: pause() would revert (or silently no-op) and cost gas for
	// nothing. 409 says "your intent is understood, there is nothing to do".
	current, err := h.svc.IsPaused(c.Request.Context(), token.ContractAddress)
	if err != nil {
		h.log.Error("Failed to read paused state", "address", token.ContractAddress, "error", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to read the token's paused state"})
		return
	}
	if current == want {
		state := "unpaused"
		if current {
			state = "paused"
		}
		c.JSON(http.StatusConflict, utils.ErrorResponse{
			Error: "token is already " + state,
		})
		return
	}

	action := "unpause"
	if want {
		action = "pause"
	}
	txHash, err := h.svc.SetPaused(c.Request.Context(), signer, token.ContractAddress, want)
	h.respond(c, txHash, err, action)
}

// Teleport moves an asset from the privacy chain to the public chain.
// @Summary Teleport tokens to the public chain
// @Description Calls teleportToPublicChain on the token contract, signed with the authenticated user's
// @Description custody wallet, after a mandatory preflight (token registered & active, caller balance/ownership).
// @Description Supports only ERC20, ERC721 and ERC1155. Body fields depend on the standard: ERC20 uses
// @Description {to,amount}; ERC721 uses {to,tokenId}; ERC1155 uses {to,tokenId,amount,data}. amount is in
// @Description base units. {from} is required and selects which of the caller's wallets signs (msg.sender);
// @Description it must be one of the caller's active private-chain custody wallets.
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Token contract address"
// @Param request body teleportRequest true "Teleport parameters"
// @Success 200 {object} teleportResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} map[string]string
// @Failure 422 {object} utils.ErrorResponse
// @Router /api/tokens/{address}/teleport [post]
func (h *TokenActionHandler) Teleport(c *gin.Context) {
	if h.teleportSvc == nil {
		c.JSON(http.StatusNotImplemented, utils.ErrorResponse{Error: "teleport is not enabled on this deployment"})
		return
	}

	var req teleportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	std := domain.ErcStandard(req.Standard)
	if std != domain.ErcStandardERC20 && std != domain.ErcStandardERC721 && std != domain.ErcStandardERC1155 {
		c.JSON(
			http.StatusBadRequest,
			utils.ErrorResponse{
				Error: "invalid token standard (expected 1=ERC20,2=ERC721,3=ERC1155)",
				Hint:  "teleport supports only ERC20 (1), ERC721 (2) and ERC1155 (3)",
			},
		)
		return
	}

	in, detail := buildTeleportInput(std, req)
	if detail != "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: detail})
		return
	}

	// The caller chooses which of their own wallets signs (msg.sender). Identity is anchored to the
	// JWT — `from` may only select among the caller's own active private-chain custody wallets.
	ok := h.verifyTeleportSigner(c, in.From)
	if !ok {
		return
	}

	tokenAddress := c.Param("address")
	txHash, err := h.teleportSvc.Teleport(c.Request.Context(), tokenAddress, std, in)
	if err != nil {
		if RespondIfReverted(c, h.log, err) {
			return
		}
		HandleError(c, h.log, err)
		return
	}
	c.JSON(http.StatusOK, teleportResponse{TxHash: txHash})
}

// verifyTeleportSigner confirms that `from` is one of the authenticated user's active HSM signer
// wallets and is on the private chain (teleport source). The caller picks the wallet; we only
// confirm it is theirs and on the right chain. Writes the error response and returns ok=false on failure.
func (h *TokenActionHandler) verifyTeleportSigner(c *gin.Context, from string) bool {
	jwtUser, ok := middleware.GetAuthUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return false
	}
	userID, err := uuid.Parse(jwtUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user session"})
		return false
	}

	wallet, err := h.walletRepo.GetSignerWalletByAddress(c.Request.Context(), userID, from)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			// "not found" and "not yours" are intentionally indistinguishable — don't reveal others' wallets.
			c.JSON(http.StatusBadRequest, gin.H{"error": "from must be one of your custody wallets"})
			return false
		}
		h.log.Error("Failed to fetch wallet", "address", from, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch custody wallet"})
		return false
	}

	// Teleport's source is the privacy chain, so the signer must be a private-chain wallet.
	if wallet.Chain != domain.WalletChainPrivate {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from must be a private-chain custody wallet"})
		return false
	}

	return true
}

// authorize resolves the user's wallet and the token, and checks the AM permission via allowed.
// Returns the signer address and token on success; writes the error response and returns ok=false otherwise.
func (h *TokenActionHandler) authorize(
	c *gin.Context,
	allowed func(*core.TokenPermissions) bool,
	action string,
) (string, *domain.Token, bool) {
	jwtUser, ok := middleware.GetAuthUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return "", nil, false
	}
	userID, err := uuid.Parse(jwtUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user session"})
		return "", nil, false
	}

	wallet, err := h.walletRepo.GetSignerWalletForChain(c.Request.Context(), userID, domain.WalletChainPrivate)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no custody wallet for user"})
			return "", nil, false
		}
		h.log.Error("Failed to fetch wallet", "userID", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch custody wallet"})
		return "", nil, false
	}

	address := c.Param("address")
	token, err := h.tokenRepo.FindByAddress(c.Request.Context(), address)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "token not found"})
			return "", nil, false
		}
		h.log.Error("Failed to fetch token", "address", address, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch token"})
		return "", nil, false
	}

	perms, err := h.perms.GetTokenPermissions(c.Request.Context(), token.ContractAddress, wallet.RaylsAddress)
	if err != nil {
		h.log.Error("Failed to check token permissions", "address", token.ContractAddress, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check permissions"})
		return "", nil, false
	}
	if !allowed(perms) {
		c.JSON(http.StatusForbidden, gin.H{"error": "wallet is not allowed to " + action + " this token"})
		return "", nil, false
	}

	return wallet.RaylsAddress, token, true
}

// resolveSignerAndToken is authorize() without the Access Manager check: it identifies the
// caller's signing wallet and the token being addressed, and nothing more.
//
// Split out for pause/unpause, which the stablecoin gates on its own `pauser` address rather
// than an AM function permission — so there is no TokenPermissions predicate to apply.
func (h *TokenActionHandler) resolveSignerAndToken(c *gin.Context) (string, *domain.Token, bool) {
	jwtUser, ok := middleware.GetAuthUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return "", nil, false
	}
	userID, err := uuid.Parse(jwtUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user session"})
		return "", nil, false
	}

	wallet, err := h.walletRepo.GetSignerWalletForChain(c.Request.Context(), userID, domain.WalletChainPrivate)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no custody wallet for user"})
			return "", nil, false
		}
		h.log.Error("Failed to fetch wallet", "userID", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch custody wallet"})
		return "", nil, false
	}

	address := c.Param("address")
	token, err := h.tokenRepo.FindByAddress(c.Request.Context(), address)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "token not found"})
			return "", nil, false
		}
		h.log.Error("Failed to fetch token", "address", address, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch token"})
		return "", nil, false
	}

	return wallet.RaylsAddress, token, true
}

func (h *TokenActionHandler) respond(c *gin.Context, txHash string, err error, action string) {
	if err != nil {
		h.log.Error("Token "+action+" failed", "error", err)
		if errors.Is(err, core.ErrTxReverted) {
			c.JSON(
				http.StatusUnprocessableEntity,
				utils.ErrorResponse{Error: "transaction reverted on-chain", Hint: err.Error()},
			)
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to " + action + " on-chain"})
		return
	}
	c.JSON(http.StatusOK, txResponse{TxHash: txHash})
}

// isFungible reports whether the standard uses a scaled fungible amount (ERC20/Enygma/StableCoin).
// The stablecoin inherits the ERC20 handler, so its mint/burn take a decimal-scaled amount.
func isFungible(std domain.ErcStandard) bool {
	return std == domain.ErcStandardERC20 ||
		std == domain.ErcStandardEnygma ||
		std == domain.ErcStandardStableCoin
}

func is1155(std domain.ErcStandard) bool {
	return std == domain.ErcStandardERC1155 || std == domain.ErcStandardZkDvpERC1155
}

func is721(std domain.ErcStandard) bool {
	return std == domain.ErcStandardERC721 || std == domain.ErcStandardZkDvpERC721
}

func buildMintInput(token *domain.Token, req mintRequest) (core.MintInput, string) {
	if !common.IsHexAddress(req.To) {
		return core.MintInput{}, "to must be a valid address"
	}
	in := core.MintInput{To: req.To}

	switch {
	case isFungible(token.ErcStandard):
		amount, err := scaleDecimal(req.Amount, token.Decimals)
		if err != nil {
			return core.MintInput{}, err.Error()
		}
		in.Amount = amount
	case is721(token.ErcStandard):
		id, err := parseRawUint(req.TokenID)
		if err != nil {
			return core.MintInput{}, "tokenId: " + err.Error()
		}
		in.TokenID = id
	case is1155(token.ErcStandard):
		id, err := parseRawUint(req.TokenID)
		if err != nil {
			return core.MintInput{}, "tokenId: " + err.Error()
		}
		amount, err := scaleDecimal(req.Amount, token.Decimals)
		if err != nil {
			return core.MintInput{}, err.Error()
		}
		data, err := parseHexBytes(req.Data)
		if err != nil {
			return core.MintInput{}, "data: " + err.Error()
		}
		in.TokenID, in.Amount, in.Data = id, amount, data
	default:
		return core.MintInput{}, "unsupported token standard"
	}
	return in, ""
}

func buildBurnInput(token *domain.Token, req burnRequest) (core.BurnInput, string) {
	in := core.BurnInput{From: req.From}

	switch {
	case isFungible(token.ErcStandard):
		if !common.IsHexAddress(req.From) {
			return core.BurnInput{}, "from must be a valid address"
		}
		amount, err := scaleDecimal(req.Amount, token.Decimals)
		if err != nil {
			return core.BurnInput{}, err.Error()
		}
		in.Amount = amount
	case is721(token.ErcStandard):
		id, err := parseRawUint(req.TokenID)
		if err != nil {
			return core.BurnInput{}, "tokenId: " + err.Error()
		}
		in.TokenID = id
	case is1155(token.ErcStandard):
		if !common.IsHexAddress(req.From) {
			return core.BurnInput{}, "from must be a valid address"
		}
		id, err := parseRawUint(req.TokenID)
		if err != nil {
			return core.BurnInput{}, "tokenId: " + err.Error()
		}
		amount, err := scaleDecimal(req.Amount, token.Decimals)
		if err != nil {
			return core.BurnInput{}, err.Error()
		}
		in.TokenID, in.Amount = id, amount
	default:
		return core.BurnInput{}, "unsupported token standard"
	}
	return in, ""
}

// teleportRequired reports the legacy backend's presence-validation message when a required field is
// empty, matching the message text the e2e suite asserts on. Returns "" when the field is present.
func teleportRequired(value, field string, req teleportRequest, std domain.ErcStandard) string {
	if strings.TrimSpace(value) == "" {
		return fmt.Sprintf("%s is required for standard %d (%s)", field, req.Standard, std.Label())
	}
	return ""
}

// buildTeleportInput validates and parses the teleport request per standard. Unlike mint/burn, amounts
// are raw base units (no decimal scaling) — teleport does not consult the local token row. From is the
// caller's validated signer wallet; To is the public-chain destination. Missing required fields yield
// the legacy backend's message text (amount for ERC20/ERC1155, tokenId for ERC721/ERC1155).
func buildTeleportInput(std domain.ErcStandard, req teleportRequest) (core.TeleportInput, string) {
	if !common.IsHexAddress(req.To) {
		return core.TeleportInput{}, "to must be a valid address"
	}
	if !common.IsHexAddress(req.From) {
		return core.TeleportInput{}, "from must be a valid address"
	}
	in := core.TeleportInput{From: req.From, To: req.To}

	switch std {
	case domain.ErcStandardERC20:
		if msg := teleportRequired(req.Amount, "amount", req, std); msg != "" {
			return core.TeleportInput{}, msg
		}
		amount, err := parseRawUint(req.Amount)
		if err != nil {
			return core.TeleportInput{}, "amount: " + err.Error()
		}
		in.Amount = amount
	case domain.ErcStandardERC721:
		if msg := teleportRequired(req.TokenID, "tokenId", req, std); msg != "" {
			return core.TeleportInput{}, msg
		}
		id, err := parseRawUint(req.TokenID)
		if err != nil {
			return core.TeleportInput{}, "tokenId: " + err.Error()
		}
		in.TokenID = id
	case domain.ErcStandardERC1155:
		if msg := teleportRequired(req.TokenID, "tokenId", req, std); msg != "" {
			return core.TeleportInput{}, msg
		}
		if msg := teleportRequired(req.Amount, "amount", req, std); msg != "" {
			return core.TeleportInput{}, msg
		}
		id, err := parseRawUint(req.TokenID)
		if err != nil {
			return core.TeleportInput{}, "tokenId: " + err.Error()
		}
		amount, err := parseRawUint(req.Amount)
		if err != nil {
			return core.TeleportInput{}, "amount: " + err.Error()
		}
		data, err := parseHexBytes(req.Data)
		if err != nil {
			return core.TeleportInput{}, "invalid data: must be hex string"
		}
		in.TokenID, in.Amount, in.Data = id, amount, data
	default:
		return core.TeleportInput{}, "unsupported token standard"
	}
	return in, ""
}

// scaleDecimal converts a human decimal amount (e.g. "1.5") into base units (amount * 10^decimals)
// using integer math only. Rejects negatives and more fractional digits than decimals allows.
func scaleDecimal(amount string, decimals uint8) (*big.Int, error) {
	amount = strings.TrimSpace(amount)
	if amount == "" {
		return nil, errors.New("amount is required")
	}
	if strings.HasPrefix(amount, "-") {
		return nil, errors.New("amount must not be negative")
	}
	intPart, fracPart, _ := strings.Cut(amount, ".")
	if len(fracPart) > int(decimals) {
		return nil, errors.New("amount has more decimal places than the token allows")
	}
	if intPart == "" {
		intPart = "0"
	}
	digits := intPart + fracPart + strings.Repeat("0", int(decimals)-len(fracPart))
	v, ok := new(big.Int).SetString(digits, 10)
	if !ok {
		return nil, errors.New("amount is not a valid number")
	}
	return v, nil
}

// parseRawUint parses a base-10 non-negative integer (token id / raw uint256).
func parseRawUint(s string) (*big.Int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("value is required")
	}
	v, ok := new(big.Int).SetString(s, 10)
	if !ok || v.Sign() < 0 {
		return nil, errors.New("must be a non-negative integer")
	}
	return v, nil
}

// parseHexBytes decodes optional 0x-prefixed hex into bytes ("" → empty).
func parseHexBytes(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return []byte{}, nil
	}
	if !strings.HasPrefix(s, "0x") && !strings.HasPrefix(s, "0X") {
		return nil, errors.New("must be 0x-prefixed hex")
	}
	b := common.FromHex(s)
	return b, nil
}
