// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RNContractFactoryV1

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

// RNContractFactoryV1MetaData contains all meta data concerning the RNContractFactoryV1 contract.
var RNContractFactoryV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RAYLS_ENYGMA_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ENYGMA_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_DVP_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_DVP_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC20_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC20_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_DVP_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_DVP_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_STABLECOIN_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_STABLECOIN_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployEnygma\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployEnygmaAsUser\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc1155\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc1155AsUser\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc1155Dvp\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc20\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc20AsUser\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc721\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc721AsUser\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc721Dvp\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployExternal\",\"inputs\":[{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployFromTeleport\",\"inputs\":[{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"factoryKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ownerEOA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployRegistered\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployRegisteredAsUser\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployRegisteredExternal\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployStableCoin\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployStableCoinAsUser\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factoryOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBytecodeHash\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEndpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFactoryOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRaylsNodeEndpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_endpoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_raylsNodeEndpoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"authority_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raylsNodeEndpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBytecode\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEndpoint\",\"inputs\":[{\"name\":\"newEndpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFactoryOwner\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRaylsNodeEndpoint\",\"inputs\":[{\"name\":\"newEndpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenRegistry\",\"inputs\":[{\"name\":\"_tokenRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BytecodeSet\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"bytecodeHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractDeployed\",\"inputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EndpointUpdated\",\"inputs\":[{\"name\":\"oldEndpoint\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEndpoint\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FactoryOwnerUpdated\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RaylsNodeEndpointUpdated\",\"inputs\":[{\"name\":\"oldEndpoint\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEndpoint\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegisteredContractDeployed\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deployedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRegistrySet\",\"inputs\":[{\"name\":\"oldRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FactoryV1__BytecodeNotRegistered\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FactoryV1__EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FactoryV1__InitializationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FactoryV1__TokenRegistryNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FactoryV1__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitCodeStub__RuntimeTooLarge\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__MustSchedule\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:  "RNContractFactoryV1",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000030565b6200002a62000030565b620000e4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000815760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000e15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516130166200010e60003960008181611b3d01528181611b660152611c9f01526130166000f3fe6080604052600436106102685760003560e01c806395f50a6211610145578063cd39c4a8116100bc578063cd39c4a8146107b2578063d3fe9c16146107d2578063d64e6e8614610806578063dbbb415514610826578063dc905d0814610846578063e09bda6b14610866578063ee5ce29e14610884578063eff52a99146108b8578063f3cc660c146108d8578063f8c8765e146108f8578063fbc1905114610918578063fcab698214610938578063fe75427d1461096c57600080fd5b806395f50a62146106035780639c363c31146106375780639d23c4c7146106575780639ee1d00a14610677578063a0a8e46014610697578063a8126261146106ab578063ad10c149146106cb578063ad3cb1cc146106ed578063aed8e9671461072b578063b3872e4e14610749578063bd7e7d2f1461077d578063bf7e214f1461079d57600080fd5b80634f1ef286116101e45780634f1ef2861461045157806352d1902d146104645780635dd56210146104795780635e280f111461049b578063614bcfd3146104bb5780636670b8c5146104db5780636bc5d477146104fb5780636d7ffb32146105195780636e9bb21e1461053b5780636f5e71fe1461055b5780638f22b4a91461057b57806392c04d9b1461059b57806394f08898146105cf57600080fd5b8062866e721461026d57806302b6bcf2146102a35780630cd503ea146102c35780631464230b146102e357806316381b151461031357806319e011ad146103335780631c7fe28b1461035357806329b0d7011461038757806335a5af92146103bb57806339e297b2146103dd578063403be3de146103fd5780634273601c14610431575b600080fd5b34801561027957600080fd5b5061028d610288366004612437565b61098c565b60405161029a91906124aa565b60405180910390f35b3480156102af57600080fd5b5061028d6102be3660046124cd565b6109fe565b3480156102cf57600080fd5b5061028d6102de3660046124cd565b610a78565b3480156102ef57600080fd5b50610305600080516020612f6183398151915281565b60405190815260200161029a565b34801561031f57600080fd5b5061028d61032e366004612437565b610ac1565b34801561033f57600080fd5b5061028d61034e366004612437565b610bc9565b34801561035f57600080fd5b506103057fb18f47dcc713cc88c98ee8f273e72eeb31786b96acc9bb0504b62e0e092b028581565b34801561039357600080fd5b506103057fcb55a4ff4e57d8b70a6d90c4fb32866a1f659a165f1d9d8529f5fc97cadad3bb81565b3480156103c757600080fd5b506103db6103d6366004612565565b610cc1565b005b3480156103e957600080fd5b5060365461028d906001600160a01b031681565b34801561040957600080fd5b506103057fc51d77ababd4e8c9739d0e4f8e9b4996c85926cf22e47a1a3d8e8f138ef9b2c281565b34801561043d57600080fd5b5060025461028d906001600160a01b031681565b6103db61045f366004612598565b610d5a565b34801561047057600080fd5b50610305610d79565b34801561048557600080fd5b50610305600080516020612fa183398151915281565b3480156104a757600080fd5b5060015461028d906001600160a01b031681565b3480156104c757600080fd5b5061028d6104d6366004612437565b610d96565b3480156104e757600080fd5b506103056104f636600461265b565b610df1565b34801561050757600080fd5b506036546001600160a01b031661028d565b34801561052557600080fd5b50610305600080516020612f8183398151915281565b34801561054757600080fd5b5061028d610556366004612674565b610e3c565b34801561056757600080fd5b5061028d610576366004612716565b610e9e565b34801561058757600080fd5b5061028d610596366004612768565b610f2a565b3480156105a757600080fd5b506103057f5fd75318a22185c56d671bcc0d03b533a2d32884d4f19ec45e39dbc13666c40181565b3480156105db57600080fd5b506103057f899754d95ef7d990d237c5efa3cb30c0681aa503a4506a44d4f45400618d203f81565b34801561060f57600080fd5b506103057f92b3665efc99457ebd80a9331de1e6669556cd59542f5170a8c0cf141a3797dd81565b34801561064357600080fd5b506103db6106523660046127d1565b610fa7565b34801561066357600080fd5b5060375461028d906001600160a01b031681565b34801561068357600080fd5b5061028d6106923660046127d1565b611035565b3480156106a357600080fd5b506001610305565b3480156106b757600080fd5b5061028d6106c636600461281c565b6110b2565b3480156106d757600080fd5b50610305600080516020612fc183398151915281565b3480156106f957600080fd5b5061071e604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161029a91906128d7565b34801561073757600080fd5b506001546001600160a01b031661028d565b34801561075557600080fd5b506103057f22e8dcc44a0ea789b3e714eddf86e6afde2af1fa9b314ea411285159bd6c59ef81565b34801561078957600080fd5b506103db610798366004612565565b611113565b3480156107a957600080fd5b5061028d6111ac565b3480156107be57600080fd5b5061028d6107cd366004612716565b6111c5565b3480156107de57600080fd5b506103057f1246fd715e883c6fe463ca92b55f574af0bce2fb2366ab3574c6f8bdbfa0244f81565b34801561081257600080fd5b5061028d6108213660046128ea565b611231565b34801561083257600080fd5b506103db610841366004612565565b611290565b34801561085257600080fd5b5061028d610861366004612975565b611329565b34801561087257600080fd5b506002546001600160a01b031661028d565b34801561089057600080fd5b506103057fe1184adeb485c128bbddcc382768c2c1326594c0d5622b125c54e5e36aeb3e4c81565b3480156108c457600080fd5b5061028d6108d33660046128ea565b61138e565b3480156108e457600080fd5b506103db6108f3366004612565565b6113d9565b34801561090457600080fd5b506103db610913366004612a0e565b611472565b34801561092457600080fd5b5061028d6109333660046128ea565b6115b3565b34801561094457600080fd5b506103057faf96ce9190fe060687902e1d85f6c984df36addcea62deac92bf726c5b72175e81565b34801561097857600080fd5b5061028d610987366004612674565b611610565b60006109a4336000356001600160e01b03191661166f565b6109ac6117c3565b6109ea600080516020612fa1833981519152878787876040516020016109d59493929190612a93565b604051602081830303815290604052846117f9565b90505b6109f5611907565b95945050505050565b6000610a086117c3565b600480546001600160a01b03191633179055604051610a5e90600080516020612f8183398151915290610a479089908990899089908990602001612aba565b60408051601f1981840301815291905260006117f9565b600480546001600160a01b031916905590506109f5611907565b6000610a826117c3565b600480546001600160a01b03191633179055604051610a5e90600080516020612f6183398151915290610a479089908990899089908990602001612aba565b6000610ad9336000356001600160e01b03191661166f565b610ae16117c3565b6000859003610b035760405163974918c160e01b815260040160405180910390fd5b610b7886868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a018190048102820181019092528881529250889150879081908401838280828437600092019190915250879250611918915050565b905081816001600160a01b03167fb085ff794f342ed78acc7791d067e28a931e614b52476c0305795e1ff0a154bc60405160405180910390a3610bba81611a29565b156109ed576109ed8282611a9f565b6000610be1336000356001600160e01b03191661166f565b610be96117c3565b6000859003610c0b5760405163974918c160e01b815260040160405180910390fd5b610c8086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a018190048102820181019092528881529250889150879081908401838280828437600092019190915250879250611918915050565b905081816001600160a01b03167fb085ff794f342ed78acc7791d067e28a931e614b52476c0305795e1ff0a154bc60405160405180910390a36109f5611907565b610cd7336000356001600160e01b03191661166f565b6001600160a01b038116610cfe5760405163385d198b60e21b815260040160405180910390fd5b6037546040516001600160a01b038084169216907f847efc98ec1824503b43330b9ecaa4e53aba2277563bcfa73db74842fe45173390600090a3603780546001600160a01b0319166001600160a01b0392909216919091179055565b610d62611b32565b610d6b82611bc2565b610d758282611bdb565b5050565b6000610d83611c94565b50600080516020612f4183398151915290565b6000610dae336000356001600160e01b03191661166f565b610db66117c3565b6109ea7f5fd75318a22185c56d671bcc0d03b533a2d32884d4f19ec45e39dbc13666c401878787876040516020016109d59493929190612a93565b600081815260036020526040812080548190610e0c90612af7565b9050600003610e1e5750600092915050565b80604051610e2c9190612b31565b6040518091039020915050919050565b6000610e54336000356001600160e01b03191661166f565b610e5c6117c3565b610e89600080516020612fc18339815191528989898989896040516020016109d596959493929190612ba7565b9050610e93611907565b979650505050505050565b6000610eb6336000356001600160e01b03191661166f565b610ebe6117c3565b610f008585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508792506117f9915050565b9050610f0b81611a29565b15610f1a57610f1a8282611a9f565b610f22611907565b949350505050565b6000610f42336000356001600160e01b03191661166f565b610f4a6117c3565b600480546001600160a01b0319166001600160a01b038616179055604080516020601f8501819004810282018101909252838152610a5e9187919086908690819084018382808284376000920191909152508b92506117f9915050565b610fbd336000356001600160e01b03191661166f565b6000838152600360205260409020610fd6828483612c38565b50827f7cb1e35292a1f257e3116b1a5c9b2b2901410f7e08bb1594ab4d33de24823b2e821561101c57838360405161100f929190612cf8565b604051809103902061101f565b60005b60405190815260200160405180910390a2505050565b600061103f6117c3565b600480546001600160a01b03191633179055604080516020601f8501819004810282018101909252838152611091918691908690869081908401838280828437600092018290525092506117f9915050565b600480546001600160a01b031916905590506110ab611907565b9392505050565b60006110bc6117c3565b600480546001600160a01b031916331790556040516110f990600080516020612fa183398151915290610a47908890889088908890602001612a93565b600480546001600160a01b03191690559050610f22611907565b611129336000356001600160e01b03191661166f565b6001600160a01b0381166111505760405163385d198b60e21b815260040160405180910390fd5b6036546040516001600160a01b038084169216907fa1983b0b929b6bc0b362fb4a1dbfd5cc02906e7ddf83fe67de740650f5a4f94390600090a3603680546001600160a01b0319166001600160a01b0392909216919091179055565b60006111b6611cdd565b546001600160a01b0316919050565b60006111dd336000356001600160e01b03191661166f565b6111e56117c3565b6112278585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508792506117f9915050565b9050610f22611907565b6000611249336000356001600160e01b03191661166f565b6112516117c3565b61127c600080516020612f6183398151915288888888886040516020016109d5959493929190612aba565b9050611286611907565b9695505050505050565b6112a6336000356001600160e01b03191661166f565b6001600160a01b0381166112cd5760405163385d198b60e21b815260040160405180910390fd5b6001546040516001600160a01b038084169216907f241827194635b544bceb965d0e124305c0968051403ae83cb99d1ef86bd65f5890600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b60006113336117c3565b600480546001600160a01b0319163317905560405161137490600080516020612fc183398151915290610a47908a908a908a908a908a908a90602001612ba7565b600480546001600160a01b03191690559050611286611907565b60006113a6336000356001600160e01b03191661166f565b6113ae6117c3565b61127c600080516020612f8183398151915288888888886040516020016109d5959493929190612aba565b6113ef336000356001600160e01b03191661166f565b6001600160a01b0381166114165760405163385d198b60e21b815260040160405180910390fd5b6002546040516001600160a01b038084169216907f31dbecc021dab5ca391fd7d33f611f162c577acabcb9b1a7bf8e8155f571b08f90600090a3600280546001600160a01b0319166001600160a01b0392909216919091179055565b600061147c611d3f565b805490915060ff600160401b82041615906001600160401b03166000811580156114a35750825b90506000826001600160401b031660011480156114bf5750303b155b9050811580156114cd575080155b156114eb5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561151557845460ff60401b1916600160401b1785555b6001600160a01b03881661153c5760405163385d198b60e21b815260040160405180910390fd5b611547898888611d6a565b603680546001600160a01b0319166001600160a01b038a1617905583156115a857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60006115cb336000356001600160e01b03191661166f565b6115d36117c3565b61127c7f22e8dcc44a0ea789b3e714eddf86e6afde2af1fa9b314ea411285159bd6c59ef88888888886040516020016109d5959493929190612aba565b6000611628336000356001600160e01b03191661166f565b6116306117c3565b610e897f899754d95ef7d990d237c5efa3cb30c0681aa503a4506a44d4f45400618d203f8989898989896040516020016109d596959493929190612ba7565b6000611679611cdd565b80549091506001600160a01b0316806116b1576000604051638944034760e01b81526004016116a891906124aa565b60405180910390fd5b60405163b700961360e01b81526001600160a01b0385811660048301523060248301526001600160e01b031985166044830152600091829182919085169063b700961390606401606060405180830381865afa158015611715573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117399190612d18565b925092509250826117ba5780156117635760405163cc9855ad60e01b815260040160405180910390fd5b63ffffffff82161561179f5760405163a426878960e01b81526001600160a01b038816600482015263ffffffff831660248201526044016116a8565b86604051632ecd3d0360e21b81526004016116a891906124aa565b50505050505050565b60006117cd611df6565b8054909150600119016117f357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000838152600360205260408120805482919061181590612af7565b80601f016020809104026020016040519081016040528092919081815260200182805461184190612af7565b801561188e5780601f106118635761010080835404028352916020019161188e565b820191906000526020600020905b81548152906001019060200180831161187157829003601f168201915b5050505050905080516000036118ba57604051636b714b5760e01b8152600481018690526024016116a8565b6118c5818585611918565b915082826001600160a01b0316867f616141b2c322f0e10fa36f75caad9f55f2a5f42053b5b1b5f93482c2b6c0b82b60405160405180910390a4509392505050565b6000611911611df6565b6001905550565b600080600080815461192990612d7c565b9182905550905061194460008261193f88611e1a565b611f18565b9150600061195184611faf565b90506000836001600160a01b03168683604051602401611972929190612d95565b60408051601f198184030181529181526020820180516001600160e01b03166301d3f35b60e41b179052516119a79190612dfc565b6000604051808303816000865af19150503d80600081146119e4576040519150601f19603f3d011682016040523d82523d6000602084013e6119e9565b606091505b5050905080611a1f573d8015611a0557604051816000823e8181fd5b5060405163063cc07360e41b815260040160405180910390fd5b5050509392505050565b6000816001600160a01b03166335abee1a6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611a85575060408051601f3d908101601f19168201909252611a8291810190612e18565b60015b611a9157506000919050565b50600192915050565b919050565b6037546001600160a01b0316611ac8576040516314d4388560e31b815260040160405180910390fd5b6037546040516347ecc93960e01b8152600481018490526001600160a01b038381166024830152909116906347ecc93990604401600060405180830381600087803b158015611b1657600080fd5b505af1158015611b2a573d6000803e3d6000fd5b505050505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480611ba257507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611b966120f0565b6001600160a01b031614155b15611bc05760405163703e46dd60e11b815260040160405180910390fd5b565b611bd8336000356001600160e01b03191661166f565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611c35575060408051601f3d908101601f19168201909252611c3291810190612e35565b60015b611c545781604051634c9c8ce360e01b81526004016116a891906124aa565b600080516020612f418339815191528114611c8557604051632a87526960e21b8152600481018290526024016116a8565b611c8f8383612106565b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611bc05760405163703e46dd60e11b815260040160405180910390fd5b60008060ff19611d0e60017f2d9a26166e6ed6927c3e14d89c4d572a357f984344fcaed3c38c3ff5cdb83f35612e4e565b604051602001611d2091815260200190565b60408051601f1981840301815291905280516020909101201692915050565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b611d7261215c565b6001600160a01b0383161580611d8f57506001600160a01b038216155b15611dad5760405163385d198b60e21b815260040160405180910390fd5b611db5612181565b611dbd612189565b600180546001600160a01b038086166001600160a01b0319928316179092556002805492851692909116919091179055611c8f81612199565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b606061ffff801682511115611e47578151604051634fbdd44360e11b81526004016116a891815260200190565b60ff825111611ebf57604051806040016040528060128152602001710608060405234801561001057600080fd5b560741b815250606060f81b835160f81b6a8061001f6000396000f3fe60a81b85604051602001611ea9959493929190612e61565b6040516020818303038152906040529050919050565b604051806040016040528060128152602001710608060405234801561001057600080fd5b560741b815250606160f81b835160f01b6a806100206000396000f3fe60a81b85604051602001611ea9959493929190612ebe565b600083471015611f445760405163cf47918160e01b8152476004820152602481018590526044016116a8565b8151600003611f6657604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615611f88576040513d6000823e3d81fd5b6001600160a01b0381166110ab5760405163b06ebf3d60e01b815260040160405180910390fd5b6040805160c0808201835260008083526020808401829052838501829052606084018290526080840182905260a084019190915260048054855193840186526001546001600160a01b0390811680865260365482168686015287516317170e2760e31b815288519798929093169691860194909363b8b87138938082019391908290030181865afa158015612048573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061206c9190612f23565b6001600160a01b0316815260200160006001600160a01b0316836001600160a01b0316036120a5576002546001600160a01b03166120a7565b825b6001600160a01b0316815260200184815260200160006001600160a01b0316836001600160a01b0316036120db57336120de565b60005b6001600160a01b031690529392505050565b6000600080516020612f418339815191526111b6565b61210f826121da565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561215457611c8f8282612236565b610d756122a3565b6121646122c2565b611bc057604051631afcd79f60e31b815260040160405180910390fd5b611bc061215c565b61219161215c565b611bc06122dc565b60006121a3611cdd565b80549091506001600160a01b0316156121d15781604051638944034760e01b81526004016116a891906124aa565b610d75826122e4565b806001600160a01b03163b6000036122075780604051634c9c8ce360e01b81526004016116a891906124aa565b600080516020612f4183398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516122539190612dfc565b600060405180830381855af49150503d806000811461228e576040519150601f19603f3d011682016040523d82523d6000602084013e612293565b606091505b50915091506109f5858383612374565b3415611bc05760405163b398979f60e01b815260040160405180910390fd5b60006122cc611d3f565b54600160401b900460ff16919050565b61190761215c565b6001600160a01b03811661230d5780604051638944034760e01b81526004016116a891906124aa565b6000612317611cdd565b80546040519192506001600160a01b03808516929116907fa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b7638998019890600090a380546001600160a01b0319166001600160a01b0392909216919091179055565b60608261238957612384826123c7565b6110ab565b81511580156123a057506001600160a01b0384163b155b156123c05783604051639996b31560e01b81526004016116a891906124aa565b50806110ab565b8051156123d657805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b60008083601f84011261240157600080fd5b5081356001600160401b0381111561241857600080fd5b60208301915083602082850101111561243057600080fd5b9250929050565b60008060008060006060868803121561244f57600080fd5b85356001600160401b038082111561246657600080fd5b61247289838a016123ef565b9097509550602088013591508082111561248b57600080fd5b50612498888289016123ef565b96999598509660400135949350505050565b6001600160a01b0391909116815260200190565b60ff81168114611bd857600080fd5b6000806000806000606086880312156124e557600080fd5b85356001600160401b03808211156124fc57600080fd5b61250889838a016123ef565b9097509550602088013591508082111561252157600080fd5b5061252e888289016123ef565b9094509250506040860135612542816124be565b809150509295509295909350565b6001600160a01b0381168114611bd857600080fd5b60006020828403121561257757600080fd5b81356110ab81612550565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156125ab57600080fd5b82356125b681612550565b915060208301356001600160401b03808211156125d257600080fd5b818501915085601f8301126125e657600080fd5b8135818111156125f8576125f8612582565b604051601f8201601f19908116603f0116810190838211818310171561262057612620612582565b8160405282815288602084870101111561263957600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60006020828403121561266d57600080fd5b5035919050565b60008060008060008060006080888a03121561268f57600080fd5b87356001600160401b03808211156126a657600080fd5b6126b28b838c016123ef565b909950975060208a01359150808211156126cb57600080fd5b6126d78b838c016123ef565b909750955060408a01359150808211156126f057600080fd5b506126fd8a828b016123ef565b989b979a50959894979596606090950135949350505050565b6000806000806060858703121561272c57600080fd5b8435935060208501356001600160401b0381111561274957600080fd5b612755878288016123ef565b9598909750949560400135949350505050565b60008060008060006080868803121561278057600080fd5b8535945060208601359350604086013561279981612550565b925060608601356001600160401b038111156127b457600080fd5b6127c0888289016123ef565b969995985093965092949392505050565b6000806000604084860312156127e657600080fd5b8335925060208401356001600160401b0381111561280357600080fd5b61280f868287016123ef565b9497909650939450505050565b6000806000806040858703121561283257600080fd5b84356001600160401b038082111561284957600080fd5b612855888389016123ef565b9096509450602087013591508082111561286e57600080fd5b5061287b878288016123ef565b95989497509550505050565b60005b838110156128a257818101518382015260200161288a565b50506000910152565b600081518084526128c3816020860160208601612887565b601f01601f19169290920160200192915050565b6020815260006110ab60208301846128ab565b6000806000806000806080878903121561290357600080fd5b86356001600160401b038082111561291a57600080fd5b6129268a838b016123ef565b9098509650602089013591508082111561293f57600080fd5b5061294c89828a016123ef565b9095509350506040870135612960816124be565b80925050606087013590509295509295509295565b6000806000806000806060878903121561298e57600080fd5b86356001600160401b03808211156129a557600080fd5b6129b18a838b016123ef565b909850965060208901359150808211156129ca57600080fd5b6129d68a838b016123ef565b909650945060408901359150808211156129ef57600080fd5b506129fc89828a016123ef565b979a9699509497509295939492505050565b60008060008060808587031215612a2457600080fd5b8435612a2f81612550565b93506020850135612a3f81612550565b92506040850135612a4f81612550565b91506060850135612a5f81612550565b939692955090935050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000612aa7604083018688612a6a565b8281036020840152610e93818587612a6a565b606081526000612ace606083018789612a6a565b8281036020840152612ae1818688612a6a565b91505060ff831660408301529695505050505050565b600181811c90821680612b0b57607f821691505b602082108103612b2b57634e487b7160e01b600052602260045260246000fd5b50919050565b6000808354612b3f81612af7565b60018281168015612b575760018114612b6c57612b9b565b60ff1984168752821515830287019450612b9b565b8760005260208060002060005b85811015612b925781548a820152908401908201612b79565b50505082870194505b50929695505050505050565b606081526000612bbb60608301888a612a6a565b8281036020840152612bce818789612a6a565b90508281036040840152612be3818587612a6a565b9998505050505050505050565b601f821115611c8f576000816000526020600020601f850160051c81016020861015612c195750805b601f850160051c820191505b81811015611b2a57828155600101612c25565b6001600160401b03831115612c4f57612c4f612582565b612c6383612c5d8354612af7565b83612bf0565b6000601f841160018114612c975760008515612c7f5750838201355b600019600387901b1c1916600186901b178355612cf1565b600083815260209020601f19861690835b82811015612cc85786850135825560209485019460019092019101612ca8565b5086821015612ce55760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b8183823760009101908152919050565b80518015158114611a9a57600080fd5b600080600060608486031215612d2d57600080fd5b612d3684612d08565b9250602084015163ffffffff81168114612d4f57600080fd5b9150612d5d60408501612d08565b90509250925092565b634e487b7160e01b600052601160045260246000fd5b600060018201612d8e57612d8e612d66565b5060010190565b60e081526000612da860e08301856128ab565b905060018060a01b03808451166020840152806020850151166040840152806040850151166060840152806060850151166080840152608084015160a08401528060a08501511660c0840152509392505050565b60008251612e0e818460208701612887565b9190910192915050565b600060208284031215612e2a57600080fd5b81516110ab816124be565b600060208284031215612e4757600080fd5b5051919050565b81810381811115611d6457611d64612d66565b60008651612e73818460208b01612887565b6001600160f81b0319878116918401918252861660018201526001600160a81b0319851660028201528351612eaf81600d840160208801612887565b01600d01979650505050505050565b60008651612ed0818460208b01612887565b6001600160f81b031987169083019081526001600160f01b0319861660018201526001600160a81b0319851660038201528351612f1481600e840160208801612887565b01600e01979650505050505050565b600060208284031215612f3557600080fd5b81516110ab8161255056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcb99feb9478d78ab32407a7d9d661b47b23e611a2b9c80435e669608baa5096767e8359ad5da7b3984f0033ccc1a79b3e1c2db4fb05b6ad2edf4355fdd64c221d4d51fe40120c0f73a1511e2664e915719b5731e1b39ccc60a07bfa11a21a0c777d694388a0a19d180d06562499b666a88a71bb4b6df947f4ed9746b2f23bc5bfa26469706673582212202bf04bd7bf7ea9155c3fd9935848b7a97fd5e5bd09b32b722faff10ee425efee64736f6c63430008180033",
}

// RNContractFactoryV1 is an auto generated Go binding around an Ethereum contract.
type RNContractFactoryV1 struct {
	abi abi.ABI
}

// NewRNContractFactoryV1 creates a new instance of RNContractFactoryV1.
func NewRNContractFactoryV1() *RNContractFactoryV1 {
	parsed, err := RNContractFactoryV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RNContractFactoryV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RNContractFactoryV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackRAYLSENYGMAKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d7ffb32.
//
// Solidity: function RAYLS_ENYGMA_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSENYGMAKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ENYGMA_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSENYGMAKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6d7ffb32.
//
// Solidity: function RAYLS_ENYGMA_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSENYGMAKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ENYGMA_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSENYGMATESTKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c7fe28b.
//
// Solidity: function RAYLS_ENYGMA_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSENYGMATESTKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ENYGMA_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSENYGMATESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1c7fe28b.
//
// Solidity: function RAYLS_ENYGMA_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSENYGMATESTKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ENYGMA_TEST_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC1155DVPKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92c04d9b.
//
// Solidity: function RAYLS_ERC1155_DVP_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC1155DVPKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC1155_DVP_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155DVPKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x92c04d9b.
//
// Solidity: function RAYLS_ERC1155_DVP_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC1155DVPKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC1155_DVP_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC1155DVPTESTKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29b0d701.
//
// Solidity: function RAYLS_ERC1155_DVP_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC1155DVPTESTKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC1155_DVP_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155DVPTESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x29b0d701.
//
// Solidity: function RAYLS_ERC1155_DVP_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC1155DVPTESTKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC1155_DVP_TEST_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC1155KEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5dd56210.
//
// Solidity: function RAYLS_ERC1155_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC1155KEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC1155_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155KEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5dd56210.
//
// Solidity: function RAYLS_ERC1155_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC1155KEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC1155_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC1155TESTKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x403be3de.
//
// Solidity: function RAYLS_ERC1155_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC1155TESTKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC1155_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155TESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x403be3de.
//
// Solidity: function RAYLS_ERC1155_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC1155TESTKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC1155_TEST_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC20KEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1464230b.
//
// Solidity: function RAYLS_ERC20_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC20KEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC20_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC20KEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1464230b.
//
// Solidity: function RAYLS_ERC20_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC20KEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC20_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC20TESTKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xee5ce29e.
//
// Solidity: function RAYLS_ERC20_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC20TESTKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC20_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC20TESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xee5ce29e.
//
// Solidity: function RAYLS_ERC20_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC20TESTKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC20_TEST_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC721DVPKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94f08898.
//
// Solidity: function RAYLS_ERC721_DVP_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC721DVPKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC721_DVP_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721DVPKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94f08898.
//
// Solidity: function RAYLS_ERC721_DVP_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC721DVPKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC721_DVP_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC721DVPTESTKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95f50a62.
//
// Solidity: function RAYLS_ERC721_DVP_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC721DVPTESTKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC721_DVP_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721DVPTESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95f50a62.
//
// Solidity: function RAYLS_ERC721_DVP_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC721DVPTESTKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC721_DVP_TEST_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC721KEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad10c149.
//
// Solidity: function RAYLS_ERC721_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC721KEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC721_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721KEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad10c149.
//
// Solidity: function RAYLS_ERC721_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC721KEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC721_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSERC721TESTKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcab6982.
//
// Solidity: function RAYLS_ERC721_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSERC721TESTKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_ERC721_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721TESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfcab6982.
//
// Solidity: function RAYLS_ERC721_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSERC721TESTKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_ERC721_TEST_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSSTABLECOINKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb3872e4e.
//
// Solidity: function RAYLS_STABLECOIN_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSSTABLECOINKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_STABLECOIN_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSSTABLECOINKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb3872e4e.
//
// Solidity: function RAYLS_STABLECOIN_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSSTABLECOINKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_STABLECOIN_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRAYLSSTABLECOINTESTKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3fe9c16.
//
// Solidity: function RAYLS_STABLECOIN_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRAYLSSTABLECOINTESTKEY() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("RAYLS_STABLECOIN_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSSTABLECOINTESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3fe9c16.
//
// Solidity: function RAYLS_STABLECOIN_TEST_KEY() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRAYLSSTABLECOINTESTKEY(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("RAYLS_STABLECOIN_TEST_KEY", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNContractFactoryV1 *RNContractFactoryV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := rNContractFactoryV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAuthority is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackAuthority() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("authority", data)
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
func (rNContractFactoryV1 *RNContractFactoryV1) PackContractVersion() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("contractVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackContractVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackContractVersion(data []byte) (*big.Int, error) {
	out, err := rNContractFactoryV1.abi.Unpack("contractVersion", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDeploy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19e011ad.
//
// Solidity: function deploy(bytes bytecode, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeploy(bytecode []byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deploy", bytecode, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeploy is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x19e011ad.
//
// Solidity: function deploy(bytes bytecode, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeploy(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deploy", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployEnygma is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeff52a99.
//
// Solidity: function deployEnygma(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployEnygma(name string, symbol string, decimals uint8, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployEnygma", name, symbol, decimals, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployEnygma is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeff52a99.
//
// Solidity: function deployEnygma(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployEnygma(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployEnygma", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployEnygmaAsUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02b6bcf2.
//
// Solidity: function deployEnygmaAsUser(string name, string symbol, uint8 decimals) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployEnygmaAsUser(name string, symbol string, decimals uint8) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployEnygmaAsUser", name, symbol, decimals)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployEnygmaAsUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02b6bcf2.
//
// Solidity: function deployEnygmaAsUser(string name, string symbol, uint8 decimals) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployEnygmaAsUser(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployEnygmaAsUser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00866e72.
//
// Solidity: function deployErc1155(string uri, string name, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc1155(uri string, name string, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc1155", uri, name, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc1155 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00866e72.
//
// Solidity: function deployErc1155(string uri, string name, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc1155(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc1155", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc1155AsUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8126261.
//
// Solidity: function deployErc1155AsUser(string uri, string name) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc1155AsUser(uri string, name string) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc1155AsUser", uri, name)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc1155AsUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8126261.
//
// Solidity: function deployErc1155AsUser(string uri, string name) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc1155AsUser(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc1155AsUser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc1155Dvp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x614bcfd3.
//
// Solidity: function deployErc1155Dvp(string uri, string name, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc1155Dvp(uri string, name string, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc1155Dvp", uri, name, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc1155Dvp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x614bcfd3.
//
// Solidity: function deployErc1155Dvp(string uri, string name, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc1155Dvp(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc1155Dvp", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd64e6e86.
//
// Solidity: function deployErc20(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc20(name string, symbol string, decimals uint8, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc20", name, symbol, decimals, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc20 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd64e6e86.
//
// Solidity: function deployErc20(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc20(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc20", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc20AsUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0cd503ea.
//
// Solidity: function deployErc20AsUser(string name, string symbol, uint8 decimals) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc20AsUser(name string, symbol string, decimals uint8) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc20AsUser", name, symbol, decimals)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc20AsUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0cd503ea.
//
// Solidity: function deployErc20AsUser(string name, string symbol, uint8 decimals) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc20AsUser(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc20AsUser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e9bb21e.
//
// Solidity: function deployErc721(string uri, string name, string symbol, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc721(uri string, name string, symbol string, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc721", uri, name, symbol, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc721 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e9bb21e.
//
// Solidity: function deployErc721(string uri, string name, string symbol, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc721(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc721", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc721AsUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc905d08.
//
// Solidity: function deployErc721AsUser(string uri, string name, string symbol) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc721AsUser(uri string, name string, symbol string) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc721AsUser", uri, name, symbol)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc721AsUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdc905d08.
//
// Solidity: function deployErc721AsUser(string uri, string name, string symbol) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc721AsUser(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc721AsUser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployErc721Dvp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe75427d.
//
// Solidity: function deployErc721Dvp(string uri, string name, string symbol, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployErc721Dvp(uri string, name string, symbol string, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployErc721Dvp", uri, name, symbol, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc721Dvp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfe75427d.
//
// Solidity: function deployErc721Dvp(string uri, string name, string symbol, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployErc721Dvp(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployErc721Dvp", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployExternal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16381b15.
//
// Solidity: function deployExternal(bytes bytecode, bytes userArgs, bytes32 resourceId) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployExternal(bytecode []byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployExternal", bytecode, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployExternal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x16381b15.
//
// Solidity: function deployExternal(bytes bytecode, bytes userArgs, bytes32 resourceId) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployExternal(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployExternal", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployFromTeleport is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f22b4a9.
//
// Solidity: function deployFromTeleport(bytes32 resourceId, bytes32 factoryKey, address ownerEOA, bytes userArgs) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployFromTeleport(resourceId [32]byte, factoryKey [32]byte, ownerEOA common.Address, userArgs []byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployFromTeleport", resourceId, factoryKey, ownerEOA, userArgs)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployFromTeleport is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f22b4a9.
//
// Solidity: function deployFromTeleport(bytes32 resourceId, bytes32 factoryKey, address ownerEOA, bytes userArgs) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployFromTeleport(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployFromTeleport", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcd39c4a8.
//
// Solidity: function deployRegistered(bytes32 key, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployRegistered(key [32]byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployRegistered", key, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployRegistered is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcd39c4a8.
//
// Solidity: function deployRegistered(bytes32 key, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployRegistered(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployRegistered", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployRegisteredAsUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ee1d00a.
//
// Solidity: function deployRegisteredAsUser(bytes32 key, bytes userArgs) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployRegisteredAsUser(key [32]byte, userArgs []byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployRegisteredAsUser", key, userArgs)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployRegisteredAsUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ee1d00a.
//
// Solidity: function deployRegisteredAsUser(bytes32 key, bytes userArgs) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployRegisteredAsUser(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployRegisteredAsUser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployRegisteredExternal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f5e71fe.
//
// Solidity: function deployRegisteredExternal(bytes32 key, bytes userArgs, bytes32 resourceId) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployRegisteredExternal(key [32]byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployRegisteredExternal", key, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployRegisteredExternal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f5e71fe.
//
// Solidity: function deployRegisteredExternal(bytes32 key, bytes userArgs, bytes32 resourceId) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployRegisteredExternal(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployRegisteredExternal", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployStableCoin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfbc19051.
//
// Solidity: function deployStableCoin(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployStableCoin(name string, symbol string, decimals uint8, resourceId [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployStableCoin", name, symbol, decimals, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployStableCoin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfbc19051.
//
// Solidity: function deployStableCoin(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployStableCoin(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployStableCoin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployStableCoinAsUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5734c8b.
//
// Solidity: function deployStableCoinAsUser(string name, string symbol, uint8 decimals) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) PackDeployStableCoinAsUser(name string, symbol string, decimals uint8) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("deployStableCoinAsUser", name, symbol, decimals)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployStableCoinAsUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe5734c8b.
//
// Solidity: function deployStableCoinAsUser(string name, string symbol, uint8 decimals) returns(address deployed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackDeployStableCoinAsUser(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("deployStableCoinAsUser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackEndpoint() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("endpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackEndpoint(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("endpoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackFactoryOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4273601c.
//
// Solidity: function factoryOwner() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackFactoryOwner() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("factoryOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFactoryOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4273601c.
//
// Solidity: function factoryOwner() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFactoryOwner(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("factoryOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetBytecodeHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6670b8c5.
//
// Solidity: function getBytecodeHash(bytes32 key) view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackGetBytecodeHash(key [32]byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("getBytecodeHash", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetBytecodeHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6670b8c5.
//
// Solidity: function getBytecodeHash(bytes32 key) view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackGetBytecodeHash(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("getBytecodeHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaed8e967.
//
// Solidity: function getEndpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackGetEndpoint() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("getEndpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaed8e967.
//
// Solidity: function getEndpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackGetEndpoint(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("getEndpoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetFactoryOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe09bda6b.
//
// Solidity: function getFactoryOwner() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackGetFactoryOwner() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("getFactoryOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFactoryOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe09bda6b.
//
// Solidity: function getFactoryOwner() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackGetFactoryOwner(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("getFactoryOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetRaylsNodeEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6bc5d477.
//
// Solidity: function getRaylsNodeEndpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackGetRaylsNodeEndpoint() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("getRaylsNodeEndpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRaylsNodeEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6bc5d477.
//
// Solidity: function getRaylsNodeEndpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackGetRaylsNodeEndpoint(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("getRaylsNodeEndpoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8c8765e.
//
// Solidity: function initialize(address _endpoint, address _raylsNodeEndpoint, address _owner, address authority_) returns()
func (rNContractFactoryV1 *RNContractFactoryV1) PackInitialize(endpoint common.Address, raylsNodeEndpoint common.Address, owner common.Address, authority common.Address) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("initialize", endpoint, raylsNodeEndpoint, owner, authority)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) PackProxiableUUID() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := rNContractFactoryV1.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRaylsNodeEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39e297b2.
//
// Solidity: function raylsNodeEndpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackRaylsNodeEndpoint() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("raylsNodeEndpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRaylsNodeEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x39e297b2.
//
// Solidity: function raylsNodeEndpoint() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRaylsNodeEndpoint(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("raylsNodeEndpoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSetBytecode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c363c31.
//
// Solidity: function setBytecode(bytes32 key, bytes bytecode) returns()
func (rNContractFactoryV1 *RNContractFactoryV1) PackSetBytecode(key [32]byte, bytecode []byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("setBytecode", key, bytecode)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdbbb4155.
//
// Solidity: function setEndpoint(address newEndpoint) returns()
func (rNContractFactoryV1 *RNContractFactoryV1) PackSetEndpoint(newEndpoint common.Address) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("setEndpoint", newEndpoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFactoryOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3cc660c.
//
// Solidity: function setFactoryOwner(address newOwner) returns()
func (rNContractFactoryV1 *RNContractFactoryV1) PackSetFactoryOwner(newOwner common.Address) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("setFactoryOwner", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRaylsNodeEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd7e7d2f.
//
// Solidity: function setRaylsNodeEndpoint(address newEndpoint) returns()
func (rNContractFactoryV1 *RNContractFactoryV1) PackSetRaylsNodeEndpoint(newEndpoint common.Address) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("setRaylsNodeEndpoint", newEndpoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTokenRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35a5af92.
//
// Solidity: function setTokenRegistry(address _tokenRegistry) returns()
func (rNContractFactoryV1 *RNContractFactoryV1) PackSetTokenRegistry(tokenRegistry common.Address) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("setTokenRegistry", tokenRegistry)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTokenRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) PackTokenRegistry() []byte {
	enc, err := rNContractFactoryV1.abi.Pack("tokenRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackTokenRegistry(data []byte) (common.Address, error) {
	out, err := rNContractFactoryV1.abi.Unpack("tokenRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (rNContractFactoryV1 *RNContractFactoryV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := rNContractFactoryV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// RNContractFactoryV1AuthorityUpdated represents a AuthorityUpdated event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1AuthorityUpdated) ContractEventName() string {
	return RNContractFactoryV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*RNContractFactoryV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1BytecodeSet represents a BytecodeSet event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1BytecodeSet struct {
	Key          [32]byte
	BytecodeHash [32]byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1BytecodeSetEventName = "BytecodeSet"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1BytecodeSet) ContractEventName() string {
	return RNContractFactoryV1BytecodeSetEventName
}

// UnpackBytecodeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BytecodeSet(bytes32 indexed key, bytes32 bytecodeHash)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackBytecodeSetEvent(log *types.Log) (*RNContractFactoryV1BytecodeSet, error) {
	event := "BytecodeSet"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1BytecodeSet)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1ContractDeployed represents a ContractDeployed event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1ContractDeployed struct {
	DeployedAddress common.Address
	ResourceId      [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1ContractDeployedEventName = "ContractDeployed"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1ContractDeployed) ContractEventName() string {
	return RNContractFactoryV1ContractDeployedEventName
}

// UnpackContractDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ContractDeployed(address indexed deployedAddress, bytes32 indexed resourceId)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackContractDeployedEvent(log *types.Log) (*RNContractFactoryV1ContractDeployed, error) {
	event := "ContractDeployed"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1ContractDeployed)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1EndpointUpdated represents a EndpointUpdated event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1EndpointUpdated struct {
	OldEndpoint common.Address
	NewEndpoint common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1EndpointUpdatedEventName = "EndpointUpdated"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1EndpointUpdated) ContractEventName() string {
	return RNContractFactoryV1EndpointUpdatedEventName
}

// UnpackEndpointUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EndpointUpdated(address indexed oldEndpoint, address indexed newEndpoint)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackEndpointUpdatedEvent(log *types.Log) (*RNContractFactoryV1EndpointUpdated, error) {
	event := "EndpointUpdated"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1EndpointUpdated)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1FactoryOwnerUpdated represents a FactoryOwnerUpdated event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FactoryOwnerUpdated struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1FactoryOwnerUpdatedEventName = "FactoryOwnerUpdated"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1FactoryOwnerUpdated) ContractEventName() string {
	return RNContractFactoryV1FactoryOwnerUpdatedEventName
}

// UnpackFactoryOwnerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FactoryOwnerUpdated(address indexed oldOwner, address indexed newOwner)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFactoryOwnerUpdatedEvent(log *types.Log) (*RNContractFactoryV1FactoryOwnerUpdated, error) {
	event := "FactoryOwnerUpdated"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1FactoryOwnerUpdated)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1Initialized represents a Initialized event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1Initialized) ContractEventName() string {
	return RNContractFactoryV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackInitializedEvent(log *types.Log) (*RNContractFactoryV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1Initialized)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1RaylsNodeEndpointUpdated represents a RaylsNodeEndpointUpdated event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1RaylsNodeEndpointUpdated struct {
	OldEndpoint common.Address
	NewEndpoint common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1RaylsNodeEndpointUpdatedEventName = "RaylsNodeEndpointUpdated"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1RaylsNodeEndpointUpdated) ContractEventName() string {
	return RNContractFactoryV1RaylsNodeEndpointUpdatedEventName
}

// UnpackRaylsNodeEndpointUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RaylsNodeEndpointUpdated(address indexed oldEndpoint, address indexed newEndpoint)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRaylsNodeEndpointUpdatedEvent(log *types.Log) (*RNContractFactoryV1RaylsNodeEndpointUpdated, error) {
	event := "RaylsNodeEndpointUpdated"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1RaylsNodeEndpointUpdated)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1RegisteredContractDeployed represents a RegisteredContractDeployed event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1RegisteredContractDeployed struct {
	Key             [32]byte
	DeployedAddress common.Address
	ResourceId      [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1RegisteredContractDeployedEventName = "RegisteredContractDeployed"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1RegisteredContractDeployed) ContractEventName() string {
	return RNContractFactoryV1RegisteredContractDeployedEventName
}

// UnpackRegisteredContractDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RegisteredContractDeployed(bytes32 indexed key, address indexed deployedAddress, bytes32 indexed resourceId)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRegisteredContractDeployedEvent(log *types.Log) (*RNContractFactoryV1RegisteredContractDeployed, error) {
	event := "RegisteredContractDeployed"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1RegisteredContractDeployed)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1TokenRegistrySet represents a TokenRegistrySet event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1TokenRegistrySet struct {
	OldRegistry common.Address
	NewRegistry common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1TokenRegistrySetEventName = "TokenRegistrySet"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1TokenRegistrySet) ContractEventName() string {
	return RNContractFactoryV1TokenRegistrySetEventName
}

// UnpackTokenRegistrySetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenRegistrySet(address indexed oldRegistry, address indexed newRegistry)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackTokenRegistrySetEvent(log *types.Log) (*RNContractFactoryV1TokenRegistrySet, error) {
	event := "TokenRegistrySet"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1TokenRegistrySet)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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

// RNContractFactoryV1Upgraded represents a Upgraded event raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNContractFactoryV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (RNContractFactoryV1Upgraded) ContractEventName() string {
	return RNContractFactoryV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackUpgradedEvent(log *types.Log) (*RNContractFactoryV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != rNContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNContractFactoryV1Upgraded)
	if len(log.Data) > 0 {
		if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNContractFactoryV1.abi.Events[event].Inputs {
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
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["FactoryV1BytecodeNotRegistered"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackFactoryV1BytecodeNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["FactoryV1EmptyBytecode"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackFactoryV1EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["FactoryV1InitializationFailed"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackFactoryV1InitializationFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["FactoryV1TokenRegistryNotSet"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackFactoryV1TokenRegistryNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["FactoryV1ZeroAddress"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackFactoryV1ZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["InitCodeStubRuntimeTooLarge"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackInitCodeStubRuntimeTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNContractFactoryV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return rNContractFactoryV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// RNContractFactoryV1AddressEmptyCode represents a AddressEmptyCode error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func RNContractFactoryV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackAddressEmptyCodeError(raw []byte) (*RNContractFactoryV1AddressEmptyCode, error) {
	out := new(RNContractFactoryV1AddressEmptyCode)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1Create2EmptyBytecode represents a Create2EmptyBytecode error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1Create2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func RNContractFactoryV1Create2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackCreate2EmptyBytecodeError(raw []byte) (*RNContractFactoryV1Create2EmptyBytecode, error) {
	out := new(RNContractFactoryV1Create2EmptyBytecode)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func RNContractFactoryV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackERC1967InvalidImplementationError(raw []byte) (*RNContractFactoryV1ERC1967InvalidImplementation, error) {
	out := new(RNContractFactoryV1ERC1967InvalidImplementation)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func RNContractFactoryV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackERC1967NonPayableError(raw []byte) (*RNContractFactoryV1ERC1967NonPayable, error) {
	out := new(RNContractFactoryV1ERC1967NonPayable)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1FactoryV1BytecodeNotRegistered represents a FactoryV1__BytecodeNotRegistered error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FactoryV1BytecodeNotRegistered struct {
	Key [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__BytecodeNotRegistered(bytes32 key)
func RNContractFactoryV1FactoryV1BytecodeNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x6b714b57217f55da5c41092b90d46a27701a6e34bee2c6605a39b6a1cba654f7")
}

// UnpackFactoryV1BytecodeNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__BytecodeNotRegistered(bytes32 key)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFactoryV1BytecodeNotRegisteredError(raw []byte) (*RNContractFactoryV1FactoryV1BytecodeNotRegistered, error) {
	out := new(RNContractFactoryV1FactoryV1BytecodeNotRegistered)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1BytecodeNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1FactoryV1EmptyBytecode represents a FactoryV1__EmptyBytecode error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FactoryV1EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__EmptyBytecode()
func RNContractFactoryV1FactoryV1EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x974918c11d78344eb3e9a3a905bdaa5941bfbd85677f790418e428875f34cc09")
}

// UnpackFactoryV1EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__EmptyBytecode()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFactoryV1EmptyBytecodeError(raw []byte) (*RNContractFactoryV1FactoryV1EmptyBytecode, error) {
	out := new(RNContractFactoryV1FactoryV1EmptyBytecode)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1FactoryV1InitializationFailed represents a FactoryV1__InitializationFailed error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FactoryV1InitializationFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__InitializationFailed()
func RNContractFactoryV1FactoryV1InitializationFailedErrorID() common.Hash {
	return common.HexToHash("0x63cc0730a6f99cbea6ef1a62e46650dcbc73f7146679065ed2f148d9537f1e27")
}

// UnpackFactoryV1InitializationFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__InitializationFailed()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFactoryV1InitializationFailedError(raw []byte) (*RNContractFactoryV1FactoryV1InitializationFailed, error) {
	out := new(RNContractFactoryV1FactoryV1InitializationFailed)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1InitializationFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1FactoryV1TokenRegistryNotSet represents a FactoryV1__TokenRegistryNotSet error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FactoryV1TokenRegistryNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__TokenRegistryNotSet()
func RNContractFactoryV1FactoryV1TokenRegistryNotSetErrorID() common.Hash {
	return common.HexToHash("0xa6a1c428a1fda164d946b7ebd17a4a4336420a4199c9bc3a0729b49dd381e9a3")
}

// UnpackFactoryV1TokenRegistryNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__TokenRegistryNotSet()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFactoryV1TokenRegistryNotSetError(raw []byte) (*RNContractFactoryV1FactoryV1TokenRegistryNotSet, error) {
	out := new(RNContractFactoryV1FactoryV1TokenRegistryNotSet)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1TokenRegistryNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1FactoryV1ZeroAddress represents a FactoryV1__ZeroAddress error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FactoryV1ZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__ZeroAddress()
func RNContractFactoryV1FactoryV1ZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xe174662c0bcebc470566041356754029a1f1b25c1e8dac4ed279f825ed07f83e")
}

// UnpackFactoryV1ZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__ZeroAddress()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFactoryV1ZeroAddressError(raw []byte) (*RNContractFactoryV1FactoryV1ZeroAddress, error) {
	out := new(RNContractFactoryV1FactoryV1ZeroAddress)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1FailedCall represents a FailedCall error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func RNContractFactoryV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFailedCallError(raw []byte) (*RNContractFactoryV1FailedCall, error) {
	out := new(RNContractFactoryV1FailedCall)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1FailedDeployment represents a FailedDeployment error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1FailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func RNContractFactoryV1FailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackFailedDeploymentError(raw []byte) (*RNContractFactoryV1FailedDeployment, error) {
	out := new(RNContractFactoryV1FailedDeployment)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1InitCodeStubRuntimeTooLarge represents a InitCodeStub__RuntimeTooLarge error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1InitCodeStubRuntimeTooLarge struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InitCodeStub__RuntimeTooLarge(uint256 length)
func RNContractFactoryV1InitCodeStubRuntimeTooLargeErrorID() common.Hash {
	return common.HexToHash("0x9f7ba886379802fa5ad1852881971c83df47e39891aec0d97099ee3f16bd7fc6")
}

// UnpackInitCodeStubRuntimeTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InitCodeStub__RuntimeTooLarge(uint256 length)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackInitCodeStubRuntimeTooLargeError(raw []byte) (*RNContractFactoryV1InitCodeStubRuntimeTooLarge, error) {
	out := new(RNContractFactoryV1InitCodeStubRuntimeTooLarge)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "InitCodeStubRuntimeTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1InsufficientBalance represents a InsufficientBalance error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1InsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func RNContractFactoryV1InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackInsufficientBalanceError(raw []byte) (*RNContractFactoryV1InsufficientBalance, error) {
	out := new(RNContractFactoryV1InsufficientBalance)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1InvalidInitialization represents a InvalidInitialization error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func RNContractFactoryV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackInvalidInitializationError(raw []byte) (*RNContractFactoryV1InvalidInitialization, error) {
	out := new(RNContractFactoryV1InvalidInitialization)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1NotInitializing represents a NotInitializing error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func RNContractFactoryV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackNotInitializingError(raw []byte) (*RNContractFactoryV1NotInitializing, error) {
	out := new(RNContractFactoryV1NotInitializing)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func RNContractFactoryV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*RNContractFactoryV1RaylsAccessManagedContractPaused, error) {
	out := new(RNContractFactoryV1RaylsAccessManagedContractPaused)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func RNContractFactoryV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*RNContractFactoryV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(RNContractFactoryV1RaylsAccessManagedInvalidAuthority)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func RNContractFactoryV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*RNContractFactoryV1RaylsAccessManagedMustSchedule, error) {
	out := new(RNContractFactoryV1RaylsAccessManagedMustSchedule)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func RNContractFactoryV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*RNContractFactoryV1RaylsAccessManagedUnauthorized, error) {
	out := new(RNContractFactoryV1RaylsAccessManagedUnauthorized)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1ReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1ReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func RNContractFactoryV1ReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackReentrancyGuardReentrantCallError(raw []byte) (*RNContractFactoryV1ReentrancyGuardReentrantCall, error) {
	out := new(RNContractFactoryV1ReentrancyGuardReentrantCall)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func RNContractFactoryV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*RNContractFactoryV1UUPSUnauthorizedCallContext, error) {
	out := new(RNContractFactoryV1UUPSUnauthorizedCallContext)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNContractFactoryV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the RNContractFactoryV1 contract.
type RNContractFactoryV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func RNContractFactoryV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (rNContractFactoryV1 *RNContractFactoryV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*RNContractFactoryV1UUPSUnsupportedProxiableUUID, error) {
	out := new(RNContractFactoryV1UUPSUnsupportedProxiableUUID)
	if err := rNContractFactoryV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
