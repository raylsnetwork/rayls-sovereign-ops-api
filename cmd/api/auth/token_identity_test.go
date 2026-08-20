package auth_test

import (
	"context"
	"testing"
	"time"

	authlib "github.com/go-pkgz/auth/v2"
	"github.com/go-pkgz/auth/v2/avatar"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/auth"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

const testJWTSecret = "test-secret"

func newTokenService(t *testing.T) *token.Service {
	t.Helper()
	svc := authlib.NewService(authlib.Opts{
		SecretReader: token.SecretFunc(func(_ string) (string, error) { return testJWTSecret, nil }),
		Issuer:       "rayls-ops-api",
		DisableXSRF:  true,
		AvatarStore:  avatar.NewNoOp(),
	})
	return svc.TokenService()
}

func buildUser() *domain.User {
	return &domain.User{
		Model:    domain.Model{ID: uuid.New()},
		Name:     "Alice",
		Email:    "alice@example.com",
		IsActive: true,
		Status:   domain.UserStatusRoleAssigned,
	}
}

// parseAttrs mints a token and returns the attributes the verifier sees.
func parseAttrs(t *testing.T, w *auth.TokenWrapper, user *domain.User, roles []string) map[string]any {
	t.Helper()
	str, err := w.MintToken(context.Background(), user, "google", roles, time.Minute)
	require.NoError(t, err)

	claims, err := newTokenService(t).Parse(str)
	require.NoError(t, err)
	require.NotNil(t, claims.User)
	return claims.User.Attributes
}

func TestTokenWrapper_MintToken_IdentityServiceOmitsRoles(t *testing.T) {
	// The shared identity service must not assert roles: they are granted per chain.
	wrapper := auth.NewIdentityTokenWrapper(newTokenService(t), testJWTSecret, nil, "")

	attrs := parseAttrs(t, wrapper, buildUser(), nil)

	assert.NotContains(t, attrs, "roles")
	assert.Equal(t, "google", attrs["auth_method"])
}

func TestTokenWrapper_MintToken_IdentityServiceCarriesCustodyWallet(t *testing.T) {
	// One custody wallet per user (an EVM keypair works on every chain), so identity owns
	// it and puts it in the token — unlike roles, it is not a per-chain fact.
	user := buildUser()
	walletRepo := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{{
		UserID:          user.ID,
		RaylsAddress:    "0xabc",
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPrivate,
		IsActive:        true,
	}}}
	wrapper := auth.NewIdentityTokenWrapper(newTokenService(t), testJWTSecret, walletRepo, "")

	attrs := parseAttrs(t, wrapper, user, nil)

	assert.Equal(t, "0xabc", attrs["custody_wallet_address"])
	assert.NotContains(t, attrs, "roles")
}

func TestTokenWrapper_MintToken_IdentityServiceDropsSuppliedRoles(t *testing.T) {
	// The chain-less login path hands back a placeholder operator role to keep dev
	// instances usable; it must never reach a token that every chain will honour.
	wrapper := auth.NewIdentityTokenWrapper(newTokenService(t), testJWTSecret, nil, "")

	attrs := parseAttrs(t, wrapper, buildUser(), []string{domain.RolePrivacyNodeOperator})

	assert.NotContains(t, attrs, "roles")
}

func TestTokenWrapper_MintToken_IdentityServiceStillCarriesIdentity(t *testing.T) {
	// Dropping the chain claims must not cost the identity ones the per-chain API needs.
	user := buildUser()
	wrapper := auth.NewIdentityTokenWrapper(newTokenService(t), testJWTSecret, nil, user.Email)

	str, err := wrapper.MintToken(context.Background(), user, "siwe", nil, time.Minute)
	require.NoError(t, err)

	claims, err := newTokenService(t).Parse(str)
	require.NoError(t, err)

	assert.Equal(t, user.ID.String(), claims.User.ID)
	assert.Equal(t, "Alice", claims.User.Name)
	assert.Equal(t, "alice@example.com", claims.User.Email)
	assert.Equal(t, true, claims.User.Attributes["is_admin"])
}

func TestTokenWrapper_MintToken_PerChainAPIStillEmbedsWalletAndRoles(t *testing.T) {
	// The per-chain ops-api keeps both claims — this split must not change its behaviour.
	user := buildUser()
	walletRepo := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{{
		UserID:          user.ID,
		RaylsAddress:    "0xabc",
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPrivate,
		IsActive:        true,
	}}}
	wrapper := auth.NewTokenWrapper(newTokenService(t), testJWTSecret, walletRepo, "")

	attrs := parseAttrs(t, wrapper, user, []string{domain.RolePrivacyNodeOperator})

	assert.Equal(t, "0xabc", attrs["custody_wallet_address"])
	assert.Contains(t, attrs, "roles")
}
