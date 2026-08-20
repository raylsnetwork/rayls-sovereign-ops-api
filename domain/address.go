package domain

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// NormalizeAddress canonicalizes an EVM address for storage and lookup: trimmed,
// lowercased, and 0x-prefixed. Using a single canonical form avoids duplicate rows when
// different sources report the same address in different casings (e.g. a checksummed
// address from a deploy vs a lowercase address from the Blockscout indexer).
func NormalizeAddress(addr string) string {
	a := strings.ToLower(strings.TrimSpace(addr))
	a = strings.TrimPrefix(a, "0x")
	return "0x" + a
}

// ChecksumAddress returns the EIP-55 mixed-case checksummed form of an EVM address for
// display in API responses. Storage/lookup keep the lowercase form (see NormalizeAddress).
// An empty input returns an empty string rather than the zero address.
func ChecksumAddress(addr string) string {
	if strings.TrimSpace(addr) == "" {
		return ""
	}
	return common.HexToAddress(addr).Hex()
}
