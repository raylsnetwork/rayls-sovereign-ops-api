package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// stubRoleResolver stands in for the per-chain ChainRoleService.
type stubRoleResolver struct {
	roles  []string
	err    error
	called bool
}

func (s *stubRoleResolver) RolesForUser(_ context.Context, _ uuid.UUID) ([]string, error) {
	s.called = true
	return s.roles, s.err
}

// runRequireRole exercises the middleware with a user already in the context, as
// RequireAuth would have left it.
func runRequireRole(
	t *testing.T,
	resolver ChainRoleResolver,
	claimRoles []interface{},
	userID string,
) *httptest.ResponseRecorder {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/protected", nil)

	attrs := map[string]interface{}{}
	if claimRoles != nil {
		attrs["roles"] = claimRoles
	}
	c.Set(userContextKey, &token.User{ID: userID, Attributes: attrs})

	RequireRoleWithResolver(resolver, &testutil.StubLogger{}, domain.RolePrivacyNodeOperator)(c)
	return w
}

func TestRequireRole_ResolverGrantsAccessWithoutRolesClaim(t *testing.T) {
	// The whole point of the split: an identity token carries no roles, and the chain
	// resolves them itself.
	resolver := &stubRoleResolver{roles: []string{domain.RolePrivacyNodeOperator}}

	w := runRequireRole(t, resolver, nil, uuid.New().String())

	assert.True(t, resolver.called, "resolver was not consulted")
	assert.NotEqual(t, http.StatusForbidden, w.Code)
}

func TestRequireRole_ResolverDeniesWhenUserHasNoRoleOnThisChain(t *testing.T) {
	// A user valid on another chain holds nothing here.
	resolver := &stubRoleResolver{roles: nil}

	w := runRequireRole(t, resolver, nil, uuid.New().String())

	require.Equal(t, http.StatusForbidden, w.Code)
}

func TestRequireRole_ResolverIgnoresRolesClaim(t *testing.T) {
	// A claim must never grant access once the chain is authoritative — otherwise a token
	// minted for one chain would authorize actions on another.
	resolver := &stubRoleResolver{roles: nil}

	w := runRequireRole(t, resolver,
		[]interface{}{domain.RolePrivacyNodeOperator}, uuid.New().String())

	require.Equal(t, http.StatusForbidden, w.Code)
}

func TestRequireRole_ResolverErrorIsNotAnAccessGrant(t *testing.T) {
	// A failing lookup must never fall open.
	resolver := &stubRoleResolver{err: errors.New("db down")}

	w := runRequireRole(t, resolver, nil, uuid.New().String())

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestRequireRole_StaleSessionReturns401NotServerError(t *testing.T) {
	// A token naming a user identity no longer has (the identity DB was recreated) must
	// surface as 401 so the browser drops the dead cookie. As a 500 the user is stuck
	// logged in with no permissions and no signal to sign in again.
	resolver := &stubRoleResolver{
		err: core.NewUnauthorizedError("session is no longer valid, please sign in again"),
	}

	w := runRequireRole(t, resolver, nil, uuid.New().String())

	require.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "sign in again")
}

func TestRequireRole_ResolverDeniesUnparseableUserID(t *testing.T) {
	// A token whose subject is not one of our user ids maps to no wallet, so no roles.
	resolver := &stubRoleResolver{roles: []string{domain.RolePrivacyNodeOperator}}

	w := runRequireRole(t, resolver, nil, "not-a-uuid")

	require.Equal(t, http.StatusForbidden, w.Code)
	assert.False(t, resolver.called, "resolver was queried with an invalid user id")
}

func TestRequireRole_FallsBackToClaimWhenNoResolver(t *testing.T) {
	// Without a resolver (self-contained ops-api) the previous claim-based behaviour stands.
	w := runRequireRole(t, nil,
		[]interface{}{domain.RolePrivacyNodeOperator}, uuid.New().String())

	assert.NotEqual(t, http.StatusForbidden, w.Code)
}

func TestRequireRole_FallbackDeniesWhenClaimLacksRole(t *testing.T) {
	w := runRequireRole(t, nil, []interface{}{"SOME_OTHER_ROLE"}, uuid.New().String())

	require.Equal(t, http.StatusForbidden, w.Code)
}
