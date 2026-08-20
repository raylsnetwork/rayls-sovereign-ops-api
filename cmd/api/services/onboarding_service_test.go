package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// fakeSeqCustody returns a distinct address/externalID per CreateWallet call.
type fakeSeqCustody struct {
	addrs  []string
	extIDs []string
	calls  int
}

func (f *fakeSeqCustody) CreateWallet(_ context.Context, _ uuid.UUID) (string, string, error) {
	i := f.calls
	f.calls++
	return f.addrs[i], f.extIDs[i], nil
}

func (f *fakeSeqCustody) SignAndTransact(_ context.Context, _ []byte, _ string, _ string) (string, error) {
	return "", nil
}

// fakeOperatorResolver returns a fixed operator address (or an error).
type fakeOperatorResolver struct {
	addr string
	err  error
}

func (f *fakeOperatorResolver) Resolve(_ context.Context) (string, error) { return f.addr, f.err }

// fakeUserGov models on-chain RNUserGovernance state so EnsureUser idempotency can be asserted:
// CreateUser is only triggered when the on-chain user does not yet exist.
type fakeUserGov struct {
	hasUser         bool
	createUserCalls int
	addPairCalls    int
	pairs           [][2]string
	lastOperator    string
	pending         []core.OnChainAddressPair
	approved        []core.OnChainAddressPair
	allPending      []core.OnChainPendingGroup
	setStatusCalls  int
	lastStatus      domain.ApprovalStatus
	lastStatusPair  [2]string
	lastStatusUser  [32]byte
	addPairErr      error
}

func (f *fakeUserGov) ListPending(_ context.Context, _ [32]byte) ([]core.OnChainAddressPair, error) {
	return f.pending, nil
}

func (f *fakeUserGov) ListApproved(_ context.Context, _ [32]byte) ([]core.OnChainAddressPair, error) {
	return f.approved, nil
}

func (f *fakeUserGov) EnsureUser(_ context.Context, operator string, _ [32]byte) error {
	f.lastOperator = operator
	if !f.hasUser {
		f.createUserCalls++
		f.hasUser = true
	}
	return nil
}

func (f *fakeUserGov) AddAddressPair(
	_ context.Context,
	operator string,
	_ [32]byte,
	publicAddr, privateAddr string,
) (string, error) {
	f.lastOperator = operator
	if f.addPairErr != nil {
		return "", f.addPairErr
	}
	f.addPairCalls++
	f.pairs = append(f.pairs, [2]string{publicAddr, privateAddr})
	return "0xhash", nil
}

func (f *fakeUserGov) ListAllPending(_ context.Context) ([]core.OnChainPendingGroup, error) {
	return f.allPending, nil
}

func (f *fakeUserGov) SetApprovalStatus(
	_ context.Context,
	operator string,
	onChainUserID [32]byte,
	publicAddr, privateAddr string,
	status domain.ApprovalStatus,
) (string, error) {
	f.lastOperator = operator
	f.setStatusCalls++
	f.lastStatus = status
	f.lastStatusPair = [2]string{publicAddr, privateAddr}
	f.lastStatusUser = onChainUserID
	return "0xhash", nil
}

func TestOnboardingService_AddAddressPair_CreatesPairAndReturnsPending(t *testing.T) {
	// A call creates a fresh private+public HSM wallet pair, persists both, and registers the pair as pending.
	userID := uuid.New()
	custody := &fakeSeqCustody{
		addrs:  []string{"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222"},
		extIDs: []string{"ext-priv", "ext-pub"},
	}
	wallets := &testutil.FakeUserWalletRepository{}
	users := testutil.NewFakeUserRepository()
	resolver := &fakeOperatorResolver{addr: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	gov := &fakeUserGov{}

	svc := NewOnboardingService(
		custody,
		wallets,
		users,
		resolver,
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	pair, err := svc.AddAddressPair(context.Background(), userID)

	require.NoError(t, err)
	assert.Equal(t, domain.ApprovalStatusPending, pair.Status)
	assert.Equal(t, "0x1111111111111111111111111111111111111111", pair.PrivateChainAddress)
	assert.Equal(t, "0x2222222222222222222222222222222222222222", pair.PublicChainAddress)

	require.Len(t, wallets.Wallets, 2)
	assert.Equal(t, domain.WalletChainPrivate, wallets.Wallets[0].Chain)
	assert.Equal(t, domain.WalletChainPublic, wallets.Wallets[1].Chain)
	assert.Equal(t, domain.CustodyProviderRaylsHSM, wallets.Wallets[0].CustodyProvider)
	assert.Equal(t, "ext-priv", wallets.Wallets[0].CustodyExternalID)
	assert.True(t, wallets.Wallets[0].IsActive)

	assert.Equal(t, 1, gov.createUserCalls)
	assert.Equal(t, 1, gov.addPairCalls)
	assert.Equal(t, "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", gov.lastOperator)
	assert.Equal(t, [2]string{pair.PublicChainAddress, pair.PrivateChainAddress}, gov.pairs[0])
}

func TestOnboardingService_AddAddressPair_EnsureUserIdempotent(t *testing.T) {
	// Calling onboarding twice for the same user creates the on-chain user once but adds two pairs.
	userID := uuid.New()
	custody := &fakeSeqCustody{
		addrs: []string{
			"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222",
			"0x3333333333333333333333333333333333333333", "0x4444444444444444444444444444444444444444",
		},
		extIDs: []string{"a", "b", "c", "d"},
	}
	wallets := &testutil.FakeUserWalletRepository{}
	users := testutil.NewFakeUserRepository()
	resolver := &fakeOperatorResolver{addr: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	gov := &fakeUserGov{}

	svc := NewOnboardingService(
		custody,
		wallets,
		users,
		resolver,
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	_, err1 := svc.AddAddressPair(context.Background(), userID)
	_, err2 := svc.AddAddressPair(context.Background(), userID)

	require.NoError(t, err1)
	require.NoError(t, err2)
	assert.Equal(t, 1, gov.createUserCalls, "on-chain user must be created only once")
	assert.Equal(t, 2, gov.addPairCalls, "each call adds a fresh pair")
	assert.Len(t, wallets.Wallets, 4, "two wallets per call")
}

func TestOnboardingService_AddAddressPair_PersistsOnChainUserID(t *testing.T) {
	// Onboarding stores keccak256(user.ID) on the user row so admin discovery can reverse-map it.
	userID := uuid.New()
	custody := &fakeSeqCustody{
		addrs:  []string{"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222"},
		extIDs: []string{"a", "b"},
	}
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: userID}}}
	svc := NewOnboardingService(
		custody,
		&testutil.FakeUserWalletRepository{},
		users,
		&fakeOperatorResolver{addr: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		&fakeUserGov{},
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	_, err := svc.AddAddressPair(context.Background(), userID)

	require.NoError(t, err)
	expected := domain.OnChainUserID(userID)
	assert.Equal(t, expected[:], users.Users[0].OnChainUserID)
}

func TestOnboardingService_AddAddressPair_OnChainUserIDIdempotent(t *testing.T) {
	// Onboarding the same user twice leaves the stored hash unchanged and does not error.
	userID := uuid.New()
	custody := &fakeSeqCustody{
		addrs: []string{
			"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222",
			"0x3333333333333333333333333333333333333333", "0x4444444444444444444444444444444444444444",
		},
		extIDs: []string{"a", "b", "c", "d"},
	}
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: userID}}}
	svc := NewOnboardingService(
		custody,
		&testutil.FakeUserWalletRepository{},
		users,
		&fakeOperatorResolver{addr: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		&fakeUserGov{},
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	_, err1 := svc.AddAddressPair(context.Background(), userID)
	_, err2 := svc.AddAddressPair(context.Background(), userID)

	require.NoError(t, err1)
	require.NoError(t, err2)
	expected := domain.OnChainUserID(userID)
	assert.Equal(t, expected[:], users.Users[0].OnChainUserID)
}

func TestOnboardingService_AddAddressPair_NoOperatorSigner_PropagatesError(t *testing.T) {
	// When no operator signer can be resolved, the service surfaces NoOperatorSignerError (→ 503).
	custody := &fakeSeqCustody{
		addrs:  []string{"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222"},
		extIDs: []string{"a", "b"},
	}
	wallets := &testutil.FakeUserWalletRepository{}
	users := testutil.NewFakeUserRepository()
	resolver := &fakeOperatorResolver{err: &core.NoOperatorSignerError{}}
	gov := &fakeUserGov{}

	svc := NewOnboardingService(
		custody,
		wallets,
		users,
		resolver,
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	_, err := svc.AddAddressPair(context.Background(), uuid.New())

	var noSigner *core.NoOperatorSignerError
	require.ErrorAs(t, err, &noSigner)
	assert.Equal(t, 0, gov.addPairCalls, "no governance write when the operator is unavailable")
}

func TestOnboardingService_AddAddressPair_ChainFailure_PersistsNoWallets(t *testing.T) {
	// When the on-chain pair registration fails, no wallets are persisted, so a failed onboarding
	// leaves no orphaned active rows for the next retry to multiply.
	userID := uuid.New()
	custody := &fakeSeqCustody{
		addrs:  []string{"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222"},
		extIDs: []string{"a", "b"},
	}
	wallets := &testutil.FakeUserWalletRepository{}
	users := testutil.NewFakeUserRepository()
	resolver := &fakeOperatorResolver{addr: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	gov := &fakeUserGov{addPairErr: errors.New("chain reverted")}

	svc := NewOnboardingService(
		custody,
		wallets,
		users,
		resolver,
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	_, err := svc.AddAddressPair(context.Background(), userID)

	require.Error(t, err)
	assert.Empty(t, wallets.Wallets, "no wallets persisted when the on-chain write fails")
}

func TestOnboardingService_ListMine_NoFilterReturnsPendingAndApproved(t *testing.T) {
	// Omitting the status filter returns the union of pending and approved pairs (pending first).
	gov := &fakeUserGov{
		pending:  []core.OnChainAddressPair{{PublicChainAddress: "0xpub1", Status: domain.ApprovalStatusPending}},
		approved: []core.OnChainAddressPair{{PublicChainAddress: "0xpub2", Status: domain.ApprovalStatusApproved}},
	}
	svc := NewOnboardingService(
		&fakeSeqCustody{},
		&testutil.FakeUserWalletRepository{},
		testutil.NewFakeUserRepository(),
		&fakeOperatorResolver{},
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	pairs, err := svc.ListMine(context.Background(), uuid.New(), nil)

	require.NoError(t, err)
	require.Len(t, pairs, 2)
	assert.Equal(t, domain.ApprovalStatusPending, pairs[0].Status)
	assert.Equal(t, domain.ApprovalStatusApproved, pairs[1].Status)
}

func TestOnboardingService_ListMine_PendingFilterReturnsOnlyPending(t *testing.T) {
	// Filtering by pending returns only the pending pairs and never reads approved.
	gov := &fakeUserGov{
		pending:  []core.OnChainAddressPair{{PublicChainAddress: "0xpub1", Status: domain.ApprovalStatusPending}},
		approved: []core.OnChainAddressPair{{PublicChainAddress: "0xpub2", Status: domain.ApprovalStatusApproved}},
	}
	svc := NewOnboardingService(
		&fakeSeqCustody{},
		&testutil.FakeUserWalletRepository{},
		testutil.NewFakeUserRepository(),
		&fakeOperatorResolver{},
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	pending := domain.ApprovalStatusPending
	pairs, err := svc.ListMine(context.Background(), uuid.New(), &pending)

	require.NoError(t, err)
	require.Len(t, pairs, 1)
	assert.Equal(t, domain.ApprovalStatusPending, pairs[0].Status)
}

func TestOnboardingService_ListAllPending_ResolvesHashesToUUIDs(t *testing.T) {
	// Each pending group's keccak256 hash is reverse-mapped to the owning ops-api user UUID.
	userID := uuid.New()
	onChainID := domain.OnChainUserID(userID)
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: userID}, OnChainUserID: onChainID[:]}}
	gov := &fakeUserGov{allPending: []core.OnChainPendingGroup{
		{
			OnChainUserID: onChainID,
			AddressPairs: []core.OnChainAddressPair{
				{PublicChainAddress: "0xpub", Status: domain.ApprovalStatusPending},
			},
		},
	}}
	svc := NewOnboardingService(
		&fakeSeqCustody{},
		&testutil.FakeUserWalletRepository{},
		users,
		&fakeOperatorResolver{},
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	groups, err := svc.ListAllPending(context.Background())

	require.NoError(t, err)
	require.Len(t, groups, 1)
	assert.Equal(t, userID, groups[0].UserID)
	require.Len(t, groups[0].AddressPairs, 1)
	assert.Equal(t, "0xpub", groups[0].AddressPairs[0].PublicChainAddress)
}

func TestOnboardingService_ListAllPending_SkipsUnresolvedHash(t *testing.T) {
	// A pending group whose hash does not resolve to a known user is skipped, not fatal.
	knownID := uuid.New()
	knownHash := domain.OnChainUserID(knownID)
	unknownHash := domain.OnChainUserID(uuid.New())
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: knownID}, OnChainUserID: knownHash[:]}}
	gov := &fakeUserGov{allPending: []core.OnChainPendingGroup{
		{OnChainUserID: unknownHash, AddressPairs: []core.OnChainAddressPair{{PublicChainAddress: "0xorphan"}}},
		{OnChainUserID: knownHash, AddressPairs: []core.OnChainAddressPair{{PublicChainAddress: "0xpub"}}},
	}}
	svc := NewOnboardingService(
		&fakeSeqCustody{},
		&testutil.FakeUserWalletRepository{},
		users,
		&fakeOperatorResolver{},
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	groups, err := svc.ListAllPending(context.Background())

	require.NoError(t, err)
	require.Len(t, groups, 1)
	assert.Equal(t, knownID, groups[0].UserID)
}

func TestOnboardingService_SetApprovalStatus_Approve(t *testing.T) {
	// Approving resolves the user from the id, derives the on-chain id, and signs the status write.
	userID := uuid.New()
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: userID}}}
	gov := &fakeUserGov{}
	svc := NewOnboardingService(
		&fakeSeqCustody{},
		&testutil.FakeUserWalletRepository{},
		users,
		&fakeOperatorResolver{addr: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	err := svc.SetApprovalStatus(context.Background(), userID, "0xpub", "0xpriv", domain.ApprovalStatusApproved)

	require.NoError(t, err)
	assert.Equal(t, 1, gov.setStatusCalls)
	assert.Equal(t, domain.ApprovalStatusApproved, gov.lastStatus)
	assert.Equal(t, [2]string{"0xpub", "0xpriv"}, gov.lastStatusPair)
	assert.Equal(t, domain.OnChainUserID(userID), gov.lastStatusUser)
	assert.Equal(t, "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", gov.lastOperator)
}

func TestOnboardingService_SetApprovalStatus_Reject(t *testing.T) {
	// Rejecting passes the rejected status through to the governance write.
	userID := uuid.New()
	users := testutil.NewFakeUserRepository()
	users.Users = []domain.User{{Model: domain.Model{ID: userID}}}
	gov := &fakeUserGov{}
	svc := NewOnboardingService(
		&fakeSeqCustody{},
		&testutil.FakeUserWalletRepository{},
		users,
		&fakeOperatorResolver{addr: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	err := svc.SetApprovalStatus(context.Background(), userID, "0xpub", "0xpriv", domain.ApprovalStatusRejected)

	require.NoError(t, err)
	assert.Equal(t, domain.ApprovalStatusRejected, gov.lastStatus)
}

func TestOnboardingService_SetApprovalStatus_UnknownUser_ReturnsError(t *testing.T) {
	// An unknown user id fails before any governance write.
	gov := &fakeUserGov{}
	svc := NewOnboardingService(
		&fakeSeqCustody{},
		&testutil.FakeUserWalletRepository{},
		testutil.NewFakeUserRepository(),
		&fakeOperatorResolver{},
		gov,
		&testutil.FakeTransactor{},
		&testutil.StubLogger{},
	)

	err := svc.SetApprovalStatus(context.Background(), uuid.New(), "0xpub", "0xpriv", domain.ApprovalStatusApproved)

	require.Error(t, err)
	assert.Equal(t, 0, gov.setStatusCalls)
}
