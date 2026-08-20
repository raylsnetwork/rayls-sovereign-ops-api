package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

var _ core.BootstrapService = (*bootstrapService)(nil)

type bootstrapService struct {
	userRepo          core.UserRepository
	walletRepo        core.UserWalletRepository
	oauthProviderRepo core.UserOAuthProviderRepository
	txer              core.Transactor
	custody           core.CustodyService
	funder            core.WalletFunder // optional; nil on gasless chains (no funding needed)
	log               logger.Logger
}

func NewBootstrapService(
	userRepo core.UserRepository,
	walletRepo core.UserWalletRepository,
	oauthProviderRepo core.UserOAuthProviderRepository,
	txer core.Transactor,
	custody core.CustodyService,
	funder core.WalletFunder,
	log logger.Logger,
) core.BootstrapService {
	return &bootstrapService{
		userRepo:          userRepo,
		walletRepo:        walletRepo,
		oauthProviderRepo: oauthProviderRepo,
		txer:              txer,
		custody:           custody,
		funder:            funder,
		log:               log,
	}
}

// existingWalletFor resolves the active custody wallet of an already-bootstrapped user, for
// the address carried on BootstrapAlreadyCompletedError.
//
// Best-effort: a lookup failure (unknown email, no wallet on file) yields "" rather than an
// error, because the caller's real answer is still "already bootstrapped" — the address is
// extra information, not the result.
func (s *bootstrapService) existingWalletFor(ctx context.Context, email string) string {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return ""
	}
	wallet, err := s.walletRepo.FindByUserID(ctx, user.ID)
	if err != nil || wallet == nil {
		return ""
	}
	return wallet.RaylsAddress
}

// Bootstrap provisions the initial system administrator (SDD section 2a).
// Returns BootstrapAlreadyCompletedError if any user already exists.
func (s *bootstrapService) Bootstrap(ctx context.Context, email string) (string, error) {
	count, err := s.userRepo.Count(ctx)
	if err != nil {
		return "", core.NewInternalError("count users", err)
	}
	if count > 0 {
		// Already bootstrapped — but say WHICH wallet, so a caller provisioning another
		// chain can still grant that wallet the on-chain roles it needs there. Users are
		// shared across chains (one identity database, one custody wallet each); roles are
		// not. A bare 409 made every chain after the first come up with no roles granted.
		return "", &core.BootstrapAlreadyCompletedError{Address: s.existingWalletFor(ctx, email)}
	}

	user := &domain.User{
		Email:    email,
		IsActive: true,
		Status:   domain.UserStatusRoleAssigned,
	}

	// The user row must exist before anything can reference it: user_wallets.user_id is a
	// foreign key, so the mint intent cannot be written first the way it is elsewhere. (It
	// also means user.ID was still uuid.Nil when this passed it to CreateWallet before.)
	//
	// Both writes go in ONE transaction because the Count>0 guard above makes a half-done
	// bootstrap unrecoverable: a user row with no wallet intent still trips the guard, so
	// every retry returns BootstrapAlreadyCompletedError with an empty address and the
	// deployment can only be unblocked by deleting the row by hand.
	//
	// Intent first, then mint — so a crash in between leaves a row naming this user rather
	// than an HSM key nothing points at. See services/wallet_mint.go.
	intent := &domain.UserWallet{
		RaylsAddress:    domain.PendingAddress(uuid.New()),
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPrivate,
		IsActive:        false,
	}
	if txErr := s.txer.WithTransaction(ctx, func(txCtx context.Context) error {
		if createErr := s.userRepo.Create(txCtx, user); createErr != nil {
			return core.NewInternalError("create bootstrap admin", createErr)
		}
		intent.UserID = user.ID
		if intentErr := s.walletRepo.Create(txCtx, intent); intentErr != nil {
			return core.NewInternalError("record bootstrap admin wallet intent", intentErr)
		}
		return nil
	}); txErr != nil {
		return "", txErr
	}

	address, externalID, err := s.custody.CreateWallet(ctx, user.ID)
	if err != nil {
		if delErr := s.walletRepo.DeletePending(ctx, intent.ID); delErr != nil {
			s.log.Error("failed to delete wallet intent after a failed mint",
				"intentID", intent.ID, "error", delErr)
		}
		return "", core.NewInternalError("create custody wallet for admin", err)
	}

	if err := s.txer.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := s.walletRepo.CompletePending(txCtx, intent.ID, address, externalID); err != nil {
			return core.NewInternalError("create bootstrap admin wallet", err)
		}

		siweLink := &domain.UserOAuthProvider{
			UserID:        user.ID,
			Provider:      domain.OAuthProviderGoogle,
			OAuthID:       "",
			WalletAddress: address,
			Email:         user.Email,
		}
		if err := s.oauthProviderRepo.Create(txCtx, siweLink); err != nil {
			return core.NewInternalError("create bootstrap admin SIWE provider", err)
		}
		return nil
	}); err != nil {
		return "", fmt.Errorf("bootstrap transaction: %w", err)
	}

	// Non-gasless chain: fund the admin's custody wallet so it can pay gas (e.g. token deploy).
	// Best-effort — funding failure must not fail bootstrap. nil on gasless PNs.
	if s.funder != nil {
		if err := s.funder.FundWallet(ctx, address); err != nil {
			s.log.Warn("failed to fund bootstrap admin wallet — it will need gas before it can transact",
				"address", address, "error", err)
		}
	}

	s.log.Info("Bootstrap completed", "email", email, "address", address)
	return address, nil
}
