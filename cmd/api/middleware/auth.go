package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/auth"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

const userContextKey = "auth_user"

// RequireAuth validates the JWT token and stores the user in the Gin context.
func RequireAuth(tokenWrapper *auth.TokenWrapper, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := tokenWrapper.GetUserFromRequest(c.Request)
		if err != nil {
			log.Debug("Authentication failed", "error", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}

		if user.ID == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "account not authorized"})
			return
		}

		c.Set(userContextKey, user)
		c.Next()
	}
}

// GetAuthUser retrieves the authenticated user from the Gin context.
// Returns nil, false if RequireAuth has not run or the value is missing.
func GetAuthUser(c *gin.Context) (*token.User, bool) {
	val, exists := c.Get(userContextKey)
	if !exists {
		return nil, false
	}
	u, ok := val.(*token.User)
	return u, ok
}

// ChainRoleResolver reports the roles a user holds ON THIS CHAIN. Implemented by
// services.ChainRoleService, which reads the am_* tables the AccessManager indexer keeps
// in sync. Declared here (rather than imported) to keep middleware free of a services
// dependency.
type ChainRoleResolver interface {
	RolesForUser(ctx context.Context, userID uuid.UUID) ([]string, error)
}

// RequireRole checks that the authenticated user holds at least one of the allowed roles.
//
// Roles are resolved PER REQUEST against this chain when a resolver is wired, because a
// shared identity token cannot carry them: it is presented to every chain, and a grant on
// one says nothing about another. It also means a revoked role stops working as soon as
// the indexer sees the revocation, instead of lingering until the next login.
//
// With no resolver (a self-contained ops-api that still mints its own tokens) it falls
// back to the JWT's "roles" attribute, preserving the previous behaviour exactly.
func RequireRole(roles ...string) gin.HandlerFunc {
	return RequireRoleWithResolver(nil, nil, roles...)
}

// RequireRoleWithResolver is RequireRole with live per-chain role resolution.
func RequireRoleWithResolver(resolver ChainRoleResolver, log logger.Logger, roles ...string) gin.HandlerFunc {
	allowed := make(map[string]bool, len(roles))
	for _, r := range roles {
		allowed[r] = true
	}

	return func(c *gin.Context) {
		userVal, exists := c.Get(userContextKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}

		user, ok := userVal.(*token.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
			return
		}

		userRoles, err := rolesForRequest(c, user, resolver, log)
		if err != nil {
			// A stale session (token naming a user identity no longer has, e.g. after the
			// identity DB was recreated) is reported as Unauthorized on purpose: only a 401
			// makes the browser drop the dead cookie and re-authenticate. Collapsing it into
			// a 500 leaves the user in a permanent loop — logged in, no permissions, and no
			// signal to log in again.
			var unauthorizedErr *core.UnauthorizedError
			if errors.As(err, &unauthorizedErr) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorizedErr.Message})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not resolve permissions"})
			return
		}

		for _, r := range userRoles {
			if allowed[r] {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
	}
}

// rolesForRequest resolves the caller's roles on this chain, falling back to the JWT
// claim when no resolver is wired.
func rolesForRequest(
	c *gin.Context,
	user *token.User,
	resolver ChainRoleResolver,
	log logger.Logger,
) ([]string, error) {
	if resolver == nil {
		claimed, _ := user.Attributes["roles"].([]interface{})
		out := make([]string, 0, len(claimed))
		for _, r := range claimed {
			if s, ok := r.(string); ok {
				out = append(out, s)
			}
		}
		return out, nil
	}

	// A token whose subject is not one of our user ids cannot be mapped to a wallet, so
	// it holds nothing here. That is a denial (no roles), not a server error — hence no
	// error is propagated.
	userID, parseErr := uuid.Parse(user.ID)
	if parseErr != nil {
		if log != nil {
			log.Debug("Token subject is not a user id — no roles on this chain", "subject", user.ID)
		}
		return nil, nil //nolint:nilerr // unparseable subject = no roles, deliberately not an error
	}

	resolved, err := resolver.RolesForUser(c.Request.Context(), userID)
	if err != nil {
		if log != nil {
			log.Error("Failed to resolve chain roles", "userID", user.ID, "error", err)
		}
		return nil, err
	}
	return resolved, nil
}
