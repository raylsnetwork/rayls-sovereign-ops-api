package app

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/handlers"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/docs"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/indexer"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/di"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

const corsMaxAgeHours = 12

func RunAPI(configPath string) error {
	container, err := di.New(configPath)
	if err != nil {
		return fmt.Errorf("failed to create DI container: %w", err)
	}

	defer func() {
		if cleanupErr := container.Close(); cleanupErr != nil {
			logger.Error("Error cleaning up container", "error", cleanupErr)
		}
	}()

	conf := container.Config

	router := gin.Default()

	docs.SwaggerInfo.Title = "Rayls Ops API"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	// Temporary: serves test/index.html for manual auth flow testing.
	// Remove once the frontend is implemented.
	router.Static("/test", "./test")

	router.GET("/health", container.HealthHandler.HealthCheck)

	// Bootstrap endpoint — unauthenticated, protected at the network/infra level.
	router.POST("/admin/bootstrap", container.BootstrapHandler.Bootstrap)

	// Single wildcard handles all /auth/* to avoid Gin conflicts between static and wildcard routes.
	authHTTPHandler, _ := container.PkgzAuth.Handlers()
	router.Any("/auth/*path", func(c *gin.Context) {
		path := c.Param("path")
		method := c.Request.Method

		// Custom refresh / logout — intercept before go-pkgz/auth.
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

		// Custom OAuth2 handlers — intercept before go-pkgz/auth.
		if container.OAuthHandler != nil {
			// Email self sign-up — POST only, reuses the OAuth find-or-create + IssueToken tail.
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

		// Validate the ?from= redirect URL on login requests to prevent open redirect attacks.
		if strings.HasSuffix(path, "/login") {
			if from := c.Query("from"); from != "" {
				if !isAllowedRedirectURL(from, allowedOrigins) {
					c.JSON(http.StatusBadRequest, gin.H{"error": "redirect URL not allowed"})
					return
				}
			}
		}

		authHTTPHandler.ServeHTTP(c.Writer, c.Request)
	})

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.RequireAuth(container.TokenWrapper, container.Logger))

	// Roles are resolved per request against THIS chain (container.ChainRoleService), so a
	// shared identity token — which carries no roles, because a grant on one chain says
	// nothing about another — is still authorized correctly. Nil resolver (Access Manager
	// not wired) falls back to the JWT claim, preserving the single-instance behaviour.
	//
	// The nil check is load-bearing: ChainRoleService is a concrete pointer, and a nil
	// pointer stored in an interface is NOT interface-nil, so passing it straight through
	// would make the middleware call methods on nil instead of taking the fallback.
	var roleResolver middleware.ChainRoleResolver
	if container.ChainRoleService != nil {
		roleResolver = container.ChainRoleService
	}
	requireOperator := func() gin.HandlerFunc {
		return middleware.RequireRoleWithResolver(
			roleResolver, container.Logger, domain.RolePrivacyNodeOperator,
		)
	}

	me := api.Group("/me")
	// TODO: Allow /me group to be RequireAuth instead of operator gated
	me.Use(requireOperator())
	me.GET("", container.UserHandler.Me)

	// Onboarding — add an address pair. Authenticated-only (any logged-in user); the on-chain write
	// is operator-signed. Registered on the `api` group (RequireAuth), NOT the role-gated `me` group.
	if container.OnboardingHandler != nil {
		api.POST("/me/address-pairs", container.OnboardingHandler.AddAddressPair)
		api.GET("/me/address-pairs",
			middleware.ValidateQueryParams(handlers.AddressPairListFilter{}),
			container.OnboardingHandler.ListMine)
	}

	tokens := api.Group("/tokens")
	tokens.Use(requireOperator())
	tokens.GET("", container.TokenHandler.List)
	tokens.GET("/:address", container.TokenHandler.Details)

	// Token deploy is open to any authenticated user — registered on the `api` group
	// (RequireAuth only), not the role-gated `tokens` group.
	if container.TokenDeployHandler != nil {
		api.POST("/tokens", container.TokenDeployHandler.Deploy)
		api.POST("/tokens/estimate", container.TokenDeployHandler.Estimate)
	}

	// Per-user token permissions (what the logged-in user can do on a token) — only when the
	// Access Manager is wired. Authenticated-only (self-scoped to the user's own wallet).
	if container.TokenPermissionHandler != nil {
		api.GET("/tokens/:address/permissions", container.TokenPermissionHandler.Get)
	}

	// Mint/burn — authenticated; the handler enforces the AM permission before signing.
	// Teleport — authenticated; the service runs the registry + balance/ownership preflight before signing.
	// Pause — authenticated; stablecoin only, and gated on the contract's own `pauser` address
	// rather than an AM role, so the handler checks msg.sender equality instead.
	if container.TokenActionHandler != nil {
		api.POST("/tokens/:address/mint", container.TokenActionHandler.Mint)
		api.POST("/tokens/:address/burn", container.TokenActionHandler.Burn)
		api.POST("/tokens/:address/pause", container.TokenActionHandler.Pause)
		api.POST("/tokens/:address/teleport", container.TokenActionHandler.Teleport)
	}

	// Token registry — register an already-deployed token into the TokenRegistry. Authenticated-only
	// (any logged-in user); the on-chain write is operator-signed. Registered on the `api` group
	// (RequireAuth), NOT the role-gated `tokens` group.
	if container.TokenRegistryHandler != nil {
		api.POST("/tokens/:address/register", container.TokenRegistryHandler.Register)
		api.GET("/tokens/registry", container.TokenRegistryHandler.List)
		api.GET("/tokens/registry/pending", container.TokenRegistryHandler.ListPending)
	}

	// SSE stream of curated token events — only when NATS is configured (the worker
	// publishes the curated event; the API subscribes and fans out to connected clients).
	if container.NATSManager != nil && container.TokenStreamHandler != nil {
		api.GET("/tokens/stream", container.TokenStreamHandler.Stream)
		sub, subErr := container.NATSManager.Subscribe(
			indexer.InstanceSubject(container.Config.InstanceName, indexer.SubjectTokenSSE),
			func(data []byte) { container.TokenStreamHub.Broadcast(data) },
		)
		if subErr != nil {
			container.Logger.Warn("Failed to subscribe to token SSE subject — live updates disabled", "error", subErr)
		} else {
			defer func() { _ = sub.Unsubscribe() }()
		}
	}

	// Operator-only admin routes.
	admin := api.Group("/v1/admin")
	admin.Use(requireOperator())

	// Onboarding admin — discover all pending pairs and approve/reject them. Identity for the PATCH
	// comes from the path :id, never the body.
	if container.OnboardingHandler != nil {
		admin.GET("/address-pairs/pending", container.OnboardingHandler.ListAllPending)
		admin.PATCH("/users/:id/address-pairs/status", container.OnboardingHandler.SetApprovalStatus)
	}

	// Token registry admin — move a registered token through the governance status lifecycle. The
	// token address comes from the path :address, never the body; the on-chain write is operator-signed.
	if container.TokenRegistryHandler != nil {
		admin.PATCH("/tokens/:address/status", container.TokenRegistryHandler.SetStatus)
		admin.POST("/tokens/:address/freeze", container.TokenRegistryHandler.Freeze)
		admin.POST("/tokens/:address/unfreeze", container.TokenRegistryHandler.Unfreeze)
		admin.POST("/tokens/:address/submit", container.TokenRegistryHandler.Submit)
	}
	// Wallet balances — list endpoint is authenticated; the SSE stream is authenticated
	// and per-handler filtered to the caller's own wallet.
	wallets := api.Group("/wallets")
	wallets.GET("/:address/balances", container.WalletBalanceHandler.List)
	wallets.GET("/:address/balances/:tokenAddress", container.WalletBalanceHandler.Details)
	if container.NATSManager != nil && container.WalletBalanceStreamHandler != nil {
		wallets.GET("/balances/stream", container.WalletBalanceStreamHandler.Stream)
		sub, subErr := container.NATSManager.Subscribe(
			indexer.InstanceSubject(container.Config.InstanceName, indexer.SubjectWalletBalanceSSE),
			func(data []byte) { container.WalletBalanceStreamHub.Broadcast(data) },
		)
		if subErr != nil {
			container.Logger.Warn(
				"Failed to subscribe to wallet balance SSE subject — live updates disabled",
				"error",
				subErr,
			)
		} else {
			defer func() { _ = sub.Unsubscribe() }()
		}
	}

	if container.AccessManagerHandler != nil {
		admin.GET("/roles", container.AccessManagerHandler.ListRoles)
		admin.GET("/roles/:roleId", container.AccessManagerHandler.GetRole)
		admin.GET("/roles/:roleId/members", container.AccessManagerHandler.ListRoleMembers)
		admin.GET("/accounts/:address/roles", container.AccessManagerHandler.ListAccountRoles)
		admin.GET("/managed-contracts", container.AccessManagerHandler.ListManagedContracts)
		admin.GET("/managed-contracts/:address", container.AccessManagerHandler.GetManagedContract)
		admin.GET("/managed-contracts/:address/permissions", container.AccessManagerHandler.ListFunctionPermissions)
		admin.GET("/scheduled-operations", container.AccessManagerHandler.ListScheduledOperations)
		admin.GET("/events", container.AccessManagerHandler.ListEvents)
	}

	if err = router.Run(":8080"); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

// isAllowedRedirectURL checks that the given URL's origin matches one of the allowed CORS origins.
// This prevents open redirect attacks via the ?from= parameter.
func isAllowedRedirectURL(rawURL string, allowedOrigins []string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return false
	}
	origin := parsed.Scheme + "://" + parsed.Host
	for _, allowed := range allowedOrigins {
		if strings.EqualFold(origin, strings.TrimRight(allowed, "/")) {
			return true
		}
	}
	return false
}
