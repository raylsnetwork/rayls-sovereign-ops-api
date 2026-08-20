package testutil

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// FakeTokenRepository is an in-memory TokenRepository for unit tests.
type FakeTokenRepository struct {
	Tokens []domain.Token
}

func (r *FakeTokenRepository) Upsert(_ context.Context, token *domain.Token) error {
	addr := domain.NormalizeAddress(token.ContractAddress)
	for i := range r.Tokens {
		if strings.EqualFold(r.Tokens[i].ContractAddress, addr) {
			r.Tokens[i] = *token
			r.Tokens[i].ContractAddress = addr
			return nil
		}
	}
	if token.ID == uuid.Nil {
		token.ID = uuid.New()
	}
	copy := *token
	copy.ContractAddress = addr
	r.Tokens = append(r.Tokens, copy)
	return nil
}

func (r *FakeTokenRepository) UpdateSupplyAndHolders(
	_ context.Context,
	address, totalSupply string,
	holderCount int,
) error {
	addr := domain.NormalizeAddress(address)
	for i := range r.Tokens {
		if strings.EqualFold(r.Tokens[i].ContractAddress, addr) {
			r.Tokens[i].TotalSupply = totalSupply
			r.Tokens[i].HolderCount = holderCount
			return nil
		}
	}
	return core.ErrRecordNotFound
}

func (r *FakeTokenRepository) FindByAddress(_ context.Context, address string) (*domain.Token, error) {
	addr := domain.NormalizeAddress(address)
	for i := range r.Tokens {
		if strings.EqualFold(r.Tokens[i].ContractAddress, addr) {
			t := r.Tokens[i]
			return &t, nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeTokenRepository) List(_ context.Context, _ core.TokenFilter) ([]*domain.Token, int64, error) {
	out := make([]*domain.Token, len(r.Tokens))
	for i := range r.Tokens {
		t := r.Tokens[i]
		out[i] = &t
	}
	return out, int64(len(r.Tokens)), nil
}
