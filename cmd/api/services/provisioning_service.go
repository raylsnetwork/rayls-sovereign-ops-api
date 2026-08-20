package services

import (
	"context"
	"fmt"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

var _ core.ProvisioningService = (*provisioningService)(nil)

type provisioningService struct {
	userRepo     core.UserRepository
	walletRepo   core.UserWalletRepository
	custody      core.CustodyService
	providerType domain.CustodyProviderType
	funder       core.WalletFunder // optional; nil on gasless chains (no funding needed)
	granter      core.RoleGranter  // optional; nil when role assignment is out of band (no grantor key)
	log          logger.Logger
}

func NewProvisioningService(
	userRepo core.UserRepository,
	walletRepo core.UserWalletRepository,
	custody core.CustodyService,
	custodyProvider domain.CustodyProviderType,
	funder core.WalletFunder,
	granter core.RoleGranter,
	log logger.Logger,
) core.ProvisioningService {
	return &provisioningService{
		userRepo:     userRepo,
		walletRepo:   walletRepo,
		custody:      custody,
		providerType: custodyProvider,
		funder:       funder,
		granter:      granter,
		log:          log,
	}
}

// Provision ensures the user has a custody wallet and, when a role granter is wired, that the
// wallet holds the on-chain roles it needs (FACTORY_DEPLOYER to deploy tokens, PRIVACY_NODE_OPERATOR
// for the login role check). Without a granter, on-chain role assignment stays out of band (the
// deploy/ops tooling grants it).
//
// It is idempotent — safe to call on every login. The grants run on EVERY login (not only first
// provision): the granter no-ops per role the wallet already holds, so this is cheap and self-heals
// wallets provisioned before auto-granting existed.
func (s *provisioningService) Provision(ctx context.Context, user *domain.User) error {
	wallet, err := s.ensureWallet(ctx, user)
	if err != nil {
		return fmt.Errorf("ensure wallet: %w", err)
	}

	// Fund + grant on every login. Both are idempotent (fund tops up only the shortfall; grant
	// checks hasRole first) and best-effort — a failure is logged but must not block login.
	s.prepareWalletForChain(ctx, wallet.RaylsAddress)

	if user.Status != domain.UserStatusRoleAssigned {
		user.Status = domain.UserStatusRoleAssigned
		if err := s.userRepo.Update(ctx, user); err != nil {
			return fmt.Errorf("update user status: %w", err)
		}
		s.log.Info("user provisioned", "userID", user.ID, "address", wallet.RaylsAddress)
	}
	return nil
}

// ensureWallet returns the user's existing wallet, or mints one write-ahead if absent.
func (s *provisioningService) ensureWallet(ctx context.Context, user *domain.User) (*domain.UserWallet, error) {
	wallet, err := ensureWalletFor(ctx, mintDeps{
		custody:      s.custody,
		wallets:      s.walletRepo,
		log:          s.log,
		providerType: s.providerType,
	}, user.ID, domain.WalletChainPrivate)
	if err != nil {
		return nil, err
	}

	s.log.Info("custody wallet ready", "userID", user.ID, "address", wallet.RaylsAddress)

	return wallet, nil
}

// prepareWalletForChain readies a freshly-created custody wallet for on-chain use: it funds the
// wallet for gas (non-gasless chains) and grants it the FACTORY_DEPLOYER role so the user can deploy
// tokens. Both steps are optional (nil dependency = skipped) and best-effort — a failure is logged
// but must not block login; the user can be funded/granted out of band. Funding runs first so the
// grant, if the grantor and the wallet are the same account, has gas.
func (s *provisioningService) prepareWalletForChain(ctx context.Context, address string) {
	if s.funder != nil {
		if err := s.funder.FundWallet(ctx, address); err != nil {
			s.log.Warn("failed to fund new custody wallet — it will need gas before it can transact",
				"address", address, "error", err)
		}
	}

	if s.granter != nil {
		roles := []string{domain.RoleFactoryDeployer, domain.RolePrivacyNodeOperator}
		if err := s.granter.GrantRoles(ctx, address, roles); err != nil {
			s.log.Warn(
				"failed to grant deploy/operator roles to new custody wallet — token deploy or login will be refused until granted out of band",
				"address",
				address,
				"roles",
				roles,
				"error",
				err,
			)
		}
	}
}
