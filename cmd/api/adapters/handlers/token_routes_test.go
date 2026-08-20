package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTokenRoutes_StaticAndParamCoexist guards the route layout: the static GET paths
// (/api/tokens/stream, /api/tokens/registry, /api/tokens/registry/pending) and the param paths
// (/api/tokens/:address) must coexist without a registration panic, with the static paths taking
// priority.
func TestTokenRoutes_StaticAndParamCoexist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	require.NotPanics(t, func() {
		api := r.Group("/api")
		tokens := api.Group("/tokens")
		tokens.GET("", func(c *gin.Context) { c.String(http.StatusOK, "list") })
		tokens.GET("/:address", func(c *gin.Context) { c.String(http.StatusOK, "details:"+c.Param("address")) })
		api.GET("/tokens/stream", func(c *gin.Context) { c.String(http.StatusOK, "stream") })
		api.GET("/tokens/registry", func(c *gin.Context) { c.String(http.StatusOK, "registry") })
		api.GET("/tokens/registry/pending", func(c *gin.Context) { c.String(http.StatusOK, "pending") })
	})

	cases := map[string]string{
		"/api/tokens":                  "list",
		"/api/tokens/stream":           "stream",
		"/api/tokens/registry":         "registry",
		"/api/tokens/registry/pending": "pending",
		"/api/tokens/0xabc":            "details:0xabc",
	}
	for path, want := range cases {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
		assert.Equal(t, http.StatusOK, w.Code, path)
		assert.Equal(t, want, w.Body.String(), path)
	}
}
