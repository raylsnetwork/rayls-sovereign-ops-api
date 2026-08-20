// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AbstractContractFactoryV1

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

// AbstractContractFactoryV1MetaData contains all meta data concerning the AbstractContractFactoryV1 contract.
var AbstractContractFactoryV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"RAYLS_ENYGMA_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ENYGMA_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_DVP_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_DVP_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC1155_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC20_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC20_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_DVP_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_DVP_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_ERC721_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_STABLECOIN_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RAYLS_STABLECOIN_TEST_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployEnygma\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc1155\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc1155Dvp\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc20\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc721\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployErc721Dvp\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployExternal\",\"inputs\":[{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployRegistered\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployRegisteredExternal\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"userArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployStableCoin\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factoryOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBytecodeHash\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEndpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFactoryOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBytecode\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEndpoint\",\"inputs\":[{\"name\":\"newEndpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFactoryOwner\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BytecodeSet\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"bytecodeHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractDeployed\",\"inputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EndpointUpdated\",\"inputs\":[{\"name\":\"oldEndpoint\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEndpoint\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FactoryOwnerUpdated\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegisteredContractDeployed\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deployedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FactoryV1__BytecodeNotRegistered\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FactoryV1__EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FactoryV1__InitializationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FactoryV1__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitCodeStub__RuntimeTooLarge\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__MustSchedule\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:  "AbstractContractFactoryV1",
}

// AbstractContractFactoryV1 is an auto generated Go binding around an Ethereum contract.
type AbstractContractFactoryV1 struct {
	abi abi.ABI
}

// NewAbstractContractFactoryV1 creates a new instance of AbstractContractFactoryV1.
func NewAbstractContractFactoryV1() *AbstractContractFactoryV1 {
	parsed, err := AbstractContractFactoryV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AbstractContractFactoryV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AbstractContractFactoryV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackRAYLSENYGMAKEY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d7ffb32.
//
// Solidity: function RAYLS_ENYGMA_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSENYGMAKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ENYGMA_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSENYGMAKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6d7ffb32.
//
// Solidity: function RAYLS_ENYGMA_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSENYGMAKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ENYGMA_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSENYGMATESTKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ENYGMA_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSENYGMATESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1c7fe28b.
//
// Solidity: function RAYLS_ENYGMA_TEST_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSENYGMATESTKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ENYGMA_TEST_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC1155DVPKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC1155_DVP_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155DVPKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x92c04d9b.
//
// Solidity: function RAYLS_ERC1155_DVP_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC1155DVPKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC1155_DVP_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC1155DVPTESTKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC1155_DVP_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155DVPTESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x29b0d701.
//
// Solidity: function RAYLS_ERC1155_DVP_TEST_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC1155DVPTESTKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC1155_DVP_TEST_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC1155KEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC1155_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155KEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5dd56210.
//
// Solidity: function RAYLS_ERC1155_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC1155KEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC1155_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC1155TESTKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC1155_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC1155TESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x403be3de.
//
// Solidity: function RAYLS_ERC1155_TEST_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC1155TESTKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC1155_TEST_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC20KEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC20_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC20KEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1464230b.
//
// Solidity: function RAYLS_ERC20_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC20KEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC20_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC20TESTKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC20_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC20TESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xee5ce29e.
//
// Solidity: function RAYLS_ERC20_TEST_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC20TESTKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC20_TEST_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC721DVPKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC721_DVP_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721DVPKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94f08898.
//
// Solidity: function RAYLS_ERC721_DVP_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC721DVPKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC721_DVP_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC721DVPTESTKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC721_DVP_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721DVPTESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95f50a62.
//
// Solidity: function RAYLS_ERC721_DVP_TEST_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC721DVPTESTKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC721_DVP_TEST_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC721KEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC721_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721KEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad10c149.
//
// Solidity: function RAYLS_ERC721_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC721KEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC721_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSERC721TESTKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_ERC721_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSERC721TESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfcab6982.
//
// Solidity: function RAYLS_ERC721_TEST_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSERC721TESTKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_ERC721_TEST_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSSTABLECOINKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_STABLECOIN_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSSTABLECOINKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb3872e4e.
//
// Solidity: function RAYLS_STABLECOIN_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSSTABLECOINKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_STABLECOIN_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackRAYLSSTABLECOINTESTKEY() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("RAYLS_STABLECOIN_TEST_KEY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRAYLSSTABLECOINTESTKEY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3fe9c16.
//
// Solidity: function RAYLS_STABLECOIN_TEST_KEY() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRAYLSSTABLECOINTESTKEY(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("RAYLS_STABLECOIN_TEST_KEY", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackAuthority() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("authority", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackContractVersion() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("contractVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackContractVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackContractVersion(data []byte) (*big.Int, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("contractVersion", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeploy(bytecode []byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deploy", bytecode, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeploy is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x19e011ad.
//
// Solidity: function deploy(bytes bytecode, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeploy(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deploy", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployEnygma(name string, symbol string, decimals uint8, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployEnygma", name, symbol, decimals, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployEnygma is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeff52a99.
//
// Solidity: function deployEnygma(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployEnygma(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployEnygma", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployErc1155(uri string, name string, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployErc1155", uri, name, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc1155 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00866e72.
//
// Solidity: function deployErc1155(string uri, string name, bytes32 resourceId) returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployErc1155(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployErc1155", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployErc1155Dvp(uri string, name string, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployErc1155Dvp", uri, name, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc1155Dvp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x614bcfd3.
//
// Solidity: function deployErc1155Dvp(string uri, string name, bytes32 resourceId) returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployErc1155Dvp(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployErc1155Dvp", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployErc20(name string, symbol string, decimals uint8, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployErc20", name, symbol, decimals, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc20 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd64e6e86.
//
// Solidity: function deployErc20(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployErc20(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployErc20", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployErc721(uri string, name string, symbol string, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployErc721", uri, name, symbol, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc721 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e9bb21e.
//
// Solidity: function deployErc721(string uri, string name, string symbol, bytes32 resourceId) returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployErc721(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployErc721", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployErc721Dvp(uri string, name string, symbol string, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployErc721Dvp", uri, name, symbol, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployErc721Dvp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfe75427d.
//
// Solidity: function deployErc721Dvp(string uri, string name, string symbol, bytes32 resourceId) returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployErc721Dvp(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployErc721Dvp", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployExternal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16381b15.
//
// Solidity: function deployExternal(bytes bytecode, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployExternal(bytecode []byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployExternal", bytecode, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployExternal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x16381b15.
//
// Solidity: function deployExternal(bytes bytecode, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployExternal(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployExternal", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployRegistered(key [32]byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployRegistered", key, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployRegistered is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcd39c4a8.
//
// Solidity: function deployRegistered(bytes32 key, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployRegistered(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployRegistered", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeployRegisteredExternal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f5e71fe.
//
// Solidity: function deployRegisteredExternal(bytes32 key, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployRegisteredExternal(key [32]byte, userArgs []byte, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployRegisteredExternal", key, userArgs, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployRegisteredExternal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f5e71fe.
//
// Solidity: function deployRegisteredExternal(bytes32 key, bytes userArgs, bytes32 resourceId) returns(address deployedAddress)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployRegisteredExternal(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployRegisteredExternal", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackDeployStableCoin(name string, symbol string, decimals uint8, resourceId [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("deployStableCoin", name, symbol, decimals, resourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeployStableCoin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfbc19051.
//
// Solidity: function deployStableCoin(string name, string symbol, uint8 decimals, bytes32 resourceId) returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackDeployStableCoin(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("deployStableCoin", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackEndpoint() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("endpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackEndpoint(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("endpoint", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackFactoryOwner() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("factoryOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFactoryOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4273601c.
//
// Solidity: function factoryOwner() view returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFactoryOwner(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("factoryOwner", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackGetBytecodeHash(key [32]byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("getBytecodeHash", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetBytecodeHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6670b8c5.
//
// Solidity: function getBytecodeHash(bytes32 key) view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackGetBytecodeHash(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("getBytecodeHash", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackGetEndpoint() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("getEndpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaed8e967.
//
// Solidity: function getEndpoint() view returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackGetEndpoint(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("getEndpoint", data)
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackGetFactoryOwner() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("getFactoryOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFactoryOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe09bda6b.
//
// Solidity: function getFactoryOwner() view returns(address)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackGetFactoryOwner(data []byte) (common.Address, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("getFactoryOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackProxiableUUID() []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := abstractContractFactoryV1.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackSetBytecode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c363c31.
//
// Solidity: function setBytecode(bytes32 key, bytes bytecode) returns()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackSetBytecode(key [32]byte, bytecode []byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("setBytecode", key, bytecode)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdbbb4155.
//
// Solidity: function setEndpoint(address newEndpoint) returns()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackSetEndpoint(newEndpoint common.Address) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("setEndpoint", newEndpoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFactoryOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3cc660c.
//
// Solidity: function setFactoryOwner(address newOwner) returns()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackSetFactoryOwner(newOwner common.Address) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("setFactoryOwner", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := abstractContractFactoryV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// AbstractContractFactoryV1AuthorityUpdated represents a AuthorityUpdated event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1AuthorityUpdated) ContractEventName() string {
	return AbstractContractFactoryV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*AbstractContractFactoryV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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

// AbstractContractFactoryV1BytecodeSet represents a BytecodeSet event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1BytecodeSet struct {
	Key          [32]byte
	BytecodeHash [32]byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1BytecodeSetEventName = "BytecodeSet"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1BytecodeSet) ContractEventName() string {
	return AbstractContractFactoryV1BytecodeSetEventName
}

// UnpackBytecodeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BytecodeSet(bytes32 indexed key, bytes32 bytecodeHash)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackBytecodeSetEvent(log *types.Log) (*AbstractContractFactoryV1BytecodeSet, error) {
	event := "BytecodeSet"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1BytecodeSet)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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

// AbstractContractFactoryV1ContractDeployed represents a ContractDeployed event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1ContractDeployed struct {
	DeployedAddress common.Address
	ResourceId      [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1ContractDeployedEventName = "ContractDeployed"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1ContractDeployed) ContractEventName() string {
	return AbstractContractFactoryV1ContractDeployedEventName
}

// UnpackContractDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ContractDeployed(address indexed deployedAddress, bytes32 indexed resourceId)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackContractDeployedEvent(log *types.Log) (*AbstractContractFactoryV1ContractDeployed, error) {
	event := "ContractDeployed"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1ContractDeployed)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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

// AbstractContractFactoryV1EndpointUpdated represents a EndpointUpdated event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1EndpointUpdated struct {
	OldEndpoint common.Address
	NewEndpoint common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1EndpointUpdatedEventName = "EndpointUpdated"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1EndpointUpdated) ContractEventName() string {
	return AbstractContractFactoryV1EndpointUpdatedEventName
}

// UnpackEndpointUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EndpointUpdated(address indexed oldEndpoint, address indexed newEndpoint)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackEndpointUpdatedEvent(log *types.Log) (*AbstractContractFactoryV1EndpointUpdated, error) {
	event := "EndpointUpdated"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1EndpointUpdated)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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

// AbstractContractFactoryV1FactoryOwnerUpdated represents a FactoryOwnerUpdated event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1FactoryOwnerUpdated struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1FactoryOwnerUpdatedEventName = "FactoryOwnerUpdated"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1FactoryOwnerUpdated) ContractEventName() string {
	return AbstractContractFactoryV1FactoryOwnerUpdatedEventName
}

// UnpackFactoryOwnerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FactoryOwnerUpdated(address indexed oldOwner, address indexed newOwner)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFactoryOwnerUpdatedEvent(log *types.Log) (*AbstractContractFactoryV1FactoryOwnerUpdated, error) {
	event := "FactoryOwnerUpdated"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1FactoryOwnerUpdated)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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

// AbstractContractFactoryV1Initialized represents a Initialized event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1Initialized) ContractEventName() string {
	return AbstractContractFactoryV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackInitializedEvent(log *types.Log) (*AbstractContractFactoryV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1Initialized)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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

// AbstractContractFactoryV1RegisteredContractDeployed represents a RegisteredContractDeployed event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1RegisteredContractDeployed struct {
	Key             [32]byte
	DeployedAddress common.Address
	ResourceId      [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1RegisteredContractDeployedEventName = "RegisteredContractDeployed"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1RegisteredContractDeployed) ContractEventName() string {
	return AbstractContractFactoryV1RegisteredContractDeployedEventName
}

// UnpackRegisteredContractDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RegisteredContractDeployed(bytes32 indexed key, address indexed deployedAddress, bytes32 indexed resourceId)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRegisteredContractDeployedEvent(log *types.Log) (*AbstractContractFactoryV1RegisteredContractDeployed, error) {
	event := "RegisteredContractDeployed"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1RegisteredContractDeployed)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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

// AbstractContractFactoryV1Upgraded represents a Upgraded event raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const AbstractContractFactoryV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (AbstractContractFactoryV1Upgraded) ContractEventName() string {
	return AbstractContractFactoryV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackUpgradedEvent(log *types.Log) (*AbstractContractFactoryV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != abstractContractFactoryV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AbstractContractFactoryV1Upgraded)
	if len(log.Data) > 0 {
		if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range abstractContractFactoryV1.abi.Events[event].Inputs {
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
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["FactoryV1BytecodeNotRegistered"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackFactoryV1BytecodeNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["FactoryV1EmptyBytecode"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackFactoryV1EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["FactoryV1InitializationFailed"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackFactoryV1InitializationFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["FactoryV1ZeroAddress"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackFactoryV1ZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["InitCodeStubRuntimeTooLarge"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackInitCodeStubRuntimeTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], abstractContractFactoryV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return abstractContractFactoryV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// AbstractContractFactoryV1AddressEmptyCode represents a AddressEmptyCode error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func AbstractContractFactoryV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackAddressEmptyCodeError(raw []byte) (*AbstractContractFactoryV1AddressEmptyCode, error) {
	out := new(AbstractContractFactoryV1AddressEmptyCode)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1Create2EmptyBytecode represents a Create2EmptyBytecode error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1Create2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func AbstractContractFactoryV1Create2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackCreate2EmptyBytecodeError(raw []byte) (*AbstractContractFactoryV1Create2EmptyBytecode, error) {
	out := new(AbstractContractFactoryV1Create2EmptyBytecode)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func AbstractContractFactoryV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackERC1967InvalidImplementationError(raw []byte) (*AbstractContractFactoryV1ERC1967InvalidImplementation, error) {
	out := new(AbstractContractFactoryV1ERC1967InvalidImplementation)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func AbstractContractFactoryV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackERC1967NonPayableError(raw []byte) (*AbstractContractFactoryV1ERC1967NonPayable, error) {
	out := new(AbstractContractFactoryV1ERC1967NonPayable)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1FactoryV1BytecodeNotRegistered represents a FactoryV1__BytecodeNotRegistered error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1FactoryV1BytecodeNotRegistered struct {
	Key [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__BytecodeNotRegistered(bytes32 key)
func AbstractContractFactoryV1FactoryV1BytecodeNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x6b714b57217f55da5c41092b90d46a27701a6e34bee2c6605a39b6a1cba654f7")
}

// UnpackFactoryV1BytecodeNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__BytecodeNotRegistered(bytes32 key)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFactoryV1BytecodeNotRegisteredError(raw []byte) (*AbstractContractFactoryV1FactoryV1BytecodeNotRegistered, error) {
	out := new(AbstractContractFactoryV1FactoryV1BytecodeNotRegistered)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1BytecodeNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1FactoryV1EmptyBytecode represents a FactoryV1__EmptyBytecode error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1FactoryV1EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__EmptyBytecode()
func AbstractContractFactoryV1FactoryV1EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x974918c11d78344eb3e9a3a905bdaa5941bfbd85677f790418e428875f34cc09")
}

// UnpackFactoryV1EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__EmptyBytecode()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFactoryV1EmptyBytecodeError(raw []byte) (*AbstractContractFactoryV1FactoryV1EmptyBytecode, error) {
	out := new(AbstractContractFactoryV1FactoryV1EmptyBytecode)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1FactoryV1InitializationFailed represents a FactoryV1__InitializationFailed error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1FactoryV1InitializationFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__InitializationFailed()
func AbstractContractFactoryV1FactoryV1InitializationFailedErrorID() common.Hash {
	return common.HexToHash("0x63cc0730a6f99cbea6ef1a62e46650dcbc73f7146679065ed2f148d9537f1e27")
}

// UnpackFactoryV1InitializationFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__InitializationFailed()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFactoryV1InitializationFailedError(raw []byte) (*AbstractContractFactoryV1FactoryV1InitializationFailed, error) {
	out := new(AbstractContractFactoryV1FactoryV1InitializationFailed)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1InitializationFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1FactoryV1ZeroAddress represents a FactoryV1__ZeroAddress error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1FactoryV1ZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FactoryV1__ZeroAddress()
func AbstractContractFactoryV1FactoryV1ZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xe174662c0bcebc470566041356754029a1f1b25c1e8dac4ed279f825ed07f83e")
}

// UnpackFactoryV1ZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FactoryV1__ZeroAddress()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFactoryV1ZeroAddressError(raw []byte) (*AbstractContractFactoryV1FactoryV1ZeroAddress, error) {
	out := new(AbstractContractFactoryV1FactoryV1ZeroAddress)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "FactoryV1ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1FailedCall represents a FailedCall error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func AbstractContractFactoryV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFailedCallError(raw []byte) (*AbstractContractFactoryV1FailedCall, error) {
	out := new(AbstractContractFactoryV1FailedCall)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1FailedDeployment represents a FailedDeployment error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1FailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func AbstractContractFactoryV1FailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackFailedDeploymentError(raw []byte) (*AbstractContractFactoryV1FailedDeployment, error) {
	out := new(AbstractContractFactoryV1FailedDeployment)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1InitCodeStubRuntimeTooLarge represents a InitCodeStub__RuntimeTooLarge error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1InitCodeStubRuntimeTooLarge struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InitCodeStub__RuntimeTooLarge(uint256 length)
func AbstractContractFactoryV1InitCodeStubRuntimeTooLargeErrorID() common.Hash {
	return common.HexToHash("0x9f7ba886379802fa5ad1852881971c83df47e39891aec0d97099ee3f16bd7fc6")
}

// UnpackInitCodeStubRuntimeTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InitCodeStub__RuntimeTooLarge(uint256 length)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackInitCodeStubRuntimeTooLargeError(raw []byte) (*AbstractContractFactoryV1InitCodeStubRuntimeTooLarge, error) {
	out := new(AbstractContractFactoryV1InitCodeStubRuntimeTooLarge)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "InitCodeStubRuntimeTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1InsufficientBalance represents a InsufficientBalance error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1InsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func AbstractContractFactoryV1InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackInsufficientBalanceError(raw []byte) (*AbstractContractFactoryV1InsufficientBalance, error) {
	out := new(AbstractContractFactoryV1InsufficientBalance)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1InvalidInitialization represents a InvalidInitialization error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func AbstractContractFactoryV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackInvalidInitializationError(raw []byte) (*AbstractContractFactoryV1InvalidInitialization, error) {
	out := new(AbstractContractFactoryV1InvalidInitialization)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1NotInitializing represents a NotInitializing error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func AbstractContractFactoryV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackNotInitializingError(raw []byte) (*AbstractContractFactoryV1NotInitializing, error) {
	out := new(AbstractContractFactoryV1NotInitializing)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func AbstractContractFactoryV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*AbstractContractFactoryV1RaylsAccessManagedContractPaused, error) {
	out := new(AbstractContractFactoryV1RaylsAccessManagedContractPaused)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func AbstractContractFactoryV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*AbstractContractFactoryV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(AbstractContractFactoryV1RaylsAccessManagedInvalidAuthority)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func AbstractContractFactoryV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*AbstractContractFactoryV1RaylsAccessManagedMustSchedule, error) {
	out := new(AbstractContractFactoryV1RaylsAccessManagedMustSchedule)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func AbstractContractFactoryV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*AbstractContractFactoryV1RaylsAccessManagedUnauthorized, error) {
	out := new(AbstractContractFactoryV1RaylsAccessManagedUnauthorized)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1ReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1ReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func AbstractContractFactoryV1ReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackReentrancyGuardReentrantCallError(raw []byte) (*AbstractContractFactoryV1ReentrancyGuardReentrantCall, error) {
	out := new(AbstractContractFactoryV1ReentrancyGuardReentrantCall)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func AbstractContractFactoryV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*AbstractContractFactoryV1UUPSUnauthorizedCallContext, error) {
	out := new(AbstractContractFactoryV1UUPSUnauthorizedCallContext)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AbstractContractFactoryV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the AbstractContractFactoryV1 contract.
type AbstractContractFactoryV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func AbstractContractFactoryV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (abstractContractFactoryV1 *AbstractContractFactoryV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*AbstractContractFactoryV1UUPSUnsupportedProxiableUUID, error) {
	out := new(AbstractContractFactoryV1UUPSUnsupportedProxiableUUID)
	if err := abstractContractFactoryV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
