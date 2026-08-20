package blockchain

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

func sel(sig string) []byte { return crypto.Keccak256([]byte(sig))[:4] }

func TestMintCalldata_ERC20(t *testing.T) {
	// ERC20 mint dispatches mint(address,uint256) with the recipient and amount.
	cd, err := mintCalldata(domain.ErcStandardERC20, core.MintInput{
		To: "0x0000000000000000000000000000000000000001", Amount: big.NewInt(1000),
	})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("mint(address,uint256)")))
}

func TestMintCalldata_StableCoin_UsesERC20Signature(t *testing.T) {
	// The stablecoin inherits the ERC20 handler, so mint dispatches mint(address,uint256) —
	// byte-identical to the ERC20 path for the same args.
	to, amt := "0x0000000000000000000000000000000000000001", big.NewInt(1000)
	stable, err := mintCalldata(domain.ErcStandardStableCoin, core.MintInput{To: to, Amount: amt})
	require.NoError(t, err)
	erc20, err := mintCalldata(domain.ErcStandardERC20, core.MintInput{To: to, Amount: amt})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(stable, sel("mint(address,uint256)")))
	assert.Equal(t, erc20, stable)
}

func TestBurnCalldata_StableCoin_UsesERC20Signature(t *testing.T) {
	// The stablecoin burn dispatches burn(address,uint256), matching the ERC20 path.
	from, amt := "0x0000000000000000000000000000000000000002", big.NewInt(500)
	stable, err := burnCalldata(domain.ErcStandardStableCoin, core.BurnInput{From: from, Amount: amt})
	require.NoError(t, err)
	erc20, err := burnCalldata(domain.ErcStandardERC20, core.BurnInput{From: from, Amount: amt})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(stable, sel("burn(address,uint256)")))
	assert.Equal(t, erc20, stable)
}

func TestMintCalldata_ERC1155(t *testing.T) {
	// ERC1155 mint dispatches mint(address,uint256,uint256,bytes).
	cd, err := mintCalldata(domain.ErcStandardERC1155, core.MintInput{
		To: "0x0000000000000000000000000000000000000001", TokenID: big.NewInt(7), Amount: big.NewInt(5), Data: []byte{},
	})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("mint(address,uint256,uint256,bytes)")))
}

func TestMintCalldata_DvpERC721_EmptyExtraDataEncodes(t *testing.T) {
	// DVP ERC721 mint dispatches mint(address,uint256,(string,string,bool)[]) with an empty array.
	cd, err := mintCalldata(domain.ErcStandardZkDvpERC721, core.MintInput{
		To: "0x0000000000000000000000000000000000000001", TokenID: big.NewInt(9),
	})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("mint(address,uint256,(string,string,bool)[])")))
}

func TestMintCalldata_DvpERC1155_EmptyExtraDataEncodes(t *testing.T) {
	// DVP ERC1155 mint dispatches mint(address,uint256,uint256,bytes,(string,string,bool)[]).
	cd, err := mintCalldata(domain.ErcStandardZkDvpERC1155, core.MintInput{
		To: "0x0000000000000000000000000000000000000001", TokenID: big.NewInt(9), Amount: big.NewInt(3), Data: []byte{},
	})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("mint(address,uint256,uint256,bytes,(string,string,bool)[])")))
}

func TestBurnCalldata_ERC721(t *testing.T) {
	// ERC721 burn dispatches burn(uint256) with just the token id.
	cd, err := burnCalldata(domain.ErcStandardERC721, core.BurnInput{TokenID: big.NewInt(42)})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("burn(uint256)")))
}

func TestBurnCalldata_ERC1155(t *testing.T) {
	// ERC1155 burn dispatches burn(address,uint256,uint256).
	cd, err := burnCalldata(domain.ErcStandardERC1155, core.BurnInput{
		From: "0x0000000000000000000000000000000000000001", TokenID: big.NewInt(7), Amount: big.NewInt(5),
	})
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("burn(address,uint256,uint256)")))
}

func TestMintCalldata_UnsupportedStandardErrors(t *testing.T) {
	// The Custom standard has no mint mapping.
	_, err := mintCalldata(domain.ErcStandardCustom, core.MintInput{To: "0x0000000000000000000000000000000000000001"})
	require.Error(t, err)
}

func TestTeleportERC20Calldata(t *testing.T) {
	// ERC20 teleport dispatches teleportToPublicChain(address,uint256,uint256) with recipient, amount, destChainID.
	cd, err := teleportERC20Calldata("0x0000000000000000000000000000000000000002", big.NewInt(1000), big.NewInt(99))
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("teleportToPublicChain(address,uint256,uint256)")))
}

func TestTeleportERC721Calldata(t *testing.T) {
	// ERC721 teleport dispatches teleportToPublicChain(address,uint256,uint256) with recipient, id, destChainID.
	cd, err := teleportERC721Calldata("0x0000000000000000000000000000000000000002", big.NewInt(7), big.NewInt(99))
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("teleportToPublicChain(address,uint256,uint256)")))
}

func TestTeleportERC1155Calldata(t *testing.T) {
	// ERC1155 teleport dispatches teleportToPublicChain(address,uint256,uint256,uint256,bytes).
	cd, err := teleportERC1155Calldata(
		"0x0000000000000000000000000000000000000002",
		big.NewInt(7),
		big.NewInt(5),
		big.NewInt(99),
		[]byte{},
	)
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(cd, sel("teleportToPublicChain(address,uint256,uint256,uint256,bytes)")))
}

func TestTeleportERC20Calldata_MissingChainIDErrors(t *testing.T) {
	// A nil destination chain ID (PUBLIC_CHAIN_CHAIN_ID unset) is rejected before packing.
	_, err := teleportERC20Calldata("0x0000000000000000000000000000000000000002", big.NewInt(1), nil)
	require.Error(t, err)
}

func TestMintCalldata_ERC20_MissingAmountErrors(t *testing.T) {
	// ERC20 mint requires an amount; a nil amount is rejected before packing.
	_, err := mintCalldata(domain.ErcStandardERC20, core.MintInput{To: "0x0000000000000000000000000000000000000001"})
	require.Error(t, err)
}

func TestMintCalldata_ERC721_MissingTokenIDErrors(t *testing.T) {
	// ERC721 mint requires a tokenId; a nil tokenId is rejected before packing.
	_, err := mintCalldata(domain.ErcStandardERC721, core.MintInput{To: "0x0000000000000000000000000000000000000001"})
	require.Error(t, err)
}

func TestMintCalldata_ERC1155_MissingAmountErrors(t *testing.T) {
	// ERC1155 mint requires both id and amount; a nil amount is rejected before packing.
	_, err := mintCalldata(domain.ErcStandardERC1155, core.MintInput{
		To: "0x0000000000000000000000000000000000000001", TokenID: big.NewInt(1),
	})
	require.Error(t, err)
}

func TestMintCalldata_InvalidRecipientErrors(t *testing.T) {
	// A non-hex recipient address is rejected before packing.
	_, err := mintCalldata(domain.ErcStandardERC20, core.MintInput{To: "not-an-address", Amount: big.NewInt(1)})
	require.Error(t, err)
}

func TestBurnCalldata_ERC20_MissingAmountErrors(t *testing.T) {
	// ERC20 burn requires an amount; a nil amount is rejected before packing.
	_, err := burnCalldata(domain.ErcStandardERC20, core.BurnInput{From: "0x0000000000000000000000000000000000000001"})
	require.Error(t, err)
}

func TestBurnCalldata_ERC1155_MissingAmountErrors(t *testing.T) {
	// ERC1155 burn requires both id and amount; a nil amount is rejected before packing.
	_, err := burnCalldata(domain.ErcStandardERC1155, core.BurnInput{
		From: "0x0000000000000000000000000000000000000001", TokenID: big.NewInt(1),
	})
	require.Error(t, err)
}

func TestBurnCalldata_UnsupportedStandardErrors(t *testing.T) {
	// The Custom standard has no burn mapping.
	_, err := burnCalldata(domain.ErcStandardCustom, core.BurnInput{TokenID: big.NewInt(1)})
	require.Error(t, err)
}

func TestTeleportERC1155Calldata_MissingChainIDErrors(t *testing.T) {
	// A nil destination chain ID is rejected before packing for ERC1155 teleport.
	_, err := teleportERC1155Calldata(
		"0x0000000000000000000000000000000000000002",
		big.NewInt(1),
		big.NewInt(1),
		nil,
		[]byte{},
	)
	require.Error(t, err)
}
