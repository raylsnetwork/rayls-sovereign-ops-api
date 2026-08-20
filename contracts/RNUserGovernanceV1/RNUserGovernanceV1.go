// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RNUserGovernanceV1

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

// IUserGovernanceAddressPair is an auto generated low-level Go binding around an user-defined struct.
type IUserGovernanceAddressPair struct {
	PublicAddress  common.Address
	PrivateAddress common.Address
	CreatedAt      *big.Int
	IsActive       bool
	ApprovalStatus uint8
}

// RNUserGovernanceV1MetaData contains all meta data concerning the RNUserGovernanceV1 contract.
var RNUserGovernanceV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addAddressPair\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveUser\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkUserIsApprovedByPrivateAddress\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createUser\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveAddressPairCount\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveAddressPairs\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIUserGovernance.AddressPair[]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressPairApprovalStatus\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressPairsByApprovalStatus\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIUserGovernance.AddressPair[]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllPendingAddressPairs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structIUserGovernance.AddressPair[][]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllUsers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApprovedAddressPairCount\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApprovedAddressPairs\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIUserGovernance.AddressPair[]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAddressPairCount\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAddressPairs\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIUserGovernance.AddressPair[]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrivateAddressFromPublic\",\"inputs\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublicAddressFromPrivate\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRejectedAddressPairs\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIUserGovernance.AddressPair[]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserAddressPairCount\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserAddressPairs\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIUserGovernance.AddressPair[]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserIdByPrivateAddress\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserIdByPublicAddress\",\"inputs\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasUser\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAddressPairActive\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAddressPairApproved\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPrivateAddressMapped\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPublicAddressMapped\",\"inputs\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"privateAddressToUserId\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicAddressToUserId\",\"inputs\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rejectUser\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAddressPair\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeUser\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddressPairApprovalStatus\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"userAddressPairs\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIUserGovernance.AddressPair[]\",\"components\":[{\"name\":\"publicAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userExists\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddressPairAdded\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressPairApprovalChanged\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIUserGovernance.ApprovalStatus\"},{\"name\":\"newStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIUserGovernance.ApprovalStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressPairRemoved\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"publicAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"privateAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserCreated\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserRemoved\",\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__AddressPairByPrivateKeyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__AddressPairNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__InvalidPrivateAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__InvalidPublicAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__InvalidUserId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__PrivateAddressAlreadyMapped\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__PrivateAddressNotMapped\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__PrivateAddressNotMappedToUser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__PublicAddressAlreadyMapped\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__PublicAddressNotMapped\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__PublicAddressNotMappedToUser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__UserAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__UserDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNUserGovernanceV1__UserHasNoAddressPairs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__MustSchedule\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:  "RNUserGovernanceV1",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051613f77620001046000396000818161342801528181613451015261358a0152613f776000f3fe6080604052600436106102065760003560e01c80639c69903811610119578063c4ba03ed116100a6578063c4ba03ed146105eb578063c4d66de81461060b578063c92c94811461062b578063caa60d4c1461064b578063d737cd9e1461066e578063dd77a8ae1461068e578063e2842d79146106ae578063e5080035146106d0578063f1c288cb1461032e578063fbc56938146106f0578063fccb75831461068e57600080fd5b80639c699038146104c3578063a1eb2f70146104e3578063a2e8452c14610463578063abf7bfd814610503578063ad3cb1cc14610523578063b26b32c214610561578063b4984fd714610581578063b5cb15f7146105a1578063b97f100f146105b6578063bf7e214f146105d657600080fd5b80635a9054d8116101975780635a9054d8146103765780635d3070b114610396578063741827de146103b6578063748ec164146103d657806374ab379f146104035780637be91a72146104235780638ee126a5146104435780639686e8841461046357806399a43bb0146104835780639ad5e2fb146104a357600080fd5b80630488438c1461020b57806306096d601461022d578063166166e31461026357806318325426146102835780632463a1be146102a35780632861a4bb146102d05780634635fd681461030057806347369a1f1461032e5780634f1ef2861461034e57806352d1902d14610361575b600080fd5b34801561021757600080fd5b5061022b6102263660046139eb565b610710565b005b34801561023957600080fd5b5061024d610248366004613a27565b610a18565b60405161025a9190613ac5565b60405180910390f35b34801561026f57600080fd5b5061024d61027e366004613a27565b610b2c565b34801561028f57600080fd5b5061024d61029e366004613a27565b610d69565b3480156102af57600080fd5b506102c36102be366004613b0e565b610d7c565b60405161025a9190613b29565b3480156102dc57600080fd5b506102f06102eb366004613b0e565b610f79565b604051901515815260200161025a565b34801561030c57600080fd5b5061032061031b366004613a27565b610fb4565b60405190815260200161025a565b34801561033a57600080fd5b50610320610349366004613b0e565b610fe4565b61022b61035c366004613b53565b611010565b34801561036d57600080fd5b5061032061102f565b34801561038257600080fd5b5061022b610391366004613a27565b61104c565b3480156103a257600080fd5b5061024d6103b1366004613a27565b611263565b3480156103c257600080fd5b506103206103d1366004613a27565b611270565b3480156103e257600080fd5b506103f66103f13660046139eb565b6113e8565b60405161025a9190613c14565b34801561040f57600080fd5b5061024d61041e366004613a27565b6115ca565b34801561042f57600080fd5b5061032061043e366004613a27565b6116a8565b34801561044f57600080fd5b5061032061045e366004613a27565b611803565b34801561046f57600080fd5b506102f061047e366004613a27565b611853565b34801561048f57600080fd5b5061024d61049e366004613c31565b611874565b3480156104af57600080fd5b506102f06104be3660046139eb565b611af8565b3480156104cf57600080fd5b506103206104de366004613a27565b611cdb565b3480156104ef57600080fd5b506102c36104fe366004613b0e565b611e49565b34801561050f57600080fd5b5061022b61051e366004613a27565b612039565b34801561052f57600080fd5b50610554604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161025a9190613c81565b34801561056d57600080fd5b5061022b61057c366004613cb4565b61210c565b34801561058d57600080fd5b5061022b61059c366004613a27565b6123c8565b3480156105ad57600080fd5b506103206125fa565b3480156105c257600080fd5b506102f06105d13660046139eb565b61260d565b3480156105e257600080fd5b506102c36127ce565b3480156105f757600080fd5b506102f0610606366004613b0e565b6127e7565b34801561061757600080fd5b5061022b610626366004613b0e565b6129a3565b34801561063757600080fd5b5061022b610646366004613a27565b612aa5565b34801561065757600080fd5b50610660612cd0565b60405161025a929190613d3d565b34801561067a57600080fd5b506102f0610689366004613b0e565b612f57565b34801561069a57600080fd5b506103206106a9366004613b0e565b612f92565b3480156106ba57600080fd5b506106c3612fbe565b60405161025a9190613ddb565b3480156106dc57600080fd5b5061022b6106eb3660046139eb565b61301f565b3480156106fc57600080fd5b5061024d61070b366004613a27565b613298565b610726336000356001600160e01b0319166132a5565b60006107306133f9565b600085815260038201602052604090205490915060ff1661076457604051633024159f60e21b815260040160405180910390fd5b6001600160a01b0383166000908152600182016020526040902054841461079e57604051635e26379f60e01b815260040160405180910390fd5b6001600160a01b038216600090815260028201602052604090205484146107d85760405163b623c71360e01b815260040160405180910390fd5b6000848152602082905260408120805490915b818110156109b457856001600160a01b031683828154811061080f5761080f613dee565b60009182526020909120600490910201546001600160a01b031614801561086c5750846001600160a01b031683828154811061084d5761084d613dee565b60009182526020909120600160049092020101546001600160a01b0316145b156109ac578261087d600184613e1a565b8154811061088d5761088d613dee565b90600052602060002090600402018382815481106108ad576108ad613dee565b60009182526020909120825460049092020180546001600160a01b03199081166001600160a01b0393841617825560018085015490830180549092169316929092179091556002808301548183015560038084018054918401805460ff938416151560ff1982168117835592546101009081900490941694919361ff001990931661ffff19909116179190849081111561094957610949613a40565b02179055509050508280548061096157610961613e2d565b60008281526020812060046000199093019283020180546001600160a01b0319908116825560018201805490911690556002810191909155600301805461ffff1916905590556109b4565b6001016107eb565b506001600160a01b038086166000818152600186016020908152604080832083905593881680835260028801909152838220829055925189917f5f25a19db9b6e1520d9ef11946b16680a84d91db16f49e42583e97a81fa6416d91a4505050505050565b60606000610a246133f9565b600084815260038201602052604090205490915060ff16610a5857604051633024159f60e21b815260040160405180910390fd5b60008381526020828152604080832080548251818502810185019093528083529193909284015b82821015610b205760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff8082161515606085015292939260808501926101009092041690811115610afc57610afc613a40565b6002811115610b0d57610b0d613a40565b8152505081526020019060010190610a7f565b50505050915050919050565b60606000610b386133f9565b600084815260038201602052604090205490915060ff16610b6c57604051633024159f60e21b815260040160405180910390fd5b60008381526020828152604080832080548251818502810185019093528083529192909190849084015b82821015610c375760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff8082161515606085015292939260808501926101009092041690811115610c1357610c13613a40565b6002811115610c2457610c24613a40565b8152505081526020019060010190610b96565b5050505090506000815190506000805b82811015610c8857838181518110610c6157610c61613dee565b60200260200101516060015115610c805781610c7c81613e43565b9250505b600101610c47565b506000816001600160401b03811115610ca357610ca3613b3d565b604051908082528060200260200182016040528015610cdc57816020015b610cc9613943565b815260200190600190039081610cc15790505b5090506000805b84811015610d5c57858181518110610cfd57610cfd613dee565b60200260200101516060015115610d5457858181518110610d2057610d20613dee565b6020026020010151838381518110610d3a57610d3a613dee565b60200260200101819052508180610d5090613e43565b9250505b600101610ce3565b5090979650505050505050565b6060610d76826001611874565b92915050565b600080610d876133f9565b6001600160a01b038416600090815260028201602052604090205490915080610dc3576040516307c5e32d60e41b815260040160405180910390fd5b60008181526020838152604080832080548251818502810185019093528083529192909190849084015b82821015610e8e5760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff8082161515606085015292939260808501926101009092041690811115610e6a57610e6a613a40565b6002811115610e7b57610e7b613a40565b8152505081526020019060010190610ded565b5050825192935060009150505b81811015610f6c57866001600160a01b0316838281518110610ebf57610ebf613dee565b6020026020010151602001516001600160a01b0316148015610efb5750828181518110610eee57610eee613dee565b6020026020010151606001515b8015610f3557506001838281518110610f1657610f16613dee565b6020026020010151608001516002811115610f3357610f33613a40565b145b15610f6457828181518110610f4c57610f4c613dee565b60200260200101516000015195505050505050919050565b600101610e9b565b5060009695505050505050565b600080610f846133f9565b6001016000846001600160a01b03166001600160a01b031681526020019081526020016000205414159050919050565b6000610fbe6133f9565b6004018281548110610fd257610fd2613dee565b90600052602060002001549050919050565b6000610fee6133f9565b6001600160a01b03909216600090815260029290920160205250604090205490565b61101861341d565b611021826134ad565b61102b82826134c6565b5050565b600061103961357f565b50600080516020613f0283398151915290565b611062336000356001600160e01b0319166132a5565b600061106c6133f9565b600083815260038201602052604090205490915060ff166110a057604051633024159f60e21b815260040160405180910390fd5b6000828152602082905260408120805490915b8181101561114d578360010160008483815481106110d3576110d3613dee565b600091825260208083206004909202909101546001600160a01b031683528201929092526040018120819055835460028601919085908490811061111957611119613dee565b6000918252602080832060016004909302018201546001600160a01b031684528301939093526040909101812055016110b3565b5060008481526020849052604081206111659161396f565b600483015460005b8181101561121c578585600401828154811061118b5761118b613dee565b90600052602060002001540361121457600485016111aa600184613e1a565b815481106111ba576111ba613dee565b90600052602060002001548560040182815481106111da576111da613dee565b600091825260209091200155600485018054806111f9576111f9613e2d565b6001900381819060005260206000200160009055905561121c565b60010161116d565b506000858152600385016020526040808220805460ff191690555186917fa93796be3e2308b3f9614f44a3675cabfe46b3f799e867359bb1c3bdd7ec1cb091a25050505050565b6060610d76826000611874565b60008061127b6133f9565b600084815260038201602052604090205490915060ff166112af57604051633024159f60e21b815260040160405180910390fd5b60008381526020828152604080832080548251818502810185019093528083529192909190849084015b8282101561137a5760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff808216151560608501529293926080850192610100909204169081111561135657611356613a40565b600281111561136757611367613a40565b81525050815260200190600101906112d9565b5050505090506000815190506000805b828110156113de5760018482815181106113a6576113a6613dee565b60200260200101516080015160028111156113c3576113c3613a40565b036113d657816113d281613e43565b9250505b60010161138a565b5095945050505050565b6000806113f36133f9565b600086815260038201602052604090205490915060ff1661142757604051633024159f60e21b815260040160405180910390fd5b60008581526020828152604080832080548251818502810185019093528083529192909190849084015b828210156114f25760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff80821615156060850152929392608085019261010090920416908111156114ce576114ce613a40565b60028111156114df576114df613a40565b8152505081526020019060010190611451565b5050825192935060009150505b818110156115a957866001600160a01b031683828151811061152357611523613dee565b6020026020010151600001516001600160a01b03161480156115735750856001600160a01b031683828151811061155c5761155c613dee565b6020026020010151602001516001600160a01b0316145b156115a15782818151811061158a5761158a613dee565b6020026020010151608001519450505050506115c3565b6001016114ff565b50604051630ef1d6c160e01b815260040160405180910390fd5b9392505050565b60606115d46133f9565b6000838152602091825260408082208054825181860281018601909352808352919390929084015b8282101561169d5760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff808216151560608501529293926080850192610100909204169081111561167957611679613a40565b600281111561168a5761168a613a40565b81525050815260200190600101906115fc565b505050509050919050565b6000806116b36133f9565b600084815260038201602052604090205490915060ff166116e757604051633024159f60e21b815260040160405180910390fd5b60008381526020828152604080832080548251818502810185019093528083529192909190849084015b828210156117b25760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff808216151560608501529293926080850192610100909204169081111561178e5761178e613a40565b600281111561179f5761179f613a40565b8152505081526020019060010190611711565b5050505090506000815190506000805b828110156113de578381815181106117dc576117dc613dee565b602002602001015160600151156117fb57816117f781613e43565b9250505b6001016117c2565b60008061180e6133f9565b600084815260038201602052604090205490915060ff1661184257604051633024159f60e21b815260040160405180910390fd5b600092835260205250604090205490565b600061185d6133f9565b600092835260030160205250604090205460ff1690565b606060006118806133f9565b600085815260038201602052604090205490915060ff166118b457604051633024159f60e21b815260040160405180910390fd5b60008481526020828152604080832080548251818502810185019093528083529192909190849084015b8282101561197f5760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff808216151560608501529293926080850192610100909204169081111561195b5761195b613a40565b600281111561196c5761196c613a40565b81525050815260200190600101906118de565b5050505090506000815190506000805b828110156119f3578660028111156119a9576119a9613a40565b8482815181106119bb576119bb613dee565b60200260200101516080015160028111156119d8576119d8613a40565b036119eb57816119e781613e43565b9250505b60010161198f565b506000816001600160401b03811115611a0e57611a0e613b3d565b604051908082528060200260200182016040528015611a4757816020015b611a34613943565b815260200190600190039081611a2c5790505b5090506000805b84811015611aea57886002811115611a6857611a68613a40565b868281518110611a7a57611a7a613dee565b6020026020010151608001516002811115611a9757611a97613a40565b03611ae257858181518110611aae57611aae613dee565b6020026020010151838381518110611ac857611ac8613dee565b60200260200101819052508180611ade90613e43565b9250505b600101611a4e565b509098975050505050505050565b600080611b036133f9565b600086815260038201602052604090205490915060ff16611b3757604051633024159f60e21b815260040160405180910390fd5b60008581526020828152604080832080548251818502810185019093528083529192909190849084015b82821015611c025760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff8082161515606085015292939260808501926101009092041690811115611bde57611bde613a40565b6002811115611bef57611bef613a40565b8152505081526020019060010190611b61565b5050825192935060009150505b81811015611ccd57866001600160a01b0316838281518110611c3357611c33613dee565b6020026020010151600001516001600160a01b0316148015611c835750856001600160a01b0316838281518110611c6c57611c6c613dee565b6020026020010151602001516001600160a01b0316145b15611cc5576001838281518110611c9c57611c9c613dee565b6020026020010151608001516002811115611cb957611cb9613a40565b149450505050506115c3565b600101611c0f565b506000979650505050505050565b600080611ce66133f9565b600084815260038201602052604090205490915060ff16611d1a57604051633024159f60e21b815260040160405180910390fd5b60008381526020828152604080832080548251818502810185019093528083529192909190849084015b82821015611de55760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff8082161515606085015292939260808501926101009092041690811115611dc157611dc1613a40565b6002811115611dd257611dd2613a40565b8152505081526020019060010190611d44565b5050505090506000815190506000805b828110156113de576000848281518110611e1157611e11613dee565b6020026020010151608001516002811115611e2e57611e2e613a40565b03611e415781611e3d81613e43565b9250505b600101611df5565b600080611e546133f9565b6001600160a01b038416600090815260018201602052604090205490915080611e9057604051637484912160e01b815260040160405180910390fd5b60008181526020838152604080832080548251818502810185019093528083529192909190849084015b82821015611f5b5760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff8082161515606085015292939260808501926101009092041690811115611f3757611f37613a40565b6002811115611f4857611f48613a40565b8152505081526020019060010190611eba565b5050825192935060009150505b81811015610f6c57866001600160a01b0316838281518110611f8c57611f8c613dee565b6020026020010151600001516001600160a01b0316148015611fc85750828181518110611fbb57611fbb613dee565b6020026020010151606001515b801561200257506001838281518110611fe357611fe3613dee565b602002602001015160800151600281111561200057612000613a40565b145b156120315782818151811061201957612019613dee565b60200260200101516020015195505050505050919050565b600101611f68565b61204f336000356001600160e01b0319166132a5565b60006120596133f9565b90508161207957604051631fb9890760e21b815260040160405180910390fd5b600082815260038201602052604090205460ff16156120ab576040516359e0145760e11b815260040160405180910390fd5b60008281526003820160209081526040808320805460ff19166001908117909155600485018054918201815584529183209091018490555183917f0808387569c40466c01100158dd3d6f79cdbd57e12063e1d81d1cd818d9639cd91a25050565b612122336000356001600160e01b0319166132a5565b600061212c6133f9565b600086815260038201602052604090205490915060ff1661216057604051633024159f60e21b815260040160405180910390fd5b6001600160a01b0384166000908152600182016020526040902054851461219a57604051635e26379f60e01b815260040160405180910390fd5b6001600160a01b038316600090815260028201602052604090205485146121d45760405163b623c71360e01b815260040160405180910390fd5b6000858152602082905260408120805490915b818110156115a957866001600160a01b031683828154811061220b5761220b613dee565b60009182526020909120600490910201546001600160a01b03161480156122685750856001600160a01b031683828154811061224957612249613dee565b60009182526020909120600160049092020101546001600160a01b0316145b156123ba57600083828154811061228157612281613dee565b906000526020600020906004020160030160019054906101000a900460ff169050858483815481106122b5576122b5613dee565b60009182526020909120600360049092020101805461ff0019166101008360028111156122e4576122e4613a40565b021790555060018660028111156122fd576122fd613a40565b0361233d57600184838154811061231657612316613dee565b60009182526020909120600490910201600301805460ff1916911515919091179055612374565b600084838154811061235157612351613dee565b60009182526020909120600490910201600301805460ff19169115159190911790555b866001600160a01b0316886001600160a01b03168a600080516020613f22833981519152848a6040516123a8929190613e5c565b60405180910390a450505050506123c2565b6001016121e7565b50505050565b6123de336000356001600160e01b0319166132a5565b60006123e86133f9565b600083815260038201602052604090205490915060ff1661241c57604051633024159f60e21b815260040160405180910390fd5b60008281526020829052604081208054909181900361244e576040516368c9117d60e11b815260040160405180910390fd5b60005b818110156125f357600083828154811061246d5761246d613dee565b906000526020600020906004020160030160019054906101000a900460ff16600281111561249d5761249d613a40565b036125eb5760008382815481106124b6576124b6613dee565b906000526020600020906004020160030160019054906101000a900460ff16905060018483815481106124eb576124eb613dee565b60009182526020909120600360049092020101805461ff00191661010083600281111561251a5761251a613a40565b0217905550600184838154811061253357612533613dee565b906000526020600020906004020160030160006101000a81548160ff02191690831515021790555083828154811061256d5761256d613dee565b600091825260209091206001600490920201015484546001600160a01b03909116908590849081106125a1576125a1613dee565b60009182526020909120600490910201546040516001600160a01b03909116908890600080516020613f22833981519152906125e1908690600190613e5c565b60405180910390a4505b600101612451565b5050505050565b60006126046133f9565b60040154919050565b6000806126186133f9565b600086815260038201602052604090205490915060ff1661264c57604051633024159f60e21b815260040160405180910390fd5b60008581526020828152604080832080548251818502810185019093528083529192909190849084015b828210156127175760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff80821615156060850152929392608085019261010090920416908111156126f3576126f3613a40565b600281111561270457612704613a40565b8152505081526020019060010190612676565b5050825192935060009150505b81811015611ccd57866001600160a01b031683828151811061274857612748613dee565b6020026020010151600001516001600160a01b03161480156127985750856001600160a01b031683828151811061278157612781613dee565b6020026020010151602001516001600160a01b0316145b156127c6578281815181106127af576127af613dee565b6020026020010151606001519450505050506115c3565b600101612724565b60006127d86135c8565b546001600160a01b0316919050565b6000806127f26133f9565b6001600160a01b03841660009081526002820160205260409020549091508061282e576040516307c5e32d60e41b815260040160405180910390fd5b60008181526020838152604080832080548251818502810185019093528083529192909190849084015b828210156128f95760008481526020908190206040805160a0810182526004860290920180546001600160a01b0390811684526001820154169383019390935260028084015491830191909152600383015460ff80821615156060850152929392608085019261010090920416908111156128d5576128d5613a40565b60028111156128e6576128e6613a40565b8152505081526020019060010190612858565b5050825192935060009150505b8181101561298957866001600160a01b031683828151811061292a5761292a613dee565b6020026020010151602001516001600160a01b03160361298157600183828151811061295857612958613dee565b602002602001015160800151600281111561297557612975613a40565b14979650505050505050565b600101612906565b50604051631ffeda3f60e11b815260040160405180910390fd5b60006129ad61362a565b805490915060ff600160401b82041615906001600160401b03166000811580156129d45750825b90506000826001600160401b031660011480156129f05750303b155b9050811580156129fe575080155b15612a1c5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315612a4657845460ff60401b1916600160401b1785555b612a4e613653565b612a578661365b565b8315612a9d57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b612abb336000356001600160e01b0319166132a5565b6000612ac56133f9565b600083815260038201602052604090205490915060ff16612af957604051633024159f60e21b815260040160405180910390fd5b600082815260208290526040812080549091819003612b2b576040516368c9117d60e11b815260040160405180910390fd5b60005b818110156125f3576000838281548110612b4a57612b4a613dee565b906000526020600020906004020160030160019054906101000a900460ff166002811115612b7a57612b7a613a40565b03612cc8576000838281548110612b9357612b93613dee565b906000526020600020906004020160030160019054906101000a900460ff1690506002848381548110612bc857612bc8613dee565b60009182526020909120600360049092020101805461ff001916610100836002811115612bf757612bf7613a40565b02179055506000848381548110612c1057612c10613dee565b906000526020600020906004020160030160006101000a81548160ff021916908315150217905550838281548110612c4a57612c4a613dee565b600091825260209091206001600490920201015484546001600160a01b0390911690859084908110612c7e57612c7e613dee565b60009182526020909120600490910201546040516001600160a01b03909116908890600080516020613f2283398151915290612cbe908690600290613e5c565b60405180910390a4505b600101612b2e565b6060806000612cdd6133f9565b60048101549091506000816001600160401b03811115612cff57612cff613b3d565b604051908082528060200260200182016040528015612d28578160200160208202803683370190505b5090506000826001600160401b03811115612d4557612d45613b3d565b604051908082528060200260200182016040528015612d7857816020015b6060815260200190600190039081612d635790505b5090506000805b84811015612e2f576000612db1876004018381548110612da157612da1613dee565b9060005260206000200154611263565b805190915015612e2657866004018281548110612dd057612dd0613dee565b9060005260206000200154858481518110612ded57612ded613dee565b60200260200101818152505080848481518110612e0c57612e0c613dee565b60200260200101819052508280612e2290613e43565b9350505b50600101612d7f565b506000816001600160401b03811115612e4a57612e4a613b3d565b604051908082528060200260200182016040528015612e73578160200160208202803683370190505b5090506000826001600160401b03811115612e9057612e90613b3d565b604051908082528060200260200182016040528015612ec357816020015b6060815260200190600190039081612eae5790505b50905060005b83811015612f4857858181518110612ee357612ee3613dee565b6020026020010151838281518110612efd57612efd613dee565b602002602001018181525050848181518110612f1b57612f1b613dee565b6020026020010151828281518110612f3557612f35613dee565b6020908102919091010152600101612ec9565b50909890975095505050505050565b600080612f626133f9565b6002016000846001600160a01b03166001600160a01b031681526020019081526020016000205414159050919050565b6000612f9c6133f9565b6001600160a01b03909216600090815260019290920160205250604090205490565b6060612fc86133f9565b60040180548060200260200160405190810160405280929190818152602001828054801561301557602002820191906000526020600020905b815481526020019060010190808311613001575b5050505050905090565b613035336000356001600160e01b0319166132a5565b600061303f6133f9565b600085815260038201602052604090205490915060ff1661307357604051633024159f60e21b815260040160405180910390fd5b6001600160a01b03831661309a57604051635823d4a360e01b815260040160405180910390fd5b6001600160a01b0382166130c1576040516330f570db60e01b815260040160405180910390fd5b6001600160a01b0383166000908152600182016020526040902054156130fa57604051631827f87960e21b815260040160405180910390fd5b6001600160a01b03821660009081526002820160205260409020541561313357604051636d43b82f60e11b815260040160405180910390fd5b60006040518060a00160405280856001600160a01b03168152602001846001600160a01b031681526020014281526020016000151581526020016000600281111561318057613180613a40565b905260008681526020848152604080832080546001808201835591855293839020855160049095020180546001600160a01b03199081166001600160a01b039687161782559386015191810180549094169190941617909155820151600280830191909155606083015160038301805460ff19811692151592831782556080860151959650869593919261ffff1990911661ff0019909116179061010090849081111561322f5761322f613a40565b021790555050506001600160a01b03808516600081815260018501602090815260408083208a905593871680835260028701909152838220899055925188917f25f808845dcdaa39e4f19b0bc859ca9b29af7f32cf74c9824adb550b3344ea9a91a45050505050565b6060610d76826002611874565b60006132af6135c8565b80549091506001600160a01b0316806132e7576000604051638944034760e01b81526004016132de9190613b29565b60405180910390fd5b60405163b700961360e01b81526001600160a01b0385811660048301523060248301526001600160e01b031985166044830152600091829182919085169063b700961390606401606060405180830381865afa15801561334b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061336f9190613e87565b925092509250826133f05780156133995760405163cc9855ad60e01b815260040160405180910390fd5b63ffffffff8216156133d55760405163a426878960e01b81526001600160a01b038816600482015263ffffffff831660248201526044016132de565b86604051632ecd3d0360e21b81526004016132de9190613b29565b50505050505050565b7fe2b86e7bc31042496c32365d0dfa6ad28cbfffada4949d485907b1975509dd0090565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061348d57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661348161369c565b6001600160a01b031614155b156134ab5760405163703e46dd60e11b815260040160405180910390fd5b565b6134c3336000356001600160e01b0319166132a5565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613520575060408051601f3d908101601f1916820190925261351d91810190613ecc565b60015b61353f5781604051634c9c8ce360e01b81526004016132de9190613b29565b600080516020613f02833981519152811461357057604051632a87526960e21b8152600481018290526024016132de565b61357a83836136b2565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146134ab5760405163703e46dd60e11b815260040160405180910390fd5b60008060ff196135f960017f2d9a26166e6ed6927c3e14d89c4d572a357f984344fcaed3c38c3ff5cdb83f35613e1a565b60405160200161360b91815260200190565b60408051601f1981840301815291905280516020909101201692915050565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610d76565b6134ab613708565b60006136656135c8565b80549091506001600160a01b0316156136935781604051638944034760e01b81526004016132de9190613b29565b61102b8261372d565b6000600080516020613f028339815191526127d8565b6136bb826137bd565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156137005761357a8282613819565b61102b61388f565b6137106138ae565b6134ab57604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b0381166137565780604051638944034760e01b81526004016132de9190613b29565b60006137606135c8565b80546040519192506001600160a01b03808516929116907fa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b7638998019890600090a380546001600160a01b0319166001600160a01b0392909216919091179055565b806001600160a01b03163b6000036137ea5780604051634c9c8ce360e01b81526004016132de9190613b29565b600080516020613f0283398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516138369190613ee5565b600060405180830381855af49150503d8060008114613871576040519150601f19603f3d011682016040523d82523d6000602084013e613876565b606091505b50915091506138868583836138c8565b95945050505050565b34156134ab5760405163b398979f60e01b815260040160405180910390fd5b60006138b861362a565b54600160401b900460ff16919050565b6060826138dd576138d88261391b565b6115c3565b81511580156138f457506001600160a01b0384163b155b156139145783604051639996b31560e01b81526004016132de9190613b29565b50806115c3565b80511561392a57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040805160a0810182526000808252602082018190529181018290526060810182905290608082015290565b50805460008255600402906000526020600020908101906134c391905b808211156139cb5780546001600160a01b0319908116825560018201805490911690556000600282015560038101805461ffff1916905560040161398c565b5090565b80356001600160a01b03811681146139e657600080fd5b919050565b600080600060608486031215613a0057600080fd5b83359250613a10602085016139cf565b9150613a1e604085016139cf565b90509250925092565b600060208284031215613a3957600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b60038110613a7457634e487b7160e01b600052602160045260246000fd5b9052565b600060018060a01b0380835116845280602084015116602085015250604082015160408401526060820151151560608401526080820151613abc6080850182613a56565b50505060a00190565b6020808252825182820181905260009190848201906040850190845b81811015613b0257613af4838551613a78565b938501939250600101613ae1565b50909695505050505050565b600060208284031215613b2057600080fd5b6115c3826139cf565b6001600160a01b0391909116815260200190565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215613b6657600080fd5b613b6f836139cf565b915060208301356001600160401b0380821115613b8b57600080fd5b818501915085601f830112613b9f57600080fd5b813581811115613bb157613bb1613b3d565b604051601f8201601f19908116603f01168101908382118183101715613bd957613bd9613b3d565b81604052828152886020848701011115613bf257600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60208101610d768284613a56565b8035600381106139e657600080fd5b60008060408385031215613c4457600080fd5b82359150613c5460208401613c22565b90509250929050565b60005b83811015613c78578181015183820152602001613c60565b50506000910152565b6020815260008251806020840152613ca0816040850160208701613c5d565b601f01601f19169190910160400192915050565b60008060008060808587031215613cca57600080fd5b84359350613cda602086016139cf565b9250613ce8604086016139cf565b9150613cf660608601613c22565b905092959194509250565b60008151808452602080850194506020840160005b83811015613d3257815187529582019590820190600101613d16565b509495945050505050565b604081526000613d506040830185613d01565b6020838203818501528185518084528284019150828160051b8501018388016000805b84811015613dcb57878403601f19018652825180518086529088019088860190845b81811015613db657613da8838551613a78565b938b01939250600101613d95565b50509688019694505091860191600101613d73565b50919a9950505050505050505050565b6020815260006115c36020830184613d01565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115610d7657610d76613e04565b634e487b7160e01b600052603160045260246000fd5b600060018201613e5557613e55613e04565b5060010190565b60408101613e6a8285613a56565b6115c36020830184613a56565b805180151581146139e657600080fd5b600080600060608486031215613e9c57600080fd5b613ea584613e77565b9250602084015163ffffffff81168114613ebe57600080fd5b9150613a1e60408501613e77565b600060208284031215613ede57600080fd5b5051919050565b60008251613ef7818460208701613c5d565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbce9820ba93cb8fa6b7deb75828222a93826e0786c84473c6a1d2722cdab2cc634a2646970667358221220f597d755df6057a0b958386efe69bd8b782ad0fc9b3a0cb72910aa34da2b5d4b64736f6c63430008180033",
}

// RNUserGovernanceV1 is an auto generated Go binding around an Ethereum contract.
type RNUserGovernanceV1 struct {
	abi abi.ABI
}

// NewRNUserGovernanceV1 creates a new instance of RNUserGovernanceV1.
func NewRNUserGovernanceV1() *RNUserGovernanceV1 {
	parsed, err := RNUserGovernanceV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RNUserGovernanceV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RNUserGovernanceV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAddAddressPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5080035.
//
// Solidity: function addAddressPair(bytes32 userId, address publicAddress, address privateAddress) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackAddAddressPair(userId [32]byte, publicAddress common.Address, privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("addAddressPair", userId, publicAddress, privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackApproveUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4984fd7.
//
// Solidity: function approveUser(bytes32 userId) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackApproveUser(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("approveUser", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAuthority is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackAuthority() []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("authority", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackCheckUserIsApprovedByPrivateAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4ba03ed.
//
// Solidity: function checkUserIsApprovedByPrivateAddress(address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackCheckUserIsApprovedByPrivateAddress(privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("checkUserIsApprovedByPrivateAddress", privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCheckUserIsApprovedByPrivateAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc4ba03ed.
//
// Solidity: function checkUserIsApprovedByPrivateAddress(address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackCheckUserIsApprovedByPrivateAddress(data []byte) (bool, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("checkUserIsApprovedByPrivateAddress", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackCreateUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xabf7bfd8.
//
// Solidity: function createUser(bytes32 userId) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackCreateUser(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("createUser", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetActiveAddressPairCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7be91a72.
//
// Solidity: function getActiveAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetActiveAddressPairCount(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getActiveAddressPairCount", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetActiveAddressPairCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7be91a72.
//
// Solidity: function getActiveAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetActiveAddressPairCount(data []byte) (*big.Int, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getActiveAddressPairCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetActiveAddressPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x166166e3.
//
// Solidity: function getActiveAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetActiveAddressPairs(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getActiveAddressPairs", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetActiveAddressPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x166166e3.
//
// Solidity: function getActiveAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetActiveAddressPairs(data []byte) ([]IUserGovernanceAddressPair, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getActiveAddressPairs", data)
	if err != nil {
		return *new([]IUserGovernanceAddressPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]IUserGovernanceAddressPair)).(*[]IUserGovernanceAddressPair)
	return out0, err
}

// PackGetAddressPairApprovalStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x748ec164.
//
// Solidity: function getAddressPairApprovalStatus(bytes32 userId, address publicAddress, address privateAddress) view returns(uint8)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetAddressPairApprovalStatus(userId [32]byte, publicAddress common.Address, privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getAddressPairApprovalStatus", userId, publicAddress, privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAddressPairApprovalStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x748ec164.
//
// Solidity: function getAddressPairApprovalStatus(bytes32 userId, address publicAddress, address privateAddress) view returns(uint8)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetAddressPairApprovalStatus(data []byte) (uint8, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getAddressPairApprovalStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackGetAddressPairsByApprovalStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x99a43bb0.
//
// Solidity: function getAddressPairsByApprovalStatus(bytes32 userId, uint8 status) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetAddressPairsByApprovalStatus(userId [32]byte, status uint8) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getAddressPairsByApprovalStatus", userId, status)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAddressPairsByApprovalStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x99a43bb0.
//
// Solidity: function getAddressPairsByApprovalStatus(bytes32 userId, uint8 status) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetAddressPairsByApprovalStatus(data []byte) ([]IUserGovernanceAddressPair, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getAddressPairsByApprovalStatus", data)
	if err != nil {
		return *new([]IUserGovernanceAddressPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]IUserGovernanceAddressPair)).(*[]IUserGovernanceAddressPair)
	return out0, err
}

// PackGetAllPendingAddressPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcaa60d4c.
//
// Solidity: function getAllPendingAddressPairs() view returns(bytes32[], (address,address,uint256,bool,uint8)[][])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetAllPendingAddressPairs() []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getAllPendingAddressPairs")
	if err != nil {
		panic(err)
	}
	return enc
}

// GetAllPendingAddressPairsOutput serves as a container for the return parameters of contract
// method GetAllPendingAddressPairs.
type GetAllPendingAddressPairsOutput struct {
	Arg0 [][32]byte
	Arg1 [][]IUserGovernanceAddressPair
}

// UnpackGetAllPendingAddressPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcaa60d4c.
//
// Solidity: function getAllPendingAddressPairs() view returns(bytes32[], (address,address,uint256,bool,uint8)[][])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetAllPendingAddressPairs(data []byte) (GetAllPendingAddressPairsOutput, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getAllPendingAddressPairs", data)
	outstruct := new(GetAllPendingAddressPairsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([][]IUserGovernanceAddressPair)).(*[][]IUserGovernanceAddressPair)
	return *outstruct, err

}

// PackGetAllUsers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(bytes32[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetAllUsers() []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getAllUsers")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAllUsers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(bytes32[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetAllUsers(data []byte) ([][32]byte, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getAllUsers", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, err
}

// PackGetApprovedAddressPairCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x741827de.
//
// Solidity: function getApprovedAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetApprovedAddressPairCount(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getApprovedAddressPairCount", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApprovedAddressPairCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x741827de.
//
// Solidity: function getApprovedAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetApprovedAddressPairCount(data []byte) (*big.Int, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getApprovedAddressPairCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetApprovedAddressPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18325426.
//
// Solidity: function getApprovedAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetApprovedAddressPairs(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getApprovedAddressPairs", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApprovedAddressPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18325426.
//
// Solidity: function getApprovedAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetApprovedAddressPairs(data []byte) ([]IUserGovernanceAddressPair, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getApprovedAddressPairs", data)
	if err != nil {
		return *new([]IUserGovernanceAddressPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]IUserGovernanceAddressPair)).(*[]IUserGovernanceAddressPair)
	return out0, err
}

// PackGetPendingAddressPairCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c699038.
//
// Solidity: function getPendingAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetPendingAddressPairCount(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getPendingAddressPairCount", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPendingAddressPairCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c699038.
//
// Solidity: function getPendingAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetPendingAddressPairCount(data []byte) (*big.Int, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getPendingAddressPairCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetPendingAddressPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d3070b1.
//
// Solidity: function getPendingAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetPendingAddressPairs(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getPendingAddressPairs", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPendingAddressPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5d3070b1.
//
// Solidity: function getPendingAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetPendingAddressPairs(data []byte) ([]IUserGovernanceAddressPair, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getPendingAddressPairs", data)
	if err != nil {
		return *new([]IUserGovernanceAddressPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]IUserGovernanceAddressPair)).(*[]IUserGovernanceAddressPair)
	return out0, err
}

// PackGetPrivateAddressFromPublic is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eb2f70.
//
// Solidity: function getPrivateAddressFromPublic(address publicAddress) view returns(address)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetPrivateAddressFromPublic(publicAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getPrivateAddressFromPublic", publicAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPrivateAddressFromPublic is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eb2f70.
//
// Solidity: function getPrivateAddressFromPublic(address publicAddress) view returns(address)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetPrivateAddressFromPublic(data []byte) (common.Address, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getPrivateAddressFromPublic", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetPublicAddressFromPrivate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2463a1be.
//
// Solidity: function getPublicAddressFromPrivate(address privateAddress) view returns(address)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetPublicAddressFromPrivate(privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getPublicAddressFromPrivate", privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPublicAddressFromPrivate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2463a1be.
//
// Solidity: function getPublicAddressFromPrivate(address privateAddress) view returns(address)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetPublicAddressFromPrivate(data []byte) (common.Address, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getPublicAddressFromPrivate", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetRejectedAddressPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfbc56938.
//
// Solidity: function getRejectedAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetRejectedAddressPairs(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getRejectedAddressPairs", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRejectedAddressPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfbc56938.
//
// Solidity: function getRejectedAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetRejectedAddressPairs(data []byte) ([]IUserGovernanceAddressPair, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getRejectedAddressPairs", data)
	if err != nil {
		return *new([]IUserGovernanceAddressPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]IUserGovernanceAddressPair)).(*[]IUserGovernanceAddressPair)
	return out0, err
}

// PackGetUserAddressPairCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ee126a5.
//
// Solidity: function getUserAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetUserAddressPairCount(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getUserAddressPairCount", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserAddressPairCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ee126a5.
//
// Solidity: function getUserAddressPairCount(bytes32 userId) view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetUserAddressPairCount(data []byte) (*big.Int, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getUserAddressPairCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetUserAddressPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06096d60.
//
// Solidity: function getUserAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetUserAddressPairs(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getUserAddressPairs", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserAddressPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06096d60.
//
// Solidity: function getUserAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetUserAddressPairs(data []byte) ([]IUserGovernanceAddressPair, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getUserAddressPairs", data)
	if err != nil {
		return *new([]IUserGovernanceAddressPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]IUserGovernanceAddressPair)).(*[]IUserGovernanceAddressPair)
	return out0, err
}

// PackGetUserCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetUserCount() []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getUserCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetUserCount(data []byte) (*big.Int, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getUserCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetUserIdByPrivateAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47369a1f.
//
// Solidity: function getUserIdByPrivateAddress(address privateAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetUserIdByPrivateAddress(privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getUserIdByPrivateAddress", privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserIdByPrivateAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x47369a1f.
//
// Solidity: function getUserIdByPrivateAddress(address privateAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetUserIdByPrivateAddress(data []byte) ([32]byte, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getUserIdByPrivateAddress", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetUserIdByPublicAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfccb7583.
//
// Solidity: function getUserIdByPublicAddress(address publicAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackGetUserIdByPublicAddress(publicAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("getUserIdByPublicAddress", publicAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserIdByPublicAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfccb7583.
//
// Solidity: function getUserIdByPublicAddress(address publicAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackGetUserIdByPublicAddress(data []byte) ([32]byte, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("getUserIdByPublicAddress", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackHasUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9686e884.
//
// Solidity: function hasUser(bytes32 userId) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackHasUser(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("hasUser", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9686e884.
//
// Solidity: function hasUser(bytes32 userId) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackHasUser(data []byte) (bool, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("hasUser", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address authority_) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackInitialize(authority common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("initialize", authority)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsAddressPairActive is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb97f100f.
//
// Solidity: function isAddressPairActive(bytes32 userId, address publicAddress, address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackIsAddressPairActive(userId [32]byte, publicAddress common.Address, privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("isAddressPairActive", userId, publicAddress, privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsAddressPairActive is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb97f100f.
//
// Solidity: function isAddressPairActive(bytes32 userId, address publicAddress, address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackIsAddressPairActive(data []byte) (bool, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("isAddressPairActive", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsAddressPairApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ad5e2fb.
//
// Solidity: function isAddressPairApproved(bytes32 userId, address publicAddress, address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackIsAddressPairApproved(userId [32]byte, publicAddress common.Address, privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("isAddressPairApproved", userId, publicAddress, privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsAddressPairApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ad5e2fb.
//
// Solidity: function isAddressPairApproved(bytes32 userId, address publicAddress, address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackIsAddressPairApproved(data []byte) (bool, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("isAddressPairApproved", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsPrivateAddressMapped is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd737cd9e.
//
// Solidity: function isPrivateAddressMapped(address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackIsPrivateAddressMapped(privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("isPrivateAddressMapped", privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsPrivateAddressMapped is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd737cd9e.
//
// Solidity: function isPrivateAddressMapped(address privateAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackIsPrivateAddressMapped(data []byte) (bool, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("isPrivateAddressMapped", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsPublicAddressMapped is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2861a4bb.
//
// Solidity: function isPublicAddressMapped(address publicAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackIsPublicAddressMapped(publicAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("isPublicAddressMapped", publicAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsPublicAddressMapped is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2861a4bb.
//
// Solidity: function isPublicAddressMapped(address publicAddress) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackIsPublicAddressMapped(data []byte) (bool, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("isPublicAddressMapped", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackPrivateAddressToUserId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1c288cb.
//
// Solidity: function privateAddressToUserId(address privateAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackPrivateAddressToUserId(privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("privateAddressToUserId", privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPrivateAddressToUserId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1c288cb.
//
// Solidity: function privateAddressToUserId(address privateAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackPrivateAddressToUserId(data []byte) ([32]byte, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("privateAddressToUserId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackProxiableUUID() []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackPublicAddressToUserId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd77a8ae.
//
// Solidity: function publicAddressToUserId(address publicAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackPublicAddressToUserId(publicAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("publicAddressToUserId", publicAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPublicAddressToUserId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd77a8ae.
//
// Solidity: function publicAddressToUserId(address publicAddress) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackPublicAddressToUserId(data []byte) ([32]byte, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("publicAddressToUserId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRejectUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc92c9481.
//
// Solidity: function rejectUser(bytes32 userId) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackRejectUser(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("rejectUser", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemoveAddressPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0488438c.
//
// Solidity: function removeAddressPair(bytes32 userId, address publicAddress, address privateAddress) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackRemoveAddressPair(userId [32]byte, publicAddress common.Address, privateAddress common.Address) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("removeAddressPair", userId, publicAddress, privateAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemoveUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a9054d8.
//
// Solidity: function removeUser(bytes32 userId) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackRemoveUser(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("removeUser", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetAddressPairApprovalStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb26b32c2.
//
// Solidity: function setAddressPairApprovalStatus(bytes32 userId, address publicAddress, address privateAddress, uint8 newStatus) returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackSetAddressPairApprovalStatus(userId [32]byte, publicAddress common.Address, privateAddress common.Address, newStatus uint8) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("setAddressPairApprovalStatus", userId, publicAddress, privateAddress, newStatus)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUserAddressPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74ab379f.
//
// Solidity: function userAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackUserAddressPairs(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("userAddressPairs", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUserAddressPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x74ab379f.
//
// Solidity: function userAddressPairs(bytes32 userId) view returns((address,address,uint256,bool,uint8)[])
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUserAddressPairs(data []byte) ([]IUserGovernanceAddressPair, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("userAddressPairs", data)
	if err != nil {
		return *new([]IUserGovernanceAddressPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]IUserGovernanceAddressPair)).(*[]IUserGovernanceAddressPair)
	return out0, err
}

// PackUserExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2e8452c.
//
// Solidity: function userExists(bytes32 userId) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackUserExists(userId [32]byte) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("userExists", userId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUserExists is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa2e8452c.
//
// Solidity: function userExists(bytes32 userId) view returns(bool)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUserExists(data []byte) (bool, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("userExists", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUserIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4635fd68.
//
// Solidity: function userIds(uint256 index) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) PackUserIds(index *big.Int) []byte {
	enc, err := rNUserGovernanceV1.abi.Pack("userIds", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUserIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4635fd68.
//
// Solidity: function userIds(uint256 index) view returns(bytes32)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUserIds(data []byte) ([32]byte, error) {
	out, err := rNUserGovernanceV1.abi.Unpack("userIds", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// RNUserGovernanceV1AddressPairAdded represents a AddressPairAdded event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1AddressPairAdded struct {
	UserId         [32]byte
	PublicAddress  common.Address
	PrivateAddress common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1AddressPairAddedEventName = "AddressPairAdded"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1AddressPairAdded) ContractEventName() string {
	return RNUserGovernanceV1AddressPairAddedEventName
}

// UnpackAddressPairAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AddressPairAdded(bytes32 indexed userId, address indexed publicAddress, address indexed privateAddress)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackAddressPairAddedEvent(log *types.Log) (*RNUserGovernanceV1AddressPairAdded, error) {
	event := "AddressPairAdded"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1AddressPairAdded)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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

// RNUserGovernanceV1AddressPairApprovalChanged represents a AddressPairApprovalChanged event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1AddressPairApprovalChanged struct {
	UserId         [32]byte
	PublicAddress  common.Address
	PrivateAddress common.Address
	OldStatus      uint8
	NewStatus      uint8
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1AddressPairApprovalChangedEventName = "AddressPairApprovalChanged"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1AddressPairApprovalChanged) ContractEventName() string {
	return RNUserGovernanceV1AddressPairApprovalChangedEventName
}

// UnpackAddressPairApprovalChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AddressPairApprovalChanged(bytes32 indexed userId, address indexed publicAddress, address indexed privateAddress, uint8 oldStatus, uint8 newStatus)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackAddressPairApprovalChangedEvent(log *types.Log) (*RNUserGovernanceV1AddressPairApprovalChanged, error) {
	event := "AddressPairApprovalChanged"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1AddressPairApprovalChanged)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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

// RNUserGovernanceV1AddressPairRemoved represents a AddressPairRemoved event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1AddressPairRemoved struct {
	UserId         [32]byte
	PublicAddress  common.Address
	PrivateAddress common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1AddressPairRemovedEventName = "AddressPairRemoved"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1AddressPairRemoved) ContractEventName() string {
	return RNUserGovernanceV1AddressPairRemovedEventName
}

// UnpackAddressPairRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AddressPairRemoved(bytes32 indexed userId, address indexed publicAddress, address indexed privateAddress)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackAddressPairRemovedEvent(log *types.Log) (*RNUserGovernanceV1AddressPairRemoved, error) {
	event := "AddressPairRemoved"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1AddressPairRemoved)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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

// RNUserGovernanceV1AuthorityUpdated represents a AuthorityUpdated event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1AuthorityUpdated) ContractEventName() string {
	return RNUserGovernanceV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*RNUserGovernanceV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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

// RNUserGovernanceV1Initialized represents a Initialized event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1Initialized) ContractEventName() string {
	return RNUserGovernanceV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackInitializedEvent(log *types.Log) (*RNUserGovernanceV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1Initialized)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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

// RNUserGovernanceV1Upgraded represents a Upgraded event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1Upgraded) ContractEventName() string {
	return RNUserGovernanceV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUpgradedEvent(log *types.Log) (*RNUserGovernanceV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1Upgraded)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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

// RNUserGovernanceV1UserCreated represents a UserCreated event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1UserCreated struct {
	UserId [32]byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1UserCreatedEventName = "UserCreated"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1UserCreated) ContractEventName() string {
	return RNUserGovernanceV1UserCreatedEventName
}

// UnpackUserCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserCreated(bytes32 indexed userId)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUserCreatedEvent(log *types.Log) (*RNUserGovernanceV1UserCreated, error) {
	event := "UserCreated"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1UserCreated)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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

// RNUserGovernanceV1UserRemoved represents a UserRemoved event raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1UserRemoved struct {
	UserId [32]byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const RNUserGovernanceV1UserRemovedEventName = "UserRemoved"

// ContractEventName returns the user-defined event name.
func (RNUserGovernanceV1UserRemoved) ContractEventName() string {
	return RNUserGovernanceV1UserRemovedEventName
}

// UnpackUserRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserRemoved(bytes32 indexed userId)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUserRemovedEvent(log *types.Log) (*RNUserGovernanceV1UserRemoved, error) {
	event := "UserRemoved"
	if log.Topics[0] != rNUserGovernanceV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNUserGovernanceV1UserRemoved)
	if len(log.Data) > 0 {
		if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNUserGovernanceV1.abi.Events[event].Inputs {
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
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1AddressPairByPrivateKeyNotFound"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1AddressPairByPrivateKeyNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1AddressPairNotFound"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1AddressPairNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1InvalidPrivateAddress"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1InvalidPrivateAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1InvalidPublicAddress"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1InvalidPublicAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1InvalidUserId"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1InvalidUserIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1PrivateAddressAlreadyMapped"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1PrivateAddressAlreadyMappedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1PrivateAddressNotMapped"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1PrivateAddressNotMappedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1PrivateAddressNotMappedToUser"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1PrivateAddressNotMappedToUserError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1PublicAddressAlreadyMapped"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1PublicAddressAlreadyMappedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1PublicAddressNotMapped"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1PublicAddressNotMappedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1PublicAddressNotMappedToUser"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1PublicAddressNotMappedToUserError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1UserAlreadyExists"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1UserAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1UserDoesNotExist"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1UserDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RNUserGovernanceV1UserHasNoAddressPairs"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRNUserGovernanceV1UserHasNoAddressPairsError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNUserGovernanceV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return rNUserGovernanceV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// RNUserGovernanceV1AddressEmptyCode represents a AddressEmptyCode error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func RNUserGovernanceV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackAddressEmptyCodeError(raw []byte) (*RNUserGovernanceV1AddressEmptyCode, error) {
	out := new(RNUserGovernanceV1AddressEmptyCode)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func RNUserGovernanceV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackERC1967InvalidImplementationError(raw []byte) (*RNUserGovernanceV1ERC1967InvalidImplementation, error) {
	out := new(RNUserGovernanceV1ERC1967InvalidImplementation)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func RNUserGovernanceV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackERC1967NonPayableError(raw []byte) (*RNUserGovernanceV1ERC1967NonPayable, error) {
	out := new(RNUserGovernanceV1ERC1967NonPayable)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1FailedCall represents a FailedCall error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func RNUserGovernanceV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackFailedCallError(raw []byte) (*RNUserGovernanceV1FailedCall, error) {
	out := new(RNUserGovernanceV1FailedCall)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1InvalidInitialization represents a InvalidInitialization error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func RNUserGovernanceV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackInvalidInitializationError(raw []byte) (*RNUserGovernanceV1InvalidInitialization, error) {
	out := new(RNUserGovernanceV1InvalidInitialization)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1NotInitializing represents a NotInitializing error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func RNUserGovernanceV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackNotInitializingError(raw []byte) (*RNUserGovernanceV1NotInitializing, error) {
	out := new(RNUserGovernanceV1NotInitializing)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1AddressPairByPrivateKeyNotFound represents a RNUserGovernanceV1__AddressPairByPrivateKeyNotFound error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1AddressPairByPrivateKeyNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__AddressPairByPrivateKeyNotFound()
func RNUserGovernanceV1RNUserGovernanceV1AddressPairByPrivateKeyNotFoundErrorID() common.Hash {
	return common.HexToHash("0x3ffdb47e9df8732210100989aa3cb247c325cc0a6dc943f29cffaafd3fc2bdad")
}

// UnpackRNUserGovernanceV1AddressPairByPrivateKeyNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__AddressPairByPrivateKeyNotFound()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1AddressPairByPrivateKeyNotFoundError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1AddressPairByPrivateKeyNotFound, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1AddressPairByPrivateKeyNotFound)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1AddressPairByPrivateKeyNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1AddressPairNotFound represents a RNUserGovernanceV1__AddressPairNotFound error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1AddressPairNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__AddressPairNotFound()
func RNUserGovernanceV1RNUserGovernanceV1AddressPairNotFoundErrorID() common.Hash {
	return common.HexToHash("0x0ef1d6c146148bc2f4e7f4c7815de2e3ade483a577b263d8320b1f0c28c2233d")
}

// UnpackRNUserGovernanceV1AddressPairNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__AddressPairNotFound()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1AddressPairNotFoundError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1AddressPairNotFound, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1AddressPairNotFound)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1AddressPairNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1InvalidPrivateAddress represents a RNUserGovernanceV1__InvalidPrivateAddress error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1InvalidPrivateAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__InvalidPrivateAddress()
func RNUserGovernanceV1RNUserGovernanceV1InvalidPrivateAddressErrorID() common.Hash {
	return common.HexToHash("0x30f570db0273b5f8e2f4d5f3b80ead42c142e0152be36debd5ddc8fcf7f69a98")
}

// UnpackRNUserGovernanceV1InvalidPrivateAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__InvalidPrivateAddress()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1InvalidPrivateAddressError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1InvalidPrivateAddress, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1InvalidPrivateAddress)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1InvalidPrivateAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1InvalidPublicAddress represents a RNUserGovernanceV1__InvalidPublicAddress error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1InvalidPublicAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__InvalidPublicAddress()
func RNUserGovernanceV1RNUserGovernanceV1InvalidPublicAddressErrorID() common.Hash {
	return common.HexToHash("0x5823d4a32d858fa2384c77da71dbe27b8101a742457651e7d41818948d19d796")
}

// UnpackRNUserGovernanceV1InvalidPublicAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__InvalidPublicAddress()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1InvalidPublicAddressError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1InvalidPublicAddress, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1InvalidPublicAddress)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1InvalidPublicAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1InvalidUserId represents a RNUserGovernanceV1__InvalidUserId error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1InvalidUserId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__InvalidUserId()
func RNUserGovernanceV1RNUserGovernanceV1InvalidUserIdErrorID() common.Hash {
	return common.HexToHash("0x7ee6241cf6fcf0c67e5920d689b8d148da52dffe8ce322208efd6d7975e9e5d8")
}

// UnpackRNUserGovernanceV1InvalidUserIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__InvalidUserId()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1InvalidUserIdError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1InvalidUserId, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1InvalidUserId)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1InvalidUserId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1PrivateAddressAlreadyMapped represents a RNUserGovernanceV1__PrivateAddressAlreadyMapped error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1PrivateAddressAlreadyMapped struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__PrivateAddressAlreadyMapped()
func RNUserGovernanceV1RNUserGovernanceV1PrivateAddressAlreadyMappedErrorID() common.Hash {
	return common.HexToHash("0xda87705e2d1c3cc43a432f6898c1ae8d02c01c435bd53495d4a606f72302b957")
}

// UnpackRNUserGovernanceV1PrivateAddressAlreadyMappedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__PrivateAddressAlreadyMapped()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1PrivateAddressAlreadyMappedError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1PrivateAddressAlreadyMapped, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1PrivateAddressAlreadyMapped)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1PrivateAddressAlreadyMapped", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMapped represents a RNUserGovernanceV1__PrivateAddressNotMapped error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMapped struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__PrivateAddressNotMapped()
func RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMappedErrorID() common.Hash {
	return common.HexToHash("0x7c5e32d03648ebaabd275e753dfd06020730f364ee51d61cfe1974265d964a80")
}

// UnpackRNUserGovernanceV1PrivateAddressNotMappedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__PrivateAddressNotMapped()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1PrivateAddressNotMappedError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMapped, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMapped)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1PrivateAddressNotMapped", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMappedToUser represents a RNUserGovernanceV1__PrivateAddressNotMappedToUser error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMappedToUser struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__PrivateAddressNotMappedToUser()
func RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMappedToUserErrorID() common.Hash {
	return common.HexToHash("0xb623c713f20c3e7594861e67e9ebb560edb2616e5b0d0714a36534dbcfb9b708")
}

// UnpackRNUserGovernanceV1PrivateAddressNotMappedToUserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__PrivateAddressNotMappedToUser()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1PrivateAddressNotMappedToUserError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMappedToUser, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1PrivateAddressNotMappedToUser)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1PrivateAddressNotMappedToUser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1PublicAddressAlreadyMapped represents a RNUserGovernanceV1__PublicAddressAlreadyMapped error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1PublicAddressAlreadyMapped struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__PublicAddressAlreadyMapped()
func RNUserGovernanceV1RNUserGovernanceV1PublicAddressAlreadyMappedErrorID() common.Hash {
	return common.HexToHash("0x609fe1e448d69e159590389e2c4f4ee65a102eae75c631fcab2b1191c4850b30")
}

// UnpackRNUserGovernanceV1PublicAddressAlreadyMappedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__PublicAddressAlreadyMapped()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1PublicAddressAlreadyMappedError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1PublicAddressAlreadyMapped, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1PublicAddressAlreadyMapped)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1PublicAddressAlreadyMapped", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMapped represents a RNUserGovernanceV1__PublicAddressNotMapped error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMapped struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__PublicAddressNotMapped()
func RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMappedErrorID() common.Hash {
	return common.HexToHash("0x74849121ec37d8a058af8255cc9ec8b909689891fd5a250af2c52fab088d5303")
}

// UnpackRNUserGovernanceV1PublicAddressNotMappedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__PublicAddressNotMapped()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1PublicAddressNotMappedError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMapped, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMapped)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1PublicAddressNotMapped", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMappedToUser represents a RNUserGovernanceV1__PublicAddressNotMappedToUser error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMappedToUser struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__PublicAddressNotMappedToUser()
func RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMappedToUserErrorID() common.Hash {
	return common.HexToHash("0x5e26379f5278050e9b7e67060f5d22c8b2082b75d440f442795f27c0d090d06d")
}

// UnpackRNUserGovernanceV1PublicAddressNotMappedToUserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__PublicAddressNotMappedToUser()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1PublicAddressNotMappedToUserError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMappedToUser, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1PublicAddressNotMappedToUser)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1PublicAddressNotMappedToUser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1UserAlreadyExists represents a RNUserGovernanceV1__UserAlreadyExists error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1UserAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__UserAlreadyExists()
func RNUserGovernanceV1RNUserGovernanceV1UserAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xb3c028ae8b1230637f5bce23d25c5dd11081b79b5f9a03f04da270d0fb55396d")
}

// UnpackRNUserGovernanceV1UserAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__UserAlreadyExists()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1UserAlreadyExistsError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1UserAlreadyExists, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1UserAlreadyExists)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1UserAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1UserDoesNotExist represents a RNUserGovernanceV1__UserDoesNotExist error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1UserDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__UserDoesNotExist()
func RNUserGovernanceV1RNUserGovernanceV1UserDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0xc090567c700ade0c0e93290e1be6494cb885dcfae0778f13e9c03911cb592c4c")
}

// UnpackRNUserGovernanceV1UserDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__UserDoesNotExist()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1UserDoesNotExistError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1UserDoesNotExist, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1UserDoesNotExist)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1UserDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RNUserGovernanceV1UserHasNoAddressPairs represents a RNUserGovernanceV1__UserHasNoAddressPairs error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RNUserGovernanceV1UserHasNoAddressPairs struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNUserGovernanceV1__UserHasNoAddressPairs()
func RNUserGovernanceV1RNUserGovernanceV1UserHasNoAddressPairsErrorID() common.Hash {
	return common.HexToHash("0xd19222fa3046443e068d70cc7b7eb9c30a99a2a5235a5360b19553fa70df469e")
}

// UnpackRNUserGovernanceV1UserHasNoAddressPairsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNUserGovernanceV1__UserHasNoAddressPairs()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRNUserGovernanceV1UserHasNoAddressPairsError(raw []byte) (*RNUserGovernanceV1RNUserGovernanceV1UserHasNoAddressPairs, error) {
	out := new(RNUserGovernanceV1RNUserGovernanceV1UserHasNoAddressPairs)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RNUserGovernanceV1UserHasNoAddressPairs", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func RNUserGovernanceV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*RNUserGovernanceV1RaylsAccessManagedContractPaused, error) {
	out := new(RNUserGovernanceV1RaylsAccessManagedContractPaused)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func RNUserGovernanceV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*RNUserGovernanceV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(RNUserGovernanceV1RaylsAccessManagedInvalidAuthority)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func RNUserGovernanceV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*RNUserGovernanceV1RaylsAccessManagedMustSchedule, error) {
	out := new(RNUserGovernanceV1RaylsAccessManagedMustSchedule)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func RNUserGovernanceV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*RNUserGovernanceV1RaylsAccessManagedUnauthorized, error) {
	out := new(RNUserGovernanceV1RaylsAccessManagedUnauthorized)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func RNUserGovernanceV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*RNUserGovernanceV1UUPSUnauthorizedCallContext, error) {
	out := new(RNUserGovernanceV1UUPSUnauthorizedCallContext)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNUserGovernanceV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the RNUserGovernanceV1 contract.
type RNUserGovernanceV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func RNUserGovernanceV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (rNUserGovernanceV1 *RNUserGovernanceV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*RNUserGovernanceV1UUPSUnsupportedProxiableUUID, error) {
	out := new(RNUserGovernanceV1UUPSUnsupportedProxiableUUID)
	if err := rNUserGovernanceV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
