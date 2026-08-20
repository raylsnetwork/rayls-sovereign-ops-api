package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

func newDetailsContext(t *testing.T, address string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/tokens/"+address, nil)
	c.Params = gin.Params{{Key: "address", Value: address}}
	return c, w
}

func TestTokenHandler_Details_ReturnsToken(t *testing.T) {
	// Details returns the token matching the contract address in the path.
	addr := "0xabc"
	repo := &fakeTokenRepo{byAddress: map[string]*domain.Token{
		addr: {
			Model:           domain.Model{ID: uuid.New()},
			Name:            "Token",
			Symbol:          "TKN",
			ContractAddress: addr,
			ErcStandard:     domain.ErcStandardERC20,
		},
	}}
	h := NewTokenHandler(repo, &testutil.StubLogger{})

	c, w := newDetailsContext(t, addr)
	h.Details(c)

	require.Equal(t, http.StatusOK, w.Code)
	var resp tokenResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, addr, resp.ContractAddress)
	assert.Equal(t, "TKN", resp.Symbol)
}

func TestTokenHandler_Details_NotFoundReturns404(t *testing.T) {
	// An unknown contract address yields 404.
	repo := &fakeTokenRepo{}
	h := NewTokenHandler(repo, &testutil.StubLogger{})

	c, w := newDetailsContext(t, "0xdoesnotexist")
	h.Details(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
