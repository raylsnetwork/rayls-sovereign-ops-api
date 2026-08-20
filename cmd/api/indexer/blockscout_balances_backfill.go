package indexer

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// BlockscoutBalancesBackfiller walks Blockscout's address_current_token_balances
// table from a stored cursor onwards, upserting balances only for addresses that
// belong to a known user_wallets row. The cursor is the bigserial `id` column
// (Blockscout schema): a single monotonic value, persisted in indexer_state under
// CursorBlockscoutBalancesID.
type BlockscoutBalancesBackfiller struct {
	connStr     string
	balanceRepo core.WalletBalanceRepository
	walletRepo  core.UserWalletRepository
	stateRepo   core.IndexerStateRepository
	batch       int
	log         logger.Logger
}

func NewBlockscoutBalancesBackfiller(
	connStr string,
	balanceRepo core.WalletBalanceRepository,
	walletRepo core.UserWalletRepository,
	stateRepo core.IndexerStateRepository,
	batch int,
	log logger.Logger,
) *BlockscoutBalancesBackfiller {
	if batch <= 0 {
		batch = defaultBackfillBatch
	}
	return &BlockscoutBalancesBackfiller{
		connStr:     connStr,
		balanceRepo: balanceRepo,
		walletRepo:  walletRepo,
		stateRepo:   stateRepo,
		batch:       batch,
		log:         log,
	}
}

// Run fetches balance rows from Blockscout in cursor order, upserts those whose
// address_hash matches a known user_wallets.rayls_address, and advances the
// cursor past every row inspected (matched or not). Synchronous — call before
// starting the listener goroutine.
func (b *BlockscoutBalancesBackfiller) Run(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, b.connStr)
	if err != nil {
		return fmt.Errorf("balances backfill connect: %w", err)
	}
	defer func() { _ = conn.Close(ctx) }()

	cursorID := int64(0)
	if v, stateErr := b.stateRepo.Get(ctx, core.CursorBlockscoutBalancesID); stateErr == nil {
		if parsed, parseErr := strconv.ParseInt(v, 10, 64); parseErr == nil {
			cursorID = parsed
		}
	} else if !errors.Is(stateErr, core.ErrRecordNotFound) {
		return fmt.Errorf("read balances cursor: %w", stateErr)
	}

	b.log.Info("Blockscout balances backfill starting", "cursor_id", cursorID)

	total := 0
	for {
		rows, queryErr := conn.Query(ctx, `
			SELECT id,
			       encode(address_hash, 'hex'),
			       encode(token_contract_address_hash, 'hex'),
			       COALESCE(value::text, '0'),
			       COALESCE(block_number, 0)
			FROM address_current_token_balances
			WHERE id > $1
			ORDER BY id
			LIMIT $2`,
			cursorID, b.batch,
		)
		if queryErr != nil {
			return fmt.Errorf("balances backfill query: %w", queryErr)
		}

		type balanceRow struct {
			id           int64
			addressHash  string
			tokenAddress string
			value        string
			blockNumber  int64
		}
		var batch []balanceRow
		for rows.Next() {
			var r balanceRow
			if scanErr := rows.Scan(&r.id, &r.addressHash, &r.tokenAddress, &r.value, &r.blockNumber); scanErr != nil {
				rows.Close()
				return fmt.Errorf("balances backfill scan: %w", scanErr)
			}
			batch = append(batch, r)
		}
		rows.Close()
		if rows.Err() != nil {
			return fmt.Errorf("balances backfill rows: %w", rows.Err())
		}

		if len(batch) == 0 {
			break
		}

		matched := 0
		for _, r := range batch {
			walletAddr := domain.NormalizeAddress(r.addressHash)
			if _, lookupErr := b.walletRepo.FindByRaylsAddress(ctx, walletAddr); lookupErr != nil {
				if errors.Is(lookupErr, core.ErrRecordNotFound) {
					continue
				}
				return fmt.Errorf("balances backfill wallet lookup %s: %w", walletAddr, lookupErr)
			}

			balance := &domain.WalletBalance{
				WalletAddress: walletAddr,
				TokenAddress:  domain.NormalizeAddress(r.tokenAddress),
				Balance:       r.value,
				BlockNumber:   uint64(r.blockNumber),
			}
			if upsertErr := b.balanceRepo.Upsert(ctx, balance); upsertErr != nil {
				return fmt.Errorf("balances backfill upsert %s/%s: %w", walletAddr, balance.TokenAddress, upsertErr)
			}
			matched++
		}

		cursorID = batch[len(batch)-1].id
		if setErr := b.stateRepo.Set(
			ctx,
			core.CursorBlockscoutBalancesID,
			strconv.FormatInt(cursorID, 10),
		); setErr != nil {
			return fmt.Errorf("save balances cursor: %w", setErr)
		}

		total += matched
		b.log.Info("Balances backfill batch processed",
			"scanned", len(batch), "matched", matched, "total_matched", total, "cursor_id", cursorID)

		if len(batch) < b.batch {
			break
		}
	}

	b.log.Info("Blockscout balances backfill complete", "total_matched", total)
	return nil
}
