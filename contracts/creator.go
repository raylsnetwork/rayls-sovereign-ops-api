package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// EnsureCode verifies that contract bytecode exists at address (using the caller's ctx).
//
// abigen v2 bindings are instantiated without an address — New<Name>() returns a stateless
// wrapper and the address is supplied per call (CallContract with Pack/Unpack, or Instance).
// Callers therefore validate code presence separately before issuing calls.
func EnsureCode(ctx context.Context, client bind.ContractBackend, address common.Address) error {
	code, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		return fmt.Errorf("check code at address %s: %w", address.Hex(), err)
	}
	if len(code) == 0 {
		return fmt.Errorf("no contract code at address %s", address.Hex())
	}
	return nil
}
