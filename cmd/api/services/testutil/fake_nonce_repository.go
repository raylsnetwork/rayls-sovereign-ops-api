package testutil

import (
	"context"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// FakeNonceRepository is an in-memory NonceRepository for unit tests.
type FakeNonceRepository struct {
	Nonce   *domain.Nonce // returned by FindValidAndMarkUsed when non-nil
	FindErr error         // returned by FindValidAndMarkUsed when non-nil
}

func (r *FakeNonceRepository) Create(_ context.Context, _ *domain.Nonce) error {
	return nil
}

func (r *FakeNonceRepository) FindValidAndMarkUsed(_ context.Context, _, _ string) (*domain.Nonce, error) {
	if r.FindErr != nil {
		return nil, r.FindErr
	}
	if r.Nonce != nil {
		return r.Nonce, nil
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeNonceRepository) DeleteExpired(_ context.Context) error {
	return nil
}
