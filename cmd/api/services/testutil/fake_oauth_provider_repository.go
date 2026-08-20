package testutil

import (
	"context"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// FakeUserOAuthProviderRepository is an in-memory UserOAuthProviderRepository for unit tests.
type FakeUserOAuthProviderRepository struct {
	Providers []domain.UserOAuthProvider
	CreateErr error // if set, returned by Create instead of appending
}

func (r *FakeUserOAuthProviderRepository) Create(_ context.Context, provider *domain.UserOAuthProvider) error {
	if r.CreateErr != nil {
		return r.CreateErr
	}
	r.Providers = append(r.Providers, *provider)
	return nil
}

func (r *FakeUserOAuthProviderRepository) FindByProviderAndID(
	_ context.Context,
	provider domain.OAuthProvider,
	oauthID string,
) (*domain.UserOAuthProvider, error) {
	for i := range r.Providers {
		if r.Providers[i].Provider == provider && r.Providers[i].OAuthID == oauthID {
			return &r.Providers[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeUserOAuthProviderRepository) FindByProviderAndUserID(
	_ context.Context,
	provider domain.OAuthProvider,
	userID uuid.UUID,
) (*domain.UserOAuthProvider, error) {
	for i := range r.Providers {
		if r.Providers[i].Provider == provider && r.Providers[i].UserID == userID {
			return &r.Providers[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeUserOAuthProviderRepository) UpdateOAuthID(_ context.Context, id uuid.UUID, oauthID string) error {
	for i := range r.Providers {
		if r.Providers[i].ID == id {
			r.Providers[i].OAuthID = oauthID
			return nil
		}
	}
	return core.ErrRecordNotFound
}

func (r *FakeUserOAuthProviderRepository) CountByUserID(_ context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	for i := range r.Providers {
		if r.Providers[i].UserID == userID {
			count++
		}
	}
	return count, nil
}
