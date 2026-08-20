package indexer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

const (
	subjectWalletBalanceUpdated = "ops.wallet_balances.updated"
	StreamWalletBalances        = "WALLET_BALANCE_EVENTS"

	// SubjectWalletBalanceSSE is the core-NATS subject for the live wallet-balance event
	// consumed by the API and fanned out to SSE clients. It sits outside the
	// WALLET_BALANCE_EVENTS stream filter (ops.<inst>.wallet_balances.>) so it is not
	// persisted in the WorkQueue stream.
	SubjectWalletBalanceSSE = "ops.sse.wallet_balances"
)

// WalletBalanceEvent is the curated payload pushed for live SSE fan-out and for
// the durable JetStream event. The worker owns this shape — the frontend reads
// it directly.
type WalletBalanceEvent struct {
	Type          string `json:"type"` // "balance_updated"
	WalletAddress string `json:"walletAddress"`
	TokenAddress  string `json:"tokenAddress"`
	Balance       string `json:"balance"`
	BlockNumber   uint64 `json:"blockNumber"`
}

type balanceChangePayload struct {
	Op                       string `json:"op"`
	AddressHash              string `json:"address_hash"`
	TokenContractAddressHash string `json:"token_contract_address_hash"`
	Value                    string `json:"value"`
	BlockNumber              int64  `json:"block_number"`
}

// BlockscoutBalancesListener subscribes to pg_notify('balance_change') on the
// Blockscout database and upserts matching balances into wallet_balances. Only
// balances for addresses present in user_wallets are stored or broadcast.
// Reconnects with exponential backoff (1s → 30s) on disconnect.
type BlockscoutBalancesListener struct {
	connStr       string
	balanceRepo   core.WalletBalanceRepository
	walletRepo    core.UserWalletRepository
	publisher     core.EventPublisher
	livePublisher core.LiveEventPublisher
	log           logger.Logger
	instance      string
}

func NewBlockscoutBalancesListener(
	connStr string,
	balanceRepo core.WalletBalanceRepository,
	walletRepo core.UserWalletRepository,
	publisher core.EventPublisher,
	livePublisher core.LiveEventPublisher,
	log logger.Logger,
	instance string,
) *BlockscoutBalancesListener {
	return &BlockscoutBalancesListener{
		connStr:       connStr,
		balanceRepo:   balanceRepo,
		walletRepo:    walletRepo,
		publisher:     publisher,
		livePublisher: livePublisher,
		log:           log,
		instance:      instance,
	}
}

// Start listens for balance_change notifications until ctx is cancelled.
func (l *BlockscoutBalancesListener) Start(ctx context.Context) {
	l.log.Info("Blockscout balances listener starting")
	backoff := time.Second
	for {
		if err := l.listen(ctx); err != nil {
			if ctx.Err() != nil {
				l.log.Info("Blockscout balances listener stopped")
				return
			}
			l.log.Warn("Blockscout balances listener disconnected, reconnecting",
				"error", err, "backoff", backoff)
			select {
			case <-ctx.Done():
				return
			case <-time.After(backoff):
			}
			if backoff < 30*time.Second {
				backoff *= 2
			}
		} else {
			backoff = time.Second
		}
	}
}

func (l *BlockscoutBalancesListener) listen(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, l.connStr)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer func() { _ = conn.Close(ctx) }()

	if _, err = conn.Exec(ctx, "LISTEN balance_change"); err != nil {
		return fmt.Errorf("LISTEN: %w", err)
	}

	l.log.Info("Blockscout balances listener ready, waiting for balance_change notifications")

	for {
		notification, err := conn.WaitForNotification(ctx)
		if err != nil {
			return err
		}
		l.handle(ctx, notification.Payload)
	}
}

func (l *BlockscoutBalancesListener) handle(ctx context.Context, raw string) {
	var p balanceChangePayload
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		l.log.Error("Failed to parse balance_change payload", "payload", raw, "error", err)
		return
	}

	walletAddr := domain.NormalizeAddress(p.AddressHash)
	tokenAddr := domain.NormalizeAddress(p.TokenContractAddressHash)

	if _, lookupErr := l.walletRepo.FindByRaylsAddress(ctx, walletAddr); lookupErr != nil {
		if errors.Is(lookupErr, core.ErrRecordNotFound) {
			return
		}
		l.log.Error("Failed to look up wallet for balance change", "address", walletAddr, "error", lookupErr)
		return
	}

	value := p.Value
	if value == "" {
		value = "0"
	}

	balance := &domain.WalletBalance{
		WalletAddress: walletAddr,
		TokenAddress:  tokenAddr,
		Balance:       value,
		BlockNumber:   uint64(p.BlockNumber),
	}
	if err := l.balanceRepo.Upsert(ctx, balance); err != nil {
		l.log.Error("Failed to upsert wallet balance",
			"wallet", walletAddr, "token", tokenAddr, "error", err)
		return
	}

	l.log.Info("Wallet balance updated",
		"wallet", walletAddr, "token", tokenAddr,
		"balance", value, "block", p.BlockNumber, "op", p.Op)

	evt := WalletBalanceEvent{
		Type:          "balance_updated",
		WalletAddress: walletAddr,
		TokenAddress:  tokenAddr,
		Balance:       value,
		BlockNumber:   uint64(p.BlockNumber),
	}
	l.publish(ctx, subjectWalletBalanceUpdated, evt)
	l.publishSSE(evt)
}

func (l *BlockscoutBalancesListener) publish(ctx context.Context, subject string, payload any) {
	if l.publisher == nil {
		return
	}
	scoped := InstanceSubject(l.instance, subject)
	if err := l.publisher.Publish(ctx, scoped, payload); err != nil {
		l.log.Warn("Failed to publish wallet balance event", "subject", scoped, "error", err)
	}
}

func (l *BlockscoutBalancesListener) publishSSE(evt WalletBalanceEvent) {
	if l.livePublisher == nil {
		return
	}
	data, err := json.Marshal(evt)
	if err != nil {
		l.log.Warn("Failed to marshal SSE wallet balance event",
			"wallet", evt.WalletAddress, "token", evt.TokenAddress, "error", err)
		return
	}
	subject := InstanceSubject(l.instance, SubjectWalletBalanceSSE)
	if err := l.livePublisher.PublishLive(subject, data); err != nil {
		l.log.Warn("Failed to publish SSE wallet balance event", "subject", subject, "error", err)
	}
}
