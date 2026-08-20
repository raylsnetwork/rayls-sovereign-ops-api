package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseErcStandard_MapsRaylsFetcherStrings(t *testing.T) {
	// The Blockscout RaylsTokenDiscovery fetcher writes "Rayls-*" strings into tokens.type;
	// these are the authoritative classification and MUST resolve to the matching standard.
	assert.Equal(t, ErcStandardERC20, ParseErcStandard("Rayls-ERC-20"))
	assert.Equal(t, ErcStandardERC721, ParseErcStandard("Rayls-ERC-721"))
	assert.Equal(t, ErcStandardERC1155, ParseErcStandard("Rayls-ERC-1155"))
	assert.Equal(t, ErcStandardEnygma, ParseErcStandard("Rayls-Enygma"))
	assert.Equal(t, ErcStandardZkDvpERC721, ParseErcStandard("Rayls-ERC-721-DVP"))
	assert.Equal(t, ErcStandardZkDvpERC1155, ParseErcStandard("Rayls-ERC-1155-DVP"))
	assert.Equal(t, ErcStandardStableCoin, ParseErcStandard("Rayls-StableCoin"))
}

func TestParseErcStandard_MapsPlainErcStrings(t *testing.T) {
	// Non-Rayls tokens the fetcher never overwrote keep their plain ERC type.
	assert.Equal(t, ErcStandardERC20, ParseErcStandard("ERC-20"))
	assert.Equal(t, ErcStandardERC721, ParseErcStandard("ERC-721"))
	assert.Equal(t, ErcStandardERC1155, ParseErcStandard("ERC-1155"))
}

func TestParseErcStandard_MapsCanonicalLabels(t *testing.T) {
	// An API ?type= filter can use the same labels the responses expose.
	assert.Equal(t, ErcStandardERC20, ParseErcStandard("RAYLS_ERC20"))
	assert.Equal(t, ErcStandardEnygma, ParseErcStandard("RAYLS_ENYGMA"))
	assert.Equal(t, ErcStandardZkDvpERC1155, ParseErcStandard("RAYLS_ERC1155_DVP"))
	assert.Equal(t, ErcStandardStableCoin, ParseErcStandard("RAYLS_STABLECOIN"))
}

func TestErcStandard_StableCoin_LabelRoundTrips(t *testing.T) {
	// The new stablecoin standard labels as RAYLS_STABLECOIN and parses back to itself.
	assert.Equal(t, "RAYLS_STABLECOIN", ErcStandardStableCoin.Label())
	assert.Equal(t, ErcStandardStableCoin, ParseErcStandard(ErcStandardStableCoin.Label()))
}

func TestParseErcStandard_UnknownMapsToCustom(t *testing.T) {
	// An unrecognised type string falls back to Custom rather than erroring.
	assert.Equal(t, ErcStandardCustom, ParseErcStandard("Dogecoin"))
	assert.Equal(t, ErcStandardCustom, ParseErcStandard(""))
}

func TestParseErcStandard_RoundTripsWithLabel(t *testing.T) {
	// Every standard's canonical Label() parses back to that same standard.
	for _, std := range []ErcStandard{
		ErcStandardERC20, ErcStandardERC721, ErcStandardERC1155,
		ErcStandardEnygma, ErcStandardZkDvpERC721, ErcStandardZkDvpERC1155,
		ErcStandardStableCoin,
	} {
		assert.Equal(t, std, ParseErcStandard(std.Label()), "round-trip failed for %d", std)
	}
}
