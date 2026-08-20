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

func TestWalletBalanceService_ListForWallet_UnknownWalletReturnsErrWalletNotFound(t *testing.T) {
	// Listing balances for an address that is not in user_wallets must surface ErrWalletNotFound.
	balances := &testutil.FakeWalletBalanceRepository{}
	wallets := &testutil.FakeUserWalletRepository{}
	tokens := &testutil.FakeTokenRepository{}

	svc := NewWalletBalanceService(balances, wallets, tokens, &testutil.StubLogger{})

	_, err := svc.ListForWallet(context.Background(), "0x1111111111111111111111111111111111111111")

	require.Error(t, err)
	assert.True(t, errors.Is(err, core.ErrWalletNotFound))
}

func TestWalletBalanceService_ListForWallet_KnownWalletReturnsEnrichedView(t *testing.T) {
	// A known wallet returns each stored balance enriched with token symbol, name, and decimals.
	walletAddr := domain.NormalizeAddress("0xAAAaaAAaaaAaaaAAaaAaAAAaaaAaaAAAAaAaAaaA")
	tokenAddr := domain.NormalizeAddress("0xBBBbbBBbbbBbbbBBbbBbBBBbbbBbbBBBBbBbBbbB")

	wallets := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			Model:        domain.Model{ID: uuid.New()},
			UserID:       uuid.New(),
			RaylsAddress: walletAddr,
			IsActive:     true,
		}},
	}
	balances := &testutil.FakeWalletBalanceRepository{}
	require.NoError(t, balances.Upsert(context.Background(), &domain.WalletBalance{
		WalletAddress: walletAddr,
		TokenAddress:  tokenAddr,
		Balance:       "12345",
		BlockNumber:   100,
	}))
	tokens := &testutil.FakeTokenRepository{}
	require.NoError(t, tokens.Upsert(context.Background(), &domain.Token{
		ContractAddress: tokenAddr,
		Symbol:          "TKN",
		Name:            "Test Token",
		Decimals:        18,
	}))

	svc := NewWalletBalanceService(balances, wallets, tokens, &testutil.StubLogger{})

	views, err := svc.ListForWallet(context.Background(), walletAddr)

	require.NoError(t, err)
	require.Len(t, views, 1)
	assert.Equal(t, walletAddr, views[0].WalletAddress)
	assert.Equal(t, tokenAddr, views[0].TokenAddress)
	assert.Equal(t, "12345", views[0].Balance)
	assert.Equal(t, uint64(100), views[0].BlockNumber)
	assert.Equal(t, "TKN", views[0].TokenSymbol)
	assert.Equal(t, "Test Token", views[0].TokenName)
	assert.Equal(t, uint8(18), views[0].Decimals)
}

func TestWalletBalanceService_GetForWalletAndToken_UnknownWalletReturnsErrWalletNotFound(t *testing.T) {
	// Fetching by (wallet, token) for a wallet that is not in user_wallets must surface ErrWalletNotFound.
	balances := &testutil.FakeWalletBalanceRepository{}
	wallets := &testutil.FakeUserWalletRepository{}
	tokens := &testutil.FakeTokenRepository{}

	svc := NewWalletBalanceService(balances, wallets, tokens, &testutil.StubLogger{})

	_, err := svc.GetForWalletAndToken(context.Background(),
		"0x1111111111111111111111111111111111111111",
		"0x2222222222222222222222222222222222222222",
	)

	require.Error(t, err)
	assert.True(t, errors.Is(err, core.ErrWalletNotFound))
}

func TestWalletBalanceService_GetForWalletAndToken_KnownPairReturnsEnrichedView(t *testing.T) {
	// A known (wallet, token) pair returns a single balance enriched with token metadata.
	walletAddr := domain.NormalizeAddress("0xAAAaaAAaaaAaaaAAaaAaAAAaaaAaaAAAAaAaAaaA")
	tokenAddr := domain.NormalizeAddress("0xBBBbbBBbbbBbbbBBbbBbBBBbbbBbbBBBBbBbBbbB")

	wallets := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			Model:        domain.Model{ID: uuid.New()},
			UserID:       uuid.New(),
			RaylsAddress: walletAddr,
			IsActive:     true,
		}},
	}
	balances := &testutil.FakeWalletBalanceRepository{}
	require.NoError(t, balances.Upsert(context.Background(), &domain.WalletBalance{
		WalletAddress: walletAddr,
		TokenAddress:  tokenAddr,
		Balance:       "777",
		BlockNumber:   200,
	}))
	tokens := &testutil.FakeTokenRepository{}
	require.NoError(t, tokens.Upsert(context.Background(), &domain.Token{
		ContractAddress: tokenAddr,
		Symbol:          "TKN",
		Name:            "Test Token",
		Decimals:        6,
	}))

	svc := NewWalletBalanceService(balances, wallets, tokens, &testutil.StubLogger{})

	view, err := svc.GetForWalletAndToken(context.Background(), walletAddr, tokenAddr)

	require.NoError(t, err)
	require.NotNil(t, view)
	assert.Equal(t, walletAddr, view.WalletAddress)
	assert.Equal(t, tokenAddr, view.TokenAddress)
	assert.Equal(t, "777", view.Balance)
	assert.Equal(t, uint64(200), view.BlockNumber)
	assert.Equal(t, "TKN", view.TokenSymbol)
	assert.Equal(t, "Test Token", view.TokenName)
	assert.Equal(t, uint8(6), view.Decimals)
}

func TestWalletBalanceService_GetForWalletAndToken_MissingBalanceReturnsNotFoundError(t *testing.T) {
	// A known wallet that holds no balance for the queried token returns a typed *NotFoundError
	// so HandleError can produce a 404 with a useful resource identifier.
	walletAddr := domain.NormalizeAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeE")
	tokenAddr := domain.NormalizeAddress("0xfffffffffffffffffffffffffffffffffffffffF")

	wallets := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			Model: domain.Model{ID: uuid.New()}, UserID: uuid.New(),
			RaylsAddress: walletAddr, IsActive: true,
		}},
	}
	balances := &testutil.FakeWalletBalanceRepository{}
	tokens := &testutil.FakeTokenRepository{}

	svc := NewWalletBalanceService(balances, wallets, tokens, &testutil.StubLogger{})

	_, err := svc.GetForWalletAndToken(context.Background(), walletAddr, tokenAddr)

	require.Error(t, err)
	var nfErr *core.NotFoundError
	require.True(t, errors.As(err, &nfErr))
	assert.Equal(t, "wallet balance", nfErr.Resource)
	assert.Contains(t, nfErr.ID, walletAddr)
	assert.Contains(t, nfErr.ID, tokenAddr)
}

func TestWalletBalanceService_ListForWallet_TokenMissingDoesNotFail(t *testing.T) {
	// A balance whose token row hasn't been indexed yet is returned without enrichment, not as an error.
	walletAddr := domain.NormalizeAddress("0xccccccccccccccccccccccccccccccccccccccCC")
	tokenAddr := domain.NormalizeAddress("0xdddddddddddddddddddddddddddddddddddddddD")

	wallets := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			Model: domain.Model{ID: uuid.New()}, UserID: uuid.New(),
			RaylsAddress: walletAddr, IsActive: true,
		}},
	}
	balances := &testutil.FakeWalletBalanceRepository{}
	require.NoError(t, balances.Upsert(context.Background(), &domain.WalletBalance{
		WalletAddress: walletAddr, TokenAddress: tokenAddr, Balance: "1", BlockNumber: 1,
	}))
	tokens := &testutil.FakeTokenRepository{}

	svc := NewWalletBalanceService(balances, wallets, tokens, &testutil.StubLogger{})

	views, err := svc.ListForWallet(context.Background(), walletAddr)

	require.NoError(t, err)
	require.Len(t, views, 1)
	assert.Equal(t, "", views[0].TokenSymbol)
	assert.Equal(t, "1", views[0].Balance)
}
