package handlers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services/testutil"
)

// stubBootstrapService records whether the bootstrap actually ran.
type stubBootstrapService struct{ called bool }

func (s *stubBootstrapService) Bootstrap(_ context.Context, _ string) (string, error) {
	s.called = true
	return "0xadmin", nil
}

func postBootstrap(t *testing.T, h *BootstrapHandler, authHeader string) (*httptest.ResponseRecorder, *gin.Context) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/admin/bootstrap",
		bytes.NewBufferString(`{"email":"admin@example.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")
	if authHeader != "" {
		c.Request.Header.Set("Authorization", authHeader)
	}
	return w, c
}

func TestBootstrapHandler_Bootstrap_RejectsWrongToken(t *testing.T) {
	// A bad token must not create the first admin — whoever bootstraps owns the deployment.
	svc := &stubBootstrapService{}
	h := NewBootstrapHandler(svc, "correct-token", &testutil.StubLogger{})

	w, c := postBootstrap(t, h, "Bearer wrong-token")
	h.Bootstrap(c)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, svc.called, "bootstrap ran despite an invalid token")
}

func TestBootstrapHandler_Bootstrap_RejectsMissingToken(t *testing.T) {
	// No Authorization header at all is equally unauthorized when a token is configured.
	svc := &stubBootstrapService{}
	h := NewBootstrapHandler(svc, "correct-token", &testutil.StubLogger{})

	w, c := postBootstrap(t, h, "")
	h.Bootstrap(c)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, svc.called, "bootstrap ran with no token presented")
}

func TestBootstrapHandler_Bootstrap_AcceptsCorrectToken(t *testing.T) {
	// The configured token lets the bootstrap through.
	svc := &stubBootstrapService{}
	h := NewBootstrapHandler(svc, "correct-token", &testutil.StubLogger{})

	w, c := postBootstrap(t, h, "Bearer correct-token")
	h.Bootstrap(c)

	require.Equal(t, http.StatusCreated, w.Code)
	assert.True(t, svc.called)
}

func TestBootstrapHandler_Bootstrap_UnauthenticatedWhenNoTokenConfigured(t *testing.T) {
	// Empty token preserves the previous behaviour for dev / network-isolated deployments.
	svc := &stubBootstrapService{}
	h := NewBootstrapHandler(svc, "", &testutil.StubLogger{})

	w, c := postBootstrap(t, h, "")
	h.Bootstrap(c)

	require.Equal(t, http.StatusCreated, w.Code)
	assert.True(t, svc.called)
}
