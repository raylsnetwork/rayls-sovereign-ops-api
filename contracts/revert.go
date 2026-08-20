package contracts

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SimulateRevertReason re-runs a failed call via eth_call at the given block to recover the revert
// reason. A transaction receipt carries no reason, so re-executing the same call against the block's
// state is the only way to surface why the contract rejected it. It returns a concise, client-safe
// reason, or "" when no reason can be determined (e.g. the call no longer reverts at that block).
func SimulateRevertReason(
	ctx context.Context,
	client *ethclient.Client,
	msg ethereum.CallMsg,
	blockNumber *big.Int,
) string {
	out, err := client.CallContract(ctx, msg, blockNumber)
	if err != nil {
		// Prefer the structured revert data on the error (Geth JSON-RPC `data`): decodes custom
		// errors to a readable name regardless of the error message text.
		if data, ok := ethclient.RevertErrorData(err); ok {
			if reason, known := decodeKnownRevert(data); known {
				return reason
			}
		}
		// Some RPC providers return the revert data as the call's return bytes...
		if len(out) > 0 {
			if reason, decodeErr := decodeRevertReason(out); decodeErr == nil {
				return reason
			}
		}
		// ...others embed it in the error (JSON-RPC `data` field or the message text).
		if reason := tryExtractRevertReason(err); reason != "" {
			return SanitizeRevertForClient(errors.New(reason))
		}
		return SanitizeRevertForClient(err)
	}
	// No error but data present: decode any inline revert payload.
	if len(out) > 0 {
		if reason, decodeErr := decodeRevertReason(out); decodeErr == nil {
			return reason
		}
	}
	return ""
}

// SanitizeRevertForClient returns a concise, user-facing reason string without internal prefixes or
// stack traces.
func SanitizeRevertForClient(err error) string {
	if err == nil {
		return ""
	}
	// Prefer the structured revert data carried by Geth's JSON-RPC error: it holds the exact
	// selector + args, so a custom error decodes to a readable name even when the error message is
	// just "execution reverted" (which the text heuristics below would otherwise return verbatim).
	if data, ok := ethclient.RevertErrorData(err); ok {
		if reason, known := decodeKnownRevert(data); known {
			return reason
		}
	}
	// First, try structured extraction.
	if reason := ExtractRevertReasonFromError(err); reason != "" {
		return stripQuotes(trimReasonPrefixes(singleLine(reason)))
	}
	// Next, search commonly used phrases inside the error message.
	m := singleLine(err.Error())
	keys := []string{
		"execution reverted:",
		"reverted with reason string",
		"contract call reverted:",
		"contract call reverted (raw):",
		"reverted",
		"execution failure",
		"execution error",
	}
	for _, k := range keys {
		if i := strings.Index(m, k); i >= 0 {
			return stripQuotes(trimReasonPrefixes(strings.TrimSpace(m[i+len(k):])))
		}
	}
	// If we can find a hex blob, show that only.
	if hexStr := extractHexDataFromMessage(m); hexStr != "" {
		return hexStr
	}
	// Last resort: return the first line without stack info.
	return stripQuotes(singleLine(m))
}

// ExtractRevertReasonFromError exposes revert reason extraction for callers outside this package.
func ExtractRevertReasonFromError(err error) string {
	if err == nil {
		return ""
	}
	return tryExtractRevertReason(err)
}

// IsRevertError reports whether err represents an on-chain execution revert. It exists for failures
// that surface at gas-estimation time: eth_estimateGas re-executes the call, so a call that will
// revert fails here — before any transaction is signed or sent — and thus produces no receipt with
// status 0 to inspect. Such a failure is a client-correctable condition and must be classified as a
// revert (mapped to 422 upstream), not swallowed as a generic 500.
func IsRevertError(err error) bool {
	if err == nil {
		return false
	}
	// Prefer the structured revert data Geth carries on the JSON-RPC error: it reliably marks an
	// on-chain revert regardless of the message text. Fall back to the string heuristic only when
	// absent, so unrelated RPC/network errors that merely contain "revert" aren't misclassified.
	if _, ok := ethclient.RevertErrorData(err); ok {
		return true
	}
	return strings.Contains(strings.ToLower(err.Error()), "revert")
}

type jsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (err *jsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc error %d", err.Code)
	}
	return err.Message
}

// tryExtractRevertReason attempts to get a human-readable revert reason from various commonly
// observed JSON-RPC error shapes and messages.
func tryExtractRevertReason(txErr error) string {
	// 1) Direct message contains the reason (geth-style): "execution reverted: <reason>".
	msg := txErr.Error()
	if i := strings.Index(msg, "execution reverted"); i >= 0 {
		return msg[i:]
	}
	if i := strings.Index(msg, "reverted with reason string"); i >= 0 {
		return msg[i:]
	}

	// 2) Marshal/unmarshal into our jsonError mirror and inspect the Data field variants.
	var jsonErr jsonError
	if jsonBytes, marshalErr := json.Marshal(txErr); marshalErr == nil {
		_ = json.Unmarshal(jsonBytes, &jsonErr)
	}
	if reason, ok := reasonFromJSONData(jsonErr.Data); ok {
		return withMessagePrefix(jsonErr.Message, reason)
	}

	// 3) Fallbacks: decode any hex in the message; else return the message.
	if reason, ok := decodeHexReason(extractHexDataFromMessage(msg)); ok {
		return withMessagePrefix(jsonErr.Message, reason)
	}
	return jsonErr.Message
}

// reasonFromJSONData inspects a JSON-RPC error `data` field — a hex string, or a nested object
// carrying one under "data"/"originalError" — and returns the decoded reason, or the raw payload
// when it cannot be decoded. The bool reports whether any data payload was found.
func reasonFromJSONData(data any) (string, bool) {
	switch d := data.(type) {
	case string:
		return reasonFromHexField(d)
	case map[string]any:
		return reasonFromDataObject(d)
	}
	return "", false
}

// reasonFromDataObject extracts a revert reason from a JSON-RPC error data object, checking both the
// direct "data" field (a hex string or a nested {data: hex} object) and an "originalError" wrapper.
func reasonFromDataObject(m map[string]any) (string, bool) {
	if v, ok := m["data"]; ok {
		if ds, ok := v.(string); ok {
			if r, found := reasonFromHexField(ds); found {
				return r, true
			}
		}
		if nested, ok := v.(map[string]any); ok {
			if ds, ok := nested["data"].(string); ok {
				if r, found := reasonFromHexField(ds); found {
					return r, true
				}
			}
		}
	}
	if oe, ok := m["originalError"].(map[string]any); ok {
		if ds, ok := oe["data"].(string); ok {
			if r, found := reasonFromHexField(ds); found {
				return r, true
			}
		}
	}
	return "", false
}

// reasonFromHexField decodes a single hex revert payload, falling back to a raw "data=..." marker
// when the bytes are not a recognized revert format. An empty input reports not-found.
func reasonFromHexField(hexStr string) (string, bool) {
	if hexStr == "" {
		return "", false
	}
	if reason, err := decodeRevertReason(common.FromHex(hexStr)); err == nil {
		return reason, true
	}
	return "data=" + hexStr, true
}

// decodeHexReason decodes a hex revert payload into a reason, reporting false when the input is
// empty or the bytes are not a decodable revert format.
func decodeHexReason(hexStr string) (string, bool) {
	if hexStr == "" {
		return "", false
	}
	if reason, err := decodeRevertReason(common.FromHex(hexStr)); err == nil {
		return reason, true
	}
	return "", false
}

// withMessagePrefix joins a JSON-RPC error message with a decoded reason, omitting the prefix when
// the message is empty.
func withMessagePrefix(message, reason string) string {
	if message != "" {
		return message + ", " + reason
	}
	return reason
}

func singleLine(s string) string {
	if idx := strings.IndexByte(s, '\n'); idx >= 0 {
		return strings.TrimSpace(s[:idx])
	}
	return strings.TrimSpace(s)
}

func trimReasonPrefixes(s string) string {
	prefixes := []string{
		"execution reverted:",
		"reverted with reason string",
		"contract call reverted:",
		"contract call reverted (raw):",
		"reverted",
		"execution failure",
		"execution error",
	}
	out := strings.TrimSpace(s)
	for _, p := range prefixes {
		if strings.HasPrefix(strings.ToLower(out), strings.ToLower(p)) {
			out = strings.TrimSpace(out[len(p):])
		}
	}
	return out
}

func stripQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '\'' && s[len(s)-1] == '\'') || (s[0] == '"' && s[len(s)-1] == '"') {
			return strings.TrimSpace(s[1 : len(s)-1])
		}
	}
	return s
}

func extractHexDataFromMessage(msg string) string {
	// Match 0x followed by at least 8 hex chars (selector + ...).
	re := regexp.MustCompile(`0x[0-9a-fA-F]{8,}`)
	loc := re.FindStringIndex(msg)
	if loc == nil {
		return ""
	}
	return msg[loc[0]:loc[1]]
}

// selectorLength is the size in bytes of a Solidity function/error selector.
const selectorLength = 4

func decodeRevertReason(data []byte) (string, error) {
	if len(data) < selectorLength {
		return "", fmt.Errorf("data too short to contain a revert reason")
	}
	if reason, ok := decodeKnownRevert(data); ok {
		return reason, nil
	}
	// If no known format is detected, return the raw data as a hex string.
	return "0x" + hex.EncodeToString(data), nil
}

// decodeKnownRevert decodes revert returndata into a readable reason for the three formats we
// recognize: standard Error(string), Panic(uint256), and registered custom errors. It reports
// ok=false when the payload matches none of them, so callers can fall back to the raw hex.
func decodeKnownRevert(data []byte) (string, bool) {
	if len(data) < selectorLength {
		return "", false
	}

	// Standard revert reason selector (0x08c379a0) encoded "Error(string)".
	revertSelector := []byte{0x08, 0xc3, 0x79, 0xa0}
	// Solidity Panic(uint256) selector (0x4e487b71).
	panicSelector := []byte{0x4e, 0x48, 0x7b, 0x71}

	// Valid ABI-encoded revert data always begins with its selector, so match a prefix (not
	// bytes.Index) — a selector appearing inside a custom error's string/bytes argument must not be
	// decoded from the wrong offset. This mirrors decodeCustomError, which reads the leading 4 bytes.
	if bytes.HasPrefix(data, revertSelector) {
		if reason, err := decodeStandardRevertReason(data[selectorLength:]); err == nil {
			return reason, true
		}
	}
	if bytes.HasPrefix(data, panicSelector) {
		if reason, err := decodePanicReason(data[selectorLength:]); err == nil {
			return reason, true
		}
	}
	// Custom Solidity errors (error Foo(...)) registered by the adapters — decode to "Foo(arg=…)"
	// instead of a bare selector.
	if reason, ok := decodeCustomError(data); ok {
		return reason, true
	}
	return "", false
}

func decodeStandardRevertReason(data []byte) (string, error) {
	stringTy, err := abi.NewType("string", "", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create ABI type: %w", err)
	}
	arguments := abi.Arguments{{Type: stringTy}}

	reason, err := arguments.Unpack(data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack revert reason: %w", err)
	}
	if len(reason) == 0 {
		return "", fmt.Errorf("no revert reason found")
	}

	s, ok := reason[0].(string)
	if !ok {
		return "", fmt.Errorf("revert reason is not a string")
	}
	return s, nil
}

// decodePanicReason decodes Solidity Panic(uint256) reasons.
// See https://docs.soliditylang.org/en/latest/control-structures.html#panic-via-assert-and-error-via-require
func decodePanicReason(data []byte) (string, error) {
	uintTy, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create ABI type: %w", err)
	}
	args := abi.Arguments{{Type: uintTy}}
	vals, err := args.Unpack(data)
	if err != nil || len(vals) == 0 {
		return "", fmt.Errorf("failed to unpack panic code")
	}
	if bi, ok := vals[0].(*big.Int); ok {
		return humanPanicCode(bi.Uint64()), nil
	}
	return "", fmt.Errorf("unexpected panic code type")
}

// Solidity Panic(uint256) codes, per
// https://docs.soliditylang.org/en/latest/control-structures.html#panic-via-assert-and-error-via-require
const (
	panicAssertFalse         = 0x01
	panicArithmeticOverflow  = 0x11
	panicDivideByZero        = 0x12
	panicInvalidEnum         = 0x21
	panicBadStorageEncoding  = 0x22
	panicEmptyArrayPop       = 0x31
	panicArrayOutOfBounds    = 0x32
	panicExcessiveAllocation = 0x41
	panicUninitializedFunc   = 0x51
)

func humanPanicCode(code uint64) string {
	switch code {
	case panicAssertFalse:
		return "Panic(0x01): assert(false)"
	case panicArithmeticOverflow:
		return "Panic(0x11): arithmetic overflow/underflow"
	case panicDivideByZero:
		return "Panic(0x12): division or modulo by zero"
	case panicInvalidEnum:
		return "Panic(0x21): invalid enum value"
	case panicBadStorageEncoding:
		return "Panic(0x22): storage byte array that is incorrectly encoded"
	case panicEmptyArrayPop:
		return "Panic(0x31): pop on an empty array"
	case panicArrayOutOfBounds:
		return "Panic(0x32): out-of-bounds array access"
	case panicExcessiveAllocation:
		return "Panic(0x41): too much memory allocated"
	case panicUninitializedFunc:
		return "Panic(0x51): zero-initialized internal function"
	default:
		return fmt.Sprintf("Panic(0x%x)", code)
	}
}
