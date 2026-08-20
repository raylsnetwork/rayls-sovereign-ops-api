package testutil

import (
	"context"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// FakeUserSignupDetailsRepository is an in-memory UserSignupDetailsRepository for unit tests.
type FakeUserSignupDetailsRepository struct {
	Details   []domain.UserSignupDetails
	UpsertErr error // if set, returned by Upsert instead of storing
}

var _ core.UserSignupDetailsRepository = (*FakeUserSignupDetailsRepository)(nil)

func (r *FakeUserSignupDetailsRepository) Upsert(_ context.Context, details *domain.UserSignupDetails) error {
	if r.UpsertErr != nil {
		return r.UpsertErr
	}
	if details.ID == uuid.Nil {
		details.ID = uuid.New()
	}
	for i := range r.Details {
		if r.Details[i].UserID == details.UserID {
			r.Details[i] = *details
			return nil
		}
	}
	r.Details = append(r.Details, *details)
	return nil
}

func (r *FakeUserSignupDetailsRepository) FindByUserID(
	_ context.Context,
	userID uuid.UUID,
) (*domain.UserSignupDetails, error) {
	for i := range r.Details {
		if r.Details[i].UserID == userID {
			return &r.Details[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}
