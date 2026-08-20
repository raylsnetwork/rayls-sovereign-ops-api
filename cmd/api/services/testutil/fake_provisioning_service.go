package testutil

import (
	"context"

	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// FakeProvisioningService simulates provisioning: it records calls and, unless it
// returns an error, advances the user to role_assigned like the real service does.
type FakeProvisioningService struct {
	Err   error
	Calls []*domain.User
}

func NewFakeProvisioningService() *FakeProvisioningService { return &FakeProvisioningService{} }

func (f *FakeProvisioningService) Provision(_ context.Context, user *domain.User) error {
	f.Calls = append(f.Calls, user)
	if f.Err != nil {
		return f.Err
	}
	user.Status = domain.UserStatusRoleAssigned
	return nil
}
