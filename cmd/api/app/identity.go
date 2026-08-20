package app

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/di"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// RunIdentity starts the SHARED IDENTITY SERVICE.
//
// One instance serves every chain. It owns logging in (SIWE, Google/Microsoft OAuth,
// email signup), the user record, and the session JWT — and nothing else. There is no
// chain here: no RPC, no AccessManager, no indexers, no tokens.
//
// The token it mints is deliberately identity-only (id, name, email, auth_method,
// is_admin). Roles and signing wallets belong to a single chain, so a token that is
// valid across all of them must not assert either; each chain's ops-api authorizes
// requests against its own database instead.
func RunIdentity(configPath string) error {
	container, err := di.NewIdentity(configPath)
	if err != nil {
		return fmt.Errorf("failed to create identity container: %w", err)
	}

	defer func() {
		if cleanupErr := container.Close(); cleanupErr != nil {
			logger.Error("Error cleaning up identity container", "error", cleanupErr)
		}
	}()

	// Migrations ran in NewIdentity, against the identity schema (see the container).
	conf := container.Config

	router := gin.Default()

	allowedOrigins := strings.Split(conf.CorsUrls, ";")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "X-JWT", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           corsMaxAgeHours * time.Hour,
	}))

	router.Use(middleware.ExtractBearerToken())
	router.Use(middleware.ValidateQueryEncoding())

	router.GET("/health", container.HealthHandler.HealthCheck)

	// Bootstrap the first admin. Unauthenticated by design and guarded only by "no users
	// exist yet" — it must therefore stay unreachable from the public internet.
	router.POST("/admin/bootstrap", container.BootstrapHandler.Bootstrap)

	// One wildcard for all of /auth/*, avoiding Gin's static-vs-wildcard route conflict.
	// Mirrors RunAPI's dispatch so both services expose an identical auth surface.
	authHTTPHandler, _ := container.PkgzAuth.Handlers()
	router.Any("/auth/*path", func(c *gin.Context) {
		path := c.Param("path")
		method := c.Request.Method

		if method == http.MethodPost {
			switch path {
			case "/refresh":
				container.AuthHandler.Refresh(c)
				return
			case "/logout":
				container.AuthHandler.Logout(c)
				return
			}
		}

		if container.OAuthHandler != nil {
			if method == http.MethodPost && path == "/signup" {
				container.OAuthHandler.EmailSignup(c.Writer, c.Request)
				return
			}
			switch path {
			case "/google/login":
				container.OAuthHandler.GoogleLogin(c.Writer, c.Request)
				return
			case "/google/callback":
				container.OAuthHandler.GoogleCallback(c.Writer, c.Request)
				return
			case "/microsoft/login":
				container.OAuthHandler.MicrosoftLogin(c.Writer, c.Request)
				return
			case "/microsoft/callback":
				container.OAuthHandler.MicrosoftCallback(c.Writer, c.Request)
				return
			}
		}

		// Open-redirect guard on ?from= (same rule as the ops-api).
		if strings.HasSuffix(path, "/login") {
			if from := c.Query("from"); from != "" {
				if !isAllowedRedirectURL(from, allowedOrigins) {
					c.JSON(http.StatusBadRequest, gin.H{"error": "redirect URL not allowed"})
					return
				}
			}
		}

		// go-pkgz built-ins: /auth/user, /auth/status, /auth/list, SIWE login/callback.
		authHTTPHandler.ServeHTTP(c.Writer, c.Request)
	})

	// Authenticated identity endpoints. NOT role-gated: roles are per-chain, and this
	// service has no chain to read them from. Being authenticated is the whole check.
	api := router.Group("/api")
	api.Use(middleware.RequireAuth(container.TokenWrapper, container.Logger))
	api.GET("/me", container.UserHandler.Me)

	port := conf.IdentityPort
	if port == "" {
		port = "8090"
	}

	container.Logger.Info("Identity service listening", "port", port)
	if err := router.Run(":" + port); err != nil {
		return fmt.Errorf("failed to run identity service: %w", err)
	}
	return nil
}
