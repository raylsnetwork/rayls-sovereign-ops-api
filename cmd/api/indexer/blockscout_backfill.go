package indexer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

const defaultBackfillBatch = 50

type BlockscoutBackfiller struct {
	connStr   string
	tokenRepo core.TokenRepository
	stateRepo core.IndexerStateRepository
	batch     int
	log       logger.Logger
	chainID   string // chain this instance indexes; stamped as issuerId on discovery
}

func NewBlockscoutBackfiller(
	connStr string,
	tokenRepo core.TokenRepository,
	stateRepo core.IndexerStateRepository,
	batch int,
	log logger.Logger,
	chainID string,
) *BlockscoutBackfiller {
	if batch <= 0 {
		batch = defaultBackfillBatch
	}
	return &BlockscoutBackfiller{
		connStr:   connStr,
		tokenRepo: tokenRepo,
		stateRepo: stateRepo,
		batch:     batch,
		log:       log,
		chainID:   chainID,
	}
}

// Run fetches all tokens from Blockscout inserted after the stored cursor,
// upserts them into the ops-api tokens table, and advances the cursor.
// Runs synchronously — call before starting the listener goroutine.
func (b *BlockscoutBackfiller) Run(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, b.connStr)
	if err != nil {
		return fmt.Errorf("backfill connect: %w", err)
	}
	defer func() { _ = conn.Close(ctx) }()

	cursorTime := time.Time{}
	cursorAddr := ""

	if ts, stateErr := b.stateRepo.Get(ctx, core.CursorBlockscoutInsertedAt); stateErr == nil {
		if t, parseErr := time.Parse(time.RFC3339Nano, ts); parseErr == nil {
			cursorTime = t
		}
	} else if !errors.Is(stateErr, core.ErrRecordNotFound) {
		return fmt.Errorf("read cursor time: %w", stateErr)
	}

	if addr, stateErr := b.stateRepo.Get(ctx, core.CursorBlockscoutAddress); stateErr == nil {
		cursorAddr = addr
	} else if !errors.Is(stateErr, core.ErrRecordNotFound) {
		return fmt.Errorf("read cursor addr: %w", stateErr)
	}

	b.log.Info("Blockscout backfill starting", "cursor_time", cursorTime, "cursor_addr", cursorAddr)

	total := 0
	for {
		rows, queryErr := conn.Query(ctx, `
			SELECT inserted_at,
			       encode(contract_address_hash, 'hex'),
			       COALESCE(name, ''),
			       COALESCE(symbol, ''),
			       COALESCE(type, ''),
			       COALESCE(decimals, 0)::int,
			       COALESCE(total_supply::text, ''),
			       COALESCE(holder_count, 0)
			FROM tokens
			WHERE (inserted_at, contract_address_hash) > ($1, decode($2, 'hex'))
			ORDER BY inserted_at, contract_address_hash
			LIMIT $3`,
			cursorTime, stripHexPrefix(cursorAddr), b.batch,
		)
		if queryErr != nil {
			return fmt.Errorf("backfill query: %w", queryErr)
		}

		type tokenRow struct {
			insertedAt  time.Time
			address     string
			name        string
			symbol      string
			ercType     string
			decimals    int
			totalSupply string
			holderCount int
		}
		var batch []tokenRow
		for rows.Next() {
			var r tokenRow
			if scanErr := rows.Scan(
				&r.insertedAt,
				&r.address,
				&r.name,
				&r.symbol,
				&r.ercType,
				&r.decimals,
				&r.totalSupply,
				&r.holderCount,
			); scanErr != nil {
				rows.Close()
				return fmt.Errorf("backfill scan: %w", scanErr)
			}
			batch = append(batch, r)
		}
		rows.Close()
		if rows.Err() != nil {
			return fmt.Errorf("backfill rows: %w", rows.Err())
		}

		if len(batch) == 0 {
			break
		}

		for _, r := range batch {
			addr := domain.NormalizeAddress(r.address)
			token := &domain.Token{
				Name:            r.name,
				Symbol:          r.symbol,
				ErcStandard:     domain.ParseErcStandard(r.ercType),
				ContractAddress: addr,
				Decimals:        clampDecimals(r.decimals),
				TotalSupply:     r.totalSupply,
				HolderCount:     r.holderCount,
				Status:          domain.TokenStatusActive,
				TokenClass:      "unknown",
				// One chain per instance; Upsert only fills issuer_id when empty (keeps
				// API-deployed values).
				IssuerID: b.chainID,
			}
			if upsertErr := b.tokenRepo.Upsert(ctx, token); upsertErr != nil {
				return fmt.Errorf("backfill upsert %s: %w", addr, upsertErr)
			}
		}

		last := batch[len(batch)-1]
		cursorTime = last.insertedAt
		cursorAddr = domain.NormalizeAddress(last.address)

		if setErr := b.stateRepo.Set(
			ctx,
			core.CursorBlockscoutInsertedAt,
			cursorTime.UTC().Format(time.RFC3339Nano),
		); setErr != nil {
			return fmt.Errorf("save cursor time: %w", setErr)
		}
		if setErr := b.stateRepo.Set(ctx, core.CursorBlockscoutAddress, cursorAddr); setErr != nil {
			return fmt.Errorf("save cursor addr: %w", setErr)
		}

		total += len(batch)
		b.log.Info("Backfill batch processed", "count", len(batch), "total", total, "cursor_time", cursorTime)

		if len(batch) < b.batch {
			break
		}
	}

	b.log.Info("Blockscout backfill complete", "total", total)
	return nil
}

// stripHexPrefix removes 0x/0X prefix for use in SQL decode(addr, 'hex').
func stripHexPrefix(addr string) string {
	if len(addr) >= 2 && (addr[:2] == "0x" || addr[:2] == "0X") {
		return addr[2:]
	}
	return addr
}

// clampDecimals safely narrows the Blockscout decimals column (scanned as int) to
// the uint8 the token model uses. ERC token decimals are always within 0..255;
// anything outside that range is treated as unknown (0), which the Upsert guards.
func clampDecimals(d int) uint8 {
	if d < 0 || d > 255 {
		return 0
	}
	return uint8(d)
}
