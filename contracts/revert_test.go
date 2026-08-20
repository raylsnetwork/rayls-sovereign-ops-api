package contracts

import (
	"encoding/hex"
	"errors"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSanitizeRevertForClient_CommonPhrases(t *testing.T) {
	// SanitizeRevertForClient strips RPC/revert prefixes and quotes to leave a bare, client-safe reason.
	cases := []struct {
		name string
		err  error
		exp  string
	}{
		{"exec_reverted_with_reason", errors.New("execution reverted: not allowed"), "not allowed"},
		{"reverted_with_reason_string", errors.New("reverted with reason string 'bad'"), "bad"},
		{"contract_call_reverted_prefixed", errors.New("contract call reverted: already exists"), "already exists"},
		{
			"raw_hex_only",
			errors.New("contract call reverted (raw): 0x08c379a000000000000000"),
			"0x08c379a000000000000000",
		},
		{
			"jsonrpc_with_hex",
			errors.New("rpc error: code = Unknown desc = execution reverted: 0x08c379a00000ff"),
			"0x08c379a00000ff",
		},
		{
			"json_message_reason_then_hex",
			errors.New("rpc error: execution reverted: insufficient balance (data: 0xdeadbeef)"),
			"insufficient balance (data: 0xdeadbeef)",
		},
		{"unknown_message", errors.New("some unrelated network error"), "some unrelated network error"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeRevertForClient(tt.err)
			assert.Equal(t, tt.exp, got)
		})
	}
}

func TestExtractHexDataFromMessage(t *testing.T) {
	// extractHexDataFromMessage returns the first 0x-prefixed hex blob (selector or longer), or "".
	cases := []struct{ in, exp string }{
		{"no hex here", ""},
		{"something 0xdeadBEEF more", "0xdeadBEEF"},
		{"0x08c379a0 message first", "0x08c379a0"},
	}
	for _, c := range cases {
		assert.Equal(t, c.exp, extractHexDataFromMessage(c.in))
	}
}

func TestDecodeRevertReason_StandardErrorString(t *testing.T) {
	// A standard Error(string) revert payload decodes to its human-readable reason.
	// Selector 0x08c379a0 + ABI-encoded string "boom".
	data := common.FromHex(
		"08c379a0" +
			"0000000000000000000000000000000000000000000000000000000000000020" +
			"0000000000000000000000000000000000000000000000000000000000000004" +
			"626f6f6d00000000000000000000000000000000000000000000000000000000",
	)

	reason, err := decodeRevertReason(data)

	assert.NoError(t, err)
	assert.Equal(t, "boom", reason)
}

func TestDecodeRevertReason_Panic(t *testing.T) {
	// A Panic(uint256) payload decodes to a human-readable panic description.
	// Selector 0x4e487b71 + panic code 0x11 (arithmetic overflow/underflow).
	data := common.FromHex(
		"4e487b71" +
			"0000000000000000000000000000000000000000000000000000000000000011",
	)

	reason, err := decodeRevertReason(data)

	assert.NoError(t, err)
	assert.Equal(t, "Panic(0x11): arithmetic overflow/underflow", reason)
}

func TestSanitizeRevertForClient_NilError(t *testing.T) {
	// A nil error yields an empty string.
	assert.Equal(t, "", SanitizeRevertForClient(nil))
}

func TestDecodeRevertReason_UnknownFormatReturnsHex(t *testing.T) {
	// Bytes matching no known selector are returned as a 0x-prefixed hex string.
	reason, err := decodeRevertReason([]byte{0xde, 0xad, 0xbe, 0xef})

	assert.NoError(t, err)
	assert.Equal(t, "0xdeadbeef", reason)
}

func TestDecodeRevertReason_TooShort(t *testing.T) {
	// A payload shorter than a selector cannot contain a reason and errors.
	_, err := decodeRevertReason([]byte{0x01, 0x02})

	assert.Error(t, err)
}

func TestHumanPanicCode_UnknownCode(t *testing.T) {
	// An unrecognized panic code falls back to a generic hex-formatted label.
	assert.Equal(t, "Panic(0x99)", humanPanicCode(0x99))
}

func TestDecodeRevertReason_SelectorInArgNotMisdecoded(t *testing.T) {
	// A custom error carrying a bytes argument that embeds a valid Error(string) encoding must decode
	// to the outer custom error, not the embedded string — the selector is matched as a prefix, not
	// anywhere in the payload.
	const boomABI = `[{"type":"error","name":"Boom","inputs":[{"name":"payload","type":"bytes"}]}]`
	parsed, err := abi.JSON(strings.NewReader(boomABI))
	require.NoError(t, err)
	RegisterErrorABI(parsed)

	stringTy, err := abi.NewType("string", "", nil)
	require.NoError(t, err)
	embeddedStr, err := abi.Arguments{{Type: stringTy}}.Pack("gotcha")
	require.NoError(t, err)
	embedded := append([]byte{0x08, 0xc3, 0x79, 0xa0}, embeddedStr...)

	boom := parsed.Errors["Boom"]
	packed, err := boom.Inputs.Pack(embedded)
	require.NoError(t, err)
	data := append(append([]byte{}, boom.ID[:selectorLength]...), packed...)

	reason, err := decodeRevertReason(data)

	require.NoError(t, err)
	assert.Equal(t, "Boom(payload=0x"+hex.EncodeToString(embedded)+")", reason)
	assert.NotEqual(t, "gotcha", reason)
}

func TestIsRevertError_TrueForStructuredData(t *testing.T) {
	// An error carrying structured JSON-RPC revert data is a revert even when the message lacks "revert".
	err := &dataError{msg: "execution failed", data: "0x08c379a0"}

	assert.True(t, IsRevertError(err))
}

func TestIsRevertError_FalseForPlainNetworkError(t *testing.T) {
	// A generic network error without structured data and without the word "revert" is not a revert.
	assert.False(t, IsRevertError(errors.New("connection reset by peer")))
}

func TestIsRevertError_TrueForMessageHeuristicFallback(t *testing.T) {
	// The string heuristic still classifies a plain error mentioning "reverted" as a revert.
	assert.True(t, IsRevertError(errors.New("execution reverted: bad state")))
}
