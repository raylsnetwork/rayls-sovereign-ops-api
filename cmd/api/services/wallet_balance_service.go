package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// WalletBalanceView enriches a stored balance with the token's display fields so
// frontends can render holdings without a second round trip.
type WalletBalanceView struct {
	WalletAddress string `json:"walletAddress"`
	TokenAddress  string `json:"tokenAddress"`
	TokenSymbol   string `json:"tokenSymbol"`
	TokenName     string `json:"tokenName"`
	Decimals      uint8  `json:"decimals"`
	Balance       string `json:"balance"`
	BlockNumber   uint64 `json:"blockNumber"`
	UpdatedAt     string `json:"updatedAt"`
}

// WalletBalanceService reads per-wallet balances and joins them with token
// metadata for API responses.
type WalletBalanceService struct {
	balances core.WalletBalanceRepository
	wallets  core.UserWalletRepository
	tokens   core.TokenRepository
	log      logger.Logger
}

func NewWalletBalanceService(
	balances core.WalletBalanceRepository,
	wallets core.UserWalletRepository,
	tokens core.TokenRepository,
	log logger.Logger,
) *WalletBalanceService {
	return &WalletBalanceService{balances: balances, wallets: wallets, tokens: tokens, log: log}
}

// ListForWallet returns the current balances for a wallet, enriched with token
// metadata. Returns core.ErrWalletNotFound if no user_wallets row exists for the
// given address.
func (s *WalletBalanceService) ListForWallet(ctx context.Context, walletAddress string) ([]WalletBalanceView, error) {
	addr := domain.NormalizeAddress(walletAddress)

	if _, err := s.wallets.FindByRaylsAddress(ctx, addr); err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrWalletNotFound
		}
		return nil, fmt.Errorf("lookup wallet: %w", err)
	}

	balances, err := s.balances.ListByWallet(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("list balances: %w", err)
	}

	out := make([]WalletBalanceView, 0, len(balances))
	for _, b := range balances {
		out = append(out, s.enrich(ctx, b))
	}
	return out, nil
}

// GetForWalletAndToken returns the stored balance for a specific (wallet, token)
// pair, enriched with token metadata. Returns core.ErrWalletNotFound when the
// wallet is unknown and a typed *core.NotFoundError (which HandleError maps to
// HTTP 404) when the wallet is known but holds no balance for the token.
func (s *WalletBalanceService) GetForWalletAndToken(
	ctx context.Context,
	walletAddress, tokenAddress string,
) (*WalletBalanceView, error) {
	wallet := domain.NormalizeAddress(walletAddress)
	token := domain.NormalizeAddress(tokenAddress)

	if _, err := s.wallets.FindByRaylsAddress(ctx, wallet); err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrWalletNotFound
		}
		return nil, fmt.Errorf("lookup wallet: %w", err)
	}

	balance, err := s.balances.GetByWalletAndToken(ctx, wallet, token)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("wallet balance", wallet+"/"+token)
		}
		return nil, fmt.Errorf("get balance: %w", err)
	}

	view := s.enrich(ctx, balance)
	return &view, nil
}

// enrich joins a stored balance with the token's display fields. Missing token
// metadata is non-fatal — the balance is returned with empty symbol/name/decimals.
func (s *WalletBalanceService) enrich(ctx context.Context, b *domain.WalletBalance) WalletBalanceView {
	view := WalletBalanceView{
		WalletAddress: b.WalletAddress,
		TokenAddress:  b.TokenAddress,
		Balance:       b.Balance,
		BlockNumber:   b.BlockNumber,
		UpdatedAt:     b.UpdatedAt.UTC().Format("2006-01-02T15:04:05.000000Z"),
	}
	if tok, tokErr := s.tokens.FindByAddress(ctx, b.TokenAddress); tokErr == nil && tok != nil {
		view.TokenSymbol = tok.Symbol
		view.TokenName = tok.Name
		view.Decimals = tok.Decimals
	} else if tokErr != nil && !errors.Is(tokErr, core.ErrRecordNotFound) {
		s.log.Warn("Failed to load token metadata for balance view",
			"token", b.TokenAddress, "error", tokErr)
	}
	return view
}
