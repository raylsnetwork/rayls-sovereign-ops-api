package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionSelector_KnownValues(t *testing.T) {
	// Canonical selectors must match keccak256(signature)[:4].
	assert.Equal(t, "0x40c10f19", functionSelector("mint(address,uint256)"))
	assert.Equal(t, "0x9dc29fac", functionSelector("burn(address,uint256)"))
}

func TestTokenFunctionName(t *testing.T) {
	// Selectors resolve to friendly names (case-insensitive); unknown selectors return "".
	assert.Equal(t, fnMint, tokenFunctionName("0x40c10f19"))
	assert.Equal(t, fnMint, tokenFunctionName("0x40C10F19"))
	assert.Equal(t, fnBurn, tokenFunctionName("0x9dc29fac"))
	assert.Equal(t, "", tokenFunctionName("0xdeadbeef"))
}
