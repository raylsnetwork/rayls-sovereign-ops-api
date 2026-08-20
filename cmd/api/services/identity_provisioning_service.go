package services

import (
	"context"
	"fmt"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// identityProvisioningService activates a new account in the SHARED IDENTITY SERVICE and
// gives it its custody wallet.
//
// Identity owns wallets. A custody wallet is an EVM keypair, so ONE serves every chain —
// only its on-chain state (balance, roles) differs, and those are per-chain facts the
// chain itself supplies. Minting here (rather than per chain) is what makes the address
// stable everywhere, and it is the only way a chain can resolve a wallet for a user who
// has never touched it.
//
// What this does NOT do is fund the wallet or grant it roles: both need a chain, and this
// service deliberately has none. The per-chain ops-api does that when the user arrives.
//
// Without this, every first login 403s: applyLoginDecisionTree returns
// RoleAssignmentPendingError for waiting_role_assignment when no provisioner is wired.
type identityProvisioningService struct {
	userRepo     core.UserRepository
	walletRepo   core.UserWalletRepository
	custody      core.CustodyService
	providerType domain.CustodyProviderType
	log          logger.Logger
}

var _ core.ProvisioningService = (*identityProvisioningService)(nil)

// NewIdentityProvisioningService creates the identity service's account activator.
func NewIdentityProvisioningService(
	userRepo core.UserRepository,
	walletRepo core.UserWalletRepository,
	custody core.CustodyService,
	providerType domain.CustodyProviderType,
	log logger.Logger,
) core.ProvisioningService {
	return &identityProvisioningService{
		userRepo:     userRepo,
		walletRepo:   walletRepo,
		custody:      custody,
		providerType: providerType,
		log:          log,
	}
}

// Provision gives the user a custody wallet (if they have none) and marks the account
// usable. Idempotent: safe on every login.
func (s *identityProvisioningService) Provision(ctx context.Context, user *domain.User) error {
	if err := s.ensureWallet(ctx, user); err != nil {
		return fmt.Errorf("ensure wallet: %w", err)
	}

	if user.Status == domain.UserStatusRoleAssigned {
		return nil
	}

	user.Status = domain.UserStatusRoleAssigned
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("update user status: %w", err)
	}

	s.log.Info("identity account activated", "userID", user.ID)
	return nil
}

// ensureWallet mints the user's single custody wallet on first login. Runs before the
// status check so an account activated before identity owned wallets still gets one.
func (s *identityProvisioningService) ensureWallet(ctx context.Context, user *domain.User) error {
	if s.custody == nil || s.walletRepo == nil {
		return nil
	}

	wallet, err := ensureWalletFor(ctx, mintDeps{
		custody:      s.custody,
		wallets:      s.walletRepo,
		log:          s.log,
		providerType: s.providerType,
	}, user.ID, domain.WalletChainPrivate)
	if err != nil {
		return err
	}

	s.log.Info("custody wallet ready", "userID", user.ID, "address", wallet.RaylsAddress)
	return nil
}
