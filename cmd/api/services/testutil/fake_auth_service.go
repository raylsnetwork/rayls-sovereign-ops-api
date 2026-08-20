package testutil

import (
	"context"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// FakeAuthService is a stub AuthService for unit tests. Each method returns the
// pre-set result; unset error fields mean success.
type FakeAuthService struct {
	User  *domain.User
	Roles []string
	Err   error

	Message string
	Nonce   string
}

var _ core.AuthService = (*FakeAuthService)(nil)

func (s *FakeAuthService) GenerateChallenge(_ context.Context, _ string) (string, string, error) {
	return s.Message, s.Nonce, s.Err
}

func (s *FakeAuthService) VerifySIWE(
	_ context.Context,
	_, _, _ string,
) (*domain.User, []string, error) {
	if s.Err != nil {
		return nil, nil, s.Err
	}
	return s.User, s.Roles, nil
}

func (s *FakeAuthService) FindOrCreateOAuthUser(
	_ context.Context,
	_ domain.OAuthProvider,
	_, _, _ string,
	_ bool,
) (*domain.User, []string, error) {
	if s.Err != nil {
		return nil, nil, s.Err
	}
	return s.User, s.Roles, nil
}
