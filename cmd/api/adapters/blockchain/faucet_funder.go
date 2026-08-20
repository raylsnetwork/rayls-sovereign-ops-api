package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
)

var _ core.WalletFunder = (*FaucetFunder)(nil)

// defaultFaucetFundWei is the target balance topped up to a new wallet when no explicit
// amount is configured: 5 ETH — enough for several token deploys on a 100 gwei dev chain.
var defaultFaucetFundWei, _ = new(big.Int).SetString("5000000000000000000", 10)

// FaucetFunder funds new custody wallets from a configured faucet key so they can pay gas
// on a non-gasless chain. DEV-ONLY: the faucet private key lives in config, not the HSM.
// Transfers only the shortfall up to the target balance and is a no-op when the wallet is
// already funded. A mutex serializes transfers so concurrent logins don't reuse a nonce.
type FaucetFunder struct {
	client  *ethclient.Client
	key     *ecdsa.PrivateKey
	from    common.Address
	chainID *big.Int
	target  *big.Int
	mu      sync.Mutex
}

// NewFaucetFunder builds a funder from a hex private key (with or without 0x prefix) and an
// optional decimal-wei target ("" → 5 ETH). Reuses an existing RPC client to avoid a second dial.
func NewFaucetFunder(client *ethclient.Client, privateKeyHex, targetWei string) (*FaucetFunder, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(privateKeyHex), "0x"))
	if err != nil {
		return nil, fmt.Errorf("parse faucet private key: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain ID for faucet: %w", err)
	}

	target := defaultFaucetFundWei
	if strings.TrimSpace(targetWei) != "" {
		t, ok := new(big.Int).SetString(strings.TrimSpace(targetWei), 10)
		if !ok {
			return nil, fmt.Errorf("invalid faucet target wei: %q", targetWei)
		}
		target = t
	}

	return &FaucetFunder{
		client:  client,
		key:     key,
		from:    crypto.PubkeyToAddress(key.PublicKey),
		chainID: chainID,
		target:  target,
	}, nil
}

// FundWallet tops up address to the target balance, transferring only the shortfall.
func (f *FaucetFunder) FundWallet(ctx context.Context, address string) error {
	if !common.IsHexAddress(address) {
		return fmt.Errorf("invalid wallet address: %q", address)
	}
	to := common.HexToAddress(address)

	f.mu.Lock()
	defer f.mu.Unlock()

	balance, err := f.client.BalanceAt(ctx, to, nil)
	if err != nil {
		return fmt.Errorf("read wallet balance: %w", err)
	}
	if balance.Cmp(f.target) >= 0 {
		return nil // already funded
	}
	shortfall := new(big.Int).Sub(f.target, balance)

	nonce, err := f.client.PendingNonceAt(ctx, f.from)
	if err != nil {
		return fmt.Errorf("get faucet nonce: %w", err)
	}
	gasPrice, err := f.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("suggest gas price: %w", err)
	}

	tx := types.NewTransaction(nonce, to, shortfall, 21000, gasPrice, nil)
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(f.chainID), f.key)
	if err != nil {
		return fmt.Errorf("sign faucet tx: %w", err)
	}
	if err := f.client.SendTransaction(ctx, signed); err != nil {
		return fmt.Errorf("send faucet tx: %w", err)
	}

	// Wait briefly for inclusion so the balance is usable on the immediately-following deploy.
	receiptCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for {
		receipt, err := f.client.TransactionReceipt(receiptCtx, signed.Hash())
		if err == nil {
			if receipt.Status == 0 {
				return fmt.Errorf("faucet tx reverted (tx %s)", signed.Hash().Hex())
			}
			return nil
		}
		select {
		case <-receiptCtx.Done():
			return fmt.Errorf("timed out waiting for faucet tx %s: %w", signed.Hash().Hex(), receiptCtx.Err())
		case <-time.After(time.Second):
		}
	}
}
