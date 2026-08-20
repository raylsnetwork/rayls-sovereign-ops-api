package auth

import (
	"context"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// FetchRole looks up the user's wallet then queries the on-chain role service.
// Returns "" on any error — callers issue the token without a role claim rather than blocking login.
func FetchRole(
	ctx context.Context,
	userID uuid.UUID,
	walletRepo core.UserWalletRepository,
	roleService core.RoleService,
	log logger.Logger,
) string {
	wallet, err := walletRepo.FindByUserID(ctx, userID)
	if err != nil {
		return ""
	}
	role, err := roleService.GetUserRole(ctx, userID, wallet.RaylsAddress)
	if err != nil {
		log.Warn("failed to fetch on-chain role", "userID", userID, "error", err)
		return ""
	}
	return role
}
