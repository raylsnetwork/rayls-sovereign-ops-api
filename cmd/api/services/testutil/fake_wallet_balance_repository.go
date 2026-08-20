package testutil

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// FakeWalletBalanceRepository is an in-memory WalletBalanceRepository for unit tests.
// Upsert enforces the same block_number guard as the real repository: the incoming
// balance only overwrites the stored one when its block_number is greater than or
// equal to the stored one.
type FakeWalletBalanceRepository struct {
	Balances []domain.WalletBalance
}

func (r *FakeWalletBalanceRepository) Upsert(_ context.Context, balance *domain.WalletBalance) error {
	walletAddr := domain.NormalizeAddress(balance.WalletAddress)
	tokenAddr := domain.NormalizeAddress(balance.TokenAddress)
	for i := range r.Balances {
		if strings.EqualFold(r.Balances[i].WalletAddress, walletAddr) &&
			strings.EqualFold(r.Balances[i].TokenAddress, tokenAddr) {
			if balance.BlockNumber >= r.Balances[i].BlockNumber {
				r.Balances[i].Balance = balance.Balance
				r.Balances[i].BlockNumber = balance.BlockNumber
			}
			return nil
		}
	}
	if balance.ID == uuid.Nil {
		balance.ID = uuid.New()
	}
	copy := *balance
	copy.WalletAddress = walletAddr
	copy.TokenAddress = tokenAddr
	r.Balances = append(r.Balances, copy)
	return nil
}

func (r *FakeWalletBalanceRepository) ListByWallet(
	_ context.Context,
	walletAddress string,
) ([]*domain.WalletBalance, error) {
	addr := domain.NormalizeAddress(walletAddress)
	var out []*domain.WalletBalance
	for i := range r.Balances {
		if strings.EqualFold(r.Balances[i].WalletAddress, addr) {
			b := r.Balances[i]
			out = append(out, &b)
		}
	}
	return out, nil
}

func (r *FakeWalletBalanceRepository) GetByWalletAndToken(
	_ context.Context,
	walletAddress, tokenAddress string,
) (*domain.WalletBalance, error) {
	wa := domain.NormalizeAddress(walletAddress)
	ta := domain.NormalizeAddress(tokenAddress)
	for i := range r.Balances {
		if strings.EqualFold(r.Balances[i].WalletAddress, wa) &&
			strings.EqualFold(r.Balances[i].TokenAddress, ta) {
			b := r.Balances[i]
			return &b, nil
		}
	}
	return nil, core.ErrRecordNotFound
}
