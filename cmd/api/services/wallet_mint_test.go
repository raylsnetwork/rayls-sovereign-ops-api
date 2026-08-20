package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

func mintTestDeps(custody *testutil.FakeCustodyService, wallets *testutil.FakeUserWalletRepository) mintDeps {
	return mintDeps{
		custody:      custody,
		wallets:      wallets,
		log:          &testutil.StubLogger{},
		providerType: domain.CustodyProviderRaylsHSM,
	}
}

func TestMintWallet_WritesTheIntentRowBeforeMinting(t *testing.T) {
	// The whole point of the fix: the database is authoritative before the irreversible
	// HSM call, so a crash mid-mint leaves a row naming the user.
	wallets := &testutil.FakeUserWalletRepository{}
	custody := testutil.NewFakeCustodyService("0xabc", "ext-1")
	userID := uuid.New()

	var pendingAtMintTime int
	custody.OnCreate = func() {
		found, err := wallets.FindPendingByUserID(context.Background(), userID)
		require.NoError(t, err)
		pendingAtMintTime = len(found)
	}

	_, err := mintWallet(context.Background(), mintTestDeps(custody, wallets), userID, domain.WalletChainPrivate)

	require.NoError(t, err)
	assert.Equal(t, 1, pendingAtMintTime, "intent row must exist before the HSM is called")
}

func TestMintWallet_CompletesTheIntentInPlace(t *testing.T) {
	// Completion fills in the minted address rather than inserting a second row.
	wallets := &testutil.FakeUserWalletRepository{}
	custody := testutil.NewFakeCustodyService("0xabc", "ext-1")

	wallet, err := mintWallet(
		context.Background(),
		mintTestDeps(custody, wallets),
		uuid.New(),
		domain.WalletChainPrivate,
	)

	require.NoError(t, err)
	assert.Len(t, wallets.Wallets, 1)
	assert.Equal(t, "0xabc", wallets.Wallets[0].RaylsAddress)
	assert.Equal(t, "ext-1", wallets.Wallets[0].CustodyExternalID)
	assert.True(t, wallets.Wallets[0].IsActive)
	assert.False(t, wallets.Wallets[0].IsPending())
	assert.Equal(t, "0xabc", wallet.RaylsAddress)
}

func TestMintWallet_DeletesTheIntentWhenTheMintFails(t *testing.T) {
	// No key was minted, so the intent is meaningless and must not look stranded later.
	wallets := &testutil.FakeUserWalletRepository{}
	custody := testutil.NewFakeCustodyService("0xabc", "ext-1")
	custody.Err = errors.New("hsm down")

	_, err := mintWallet(context.Background(), mintTestDeps(custody, wallets), uuid.New(), domain.WalletChainPrivate)

	require.Error(t, err)
	assert.Empty(t, wallets.Wallets, "a failed mint must leave no rows behind")
}

func TestMintWallet_KeepsTheIntentWhenCompletionFails(t *testing.T) {
	// This is the one window that still orphans a key. The intent row must survive: it is
	// the only durable record that the key belongs to this user.
	wallets := &testutil.FakeUserWalletRepository{}
	wallets.CompletePendingErr = errors.New("db gone")
	custody := testutil.NewFakeCustodyService("0xabc", "ext-1")
	userID := uuid.New()

	_, err := mintWallet(context.Background(), mintTestDeps(custody, wallets), userID, domain.WalletChainPrivate)

	require.Error(t, err)
	pending, findErr := wallets.FindPendingByUserID(context.Background(), userID)
	require.NoError(t, findErr)
	assert.Len(t, pending, 1, "the intent must survive so the orphaned key stays traceable")
}

func TestMintWallet_DoesNotMintWhenTheIntentCannotBePersisted(t *testing.T) {
	// Fail before the side effect, not after — nothing to orphan.
	wallets := &testutil.FakeUserWalletRepository{}
	wallets.CreateErr = errors.New("db gone")
	custody := testutil.NewFakeCustodyService("0xabc", "ext-1")

	_, err := mintWallet(context.Background(), mintTestDeps(custody, wallets), uuid.New(), domain.WalletChainPrivate)

	require.Error(t, err)
	assert.Zero(t, custody.Calls, "the HSM must not be called if the intent did not persist")
}

func TestEnsureWalletFor_ReturnsTheExistingWalletWithoutMinting(t *testing.T) {
	// The common path on every login: one wallet per user, reused across chains.
	userID := uuid.New()
	wallets := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			Model:           domain.Model{ID: uuid.New()},
			UserID:          userID,
			RaylsAddress:    "0xexisting",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
			IsActive:        true,
		}},
	}
	custody := testutil.NewFakeCustodyService("0xnew", "ext-2")

	wallet, err := ensureWalletFor(
		context.Background(),
		mintTestDeps(custody, wallets),
		userID,
		domain.WalletChainPrivate,
	)

	require.NoError(t, err)
	assert.Equal(t, "0xexisting", wallet.RaylsAddress)
	assert.Zero(t, custody.Calls)
}

func TestEnsureWalletFor_DoesNotTreatAStrandedIntentAsAWallet(t *testing.T) {
	// A pending row is not a usable wallet: returning one would hand out a placeholder
	// address as if it were a real signer.
	userID := uuid.New()
	wallets := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			Model:        domain.Model{ID: uuid.New()},
			UserID:       userID,
			RaylsAddress: domain.PendingAddress(uuid.New()),
			IsActive:     false,
		}},
	}
	custody := testutil.NewFakeCustodyService("0xnew", "ext-2")

	wallet, err := ensureWalletFor(
		context.Background(),
		mintTestDeps(custody, wallets),
		userID,
		domain.WalletChainPrivate,
	)

	require.NoError(t, err)
	assert.Equal(t, "0xnew", wallet.RaylsAddress)
	assert.False(t, wallet.IsPending())
}

func TestPendingAddress_FitsTheColumnAndNeverLooksLikeAnAddress(t *testing.T) {
	// rayls_address is VARCHAR(42) and uniquely indexed; the placeholder has to fit and must
	// never collide with a real 0x address.
	addr := domain.PendingAddress(uuid.New())

	assert.LessOrEqual(t, len(addr), 42)
	assert.NotContains(t, addr, "0x")
	assert.True(t, (&domain.UserWallet{RaylsAddress: addr}).IsPending())
}

func TestPendingAddress_IsUniquePerAttempt(t *testing.T) {
	// Two concurrent mints for the same user must not collide on uq_rayls_address.
	assert.NotEqual(t, domain.PendingAddress(uuid.New()), domain.PendingAddress(uuid.New()))
}

func TestUserWallet_IsPendingIsFalseForARealAddress(t *testing.T) {
	assert.False(t, (&domain.UserWallet{RaylsAddress: "0xabc"}).IsPending())
}
