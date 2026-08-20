package custody

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newHSMStub serves the three endpoints SignAndTransact touches (auth, wallet lookup, transaction)
// and captures the decoded transaction request body for assertions.
func newHSMStub(t *testing.T, captured *map[string]any) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api/auth":
			_, err := w.Write([]byte(`{"Token":"stub-token"}`))
			require.NoError(t, err)
		case r.URL.Path == "/api/transaction":
			require.NoError(t, json.NewDecoder(r.Body).Decode(captured))
			_, err := w.Write([]byte(`{"TxHash":"0xdeadbeef"}`))
			require.NoError(t, err)
		default: // /api/wallet/address/{addr}
			_, err := w.Write([]byte(`{"id":"wallet-1"}`))
			require.NoError(t, err)
		}
	}))
}

// signedPayload builds the RLP-encoded unsigned tx that callers hand to SignAndTransact.
func signedPayload(t *testing.T) []byte {
	t.Helper()

	tx := types.NewTransaction(
		7,
		common.HexToAddress("0x00000000000000000000000000000000000000aa"),
		big.NewInt(0),
		21000,
		big.NewInt(1),
		[]byte{0x01, 0x02},
	)
	payload, err := tx.MarshalBinary()
	require.NoError(t, err)
	return payload
}

func TestRaylsHSM_SignAndTransact_SendsChainIDAndRPCURL(t *testing.T) {
	// Custody serves chains created at runtime, so each signing request must name its target
	// chain rather than relying on custody's process-wide config.
	var body map[string]any
	server := newHSMStub(t, &body)
	defer server.Close()

	hsm := NewRaylsHSM(server.URL, "api-key", "pw", "http://chain.local:8545")

	txHash, err := hsm.SignAndTransact(context.Background(), signedPayload(t), "0xabc", "600042")

	require.NoError(t, err)
	assert.Equal(t, "0xdeadbeef", txHash)
	assert.Equal(t, "600042", body["ChainId"])
	assert.Equal(t, "http://chain.local:8545", body["RpcUrl"])
}

func TestRaylsHSM_SignAndTransact_OmitsRoutingFieldsWhenUnset(t *testing.T) {
	// With no chain ID or RPC URL to send, both fields are omitted so custody falls back to its
	// own configuration instead of receiving empty strings it would reject.
	var body map[string]any
	server := newHSMStub(t, &body)
	defer server.Close()

	hsm := NewRaylsHSM(server.URL, "api-key", "pw", "")

	_, err := hsm.SignAndTransact(context.Background(), signedPayload(t), "0xabc", "")

	require.NoError(t, err)
	assert.NotContains(t, body, "ChainId")
	assert.NotContains(t, body, "RpcUrl")
}
