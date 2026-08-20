package blockchain

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts/RNContractFactoryV1"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

func newFactoryService() *RaylsContractFactoryService {
	return &RaylsContractFactoryService{factory: RNContractFactoryV1.NewRNContractFactoryV1()}
}

// resourceId is always bytes32(0) at deploy time.
var zeroResource [32]byte

func TestRaylsContractFactoryService_DeployCalldata_ERC20RoutesToDeployErc20(t *testing.T) {
	// An ERC20 spec produces the same calldata as the typed PackDeployErc20 binding call (resourceId = 0).
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardERC20, Name: "Token", Symbol: "TKN", Decimals: 18,
	})

	require.NoError(t, err)
	want := RNContractFactoryV1.NewRNContractFactoryV1().PackDeployErc20("Token", "TKN", 18, zeroResource)
	assert.Equal(t, want, got)
}

func TestRaylsContractFactoryService_DeployCalldata_EnygmaRoutesToDeployEnygma(t *testing.T) {
	// An Enygma spec produces the same calldata as the typed PackDeployEnygma binding call.
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardEnygma, Name: "Eny", Symbol: "ENY", Decimals: 6,
	})

	require.NoError(t, err)
	want := RNContractFactoryV1.NewRNContractFactoryV1().PackDeployEnygma("Eny", "ENY", 6, zeroResource)
	assert.Equal(t, want, got)
}

func TestRaylsContractFactoryService_DeployCalldata_ERC721RoutesToDeployErc721(t *testing.T) {
	// An ERC721 spec produces the same calldata as the typed PackDeployErc721 binding call.
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardERC721, URI: "ipfs://x", Name: "NFT", Symbol: "NFT",
	})

	require.NoError(t, err)
	want := RNContractFactoryV1.NewRNContractFactoryV1().PackDeployErc721("ipfs://x", "NFT", "NFT", zeroResource)
	assert.Equal(t, want, got)
}

func TestRaylsContractFactoryService_DeployCalldata_ERC1155RoutesToDeployErc1155(t *testing.T) {
	// An ERC1155 spec produces the same calldata as the typed PackDeployErc1155 binding call.
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardERC1155, URI: "ipfs://y", Name: "Multi",
	})

	require.NoError(t, err)
	want := RNContractFactoryV1.NewRNContractFactoryV1().PackDeployErc1155("ipfs://y", "Multi", zeroResource)
	assert.Equal(t, want, got)
}

func TestRaylsContractFactoryService_DeployCalldata_ZkDvpERC721RoutesToDeployErc721Dvp(t *testing.T) {
	// A ZkDvp ERC721 spec produces the same calldata as the typed PackDeployErc721Dvp binding call.
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardZkDvpERC721, URI: "ipfs://z", Name: "Dvp", Symbol: "DVP",
	})

	require.NoError(t, err)
	want := RNContractFactoryV1.NewRNContractFactoryV1().PackDeployErc721Dvp("ipfs://z", "Dvp", "DVP", zeroResource)
	assert.Equal(t, want, got)
}

func TestRaylsContractFactoryService_DeployCalldata_ZkDvpERC1155RoutesToDeployErc1155Dvp(t *testing.T) {
	// A ZkDvp ERC1155 spec produces the same calldata as the typed PackDeployErc1155Dvp binding call.
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardZkDvpERC1155, URI: "ipfs://w", Name: "DvpMulti",
	})

	require.NoError(t, err)
	want := RNContractFactoryV1.NewRNContractFactoryV1().PackDeployErc1155Dvp("ipfs://w", "DvpMulti", zeroResource)
	assert.Equal(t, want, got)
}

func TestRaylsContractFactoryService_DeployCalldata_StableCoinRoutesToDeployAsUser(t *testing.T) {
	// A StableCoin routes to the *AsUser* wrapper so the DEPLOYING WALLET becomes trusted.owner.
	// The plain deployRegistered path left owner = factoryOwner, and since the handler seeds
	// pauser/masterMinter/blacklister from trusted.owner, the deployer could not pause their
	// own token. Routing here is what makes the deployer the pauser.
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardStableCoin, Name: "USD Coin", Symbol: "USDC", Decimals: 6,
	})

	require.NoError(t, err)
	want := RNContractFactoryV1.NewRNContractFactoryV1().
		PackDeployStableCoinAsUser("USD Coin", "USDC", 6)
	assert.Equal(t, want, got)
}

func TestRaylsContractFactoryService_DeployCalldata_StableCoinUsesOnChainSelector(t *testing.T) {
	// The stablecoin calldata must carry the selector the deployed factory actually exposes.
	// Comparing calldata against the binding alone cannot catch a stale ABI: a wrong binding
	// matches itself. Pinning the keccak of the literal on-chain signature does catch it —
	// a 4-arg ABI (with a resourceId the *AsUser wrapper does not take) yields 0x7191dbe9,
	// which has no implementation on the factory and reverts inside eth_estimateGas.
	svc := newFactoryService()

	got, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardStableCoin, Name: "USD Coin", Symbol: "USDC", Decimals: 6,
	})

	require.NoError(t, err)
	want := crypto.Keccak256([]byte("deployStableCoinAsUser(string,string,uint8)"))[:4]
	assert.Equal(t, want, got[:4])
}

func TestRaylsContractFactoryService_DeployCalldata_UnsupportedStandardErrors(t *testing.T) {
	// The Custom standard has no factory deploy method and must error.
	svc := newFactoryService()

	_, err := svc.deployCalldata(core.TokenDeploySpec{
		ErcStandard: domain.ErcStandardCustom, Name: "X", Symbol: "X",
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported token standard")
}
