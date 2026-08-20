package indexer

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

func TestBlockscoutBalancesListener_Handle_UpsertsAndPublishesForKnownWallet(t *testing.T) {
	// A balance_change notification for a wallet in user_wallets upserts the balance and emits SSE.
	walletAddr := domain.NormalizeAddress("0xAAAaaAAaaaAaaaAAaaAaAAAaaaAaaAAAAaAaAaaA")
	tokenHex := "BBBbbBBbbbBbbbBBbbBbBBBbbbBbbBBBBbBbBbbB"
	walletHex := "AAAaaAAaaaAaaaAAaaAaAAAaaaAaaAAAAaAaAaaA"

	wallets := &testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			Model: domain.Model{ID: uuid.New()}, UserID: uuid.New(),
			RaylsAddress: walletAddr, IsActive: true,
		}},
	}
	balances := &testutil.FakeWalletBalanceRepository{}
	lp := &fakeLivePublisher{}

	l := NewBlockscoutBalancesListener("", balances, wallets, nil, lp, &testutil.StubLogger{}, "inst")

	payload, _ := json.Marshal(balanceChangePayload{
		Op:                       "INSERT",
		AddressHash:              walletHex,
		TokenContractAddressHash: tokenHex,
		Value:                    "42000000",
		BlockNumber:              99,
	})

	l.handle(context.Background(), string(payload))

	require.Len(t, balances.Balances, 1)
	stored := balances.Balances[0]
	assert.Equal(t, walletAddr, stored.WalletAddress)
	assert.Equal(t, domain.NormalizeAddress(tokenHex), stored.TokenAddress)
	assert.Equal(t, "42000000", stored.Balance)
	assert.Equal(t, uint64(99), stored.BlockNumber)

	require.Equal(t, 1, lp.calls)
	assert.Equal(t, "ops.inst.sse.wallet_balances", lp.subject)

	var evt WalletBalanceEvent
	require.NoError(t, json.Unmarshal(lp.data, &evt))
	assert.Equal(t, "balance_updated", evt.Type)
	assert.Equal(t, walletAddr, evt.WalletAddress)
	assert.Equal(t, "42000000", evt.Balance)
	assert.Equal(t, uint64(99), evt.BlockNumber)
}

func TestBlockscoutBalancesListener_Handle_SkipsUnknownWallet(t *testing.T) {
	// A notification for a wallet not in user_wallets is dropped: no upsert, no publish.
	wallets := &testutil.FakeUserWalletRepository{}
	balances := &testutil.FakeWalletBalanceRepository{}
	lp := &fakeLivePublisher{}

	l := NewBlockscoutBalancesListener("", balances, wallets, nil, lp, &testutil.StubLogger{}, "")

	payload, _ := json.Marshal(balanceChangePayload{
		Op:                       "INSERT",
		AddressHash:              "1111111111111111111111111111111111111111",
		TokenContractAddressHash: "2222222222222222222222222222222222222222",
		Value:                    "1",
		BlockNumber:              1,
	})

	l.handle(context.Background(), string(payload))

	assert.Empty(t, balances.Balances)
	assert.Equal(t, 0, lp.calls)
}

func TestBlockscoutBalancesListener_Handle_MalformedPayloadDoesNotPanic(t *testing.T) {
	// An unparseable payload is logged and dropped — handler returns cleanly.
	wallets := &testutil.FakeUserWalletRepository{}
	balances := &testutil.FakeWalletBalanceRepository{}
	lp := &fakeLivePublisher{}

	l := NewBlockscoutBalancesListener("", balances, wallets, nil, lp, &testutil.StubLogger{}, "")

	l.handle(context.Background(), "{not json")

	assert.Empty(t, balances.Balances)
	assert.Equal(t, 0, lp.calls)
}

func TestBlockscoutBalancesListener_Handle_BlockNumberGuardKeepsFresher(t *testing.T) {
	// An older block_number must NOT overwrite a stored fresher balance.
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
		WalletAddress: walletAddr, TokenAddress: tokenAddr, Balance: "100", BlockNumber: 50,
	}))

	l := NewBlockscoutBalancesListener("", balances, wallets, nil, &fakeLivePublisher{}, &testutil.StubLogger{}, "")

	payload, _ := json.Marshal(balanceChangePayload{
		Op:                       "UPDATE",
		AddressHash:              walletAddr[2:],
		TokenContractAddressHash: tokenAddr[2:],
		Value:                    "1",
		BlockNumber:              10,
	})

	l.handle(context.Background(), string(payload))

	require.Len(t, balances.Balances, 1)
	assert.Equal(t, "100", balances.Balances[0].Balance)
	assert.Equal(t, uint64(50), balances.Balances[0].BlockNumber)
}
