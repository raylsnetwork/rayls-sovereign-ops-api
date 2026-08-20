package auth

import (
	authlib "github.com/go-pkgz/auth/v2"
	"github.com/go-pkgz/auth/v2/avatar"
	"github.com/go-pkgz/auth/v2/token"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/config"
)

// NewPkgzAuthService creates the go-pkgz/auth service used only for JWT token issuance.
// OAuth providers (Google, Microsoft) are handled by OAuthHandler; SIWE is added via AddCustomHandler.
func NewPkgzAuthService(cfg config.Auth, url string) *authlib.Service {
	return authlib.NewService(authlib.Opts{
		SecretReader:  token.SecretFunc(func(_ string) (string, error) { return cfg.JWTSecret, nil }),
		URL:           url,
		AvatarStore:   avatar.NewNoOp(),
		Issuer:        "rayls-ops-api",
		SecureCookies: cfg.SecureCookies,
		DisableXSRF:   true,
		// Share the session cookie across subdomains so the per-chain playground
		// can forward it to this api (empty keeps the previous host-only cookie).
		JWTCookieDomain: cfg.CookieDomain,
	})
}
