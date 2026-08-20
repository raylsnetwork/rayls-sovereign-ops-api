// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RNMessageLib

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

// RNMessageLibMetaData contains all meta data concerning the RNMessageLib contract.
var RNMessageLibMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "RNMessageLib",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204939368887c3803d74251be24a544755fd85a5f8f928b1b8afd633627660383764736f6c63430008180033",
}

// RNMessageLib is an auto generated Go binding around an Ethereum contract.
type RNMessageLib struct {
	abi abi.ABI
}

// NewRNMessageLib creates a new instance of RNMessageLib.
func NewRNMessageLib() *RNMessageLib {
	parsed, err := RNMessageLibMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RNMessageLib{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RNMessageLib) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}
