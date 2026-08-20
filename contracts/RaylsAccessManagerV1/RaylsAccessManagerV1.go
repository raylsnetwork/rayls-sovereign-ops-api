// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RaylsAccessManagerV1

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

// IRaylsAccessManagerRoleInfo is an auto generated low-level Go binding around an user-defined struct.
type IRaylsAccessManagerRoleInfo struct {
	RoleId       uint64
	Label        string
	AdminRole    uint64
	GuardianRole uint64
	GrantDelay   uint32
	MemberCount  *big.Int
}

// IRaylsAccessManagerSelectorRoleMapping is an auto generated low-level Go binding around an user-defined struct.
type IRaylsAccessManagerSelectorRoleMapping struct {
	RoleName  string
	Selectors [][4]byte
}

// AccessManagerAuthLibMetaData contains all meta data concerning the AccessManagerAuthLib contract.
var AccessManagerAuthLibMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "bf60ea6c405102ea4f38b2dc2b36bfe85c",
}

// AccessManagerAuthLib is an auto generated Go binding around an Ethereum contract.
type AccessManagerAuthLib struct {
	abi abi.ABI
}

// NewAccessManagerAuthLib creates a new instance of AccessManagerAuthLib.
func NewAccessManagerAuthLib() *AccessManagerAuthLib {
	parsed, err := AccessManagerAuthLibMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AccessManagerAuthLib{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AccessManagerAuthLib) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// AccessManagerContractScopedLibMetaData contains all meta data concerning the AccessManagerContractScopedLib contract.
var AccessManagerContractScopedLibMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "8b8cf5c657cd0bf2c15e86850df9221a20",
}

// AccessManagerContractScopedLib is an auto generated Go binding around an Ethereum contract.
type AccessManagerContractScopedLib struct {
	abi abi.ABI
}

// NewAccessManagerContractScopedLib creates a new instance of AccessManagerContractScopedLib.
func NewAccessManagerContractScopedLib() *AccessManagerContractScopedLib {
	parsed, err := AccessManagerContractScopedLibMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AccessManagerContractScopedLib{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AccessManagerContractScopedLib) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// AccessManagerEnumerationLibMetaData contains all meta data concerning the AccessManagerEnumerationLib contract.
var AccessManagerEnumerationLibMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "2ed7d8868861269722eac54306f1710b68",
}

// AccessManagerEnumerationLib is an auto generated Go binding around an Ethereum contract.
type AccessManagerEnumerationLib struct {
	abi abi.ABI
}

// NewAccessManagerEnumerationLib creates a new instance of AccessManagerEnumerationLib.
func NewAccessManagerEnumerationLib() *AccessManagerEnumerationLib {
	parsed, err := AccessManagerEnumerationLibMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AccessManagerEnumerationLib{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AccessManagerEnumerationLib) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// AccessManagerRoleConfigLibMetaData contains all meta data concerning the AccessManagerRoleConfigLib contract.
var AccessManagerRoleConfigLibMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "b4ea0fc60807818dfa02c533edf487ccb8",
}

// AccessManagerRoleConfigLib is an auto generated Go binding around an Ethereum contract.
type AccessManagerRoleConfigLib struct {
	abi abi.ABI
}

// NewAccessManagerRoleConfigLib creates a new instance of AccessManagerRoleConfigLib.
func NewAccessManagerRoleConfigLib() *AccessManagerRoleConfigLib {
	parsed, err := AccessManagerRoleConfigLibMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AccessManagerRoleConfigLib{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AccessManagerRoleConfigLib) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// AccessManagerScheduleLibMetaData contains all meta data concerning the AccessManagerScheduleLib contract.
var AccessManagerScheduleLibMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "f4bcc756bc578e056023d7c6ee59f861d4",
}

// AccessManagerScheduleLib is an auto generated Go binding around an Ethereum contract.
type AccessManagerScheduleLib struct {
	abi abi.ABI
}

// NewAccessManagerScheduleLib creates a new instance of AccessManagerScheduleLib.
func NewAccessManagerScheduleLib() *AccessManagerScheduleLib {
	parsed, err := AccessManagerScheduleLibMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AccessManagerScheduleLib{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AccessManagerScheduleLib) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// RaylsAccessManagerV1MetaData contains all meta data concerning the RaylsAccessManagerV1 contract.
var RaylsAccessManagerV1MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADMIN\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXPIRATION\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_OWNER\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"},{\"internalType\":\"uint64[]\",\"name\":\"roleIds\",\"type\":\"uint64[]\"}],\"name\":\"addFunctionAllowedRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"}],\"name\":\"getAccountContractScopedRoles\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountRoles\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountRolesWithInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"adminRole\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"guardianRole\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"grantDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"memberCount\",\"type\":\"uint256\"}],\"internalType\":\"structIRaylsAccessManager.RoleInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRoles\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"adminRole\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"guardianRole\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"grantDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"memberCount\",\"type\":\"uint256\"}],\"internalType\":\"structIRaylsAccessManager.RoleInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"}],\"name\":\"getContractAuthority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"}],\"name\":\"getContractScopedRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getFunctionAllowedRoles\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getFunctionAllowedRolesWithInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"adminRole\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"guardianRole\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"grantDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"memberCount\",\"type\":\"uint256\"}],\"internalType\":\"structIRaylsAccessManager.RoleInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredRoleCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"name\":\"getRoleGuardian\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getRoleIdByName\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"name\":\"getRoleInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"adminRole\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"guardianRole\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"grantDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"memberCount\",\"type\":\"uint256\"}],\"internalType\":\"structIRaylsAccessManager.RoleInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"roleIds\",\"type\":\"uint64[]\"}],\"name\":\"getRoleInfoBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"adminRole\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"guardianRole\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"grantDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"memberCount\",\"type\":\"uint256\"}],\"internalType\":\"structIRaylsAccessManager.RoleInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"name\":\"getRoleLabel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"}],\"name\":\"getSchedule\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"when\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionDelay\",\"type\":\"uint32\"}],\"name\":\"grantContractScopedRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionDelay\",\"type\":\"uint32\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantSelfTokenOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"}],\"name\":\"hasContractScopedRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isMember\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"executionDelay\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isMember\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"executionDelay\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoleByName\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialAdmin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"}],\"name\":\"isContractPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"labelRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerRole\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"registerRoleAndLabel\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"},{\"internalType\":\"uint64[]\",\"name\":\"roleIds\",\"type\":\"uint64[]\"}],\"name\":\"removeFunctionAllowedRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"}],\"name\":\"revokeContractScopedRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"when\",\"type\":\"uint48\"}],\"name\":\"schedule\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"ownerSelectors\",\"type\":\"bytes4[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"roleName\",\"type\":\"string\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIRaylsAccessManager.SelectorRoleMapping[]\",\"name\":\"roleMappings\",\"type\":\"tuple[]\"}],\"name\":\"selfRegisterManagedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setContractPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"newDelay\",\"type\":\"uint32\"}],\"name\":\"setGrantDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"adminRoleId\",\"type\":\"uint64\"}],\"name\":\"setRoleAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"guardianRoleId\",\"type\":\"uint64\"}],\"name\":\"setRoleGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAuthority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"ContractPauseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"executionDelay\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"activeSince\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"grantor\",\"type\":\"address\"}],\"name\":\"ContractScopedRoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"}],\"name\":\"ContractScopedRoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"name\":\"FunctionAllowedRoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"}],\"name\":\"FunctionAllowedRoleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAuthority\",\"type\":\"address\"}],\"name\":\"ManagedContractRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"}],\"name\":\"OperationCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"}],\"name\":\"OperationExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managedContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"executeAfter\",\"type\":\"uint48\"}],\"name\":\"OperationScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newAdmin\",\"type\":\"uint64\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newDelay\",\"type\":\"uint32\"}],\"name\":\"RoleGrantDelayChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"executionDelay\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"activeSince\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantor\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newGuardian\",\"type\":\"uint64\"}],\"name\":\"RoleGuardianChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"RoleLabelSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"RoleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"roleId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaylsAccessManaged__ContractPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"RaylsAccessManaged__MustSchedule\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RaylsAccessManaged__Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaylsAccessManagerV1__ContractPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RaylsAccessManagerV1__Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	ID:  "270a97eae1f28152c7a5adb2bfcc8ebaf0",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516137956200010460003960008181611f1401528181612002015261202b01526137956000f3fe6080604052600436106102695760003560e01c8063a5808e2f11610145578063c5050c1f116100bc578063c5050c1f146107b3578063c554ae82146107c8578063cacebadc146107e8578063d1f856ee14610815578063d5c4ca0014610835578063d6bb62c614610855578063f11545ec14610875578063f222bbfa14610895578063f2bcac3d146108b5578063f317a850146108ca578063f801a698146108ea578063fe0776f51461090a578063ff6c33e61461092a57600080fd5b8063a5808e2f14610603578063a64d95ce14610630578063ad3cb1cc14610650578063b0f008eb1461068e578063b7009613146106ae578063b7d2b162146106f2578063ba126a5814610712578063bb4b573414610732578063bf7e214f14610749578063c1179a7f1461075e578063c349783714610773578063c4d66de81461079357600080fd5b80633bd37148116101e45780633bd37148146104375780634f1ef28614610457578063529629521461046a57806352d1902d1461048a578063530dd456146104ad57806357087597146104cd5780635bfa5027146104ed5780635cdd11a11461051a5780635e80a70b1461053a57806365e7d5c71461055a5780636f5e800e14610587578063804af8c5146105a7578063853551b8146105e357600080fd5b80630140c5211461026e5780630649fea41461029057806308c47967146102c65780630b0a93ba146102e657806312c1e07214610313578063180cb42b146103335780631b9bfca9146103635780631cff79cd1461038357806325c471a0146103ab5780632a0acc6a146103cb57806330cae187146103e05780633adc277a14610400575b600080fd5b34801561027a57600080fd5b5061028e6102893660046124ac565b61093f565b005b34801561029c57600080fd5b506102b06102ab36600461254b565b6109b5565b6040516102bd9190612638565b60405180910390f35b3480156102d257600080fd5b5061028e6102e13660046126b1565b610a3f565b3480156102f257600080fd5b506103066103013660046126fc565b610aaf565b6040516102bd9190612719565b34801561031f57600080fd5b5061028e61032e36600461273f565b610b2a565b34801561033f57600080fd5b5061035361034e3660046127dc565b610bbe565b60405190151581526020016102bd565b34801561036f57600080fd5b5061030661037e366004612827565b610c47565b610396610391366004612892565b610cd1565b60405163ffffffff90911681526020016102bd565b3480156103b757600080fd5b5061028e6103c63660046128e6565b610f41565b3480156103d757600080fd5b50610306600081565b3480156103ec57600080fd5b5061028e6103fb366004612926565b610f9c565b34801561040c57600080fd5b5061042061041b36600461295f565b611014565b60405165ffffffffffff90911681526020016102bd565b34801561044357600080fd5b506102b0610452366004612978565b61108b565b61028e610465366004612a26565b6110c7565b34801561047657600080fd5b5061028e610485366004612926565b6110d5565b34801561049657600080fd5b5061049f61111d565b6040519081526020016102bd565b3480156104b957600080fd5b506103066104c83660046126fc565b61113a565b3480156104d957600080fd5b506103066104e8366004612ab8565b611174565b3480156104f957600080fd5b5061050d6105083660046126fc565b6111f1565b6040516102bd9190612aed565b34801561052657600080fd5b5061028e6105353660046124ac565b61129b565b34801561054657600080fd5b5061028e610555366004612b00565b6112da565b34801561056657600080fd5b5061057a610575366004612b00565b611344565b6040516102bd9190612b1d565b34801561059357600080fd5b5061028e6105a23660046124ac565b6113bf565b3480156105b357600080fd5b506105c76105c23660046126b1565b6113fe565b60408051921515835263ffffffff9091166020830152016102bd565b3480156105ef57600080fd5b5061028e6105fe366004612b31565b611488565b34801561060f57600080fd5b5061062361061e3660046126fc565b6114c3565b6040516102bd9190612b51565b34801561063c57600080fd5b5061028e61064b366004612b9e565b611542565b34801561065c57600080fd5b50610681604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516102bd9190612bcc565b34801561069a57600080fd5b506103536106a9366004612b00565b61158e565b3480156106ba57600080fd5b506106ce6106c9366004612bdf565b611609565b60408051931515845263ffffffff90921660208401521515908201526060016102bd565b3480156106fe57600080fd5b5061028e61070d366004612c26565b611699565b34801561071e57600080fd5b5061028e61072d366004612c62565b6116d2565b34801561073e57600080fd5b5061042062093a8081565b34801561075557600080fd5b5061057a61171a565b34801561076a57600080fd5b50610306600181565b34801561077f57600080fd5b506102b061078e366004612b00565b611733565b34801561079f57600080fd5b5061028e6107ae366004612b00565b6117b2565b3480156107bf57600080fd5b506103066119ae565b3480156107d457600080fd5b506106816107e33660046126fc565b6119ca565b3480156107f457600080fd5b50610808610803366004612b00565b611a49565b6040516102bd9190612c90565b34801561082157600080fd5b506105c7610830366004612c26565b611ac8565b34801561084157600080fd5b5061080861085036600461254b565b611b50565b34801561086157600080fd5b50610396610870366004612cd1565b611bd1565b34801561088157600080fd5b50610808610890366004612d29565b611c52565b3480156108a157600080fd5b506106236108b0366004612c26565b611c9d565b3480156108c157600080fd5b506102b0611d1e565b3480156108d657600080fd5b506103066108e5366004612ab8565b611d96565b3480156108f657600080fd5b5061049f610905366004612d5b565b611dd2565b34801561091657600080fd5b5061028e610925366004612c26565b611e53565b34801561093657600080fd5b50610306600281565b604051630140c52160e01b815273__$b4ea0fc60807818dfa02c533edf487ccb8$__90630140c5219061097e9088908890889088908890600401612e3d565b60006040518083038186803b15801561099657600080fd5b505af41580156109aa573d6000803e3d6000fd5b505050505050505050565b6040516301927fa960e21b815260609073__$2ed7d8868861269722eac54306f1710b68$__90630649fea4906109f19086908690600401612e81565b600060405180830381865af4158015610a0e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a369190810190612fd9565b90505b92915050565b6040516308c4796760e01b815273__$8b8cf5c657cd0bf2c15e86850df9221a20$__906308c4796790610a7a90869086908690600401613089565b60006040518083038186803b158015610a9257600080fd5b505af4158015610aa6573d6000803e3d6000fd5b50505050505050565b60405163058549dd60e11b815260009073__$b4ea0fc60807818dfa02c533edf487ccb8$__90630b0a93ba90610ae9908590600401612719565b602060405180830381865af4158015610b06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3991906130b3565b604051630960f03960e11b81526001600160401b03851660048201526001600160a01b0380851660248301528316604482015263ffffffff8216606482015273__$8b8cf5c657cd0bf2c15e86850df9221a20$__906312c1e0729060840160006040518083038186803b158015610ba057600080fd5b505af4158015610bb4573d6000803e3d6000fd5b5050505050505050565b60405163180cb42b60e01b815260009073__$b4ea0fc60807818dfa02c533edf487ccb8$__9063180cb42b90610bfc908790879087906004016130f9565b602060405180830381865af4158015610c19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3d9190613125565b90505b9392505050565b604051631b9bfca960e01b815260009073__$b4ea0fc60807818dfa02c533edf487ccb8$__90631b9bfca990610c87908890889088908890600401613142565b602060405180830381865af4158015610ca4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc891906130b3565b95945050505050565b60008073__$f4bcc756bc578e056023d7c6ee59f861d4$__6333ddcca1338787876040518563ffffffff1660e01b8152600401610d119493929190613174565b602060405180830381865af4158015610d2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5291906131ab565b905060008073__$bf60ea6c405102ea4f38b2dc2b36bfe85c$__63b70096133389610d806004868b8d6131c4565b610d89916131ee565b6040518463ffffffff1660e01b8152600401610da79392919061321e565b606060405180830381865af4158015610dc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de8919061324b565b90955090925090508015610e0f57604051633dbfd5d960e11b815260040160405180910390fd5b81158015610e21575063ffffffff8416155b15610e4a573360405163f96b0e6560e01b8152600401610e419190612b1d565b60405180910390fd5b6000610e54611e8c565b90506001816009016000828254610e6b91906132a3565b92505081905550600080896001600160a01b0316348a8a604051610e909291906132b6565b60006040518083038185875af1925050503d8060008114610ecd576040519150601f19603f3d011682016040523d82523d6000602084013e610ed2565b606091505b50915091506001836009016000828254610eec91906132c6565b90915550829050610eff57805160208201fd5b6040516001600160a01b038b169087907e1d1d6b8d664ddc069841e54d003c4a632a7392aa33ef7f83806380c2bc687390600090a35050505050509392505050565b60405163012e238d60e51b81526001600160401b03841660048201526001600160a01b038316602482015263ffffffff8216604482015273__$b4ea0fc60807818dfa02c533edf487ccb8$__906325c471a090606401610a7a565b6040516330cae18760e01b81526001600160401b0380841660048301528216602482015273__$b4ea0fc60807818dfa02c533edf487ccb8$__906330cae187906044015b60006040518083038186803b158015610ff857600080fd5b505af415801561100c573d6000803e3d6000fd5b505050505050565b604051631d6e13bd60e11b81526004810182905260009073__$f4bcc756bc578e056023d7c6ee59f861d4$__90633adc277a90602401602060405180830381865af4158015611067573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3991906132d9565b60405163077a6e2960e31b815260609073__$2ed7d8868861269722eac54306f1710b68$__90633bd37148906109f190869086906004016132f6565b6110d18282611eee565b5050565b60405163294b14a960e11b81526001600160401b0380841660048301528216602482015273__$b4ea0fc60807818dfa02c533edf487ccb8$__90635296295290604401610fe0565b6000611127611f09565b5060008051602061374083398151915290565b604051632986ea2b60e11b815260009073__$b4ea0fc60807818dfa02c533edf487ccb8$__9063530dd45690610ae9908590600401612719565b604051635708759760e01b815260009073__$b4ea0fc60807818dfa02c533edf487ccb8$__906357087597906111b0908690869060040161330a565b602060405180830381865af41580156111cd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3691906130b3565b6040805160c081018252600080825260606020830181905282840182905282018190526080820181905260a08201529051635bfa502760e01b815273__$2ed7d8868861269722eac54306f1710b68$__90635bfa502790611256908590600401612719565b600060405180830381865af4158015611273573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a39919081019061331e565b604051635cdd11a160e01b815273__$b4ea0fc60807818dfa02c533edf487ccb8$__90635cdd11a19061097e9088908890889088908890600401612e3d565b604051635e80a70b60e01b815273__$8b8cf5c657cd0bf2c15e86850df9221a20$__90635e80a70b90611311908490600401612b1d565b60006040518083038186803b15801561132957600080fd5b505af415801561133d573d6000803e3d6000fd5b5050505050565b6040516365e7d5c760e01b815260009073__$8b8cf5c657cd0bf2c15e86850df9221a20$__906365e7d5c79061137e908590600401612b1d565b602060405180830381865af415801561139b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a399190613352565b604051639b40ef6d60e01b815273__$8b8cf5c657cd0bf2c15e86850df9221a20$__90639b40ef6d9061097e908890889088908890889060040161336f565b60008073__$8b8cf5c657cd0bf2c15e86850df9221a20$__63804af8c58686866040518463ffffffff1660e01b815260040161143c93929190613089565b6040805180830381865af4158015611458573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147c91906134aa565b91509150935093915050565b6040516310a6aa3760e31b815273__$b4ea0fc60807818dfa02c533edf487ccb8$__9063853551b890610a7a908690869086906004016134d9565b60405163a5808e2f60e01b815260609073__$2ed7d8868861269722eac54306f1710b68$__9063a5808e2f906114fd908590600401612719565b600060405180830381865af415801561151a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a3991908101906134fc565b604051635326cae760e11b81526001600160401b038316600482015263ffffffff8216602482015273__$b4ea0fc60807818dfa02c533edf487ccb8$__9063a64d95ce90604401610fe0565b60405163b0f008eb60e01b815260009073__$b4ea0fc60807818dfa02c533edf487ccb8$__9063b0f008eb906115c8908590600401612b1d565b602060405180830381865af41580156115e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a399190613125565b600080600073__$bf60ea6c405102ea4f38b2dc2b36bfe85c$__63b70096138787876040518463ffffffff1660e01b81526004016116499392919061321e565b606060405180830381865af4158015611666573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061168a919061324b565b92509250925093509350939050565b604051635be958b160e11b815273__$b4ea0fc60807818dfa02c533edf487ccb8$__9063b7d2b16290610fe0908590859060040161358a565b6040516317424d4b60e31b81526001600160a01b0383166004820152811515602482015273__$b4ea0fc60807818dfa02c533edf487ccb8$__9063ba126a5890604401610fe0565b6000611724611f54565b546001600160a01b0316919050565b60405163c349783760e01b815260609073__$2ed7d8868861269722eac54306f1710b68$__9063c34978379061176d908590600401612b1d565b600060405180830381865af415801561178a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a399190810190612fd9565b60006117bc611f85565b805490915060ff600160401b82041615906001600160401b03166000811580156117e35750825b90506000826001600160401b031660011480156117ff5750303b155b90508115801561180d575080155b1561182b5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561185557845460ff60401b1916600160401b1785555b61185d611fae565b61186630611fb6565b604080516101408101825263012e238d60e51b8152635be958b160e11b6020820152630960f03960e11b818301526308c4796760e01b606082015263fe0776f560e01b60808201526337af400760e11b60a0820152635e80a70b60e01b60c0820152631f0034d360e31b60e0820152631cff79cd60e01b610100820152636b5db16360e11b610120820152905163d584c82160e01b815273__$b4ea0fc60807818dfa02c533edf487ccb8$__9063d584c8219061192b908a90309086906004016135ac565b60006040518083038186803b15801561194357600080fd5b505af4158015611957573d6000803e3d6000fd5b5050505050831561100c57845460ff60401b191685556040517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29061199e90600190612719565b60405180910390a1505050505050565b60006119b8611e8c565b600401546001600160401b0316919050565b6040516362aa574160e11b815260609073__$b4ea0fc60807818dfa02c533edf487ccb8$__9063c554ae8290611a04908590600401612719565b600060405180830381865af4158015611a21573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a399190810190613606565b6040516332b3aeb760e21b815260609073__$2ed7d8868861269722eac54306f1710b68$__9063cacebadc90611a83908590600401612b1d565b600060405180830381865af4158015611aa0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a39919081019061363a565b60008073__$b4ea0fc60807818dfa02c533edf487ccb8$__63d1f856ee85856040518363ffffffff1660e01b8152600401611b0492919061358a565b6040805180830381865af4158015611b20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4491906134aa565b915091505b9250929050565b604051636bcd4f2760e11b815260609073__$2ed7d8868861269722eac54306f1710b68$__9063d79a9e4e90611b8c9086908690600401612e81565b600060405180830381865af4158015611ba9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a36919081019061363a565b604051636b5db16360e11b815260009073__$f4bcc756bc578e056023d7c6ee59f861d4$__9063d6bb62c690611c11908890889088908890600401613174565b602060405180830381865af4158015611c2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc891906136c8565b604051633c45517b60e21b81526001600160a01b0380841660048301528216602482015260609073__$2ed7d8868861269722eac54306f1710b68$__9063f11545ec90604401611b8c565b6040516379115dfd60e11b815260609073__$2ed7d8868861269722eac54306f1710b68$__9063f222bbfa90611cd9908690869060040161358a565b600060405180830381865af4158015611cf6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a3691908101906134fc565b606073__$2ed7d8868861269722eac54306f1710b68$__63f2bcac3d6040518163ffffffff1660e01b8152600401600060405180830381865af4158015611d69573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d919190810190612fd9565b905090565b604051630f317a8560e41b815260009073__$b4ea0fc60807818dfa02c533edf487ccb8$__9063f317a850906111b0908690869060040161330a565b604051631f0034d360e31b815260009073__$f4bcc756bc578e056023d7c6ee59f861d4$__9063f801a69890611e129088908890889088906004016136e5565b602060405180830381865af4158015611e2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc891906131ab565b60405163fe0776f560e01b815273__$b4ea0fc60807818dfa02c533edf487ccb8$__9063fe0776f590610fe0908590859060040161358a565b60008060ff19611ebd60017ffc9aaa9eebaa8d3980e1146ee9e228a3853058d792a5753f711093f99240e5ae6132c6565b604051602001611ecf91815260200190565b60408051601f1981840301815291905280516020909101201692915050565b611ef6611ff7565b611eff82612085565b6110d18282612099565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611f525760405163703e46dd60e11b815260040160405180910390fd5b565b60008060ff19611ebd60017f2d9a26166e6ed6927c3e14d89c4d572a357f984344fcaed3c38c3ff5cdb83f356132c6565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610a39565b611f52612152565b6000611fc0611f54565b80549091506001600160a01b031615611fee5781604051638944034760e01b8152600401610e419190612b1d565b6110d182612177565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061206757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661205b612207565b6001600160a01b031614155b15611f525760405163703e46dd60e11b815260040160405180910390fd5b612096612090611e8c565b3361221d565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156120f3575060408051601f3d908101601f191682019092526120f0918101906131ab565b60015b6121125781604051634c9c8ce360e01b8152600401610e419190612b1d565b600080516020613740833981519152811461214357604051632a87526960e21b815260048101829052602401610e41565b61214d8383612280565b505050565b61215a6122d6565b611f5257604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b0381166121a05780604051638944034760e01b8152600401610e419190612b1d565b60006121aa611f54565b80546040519192506001600160a01b03808516929116907fa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b7638998019890600090a380546001600160a01b0319166001600160a01b0392909216919091179055565b6000600080516020613740833981519152611724565b6000808052602083815260408083206001600160a01b03851684526002019091529020805465ffffffffffff161580612260575080544265ffffffffffff909116115b1561214d578160405163f96b0e6560e01b8152600401610e419190612b1d565b612289826122f0565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156122ce5761214d828261234c565b6110d16123b9565b60006122e0611f85565b54600160401b900460ff16919050565b806001600160a01b03163b60000361231d5780604051634c9c8ce360e01b8152600401610e419190612b1d565b60008051602061374083398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516123699190613723565b600060405180830381855af49150503d80600081146123a4576040519150601f19603f3d011682016040523d82523d6000602084013e6123a9565b606091505b5091509150610cc88583836123d8565b3415611f525760405163b398979f60e01b815260040160405180910390fd5b6060826123ed576123e88261242b565b610c40565b815115801561240457506001600160a01b0384163b155b156124245783604051639996b31560e01b8152600401610e419190612b1d565b5080610c40565b80511561243a57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b6001600160a01b038116811461209657600080fd5b60008083601f84011261247a57600080fd5b5081356001600160401b0381111561249157600080fd5b6020830191508360208260051b8501011115611b4957600080fd5b6000806000806000606086880312156124c457600080fd5b85356124cf81612453565b945060208601356001600160401b03808211156124eb57600080fd5b6124f789838a01612468565b9096509450604088013591508082111561251057600080fd5b5061251d88828901612468565b969995985093965092949392505050565b80356001600160e01b03198116811461254657600080fd5b919050565b6000806040838503121561255e57600080fd5b823561256981612453565b91506125776020840161252e565b90509250929050565b60005b8381101561259b578181015183820152602001612583565b50506000910152565b600081518084526125bc816020860160208601612580565b601f01601f19169290920160200192915050565b60006001600160401b03808351168452602083015160c060208601526125f960c08601826125a4565b905081604085015116604086015281606085015116606086015263ffffffff608085015116608086015260a084015160a0860152809250505092915050565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b8281101561268f57603f1988860301845261267d8583516125d0565b94509285019290850190600101612661565b5092979650505050505050565b6001600160401b038116811461209657600080fd5b6000806000606084860312156126c657600080fd5b83356126d18161269c565b925060208401356126e181612453565b915060408401356126f181612453565b809150509250925092565b60006020828403121561270e57600080fd5b8135610c408161269c565b6001600160401b0391909116815260200190565b63ffffffff8116811461209657600080fd5b6000806000806080858703121561275557600080fd5b84356127608161269c565b9350602085013561277081612453565b9250604085013561278081612453565b915060608501356127908161272d565b939692955090935050565b60008083601f8401126127ad57600080fd5b5081356001600160401b038111156127c457600080fd5b602083019150836020828501011115611b4957600080fd5b6000806000604084860312156127f157600080fd5b83356001600160401b0381111561280757600080fd5b6128138682870161279b565b90945092505060208401356126f181612453565b6000806000806040858703121561283d57600080fd5b84356001600160401b038082111561285457600080fd5b6128608883890161279b565b9096509450602087013591508082111561287957600080fd5b506128868782880161279b565b95989497509550505050565b6000806000604084860312156128a757600080fd5b83356128b281612453565b925060208401356001600160401b038111156128cd57600080fd5b6128d98682870161279b565b9497909650939450505050565b6000806000606084860312156128fb57600080fd5b83356129068161269c565b9250602084013561291681612453565b915060408401356126f18161272d565b6000806040838503121561293957600080fd5b82356129448161269c565b915060208301356129548161269c565b809150509250929050565b60006020828403121561297157600080fd5b5035919050565b6000806020838503121561298b57600080fd5b82356001600160401b038111156129a157600080fd5b6129ad85828601612468565b90969095509350505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156129f7576129f76129b9565b604052919050565b60006001600160401b03821115612a1857612a186129b9565b50601f01601f191660200190565b60008060408385031215612a3957600080fd5b8235612a4481612453565b915060208301356001600160401b03811115612a5f57600080fd5b8301601f81018513612a7057600080fd5b8035612a83612a7e826129ff565b6129cf565b818152866020838501011115612a9857600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008060208385031215612acb57600080fd5b82356001600160401b03811115612ae157600080fd5b6129ad8582860161279b565b602081526000610a3660208301846125d0565b600060208284031215612b1257600080fd5b8135610c4081612453565b6001600160a01b0391909116815260200190565b600080600060408486031215612b4657600080fd5b83356128b28161269c565b6020808252825182820181905260009190848201906040850190845b81811015612b925783516001600160a01b031683529284019291840191600101612b6d565b50909695505050505050565b60008060408385031215612bb157600080fd5b8235612bbc8161269c565b915060208301356129548161272d565b602081526000610a3660208301846125a4565b600080600060608486031215612bf457600080fd5b8335612bff81612453565b92506020840135612c0f81612453565b9150612c1d6040850161252e565b90509250925092565b60008060408385031215612c3957600080fd5b8235612c448161269c565b9150602083013561295481612453565b801515811461209657600080fd5b60008060408385031215612c7557600080fd5b8235612c8081612453565b9150602083013561295481612c54565b6020808252825182820181905260009190848201906040850190845b81811015612b925783516001600160401b031683529284019291840191600101612cac565b60008060008060608587031215612ce757600080fd5b8435612cf281612453565b93506020850135612d0281612453565b925060408501356001600160401b03811115612d1d57600080fd5b6128868782880161279b565b60008060408385031215612d3c57600080fd5b8235612c4481612453565b65ffffffffffff8116811461209657600080fd5b60008060008060608587031215612d7157600080fd5b8435612d7c81612453565b935060208501356001600160401b03811115612d9757600080fd5b612da38782880161279b565b909450925050604085013561279081612d47565b8183526000602080850194508260005b85811015612df4576001600160e01b0319612de18361252e565b1687529582019590820190600101612dc7565b509495945050505050565b8183526000602080850194508260005b85811015612df4578135612e228161269c565b6001600160401b031687529582019590820190600101612e0f565b6001600160a01b0386168152606060208201819052600090612e629083018688612db7565b8281036040840152612e75818587612dff565b98975050505050505050565b6001600160a01b039290921682526001600160e01b031916602082015260400190565b60006001600160401b03821115612ebd57612ebd6129b9565b5060051b60200190565b600082601f830112612ed857600080fd5b8151612ee6612a7e826129ff565b818152846020838601011115612efb57600080fd5b612f0c826020830160208701612580565b949350505050565b80516125468161272d565b600060c08284031215612f3157600080fd5b60405160c081016001600160401b038282108183111715612f5457612f546129b9565b8160405282935084519150612f688261269c565b90825260208401519080821115612f7e57600080fd5b50612f8b85828601612ec7565b6020830152506040830151612f9f8161269c565b60408201526060830151612fb28161269c565b6060820152612fc360808401612f14565b608082015260a083015160a08201525092915050565b60006020808385031215612fec57600080fd5b82516001600160401b038082111561300357600080fd5b818501915085601f83011261301757600080fd5b8151613025612a7e82612ea4565b81815260059190911b8301840190848101908883111561304457600080fd5b8585015b8381101561307c578051858111156130605760008081fd5b61306e8b89838a0101612f1f565b845250918601918601613048565b5098975050505050505050565b6001600160401b039390931683526001600160a01b03918216602084015216604082015260600190565b6000602082840312156130c557600080fd5b8151610c408161269c565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60408152600061310d6040830185876130d0565b905060018060a01b0383166020830152949350505050565b60006020828403121561313757600080fd5b8151610c4081612c54565b6040815260006131566040830186886130d0565b82810360208401526131698185876130d0565b979650505050505050565b6001600160a01b038581168252841660208201526060604082018190526000906131a190830184866130d0565b9695505050505050565b6000602082840312156131bd57600080fd5b5051919050565b600080858511156131d457600080fd5b838611156131e157600080fd5b5050820193919092039150565b6001600160e01b031981358181169160048510156132165780818660040360031b1b83161692505b505092915050565b6001600160a01b0393841681529190921660208201526001600160e01b0319909116604082015260600190565b60008060006060848603121561326057600080fd5b835161326b81612c54565b602085015190935061327c8161272d565b60408501519092506126f181612c54565b634e487b7160e01b600052601160045260246000fd5b80820180821115610a3957610a3961328d565b8183823760009101908152919050565b81810381811115610a3957610a3961328d565b6000602082840312156132eb57600080fd5b8151610c4081612d47565b602081526000610c3d602083018486612dff565b602081526000610c3d6020830184866130d0565b60006020828403121561333057600080fd5b81516001600160401b0381111561334657600080fd5b612f0c84828501612f1f565b60006020828403121561336457600080fd5b8151610c4081612453565b6001600160a01b0386168152606060208083018290526000916133959084018789612db7565b604084820360408601528186835283830190506005848860051b8501018960005b8a81101561349757868303601f190185528135368d9003603e190181126133dc57600080fd5b8c01803536829003601e19018082126133f457600080fd5b9082018a81019190356001600160401b038082111561341257600080fd5b81360384131561342157600080fd5b8a88526134318b890183866130d0565b93508c850135915082821261344557600080fd5b9381018c810194903592508083111561345d57600080fd5b505080871b360383131561347057600080fd5b8582038b870152613482828285612db7565b978b01979550505091880191506001016133b6565b50909d9c50505050505050505050505050565b600080604083850312156134bd57600080fd5b82516134c881612c54565b60208401519092506129548161272d565b6001600160401b0384168152604060208201526000610cc86040830184866130d0565b6000602080838503121561350f57600080fd5b82516001600160401b0381111561352557600080fd5b8301601f8101851361353657600080fd5b8051613544612a7e82612ea4565b81815260059190911b8201830190838101908783111561356357600080fd5b928401925b8284101561316957835161357b81612453565b82529284019290840190613568565b6001600160401b039290921682526001600160a01b0316602082015260400190565b6001600160a01b038481168252831660208083019190915261018082019060408301908460005b600a8110156135fa5781516001600160e01b031916845292820192908201906001016135d3565b50505050949350505050565b60006020828403121561361857600080fd5b81516001600160401b0381111561362e57600080fd5b612f0c84828501612ec7565b6000602080838503121561364d57600080fd5b82516001600160401b0381111561366357600080fd5b8301601f8101851361367457600080fd5b8051613682612a7e82612ea4565b81815260059190911b820183019083810190878311156136a157600080fd5b928401925b828410156131695783516136b98161269c565b825292840192908401906136a6565b6000602082840312156136da57600080fd5b8151610c408161272d565b6001600160a01b038516815260606020820181905260009061370a90830185876130d0565b905065ffffffffffff8316604083015295945050505050565b60008251613735818460208701612580565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122018bca46518b3e43f4f61c16f0bde38db078aa1e0bdfb79599448e17151a69cdd64736f6c63430008180033",
	Deps: []*bind.MetaData{
		&AccessManagerAuthLibMetaData,
		&AccessManagerContractScopedLibMetaData,
		&AccessManagerEnumerationLibMetaData,
		&AccessManagerRoleConfigLibMetaData,
		&AccessManagerScheduleLibMetaData,
	},
}

// RaylsAccessManagerV1 is an auto generated Go binding around an Ethereum contract.
type RaylsAccessManagerV1 struct {
	abi abi.ABI
}

// NewRaylsAccessManagerV1 creates a new instance of RaylsAccessManagerV1.
func NewRaylsAccessManagerV1() *RaylsAccessManagerV1 {
	parsed, err := RaylsAccessManagerV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RaylsAccessManagerV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RaylsAccessManagerV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackADMIN is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackADMIN() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("ADMIN")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackADMIN is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackADMIN(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("ADMIN", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackEXPIRATION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb4b5734.
//
// Solidity: function EXPIRATION() view returns(uint48)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackEXPIRATION() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("EXPIRATION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEXPIRATION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbb4b5734.
//
// Solidity: function EXPIRATION() view returns(uint48)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackEXPIRATION(data []byte) (*big.Int, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("EXPIRATION", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPUBLIC is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1179a7f.
//
// Solidity: function PUBLIC() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackPUBLIC() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("PUBLIC")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPUBLIC is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc1179a7f.
//
// Solidity: function PUBLIC() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackPUBLIC(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("PUBLIC", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackTOKENOWNER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xff6c33e6.
//
// Solidity: function TOKEN_OWNER() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackTOKENOWNER() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("TOKEN_OWNER")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTOKENOWNER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xff6c33e6.
//
// Solidity: function TOKEN_OWNER() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackTOKENOWNER(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("TOKEN_OWNER", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAddFunctionAllowedRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5cdd11a1.
//
// Solidity: function addFunctionAllowedRoles(address managedContract, bytes4[] selectors, uint64[] roleIds) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackAddFunctionAllowedRoles(managedContract common.Address, selectors [][4]byte, roleIds []uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("addFunctionAllowedRoles", managedContract, selectors, roleIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAuthority is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackAuthority() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("authority", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackCanCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7009613.
//
// Solidity: function canCall(address caller, address managedContract, bytes4 selector) view returns(bool allowed, uint32 delay, bool paused)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackCanCall(caller common.Address, managedContract common.Address, selector [4]byte) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("canCall", caller, managedContract, selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// CanCallOutput serves as a container for the return parameters of contract
// method CanCall.
type CanCallOutput struct {
	Allowed bool
	Delay   uint32
	Paused  bool
}

// UnpackCanCall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb7009613.
//
// Solidity: function canCall(address caller, address managedContract, bytes4 selector) view returns(bool allowed, uint32 delay, bool paused)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackCanCall(data []byte) (CanCallOutput, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("canCall", data)
	outstruct := new(CanCallOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Allowed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Delay = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Paused = *abi.ConvertType(out[2], new(bool)).(*bool)
	return *outstruct, err

}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd6bb62c6.
//
// Solidity: function cancel(address caller, address managedContract, bytes data) returns(uint32)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackCancel(caller common.Address, managedContract common.Address, data []byte) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("cancel", caller, managedContract, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCancel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd6bb62c6.
//
// Solidity: function cancel(address caller, address managedContract, bytes data) returns(uint32)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackCancel(data []byte) (uint32, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("cancel", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, err
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1cff79cd.
//
// Solidity: function execute(address managedContract, bytes data) payable returns(uint32 delay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackExecute(managedContract common.Address, data []byte) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("execute", managedContract, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecute is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1cff79cd.
//
// Solidity: function execute(address managedContract, bytes data) payable returns(uint32 delay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackExecute(data []byte) (uint32, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("execute", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, err
}

// PackGetAccountContractScopedRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf11545ec.
//
// Solidity: function getAccountContractScopedRoles(address account, address managedContract) view returns(uint64[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetAccountContractScopedRoles(account common.Address, managedContract common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getAccountContractScopedRoles", account, managedContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAccountContractScopedRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf11545ec.
//
// Solidity: function getAccountContractScopedRoles(address account, address managedContract) view returns(uint64[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetAccountContractScopedRoles(data []byte) ([]uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getAccountContractScopedRoles", data)
	if err != nil {
		return *new([]uint64), err
	}
	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	return out0, err
}

// PackGetAccountRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcacebadc.
//
// Solidity: function getAccountRoles(address account) view returns(uint64[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetAccountRoles(account common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getAccountRoles", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAccountRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcacebadc.
//
// Solidity: function getAccountRoles(address account) view returns(uint64[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetAccountRoles(data []byte) ([]uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getAccountRoles", data)
	if err != nil {
		return *new([]uint64), err
	}
	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	return out0, err
}

// PackGetAccountRolesWithInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3497837.
//
// Solidity: function getAccountRolesWithInfo(address account) view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetAccountRolesWithInfo(account common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getAccountRolesWithInfo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAccountRolesWithInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc3497837.
//
// Solidity: function getAccountRolesWithInfo(address account) view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetAccountRolesWithInfo(data []byte) ([]IRaylsAccessManagerRoleInfo, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getAccountRolesWithInfo", data)
	if err != nil {
		return *new([]IRaylsAccessManagerRoleInfo), err
	}
	out0 := *abi.ConvertType(out[0], new([]IRaylsAccessManagerRoleInfo)).(*[]IRaylsAccessManagerRoleInfo)
	return out0, err
}

// PackGetAllRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2bcac3d.
//
// Solidity: function getAllRoles() view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetAllRoles() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getAllRoles")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAllRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2bcac3d.
//
// Solidity: function getAllRoles() view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetAllRoles(data []byte) ([]IRaylsAccessManagerRoleInfo, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getAllRoles", data)
	if err != nil {
		return *new([]IRaylsAccessManagerRoleInfo), err
	}
	out0 := *abi.ConvertType(out[0], new([]IRaylsAccessManagerRoleInfo)).(*[]IRaylsAccessManagerRoleInfo)
	return out0, err
}

// PackGetContractAuthority is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65e7d5c7.
//
// Solidity: function getContractAuthority(address managedContract) view returns(address)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetContractAuthority(managedContract common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getContractAuthority", managedContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetContractAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x65e7d5c7.
//
// Solidity: function getContractAuthority(address managedContract) view returns(address)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetContractAuthority(data []byte) (common.Address, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getContractAuthority", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetContractScopedRoleMembers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf222bbfa.
//
// Solidity: function getContractScopedRoleMembers(uint64 roleId, address managedContract) view returns(address[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetContractScopedRoleMembers(roleId uint64, managedContract common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getContractScopedRoleMembers", roleId, managedContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetContractScopedRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf222bbfa.
//
// Solidity: function getContractScopedRoleMembers(uint64 roleId, address managedContract) view returns(address[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetContractScopedRoleMembers(data []byte) ([]common.Address, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getContractScopedRoleMembers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGetFunctionAllowedRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5c4ca00.
//
// Solidity: function getFunctionAllowedRoles(address managedContract, bytes4 selector) view returns(uint64[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetFunctionAllowedRoles(managedContract common.Address, selector [4]byte) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getFunctionAllowedRoles", managedContract, selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFunctionAllowedRoles is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5c4ca00.
//
// Solidity: function getFunctionAllowedRoles(address managedContract, bytes4 selector) view returns(uint64[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetFunctionAllowedRoles(data []byte) ([]uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getFunctionAllowedRoles", data)
	if err != nil {
		return *new([]uint64), err
	}
	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	return out0, err
}

// PackGetFunctionAllowedRolesWithInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0649fea4.
//
// Solidity: function getFunctionAllowedRolesWithInfo(address managedContract, bytes4 selector) view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetFunctionAllowedRolesWithInfo(managedContract common.Address, selector [4]byte) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getFunctionAllowedRolesWithInfo", managedContract, selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFunctionAllowedRolesWithInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0649fea4.
//
// Solidity: function getFunctionAllowedRolesWithInfo(address managedContract, bytes4 selector) view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetFunctionAllowedRolesWithInfo(data []byte) ([]IRaylsAccessManagerRoleInfo, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getFunctionAllowedRolesWithInfo", data)
	if err != nil {
		return *new([]IRaylsAccessManagerRoleInfo), err
	}
	out0 := *abi.ConvertType(out[0], new([]IRaylsAccessManagerRoleInfo)).(*[]IRaylsAccessManagerRoleInfo)
	return out0, err
}

// PackGetRegisteredRoleCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5050c1f.
//
// Solidity: function getRegisteredRoleCount() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRegisteredRoleCount() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRegisteredRoleCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRegisteredRoleCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc5050c1f.
//
// Solidity: function getRegisteredRoleCount() view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRegisteredRoleCount(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRegisteredRoleCount", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x530dd456.
//
// Solidity: function getRoleAdmin(uint64 roleId) view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRoleAdmin(roleId uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRoleAdmin", roleId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x530dd456.
//
// Solidity: function getRoleAdmin(uint64 roleId) view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRoleAdmin(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackGetRoleGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b0a93ba.
//
// Solidity: function getRoleGuardian(uint64 roleId) view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRoleGuardian(roleId uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRoleGuardian", roleId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleGuardian is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0b0a93ba.
//
// Solidity: function getRoleGuardian(uint64 roleId) view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRoleGuardian(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRoleGuardian", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackGetRoleIdByName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf317a850.
//
// Solidity: function getRoleIdByName(string name) view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRoleIdByName(name string) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRoleIdByName", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleIdByName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf317a850.
//
// Solidity: function getRoleIdByName(string name) view returns(uint64)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRoleIdByName(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRoleIdByName", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackGetRoleInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5bfa5027.
//
// Solidity: function getRoleInfo(uint64 roleId) view returns((uint64,string,uint64,uint64,uint32,uint256))
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRoleInfo(roleId uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRoleInfo", roleId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5bfa5027.
//
// Solidity: function getRoleInfo(uint64 roleId) view returns((uint64,string,uint64,uint64,uint32,uint256))
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRoleInfo(data []byte) (IRaylsAccessManagerRoleInfo, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRoleInfo", data)
	if err != nil {
		return *new(IRaylsAccessManagerRoleInfo), err
	}
	out0 := *abi.ConvertType(out[0], new(IRaylsAccessManagerRoleInfo)).(*IRaylsAccessManagerRoleInfo)
	return out0, err
}

// PackGetRoleInfoBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bd37148.
//
// Solidity: function getRoleInfoBatch(uint64[] roleIds) view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRoleInfoBatch(roleIds []uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRoleInfoBatch", roleIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleInfoBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3bd37148.
//
// Solidity: function getRoleInfoBatch(uint64[] roleIds) view returns((uint64,string,uint64,uint64,uint32,uint256)[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRoleInfoBatch(data []byte) ([]IRaylsAccessManagerRoleInfo, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRoleInfoBatch", data)
	if err != nil {
		return *new([]IRaylsAccessManagerRoleInfo), err
	}
	out0 := *abi.ConvertType(out[0], new([]IRaylsAccessManagerRoleInfo)).(*[]IRaylsAccessManagerRoleInfo)
	return out0, err
}

// PackGetRoleLabel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc554ae82.
//
// Solidity: function getRoleLabel(uint64 roleId) view returns(string)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRoleLabel(roleId uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRoleLabel", roleId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleLabel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc554ae82.
//
// Solidity: function getRoleLabel(uint64 roleId) view returns(string)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRoleLabel(data []byte) (string, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRoleLabel", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackGetRoleMembers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5808e2f.
//
// Solidity: function getRoleMembers(uint64 roleId) view returns(address[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetRoleMembers(roleId uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getRoleMembers", roleId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa5808e2f.
//
// Solidity: function getRoleMembers(uint64 roleId) view returns(address[])
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getRoleMembers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGetSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3adc277a.
//
// Solidity: function getSchedule(bytes32 operationId) view returns(uint48 when)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGetSchedule(operationId [32]byte) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("getSchedule", operationId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3adc277a.
//
// Solidity: function getSchedule(bytes32 operationId) view returns(uint48 when)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackGetSchedule(data []byte) (*big.Int, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("getSchedule", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGrantContractScopedRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12c1e072.
//
// Solidity: function grantContractScopedRole(uint64 roleId, address account, address managedContract, uint32 executionDelay) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGrantContractScopedRole(roleId uint64, account common.Address, managedContract common.Address, executionDelay uint32) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("grantContractScopedRole", roleId, account, managedContract, executionDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25c471a0.
//
// Solidity: function grantRole(uint64 roleId, address account, uint32 executionDelay) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGrantRole(roleId uint64, account common.Address, executionDelay uint32) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("grantRole", roleId, account, executionDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGrantSelfTokenOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e80a70b.
//
// Solidity: function grantSelfTokenOwner(address account) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackGrantSelfTokenOwner(account common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("grantSelfTokenOwner", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasContractScopedRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x804af8c5.
//
// Solidity: function hasContractScopedRole(uint64 roleId, address account, address managedContract) view returns(bool isMember, uint32 executionDelay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackHasContractScopedRole(roleId uint64, account common.Address, managedContract common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("hasContractScopedRole", roleId, account, managedContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// HasContractScopedRoleOutput serves as a container for the return parameters of contract
// method HasContractScopedRole.
type HasContractScopedRoleOutput struct {
	IsMember       bool
	ExecutionDelay uint32
}

// UnpackHasContractScopedRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x804af8c5.
//
// Solidity: function hasContractScopedRole(uint64 roleId, address account, address managedContract) view returns(bool isMember, uint32 executionDelay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackHasContractScopedRole(data []byte) (HasContractScopedRoleOutput, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("hasContractScopedRole", data)
	outstruct := new(HasContractScopedRoleOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.IsMember = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ExecutionDelay = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	return *outstruct, err

}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1f856ee.
//
// Solidity: function hasRole(uint64 roleId, address account) view returns(bool isMember, uint32 executionDelay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackHasRole(roleId uint64, account common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("hasRole", roleId, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// HasRoleOutput serves as a container for the return parameters of contract
// method HasRole.
type HasRoleOutput struct {
	IsMember       bool
	ExecutionDelay uint32
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd1f856ee.
//
// Solidity: function hasRole(uint64 roleId, address account) view returns(bool isMember, uint32 executionDelay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackHasRole(data []byte) (HasRoleOutput, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("hasRole", data)
	outstruct := new(HasRoleOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.IsMember = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ExecutionDelay = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	return *outstruct, err

}

// PackHasRoleByName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x180cb42b.
//
// Solidity: function hasRoleByName(string name, address account) view returns(bool)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackHasRoleByName(name string, account common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("hasRoleByName", name, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRoleByName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x180cb42b.
//
// Solidity: function hasRoleByName(string name, address account) view returns(bool)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackHasRoleByName(data []byte) (bool, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("hasRoleByName", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address initialAdmin) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackInitialize(initialAdmin common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("initialize", initialAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsContractPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0f008eb.
//
// Solidity: function isContractPaused(address managedContract) view returns(bool)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackIsContractPaused(managedContract common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("isContractPaused", managedContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsContractPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb0f008eb.
//
// Solidity: function isContractPaused(address managedContract) view returns(bool)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackIsContractPaused(data []byte) (bool, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("isContractPaused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackLabelRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x853551b8.
//
// Solidity: function labelRole(uint64 roleId, string label) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackLabelRole(roleId uint64, label string) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("labelRole", roleId, label)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackProxiableUUID() []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRegisterRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57087597.
//
// Solidity: function registerRole(string name) returns(uint64 roleId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackRegisterRole(name string) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("registerRole", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRegisterRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57087597.
//
// Solidity: function registerRole(string name) returns(uint64 roleId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRegisterRole(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("registerRole", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackRegisterRoleAndLabel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b9bfca9.
//
// Solidity: function registerRoleAndLabel(string name, string label) returns(uint64 roleId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackRegisterRoleAndLabel(name string, label string) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("registerRoleAndLabel", name, label)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRegisterRoleAndLabel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1b9bfca9.
//
// Solidity: function registerRoleAndLabel(string name, string label) returns(uint64 roleId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRegisterRoleAndLabel(data []byte) (uint64, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("registerRoleAndLabel", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackRemoveFunctionAllowedRoles is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0140c521.
//
// Solidity: function removeFunctionAllowedRoles(address managedContract, bytes4[] selectors, uint64[] roleIds) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackRemoveFunctionAllowedRoles(managedContract common.Address, selectors [][4]byte, roleIds []uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("removeFunctionAllowedRoles", managedContract, selectors, roleIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe0776f5.
//
// Solidity: function renounceRole(uint64 roleId, address callerConfirmation) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackRenounceRole(roleId uint64, callerConfirmation common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("renounceRole", roleId, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeContractScopedRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08c47967.
//
// Solidity: function revokeContractScopedRole(uint64 roleId, address account, address managedContract) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackRevokeContractScopedRole(roleId uint64, account common.Address, managedContract common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("revokeContractScopedRole", roleId, account, managedContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7d2b162.
//
// Solidity: function revokeRole(uint64 roleId, address account) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackRevokeRole(roleId uint64, account common.Address) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("revokeRole", roleId, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf801a698.
//
// Solidity: function schedule(address managedContract, bytes data, uint48 when) returns(bytes32 operationId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackSchedule(managedContract common.Address, data []byte, when *big.Int) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("schedule", managedContract, data, when)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf801a698.
//
// Solidity: function schedule(address managedContract, bytes data, uint48 when) returns(bytes32 operationId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackSchedule(data []byte) ([32]byte, error) {
	out, err := raylsAccessManagerV1.abi.Unpack("schedule", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackSelfRegisterManagedContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f5e800e.
//
// Solidity: function selfRegisterManagedContract(address deployer, bytes4[] ownerSelectors, (string,bytes4[])[] roleMappings) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackSelfRegisterManagedContract(deployer common.Address, ownerSelectors [][4]byte, roleMappings []IRaylsAccessManagerSelectorRoleMapping) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("selfRegisterManagedContract", deployer, ownerSelectors, roleMappings)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetContractPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba126a58.
//
// Solidity: function setContractPaused(address managedContract, bool paused) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackSetContractPaused(managedContract common.Address, paused bool) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("setContractPaused", managedContract, paused)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetGrantDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa64d95ce.
//
// Solidity: function setGrantDelay(uint64 roleId, uint32 newDelay) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackSetGrantDelay(roleId uint64, newDelay uint32) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("setGrantDelay", roleId, newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x30cae187.
//
// Solidity: function setRoleAdmin(uint64 roleId, uint64 adminRoleId) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackSetRoleAdmin(roleId uint64, adminRoleId uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("setRoleAdmin", roleId, adminRoleId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRoleGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52962952.
//
// Solidity: function setRoleGuardian(uint64 roleId, uint64 guardianRoleId) returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackSetRoleGuardian(roleId uint64, guardianRoleId uint64) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("setRoleGuardian", roleId, guardianRoleId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := raylsAccessManagerV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// RaylsAccessManagerV1AuthorityUpdated represents a AuthorityUpdated event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1AuthorityUpdated) ContractEventName() string {
	return RaylsAccessManagerV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*RaylsAccessManagerV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1ContractPauseUpdated represents a ContractPauseUpdated event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1ContractPauseUpdated struct {
	ManagedContract common.Address
	Paused          bool
	Raw             *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1ContractPauseUpdatedEventName = "ContractPauseUpdated"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1ContractPauseUpdated) ContractEventName() string {
	return RaylsAccessManagerV1ContractPauseUpdatedEventName
}

// UnpackContractPauseUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ContractPauseUpdated(address indexed managedContract, bool paused)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackContractPauseUpdatedEvent(log *types.Log) (*RaylsAccessManagerV1ContractPauseUpdated, error) {
	event := "ContractPauseUpdated"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1ContractPauseUpdated)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1ContractScopedRoleGranted represents a ContractScopedRoleGranted event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1ContractScopedRoleGranted struct {
	RoleId          uint64
	Account         common.Address
	ManagedContract common.Address
	ExecutionDelay  uint32
	ActiveSince     *big.Int
	Grantor         common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1ContractScopedRoleGrantedEventName = "ContractScopedRoleGranted"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1ContractScopedRoleGranted) ContractEventName() string {
	return RaylsAccessManagerV1ContractScopedRoleGrantedEventName
}

// UnpackContractScopedRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ContractScopedRoleGranted(uint64 indexed roleId, address indexed account, address indexed managedContract, uint32 executionDelay, uint48 activeSince, address grantor)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackContractScopedRoleGrantedEvent(log *types.Log) (*RaylsAccessManagerV1ContractScopedRoleGranted, error) {
	event := "ContractScopedRoleGranted"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1ContractScopedRoleGranted)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1ContractScopedRoleRevoked represents a ContractScopedRoleRevoked event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1ContractScopedRoleRevoked struct {
	RoleId          uint64
	Account         common.Address
	ManagedContract common.Address
	Revoker         common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1ContractScopedRoleRevokedEventName = "ContractScopedRoleRevoked"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1ContractScopedRoleRevoked) ContractEventName() string {
	return RaylsAccessManagerV1ContractScopedRoleRevokedEventName
}

// UnpackContractScopedRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ContractScopedRoleRevoked(uint64 indexed roleId, address indexed account, address indexed managedContract, address revoker)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackContractScopedRoleRevokedEvent(log *types.Log) (*RaylsAccessManagerV1ContractScopedRoleRevoked, error) {
	event := "ContractScopedRoleRevoked"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1ContractScopedRoleRevoked)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1FunctionAllowedRoleAdded represents a FunctionAllowedRoleAdded event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1FunctionAllowedRoleAdded struct {
	ManagedContract common.Address
	Selector        [4]byte
	RoleId          uint64
	Raw             *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1FunctionAllowedRoleAddedEventName = "FunctionAllowedRoleAdded"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1FunctionAllowedRoleAdded) ContractEventName() string {
	return RaylsAccessManagerV1FunctionAllowedRoleAddedEventName
}

// UnpackFunctionAllowedRoleAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FunctionAllowedRoleAdded(address indexed managedContract, bytes4 indexed selector, uint64 indexed roleId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackFunctionAllowedRoleAddedEvent(log *types.Log) (*RaylsAccessManagerV1FunctionAllowedRoleAdded, error) {
	event := "FunctionAllowedRoleAdded"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1FunctionAllowedRoleAdded)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1FunctionAllowedRoleRemoved represents a FunctionAllowedRoleRemoved event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1FunctionAllowedRoleRemoved struct {
	ManagedContract common.Address
	Selector        [4]byte
	RoleId          uint64
	Raw             *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1FunctionAllowedRoleRemovedEventName = "FunctionAllowedRoleRemoved"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1FunctionAllowedRoleRemoved) ContractEventName() string {
	return RaylsAccessManagerV1FunctionAllowedRoleRemovedEventName
}

// UnpackFunctionAllowedRoleRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FunctionAllowedRoleRemoved(address indexed managedContract, bytes4 indexed selector, uint64 indexed roleId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackFunctionAllowedRoleRemovedEvent(log *types.Log) (*RaylsAccessManagerV1FunctionAllowedRoleRemoved, error) {
	event := "FunctionAllowedRoleRemoved"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1FunctionAllowedRoleRemoved)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1Initialized represents a Initialized event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1Initialized) ContractEventName() string {
	return RaylsAccessManagerV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackInitializedEvent(log *types.Log) (*RaylsAccessManagerV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1Initialized)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1ManagedContractRegistered represents a ManagedContractRegistered event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1ManagedContractRegistered struct {
	ManagedContract   common.Address
	ContractAuthority common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1ManagedContractRegisteredEventName = "ManagedContractRegistered"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1ManagedContractRegistered) ContractEventName() string {
	return RaylsAccessManagerV1ManagedContractRegisteredEventName
}

// UnpackManagedContractRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ManagedContractRegistered(address indexed managedContract, address indexed contractAuthority)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackManagedContractRegisteredEvent(log *types.Log) (*RaylsAccessManagerV1ManagedContractRegistered, error) {
	event := "ManagedContractRegistered"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1ManagedContractRegistered)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1OperationCanceled represents a OperationCanceled event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1OperationCanceled struct {
	OperationId [32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1OperationCanceledEventName = "OperationCanceled"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1OperationCanceled) ContractEventName() string {
	return RaylsAccessManagerV1OperationCanceledEventName
}

// UnpackOperationCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperationCanceled(bytes32 indexed operationId)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackOperationCanceledEvent(log *types.Log) (*RaylsAccessManagerV1OperationCanceled, error) {
	event := "OperationCanceled"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1OperationCanceled)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1OperationExecuted represents a OperationExecuted event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1OperationExecuted struct {
	OperationId     [32]byte
	ManagedContract common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1OperationExecutedEventName = "OperationExecuted"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1OperationExecuted) ContractEventName() string {
	return RaylsAccessManagerV1OperationExecutedEventName
}

// UnpackOperationExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperationExecuted(bytes32 indexed operationId, address indexed managedContract)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackOperationExecutedEvent(log *types.Log) (*RaylsAccessManagerV1OperationExecuted, error) {
	event := "OperationExecuted"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1OperationExecuted)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1OperationScheduled represents a OperationScheduled event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1OperationScheduled struct {
	OperationId     [32]byte
	Caller          common.Address
	ManagedContract common.Address
	ExecuteAfter    *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1OperationScheduledEventName = "OperationScheduled"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1OperationScheduled) ContractEventName() string {
	return RaylsAccessManagerV1OperationScheduledEventName
}

// UnpackOperationScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperationScheduled(bytes32 indexed operationId, address indexed caller, address indexed managedContract, uint48 executeAfter)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackOperationScheduledEvent(log *types.Log) (*RaylsAccessManagerV1OperationScheduled, error) {
	event := "OperationScheduled"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1OperationScheduled)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1RoleAdminChanged represents a RoleAdminChanged event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RoleAdminChanged struct {
	RoleId   uint64
	NewAdmin uint64
	Raw      *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1RoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1RoleAdminChanged) ContractEventName() string {
	return RaylsAccessManagerV1RoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(uint64 indexed roleId, uint64 indexed newAdmin)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRoleAdminChangedEvent(log *types.Log) (*RaylsAccessManagerV1RoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1RoleAdminChanged)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1RoleGrantDelayChanged represents a RoleGrantDelayChanged event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RoleGrantDelayChanged struct {
	RoleId   uint64
	NewDelay uint32
	Raw      *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1RoleGrantDelayChangedEventName = "RoleGrantDelayChanged"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1RoleGrantDelayChanged) ContractEventName() string {
	return RaylsAccessManagerV1RoleGrantDelayChangedEventName
}

// UnpackRoleGrantDelayChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGrantDelayChanged(uint64 indexed roleId, uint32 newDelay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRoleGrantDelayChangedEvent(log *types.Log) (*RaylsAccessManagerV1RoleGrantDelayChanged, error) {
	event := "RoleGrantDelayChanged"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1RoleGrantDelayChanged)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1RoleGranted represents a RoleGranted event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RoleGranted struct {
	RoleId         uint64
	Account        common.Address
	ExecutionDelay uint32
	ActiveSince    *big.Int
	Grantor        common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1RoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1RoleGranted) ContractEventName() string {
	return RaylsAccessManagerV1RoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(uint64 indexed roleId, address indexed account, uint32 executionDelay, uint48 activeSince, address indexed grantor)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRoleGrantedEvent(log *types.Log) (*RaylsAccessManagerV1RoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1RoleGranted)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1RoleGuardianChanged represents a RoleGuardianChanged event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RoleGuardianChanged struct {
	RoleId      uint64
	NewGuardian uint64
	Raw         *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1RoleGuardianChangedEventName = "RoleGuardianChanged"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1RoleGuardianChanged) ContractEventName() string {
	return RaylsAccessManagerV1RoleGuardianChangedEventName
}

// UnpackRoleGuardianChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGuardianChanged(uint64 indexed roleId, uint64 indexed newGuardian)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRoleGuardianChangedEvent(log *types.Log) (*RaylsAccessManagerV1RoleGuardianChanged, error) {
	event := "RoleGuardianChanged"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1RoleGuardianChanged)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1RoleLabelSet represents a RoleLabelSet event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RoleLabelSet struct {
	RoleId uint64
	Label  string
	Raw    *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1RoleLabelSetEventName = "RoleLabelSet"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1RoleLabelSet) ContractEventName() string {
	return RaylsAccessManagerV1RoleLabelSetEventName
}

// UnpackRoleLabelSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleLabelSet(uint64 indexed roleId, string label)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRoleLabelSetEvent(log *types.Log) (*RaylsAccessManagerV1RoleLabelSet, error) {
	event := "RoleLabelSet"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1RoleLabelSet)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1RoleRegistered represents a RoleRegistered event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RoleRegistered struct {
	RoleId uint64
	Name   string
	Raw    *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1RoleRegisteredEventName = "RoleRegistered"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1RoleRegistered) ContractEventName() string {
	return RaylsAccessManagerV1RoleRegisteredEventName
}

// UnpackRoleRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRegistered(uint64 indexed roleId, string name)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRoleRegisteredEvent(log *types.Log) (*RaylsAccessManagerV1RoleRegistered, error) {
	event := "RoleRegistered"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1RoleRegistered)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1RoleRevoked represents a RoleRevoked event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RoleRevoked struct {
	RoleId  uint64
	Account common.Address
	Revoker common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1RoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1RoleRevoked) ContractEventName() string {
	return RaylsAccessManagerV1RoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(uint64 indexed roleId, address indexed account, address indexed revoker)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRoleRevokedEvent(log *types.Log) (*RaylsAccessManagerV1RoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1RoleRevoked)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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

// RaylsAccessManagerV1Upgraded represents a Upgraded event raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RaylsAccessManagerV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (RaylsAccessManagerV1Upgraded) ContractEventName() string {
	return RaylsAccessManagerV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackUpgradedEvent(log *types.Log) (*RaylsAccessManagerV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != raylsAccessManagerV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RaylsAccessManagerV1Upgraded)
	if len(log.Data) > 0 {
		if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range raylsAccessManagerV1.abi.Events[event].Inputs {
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
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["RaylsAccessManagerV1ContractPaused"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackRaylsAccessManagerV1ContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["RaylsAccessManagerV1Unauthorized"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackRaylsAccessManagerV1UnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], raylsAccessManagerV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return raylsAccessManagerV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// RaylsAccessManagerV1AddressEmptyCode represents a AddressEmptyCode error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func RaylsAccessManagerV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackAddressEmptyCodeError(raw []byte) (*RaylsAccessManagerV1AddressEmptyCode, error) {
	out := new(RaylsAccessManagerV1AddressEmptyCode)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func RaylsAccessManagerV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackERC1967InvalidImplementationError(raw []byte) (*RaylsAccessManagerV1ERC1967InvalidImplementation, error) {
	out := new(RaylsAccessManagerV1ERC1967InvalidImplementation)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func RaylsAccessManagerV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackERC1967NonPayableError(raw []byte) (*RaylsAccessManagerV1ERC1967NonPayable, error) {
	out := new(RaylsAccessManagerV1ERC1967NonPayable)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1FailedCall represents a FailedCall error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func RaylsAccessManagerV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackFailedCallError(raw []byte) (*RaylsAccessManagerV1FailedCall, error) {
	out := new(RaylsAccessManagerV1FailedCall)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1InvalidInitialization represents a InvalidInitialization error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func RaylsAccessManagerV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackInvalidInitializationError(raw []byte) (*RaylsAccessManagerV1InvalidInitialization, error) {
	out := new(RaylsAccessManagerV1InvalidInitialization)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1NotInitializing represents a NotInitializing error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func RaylsAccessManagerV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackNotInitializingError(raw []byte) (*RaylsAccessManagerV1NotInitializing, error) {
	out := new(RaylsAccessManagerV1NotInitializing)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func RaylsAccessManagerV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*RaylsAccessManagerV1RaylsAccessManagedContractPaused, error) {
	out := new(RaylsAccessManagerV1RaylsAccessManagedContractPaused)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func RaylsAccessManagerV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*RaylsAccessManagerV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(RaylsAccessManagerV1RaylsAccessManagedInvalidAuthority)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func RaylsAccessManagerV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*RaylsAccessManagerV1RaylsAccessManagedMustSchedule, error) {
	out := new(RaylsAccessManagerV1RaylsAccessManagedMustSchedule)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func RaylsAccessManagerV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*RaylsAccessManagerV1RaylsAccessManagedUnauthorized, error) {
	out := new(RaylsAccessManagerV1RaylsAccessManagedUnauthorized)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1RaylsAccessManagerV1ContractPaused represents a RaylsAccessManagerV1__ContractPaused error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RaylsAccessManagerV1ContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManagerV1__ContractPaused()
func RaylsAccessManagerV1RaylsAccessManagerV1ContractPausedErrorID() common.Hash {
	return common.HexToHash("0x7b7fabb2b405042262517d916d03be98c20f077febdabd2d88d10192ac117bf7")
}

// UnpackRaylsAccessManagerV1ContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManagerV1__ContractPaused()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRaylsAccessManagerV1ContractPausedError(raw []byte) (*RaylsAccessManagerV1RaylsAccessManagerV1ContractPaused, error) {
	out := new(RaylsAccessManagerV1RaylsAccessManagerV1ContractPaused)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "RaylsAccessManagerV1ContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1RaylsAccessManagerV1Unauthorized represents a RaylsAccessManagerV1__Unauthorized error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1RaylsAccessManagerV1Unauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManagerV1__Unauthorized(address caller)
func RaylsAccessManagerV1RaylsAccessManagerV1UnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xf96b0e65e0e869b92c08051ea45e36c91d6c5e3278c8e4c4aa54e9c935c034f1")
}

// UnpackRaylsAccessManagerV1UnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManagerV1__Unauthorized(address caller)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackRaylsAccessManagerV1UnauthorizedError(raw []byte) (*RaylsAccessManagerV1RaylsAccessManagerV1Unauthorized, error) {
	out := new(RaylsAccessManagerV1RaylsAccessManagerV1Unauthorized)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "RaylsAccessManagerV1Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func RaylsAccessManagerV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*RaylsAccessManagerV1UUPSUnauthorizedCallContext, error) {
	out := new(RaylsAccessManagerV1UUPSUnauthorizedCallContext)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RaylsAccessManagerV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the RaylsAccessManagerV1 contract.
type RaylsAccessManagerV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func RaylsAccessManagerV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (raylsAccessManagerV1 *RaylsAccessManagerV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*RaylsAccessManagerV1UUPSUnsupportedProxiableUUID, error) {
	out := new(RaylsAccessManagerV1UUPSUnsupportedProxiableUUID)
	if err := raylsAccessManagerV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
