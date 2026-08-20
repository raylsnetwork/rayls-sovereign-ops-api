package services

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// Function name constants for the protocol token functions whose access is gated by the
// Access Manager. Used both for the friendly name in responses and to derive capability flags.
const (
	fnMint              = "mint"
	fnBurn              = "burn"
	fnSubmitTokenUpdate = "submitTokenUpdate"
	fnSetSwapValidity   = "setSwapValidityTime"
)

// tokenFunctionSignatures maps a canonical Solidity signature to a friendly function name.
// Selectors are derived as the first 4 bytes of keccak256(signature). Several signatures
// (overloads across token standards) collapse to the same friendly name.
//
// Source: RaylsErc20/721/1155/EnygmaHandler in rayls-privacy-contracts (restricted functions).
var tokenFunctionSignatures = map[string]string{
	// OWNER-gated (the user-relevant actions)
	"mint(address,uint256)":                    fnMint,
	"mint(address,uint256,uint256,bytes)":      fnMint,
	"burn(address,uint256)":                    fnBurn,
	"burn(uint256)":                            fnBurn,
	"burn(address,uint256,uint256)":            fnBurn,
	"submitTokenUpdate(uint8,uint256)":         fnSubmitTokenUpdate,
	"submitTokenUpdate(uint8,uint256,uint256)": fnSubmitTokenUpdate,
	"setSwapValidityTime(uint64)":              fnSetSwapValidity,

	// MESSAGE_EXECUTOR / RELAYER-gated (protocol-internal — included only so any callable
	// selector resolves to a readable name; a regular user won't hold these roles).
	"receiveTeleport(address,uint256)":                "receiveTeleport",
	"receiveTeleportAtomic(address,uint256)":          "receiveTeleportAtomic",
	"receiveTeleportFromPublicChain(address,uint256)": "receiveTeleportFromPublicChain",
	"revertTeleportMint(address,uint256)":             "revertTeleportMint",
	"revertTeleportBurn(address,uint256)":             "revertTeleportBurn",
	"revertTeleportToPublicChain(address,uint256)":    "revertTeleportToPublicChain",
	"receiveResourceId(bytes32)":                      "receiveResourceId",
	"unlock(address,uint256)":                         "unlock",
}

// selectorToName is the derived selector(0x + 8 hex) -> friendly name map, built once at init.
var selectorToName = buildSelectorToName(tokenFunctionSignatures)

func buildSelectorToName(sigs map[string]string) map[string]string {
	out := make(map[string]string, len(sigs))
	for sig, name := range sigs {
		out[functionSelector(sig)] = name
	}
	return out
}

// functionSelector returns the 4-byte selector ("0x" + 8 lowercase hex) for a canonical signature.
func functionSelector(signature string) string {
	sum := crypto.Keccak256([]byte(signature))
	return fmt.Sprintf("0x%02x%02x%02x%02x", sum[0], sum[1], sum[2], sum[3])
}

// tokenFunctionName returns the friendly name for a 4-byte selector, or "" if unknown.
func tokenFunctionName(selector string) string {
	return selectorToName[strings.ToLower(strings.TrimSpace(selector))]
}
