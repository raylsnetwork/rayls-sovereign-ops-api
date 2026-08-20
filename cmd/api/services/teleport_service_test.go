package services

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

type fakeTeleportRegistry struct {
	exists bool
	err    error
}

func (f *fakeTeleportRegistry) Exists(_ context.Context, _ string) (bool, error) {
	return f.exists, f.err
}

type fakeTokenChainClient struct {
	erc20Balance   *big.Int
	erc721Owner    string
	erc1155Balance *big.Int

	calledMethod   string // "ERC20" / "ERC721" / "ERC1155"
	gotSigner      string
	gotTo          string
	gotAmount      *big.Int
	gotTokenID     *big.Int
	gotDestChainID *big.Int
	txHash         string
	teleportErr    error
}

func (f *fakeTokenChainClient) ERC20Balance(_ context.Context, _, _ string) (*big.Int, error) {
	return f.erc20Balance, nil
}

func (f *fakeTokenChainClient) ERC721Owner(_ context.Context, _ string, _ *big.Int) (string, error) {
	return f.erc721Owner, nil
}

func (f *fakeTokenChainClient) ERC1155Balance(_ context.Context, _, _ string, _ *big.Int) (*big.Int, error) {
	return f.erc1155Balance, nil
}

func (f *fakeTokenChainClient) TeleportERC20(
	_ context.Context,
	signer, _, to string,
	amount, destChainID *big.Int,
) (string, error) {
	f.calledMethod = "ERC20"
	f.gotSigner = signer
	f.gotTo = to
	f.gotAmount = amount
	f.gotDestChainID = destChainID
	return f.txHash, f.teleportErr
}

func (f *fakeTokenChainClient) TeleportERC721(
	_ context.Context,
	signer, _, to string,
	tokenID, destChainID *big.Int,
) (string, error) {
	f.calledMethod = "ERC721"
	f.gotSigner = signer
	f.gotTo = to
	f.gotTokenID = tokenID
	f.gotDestChainID = destChainID
	return f.txHash, f.teleportErr
}

func (f *fakeTokenChainClient) TeleportERC1155(
	_ context.Context,
	signer, _, to string,
	tokenID, amount, destChainID *big.Int,
	_ []byte,
) (string, error) {
	f.calledMethod = "ERC1155"
	f.gotSigner = signer
	f.gotTo = to
	f.gotTokenID = tokenID
	f.gotAmount = amount
	f.gotDestChainID = destChainID
	return f.txHash, f.teleportErr
}

func TestTeleportService_Teleport_UnsupportedStandardRejectsBeforeChain(t *testing.T) {
	// Standards outside ERC20/721/1155 are rejected with a validation error and never reach the chain.
	chain := &fakeTokenChainClient{}
	svc := NewTeleportService(chain, &fakeTeleportRegistry{exists: true}, big.NewInt(99), &testutil.StubLogger{})

	from := "0x1111111111111111111111111111111111111111"
	_, err := svc.Teleport(context.Background(), "0xtoken", domain.ErcStandardEnygma, core.TeleportInput{
		From: from, To: from, Amount: big.NewInt(1),
	})

	var verr *core.ValidationError
	require.ErrorAs(t, err, &verr)
	assert.Empty(t, chain.calledMethod)
}

func TestTeleportService_Teleport_UnregisteredTokenRejected(t *testing.T) {
	// A token absent from the registry is rejected before any balance read or signing.
	chain := &fakeTokenChainClient{}
	svc := NewTeleportService(chain, &fakeTeleportRegistry{exists: false}, big.NewInt(99), &testutil.StubLogger{})

	from := "0x1111111111111111111111111111111111111111"
	_, err := svc.Teleport(context.Background(), "0xtoken", domain.ErcStandardERC20, core.TeleportInput{
		From: from, To: from, Amount: big.NewInt(1),
	})

	var verr *core.ValidationError
	require.ErrorAs(t, err, &verr)
	assert.Empty(t, chain.calledMethod)
}

func TestTeleportService_Teleport_InsufficientERC20BalanceRejected(t *testing.T) {
	// An ERC20 balance below the requested amount is rejected before signing.
	chain := &fakeTokenChainClient{erc20Balance: big.NewInt(500)}
	svc := NewTeleportService(chain, &fakeTeleportRegistry{exists: true}, big.NewInt(99), &testutil.StubLogger{})

	from := "0x1111111111111111111111111111111111111111"
	_, err := svc.Teleport(context.Background(), "0xtoken", domain.ErcStandardERC20, core.TeleportInput{
		From: from, To: from, Amount: big.NewInt(1000),
	})

	var verr *core.ValidationError
	require.ErrorAs(t, err, &verr)
	assert.Empty(t, chain.calledMethod)
}

func TestTeleportService_Teleport_ERC721WrongOwnerRejected(t *testing.T) {
	// An ERC721 token owned by someone else is rejected before signing.
	chain := &fakeTokenChainClient{erc721Owner: "0x2222222222222222222222222222222222222222"}
	svc := NewTeleportService(chain, &fakeTeleportRegistry{exists: true}, big.NewInt(99), &testutil.StubLogger{})

	from := "0x1111111111111111111111111111111111111111"
	_, err := svc.Teleport(context.Background(), "0xtoken", domain.ErcStandardERC721, core.TeleportInput{
		From: from, To: from, TokenID: big.NewInt(7),
	})

	var verr *core.ValidationError
	require.ErrorAs(t, err, &verr)
	assert.Empty(t, chain.calledMethod)
}

func TestTeleportService_Teleport_HappyPathDelegatesToChain(t *testing.T) {
	// A registered token with sufficient balance reaches the chain client and returns its tx hash.
	chain := &fakeTokenChainClient{erc20Balance: big.NewInt(1000), txHash: "0xtx"}
	svc := NewTeleportService(chain, &fakeTeleportRegistry{exists: true}, big.NewInt(99), &testutil.StubLogger{})

	from := "0x1111111111111111111111111111111111111111"
	in := core.TeleportInput{From: from, To: from, Amount: big.NewInt(1000)}
	txHash, err := svc.Teleport(context.Background(), "0xtoken", domain.ErcStandardERC20, in)

	require.NoError(t, err)
	assert.Equal(t, "0xtx", txHash)
	assert.Equal(t, "ERC20", chain.calledMethod)
	assert.Equal(t, from, chain.gotSigner)
	assert.Equal(t, from, chain.gotTo)
	assert.Equal(t, big.NewInt(1000), chain.gotAmount)
	assert.Equal(t, big.NewInt(99), chain.gotDestChainID)
}

func TestTeleportService_Teleport_RegistryErrorPropagates(t *testing.T) {
	// A registry lookup failure is wrapped and returned before any balance read or signing.
	chain := &fakeTokenChainClient{}
	svc := NewTeleportService(
		chain,
		&fakeTeleportRegistry{err: errors.New("registry down")},
		big.NewInt(99),
		&testutil.StubLogger{},
	)

	from := "0x1111111111111111111111111111111111111111"
	_, err := svc.Teleport(context.Background(), "0xtoken", domain.ErcStandardERC20, core.TeleportInput{
		From: from, To: from, Amount: big.NewInt(1),
	})

	require.Error(t, err)
	assert.Empty(t, chain.calledMethod)
}

func TestTeleportService_Teleport_ERC20MissingAmountRejected(t *testing.T) {
	// A nil amount fails the ERC20 preflight with a validation error before signing.
	chain := &fakeTokenChainClient{erc20Balance: big.NewInt(1000)}
	svc := NewTeleportService(chain, &fakeTeleportRegistry{exists: true}, big.NewInt(99), &testutil.StubLogger{})

	from := "0x1111111111111111111111111111111111111111"
	_, err := svc.Teleport(
		context.Background(),
		"0xtoken",
		domain.ErcStandardERC20,
		core.TeleportInput{From: from, To: from},
	)

	var verr *core.ValidationError
	require.ErrorAs(t, err, &verr)
	assert.Empty(t, chain.calledMethod)
}

func TestTeleportService_Teleport_ChainClientErrorPropagates(t *testing.T) {
	// A signing/broadcast failure from the chain client (nonce/gas/custody) propagates to the caller.
	chain := &fakeTokenChainClient{erc20Balance: big.NewInt(1000), teleportErr: errors.New("sign and transact: boom")}
	svc := NewTeleportService(chain, &fakeTeleportRegistry{exists: true}, big.NewInt(99), &testutil.StubLogger{})

	from := "0x1111111111111111111111111111111111111111"
	_, err := svc.Teleport(context.Background(), "0xtoken", domain.ErcStandardERC20, core.TeleportInput{
		From: from, To: from, Amount: big.NewInt(1000),
	})

	require.Error(t, err)
	assert.Equal(t, "ERC20", chain.calledMethod)
}
