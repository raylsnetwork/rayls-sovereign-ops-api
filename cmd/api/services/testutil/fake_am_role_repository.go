package testutil

import (
	"context"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

var _ core.AccessManagerRoleRepository = (*FakeAccessManagerRoleRepository)(nil)

// FakeAccessManagerRoleRepository is an in-memory AccessManagerRoleRepository for unit tests.
type FakeAccessManagerRoleRepository struct {
	Roles   []*domain.AccessManagerRole
	ListErr error // if set, returned by List instead of the stored roles
}

func (r *FakeAccessManagerRoleRepository) Upsert(_ context.Context, role *domain.AccessManagerRole) error {
	for i := range r.Roles {
		if r.Roles[i].RoleID == role.RoleID {
			r.Roles[i] = role
			return nil
		}
	}
	r.Roles = append(r.Roles, role)
	return nil
}

func (r *FakeAccessManagerRoleRepository) FindByID(
	_ context.Context,
	roleID uint64,
) (*domain.AccessManagerRole, error) {
	for i := range r.Roles {
		if r.Roles[i].RoleID == roleID {
			return r.Roles[i], nil
		}
	}
	return nil, core.ErrRecordNotFound
}

func (r *FakeAccessManagerRoleRepository) List(_ context.Context) ([]*domain.AccessManagerRole, error) {
	if r.ListErr != nil {
		return nil, r.ListErr
	}
	return r.Roles, nil
}
