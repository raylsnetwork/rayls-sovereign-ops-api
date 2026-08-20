package indexer

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/config"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/RaylsAccessManagerV1"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// AccessManagerBackfiller processes all historical blocks from the last saved
// cursor up to the current chain tip. Runs synchronously at startup — call
// before starting the AccessManagerListener goroutine.
type AccessManagerBackfiller struct {
	client    *ethclient.Client
	parser    *AccessManagerLogParser
	stateRepo core.IndexerStateRepository
	publisher core.EventPublisher
	cfg       *config.Config
	log       logger.Logger
}

func NewAccessManagerBackfiller(
	client *ethclient.Client,
	contract *RaylsAccessManagerV1.RaylsAccessManagerV1,
	addr common.Address,
	stateRepo core.IndexerStateRepository,
	publisher core.EventPublisher,
	cfg *config.Config,
	log logger.Logger,
) *AccessManagerBackfiller {
	return &AccessManagerBackfiller{
		client:    client,
		parser:    NewAccessManagerLogParser(client, contract, addr),
		stateRepo: stateRepo,
		publisher: publisher,
		cfg:       cfg,
		log:       log,
	}
}

// Run processes all unindexed blocks as fast as possible, advancing the cursor
// after each batch. Stops when the cursor reaches the chain tip.
func (b *AccessManagerBackfiller) Run(ctx context.Context) error {
	fromBlock, err := b.cursorBlock(ctx)
	if err != nil {
		return fmt.Errorf("read cursor: %w", err)
	}

	latest, err := b.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("get latest block: %w", err)
	}

	if fromBlock > latest {
		b.log.Info("AccessManagerBackfiller: already up to date", "block", fromBlock)
		return nil
	}

	b.log.Info("AccessManagerBackfiller started", "from", fromBlock, "to", latest)

	batchSize := uint64(b.cfg.Blockchain.BlockBatchSize)
	if batchSize == 0 {
		batchSize = uint64(defaultBlockBatchSize)
	}

	total := 0
	for fromBlock <= latest {
		toBlock := fromBlock + batchSize - 1
		if toBlock > latest {
			toBlock = latest
		}

		logs, parseErr := b.parser.ParseRange(ctx, fromBlock, toBlock)
		if parseErr != nil {
			return fmt.Errorf("parse range [%d, %d]: %w", fromBlock, toBlock, parseErr)
		}

		if b.publisher != nil {
			subject := InstanceSubject(b.cfg.InstanceName, subjectAccessManagerEvents)
			for _, cl := range logs {
				if pubErr := b.publisher.Publish(ctx, subject, cl); pubErr != nil {
					return fmt.Errorf("publish event %s tx %s: %w", cl.EventName, cl.TransactionHash, pubErr)
				}
			}
		}

		total += len(logs)
		fromBlock = toBlock + 1

		if saveErr := b.stateRepo.Set(
			ctx,
			core.CursorAccessManagerLastBlock,
			strconv.FormatUint(fromBlock, 10),
		); saveErr != nil {
			return fmt.Errorf("save cursor: %w", saveErr)
		}

		if len(logs) > 0 {
			b.log.Info(
				"AccessManagerBackfiller batch",
				"from",
				fromBlock-batchSize,
				"to",
				toBlock,
				"events",
				len(logs),
				"total",
				total,
			)
		}
	}

	b.log.Info("AccessManagerBackfiller complete", "total_events", total, "next_block", fromBlock)
	return nil
}

func (b *AccessManagerBackfiller) cursorBlock(ctx context.Context) (uint64, error) {
	val, err := b.stateRepo.Get(ctx, core.CursorAccessManagerLastBlock)
	if errors.Is(err, core.ErrRecordNotFound) {
		return b.cfg.Blockchain.StartingBlock, nil
	}
	if err != nil {
		return 0, err
	}
	n, parseErr := strconv.ParseUint(val, 10, 64)
	if parseErr != nil {
		return b.cfg.Blockchain.StartingBlock, nil
	}
	return n, nil
}
