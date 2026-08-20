package contracts

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/PNTokenRegistryV1"
)

// registerRegistryErrors parses and registers the PNTokenRegistryV1 ABI's custom errors and returns
// the parsed ABI so tests can build revert payloads from it.
func registerRegistryErrors(t *testing.T) *abi.ABI {
	t.Helper()
	parsed, err := PNTokenRegistryV1.PNTokenRegistryV1MetaData.ParseABI()
	require.NoError(t, err)
	RegisterErrorABI(*parsed)
	return parsed
}

// customErrorPayload builds revert returndata (4-byte selector + packed args) for the named error.
func customErrorPayload(t *testing.T, a *abi.ABI, name string, args ...any) []byte {
	t.Helper()
	e, ok := a.Errors[name]
	require.True(t, ok, "error %q not found in ABI", name)
	packed, err := e.Inputs.Pack(args...)
	require.NoError(t, err)
	return append(append([]byte{}, e.ID[:selectorLength]...), packed...)
}

func TestDecodeRevertReason_CustomErrorWithAddressAndEnum(t *testing.T) {
	// A registered custom error with an address + uint8 decodes to "Name(arg=value, …)".
	a := registerRegistryErrors(t)
	token := common.HexToAddress("0x00000000000000000000000000000000000000aa")

	data := customErrorPayload(t, a, "RaylsAppV1__PrivacyNodeNotActive", token, uint8(1))
	reason, err := decodeRevertReason(data)

	require.NoError(t, err)
	assert.Equal(t, "RaylsAppV1__PrivacyNodeNotActive(tokenAddress="+token.Hex()+", privacyNodeStatus=1)", reason)
}

func TestDecodeRevertReason_CustomErrorAddressArg(t *testing.T) {
	// The access-manager unauthorized error decodes with the caller address rendered as hex.
	a := registerRegistryErrors(t)
	caller := common.HexToAddress("0x00000000000000000000000000000000000000bb")

	data := customErrorPayload(t, a, "RaylsAccessManaged__Unauthorized", caller)
	reason, err := decodeRevertReason(data)

	require.NoError(t, err)
	assert.Equal(t, "RaylsAccessManaged__Unauthorized(caller="+caller.Hex()+")", reason)
}

func TestDecodeRevertReason_CustomErrorNoArgs(t *testing.T) {
	// A no-argument custom error decodes to "Name()".
	a := registerRegistryErrors(t)

	data := customErrorPayload(t, a, "TokenRegistryV1__TokenCoreNotConfigured")
	reason, err := decodeRevertReason(data)

	require.NoError(t, err)
	assert.Equal(t, "TokenRegistryV1__TokenCoreNotConfigured()", reason)
}

func TestDecodeRevertReason_UnknownSelectorFallsBackToHex(t *testing.T) {
	// An unregistered selector is returned as raw hex, unchanged from prior behavior.
	data := []byte{0xde, 0xad, 0xbe, 0xef}

	reason, err := decodeRevertReason(data)

	require.NoError(t, err)
	assert.Equal(t, "0xdeadbeef", reason)
}

func TestDecodeRevertReason_StandardErrorStringStillDecodes(t *testing.T) {
	// Adding custom-error decoding does not regress the standard Error(string) path.
	stringTy, err := abi.NewType("string", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: stringTy}}.Pack("boom")
	require.NoError(t, err)
	data := append([]byte{0x08, 0xc3, 0x79, 0xa0}, packed...)

	reason, err := decodeRevertReason(data)

	require.NoError(t, err)
	assert.Equal(t, "boom", reason)
}

func TestDecodeRevertReason_PanicStillDecodes(t *testing.T) {
	// The Panic(uint256) path is likewise preserved.
	uintTy, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: uintTy}}.Pack(big.NewInt(0x11))
	require.NoError(t, err)
	data := append([]byte{0x4e, 0x48, 0x7b, 0x71}, packed...)

	reason, err := decodeRevertReason(data)

	require.NoError(t, err)
	assert.Equal(t, "Panic(0x11): arithmetic overflow/underflow", reason)
}

// dataError implements geth's rpc.Error + rpc.DataError so ethclient.RevertErrorData can extract the
// structured revert payload — mirroring an eth_estimateGas revert on a Geth backend.
type dataError struct {
	msg  string
	data string
}

func (e *dataError) Error() string  { return e.msg }
func (e *dataError) ErrorCode() int { return 3 }
func (e *dataError) ErrorData() any { return e.data }

func TestSanitizeRevertForClient_DecodesStructuredCustomError(t *testing.T) {
	// The estimate-gas path decodes the custom error from the JSON-RPC error data even though the
	// error message is only "execution reverted".
	a := registerRegistryErrors(t)
	caller := common.HexToAddress("0x00000000000000000000000000000000000000cc")
	payload := customErrorPayload(t, a, "RaylsAccessManaged__Unauthorized", caller)
	err := &dataError{msg: "execution reverted", data: "0x" + hex.EncodeToString(payload)}

	reason := SanitizeRevertForClient(err)

	assert.Equal(t, "RaylsAccessManaged__Unauthorized(caller="+caller.Hex()+")", reason)
}
