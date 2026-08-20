package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

const (
	subjectTokenDiscovered = "ops.tokens.discovered"
	subjectTokenUpdated    = "ops.tokens.supply_updated"
	StreamTokens           = "TOKEN_EVENTS"

	// SubjectTokenSSE is the core-NATS subject for the curated live token event consumed by
	// the API and fanned out to SSE clients. It sits OUTSIDE the TOKEN_EVENTS stream filter
	// (ops.<inst>.tokens.>) so it is not persisted in the WorkQueue stream.
	SubjectTokenSSE = "ops.sse.tokens"
)

// TokenSSEEvent is the curated payload the worker pushes for live SSE fan-out.
// The worker controls exactly which fields (and which events) reach the frontend.
type TokenSSEEvent struct {
	Type        string `json:"type"` // "discovered" | "supply_updated"
	Address     string `json:"address"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	ErcStandard string `json:"ercStandard"`
	TotalSupply string `json:"totalSupply"`
	HolderCount int    `json:"holderCount"`
	Status      string `json:"status"` // lifecycle state after processing (e.g. "active")
}

type tokenChangePayload struct {
	Op      string `json:"op"`
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Type    string `json:"type"`
	// Decimals is nullable in Blockscout's tokens table (unset until the token is
	// cataloged), so model it as a pointer to distinguish "0 decimals" from "unknown".
	Decimals    *uint8 `json:"decimals"`
	TotalSupply string `json:"total_supply"`
	HolderCount int    `json:"holder_count"`
}

// BlockscoutListener subscribes to pg_notify('token_change') on the Blockscout
// database and upserts token data into the ops-api tokens table.
// It reconnects automatically on connection failure with exponential backoff.
type BlockscoutListener struct {
	connStr       string
	tokenRepo     core.TokenRepository
	publisher     core.EventPublisher     // nil when NATS is not configured
	livePublisher core.LiveEventPublisher // nil when NATS is not configured
	log           logger.Logger
	instance      string // forwarded from config.Config.InstanceName; "" means single-instance
	chainID       string // chain this instance indexes; stamped as issuerId on discovery
}

func NewBlockscoutListener(
	connStr string,
	tokenRepo core.TokenRepository,
	publisher core.EventPublisher,
	livePublisher core.LiveEventPublisher,
	log logger.Logger,
	instance string,
	chainID string,
) *BlockscoutListener {
	return &BlockscoutListener{
		connStr:       connStr,
		tokenRepo:     tokenRepo,
		publisher:     publisher,
		livePublisher: livePublisher,
		log:           log,
		instance:      instance,
		chainID:       chainID,
	}
}

// Start listens for token_change notifications until ctx is cancelled.
func (l *BlockscoutListener) Start(ctx context.Context) {
	l.log.Info("Blockscout listener starting")
	backoff := time.Second
	for {
		if err := l.listen(ctx); err != nil {
			if ctx.Err() != nil {
				l.log.Info("Blockscout listener stopped")
				return
			}
			l.log.Warn("Blockscout listener disconnected, reconnecting",
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

func (l *BlockscoutListener) listen(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, l.connStr)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer func() { _ = conn.Close(ctx) }()

	if _, err = conn.Exec(ctx, "LISTEN token_change"); err != nil {
		return fmt.Errorf("LISTEN: %w", err)
	}

	l.log.Info("Blockscout listener ready, waiting for token_change notifications")

	for {
		notification, err := conn.WaitForNotification(ctx)
		if err != nil {
			return err
		}
		l.handle(ctx, notification.Payload)
	}
}

func (l *BlockscoutListener) handle(ctx context.Context, raw string) {
	var p tokenChangePayload
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		l.log.Error("Failed to parse token_change payload", "payload", raw, "error", err)
		return
	}

	address := domain.NormalizeAddress(p.Address)

	switch p.Op {
	case "INSERT":
		token := &domain.Token{
			Name:            p.Name,
			Symbol:          p.Symbol,
			ErcStandard:     domain.ParseErcStandard(p.Type),
			ContractAddress: address,
			Decimals:        decimalsOrZero(p.Decimals),
			TotalSupply:     p.TotalSupply,
			HolderCount:     p.HolderCount,
			Status:          domain.TokenStatusActive,
			TokenClass:      "unknown",
			// Each instance indexes one chain, so a discovered token lives on it. The Upsert
			// guards issuer_id (only fills when empty), so API-deployed tokens keep their value.
			IssuerID: l.chainID,
		}
		if err := l.tokenRepo.Upsert(ctx, token); err != nil {
			l.log.Error("Failed to upsert token", "address", address, "error", err)
			return
		}
		l.log.Info("Token discovered", "address", address, "name", p.Name, "symbol", p.Symbol)
		l.publish(ctx, subjectTokenDiscovered, p)
		l.publishSSE("discovered", l.currentStatus(ctx, address), p)

	case "UPDATE":
		token := &domain.Token{
			Name:            p.Name,
			Symbol:          p.Symbol,
			ErcStandard:     domain.ParseErcStandard(p.Type),
			ContractAddress: address,
			Decimals:        decimalsOrZero(p.Decimals),
			TotalSupply:     p.TotalSupply,
			HolderCount:     p.HolderCount,
		}
		if err := l.tokenRepo.Upsert(ctx, token); err != nil {
			l.log.Error("Failed to update token", "address", address, "error", err)
			return
		}
		l.log.Info("Token updated", "address", address, "name", p.Name,
			"total_supply", p.TotalSupply, "holder_count", p.HolderCount)
		l.publish(ctx, subjectTokenUpdated, p)
		l.publishSSE("supply_updated", l.currentStatus(ctx, address), p)
	}
}

// currentStatus reads the token's stored status after upsert (Internal for API-deployed tokens,
// Active for organically discovered ones). Falls back to Active if the read fails.
func (l *BlockscoutListener) currentStatus(ctx context.Context, address string) domain.TokenStatus {
	if tok, err := l.tokenRepo.FindByAddress(ctx, address); err == nil && tok != nil {
		return tok.Status
	}
	return domain.TokenStatusActive
}

// publishSSE pushes a curated live event for SSE fan-out. The worker owns this payload —
// it decides which fields and which events reach the frontend. No-op when NATS is off.
func (l *BlockscoutListener) publishSSE(evtType string, status domain.TokenStatus, p tokenChangePayload) {
	if l.livePublisher == nil {
		return
	}
	evt := TokenSSEEvent{
		Type:        evtType,
		Address:     domain.NormalizeAddress(p.Address),
		Name:        p.Name,
		Symbol:      p.Symbol,
		ErcStandard: p.Type,
		TotalSupply: p.TotalSupply,
		HolderCount: p.HolderCount,
		Status:      status.Label(),
	}
	data, err := json.Marshal(evt)
	if err != nil {
		l.log.Warn("Failed to marshal SSE token event", "address", evt.Address, "error", err)
		return
	}
	subject := InstanceSubject(l.instance, SubjectTokenSSE)
	if err := l.livePublisher.PublishLive(subject, data); err != nil {
		l.log.Warn("Failed to publish SSE token event", "subject", subject, "error", err)
	}
}

func (l *BlockscoutListener) publish(ctx context.Context, subject string, payload any) {
	if l.publisher == nil {
		return
	}
	scoped := InstanceSubject(l.instance, subject)
	if err := l.publisher.Publish(ctx, scoped, payload); err != nil {
		l.log.Warn("Failed to publish token event", "subject", scoped, "error", err)
	}
}

// decimalsOrZero unwraps the nullable decimals from a token_change payload. A nil
// value (Blockscout has not cataloged the token's decimals yet) maps to 0; the
// Upsert guards the column so this 0 will not clobber a decimals value the deploy
// path already recorded.
func decimalsOrZero(d *uint8) uint8 {
	if d == nil {
		return 0
	}
	return *d
}
