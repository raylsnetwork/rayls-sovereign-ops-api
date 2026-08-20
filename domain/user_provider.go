package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type OAuthProvider int

const (
	OAuthProviderUnknown   OAuthProvider = iota // 0 — reserved/unset
	OAuthProviderGoogle                         // 1
	OAuthProviderMicrosoft                      // 2
	OAuthProviderSIWE                           // 3 — Sign-In With Ethereum (MetaMask)
	OAuthProviderEmail                          // 4 — email self sign-up (oauthID = email address)
)

func (p OAuthProvider) IsValid() bool {
	switch p {
	case OAuthProviderGoogle, OAuthProviderMicrosoft, OAuthProviderSIWE, OAuthProviderEmail:
		return true
	}
	return false
}

// String returns the provider's wire name — the inverse of ParseOAuthProvider. Used in error
// payloads to tell a caller which login method an account already uses. "siwe" has no
// ParseOAuthProvider counterpart (SIWE has its own endpoint, not the OAuth callback), and an
// unset provider renders as "unknown" rather than an empty string.
func (p OAuthProvider) String() string {
	switch p {
	case OAuthProviderGoogle:
		return "google"
	case OAuthProviderMicrosoft:
		return "microsoft"
	case OAuthProviderSIWE:
		return "siwe"
	case OAuthProviderEmail:
		return "email"
	}
	return "unknown"
}

// ParseOAuthProvider converts a provider name string (as returned by go-pkgz/auth)
// to its OAuthProvider enum value.
func ParseOAuthProvider(s string) (OAuthProvider, error) {
	switch s {
	case "google":
		return OAuthProviderGoogle, nil
	case "microsoft":
		return OAuthProviderMicrosoft, nil
	case "email":
		return OAuthProviderEmail, nil
	}
	return OAuthProviderUnknown, fmt.Errorf("unknown OAuth provider: %s", s)
}

type UserOAuthProvider struct {
	ImmutableModel
	UserID        uuid.UUID     `json:"userId"`
	Provider      OAuthProvider `json:"provider"`
	OAuthID       string        `json:"oauthId"`
	Email         string        `json:"email"`
	WalletAddress string        `json:"walletAddress"` // only set for OAuthProviderSIWE
}
