package services

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

var _ core.OnboardingService = (*onboardingService)(nil)

// onboardingService orchestrates self-service onboarding it creates a fresh HSM
// wallet pair, persists both wallets, and registers the pair on-chain via operator-authority
// governance writes signed by the resolved operator wallet. Each call creates a fresh pair, so a
// user accumulates many pairs over repeated calls.
type onboardingService struct {
	custody          core.CustodyService
	wallets          core.UserWalletRepository
	users            core.UserRepository
	operatorResolver core.OperatorSignerResolver
	gov              core.UserGovernanceService
	txer             core.Transactor
	log              logger.Logger
}

func NewOnboardingService(
	custody core.CustodyService,
	wallets core.UserWalletRepository,
	users core.UserRepository,
	operatorResolver core.OperatorSignerResolver,
	gov core.UserGovernanceService,
	txer core.Transactor,
	log logger.Logger,
) core.OnboardingService {
	return &onboardingService{
		custody:          custody,
		wallets:          wallets,
		users:            users,
		operatorResolver: operatorResolver,
		gov:              gov,
		txer:             txer,
		log:              log,
	}
}

// AddAddressPair creates a fresh HSM wallet pair for the user, registers the pair on-chain, and only
// then activates both wallets. The on-chain writes are ordered ahead of activation so a chain
// failure cannot leave active user_wallets rows (which would otherwise pollute login/signer
// resolution and be multiplied by every retry). Returns the created pair in pending status.
//
// Each mint is preceded by an intent row (see buildWallet). CreateWallet has no idempotency key
// and the HSM has no delete API, so a crash after minting still leaves an unreferenced key — but
// the intent row records that it belongs to this user, which is what makes it recoverable rather
// than invisible. Resolving the operator and ensuring the on-chain user happen before minting to
// keep that window as small as possible.
func (s *onboardingService) AddAddressPair(ctx context.Context, userID uuid.UUID) (*core.OnChainAddressPair, error) {
	onChainUserID := domain.OnChainUserID(userID)

	// Persist the hash before any on-chain write: a DB failure aborts cleanly with no chain mutation,
	// and every hash the contract can later return is guaranteed reverse-resolvable to this user.
	if err := s.users.SetOnChainUserID(ctx, userID, onChainUserID[:]); err != nil {
		return nil, fmt.Errorf("persist on-chain user id: %w", err)
	}

	operatorAddress, err := s.operatorResolver.Resolve(ctx)
	if err != nil {
		return nil, fmt.Errorf("resolve operator signer: %w", err)
	}

	if err := s.gov.EnsureUser(ctx, operatorAddress, onChainUserID); err != nil {
		return nil, fmt.Errorf("ensure on-chain user: %w", err)
	}

	// Mint the wallet pair but hold off persisting until the on-chain pair is registered.
	privateWallet, err := s.buildWallet(ctx, userID, domain.WalletChainPrivate)
	if err != nil {
		return nil, fmt.Errorf("create private-chain wallet: %w", err)
	}

	publicWallet, err := s.buildWallet(ctx, userID, domain.WalletChainPublic)
	if err != nil {
		return nil, fmt.Errorf("create public-chain wallet: %w", err)
	}

	if _, err := s.gov.AddAddressPair(
		ctx,
		operatorAddress,
		onChainUserID,
		publicWallet.RaylsAddress,
		privateWallet.RaylsAddress,
	); err != nil {
		// The chain rejected the pair, so no active wallet should exist for it. The keys are
		// already minted and unreclaimable, but the intent rows are not useful here.
		s.discardIntents(ctx, privateWallet, publicWallet)
		return nil, fmt.Errorf("add address pair: %w", err)
	}

	// Activate both wallets atomically only after the on-chain pair exists, so a partial write
	// never leaves a single dangling wallet. The rows already exist as intents; this fills in
	// the minted addresses and flips them active.
	if err := s.txer.WithTransaction(ctx, func(ctx context.Context) error {
		if err := s.wallets.CompletePending(
			ctx, privateWallet.ID, privateWallet.RaylsAddress, privateWallet.CustodyExternalID,
		); err != nil {
			return fmt.Errorf("persist private-chain wallet: %w", err)
		}
		if err := s.wallets.CompletePending(
			ctx, publicWallet.ID, publicWallet.RaylsAddress, publicWallet.CustodyExternalID,
		); err != nil {
			return fmt.Errorf("persist public-chain wallet: %w", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	privateWallet.IsActive = true
	publicWallet.IsActive = true

	s.log.Info(
		"address pair created",
		"userID",
		userID,
		"public",
		publicWallet.RaylsAddress,
		"private",
		privateWallet.RaylsAddress,
	)

	return &core.OnChainAddressPair{
		PublicChainAddress:  publicWallet.RaylsAddress,
		PrivateChainAddress: privateWallet.RaylsAddress,
		Status:              domain.ApprovalStatusPending,
		CreatedAt:           time.Now().UTC(),
	}, nil
}

// ListMine returns the caller's on-chain address pairs. A nil status returns all pairs (the union
// of pending and approved); a status of pending returns only the pending pairs.
func (s *onboardingService) ListMine(
	ctx context.Context,
	userID uuid.UUID,
	status *domain.ApprovalStatus,
) ([]core.OnChainAddressPair, error) {
	onChainUserID := domain.OnChainUserID(userID)

	if status != nil && *status == domain.ApprovalStatusPending {
		pending, err := s.gov.ListPending(ctx, onChainUserID)
		if err != nil {
			return nil, fmt.Errorf("list pending address pairs: %w", err)
		}
		return pending, nil
	}

	pending, err := s.gov.ListPending(ctx, onChainUserID)
	if err != nil {
		return nil, fmt.Errorf("list pending address pairs: %w", err)
	}
	approved, err := s.gov.ListApproved(ctx, onChainUserID)
	if err != nil {
		return nil, fmt.Errorf("list approved address pairs: %w", err)
	}

	pairs := make([]core.OnChainAddressPair, 0, len(pending)+len(approved))
	pairs = append(pairs, pending...)
	pairs = append(pairs, approved...)
	return pairs, nil
}

// ListAllPending reads every pending pair across all on-chain users, then reverse-maps each group's
// keccak256 hash to an ops-api UUID in one query. A hash that does not resolve to a known user (e.g.
// a user onboarded before the on_chain_user_id column existed) is skipped and logged, rather than
// failing the whole request.
func (s *onboardingService) ListAllPending(ctx context.Context) ([]core.PendingUserAddressPairs, error) {
	groups, err := s.gov.ListAllPending(ctx)
	if err != nil {
		return nil, fmt.Errorf("list all pending address pairs: %w", err)
	}

	hashes := make([][]byte, 0, len(groups))
	for _, g := range groups {
		hash := g.OnChainUserID
		hashes = append(hashes, hash[:])
	}

	resolved, err := s.users.FindByOnChainUserIDs(ctx, hashes)
	if err != nil {
		return nil, fmt.Errorf("reverse-map on-chain user ids: %w", err)
	}

	result := make([]core.PendingUserAddressPairs, 0, len(groups))
	for _, g := range groups {
		key := hex.EncodeToString(g.OnChainUserID[:])
		userID, ok := resolved[key]
		if !ok {
			s.log.Warn("pending address-pair hash did not resolve to a known user; skipping", "hash", key)
			continue
		}
		result = append(result, core.PendingUserAddressPairs{
			UserID:       userID,
			AddressPairs: g.AddressPairs,
		})
	}
	return result, nil
}

// SetApprovalStatus approves or rejects an address pair for the user identified by userID. The
// identity comes from userID (never the request body): the user is resolved, their on-chain user ID
// derived, the operator wallet resolved, and the operator-signed status write submitted.
func (s *onboardingService) SetApprovalStatus(
	ctx context.Context,
	userID uuid.UUID,
	publicAddr, privateAddr string,
	status domain.ApprovalStatus,
) error {
	if _, err := s.users.FindByID(ctx, userID); err != nil {
		return fmt.Errorf("find user: %w", err)
	}

	onChainUserID := domain.OnChainUserID(userID)

	operatorAddress, err := s.operatorResolver.Resolve(ctx)
	if err != nil {
		return fmt.Errorf("resolve operator signer: %w", err)
	}

	if _, err := s.gov.SetApprovalStatus(
		ctx,
		operatorAddress,
		onChainUserID,
		publicAddr,
		privateAddr,
		status,
	); err != nil {
		return fmt.Errorf("set approval status: %w", err)
	}

	s.log.Info("address pair approval status set", "userID", userID, "status", status.Label())
	return nil
}

// buildWallet mints a fresh HSM wallet via custody and returns the corresponding UserWallet for the
// given chain, with a normalized address. It does not persist the wallet — the caller persists it
// only after the on-chain pair is registered.
// buildWallet records a mint intent, mints, and returns the intent row with the real address
// filled in — but NOT yet completed in the database. The caller activates both wallets inside
// its transaction, after the on-chain pair is registered.
//
// Splitting it this way keeps the existing ordering guarantee (no active wallet row exists
// until the chain agrees) while still writing the intent ahead of the mint, so a crash leaves
// a durable pointer to the key instead of an untraceable orphan.
func (s *onboardingService) buildWallet(
	ctx context.Context,
	userID uuid.UUID,
	chain domain.WalletChain,
) (*domain.UserWallet, error) {
	intent := &domain.UserWallet{
		UserID:          userID,
		RaylsAddress:    domain.PendingAddress(uuid.New()),
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           chain,
		IsActive:        false,
	}
	if err := s.wallets.Create(ctx, intent); err != nil {
		return nil, fmt.Errorf("record wallet intent: %w", err)
	}

	address, externalID, err := s.custody.CreateWallet(ctx, userID)
	if err != nil {
		if delErr := s.wallets.DeletePending(ctx, intent.ID); delErr != nil {
			s.log.Error("failed to delete wallet intent after a failed mint",
				"userID", userID, "intentID", intent.ID, "error", delErr)
		}
		return nil, fmt.Errorf("create custody wallet: %w", err)
	}

	intent.RaylsAddress = domain.NormalizeAddress(address)
	intent.CustodyExternalID = externalID
	return intent, nil
}

// discardIntents drops intent rows for a pair that never reached the chain. Best-effort: the
// rows are inert (is_active=false, placeholder address), and any that survive are reported
// on the user's next provisioning pass.
func (s *onboardingService) discardIntents(ctx context.Context, wallets ...*domain.UserWallet) {
	for _, w := range wallets {
		if w == nil {
			continue
		}
		if err := s.wallets.DeletePending(ctx, w.ID); err != nil {
			s.log.Error("failed to discard wallet intent",
				"intentID", w.ID, "address", w.RaylsAddress, "error", err)
		}
	}
}
