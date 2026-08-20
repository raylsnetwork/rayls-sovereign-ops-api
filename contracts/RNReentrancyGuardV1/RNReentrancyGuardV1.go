// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RNReentrancyGuardV1

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

// RNReentrancyGuardV1MetaData contains all meta data concerning the RNReentrancyGuardV1 contract.
var RNReentrancyGuardV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNReentrancyGuardV1__ReceiveReentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNReentrancyGuardV1__SendReentrancy\",\"inputs\":[]}]",
	ID:  "RNReentrancyGuardV1",
}

// RNReentrancyGuardV1 is an auto generated Go binding around an Ethereum contract.
type RNReentrancyGuardV1 struct {
	abi abi.ABI
}

// NewRNReentrancyGuardV1 creates a new instance of RNReentrancyGuardV1.
func NewRNReentrancyGuardV1() *RNReentrancyGuardV1 {
	parsed, err := RNReentrancyGuardV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RNReentrancyGuardV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RNReentrancyGuardV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (rNReentrancyGuardV1 *RNReentrancyGuardV1) PackInitialize() []byte {
	enc, err := rNReentrancyGuardV1.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// RNReentrancyGuardV1Initialized represents a Initialized event raised by the RNReentrancyGuardV1 contract.
type RNReentrancyGuardV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const RNReentrancyGuardV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (RNReentrancyGuardV1Initialized) ContractEventName() string {
	return RNReentrancyGuardV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (rNReentrancyGuardV1 *RNReentrancyGuardV1) UnpackInitializedEvent(log *types.Log) (*RNReentrancyGuardV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != rNReentrancyGuardV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNReentrancyGuardV1Initialized)
	if len(log.Data) > 0 {
		if err := rNReentrancyGuardV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNReentrancyGuardV1.abi.Events[event].Inputs {
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
func (rNReentrancyGuardV1 *RNReentrancyGuardV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], rNReentrancyGuardV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return rNReentrancyGuardV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNReentrancyGuardV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return rNReentrancyGuardV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNReentrancyGuardV1.abi.Errors["RNReentrancyGuardV1ReceiveReentrancy"].ID.Bytes()[:4]) {
		return rNReentrancyGuardV1.UnpackRNReentrancyGuardV1ReceiveReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNReentrancyGuardV1.abi.Errors["RNReentrancyGuardV1SendReentrancy"].ID.Bytes()[:4]) {
		return rNReentrancyGuardV1.UnpackRNReentrancyGuardV1SendReentrancyError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// RNReentrancyGuardV1InvalidInitialization represents a InvalidInitialization error raised by the RNReentrancyGuardV1 contract.
type RNReentrancyGuardV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func RNReentrancyGuardV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (rNReentrancyGuardV1 *RNReentrancyGuardV1) UnpackInvalidInitializationError(raw []byte) (*RNReentrancyGuardV1InvalidInitialization, error) {
	out := new(RNReentrancyGuardV1InvalidInitialization)
	if err := rNReentrancyGuardV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNReentrancyGuardV1NotInitializing represents a NotInitializing error raised by the RNReentrancyGuardV1 contract.
type RNReentrancyGuardV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func RNReentrancyGuardV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (rNReentrancyGuardV1 *RNReentrancyGuardV1) UnpackNotInitializingError(raw []byte) (*RNReentrancyGuardV1NotInitializing, error) {
	out := new(RNReentrancyGuardV1NotInitializing)
	if err := rNReentrancyGuardV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNReentrancyGuardV1RNReentrancyGuardV1ReceiveReentrancy represents a RNReentrancyGuardV1__ReceiveReentrancy error raised by the RNReentrancyGuardV1 contract.
type RNReentrancyGuardV1RNReentrancyGuardV1ReceiveReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNReentrancyGuardV1__ReceiveReentrancy()
func RNReentrancyGuardV1RNReentrancyGuardV1ReceiveReentrancyErrorID() common.Hash {
	return common.HexToHash("0x17326c63e33303d77a4e1a51b7a04424a96a05b70c445492bc1f97ec9e33b2e8")
}

// UnpackRNReentrancyGuardV1ReceiveReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNReentrancyGuardV1__ReceiveReentrancy()
func (rNReentrancyGuardV1 *RNReentrancyGuardV1) UnpackRNReentrancyGuardV1ReceiveReentrancyError(raw []byte) (*RNReentrancyGuardV1RNReentrancyGuardV1ReceiveReentrancy, error) {
	out := new(RNReentrancyGuardV1RNReentrancyGuardV1ReceiveReentrancy)
	if err := rNReentrancyGuardV1.abi.UnpackIntoInterface(out, "RNReentrancyGuardV1ReceiveReentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNReentrancyGuardV1RNReentrancyGuardV1SendReentrancy represents a RNReentrancyGuardV1__SendReentrancy error raised by the RNReentrancyGuardV1 contract.
type RNReentrancyGuardV1RNReentrancyGuardV1SendReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNReentrancyGuardV1__SendReentrancy()
func RNReentrancyGuardV1RNReentrancyGuardV1SendReentrancyErrorID() common.Hash {
	return common.HexToHash("0x5c3d889ea3f0e73f68a749d7795ddad4411fe30d6bf3a5394d8355213c911258")
}

// UnpackRNReentrancyGuardV1SendReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNReentrancyGuardV1__SendReentrancy()
func (rNReentrancyGuardV1 *RNReentrancyGuardV1) UnpackRNReentrancyGuardV1SendReentrancyError(raw []byte) (*RNReentrancyGuardV1RNReentrancyGuardV1SendReentrancy, error) {
	out := new(RNReentrancyGuardV1RNReentrancyGuardV1SendReentrancy)
	if err := rNReentrancyGuardV1.abi.UnpackIntoInterface(out, "RNReentrancyGuardV1SendReentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}
