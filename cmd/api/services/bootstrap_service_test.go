package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

func newBootstrapService(
	users *testutil.FakeUserRepository,
	wallets *testutil.FakeUserWalletRepository,
) core.BootstrapService {
	return NewBootstrapService(
		users,
		wallets,
		&testutil.FakeUserOAuthProviderRepository{},
		&testutil.FakeTransactor{},
		testutil.NewFakeCustodyService("0xabc", "ext-1"),
		nil,
		&testutil.StubLogger{},
	)
}

func TestBootstrapService_Bootstrap_CreatesAdminOnEmptyDeployment(t *testing.T) {
	// With no users yet, bootstrap mints a custody wallet and returns its address.
	users := testutil.NewFakeUserRepository()
	wallets := &testutil.FakeUserWalletRepository{}

	svc := newBootstrapService(users, wallets)

	address, err := svc.Bootstrap(context.Background(), "admin@example.com")

	require.NoError(t, err)
	assert.Equal(t, "0xabc", address)
}

func TestBootstrapService_Bootstrap_AlreadyCompletedReturnsExistingWallet(t *testing.T) {
	// Accounts are shared across chains, so every chain after the first conflicts — but the
	// error must carry the existing wallet so the caller can still grant per-chain roles.
	userID := uuid.New()
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: userID}, Email: "admin@example.com", IsActive: true}}
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{UserID: userID, RaylsAddress: "0xdef", IsActive: true},
	}}

	svc := newBootstrapService(users, wallets)

	address, err := svc.Bootstrap(context.Background(), "admin@example.com")

	assert.Empty(t, address)
	var alreadyDone *core.BootstrapAlreadyCompletedError
	require.ErrorAs(t, err, &alreadyDone)
	assert.Equal(t, "0xdef", alreadyDone.Address)
}

func TestBootstrapService_Bootstrap_AlreadyCompletedWithNoWalletOnFile(t *testing.T) {
	// The user exists but has no wallet: still a conflict, just with nothing to grant to.
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: uuid.New()}, Email: "admin@example.com", IsActive: true}}

	svc := newBootstrapService(users, &testutil.FakeUserWalletRepository{})

	_, err := svc.Bootstrap(context.Background(), "admin@example.com")

	var alreadyDone *core.BootstrapAlreadyCompletedError
	require.ErrorAs(t, err, &alreadyDone)
	assert.Empty(t, alreadyDone.Address)
}

// newBootstrapServiceTx is newBootstrapService with a caller-supplied transactor, for the
// rollback cases below.
func newBootstrapServiceTx(
	users *testutil.FakeUserRepository,
	wallets *testutil.FakeUserWalletRepository,
	txer core.Transactor,
) core.BootstrapService {
	return NewBootstrapService(
		users,
		wallets,
		&testutil.FakeUserOAuthProviderRepository{},
		txer,
		testutil.NewFakeCustodyService("0xabc", "ext-1"),
		nil,
		&testutil.StubLogger{},
	)
}

// rollbackTransactorFor wires a RollbackTransactor that restores the user slice, which is the
// row a failed wallet-intent write would otherwise strand.
func rollbackTransactorFor(users *testutil.FakeUserRepository) *testutil.RollbackTransactor {
	var saved []domain.User
	return &testutil.RollbackTransactor{
		Snapshot: func() { saved = append([]domain.User(nil), users.Users...) },
		Restore:  func() { users.Users = saved },
	}
}

func TestBootstrapService_Bootstrap_WalletIntentFailureRollsBackTheUser(t *testing.T) {
	// A failed wallet intent must take the user row with it: the Count>0 guard treats a
	// stranded user as "already bootstrapped", which would block every retry.
	users := testutil.NewFakeUserRepository()
	wallets := &testutil.FakeUserWalletRepository{CreateErr: errors.New("intent write failed")}
	txer := rollbackTransactorFor(users)

	svc := newBootstrapServiceTx(users, wallets, txer)

	_, err := svc.Bootstrap(context.Background(), "admin@example.com")

	require.Error(t, err)
	assert.True(t, txer.Rollback, "the wallet intent failure must roll the transaction back")
	assert.Empty(t, users.Users, "no user row may survive a failed bootstrap")
}

func TestBootstrapService_Bootstrap_RetriesAfterAWalletIntentFailure(t *testing.T) {
	// The real regression: once the intent write recovers, a retry completes normally rather
	// than returning BootstrapAlreadyCompletedError{Address: ""} forever.
	users := testutil.NewFakeUserRepository()
	wallets := &testutil.FakeUserWalletRepository{CreateErr: errors.New("intent write failed")}
	txer := rollbackTransactorFor(users)

	svc := newBootstrapServiceTx(users, wallets, txer)

	_, firstErr := svc.Bootstrap(context.Background(), "admin@example.com")
	require.Error(t, firstErr)

	wallets.CreateErr = nil
	address, err := svc.Bootstrap(context.Background(), "admin@example.com")

	require.NoError(t, err)
	assert.Equal(t, "0xabc", address)
	var alreadyDone *core.BootstrapAlreadyCompletedError
	assert.NotErrorAs(t, err, &alreadyDone, "the retry must not be blocked by a stranded user row")
}
