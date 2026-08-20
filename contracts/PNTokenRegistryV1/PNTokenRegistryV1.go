// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PNTokenRegistryV1

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// TokenStructsFrozenToken is an auto generated low-level Go binding around an user-defined struct.
type TokenStructsFrozenToken struct {
	ResourceId         [32]byte
	FrozenParticipants []*big.Int
}

// TokenStructsToken is an auto generated low-level Go binding around an user-defined struct.
type TokenStructsToken struct {
	ResourceId         [32]byte
	Name               string
	Symbol             string
	Uri                string
	TokenAddress       common.Address
	PublicTokenAddress common.Address
	IssuerChainId      *big.Int
	ErcStandard        uint8
	PrivacyNodeStatus  uint8
	HubStatus          uint8
	PublicChainStatus  uint8
	CreatedAt          *big.Int
	UpdatedAt          *big.Int
}

// PNTokenRegistryV1MetaData contains all meta data concerning the PNTokenRegistryV1 contract.
var PNTokenRegistryV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activateToken\",\"inputs\":[{\"name\":\"tokenResourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ercStandard\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"endpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRaylsEndpoint\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freezeOnPrivacyNode\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOnPublicChain\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressByResourceId\",\"inputs\":[{\"name\":\"_resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTokenStructs.Token[]\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"issuerChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ercStandard\",\"type\":\"uint8\",\"internalType\":\"enumSharedObjects.ErcStandard\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"},{\"name\":\"hubStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"},{\"name\":\"publicChainStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PublicChainStatus\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFrozenTokenForParticipant\",\"inputs\":[{\"name\":\"tokenResourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHubStatus\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrivacyNodeStatus\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublicChainStatus\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PublicChainStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenByAddress\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTokenStructs.Token\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"issuerChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ercStandard\",\"type\":\"uint8\",\"internalType\":\"enumSharedObjects.ErcStandard\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"},{\"name\":\"hubStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"},{\"name\":\"publicChainStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PublicChainStatus\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenByResourceId\",\"inputs\":[{\"name\":\"tokenResourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTokenStructs.Token\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"issuerChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ercStandard\",\"type\":\"uint8\",\"internalType\":\"enumSharedObjects.ErcStandard\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"},{\"name\":\"hubStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"},{\"name\":\"publicChainStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PublicChainStatus\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBySymbol\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTokenStructs.Token\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"issuerChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ercStandard\",\"type\":\"uint8\",\"internalType\":\"enumSharedObjects.ErcStandard\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"},{\"name\":\"hubStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"},{\"name\":\"publicChainStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PublicChainStatus\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenCore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenCore\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenFreezeManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenFreezeManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokensByHubStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTokenStructs.Token[]\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"issuerChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ercStandard\",\"type\":\"uint8\",\"internalType\":\"enumSharedObjects.ErcStandard\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"},{\"name\":\"hubStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"},{\"name\":\"publicChainStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PublicChainStatus\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokensByPrivacyNodeStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTokenStructs.Token[]\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"issuerChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ercStandard\",\"type\":\"uint8\",\"internalType\":\"enumSharedObjects.ErcStandard\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"},{\"name\":\"hubStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.HubStatus\"},{\"name\":\"publicChainStatus\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PublicChainStatus\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"endpointAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"authority_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_endpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTokenActiveForHub\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTokenActiveForPublicChain\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTokenFullyOperational\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerHubToken\",\"inputs\":[{\"name\":\"tokenResourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFrozenToken\",\"inputs\":[{\"name\":\"unfrozenToken\",\"type\":\"tuple\",\"internalType\":\"structTokenStructs.FrozenToken\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"frozenParticipants\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestAllFrozenTokensDataFromPrivateHub\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resourceId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setResourceId\",\"inputs\":[{\"name\":\"_resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenCore\",\"inputs\":[{\"name\":\"tokenCoreAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenFreezeManager\",\"inputs\":[{\"name\":\"freezeManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitToHub\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitToPublicChain\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"syncFrozenTokens\",\"inputs\":[{\"name\":\"frozenTokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenStructs.FrozenToken[]\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"frozenParticipants\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenExists\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unfreezeOnPrivacyNode\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unfreezeOnPublicChain\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFrozenToken\",\"inputs\":[{\"name\":\"frozenToken\",\"type\":\"tuple\",\"internalType\":\"structTokenStructs.FrozenToken\",\"components\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"frozenParticipants\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePrivacyNodeStatus\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumTokenStructs.PrivacyNodeStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePublicTokenAddress\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validateTokenForParticipant\",\"inputs\":[{\"name\":\"tokenResourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenCoreSet\",\"inputs\":[{\"name\":\"tokenCore\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenFreezeManagerSet\",\"inputs\":[{\"name\":\"tokenFreezeManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRegistryModulesConfigured\",\"inputs\":[{\"name\":\"tokenCore\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenFreezeManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__MustSchedule\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAppV1__HubNotActive\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"hubStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"RaylsAppV1__PrivacyNodeFrozen\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAppV1__PrivacyNodeNotActive\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"RaylsAppV1__PublicChainNotActive\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privacyNodeStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"publicChainStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"RaylsAppV1__ResourceNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAppV1__TokenRegistryNotConfigured\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAppV1__UnauthorizedTokenRegistry\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRegistryV1__InvalidTokenCoreAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenRegistryV1__InvalidTokenFreezeManagerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenRegistryV1__TokenCoreNotConfigured\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenRegistryV1__TokenFreezeManagerNotConfigured\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:  "PNTokenRegistryV1",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051612b636200010460003960008181611a7601528181611a9f0152611bd60152612b636000f3fe60806040526004361061025d5760003560e01c80636f0656f011610145578063ad3cb1cc116100bc578063ad3cb1cc146106ec578063b33f78ca1461072a578063bf7e214f1461074a578063c3034ec01461075f578063c4d66de81461077f578063d8dc05101461079f578063de0c3625146107bf578063e3ade0ec146107d4578063efa74f1f146107f4578063f05ab9de14610814578063f0a599bb14610834578063f356cc0314610852578063f5aeb4341461087057600080fd5b80636f0656f01461058157806378a89567146105a157806378dd4e6f146105b65780637ae93604146105d65780637fbc532d146105f65780638a58b79f1461061657806391ded8fa1461064357806393d85727146106635780639eb7b4eb14610678578063a01afbfb14610698578063a0a8e460146106b8578063a0b9b9b1146106cc57600080fd5b80632f73e6e4116101d95780632f73e6e4146104155780633b93734914610435578063457415671461045557806347ecc93914610475578063485cc955146104955780634da8e6c3146104b55780634f1ef286146104d557806352d1902d146104e85780635e280f111461050b5780635f997c5b1461052b5780635fa7b5841461054157806364c0673a1461056157600080fd5b806214bc461461026257806301162d4314610297578063060ed143146102c457806309824a80146102e65780630a75cf4e1461030657806310dcd5201461032657806311f50c851461034657806313aa1a1d1461037357806314ae646b14610393578063222e1b16146103b357806322e594ff146103d35780632a5c792a146103f3575b600080fd5b34801561026e57600080fd5b5061028261027d366004612057565b610890565b60405190151581526020015b60405180910390f35b3480156102a357600080fd5b506102b76102b2366004612057565b61090b565b60405161028e91906120a7565b3480156102d057600080fd5b506102e46102df3660046120ba565b610981565b005b3480156102f257600080fd5b506102e4610301366004612057565b610a70565b34801561031257600080fd5b506102e4610321366004612057565b610aee565b34801561033257600080fd5b506102e4610341366004612057565b610b37565b34801561035257600080fd5b50610366610361366004612102565b610b80565b60405161028e919061211b565b34801561037f57600080fd5b506102e461038e366004612057565b610bee565b34801561039f57600080fd5b506102e46103ae366004612057565b610cb4565b3480156103bf57600080fd5b506102e46103ce36600461212f565b610d7a565b3480156103df57600080fd5b506102b76103ee366004612057565b610dfb565b3480156103ff57600080fd5b50610408610e30565b60405161028e919061230a565b34801561042157600080fd5b506102e4610430366004612057565b610ea4565b34801561044157600080fd5b5061040861045036600461237b565b610eed565b34801561046157600080fd5b50610282610470366004612057565b610f67565b34801561048157600080fd5b506102e4610490366004612398565b610f9c565b3480156104a157600080fd5b506102e46104b03660046123c8565b610ff1565b3480156104c157600080fd5b506102e46104d0366004612057565b611102565b6102e46104e33660046124ca565b61114b565b3480156104f457600080fd5b506104fd61116a565b60405190815260200161028e565b34801561051757600080fd5b50600054610366906001600160a01b031681565b34801561053757600080fd5b506104fd60015481565b34801561054d57600080fd5b506102e461055c366004612057565b611188565b34801561056d57600080fd5b5061028261057c366004612057565b6111d1565b34801561058d57600080fd5b5061040861059c36600461237b565b611206565b3480156105ad57600080fd5b506104fd61123b565b3480156105c257600080fd5b506102826105d136600461252d565b6112a6565b3480156105e257600080fd5b506102e46105f1366004612057565b61133f565b34801561060257600080fd5b506102e4610611366004612057565b611388565b34801561062257600080fd5b50610636610631366004612102565b6113d1565b60405161028e919061254f565b34801561064f57600080fd5b5061063661065e366004612057565b611453565b34801561066f57600080fd5b506104fd61148e565b34801561068457600080fd5b506102e4610693366004612562565b6114d5565b3480156106a457600080fd5b506102e46106b3366004612102565b61151e565b3480156106c457600080fd5b5060016104fd565b3480156106d857600080fd5b506102e46106e73660046123c8565b611565565b3480156106f857600080fd5b5061071d604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161028e919061259c565b34801561073657600080fd5b50610282610745366004612057565b6115b0565b34801561075657600080fd5b506103666115e5565b34801561076b57600080fd5b506102e461077a36600461252d565b6115fe565b34801561078b57600080fd5b506102e461079a366004612057565b611679565b3480156107ab57600080fd5b506102e46107ba3660046125af565b6116a3565b3480156107cb57600080fd5b506102e46116ee565b3480156107e057600080fd5b506102b76107ef366004612057565b611760565b34801561080057600080fd5b5061063661080f3660046125dd565b611795565b34801561082057600080fd5b506102e461082f366004612057565b6117d0565b34801561084057600080fd5b506002546001600160a01b0316610366565b34801561085e57600080fd5b506003546001600160a01b0316610366565b34801561087c57600080fd5b506102e461088b366004612562565b611819565b600061089a611862565b6001600160a01b03166214bc46836040518263ffffffff1660e01b81526004016108c4919061211b565b602060405180830381865afa1580156108e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109059190612642565b92915050565b6000610915611862565b6001600160a01b03166301162d43836040518263ffffffff1660e01b8152600401610940919061211b565b602060405180830381865afa15801561095d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109059190612668565b610997336000356001600160e01b03191661188b565b60006109a1611862565b60405163060ed14360e01b8152600481018690526001600160a01b03858116602483015260ff851660448301529192509082169063060ed14390606401600060405180830381600087803b1580156109f857600080fd5b505af1158015610a0c573d6000803e3d6000fd5b505060405163a01afbfb60e01b8152600481018790526001600160a01b038616925063a01afbfb9150602401600060405180830381600087803b158015610a5257600080fd5b505af1158015610a66573d6000803e3d6000fd5b5050505050505050565b610a86336000356001600160e01b03191661188b565b610a8e611862565b6001600160a01b03166309824a80826040518263ffffffff1660e01b8152600401610ab9919061211b565b600060405180830381600087803b158015610ad357600080fd5b505af1158015610ae7573d6000803e3d6000fd5b5050505050565b610b04336000356001600160e01b03191661188b565b610b0c611862565b6001600160a01b0316630a75cf4e826040518263ffffffff1660e01b8152600401610ab9919061211b565b610b4d336000356001600160e01b03191661188b565b610b556119cd565b6001600160a01b03166310dcd520826040518263ffffffff1660e01b8152600401610ab9919061211b565b600080546040516311f50c8560e01b8152600481018490526001600160a01b03909116906311f50c8590602401602060405180830381865afa158015610bca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109059190612690565b610c04336000356001600160e01b03191661188b565b6001600160a01b038116610c2b576040516365dce40760e11b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0383169081179091556040517f2dd3698b1731a899f810b168007b9bd902aec761bfe35f98fb17348bb174a92190600090a26003546002546040516001600160a01b0392831692909116907f70e091790811d89fdfce2f92748c357fdb8847d6e97f65b45fecf8d412d32deb90600090a350565b610cca336000356001600160e01b03191661188b565b6001600160a01b038116610cf15760405163743840cd60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383169081179091556040517fcd3287a073a6c154081e3c3168a5f5966880588e313e9dd388e935d177bec14b90600090a26003546002546040516001600160a01b0392831692909116907f70e091790811d89fdfce2f92748c357fdb8847d6e97f65b45fecf8d412d32deb90600090a350565b610d90336000356001600160e01b03191661188b565b610d986119cd565b6001600160a01b031663222e1b1683836040518363ffffffff1660e01b8152600401610dc5929190612731565b600060405180830381600087803b158015610ddf57600080fd5b505af1158015610df3573d6000803e3d6000fd5b505050505050565b6000610e05611862565b6001600160a01b03166322e594ff836040518263ffffffff1660e01b8152600401610940919061211b565b6060610e3a611862565b6001600160a01b0316632a5c792a6040518163ffffffff1660e01b8152600401600060405180830381865afa158015610e77573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e9f9190810190612920565b905090565b610eba336000356001600160e01b03191661188b565b610ec26119cd565b6001600160a01b0316632f73e6e4826040518263ffffffff1660e01b8152600401610ab9919061211b565b6060610ef7611862565b6001600160a01b0316633b937349836040518263ffffffff1660e01b8152600401610f2291906120a7565b600060405180830381865afa158015610f3f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109059190810190612920565b6000610f71611862565b6001600160a01b03166345741567836040518263ffffffff1660e01b81526004016108c4919061211b565b610fb2336000356001600160e01b03191661188b565b610fba611862565b6040516347ecc93960e01b8152600481018490526001600160a01b03838116602483015291909116906347ecc93990604401610dc5565b6000610ffb6119f7565b805490915060ff600160401b82041615906001600160401b03166000811580156110225750825b90506000826001600160401b0316600114801561103e5750303b155b90508115801561104c575080155b1561106a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561109457845460ff60401b1916600160401b1785555b61109c611a20565b6110a587611679565b60036001556110b386611a2a565b83156110f957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b611118336000356001600160e01b03191661188b565b6111206119cd565b6001600160a01b0316634da8e6c3826040518263ffffffff1660e01b8152600401610ab9919061211b565b611153611a6b565b61115c82611af9565b6111668282611b12565b5050565b6000611174611bcb565b50600080516020612b0e8339815191525b90565b61119e336000356001600160e01b03191661188b565b6111a6611862565b6001600160a01b0316635fa7b584826040518263ffffffff1660e01b8152600401610ab9919061211b565b60006111db611862565b6001600160a01b03166364c0673a836040518263ffffffff1660e01b81526004016108c4919061211b565b6060611210611862565b6001600160a01b0316636f0656f0836040518263ffffffff1660e01b8152600401610f2291906120a7565b6000611245611862565b6001600160a01b03166378a895676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611282573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9f91906129e2565b60006112be336000356001600160e01b03191661188b565b6112c66119cd565b6040516378dd4e6f60e01b815260048101859052602481018490526001600160a01b0391909116906378dd4e6f90604401602060405180830381865afa158015611314573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113389190612642565b9392505050565b611355336000356001600160e01b03191661188b565b61135d6119cd565b6001600160a01b0316637ae93604826040518263ffffffff1660e01b8152600401610ab9919061211b565b61139e336000356001600160e01b03191661188b565b6113a6611862565b6001600160a01b0316637fbc532d826040518263ffffffff1660e01b8152600401610ab9919061211b565b6113d9611fb4565b6113e1611862565b6001600160a01b0316638a58b79f836040518263ffffffff1660e01b815260040161140e91815260200190565b600060405180830381865afa15801561142b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261090591908101906129fb565b61145b611fb4565b611463611862565b6001600160a01b03166391ded8fa836040518263ffffffff1660e01b815260040161140e919061211b565b6000611498611862565b6001600160a01b03166393d857276040518163ffffffff1660e01b8152600401602060405180830381865afa158015611282573d6000803e3d6000fd5b6114eb336000356001600160e01b03191661188b565b6114f36119cd565b6001600160a01b0316639eb7b4eb826040518263ffffffff1660e01b8152600401610ab99190612a2f565b6000611528611c14565b9050336001600160a01b0382161461155f573381604051620d23e560e01b8152600401611556929190612a42565b60405180910390fd5b50600155565b61157b336000356001600160e01b03191661188b565b611583611862565b6001600160a01b031663a0b9b9b183836040518363ffffffff1660e01b8152600401610dc5929190612a42565b60006115ba611862565b6001600160a01b031663b33f78ca836040518263ffffffff1660e01b81526004016108c4919061211b565b60006115ef611cab565b546001600160a01b0316919050565b611614336000356001600160e01b03191661188b565b61161c6119cd565b60405163030c0d3b60e61b815260048101849052602481018390526001600160a01b03919091169063c3034ec09060440160006040518083038186803b15801561166557600080fd5b505afa158015610df3573d6000803e3d6000fd5b611681611d0d565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6116b9336000356001600160e01b03191661188b565b6116c1611862565b6001600160a01b031663d8dc051083836040518363ffffffff1660e01b8152600401610dc5929190612a5c565b611704336000356001600160e01b03191661188b565b61170c6119cd565b6001600160a01b031663de0c36256040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561174657600080fd5b505af115801561175a573d6000803e3d6000fd5b50505050565b600061176a611862565b6001600160a01b031663e3ade0ec836040518263ffffffff1660e01b8152600401610940919061211b565b61179d611fb4565b6117a5611862565b6001600160a01b031663efa74f1f836040518263ffffffff1660e01b815260040161140e919061259c565b6117e6336000356001600160e01b03191661188b565b6117ee611862565b6001600160a01b031663f05ab9de826040518263ffffffff1660e01b8152600401610ab9919061211b565b61182f336000356001600160e01b03191661188b565b6118376119cd565b6001600160a01b031663f5aeb434826040518263ffffffff1660e01b8152600401610ab99190612a2f565b6002546001600160a01b0316806111855760405162dee40560e41b815260040160405180910390fd5b6000611895611cab565b80549091506001600160a01b0316806118c4576000604051638944034760e01b8152600401611556919061211b565b60405163b700961360e01b81526001600160a01b0385811660048301523060248301526001600160e01b031985166044830152600091829182919085169063b700961390606401606060405180830381865afa158015611928573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061194c9190612a82565b925092509250826110f95780156119765760405163cc9855ad60e01b815260040160405180910390fd5b63ffffffff8216156119b25760405163a426878960e01b81526001600160a01b038816600482015263ffffffff83166024820152604401611556565b86604051632ecd3d0360e21b8152600401611556919061211b565b6003546001600160a01b03168061118557604051635e11777d60e01b815260040160405180910390fd5b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610905565b611a28611d0d565b565b6000611a34611cab565b80549091506001600160a01b031615611a625781604051638944034760e01b8152600401611556919061211b565b61116682611d32565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480611adb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611acf611dc2565b6001600160a01b031614155b15611a285760405163703e46dd60e11b815260040160405180910390fd5b611b0f336000356001600160e01b03191661188b565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611b6c575060408051601f3d908101601f19168201909252611b69918101906129e2565b60015b611b8b5781604051634c9c8ce360e01b8152600401611556919061211b565b600080516020612b0e8339815191528114611bbc57604051632a87526960e21b815260048101829052602401611556565b611bc68383611dd8565b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611a285760405163703e46dd60e11b815260040160405180910390fd5b600080546040516311f50c8560e01b8152600360048201526001600160a01b03909116906311f50c8590602401602060405180830381865afa158015611c5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c829190612690565b90506001600160a01b03811661118557604051633eba255b60e01b815260040160405180910390fd5b60008060ff19611cdc60017f2d9a26166e6ed6927c3e14d89c4d572a357f984344fcaed3c38c3ff5cdb83f35612ad0565b604051602001611cee91815260200190565b60408051601f1981840301815291905280516020909101201692915050565b611d15611e2e565b611a2857604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b038116611d5b5780604051638944034760e01b8152600401611556919061211b565b6000611d65611cab565b80546040519192506001600160a01b03808516929116907fa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b7638998019890600090a380546001600160a01b0319166001600160a01b0392909216919091179055565b6000600080516020612b0e8339815191526115ef565b611de182611e48565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115611e2657611bc68282611ea4565b611166611f1a565b6000611e386119f7565b54600160401b900460ff16919050565b806001600160a01b03163b600003611e755780604051634c9c8ce360e01b8152600401611556919061211b565b600080516020612b0e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051611ec19190612af1565b600060405180830381855af49150503d8060008114611efc576040519150601f19603f3d011682016040523d82523d6000602084013e611f01565b606091505b5091509150611f11858383611f39565b95945050505050565b3415611a285760405163b398979f60e01b815260040160405180910390fd5b606082611f4e57611f4982611f8c565b611338565b8151158015611f6557506001600160a01b0384163b155b15611f855783604051639996b31560e01b8152600401611556919061211b565b5092915050565b805115611f9b57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b604051806101a001604052806000801916815260200160608152602001606081526020016060815260200160006001600160a01b0316815260200160006001600160a01b03168152602001600081526020016000600c81111561201957612019612074565b815260200160008152602001600081526020016000815260200160008152602001600081525090565b6001600160a01b0381168114611b0f57600080fd5b60006020828403121561206957600080fd5b813561133881612042565b634e487b7160e01b600052602160045260246000fd5b60058110611b0f57611b0f612074565b6120a38161208a565b9052565b602081016120b48361208a565b91905290565b6000806000606084860312156120cf57600080fd5b8335925060208401356120e181612042565b9150604084013560ff811681146120f757600080fd5b809150509250925092565b60006020828403121561211457600080fd5b5035919050565b6001600160a01b0391909116815260200190565b6000806020838503121561214257600080fd5b82356001600160401b038082111561215957600080fd5b818501915085601f83011261216d57600080fd5b81358181111561217c57600080fd5b8660208260051b850101111561219157600080fd5b60209290920196919550909350505050565b60005b838110156121be5781810151838201526020016121a6565b50506000910152565b600081518084526121df8160208601602086016121a3565b601f01601f19169290920160200192915050565b600d81106120a3576120a3612074565b60006101a0825184526020830151816020860152612223828601826121c7565b9150506040830151848203604086015261223d82826121c7565b9150506060830151848203606086015261225782826121c7565b915050608083015161227460808601826001600160a01b03169052565b5060a083015161228f60a08601826001600160a01b03169052565b5060c083015160c085015260e08301516122ac60e08601826121f3565b50610100808401516122c08287018261209a565b5050610120808401516122d58287018261209a565b5050610140808401516122ea8287018261209a565b505061016083810151908501526101809283015192909301919091525090565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b8281101561236157603f1988860301845261234f858351612203565b94509285019290850190600101612333565b5092979650505050505050565b60058110611b0f57600080fd5b60006020828403121561238d57600080fd5b81356113388161236e565b600080604083850312156123ab57600080fd5b8235915060208301356123bd81612042565b809150509250929050565b600080604083850312156123db57600080fd5b82356123e681612042565b915060208301356123bd81612042565b634e487b7160e01b600052604160045260246000fd5b6040516101a081016001600160401b038111828210171561242f5761242f6123f6565b60405290565b604051601f8201601f191681016001600160401b038111828210171561245d5761245d6123f6565b604052919050565b60006001600160401b0382111561247e5761247e6123f6565b50601f01601f191660200190565b600061249f61249a84612465565b612435565b90508281528383830111156124b357600080fd5b828260208301376000602084830101529392505050565b600080604083850312156124dd57600080fd5b82356124e881612042565b915060208301356001600160401b0381111561250357600080fd5b8301601f8101851361251457600080fd5b6125238582356020840161248c565b9150509250929050565b6000806040838503121561254057600080fd5b50508035926020909101359150565b6020815260006113386020830184612203565b60006020828403121561257457600080fd5b81356001600160401b0381111561258a57600080fd5b82016040818503121561133857600080fd5b60208152600061133860208301846121c7565b600080604083850312156125c257600080fd5b82356125cd81612042565b915060208301356123bd8161236e565b6000602082840312156125ef57600080fd5b81356001600160401b0381111561260557600080fd5b8201601f8101841361261657600080fd5b6126258482356020840161248c565b949350505050565b8051801515811461263d57600080fd5b919050565b60006020828403121561265457600080fd5b6113388261262d565b805161263d8161236e565b60006020828403121561267a57600080fd5b81516113388161236e565b805161263d81612042565b6000602082840312156126a257600080fd5b815161133881612042565b8035825260006020820135601e198336030181126126ca57600080fd5b82016020810190356001600160401b038111156126e657600080fd5b8060051b8036038313156126f957600080fd5b60406020870181905286018290526001600160fb1b0382111561271b57600080fd5b8083606088013794909401606001949350505050565b60208082528181018390526000906040600585901b840181019084018684805b8881101561279557878503603f190184528235368b9003603e19018112612776578283fd5b612782868c83016126ad565b9550509285019291850191600101612751565b509298975050505050505050565b600082601f8301126127b457600080fd5b81516127c261249a82612465565b8181528460208386010111156127d757600080fd5b6126258260208301602087016121a3565b8051600d811061263d57600080fd5b60006101a0828403121561280a57600080fd5b61281261240c565b90508151815260208201516001600160401b038082111561283257600080fd5b61283e858386016127a3565b6020840152604084015191508082111561285757600080fd5b612863858386016127a3565b6040840152606084015191508082111561287c57600080fd5b50612889848285016127a3565b60608301525061289b60808301612685565b60808201526128ac60a08301612685565b60a082015260c082015160c08201526128c760e083016127e8565b60e08201526101006128da81840161265d565b908201526101206128ec83820161265d565b908201526101406128fe83820161265d565b9082015261016082810151908201526101809182015191810191909152919050565b6000602080838503121561293357600080fd5b82516001600160401b038082111561294a57600080fd5b818501915085601f83011261295e57600080fd5b815181811115612970576129706123f6565b8060051b61297f858201612435565b918252838101850191858101908984111561299957600080fd5b86860192505b838310156129d5578251858111156129b75760008081fd5b6129c58b89838a01016127f7565b835250918601919086019061299f565b9998505050505050505050565b6000602082840312156129f457600080fd5b5051919050565b600060208284031215612a0d57600080fd5b81516001600160401b03811115612a2357600080fd5b612625848285016127f7565b60208152600061133860208301846126ad565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b038316815260408101612a758361208a565b8260208301529392505050565b600080600060608486031215612a9757600080fd5b612aa08461262d565b9250602084015163ffffffff81168114612ab957600080fd5b9150612ac76040850161262d565b90509250925092565b8181038181111561090557634e487b7160e01b600052601160045260246000fd5b60008251612b038184602087016121a3565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122015bef19667efe5f791835d88c58ee5e02af400e80bea7aee72cda06e12ad217e64736f6c63430008180033",
}

// PNTokenRegistryV1 is an auto generated Go binding around an Ethereum contract.
type PNTokenRegistryV1 struct {
	abi abi.ABI
}

// NewPNTokenRegistryV1 creates a new instance of PNTokenRegistryV1.
func NewPNTokenRegistryV1() *PNTokenRegistryV1 {
	parsed, err := PNTokenRegistryV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &PNTokenRegistryV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *PNTokenRegistryV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackActivateToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x060ed143.
//
// Solidity: function activateToken(bytes32 tokenResourceId, address tokenAddress, uint8 ercStandard) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackActivateToken(tokenResourceId [32]byte, tokenAddress common.Address, ercStandard uint8) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("activateToken", tokenResourceId, tokenAddress, ercStandard)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAuthority is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackAuthority() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("authority", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackContractVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackContractVersion() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("contractVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackContractVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackContractVersion(data []byte) (*big.Int, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("contractVersion", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackEndpoint() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("endpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackEndpoint(data []byte) (common.Address, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("endpoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackFreezeOnPrivacyNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f73e6e4.
//
// Solidity: function freezeOnPrivacyNode(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackFreezeOnPrivacyNode(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("freezeOnPrivacyNode", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackFreezeOnPublicChain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4da8e6c3.
//
// Solidity: function freezeOnPublicChain(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackFreezeOnPublicChain(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("freezeOnPublicChain", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetActiveTokenCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x93d85727.
//
// Solidity: function getActiveTokenCount() view returns(uint256)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetActiveTokenCount() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getActiveTokenCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetActiveTokenCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x93d85727.
//
// Solidity: function getActiveTokenCount() view returns(uint256)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetActiveTokenCount(data []byte) (*big.Int, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getActiveTokenCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetAddressByResourceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11f50c85.
//
// Solidity: function getAddressByResourceId(bytes32 _resourceId) view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetAddressByResourceId(resourceId [32]byte) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getAddressByResourceId", resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAddressByResourceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11f50c85.
//
// Solidity: function getAddressByResourceId(bytes32 _resourceId) view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetAddressByResourceId(data []byte) (common.Address, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getAddressByResourceId", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetAllTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256)[])
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetAllTokens() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getAllTokens")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAllTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256)[])
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetAllTokens(data []byte) ([]TokenStructsToken, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getAllTokens", data)
	if err != nil {
		return *new([]TokenStructsToken), err
	}
	out0 := *abi.ConvertType(out[0], new([]TokenStructsToken)).(*[]TokenStructsToken)
	return out0, err
}

// PackGetFrozenTokenForParticipant is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78dd4e6f.
//
// Solidity: function getFrozenTokenForParticipant(bytes32 tokenResourceId, uint256 chainId) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetFrozenTokenForParticipant(tokenResourceId [32]byte, chainId *big.Int) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getFrozenTokenForParticipant", tokenResourceId, chainId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFrozenTokenForParticipant is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x78dd4e6f.
//
// Solidity: function getFrozenTokenForParticipant(bytes32 tokenResourceId, uint256 chainId) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetFrozenTokenForParticipant(data []byte) (bool, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getFrozenTokenForParticipant", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackGetHubStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe3ade0ec.
//
// Solidity: function getHubStatus(address tokenAddress) view returns(uint8)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetHubStatus(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getHubStatus", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetHubStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe3ade0ec.
//
// Solidity: function getHubStatus(address tokenAddress) view returns(uint8)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetHubStatus(data []byte) (uint8, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getHubStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackGetPrivacyNodeStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22e594ff.
//
// Solidity: function getPrivacyNodeStatus(address tokenAddress) view returns(uint8)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetPrivacyNodeStatus(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getPrivacyNodeStatus", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPrivacyNodeStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x22e594ff.
//
// Solidity: function getPrivacyNodeStatus(address tokenAddress) view returns(uint8)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetPrivacyNodeStatus(data []byte) (uint8, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getPrivacyNodeStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackGetPublicChainStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01162d43.
//
// Solidity: function getPublicChainStatus(address tokenAddress) view returns(uint8)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetPublicChainStatus(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getPublicChainStatus", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPublicChainStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01162d43.
//
// Solidity: function getPublicChainStatus(address tokenAddress) view returns(uint8)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetPublicChainStatus(data []byte) (uint8, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getPublicChainStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackGetTokenByAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91ded8fa.
//
// Solidity: function getTokenByAddress(address tokenAddress) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256))
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokenByAddress(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokenByAddress", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenByAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91ded8fa.
//
// Solidity: function getTokenByAddress(address tokenAddress) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256))
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokenByAddress(data []byte) (TokenStructsToken, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokenByAddress", data)
	if err != nil {
		return *new(TokenStructsToken), err
	}
	out0 := *abi.ConvertType(out[0], new(TokenStructsToken)).(*TokenStructsToken)
	return out0, err
}

// PackGetTokenByResourceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a58b79f.
//
// Solidity: function getTokenByResourceId(bytes32 tokenResourceId) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256))
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokenByResourceId(tokenResourceId [32]byte) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokenByResourceId", tokenResourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenByResourceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a58b79f.
//
// Solidity: function getTokenByResourceId(bytes32 tokenResourceId) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256))
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokenByResourceId(data []byte) (TokenStructsToken, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokenByResourceId", data)
	if err != nil {
		return *new(TokenStructsToken), err
	}
	out0 := *abi.ConvertType(out[0], new(TokenStructsToken)).(*TokenStructsToken)
	return out0, err
}

// PackGetTokenBySymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefa74f1f.
//
// Solidity: function getTokenBySymbol(string symbol) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256))
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokenBySymbol(symbol string) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokenBySymbol", symbol)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenBySymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xefa74f1f.
//
// Solidity: function getTokenBySymbol(string symbol) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256))
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokenBySymbol(data []byte) (TokenStructsToken, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokenBySymbol", data)
	if err != nil {
		return *new(TokenStructsToken), err
	}
	out0 := *abi.ConvertType(out[0], new(TokenStructsToken)).(*TokenStructsToken)
	return out0, err
}

// PackGetTokenCore is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf0a599bb.
//
// Solidity: function getTokenCore() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokenCore() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokenCore")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenCore is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf0a599bb.
//
// Solidity: function getTokenCore() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokenCore(data []byte) (common.Address, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokenCore", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetTokenCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokenCount() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokenCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokenCount(data []byte) (*big.Int, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokenCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetTokenFreezeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf356cc03.
//
// Solidity: function getTokenFreezeManager() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokenFreezeManager() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokenFreezeManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenFreezeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf356cc03.
//
// Solidity: function getTokenFreezeManager() view returns(address)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokenFreezeManager(data []byte) (common.Address, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokenFreezeManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetTokensByHubStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f0656f0.
//
// Solidity: function getTokensByHubStatus(uint8 status) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256)[])
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokensByHubStatus(status uint8) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokensByHubStatus", status)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokensByHubStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f0656f0.
//
// Solidity: function getTokensByHubStatus(uint8 status) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256)[])
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokensByHubStatus(data []byte) ([]TokenStructsToken, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokensByHubStatus", data)
	if err != nil {
		return *new([]TokenStructsToken), err
	}
	out0 := *abi.ConvertType(out[0], new([]TokenStructsToken)).(*[]TokenStructsToken)
	return out0, err
}

// PackGetTokensByPrivacyNodeStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b937349.
//
// Solidity: function getTokensByPrivacyNodeStatus(uint8 status) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256)[])
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackGetTokensByPrivacyNodeStatus(status uint8) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("getTokensByPrivacyNodeStatus", status)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokensByPrivacyNodeStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3b937349.
//
// Solidity: function getTokensByPrivacyNodeStatus(uint8 status) view returns((bytes32,string,string,string,address,address,uint256,uint8,uint8,uint8,uint8,uint256,uint256)[])
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackGetTokensByPrivacyNodeStatus(data []byte) ([]TokenStructsToken, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("getTokensByPrivacyNodeStatus", data)
	if err != nil {
		return *new([]TokenStructsToken), err
	}
	out0 := *abi.ConvertType(out[0], new([]TokenStructsToken)).(*[]TokenStructsToken)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x485cc955.
//
// Solidity: function initialize(address endpointAddress, address authority_) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackInitialize(endpointAddress common.Address, authority common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("initialize", endpointAddress, authority)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInitialize0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address _endpoint) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackInitialize0(endpoint common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("initialize0", endpoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsTokenActiveForHub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0014bc46.
//
// Solidity: function isTokenActiveForHub(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackIsTokenActiveForHub(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("isTokenActiveForHub", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsTokenActiveForHub is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0014bc46.
//
// Solidity: function isTokenActiveForHub(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackIsTokenActiveForHub(data []byte) (bool, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("isTokenActiveForHub", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsTokenActiveForPublicChain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64c0673a.
//
// Solidity: function isTokenActiveForPublicChain(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackIsTokenActiveForPublicChain(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("isTokenActiveForPublicChain", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsTokenActiveForPublicChain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x64c0673a.
//
// Solidity: function isTokenActiveForPublicChain(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackIsTokenActiveForPublicChain(data []byte) (bool, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("isTokenActiveForPublicChain", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsTokenFullyOperational is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45741567.
//
// Solidity: function isTokenFullyOperational(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackIsTokenFullyOperational(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("isTokenFullyOperational", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsTokenFullyOperational is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x45741567.
//
// Solidity: function isTokenFullyOperational(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackIsTokenFullyOperational(data []byte) (bool, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("isTokenFullyOperational", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackProxiableUUID() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRegisterHubToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47ecc939.
//
// Solidity: function registerHubToken(bytes32 tokenResourceId, address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackRegisterHubToken(tokenResourceId [32]byte, tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("registerHubToken", tokenResourceId, tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRegisterToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09824a80.
//
// Solidity: function registerToken(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackRegisterToken(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("registerToken", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRejectToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf05ab9de.
//
// Solidity: function rejectToken(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackRejectToken(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("rejectToken", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemoveFrozenToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5aeb434.
//
// Solidity: function removeFrozenToken((bytes32,uint256[]) unfrozenToken) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackRemoveFrozenToken(unfrozenToken TokenStructsFrozenToken) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("removeFrozenToken", unfrozenToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemoveToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5fa7b584.
//
// Solidity: function removeToken(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackRemoveToken(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("removeToken", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestAllFrozenTokensDataFromPrivateHub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde0c3625.
//
// Solidity: function requestAllFrozenTokensDataFromPrivateHub() returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackRequestAllFrozenTokensDataFromPrivateHub() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("requestAllFrozenTokensDataFromPrivateHub")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackResourceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f997c5b.
//
// Solidity: function resourceId() view returns(bytes32)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackResourceId() []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("resourceId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackResourceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f997c5b.
//
// Solidity: function resourceId() view returns(bytes32)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackResourceId(data []byte) ([32]byte, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("resourceId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackSetResourceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa01afbfb.
//
// Solidity: function setResourceId(bytes32 _resourceId) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackSetResourceId(resourceId [32]byte) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("setResourceId", resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTokenCore is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x14ae646b.
//
// Solidity: function setTokenCore(address tokenCoreAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackSetTokenCore(tokenCoreAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("setTokenCore", tokenCoreAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTokenFreezeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13aa1a1d.
//
// Solidity: function setTokenFreezeManager(address freezeManagerAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackSetTokenFreezeManager(freezeManagerAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("setTokenFreezeManager", freezeManagerAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSubmitToHub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fbc532d.
//
// Solidity: function submitToHub(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackSubmitToHub(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("submitToHub", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSubmitToPublicChain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a75cf4e.
//
// Solidity: function submitToPublicChain(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackSubmitToPublicChain(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("submitToPublicChain", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSyncFrozenTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x222e1b16.
//
// Solidity: function syncFrozenTokens((bytes32,uint256[])[] frozenTokens) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackSyncFrozenTokens(frozenTokens []TokenStructsFrozenToken) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("syncFrozenTokens", frozenTokens)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTokenExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb33f78ca.
//
// Solidity: function tokenExists(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackTokenExists(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("tokenExists", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenExists is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb33f78ca.
//
// Solidity: function tokenExists(address tokenAddress) view returns(bool)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenExists(data []byte) (bool, error) {
	out, err := pNTokenRegistryV1.abi.Unpack("tokenExists", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUnfreezeOnPrivacyNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10dcd520.
//
// Solidity: function unfreezeOnPrivacyNode(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackUnfreezeOnPrivacyNode(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("unfreezeOnPrivacyNode", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnfreezeOnPublicChain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ae93604.
//
// Solidity: function unfreezeOnPublicChain(address tokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackUnfreezeOnPublicChain(tokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("unfreezeOnPublicChain", tokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateFrozenToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9eb7b4eb.
//
// Solidity: function updateFrozenToken((bytes32,uint256[]) frozenToken) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackUpdateFrozenToken(frozenToken TokenStructsFrozenToken) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("updateFrozenToken", frozenToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdatePrivacyNodeStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8dc0510.
//
// Solidity: function updatePrivacyNodeStatus(address tokenAddress, uint8 status) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackUpdatePrivacyNodeStatus(tokenAddress common.Address, status uint8) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("updatePrivacyNodeStatus", tokenAddress, status)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdatePublicTokenAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0b9b9b1.
//
// Solidity: function updatePublicTokenAddress(address tokenAddress, address publicTokenAddress) returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackUpdatePublicTokenAddress(tokenAddress common.Address, publicTokenAddress common.Address) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("updatePublicTokenAddress", tokenAddress, publicTokenAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidateTokenForParticipant is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3034ec0.
//
// Solidity: function validateTokenForParticipant(bytes32 tokenResourceId, uint256 chainId) view returns()
func (pNTokenRegistryV1 *PNTokenRegistryV1) PackValidateTokenForParticipant(tokenResourceId [32]byte, chainId *big.Int) []byte {
	enc, err := pNTokenRegistryV1.abi.Pack("validateTokenForParticipant", tokenResourceId, chainId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PNTokenRegistryV1AuthorityUpdated represents a AuthorityUpdated event raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const PNTokenRegistryV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (PNTokenRegistryV1AuthorityUpdated) ContractEventName() string {
	return PNTokenRegistryV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*PNTokenRegistryV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != pNTokenRegistryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PNTokenRegistryV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pNTokenRegistryV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PNTokenRegistryV1Initialized represents a Initialized event raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const PNTokenRegistryV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (PNTokenRegistryV1Initialized) ContractEventName() string {
	return PNTokenRegistryV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackInitializedEvent(log *types.Log) (*PNTokenRegistryV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != pNTokenRegistryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PNTokenRegistryV1Initialized)
	if len(log.Data) > 0 {
		if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pNTokenRegistryV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PNTokenRegistryV1TokenCoreSet represents a TokenCoreSet event raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1TokenCoreSet struct {
	TokenCore common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const PNTokenRegistryV1TokenCoreSetEventName = "TokenCoreSet"

// ContractEventName returns the user-defined event name.
func (PNTokenRegistryV1TokenCoreSet) ContractEventName() string {
	return PNTokenRegistryV1TokenCoreSetEventName
}

// UnpackTokenCoreSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenCoreSet(address indexed tokenCore)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenCoreSetEvent(log *types.Log) (*PNTokenRegistryV1TokenCoreSet, error) {
	event := "TokenCoreSet"
	if log.Topics[0] != pNTokenRegistryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PNTokenRegistryV1TokenCoreSet)
	if len(log.Data) > 0 {
		if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pNTokenRegistryV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PNTokenRegistryV1TokenFreezeManagerSet represents a TokenFreezeManagerSet event raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1TokenFreezeManagerSet struct {
	TokenFreezeManager common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const PNTokenRegistryV1TokenFreezeManagerSetEventName = "TokenFreezeManagerSet"

// ContractEventName returns the user-defined event name.
func (PNTokenRegistryV1TokenFreezeManagerSet) ContractEventName() string {
	return PNTokenRegistryV1TokenFreezeManagerSetEventName
}

// UnpackTokenFreezeManagerSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenFreezeManagerSet(address indexed tokenFreezeManager)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenFreezeManagerSetEvent(log *types.Log) (*PNTokenRegistryV1TokenFreezeManagerSet, error) {
	event := "TokenFreezeManagerSet"
	if log.Topics[0] != pNTokenRegistryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PNTokenRegistryV1TokenFreezeManagerSet)
	if len(log.Data) > 0 {
		if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pNTokenRegistryV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PNTokenRegistryV1TokenRegistryModulesConfigured represents a TokenRegistryModulesConfigured event raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1TokenRegistryModulesConfigured struct {
	TokenCore          common.Address
	TokenFreezeManager common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const PNTokenRegistryV1TokenRegistryModulesConfiguredEventName = "TokenRegistryModulesConfigured"

// ContractEventName returns the user-defined event name.
func (PNTokenRegistryV1TokenRegistryModulesConfigured) ContractEventName() string {
	return PNTokenRegistryV1TokenRegistryModulesConfiguredEventName
}

// UnpackTokenRegistryModulesConfiguredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenRegistryModulesConfigured(address indexed tokenCore, address indexed tokenFreezeManager)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenRegistryModulesConfiguredEvent(log *types.Log) (*PNTokenRegistryV1TokenRegistryModulesConfigured, error) {
	event := "TokenRegistryModulesConfigured"
	if log.Topics[0] != pNTokenRegistryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PNTokenRegistryV1TokenRegistryModulesConfigured)
	if len(log.Data) > 0 {
		if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pNTokenRegistryV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PNTokenRegistryV1Upgraded represents a Upgraded event raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const PNTokenRegistryV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (PNTokenRegistryV1Upgraded) ContractEventName() string {
	return PNTokenRegistryV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackUpgradedEvent(log *types.Log) (*PNTokenRegistryV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != pNTokenRegistryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PNTokenRegistryV1Upgraded)
	if len(log.Data) > 0 {
		if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pNTokenRegistryV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAppV1HubNotActive"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAppV1HubNotActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAppV1PrivacyNodeFrozen"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAppV1PrivacyNodeFrozenError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAppV1PrivacyNodeNotActive"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAppV1PrivacyNodeNotActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAppV1PublicChainNotActive"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAppV1PublicChainNotActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAppV1ResourceNotApproved"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAppV1ResourceNotApprovedError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAppV1TokenRegistryNotConfigured"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAppV1TokenRegistryNotConfiguredError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["RaylsAppV1UnauthorizedTokenRegistry"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackRaylsAppV1UnauthorizedTokenRegistryError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["TokenRegistryV1InvalidTokenCoreAddress"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackTokenRegistryV1InvalidTokenCoreAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["TokenRegistryV1InvalidTokenFreezeManagerAddress"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackTokenRegistryV1InvalidTokenFreezeManagerAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["TokenRegistryV1TokenCoreNotConfigured"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackTokenRegistryV1TokenCoreNotConfiguredError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["TokenRegistryV1TokenFreezeManagerNotConfigured"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackTokenRegistryV1TokenFreezeManagerNotConfiguredError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], pNTokenRegistryV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return pNTokenRegistryV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PNTokenRegistryV1AddressEmptyCode represents a AddressEmptyCode error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func PNTokenRegistryV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackAddressEmptyCodeError(raw []byte) (*PNTokenRegistryV1AddressEmptyCode, error) {
	out := new(PNTokenRegistryV1AddressEmptyCode)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func PNTokenRegistryV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackERC1967InvalidImplementationError(raw []byte) (*PNTokenRegistryV1ERC1967InvalidImplementation, error) {
	out := new(PNTokenRegistryV1ERC1967InvalidImplementation)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func PNTokenRegistryV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackERC1967NonPayableError(raw []byte) (*PNTokenRegistryV1ERC1967NonPayable, error) {
	out := new(PNTokenRegistryV1ERC1967NonPayable)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1FailedCall represents a FailedCall error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func PNTokenRegistryV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackFailedCallError(raw []byte) (*PNTokenRegistryV1FailedCall, error) {
	out := new(PNTokenRegistryV1FailedCall)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1InvalidInitialization represents a InvalidInitialization error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func PNTokenRegistryV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackInvalidInitializationError(raw []byte) (*PNTokenRegistryV1InvalidInitialization, error) {
	out := new(PNTokenRegistryV1InvalidInitialization)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1NotInitializing represents a NotInitializing error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func PNTokenRegistryV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackNotInitializingError(raw []byte) (*PNTokenRegistryV1NotInitializing, error) {
	out := new(PNTokenRegistryV1NotInitializing)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func PNTokenRegistryV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*PNTokenRegistryV1RaylsAccessManagedContractPaused, error) {
	out := new(PNTokenRegistryV1RaylsAccessManagedContractPaused)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func PNTokenRegistryV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*PNTokenRegistryV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(PNTokenRegistryV1RaylsAccessManagedInvalidAuthority)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func PNTokenRegistryV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*PNTokenRegistryV1RaylsAccessManagedMustSchedule, error) {
	out := new(PNTokenRegistryV1RaylsAccessManagedMustSchedule)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func PNTokenRegistryV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*PNTokenRegistryV1RaylsAccessManagedUnauthorized, error) {
	out := new(PNTokenRegistryV1RaylsAccessManagedUnauthorized)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAppV1HubNotActive represents a RaylsAppV1__HubNotActive error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAppV1HubNotActive struct {
	TokenAddress      common.Address
	PrivacyNodeStatus uint8
	HubStatus         uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAppV1__HubNotActive(address tokenAddress, uint8 privacyNodeStatus, uint8 hubStatus)
func PNTokenRegistryV1RaylsAppV1HubNotActiveErrorID() common.Hash {
	return common.HexToHash("0x3fae5bbd70277aa1cd008dceb93b19a7055c2a6d29b84733371e1c41b2048b15")
}

// UnpackRaylsAppV1HubNotActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAppV1__HubNotActive(address tokenAddress, uint8 privacyNodeStatus, uint8 hubStatus)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAppV1HubNotActiveError(raw []byte) (*PNTokenRegistryV1RaylsAppV1HubNotActive, error) {
	out := new(PNTokenRegistryV1RaylsAppV1HubNotActive)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAppV1HubNotActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAppV1PrivacyNodeFrozen represents a RaylsAppV1__PrivacyNodeFrozen error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAppV1PrivacyNodeFrozen struct {
	TokenAddress common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAppV1__PrivacyNodeFrozen(address tokenAddress)
func PNTokenRegistryV1RaylsAppV1PrivacyNodeFrozenErrorID() common.Hash {
	return common.HexToHash("0xc80bd255e67000277f5aed4960b64f92e2d5a652f07a22fba7d044de6add8f0e")
}

// UnpackRaylsAppV1PrivacyNodeFrozenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAppV1__PrivacyNodeFrozen(address tokenAddress)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAppV1PrivacyNodeFrozenError(raw []byte) (*PNTokenRegistryV1RaylsAppV1PrivacyNodeFrozen, error) {
	out := new(PNTokenRegistryV1RaylsAppV1PrivacyNodeFrozen)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAppV1PrivacyNodeFrozen", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAppV1PrivacyNodeNotActive represents a RaylsAppV1__PrivacyNodeNotActive error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAppV1PrivacyNodeNotActive struct {
	TokenAddress      common.Address
	PrivacyNodeStatus uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAppV1__PrivacyNodeNotActive(address tokenAddress, uint8 privacyNodeStatus)
func PNTokenRegistryV1RaylsAppV1PrivacyNodeNotActiveErrorID() common.Hash {
	return common.HexToHash("0xfdcdd2a6e576bf1f342ce493560565ef686a59cd3e0486f6869151efb2c7853f")
}

// UnpackRaylsAppV1PrivacyNodeNotActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAppV1__PrivacyNodeNotActive(address tokenAddress, uint8 privacyNodeStatus)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAppV1PrivacyNodeNotActiveError(raw []byte) (*PNTokenRegistryV1RaylsAppV1PrivacyNodeNotActive, error) {
	out := new(PNTokenRegistryV1RaylsAppV1PrivacyNodeNotActive)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAppV1PrivacyNodeNotActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAppV1PublicChainNotActive represents a RaylsAppV1__PublicChainNotActive error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAppV1PublicChainNotActive struct {
	TokenAddress      common.Address
	PrivacyNodeStatus uint8
	PublicChainStatus uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAppV1__PublicChainNotActive(address tokenAddress, uint8 privacyNodeStatus, uint8 publicChainStatus)
func PNTokenRegistryV1RaylsAppV1PublicChainNotActiveErrorID() common.Hash {
	return common.HexToHash("0xb607961611e6e4126e09c80bcd1e35e7a1e886888daa292eecc27cd9d4e37f3f")
}

// UnpackRaylsAppV1PublicChainNotActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAppV1__PublicChainNotActive(address tokenAddress, uint8 privacyNodeStatus, uint8 publicChainStatus)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAppV1PublicChainNotActiveError(raw []byte) (*PNTokenRegistryV1RaylsAppV1PublicChainNotActive, error) {
	out := new(PNTokenRegistryV1RaylsAppV1PublicChainNotActive)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAppV1PublicChainNotActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAppV1ResourceNotApproved represents a RaylsAppV1__ResourceNotApproved error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAppV1ResourceNotApproved struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAppV1__ResourceNotApproved()
func PNTokenRegistryV1RaylsAppV1ResourceNotApprovedErrorID() common.Hash {
	return common.HexToHash("0x8f144935367c131b72d26b0320b764f69ba3639e65abb1c811084bbd46e5c731")
}

// UnpackRaylsAppV1ResourceNotApprovedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAppV1__ResourceNotApproved()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAppV1ResourceNotApprovedError(raw []byte) (*PNTokenRegistryV1RaylsAppV1ResourceNotApproved, error) {
	out := new(PNTokenRegistryV1RaylsAppV1ResourceNotApproved)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAppV1ResourceNotApproved", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAppV1TokenRegistryNotConfigured represents a RaylsAppV1__TokenRegistryNotConfigured error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAppV1TokenRegistryNotConfigured struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAppV1__TokenRegistryNotConfigured()
func PNTokenRegistryV1RaylsAppV1TokenRegistryNotConfiguredErrorID() common.Hash {
	return common.HexToHash("0x3eba255b70fc7afd9cc5be90de2023dae8350ac3c29cbd5eaf139cadd9c4292e")
}

// UnpackRaylsAppV1TokenRegistryNotConfiguredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAppV1__TokenRegistryNotConfigured()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAppV1TokenRegistryNotConfiguredError(raw []byte) (*PNTokenRegistryV1RaylsAppV1TokenRegistryNotConfigured, error) {
	out := new(PNTokenRegistryV1RaylsAppV1TokenRegistryNotConfigured)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAppV1TokenRegistryNotConfigured", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1RaylsAppV1UnauthorizedTokenRegistry represents a RaylsAppV1__UnauthorizedTokenRegistry error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1RaylsAppV1UnauthorizedTokenRegistry struct {
	Caller        common.Address
	TokenRegistry common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAppV1__UnauthorizedTokenRegistry(address caller, address tokenRegistry)
func PNTokenRegistryV1RaylsAppV1UnauthorizedTokenRegistryErrorID() common.Hash {
	return common.HexToHash("0x000d23e5a298a9951b289bd8f5eece62aa717c000d6b0509a9f77d16f67a5b7d")
}

// UnpackRaylsAppV1UnauthorizedTokenRegistryError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAppV1__UnauthorizedTokenRegistry(address caller, address tokenRegistry)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackRaylsAppV1UnauthorizedTokenRegistryError(raw []byte) (*PNTokenRegistryV1RaylsAppV1UnauthorizedTokenRegistry, error) {
	out := new(PNTokenRegistryV1RaylsAppV1UnauthorizedTokenRegistry)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "RaylsAppV1UnauthorizedTokenRegistry", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1TokenRegistryV1InvalidTokenCoreAddress represents a TokenRegistryV1__InvalidTokenCoreAddress error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1TokenRegistryV1InvalidTokenCoreAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenRegistryV1__InvalidTokenCoreAddress()
func PNTokenRegistryV1TokenRegistryV1InvalidTokenCoreAddressErrorID() common.Hash {
	return common.HexToHash("0x743840cd612f7e36e5814125efeb3c972964e780bff028a10d4e7ba4b7bd47f3")
}

// UnpackTokenRegistryV1InvalidTokenCoreAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenRegistryV1__InvalidTokenCoreAddress()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenRegistryV1InvalidTokenCoreAddressError(raw []byte) (*PNTokenRegistryV1TokenRegistryV1InvalidTokenCoreAddress, error) {
	out := new(PNTokenRegistryV1TokenRegistryV1InvalidTokenCoreAddress)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "TokenRegistryV1InvalidTokenCoreAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1TokenRegistryV1InvalidTokenFreezeManagerAddress represents a TokenRegistryV1__InvalidTokenFreezeManagerAddress error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1TokenRegistryV1InvalidTokenFreezeManagerAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenRegistryV1__InvalidTokenFreezeManagerAddress()
func PNTokenRegistryV1TokenRegistryV1InvalidTokenFreezeManagerAddressErrorID() common.Hash {
	return common.HexToHash("0xcbb9c80e5b7e9993353677692525f58d399465ade8d5206d324f23f7cca6eaf9")
}

// UnpackTokenRegistryV1InvalidTokenFreezeManagerAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenRegistryV1__InvalidTokenFreezeManagerAddress()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenRegistryV1InvalidTokenFreezeManagerAddressError(raw []byte) (*PNTokenRegistryV1TokenRegistryV1InvalidTokenFreezeManagerAddress, error) {
	out := new(PNTokenRegistryV1TokenRegistryV1InvalidTokenFreezeManagerAddress)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "TokenRegistryV1InvalidTokenFreezeManagerAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1TokenRegistryV1TokenCoreNotConfigured represents a TokenRegistryV1__TokenCoreNotConfigured error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1TokenRegistryV1TokenCoreNotConfigured struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenRegistryV1__TokenCoreNotConfigured()
func PNTokenRegistryV1TokenRegistryV1TokenCoreNotConfiguredErrorID() common.Hash {
	return common.HexToHash("0x0dee4050ae8d6c6f1f485a566a6e78c577c2f1134e26cef638d9ecc511270fa2")
}

// UnpackTokenRegistryV1TokenCoreNotConfiguredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenRegistryV1__TokenCoreNotConfigured()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenRegistryV1TokenCoreNotConfiguredError(raw []byte) (*PNTokenRegistryV1TokenRegistryV1TokenCoreNotConfigured, error) {
	out := new(PNTokenRegistryV1TokenRegistryV1TokenCoreNotConfigured)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "TokenRegistryV1TokenCoreNotConfigured", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1TokenRegistryV1TokenFreezeManagerNotConfigured represents a TokenRegistryV1__TokenFreezeManagerNotConfigured error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1TokenRegistryV1TokenFreezeManagerNotConfigured struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenRegistryV1__TokenFreezeManagerNotConfigured()
func PNTokenRegistryV1TokenRegistryV1TokenFreezeManagerNotConfiguredErrorID() common.Hash {
	return common.HexToHash("0x5e11777d4fffe4b1cbb3c3552b2d82e0fb7c9d31e45257f293bfae903a2645b0")
}

// UnpackTokenRegistryV1TokenFreezeManagerNotConfiguredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenRegistryV1__TokenFreezeManagerNotConfigured()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackTokenRegistryV1TokenFreezeManagerNotConfiguredError(raw []byte) (*PNTokenRegistryV1TokenRegistryV1TokenFreezeManagerNotConfigured, error) {
	out := new(PNTokenRegistryV1TokenRegistryV1TokenFreezeManagerNotConfigured)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "TokenRegistryV1TokenFreezeManagerNotConfigured", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func PNTokenRegistryV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*PNTokenRegistryV1UUPSUnauthorizedCallContext, error) {
	out := new(PNTokenRegistryV1UUPSUnauthorizedCallContext)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PNTokenRegistryV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the PNTokenRegistryV1 contract.
type PNTokenRegistryV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func PNTokenRegistryV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (pNTokenRegistryV1 *PNTokenRegistryV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*PNTokenRegistryV1UUPSUnsupportedProxiableUUID, error) {
	out := new(PNTokenRegistryV1UUPSUnsupportedProxiableUUID)
	if err := pNTokenRegistryV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
