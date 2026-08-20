package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-pkgz/auth/v2/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

const (
	accessTokenDuration       = 15 * time.Minute // default TTL for issued access-token cookies
	registrationTokenDuration = 5 * time.Minute
	refreshTokenDuration      = 24 * time.Hour
)

// TokenWrapper provides convenience methods around the go-pkgz/auth token service
type TokenWrapper struct {
	tokenSvc   *token.Service
	jwtSecret  string
	walletRepo core.UserWalletRepository
	// adminEmail identifies the admin user. A token minted for a user whose email
	// matches (case-insensitive) carries the `is_admin` claim. Empty = no admin.
	adminEmail string
	// identityOnly strips the per-chain `roles` claim from minted tokens. Set by the shared
	// identity service, whose token is presented to EVERY chain: a role granted on one
	// chain's AccessManager says nothing about another, so it cannot travel in the token.
	// Each chain's ops-api resolves roles per request from its own am_* mirror instead.
	//
	// `custody_wallet_address` is NOT stripped: a custody wallet is an EVM keypair, so the
	// one address works on every chain — only its on-chain state (balance, roles) differs.
	// One wallet per user, owned by identity.
	identityOnly bool
}

// NewTokenWrapper creates a new TokenWrapper.
// walletRepo is used at mint time to embed the user's custody wallet address as a
// JWT attribute (`custody_wallet_address`), so /auth/user can return it without a
// per-request DB lookup. Pass nil to disable wallet enrichment (test setups).
// adminEmail marks matching users as admin via the `is_admin` claim; pass "" for none.
func NewTokenWrapper(
	tokenSvc *token.Service,
	jwtSecret string,
	walletRepo core.UserWalletRepository,
	adminEmail string,
) *TokenWrapper {
	return &TokenWrapper{
		tokenSvc:   tokenSvc,
		jwtSecret:  jwtSecret,
		walletRepo: walletRepo,
		adminEmail: strings.ToLower(strings.TrimSpace(adminEmail)),
	}
}

// NewIdentityTokenWrapper creates a TokenWrapper for the shared identity service: it mints
// tokens carrying the user's identity and their one custody wallet, but no roles (a
// per-chain fact — see identityOnly).
//
// walletRepo is identity's own: it owns the single wallet each user has across all chains.
func NewIdentityTokenWrapper(
	tokenSvc *token.Service,
	jwtSecret string,
	walletRepo core.UserWalletRepository,
	adminEmail string,
) *TokenWrapper {
	return &TokenWrapper{
		tokenSvc:     tokenSvc,
		jwtSecret:    jwtSecret,
		walletRepo:   walletRepo,
		adminEmail:   strings.ToLower(strings.TrimSpace(adminEmail)),
		identityOnly: true,
	}
}

// buildClaims constructs the JWT claims for a user; ttl sets the token expiry explicitly.
// roles are the on-chain role names from RaylsAccessManager (e.g. ["BANK_EMPLOYEE"]); may be nil.
// Includes the user's `custody_wallet_address` in the JWT attributes when a wallet
// exists in user_wallets — surfaced to the frontend by /auth/user.
//
// The SHARED IDENTITY SERVICE mints with roles=nil and walletRepo=nil, so neither claim is
// emitted. Both are per-chain facts a cross-chain token must not assert: roles come from
// one chain's AccessManager, and the custody wallet is minted per chain. The per-chain
// ops-api resolves both from its own database on each request instead.
func (t *TokenWrapper) buildClaims(
	ctx context.Context,
	user *domain.User,
	authMethod string,
	roles []string,
	ttl time.Duration,
) (token.Claims, error) {
	attrs := map[string]any{
		"auth_method": authMethod,
	}
	// Identity-only tokens carry no roles even when the caller supplies them: the
	// chain-less login path returns a placeholder operator role to keep dev instances
	// usable, and that must not travel to every chain as an authorization claim.
	if len(roles) > 0 && !t.identityOnly {
		attrs["roles"] = roles
	}
	if t.adminEmail != "" && strings.EqualFold(strings.TrimSpace(user.Email), t.adminEmail) {
		attrs["is_admin"] = true
	}
	if t.walletRepo != nil {
		wallet, walletErr := t.walletRepo.FindByUserID(ctx, user.ID)
		if walletErr == nil && wallet != nil && wallet.RaylsAddress != "" {
			attrs["custody_wallet_address"] = wallet.RaylsAddress
		} else if walletErr != nil && !errors.Is(walletErr, core.ErrRecordNotFound) {
			return token.Claims{}, fmt.Errorf("lookup custody wallet: %w", walletErr)
		}
	}

	now := time.Now().UTC()
	return token.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{""},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
		User: &token.User{
			Name:       user.Name,
			ID:         user.ID.String(),
			Email:      user.Email,
			Attributes: attrs,
		},
	}, nil
}

// IssueToken creates a full JWT for an authenticated user and sets it as a cookie.
func (t *TokenWrapper) IssueToken(
	ctx context.Context,
	w http.ResponseWriter,
	user *domain.User,
	authMethod string,
	roles []string,
) error {
	claims, err := t.buildClaims(ctx, user, authMethod, roles, accessTokenDuration)
	if err != nil {
		return err
	}
	if _, err := t.tokenSvc.Set(w, claims); err != nil {
		return fmt.Errorf("failed to issue token: %w", err)
	}
	return nil
}

// MintToken builds the same claims as IssueToken but returns the signed token string instead of
// setting a cookie. Used by the gen-e2e-tokens dev command. roles may be nil.
func (t *TokenWrapper) MintToken(
	ctx context.Context,
	user *domain.User,
	authMethod string,
	roles []string,
	ttl time.Duration,
) (string, error) {
	claims, err := t.buildClaims(ctx, user, authMethod, roles, ttl)
	if err != nil {
		return "", err
	}
	tokenStr, err := t.tokenSvc.Token(claims)
	if err != nil {
		return "", fmt.Errorf("failed to mint token: %w", err)
	}
	return tokenStr, nil
}

// GetUserFromRequest extracts the authenticated user claims from the request
func (t *TokenWrapper) GetUserFromRequest(r *http.Request) (*token.User, error) {
	claims, _, err := t.tokenSvc.Get(r)
	if err != nil {
		return nil, fmt.Errorf("failed to get token from request: %w", err)
	}
	if claims.User == nil {
		return nil, fmt.Errorf("token contains no user information")
	}
	return claims.User, nil
}

// ClearCookie clears the auth JWT cookie.
func (t *TokenWrapper) ClearCookie(w http.ResponseWriter) {
	t.tokenSvc.Reset(w)
}

// CreateRefreshTokenString creates a signed refresh JWT and returns the token string and its jti.
// Claims: sub=userID, jti=uuid, scope="refresh", exp=24h.
func (t *TokenWrapper) CreateRefreshTokenString(userID string) (tokenStr string, jti string, err error) {
	jti = uuid.NewString()
	now := time.Now().UTC()

	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ID:        jti,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(refreshTokenDuration)),
		Issuer:    "rayls-ops-api",
		Audience:  jwt.ClaimStrings{"refresh"},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := tkn.SignedString([]byte(t.jwtSecret))
	if err != nil {
		return "", "", fmt.Errorf("sign refresh token: %w", err)
	}
	return signed, jti, nil
}

// ParseRefreshToken validates a refresh token string and returns the userID and jti.
// Returns core.UnauthorizedError for invalid, expired, or wrong-scope tokens.
func (t *TokenWrapper) ParseRefreshToken(tokenStr string) (userID string, jti string, err error) {
	tkn, parseErr := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(t.jwtSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if parseErr != nil || !tkn.Valid {
		return "", "", core.NewUnauthorizedError("invalid or expired refresh token")
	}

	claims, ok := tkn.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", "", core.NewUnauthorizedError("invalid refresh token claims")
	}

	aud := claims.Audience
	isRefresh := false
	for _, a := range aud {
		if a == "refresh" {
			isRefresh = true
			break
		}
	}
	if !isRefresh {
		return "", "", core.NewUnauthorizedError("token is not a refresh token")
	}

	if claims.Subject == "" || claims.ID == "" {
		return "", "", core.NewUnauthorizedError("refresh token missing subject or jti")
	}

	return claims.Subject, claims.ID, nil
}
