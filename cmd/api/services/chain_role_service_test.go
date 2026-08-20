package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

const chainRoleTestOperatorRoleID = uint64(7)

func operatorRoleRepo() *testutil.FakeAccessManagerRoleRepository {
	return &testutil.FakeAccessManagerRoleRepository{Roles: []*domain.AccessManagerRole{
		{RoleID: chainRoleTestOperatorRoleID, Name: domain.RolePrivacyNodeOperator, Label: "Operator"},
	}}
}

func TestChainRoleService_RolesForUser_ReturnsRolesHeldOnThisChain(t *testing.T) {
	// The user's wallet on this chain holds the operator role, so they hold it here.
	userID := uuid.New()
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{{
		UserID:          userID,
		RaylsAddress:    "0xaaa",
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPrivate,
		IsActive:        true,
	}}}
	members := &testutil.FakeAccessManagerRoleMemberRepository{
		Members: []*domain.AccessManagerRoleMember{
			{RoleID: chainRoleTestOperatorRoleID, Account: "0xaaa", IsActive: true},
		},
	}

	svc := NewChainRoleService(
		wallets,
		testutil.NewFakeUserRepository(),
		members,
		operatorRoleRepo(),
		&testutil.StubLogger{},
	)

	roles, err := svc.RolesForUser(context.Background(), userID)

	require.NoError(t, err)
	assert.Equal(t, []string{domain.RolePrivacyNodeOperator}, roles)
}

func TestChainRoleService_RolesForUser_IgnoresInactiveMembership(t *testing.T) {
	// A revoked grant takes effect as soon as the indexer sees it — no waiting for re-login.
	userID := uuid.New()
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{{
		UserID:          userID,
		RaylsAddress:    "0xaaa",
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPrivate,
		IsActive:        true,
	}}}
	members := &testutil.FakeAccessManagerRoleMemberRepository{
		Members: []*domain.AccessManagerRoleMember{
			{RoleID: chainRoleTestOperatorRoleID, Account: "0xaaa", IsActive: false},
		},
	}

	svc := NewChainRoleService(
		wallets,
		testutil.NewFakeUserRepository(),
		members,
		operatorRoleRepo(),
		&testutil.StubLogger{},
	)

	roles, err := svc.RolesForUser(context.Background(), userID)

	require.NoError(t, err)
	assert.Empty(t, roles)
}

func TestChainRoleService_RolesForUser_NoWalletHoldsNothing(t *testing.T) {
	// A KNOWN user whose identity account has no custody wallet yet holds nothing here —
	// not an error, and this service must never mint one to compensate.
	userID := uuid.New()
	wallets := &testutil.FakeUserWalletRepository{}
	users := &testutil.FakeUserRepository{Users: []domain.User{{Model: domain.Model{ID: userID}, IsActive: true}}}
	svc := NewChainRoleService(
		wallets, users,
		&testutil.FakeAccessManagerRoleMemberRepository{},
		operatorRoleRepo(), &testutil.StubLogger{},
	)

	roles, err := svc.RolesForUser(context.Background(), userID)

	require.NoError(t, err)
	assert.Empty(t, roles)
	assert.Empty(t, wallets.Wallets, "a wallet was minted — identity owns wallet creation")
}

func TestChainRoleService_RolesForUser_UnknownUserIsUnauthorized(t *testing.T) {
	// A token naming a user identity has never heard of is a stale session (the identity
	// database was recreated). It must force re-login, not present a wallet-less account.
	svc := NewChainRoleService(
		&testutil.FakeUserWalletRepository{},
		testutil.NewFakeUserRepository(),
		&testutil.FakeAccessManagerRoleMemberRepository{},
		operatorRoleRepo(), &testutil.StubLogger{},
	)

	roles, err := svc.RolesForUser(context.Background(), uuid.New())

	require.Error(t, err)
	var unauthorized *core.UnauthorizedError
	assert.ErrorAs(t, err, &unauthorized)
	assert.Nil(t, roles)
}

func TestChainRoleService_RolesForUser_UsesTheIdentityWallet(t *testing.T) {
	// One custody wallet per user, shared across chains: roles are resolved for that
	// address from THIS chain's am_* mirror.
	userID := uuid.New()
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{{
		UserID:          userID,
		RaylsAddress:    "0xidentity",
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPrivate,
		IsActive:        true,
	}}}
	members := &testutil.FakeAccessManagerRoleMemberRepository{
		Members: []*domain.AccessManagerRoleMember{
			{RoleID: chainRoleTestOperatorRoleID, Account: "0xidentity", IsActive: true},
		},
	}

	svc := NewChainRoleService(
		wallets,
		testutil.NewFakeUserRepository(),
		members,
		operatorRoleRepo(),
		&testutil.StubLogger{},
	)

	roles, err := svc.RolesForUser(context.Background(), userID)

	require.NoError(t, err)
	assert.Equal(t, []string{domain.RolePrivacyNodeOperator}, roles)
	assert.Len(t, wallets.Wallets, 1, "no wallet should be created here")
}
