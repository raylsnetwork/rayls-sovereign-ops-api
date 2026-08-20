package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/utils"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// TokenHandler handles token listing endpoints.
type TokenHandler struct {
	tokenRepo core.TokenRepository
	log       logger.Logger
}

// NewTokenHandler creates a new TokenHandler.
func NewTokenHandler(tokenRepo core.TokenRepository, log logger.Logger) *TokenHandler {
	return &TokenHandler{tokenRepo: tokenRepo, log: log}
}

type tokenResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	ContractAddress  string `json:"contractAddress"`
	ErcStandard      uint8  `json:"ercStandard"`
	ErcStandardLabel string `json:"ercStandardLabel"`
	TokenClass       string `json:"tokenClass"`
	Status           uint8  `json:"status"`
	StatusLabel      string `json:"statusLabel"`
	TotalSupply      string `json:"totalSupply"`
	HolderCount      int    `json:"holderCount"`
	Decimals         uint8  `json:"decimals"`
	MetadataURL      string `json:"metadataUrl,omitempty"`
	IssuerID         string `json:"issuerId,omitempty"`
}

// tokenListResponse is used only for Swagger type generation.
type tokenListResponse = utils.PagedResponse[tokenResponse]

func tokenToResponse(t *domain.Token) tokenResponse {
	return tokenResponse{
		ID:               t.ID.String(),
		Name:             t.Name,
		Symbol:           t.Symbol,
		ContractAddress:  t.ContractAddress,
		ErcStandard:      uint8(t.ErcStandard),
		ErcStandardLabel: t.ErcStandard.Label(),
		TokenClass:       t.TokenClass,
		Status:           uint8(t.Status),
		StatusLabel:      t.Status.Label(),
		TotalSupply:      t.TotalSupply,
		HolderCount:      t.HolderCount,
		Decimals:         t.Decimals,
		MetadataURL:      t.MetadataURL,
		IssuerID:         t.IssuerID,
	}
}

// List returns a paginated list of tokens with optional filters.
// @Summary List tokens
// @Description Returns a paginated list of tokens discovered by the indexer. Supports filtering by name, symbol and ERC type.
// @Tags tokens
// @Produce json
// @Param page   query int    false "Page number (default 1)"
// @Param limit  query int    false "Items per page (default 20, max 100)"
// @Param name     query string false "Filter by token name (case-insensitive, partial match)"
// @Param symbol   query string false "Filter by token symbol (case-insensitive, partial match)"
// @Param type     query string false "Filter by ERC standard (canonical labels)" Enums(RAYLS_ERC20, RAYLS_ERC721, RAYLS_ERC1155, RAYLS_ENYGMA, RAYLS_ERC721_DVP, RAYLS_ERC1155_DVP, RAYLS_STABLECOIN)
// @Param issuerId query string false "Filter by the chain the token was deployed on (chainId)"
// @Success 200 {object} tokenListResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tokens [get]
func (h *TokenHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}

	filter := core.TokenFilter{
		Name:      c.Query("name"),
		Symbol:    c.Query("symbol"),
		TokenType: c.Query("type"),
		IssuerID:  c.Query("issuerId"),
		Page:      page,
		Limit:     limit,
	}

	tokens, total, err := h.tokenRepo.List(c.Request.Context(), filter)
	if err != nil {
		h.log.Error("Failed to list tokens", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list tokens"})
		return
	}

	items := make([]tokenResponse, len(tokens))
	for i, t := range tokens {
		items[i] = tokenToResponse(t)
	}

	c.JSON(http.StatusOK, utils.NewPagedResponse(items, page, limit, total))
}

// Details returns a single token by its contract address.
// @Summary Get token details
// @Description Returns the token identified by its contract address (case-insensitive).
// @Tags tokens
// @Produce json
// @Param address path string true "Token contract address (0x...)"
// @Success 200 {object} tokenResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tokens/{address} [get]
func (h *TokenHandler) Details(c *gin.Context) {
	address := c.Param("address")

	token, err := h.tokenRepo.FindByAddress(c.Request.Context(), address)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "token not found"})
			return
		}
		h.log.Error("Failed to fetch token", "address", address, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch token"})
		return
	}

	c.JSON(http.StatusOK, tokenToResponse(token))
}
