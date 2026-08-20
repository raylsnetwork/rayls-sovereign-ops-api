package custody

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
)

const hsmRequestTimeout = 30 * time.Second

var _ core.CustodyService = (*RaylsHSM)(nil)

// RaylsHSM calls the Rayls HSM custody service.
//
// Auth flow: POST /api/auth {ApiKey} → Bearer token cached in memory.
// On 401, re-authenticates once and retries.
type RaylsHSM struct {
	baseURL string
	apiKey  string
	// TODO: move to a secrets manager (vault, AWS Secrets Manager, etc.)
	password string
	// rpcURL is the JSON-RPC endpoint of the chain this instance targets, sent with every
	// signing request. Custody is shared across chains that are created at runtime, so it
	// cannot resolve the endpoint from its own process-wide config — see SignAndTransact.
	rpcURL string

	httpClient *http.Client
	mu         sync.Mutex
	token      string
}

func NewRaylsHSM(baseURL, apiKey, password, rpcURL string) *RaylsHSM {
	return &RaylsHSM{
		baseURL:    baseURL,
		apiKey:     apiKey,
		password:   password,
		rpcURL:     rpcURL,
		httpClient: &http.Client{Timeout: hsmRequestTimeout},
	}
}

// --- auth ---

type hsmAuthRequest struct {
	ApiKey string `json:"ApiKey"`
}

type hsmAuthResponse struct {
	Token string `json:"Token"`
}

func (r *RaylsHSM) authenticate(ctx context.Context) error {
	body, _ := json.Marshal(hsmAuthRequest{ApiKey: r.apiKey})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.baseURL+"/api/auth", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build auth request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("auth request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("custody auth returned %d: %s", resp.StatusCode, string(respBody))
	}

	var result hsmAuthResponse
	if err := json.Unmarshal(respBody, &result); err != nil || result.Token == "" {
		return fmt.Errorf("decode custody auth response: %w", err)
	}

	r.mu.Lock()
	r.token = result.Token
	r.mu.Unlock()
	return nil
}

func (r *RaylsHSM) bearerToken(ctx context.Context) (string, error) {
	r.mu.Lock()
	t := r.token
	r.mu.Unlock()
	if t != "" {
		return t, nil
	}
	if err := r.authenticate(ctx); err != nil {
		return "", err
	}
	r.mu.Lock()
	t = r.token
	r.mu.Unlock()
	return t, nil
}

// doJSON executes a request with Bearer auth, re-authenticating once on 401.
func (r *RaylsHSM) doJSON(ctx context.Context, method, path string, reqBody, respDst any) (int, error) {
	for attempt := 0; attempt < 2; attempt++ {
		token, err := r.bearerToken(ctx)
		if err != nil {
			return 0, err
		}

		var bodyReader io.Reader
		if reqBody != nil {
			b, err := json.Marshal(reqBody)
			if err != nil {
				return 0, fmt.Errorf("marshal request: %w", err)
			}
			bodyReader = bytes.NewReader(b)
		}

		req, err := http.NewRequestWithContext(ctx, method, r.baseURL+path, bodyReader)
		if err != nil {
			return 0, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := r.httpClient.Do(req)
		if err != nil {
			return 0, fmt.Errorf("request failed: %w", err)
		}
		respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
		resp.Body.Close()

		if resp.StatusCode == http.StatusUnauthorized && attempt == 0 {
			r.mu.Lock()
			r.token = ""
			r.mu.Unlock()
			continue
		}

		// Skip decoding when there is nothing to decode (e.g. 204 No Content with an empty body),
		// so callers see the status-based error (e.g. "wallet not found") instead of a JSON decode error.
		if respDst != nil && resp.StatusCode != http.StatusNoContent && len(respBody) > 0 {
			if err := json.Unmarshal(respBody, respDst); err != nil {
				return resp.StatusCode, fmt.Errorf("decode response (status %d): %w", resp.StatusCode, err)
			}
		}
		return resp.StatusCode, nil
	}
	return 0, fmt.Errorf("custody: failed after re-authentication")
}

// --- CreateWallet ---

type hsmCreateWalletRequest struct {
	Type            string `json:"type"`
	Password        any    `json:"password"`
	AddressQuantity int    `json:"address_quantity"`
}

type hsmWalletResponse struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

func (r *RaylsHSM) CreateWallet(ctx context.Context, _ uuid.UUID) (address, externalID string, err error) {
	var result []hsmWalletResponse
	status, err := r.doJSON(ctx, http.MethodPost, "/api/wallet", hsmCreateWalletRequest{
		Type:            "KEYSTORE_V3",
		Password:        r.password,
		AddressQuantity: 1,
	}, &result)
	if err != nil {
		return "", "", fmt.Errorf("custody CreateWallet: %w", err)
	}
	if status != http.StatusOK {
		return "", "", fmt.Errorf("custody CreateWallet returned %d", status)
	}
	if len(result) == 0 || result[0].Address == "" {
		return "", "", fmt.Errorf("custody CreateWallet returned empty wallet list")
	}
	return result[0].Address, result[0].ID, nil
}

// --- SignAndTransact ---

type hsmGetWalletByAddressResponse struct {
	ID string `json:"id"`
}

type hsmTransactionRequest struct {
	WalletId    string         `json:"WalletId"`
	Password    any            `json:"Password"`
	Transaction hsmEvmTxFields `json:"Transaction"`
	ChainId     string         `json:"ChainId,omitempty"`
	RpcUrl      string         `json:"RpcUrl,omitempty"`
}

type hsmEvmTxFields struct {
	To    any    `json:"To"`
	Data  string `json:"Data"`
	Value string `json:"Value"`
	Nonce any    `json:"Nonce"`
}

type hsmTransactionResponse struct {
	TxHash string `json:"TxHash"`
}

func (r *RaylsHSM) SignAndTransact(ctx context.Context, payload []byte, signerAddress, chainID string) (string, error) {
	// Decode the RLP-encoded unsigned tx the caller built. We split it back
	// into (To, Data, Value, Nonce) because the custody HSM REST API rebuilds
	// the tx itself from those fields — passing the RLP blob as opaque `Data`
	// with To=null makes custody attempt a CREATE on the RLP bytes, and the
	// chain rejects with OpcodeNotFound (the leading RLP list byte 0xf8/0xf9
	// isn't a valid EVM opcode).
	var tx types.Transaction
	if err := tx.UnmarshalBinary(payload); err != nil {
		return "", fmt.Errorf("decode RLP tx payload: %w", err)
	}

	toStr := ""
	if tx.To() != nil {
		toStr = tx.To().Hex()
	}
	// Custody parses Value with BigInteger.Parse (decimal). Big.Int.String()
	// emits decimal; nil falls back to "0".
	valueStr := "0"
	if v := tx.Value(); v != nil {
		valueStr = v.String()
	}

	// Resolve address → internal wallet ID.
	// The HSM indexes wallets by their EIP-55 checksum address (the form CreateWallet returned),
	// and its /api/wallet/address/{addr} lookup is case-sensitive. ops-api stores/returns addresses
	// lowercased (domain.NormalizeAddress), so the inbound signerAddress is lowercase — canonicalize
	// it to checksum here or the lookup 204s. common.HexToAddress accepts any case.
	signerAddress = common.HexToAddress(signerAddress).Hex()

	var wallet hsmGetWalletByAddressResponse
	status, err := r.doJSON(ctx, http.MethodGet, "/api/wallet/address/"+signerAddress, nil, &wallet)
	if err != nil {
		return "", fmt.Errorf("custody lookup wallet by address: %w", err)
	}
	if status != http.StatusOK || wallet.ID == "" {
		return "", fmt.Errorf("custody wallet not found for address %s (status %d)", signerAddress, status)
	}

	// ChainId/RpcUrl name the target chain per request. Custody serves chains that are created
	// at runtime, so its process-wide EVM_JSON_RPC_CHAIN_ID / EVM_JSON_RPC are empty whenever it
	// booted before any chain existed — and a single pair could not serve more than one chain
	// anyway. The chain ID must match the target chain or EIP-155 replay protection rejects the
	// signature; custody falls back to its own config only when these are omitted.
	var result hsmTransactionResponse
	status, err = r.doJSON(ctx, http.MethodPost, "/api/transaction", hsmTransactionRequest{
		WalletId: wallet.ID,
		Password: r.password,
		Transaction: hsmEvmTxFields{
			To:    toStr,
			Data:  "0x" + hex.EncodeToString(tx.Data()),
			Value: valueStr,
			Nonce: strconv.FormatUint(tx.Nonce(), 10),
		},
		ChainId: chainID,
		RpcUrl:  r.rpcURL,
	}, &result)
	if err != nil {
		return "", fmt.Errorf("custody SignAndTransact: %w", err)
	}
	if status != http.StatusOK {
		return "", fmt.Errorf("custody SignAndTransact returned %d", status)
	}
	if result.TxHash == "" {
		return "", fmt.Errorf("custody SignAndTransact returned empty tx hash")
	}
	return result.TxHash, nil
}
