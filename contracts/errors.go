package contracts

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

// customErrByID is the package-level registry of custom Solidity errors, keyed by 4-byte selector.
// Adapters populate it at construction (see RegisterErrorsFromMetaData) so decodeRevertReason can turn
// an otherwise-opaque revert selector into a readable "ErrorName(arg=value, …)" string.
var (
	customErrMu   sync.RWMutex
	customErrByID = map[[4]byte]abi.Error{}
)

// RegisterErrorABI merges the custom errors declared in a into the revert-decoding registry, keyed by
// 4-byte selector. It is safe to call multiple times and from multiple goroutines; the first
// registration for a selector wins (identical selectors denote the same error signature).
func RegisterErrorABI(a abi.ABI) {
	customErrMu.Lock()
	defer customErrMu.Unlock()
	for _, e := range a.Errors {
		var sel [4]byte
		copy(sel[:], e.ID[:selectorLength])
		if _, exists := customErrByID[sel]; !exists {
			customErrByID[sel] = e
		}
	}
}

// RegisterErrorsFromMetaData parses the contract metadata's ABI and registers its custom errors. It is
// the convenience an adapter uses at construction, e.g.
// contracts.RegisterErrorsFromMetaData(&PNTokenRegistryV1.PNTokenRegistryV1MetaData).
func RegisterErrorsFromMetaData(m *bind.MetaData) error {
	parsed, err := m.ParseABI()
	if err != nil {
		return fmt.Errorf("parse ABI for error registry: %w", err)
	}
	RegisterErrorABI(*parsed)
	return nil
}

// lookupCustomError returns the registered error for the given 4-byte selector.
func lookupCustomError(sel [4]byte) (abi.Error, bool) {
	customErrMu.RLock()
	defer customErrMu.RUnlock()
	e, ok := customErrByID[sel]
	return e, ok
}

// decodeCustomError decodes revert returndata whose leading 4-byte selector matches a registered
// custom error, formatting it as "ErrorName(argName=value, …)" (or "ErrorName()" for a no-arg error,
// or when the arguments cannot be unpacked — the name alone is still useful). It returns ok=false when
// the selector is unknown, so the caller can fall back to the raw hex payload.
func decodeCustomError(data []byte) (string, bool) {
	if len(data) < selectorLength {
		return "", false
	}
	var sel [4]byte
	copy(sel[:], data[:selectorLength])

	e, ok := lookupCustomError(sel)
	if !ok {
		return "", false
	}
	if len(e.Inputs) == 0 {
		return e.Name + "()", true
	}

	vals, err := e.Inputs.Unpack(data[selectorLength:])
	if err != nil {
		return e.Name + "()", true
	}

	parts := make([]string, 0, len(e.Inputs))
	for i, in := range e.Inputs {
		if i >= len(vals) {
			break
		}
		parts = append(parts, fmt.Sprintf("%s=%s", in.Name, formatABIValue(vals[i])))
	}
	return fmt.Sprintf("%s(%s)", e.Name, strings.Join(parts, ", ")), true
}

// formatABIValue renders a decoded ABI argument as a concise, client-safe string.
func formatABIValue(v any) string {
	switch t := v.(type) {
	case common.Address:
		return t.Hex()
	case *big.Int:
		return t.String()
	case []byte:
		return "0x" + hex.EncodeToString(t)
	case [32]byte:
		return "0x" + hex.EncodeToString(t[:])
	default:
		return fmt.Sprintf("%v", t)
	}
}
