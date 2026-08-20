package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// mintDeps is the slice of a service needed to mint a custody wallet safely. Every call site
// (login provisioning, identity provisioning, onboarding, bootstrap) needs the same sequence,
// so it lives here once rather than being repeated with subtly different recovery behaviour.
type mintDeps struct {
	custody      core.CustodyService
	wallets      core.UserWalletRepository
	log          logger.Logger
	providerType domain.CustodyProviderType
}

// mintWallet creates an HSM wallet and persists it write-ahead: the intent row is committed
// BEFORE the mint, so the database is authoritative before the irreversible side effect.
//
// The HSM cannot help here. POST /api/wallet accepts only type/password/address_quantity —
// no user id, no idempotency key — and generates a fresh random key per call, and there is no
// delete API. So a naive mint-then-persist orphans a key on any failure in between, and the
// retry (keyed on a row that was never written) mints another orphan every time rather than
// converging.
//
// Ordering here instead:
//  1. write the intent row (pending placeholder address, is_active=false — invisible to
//     every existing query, so no schema change is needed to keep them correct),
//  2. mint,
//  3. fill the real address in and activate.
//
// A crash after (1) leaves a row that names the user, so a later sweep can reconcile it
// against the HSM. A failure at (2) means no key was ever minted, so the intent is deleted
// and nothing leaks. Only a crash between (2) and (3) still orphans a key — but now it
// leaves a durable pending row pointing at it, which is what makes recovery possible at all.
func mintWallet(
	ctx context.Context,
	deps mintDeps,
	userID uuid.UUID,
	chain domain.WalletChain,
) (*domain.UserWallet, error) {
	intent := &domain.UserWallet{
		UserID:          userID,
		RaylsAddress:    domain.PendingAddress(uuid.New()),
		CustodyProvider: deps.providerType,
		Chain:           chain,
		IsActive:        false,
	}
	if err := deps.wallets.Create(ctx, intent); err != nil {
		return nil, fmt.Errorf("record wallet intent: %w", err)
	}

	address, externalID, err := deps.custody.CreateWallet(ctx, userID)
	if err != nil {
		// The mint failed, so no key exists. Drop the intent so it is not mistaken for a
		// stranded one later. A failure to clean up is not fatal — the row is inert, and a
		// sweep finds no matching HSM key.
		if delErr := deps.wallets.DeletePending(ctx, intent.ID); delErr != nil {
			deps.log.Error("failed to delete wallet intent after a failed mint",
				"userID", userID, "intentID", intent.ID, "error", delErr)
		}
		return nil, fmt.Errorf("create custody wallet: %w", err)
	}

	if err := deps.wallets.CompletePending(ctx, intent.ID, address, externalID); err != nil {
		// The key exists in the HSM and the intent row survives to prove it belongs to this
		// user. Recoverable, unlike before: the address is recorded here and the intent is
		// discoverable via FindPendingByUserID.
		deps.log.Error("custody key minted but not attached to its intent row — recover via the pending intent",
			"userID", userID, "intentID", intent.ID, "address", address,
			"externalID", externalID, "error", err)
		return nil, fmt.Errorf("complete wallet intent: %w", err)
	}

	intent.RaylsAddress = address
	intent.CustodyExternalID = externalID
	intent.IsActive = true
	return intent, nil
}

// reportStrandedIntents surfaces intents left behind by an earlier crash.
//
// It deliberately does not try to reattach them. Reattachment would need to ask the HSM
// "which key did you mint for this intent?", and the custody API cannot answer: the only
// lookup is GET /api/wallet/address/{address}, which needs the very address that was lost.
// The mint carries no user id or correlation id, so nothing on the HSM side ties a key back
// to an intent. Closing this fully requires an idempotency key on POST /api/wallet.
//
// Surfacing beats silence: each row names a user and a time, which is enough for an operator
// to reconcile against the HSM's wallet list by hand.
func reportStrandedIntents(ctx context.Context, deps mintDeps, userID uuid.UUID) error {
	pending, err := deps.wallets.FindPendingByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("find pending wallet intents: %w", err)
	}
	if len(pending) == 0 {
		return nil
	}

	deps.log.Warn("stranded custody mint intents for this user",
		"userID", userID, "count", len(pending), "oldest", pending[0].CreatedAt,
		"detail", "an earlier mint died mid-flight; the HSM may hold keys nothing references")
	return nil
}

// ensureWalletFor returns the user's existing active wallet, or mints one. It is the single
// entry point every provisioning path should use.
//
// chain is always WalletChainPrivate today — login provisioning only ever needs the private
// wallet, and onboarding mints its public/private pair through buildWallet instead. It stays
// a parameter because the choice belongs to the caller, not to this function.
//
//nolint:unparam // see above: the parameter is deliberate, not dead.
func ensureWalletFor(
	ctx context.Context,
	deps mintDeps,
	userID uuid.UUID,
	chain domain.WalletChain,
) (*domain.UserWallet, error) {
	wallet, err := deps.wallets.FindByUserID(ctx, userID)
	if err == nil && wallet != nil && wallet.RaylsAddress != "" && !wallet.IsPending() {
		return wallet, nil
	}
	if err != nil && !errors.Is(err, core.ErrRecordNotFound) {
		return nil, fmt.Errorf("find wallet: %w", err)
	}

	if err := reportStrandedIntents(ctx, deps, userID); err != nil {
		return nil, err
	}

	return mintWallet(ctx, deps, userID, chain)
}
