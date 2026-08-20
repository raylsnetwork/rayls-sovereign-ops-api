package testutil

import (
	"context"

	"github.com/google/uuid"
)

// FakeRoleService is an in-memory role service for use in tests.
type FakeRoleService struct {
	// Role is returned by GetUserRole. Empty string means no roles assigned.
	Role string
	// Err, when non-nil, is returned instead of Role.
	Err error
}

func NewFakeRoleService(role string) *FakeRoleService { return &FakeRoleService{Role: role} }

func (f *FakeRoleService) GetUserRole(_ context.Context, _ uuid.UUID, _ string) (string, error) {
	if f.Err != nil {
		return "", f.Err
	}
	return f.Role, nil
}
