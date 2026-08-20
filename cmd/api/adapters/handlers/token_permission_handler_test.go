package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services/testutil"
)

type fakeTokenPermissionService struct {
	res        *core.TokenPermissions
	gotWallet  string
	gotAddress string
}

func (f *fakeTokenPermissionService) GetTokenPermissions(
	_ context.Context,
	contractAddress, walletAddress string,
) (*core.TokenPermissions, error) {
	f.gotAddress = contractAddress
	f.gotWallet = walletAddress
	return f.res, nil
}

func newPermContext(t *testing.T, address, authUserID string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/tokens/"+address+"/permissions", nil)
	c.Params = gin.Params{{Key: "address", Value: address}}
	if authUserID != "" {
		c.Set("auth_user", &token.User{ID: authUserID})
	}
	return c, w
}

func TestTokenPermissionHandler_Get_ReturnsPermissions(t *testing.T) {
	// Resolves the user's wallet and returns the service result.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0xWALLET")
	svc := &fakeTokenPermissionService{res: &core.TokenPermissions{
		ContractAddress: "0xtoken", WalletAddress: "0xwallet", CanMint: true,
		Functions: []core.TokenFunction{{Selector: "0x40c10f19", Name: "mint"}},
	}}
	h := NewTokenPermissionHandler(svc, &walletRepo, &testutil.StubLogger{})

	c, w := newPermContext(t, "0xtoken", userID.String())
	h.Get(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "0xWALLET", svc.gotWallet)
	assert.Equal(t, "0xtoken", svc.gotAddress)
	var resp core.TokenPermissions
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.CanMint)
	require.Len(t, resp.Functions, 1)
	assert.Equal(t, "mint", resp.Functions[0].Name)
}

func TestTokenPermissionHandler_Get_NoWalletReturns400(t *testing.T) {
	// A user without a custody wallet gets 400.
	userID := uuid.New()
	walletRepo := testutil.FakeUserWalletRepository{}
	h := NewTokenPermissionHandler(&fakeTokenPermissionService{}, &walletRepo, &testutil.StubLogger{})

	c, w := newPermContext(t, "0xtoken", userID.String())
	h.Get(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTokenPermissionHandler_Get_NoAuthReturns401(t *testing.T) {
	// No authenticated user -> 401.
	walletRepo := testutil.FakeUserWalletRepository{}
	h := NewTokenPermissionHandler(&fakeTokenPermissionService{}, &walletRepo, &testutil.StubLogger{})

	c, w := newPermContext(t, "0xtoken", "")
	h.Get(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
