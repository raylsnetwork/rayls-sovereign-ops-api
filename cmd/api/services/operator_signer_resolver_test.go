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

const (
	operatorRoleID uint64 = 7
	bankRoleID     uint64 = 8
)

func newResolver(
	roles *testutil.FakeAccessManagerRoleRepository,
	members *testutil.FakeAccessManagerRoleMemberRepository,
	wallets *testutil.FakeUserWalletRepository,
) core.OperatorSignerResolver {
	return NewOperatorSignerResolver(roles, members, wallets, &testutil.StubLogger{})
}

// bothRoles seeds both the PRIVACY_NODE_OPERATOR and BANK_EMPLOYEE roles.
func bothRoles() *testutil.FakeAccessManagerRoleRepository {
	return &testutil.FakeAccessManagerRoleRepository{Roles: []*domain.AccessManagerRole{
		{RoleID: operatorRoleID, Label: "PRIVACY_NODE_OPERATOR"},
		{RoleID: bankRoleID, Label: "BANK_EMPLOYEE"},
	}}
}

// dualMember returns active memberships in both roles for the same account.
func dualMember(account string) []*domain.AccessManagerRoleMember {
	return []*domain.AccessManagerRoleMember{
		{RoleID: operatorRoleID, Account: account, IsActive: true},
		{RoleID: bankRoleID, Account: account, IsActive: true},
	}
}

func TestOperatorSignerResolver_Resolve_ReturnsHSMDualRoleWallet(t *testing.T) {
	// The account holding both roles, backed by an HSM wallet, is resolved.
	roles := &testutil.FakeAccessManagerRoleRepository{Roles: []*domain.AccessManagerRole{
		{RoleID: operatorRoleID, Label: "Privacy Node Operator"},
		{RoleID: bankRoleID, Label: "Bank Employee"},
	}}
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: dualMember("0xOPERATOR")}
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{
			Model:           domain.Model{ID: uuid.New()},
			RaylsAddress:    "0xOPERATOR",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
		},
	}}

	addr, err := newResolver(roles, members, wallets).Resolve(context.Background())

	require.NoError(t, err)
	assert.Equal(t, "0xOPERATOR", addr)
}

func TestOperatorSignerResolver_Resolve_FailsClosedWhenNoDualRoleMember(t *testing.T) {
	// A member holding only the operator role (not BANK_EMPLOYEE) yields no signer.
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: []*domain.AccessManagerRoleMember{
		{RoleID: operatorRoleID, Account: "0xONLYOP", IsActive: true},
	}}
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{
			Model:           domain.Model{ID: uuid.New()},
			RaylsAddress:    "0xONLYOP",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
		},
	}}

	_, err := newResolver(bothRoles(), members, wallets).Resolve(context.Background())

	var noSigner *core.NoOperatorSignerError
	require.Error(t, err)
	assert.True(t, errors.As(err, &noSigner))
}

func TestOperatorSignerResolver_Resolve_FailsClosedWhenNoMembers(t *testing.T) {
	// Roles present but with no members fails closed with NoOperatorSignerError.
	members := &testutil.FakeAccessManagerRoleMemberRepository{}
	wallets := &testutil.FakeUserWalletRepository{}

	_, err := newResolver(bothRoles(), members, wallets).Resolve(context.Background())

	var noSigner *core.NoOperatorSignerError
	require.Error(t, err)
	assert.True(t, errors.As(err, &noSigner))
}

func TestOperatorSignerResolver_Resolve_FailsClosedWhenOperatorRoleMissing(t *testing.T) {
	// When the operator role is absent from the Access Manager, resolution fails closed.
	roles := &testutil.FakeAccessManagerRoleRepository{Roles: []*domain.AccessManagerRole{
		{RoleID: bankRoleID, Label: "BANK_EMPLOYEE"},
	}}
	members := &testutil.FakeAccessManagerRoleMemberRepository{}
	wallets := &testutil.FakeUserWalletRepository{}

	_, err := newResolver(roles, members, wallets).Resolve(context.Background())

	var noSigner *core.NoOperatorSignerError
	require.Error(t, err)
	assert.True(t, errors.As(err, &noSigner))
}

func TestOperatorSignerResolver_Resolve_FailsClosedWhenBankRoleMissing(t *testing.T) {
	// When BANK_EMPLOYEE is absent, resolution fails closed even with a valid operator member.
	roles := &testutil.FakeAccessManagerRoleRepository{Roles: []*domain.AccessManagerRole{
		{RoleID: operatorRoleID, Label: "PRIVACY_NODE_OPERATOR"},
	}}
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: []*domain.AccessManagerRoleMember{
		{RoleID: operatorRoleID, Account: "0xOPERATOR", IsActive: true},
	}}
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{
			Model:           domain.Model{ID: uuid.New()},
			RaylsAddress:    "0xOPERATOR",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
		},
	}}

	_, err := newResolver(roles, members, wallets).Resolve(context.Background())

	var noSigner *core.NoOperatorSignerError
	require.Error(t, err)
	assert.True(t, errors.As(err, &noSigner))
}

func TestOperatorSignerResolver_Resolve_FailsClosedWhenResolvedAccountHasNoWallet(t *testing.T) {
	// The dual-role account has no UserWallet row, so it is not signable → fail closed.
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: dualMember("0xNOWALLET")}
	wallets := &testutil.FakeUserWalletRepository{}

	_, err := newResolver(bothRoles(), members, wallets).Resolve(context.Background())

	var noSigner *core.NoOperatorSignerError
	require.Error(t, err)
	assert.True(t, errors.As(err, &noSigner))
}

func TestOperatorSignerResolver_Resolve_FailsClosedWhenResolvedAccountNotHSM(t *testing.T) {
	// The dual-role account's wallet is self-custodied; custody can only sign HSM wallets → fail closed.
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: dualMember("0xSELF")}
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{Model: domain.Model{ID: uuid.New()}, RaylsAddress: "0xSELF", CustodyProvider: domain.CustodyProviderSelf},
	}}

	_, err := newResolver(bothRoles(), members, wallets).Resolve(context.Background())

	var noSigner *core.NoOperatorSignerError
	require.Error(t, err)
	assert.True(t, errors.As(err, &noSigner))
}

// --- FakeAccessManagerRoleMemberRepository.FindActiveAccountWithAllRoles ---

func TestFakeMemberRepo_FindActiveAccountWithAllRoles_ReturnsDualRoleAccount(t *testing.T) {
	// An account active in every requested role is returned.
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: dualMember("0xBOTH")}

	account, err := members.FindActiveAccountWithAllRoles(context.Background(), []uint64{operatorRoleID, bankRoleID})

	require.NoError(t, err)
	assert.Equal(t, "0xBOTH", account)
}

func TestFakeMemberRepo_FindActiveAccountWithAllRoles_ExcludesSingleRoleAccount(t *testing.T) {
	// An account in only one of the requested roles does not satisfy the intersection.
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: []*domain.AccessManagerRoleMember{
		{RoleID: operatorRoleID, Account: "0xONLYOP", IsActive: true},
	}}

	_, err := members.FindActiveAccountWithAllRoles(context.Background(), []uint64{operatorRoleID, bankRoleID})

	assert.ErrorIs(t, err, core.ErrRecordNotFound)
}

func TestFakeMemberRepo_FindActiveAccountWithAllRoles_InactiveMembershipDisqualifies(t *testing.T) {
	// An inactive membership in one role means the account no longer holds both.
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: []*domain.AccessManagerRoleMember{
		{RoleID: operatorRoleID, Account: "0xX", IsActive: true},
		{RoleID: bankRoleID, Account: "0xX", IsActive: false},
	}}

	_, err := members.FindActiveAccountWithAllRoles(context.Background(), []uint64{operatorRoleID, bankRoleID})

	assert.ErrorIs(t, err, core.ErrRecordNotFound)
}

func TestFakeMemberRepo_FindActiveAccountWithAllRoles_EmptyRolesReturnsNotFound(t *testing.T) {
	// An empty role set matches nothing.
	members := &testutil.FakeAccessManagerRoleMemberRepository{Members: dualMember("0xBOTH")}

	_, err := members.FindActiveAccountWithAllRoles(context.Background(), nil)

	assert.ErrorIs(t, err, core.ErrRecordNotFound)
}
