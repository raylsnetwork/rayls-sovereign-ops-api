package testutil

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

func TestFakeUserWalletRepository_GetSignerWalletForChain_ReturnsEarliestHSMPrivate(t *testing.T) {
	// The signer lookup returns the earliest active HSM wallet on the private chain, ignoring the
	// public-chain HSM wallet and any non-HSM wallet.
	userID := uuid.New()
	repo := &FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{
			UserID:          userID,
			RaylsAddress:    "0xself",
			CustodyProvider: domain.CustodyProviderSelf,
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		},
		{
			UserID:          userID,
			RaylsAddress:    "0xpriv1",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		},
		{
			UserID:          userID,
			RaylsAddress:    "0xpub",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
			Chain:           domain.WalletChainPublic,
			IsActive:        true,
		},
		{
			UserID:          userID,
			RaylsAddress:    "0xpriv2",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		},
	}}

	wallet, err := repo.GetSignerWalletForChain(context.Background(), userID, domain.WalletChainPrivate)

	require.NoError(t, err)
	assert.Equal(t, "0xpriv1", wallet.RaylsAddress)
}

func TestFakeUserWalletRepository_GetSignerWalletForChain_ReturnsPublicWallet(t *testing.T) {
	// Requesting the public chain returns the HSM public-chain wallet.
	userID := uuid.New()
	repo := &FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{
			UserID:          userID,
			RaylsAddress:    "0xpriv",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		},
		{
			UserID:          userID,
			RaylsAddress:    "0xpub",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
			Chain:           domain.WalletChainPublic,
			IsActive:        true,
		},
	}}

	wallet, err := repo.GetSignerWalletForChain(context.Background(), userID, domain.WalletChainPublic)

	require.NoError(t, err)
	assert.Equal(t, "0xpub", wallet.RaylsAddress)
}

func TestFakeUserWalletRepository_GetSignerWalletForChain_NotFoundWhenOnlyNonHSM(t *testing.T) {
	// A user with only a self-custody wallet has no signer wallet (custody can't sign it).
	userID := uuid.New()
	repo := &FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{
			UserID:          userID,
			RaylsAddress:    "0xself",
			CustodyProvider: domain.CustodyProviderSelf,
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		},
	}}

	_, err := repo.GetSignerWalletForChain(context.Background(), userID, domain.WalletChainPrivate)

	assert.ErrorIs(t, err, core.ErrRecordNotFound)
}

func TestFakeUserWalletRepository_GetSignerWalletForChain_UnsetChainResolvesPrivate(t *testing.T) {
	// A fixture that leaves Chain unset (0) is treated as private, mirroring the DB column default.
	userID := uuid.New()
	repo := &FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{UserID: userID, RaylsAddress: "0xhsm", CustodyProvider: domain.CustodyProviderRaylsHSM, IsActive: true},
	}}

	wallet, err := repo.GetSignerWalletForChain(context.Background(), userID, domain.WalletChainPrivate)

	require.NoError(t, err)
	assert.Equal(t, "0xhsm", wallet.RaylsAddress)
}

func TestFakeUserWalletRepository_FindByUserID_ReturnsEarliestActiveProviderAgnostic(t *testing.T) {
	// The identity lookup returns the earliest active wallet regardless of custody provider.
	userID := uuid.New()
	repo := &FakeUserWalletRepository{Wallets: []domain.UserWallet{
		{
			UserID:          userID,
			RaylsAddress:    "0xlogin",
			CustodyProvider: domain.CustodyProviderSelf,
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		},
		{
			UserID:          userID,
			RaylsAddress:    "0xhsm",
			CustodyProvider: domain.CustodyProviderRaylsHSM,
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		},
	}}

	wallet, err := repo.FindByUserID(context.Background(), userID)

	require.NoError(t, err)
	assert.Equal(t, "0xlogin", wallet.RaylsAddress)
}
