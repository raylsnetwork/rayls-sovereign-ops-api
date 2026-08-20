package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// --- fakes for the four AM repositories (only the read methods carry data) ---

type fakeMemberRepo struct {
	byAccount map[string][]*domain.AccessManagerRoleMember
}

func (f *fakeMemberRepo) Upsert(context.Context, *domain.AccessManagerRoleMember) error { return nil }
func (f *fakeMemberRepo) Revoke(context.Context, uint64, string) error                  { return nil }
func (f *fakeMemberRepo) ListByRole(context.Context, uint64, bool) ([]*domain.AccessManagerRoleMember, error) {
	return nil, nil
}

func (f *fakeMemberRepo) ListByAccount(_ context.Context, account string) ([]*domain.AccessManagerRoleMember, error) {
	return f.byAccount[account], nil
}

func (f *fakeMemberRepo) FindActiveAccountWithAllRoles(context.Context, []uint64) (string, error) {
	return "", core.ErrRecordNotFound
}

type fakeScopedRepo struct {
	byAccount map[string][]*domain.AccessManagerContractScopedRoleMember
}

func (f *fakeScopedRepo) Upsert(context.Context, *domain.AccessManagerContractScopedRoleMember) error {
	return nil
}
func (f *fakeScopedRepo) Revoke(context.Context, uint64, string, string) error { return nil }

func (f *fakeScopedRepo) ListByAccount(
	_ context.Context,
	account string,
) ([]*domain.AccessManagerContractScopedRoleMember, error) {
	return f.byAccount[account], nil
}

type fakeFnPermRepo struct {
	byContract map[string][]*domain.AccessManagerFunctionPermission
}

func (f *fakeFnPermRepo) Upsert(context.Context, *domain.AccessManagerFunctionPermission) error {
	return nil
}
func (f *fakeFnPermRepo) Remove(context.Context, string, string, uint64) error { return nil }

func (f *fakeFnPermRepo) ListByContract(
	_ context.Context,
	contract string,
) ([]*domain.AccessManagerFunctionPermission, error) {
	return f.byContract[contract], nil
}

type fakeManagedContractRepo struct {
	byAddress map[string]*domain.AccessManagerManagedContract
}

func (f *fakeManagedContractRepo) Upsert(context.Context, *domain.AccessManagerManagedContract) error {
	return nil
}
func (f *fakeManagedContractRepo) SetPaused(context.Context, string, bool) error { return nil }
func (f *fakeManagedContractRepo) List(context.Context) ([]*domain.AccessManagerManagedContract, error) {
	return nil, nil
}

func (f *fakeManagedContractRepo) FindByAddress(
	_ context.Context,
	addr string,
) (*domain.AccessManagerManagedContract, error) {
	if c, ok := f.byAddress[addr]; ok {
		return c, nil
	}
	return nil, core.ErrRecordNotFound
}

// --- fakes for the pause path (resolvePause reads the token, then the contract) ---

type fakeTokenRepo struct {
	byAddress map[string]*domain.Token
}

func (f *fakeTokenRepo) Upsert(context.Context, *domain.Token) error { return nil }
func (f *fakeTokenRepo) UpdateSupplyAndHolders(context.Context, string, string, int) error {
	return nil
}

func (f *fakeTokenRepo) List(context.Context, core.TokenFilter) ([]*domain.Token, int64, error) {
	return nil, 0, nil
}

func (f *fakeTokenRepo) FindByAddress(_ context.Context, address string) (*domain.Token, error) {
	if t, ok := f.byAddress[address]; ok {
		return t, nil
	}
	return nil, core.ErrRecordNotFound
}

// fakeTokenActions returns canned pauser/paused reads; the *Err fields simulate a failed
// chain read, which resolvePause must tolerate.
type fakeTokenActions struct {
	pauser    string
	pauserErr error
	paused    bool
	pausedErr error
}

func (f *fakeTokenActions) Mint(
	context.Context, string, string, domain.ErcStandard, core.MintInput,
) (string, error) {
	return "", nil
}

func (f *fakeTokenActions) Burn(
	context.Context, string, string, domain.ErcStandard, core.BurnInput,
) (string, error) {
	return "", nil
}

func (f *fakeTokenActions) SetPaused(context.Context, string, string, bool) (string, error) {
	return "", nil
}

func (f *fakeTokenActions) Pauser(context.Context, string) (string, error) {
	return f.pauser, f.pauserErr
}

func (f *fakeTokenActions) IsPaused(context.Context, string) (bool, error) {
	return f.paused, f.pausedErr
}

func TestTokenPermissionService_IntersectsRolesAndPermissions(t *testing.T) {
	// The user can call functions whose required role it holds globally or scoped to this contract.
	const wallet, token = "0xwallet", "0xtoken"
	mintSel := functionSelector("mint(address,uint256)")
	burnSel := functionSelector("burn(address,uint256)")
	submitSel := functionSelector("submitTokenUpdate(uint8,uint256)")

	members := &fakeMemberRepo{byAccount: map[string][]*domain.AccessManagerRoleMember{
		wallet: {{RoleID: 5, Account: wallet, IsActive: true}},
	}}
	scoped := &fakeScopedRepo{byAccount: map[string][]*domain.AccessManagerContractScopedRoleMember{
		wallet: {
			{RoleID: 8, Account: wallet, ContractAddress: token, IsActive: true},     // counts (this contract)
			{RoleID: 7, Account: wallet, ContractAddress: "0xother", IsActive: true}, // ignored (other contract)
		},
	}}
	fnPerms := &fakeFnPermRepo{byContract: map[string][]*domain.AccessManagerFunctionPermission{
		token: {
			{ContractAddress: token, Selector: mintSel, RoleID: 5, IsActive: true},   // via global role 5
			{ContractAddress: token, Selector: burnSel, RoleID: 8, IsActive: true},   // via scoped role 8
			{ContractAddress: token, Selector: submitSel, RoleID: 9, IsActive: true}, // role 9 not held
			{ContractAddress: token, Selector: mintSel, RoleID: 99, IsActive: true},  // dup selector, role not held
		},
	}}
	contracts := &fakeManagedContractRepo{byAddress: map[string]*domain.AccessManagerManagedContract{
		token: {ContractAddress: token, IsPaused: true},
	}}

	// nil tokens/actions: these cases cover the AM-derived permissions only, and pause needs
	// a live chain read (see resolvePause) that a nil reader correctly skips.
	svc := NewTokenPermissionService(members, scoped, fnPerms, contracts, nil, nil)

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.True(t, res.CanMint)
	assert.True(t, res.CanBurn)
	assert.False(t, res.CanSubmitTokenUpdate)
	assert.True(t, res.IsPaused)
	assert.Len(t, res.Functions, 2) // mint + burn, deduped; submit excluded
}

func TestTokenPermissionService_NoRolesReturnsEmpty(t *testing.T) {
	// A wallet with no roles can call nothing.
	const wallet, token = "0xwallet", "0xtoken"
	svc := NewTokenPermissionService(
		&fakeMemberRepo{}, &fakeScopedRepo{},
		&fakeFnPermRepo{byContract: map[string][]*domain.AccessManagerFunctionPermission{
			token: {
				{
					ContractAddress: token,
					Selector:        functionSelector("mint(address,uint256)"),
					RoleID:          5,
					IsActive:        true,
				},
			},
		}},
		&fakeManagedContractRepo{},
		nil, nil,
	)

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.False(t, res.CanMint)
	assert.Empty(t, res.Functions)
	assert.False(t, res.IsPaused) // contract not managed -> ErrRecordNotFound tolerated
}

// newPauseService wires a permission service holding no AM roles, so the assertions below
// isolate the pause fields resolvePause fills.
func newPauseService(token *domain.Token, actions core.TokenActionService) *TokenPermissionService {
	repo := &fakeTokenRepo{}
	if token != nil {
		repo.byAddress = map[string]*domain.Token{token.ContractAddress: token}
	}
	return NewTokenPermissionService(
		&fakeMemberRepo{}, &fakeScopedRepo{}, &fakeFnPermRepo{}, &fakeManagedContractRepo{},
		repo, actions,
	)
}

func TestTokenPermissionService_ResolvePause_PauserWalletCanPause(t *testing.T) {
	// A stablecoin whose on-chain pauser is the calling wallet reports CanPause and the live paused flag.
	const wallet, token = "0xwallet", "0xtoken"

	svc := newPauseService(
		&domain.Token{ContractAddress: token, ErcStandard: domain.ErcStandardStableCoin},
		&fakeTokenActions{pauser: wallet, paused: true},
	)

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.True(t, res.CanPause)
	assert.True(t, res.IsTokenPaused)
}

func TestTokenPermissionService_ResolvePause_PauserComparisonIsCaseInsensitive(t *testing.T) {
	// The pauser address is normalized before comparison, so checksummed casing still matches.
	const wallet, token = "0xabcdef", "0xtoken"

	svc := newPauseService(
		&domain.Token{ContractAddress: token, ErcStandard: domain.ErcStandardStableCoin},
		&fakeTokenActions{pauser: "0xABCDEF"},
	)

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.True(t, res.CanPause)
}

func TestTokenPermissionService_ResolvePause_OtherWalletCannotPause(t *testing.T) {
	// A wallet that is not the contract's pauser cannot pause, but still sees the paused flag.
	const wallet, token = "0xwallet", "0xtoken"

	svc := newPauseService(
		&domain.Token{ContractAddress: token, ErcStandard: domain.ErcStandardStableCoin},
		&fakeTokenActions{pauser: "0xsomeoneelse", paused: true},
	)

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.False(t, res.CanPause)
	assert.True(t, res.IsTokenPaused)
}

func TestTokenPermissionService_ResolvePause_SkipsNonStablecoin(t *testing.T) {
	// Only stablecoins have a pause function, so a plain ERC-20 is never queried for one.
	const wallet, token = "0xwallet", "0xtoken"

	svc := newPauseService(
		&domain.Token{ContractAddress: token, ErcStandard: domain.ErcStandardERC20},
		&fakeTokenActions{pauser: wallet, paused: true},
	)

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.False(t, res.CanPause)
	assert.False(t, res.IsTokenPaused)
}

func TestTokenPermissionService_ResolvePause_UnknownTokenLeavesPauseFalse(t *testing.T) {
	// A token absent from the repository yields no pause information rather than an error.
	const wallet, token = "0xwallet", "0xtoken"

	svc := newPauseService(nil, &fakeTokenActions{pauser: wallet, paused: true})

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.False(t, res.CanPause)
	assert.False(t, res.IsTokenPaused)
}

func TestTokenPermissionService_ResolvePause_FailedChainReadIsTolerated(t *testing.T) {
	// A failed on-chain read leaves the pause fields false instead of failing the whole call.
	const wallet, token = "0xwallet", "0xtoken"

	svc := newPauseService(
		&domain.Token{ContractAddress: token, ErcStandard: domain.ErcStandardStableCoin},
		&fakeTokenActions{
			pauserErr: errors.New("rpc unavailable"),
			pausedErr: errors.New("rpc unavailable"),
		},
	)

	res, err := svc.GetTokenPermissions(context.Background(), token, wallet)

	require.NoError(t, err)
	assert.False(t, res.CanPause)
	assert.False(t, res.IsTokenPaused)
}
