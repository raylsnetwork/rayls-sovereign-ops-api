package indexer

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// VerifyChainID asserts that the chain reachable at rpcURL actually reports the
// chain ID this instance is configured to index (BLOCKCHAIN_CHAIN_ID).
//
// Why this exists: the Blockscout backfill/listener ingest EVERY token in the
// Blockscout DB and stamp each with the configured ChainID as its issuer_id. That
// stamping is trusted blindly — nothing checks that the Blockscout DB (and the
// chain behind it) is really the configured chain. If an instance is pointed at
// another environment's infrastructure (e.g. a fresh testnet ops-api still wired
// to the devnet Blockscout/RPC), tokens from that other network get relabelled
// with THIS instance's chain ID and surface as if they were deployed here.
//
// A mismatch is a hard configuration error: return it so the caller can refuse to
// backfill rather than silently poison the token list with another chain's data.
//
// expectedChainID is the raw BLOCKCHAIN_CHAIN_ID value. It is optional (empty means
// "no chain scoping configured"), in which case verification is skipped — the caller
// already tolerates unscoped tokens.
func VerifyChainID(ctx context.Context, rpcURL, expectedChainID string, log logger.Logger) error {
	expected := strings.TrimSpace(expectedChainID)
	if expected == "" {
		log.Warn("BLOCKCHAIN_CHAIN_ID is not set — skipping chain-ID verification; " +
			"ingested tokens will not be scoped to a chain")
		return nil
	}

	want, ok := new(big.Int).SetString(expected, 10)
	if !ok {
		return fmt.Errorf("BLOCKCHAIN_CHAIN_ID %q is not a valid integer", expectedChainID)
	}

	if strings.TrimSpace(rpcURL) == "" {
		return fmt.Errorf("BLOCKCHAIN_CHAIN_ID is set to %s but BLOCKCHAIN_RPC_URL is empty, "+
			"so the chain cannot be verified — refusing to backfill against an unverified chain", expected)
	}

	// Retry transient dial/read failures. A hostname like a Cloudflare-fronted RPC
	// resolves to several edge IPs; under a VPN some edges refuse the connection while
	// others work, and the round-robin DNS may hand back a bad one first. Each attempt
	// re-dials (re-resolving DNS), so a few tries reliably land on a reachable edge.
	// A chain-ID MISMATCH is a config error, not transient — return it immediately.
	const attempts = 8
	const retryDelay = 3 * time.Second

	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		got, err := readChainID(ctx, rpcURL)
		if err == nil {
			if got.Cmp(want) != 0 {
				return fmt.Errorf(
					"chain-ID mismatch: configured BLOCKCHAIN_CHAIN_ID=%s but the RPC at %s reports chain ID %s — "+
						"this instance is wired to the wrong network's infrastructure; refusing to ingest its tokens",
					want, rpcURL, got,
				)
			}
			log.Info("Chain-ID verified", "chain_id", want.String())
			return nil
		}

		lastErr = err
		log.Warn("chain-ID verification attempt failed; retrying",
			"attempt", fmt.Sprintf("%d/%d", attempt, attempts), "error", err.Error())
		if attempt < attempts {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(retryDelay):
			}
		}
	}

	return fmt.Errorf("read chain ID from RPC for verification after %d attempts: %w", attempts, lastErr)
}

// readChainID dials the RPC fresh (so DNS is re-resolved each call) and reads the
// chain ID, with a per-attempt timeout.
func readChainID(ctx context.Context, rpcURL string) (*big.Int, error) {
	dialCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(dialCtx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("dial chain RPC for chain-ID verification: %w", err)
	}
	defer client.Close()

	got, err := client.ChainID(dialCtx)
	if err != nil {
		return nil, fmt.Errorf("read chain ID from RPC for verification: %w", err)
	}
	return got, nil
}
