package testutil

import (
	"bytes"
	"context"
	"encoding/hex"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// FakeUserRepository is an in-memory UserRepository for unit tests.
type FakeUserRepository struct {
	Users     []domain.User
	CreateErr error // if set, returned by Create instead of appending
}

// NewFakeUserRepository creates an empty FakeUserRepository.
func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{}
}

func (r *FakeUserRepository) FindByID(_ context.Context, id uuid.UUID) (*domain.User, error) {
	for i := range r.Users {
		if r.Users[i].ID == id {
			return &r.Users[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeUserRepository) Create(_ context.Context, user *domain.User) error {
	if r.CreateErr != nil {
		return r.CreateErr
	}
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	r.Users = append(r.Users, *user)
	return nil
}

func (r *FakeUserRepository) Update(_ context.Context, user *domain.User) error {
	for i := range r.Users {
		if r.Users[i].ID == user.ID {
			r.Users[i] = *user
			return nil
		}
	}
	return core.ErrRecordNotFound
}

func (r *FakeUserRepository) FindByEmail(_ context.Context, email string) (*domain.User, error) {
	for i := range r.Users {
		if r.Users[i].Email == email {
			return &r.Users[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeUserRepository) Count(_ context.Context) (int64, error) {
	return int64(len(r.Users)), nil
}

// SetOnChainUserID mirrors the GORM impl: a no-op when the user is absent or already set (the WHERE
// ... IS NULL guard matches no row), never an error.
func (r *FakeUserRepository) SetOnChainUserID(_ context.Context, userID uuid.UUID, onChainUserID []byte) error {
	for i := range r.Users {
		if r.Users[i].ID == userID && r.Users[i].OnChainUserID == nil {
			r.Users[i].OnChainUserID = onChainUserID
		}
	}
	return nil
}

// FindByOnChainUserIDs mirrors the GORM impl: returns a hex-keyed map of the users whose
// OnChainUserID matches any of the input hashes. Unmatched hashes are simply absent.
func (r *FakeUserRepository) FindByOnChainUserIDs(_ context.Context, hashes [][]byte) (map[string]uuid.UUID, error) {
	result := make(map[string]uuid.UUID, len(hashes))
	for _, h := range hashes {
		for i := range r.Users {
			if r.Users[i].OnChainUserID != nil && bytes.Equal(r.Users[i].OnChainUserID, h) {
				result[hex.EncodeToString(h)] = r.Users[i].ID
			}
		}
	}
	return result, nil
}
