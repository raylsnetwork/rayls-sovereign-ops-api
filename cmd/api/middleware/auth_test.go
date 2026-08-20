package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	authlib "github.com/go-pkgz/auth/v2"
	"github.com/go-pkgz/auth/v2/avatar"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/auth"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

const testJWTSecret = "test-secret"

// newTokenWrapper builds a real go-pkgz/auth token service (matching auth.NewPkgzAuthService) so
// tests can mint valid tokens and exercise RequireAuth end-to-end.
func newTokenWrapper(t *testing.T) *auth.TokenWrapper {
	t.Helper()
	svc := authlib.NewService(authlib.Opts{
		SecretReader: token.SecretFunc(func(_ string) (string, error) { return testJWTSecret, nil }),
		Issuer:       "rayls-ops-api",
		DisableXSRF:  true,
		AvatarStore:  avatar.NewNoOp(),
	})
	return auth.NewTokenWrapper(svc.TokenService(), testJWTSecret, nil, "")
}

// newTokenWrapperWithAdmin is newTokenWrapper but with an admin email configured.
func newTokenWrapperWithAdmin(t *testing.T, adminEmail string) *auth.TokenWrapper {
	t.Helper()
	svc := authlib.NewService(authlib.Opts{
		SecretReader: token.SecretFunc(func(_ string) (string, error) { return testJWTSecret, nil }),
		Issuer:       "rayls-ops-api",
		DisableXSRF:  true,
		AvatarStore:  avatar.NewNoOp(),
	})
	return auth.NewTokenWrapper(svc.TokenService(), testJWTSecret, nil, adminEmail)
}

func TestMintToken_StampsIsAdminForMatchingEmail(t *testing.T) {
	// A token minted for the configured admin email carries is_admin=true; others don't.
	gin.SetMode(gin.TestMode)
	tw := newTokenWrapperWithAdmin(t, "Admin@Example.com")

	admin := &domain.User{Model: domain.Model{ID: uuid.New()}, Name: "Admin", Email: "admin@example.com"}
	regular := &domain.User{Model: domain.Model{ID: uuid.New()}, Name: "Reg", Email: "reg@example.com"}

	adminAttrs := mintAndGetAttributes(t, tw, admin)
	regularAttrs := mintAndGetAttributes(t, tw, regular)

	assert.Equal(t, true, adminAttrs["is_admin"])
	_, hasAdmin := regularAttrs["is_admin"]
	assert.False(t, hasAdmin)
}

// mintAndGetAttributes mints a token for user and returns the attributes recovered
// from it via RequireAuth, so the assertion exercises the real mint→verify round-trip.
func mintAndGetAttributes(t *testing.T, tw *auth.TokenWrapper, user *domain.User) map[string]interface{} {
	t.Helper()
	tokenStr, err := tw.MintToken(context.Background(), user, "email", nil, time.Hour)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/protected", nil)
	c.Request.Header.Set("X-JWT", tokenStr)
	RequireAuth(tw, &testutil.StubLogger{})(c)

	got, ok := GetAuthUser(c)
	require.True(t, ok)
	return got.Attributes
}

func TestRequireAuth_ValidTokenSetsUserAndContinues(t *testing.T) {
	// A request carrying a valid JWT stores the user in the context and calls the next handler.
	gin.SetMode(gin.TestMode)
	tw := newTokenWrapper(t)
	user := &domain.User{Model: domain.Model{ID: uuid.New()}, Name: "Alice", Email: "alice@example.com"}
	tokenStr, err := tw.MintToken(context.Background(), user, "web2", nil, time.Hour)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/protected", nil)
	c.Request.Header.Set("X-JWT", tokenStr)

	called := false
	RequireAuth(tw, &testutil.StubLogger{})(c)
	if !c.IsAborted() {
		called = true
	}

	assert.True(t, called)
	got, ok := GetAuthUser(c)
	require.True(t, ok)
	assert.Equal(t, user.ID.String(), got.ID)
}

func TestRequireAuth_MissingTokenReturns401(t *testing.T) {
	// A request with no token is aborted with 401.
	gin.SetMode(gin.TestMode)
	tw := newTokenWrapper(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/protected", nil)

	RequireAuth(tw, &testutil.StubLogger{})(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireAuth_InvalidTokenReturns401(t *testing.T) {
	// A request carrying a malformed token is aborted with 401.
	gin.SetMode(gin.TestMode)
	tw := newTokenWrapper(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/protected", nil)
	c.Request.Header.Set("X-JWT", "not-a-valid-jwt")

	RequireAuth(tw, &testutil.StubLogger{})(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestExtractBearerToken_MapsAuthorizationToXJWT(t *testing.T) {
	// A "Bearer <token>" Authorization header is copied to the X-JWT header for go-pkgz/auth.
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	c.Request.Header.Set("Authorization", "Bearer abc.def.ghi")

	ExtractBearerToken()(c)

	assert.Equal(t, "abc.def.ghi", c.Request.Header.Get("X-JWT"))
}

func TestExtractBearerToken_NoAuthorizationHeaderLeavesXJWTEmpty(t *testing.T) {
	// With no Authorization header, X-JWT is left unset and the chain continues.
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	ExtractBearerToken()(c)

	assert.Empty(t, c.Request.Header.Get("X-JWT"))
}

func TestExtractBearerToken_NonBearerAuthorizationIgnored(t *testing.T) {
	// A non-Bearer Authorization scheme (e.g. Basic) is ignored — X-JWT stays unset.
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	c.Request.Header.Set("Authorization", "Basic dXNlcjpwYXNz")

	ExtractBearerToken()(c)

	assert.Empty(t, c.Request.Header.Get("X-JWT"))
}

func TestRequireRole_AllowsMatchingRole(t *testing.T) {
	// A user holding one of the allowed roles passes through to the next handler.
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/admin", nil)
	c.Set(userContextKey, &token.User{ID: "u1", Attributes: map[string]any{"roles": []any{"ADMIN"}}})

	RequireRole("ADMIN")(c)

	assert.False(t, c.IsAborted())
}

func TestRequireRole_DeniesWithoutMatchingRole(t *testing.T) {
	// A user lacking every allowed role is aborted with 403.
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/admin", nil)
	c.Set(userContextKey, &token.User{ID: "u1", Attributes: map[string]any{"roles": []any{"USER"}}})

	RequireRole("ADMIN")(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestRequireRole_UnauthenticatedReturns401(t *testing.T) {
	// With no user in the context, RequireRole aborts with 401 before any role check.
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/admin", nil)

	RequireRole("ADMIN")(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetAuthUser_MissingReturnsFalse(t *testing.T) {
	// GetAuthUser reports not-found when RequireAuth has not populated the context.
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	_, ok := GetAuthUser(c)

	assert.False(t, ok)
}
