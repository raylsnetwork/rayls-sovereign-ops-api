package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

func identityUser() *domain.User {
	return &domain.User{
		Model:    domain.Model{ID: uuid.New()},
		Email:    "new@example.com",
		IsActive: true,
		Status:   domain.UserStatusWaitingRoleAssignment,
	}
}

func TestIdentityProvisioningService_Provision_MintsCustodyWallet(t *testing.T) {
	// A new account gets its ONE custody wallet here — identity owns wallets, and without
	// this every chain reports "no custody wallet" for a user who just signed up.
	user := identityUser()
	wallets := &testutil.FakeUserWalletRepository{}
	users := &testutil.FakeUserRepository{Users: []domain.User{*user}}
	svc := NewIdentityProvisioningService(
		users, wallets,
		testutil.NewFakeCustodyService("0xMINTED", "ext-1"),
		domain.CustodyProviderRaylsHSM, &testutil.StubLogger{},
	)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Len(t, wallets.Wallets, 1)
	assert.Equal(t, "0xMINTED", wallets.Wallets[0].RaylsAddress)
	assert.Equal(t, domain.WalletChainPrivate, wallets.Wallets[0].Chain)
	assert.Equal(t, domain.UserStatusRoleAssigned, user.Status)
}

func TestIdentityProvisioningService_Provision_IsIdempotent(t *testing.T) {
	// Runs on every login, so an existing wallet must be reused rather than re-minted.
	user := identityUser()
	user.Status = domain.UserStatusRoleAssigned
	wallets := &testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{{
		UserID:          user.ID,
		RaylsAddress:    "0xexisting",
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPrivate,
		IsActive:        true,
	}}}
	svc := NewIdentityProvisioningService(
		testutil.NewFakeUserRepository(), wallets,
		testutil.NewFakeCustodyService("0xSHOULD-NOT-BE-USED", "ext"),
		domain.CustodyProviderRaylsHSM, &testutil.StubLogger{},
	)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Len(t, wallets.Wallets, 1, "a second wallet was minted")
	assert.Equal(t, "0xexisting", wallets.Wallets[0].RaylsAddress)
}

func TestIdentityProvisioningService_Provision_BackfillsWalletForActivatedAccount(t *testing.T) {
	// Accounts activated before identity owned wallets have status=role_assigned and NO
	// wallet. The wallet check runs before the status short-circuit so they are healed on
	// their next login instead of being stuck wallet-less forever.
	user := identityUser()
	user.Status = domain.UserStatusRoleAssigned
	wallets := &testutil.FakeUserWalletRepository{}
	svc := NewIdentityProvisioningService(
		testutil.NewFakeUserRepository(), wallets,
		testutil.NewFakeCustodyService("0xBACKFILLED", "ext-2"),
		domain.CustodyProviderRaylsHSM, &testutil.StubLogger{},
	)

	err := svc.Provision(context.Background(), user)

	require.NoError(t, err)
	require.Len(t, wallets.Wallets, 1, "an already-activated account was left without a wallet")
	assert.Equal(t, "0xBACKFILLED", wallets.Wallets[0].RaylsAddress)
}
