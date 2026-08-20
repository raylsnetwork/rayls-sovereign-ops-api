package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// WalletBalanceHandler exposes per-wallet token balances synced from Blockscout.
type WalletBalanceHandler struct {
	svc *services.WalletBalanceService
	log logger.Logger
}

func NewWalletBalanceHandler(svc *services.WalletBalanceService, log logger.Logger) *WalletBalanceHandler {
	return &WalletBalanceHandler{svc: svc, log: log}
}

// walletBalanceResponse is the per-token balance entry returned by List.
type walletBalanceResponse struct {
	WalletAddress string `json:"walletAddress"`
	TokenAddress  string `json:"tokenAddress"`
	TokenSymbol   string `json:"tokenSymbol"`
	TokenName     string `json:"tokenName"`
	Decimals      uint8  `json:"decimals"`
	Balance       string `json:"balance"`
	BlockNumber   uint64 `json:"blockNumber"`
	UpdatedAt     string `json:"updatedAt"`
}

// List returns every token balance synced for the given wallet.
// @Summary List wallet balances
// @Description Returns the current per-token balances for a wallet, sourced from Blockscout via the balances indexer.
// @Tags wallets
// @Produce json
// @Param address path string true "Wallet address (0x...)"
// @Success 200 {array} walletBalanceResponse
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/wallets/{address}/balances [get]
func (h *WalletBalanceHandler) List(c *gin.Context) {
	address := c.Param("address")

	views, err := h.svc.ListForWallet(c.Request.Context(), address)
	if err != nil {
		HandleError(c, h.log, err)
		return
	}

	out := make([]walletBalanceResponse, len(views))
	for i, v := range views {
		out[i] = walletBalanceResponse{
			WalletAddress: v.WalletAddress,
			TokenAddress:  v.TokenAddress,
			TokenSymbol:   v.TokenSymbol,
			TokenName:     v.TokenName,
			Decimals:      v.Decimals,
			Balance:       v.Balance,
			BlockNumber:   v.BlockNumber,
			UpdatedAt:     v.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, out)
}

// Details returns the balance for a specific (wallet, token) pair.
// @Summary Get a single wallet balance
// @Description Returns the current balance for a single token in a wallet, sourced from Blockscout via the balances indexer.
// @Tags wallets
// @Produce json
// @Param address      path string true "Wallet address (0x...)"
// @Param tokenAddress path string true "Token contract address (0x...)"
// @Success 200 {object} walletBalanceResponse
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/wallets/{address}/balances/{tokenAddress} [get]
func (h *WalletBalanceHandler) Details(c *gin.Context) {
	address := c.Param("address")
	tokenAddress := c.Param("tokenAddress")

	view, err := h.svc.GetForWalletAndToken(c.Request.Context(), address, tokenAddress)
	if err != nil {
		HandleError(c, h.log, err)
		return
	}

	c.JSON(http.StatusOK, walletBalanceResponse{
		WalletAddress: view.WalletAddress,
		TokenAddress:  view.TokenAddress,
		TokenSymbol:   view.TokenSymbol,
		TokenName:     view.TokenName,
		Decimals:      view.Decimals,
		Balance:       view.Balance,
		BlockNumber:   view.BlockNumber,
		UpdatedAt:     view.UpdatedAt,
	})
}
