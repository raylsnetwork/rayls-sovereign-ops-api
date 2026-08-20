package indexer

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/config"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts/RaylsAccessManagerV1"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

const defaultBlockBatchSize = int64(100)

// AccessManagerListener polls the blockchain for RaylsAccessManagerV1 events
// and publishes each decoded log to NATS for downstream processing.
type AccessManagerListener struct {
	client    *ethclient.Client
	parser    *AccessManagerLogParser
	stateRepo core.IndexerStateRepository
	publisher core.EventPublisher
	cfg       *config.Config
	log       logger.Logger
}

func NewAccessManagerListener(
	client *ethclient.Client,
	contract *RaylsAccessManagerV1.RaylsAccessManagerV1,
	addr common.Address,
	stateRepo core.IndexerStateRepository,
	publisher core.EventPublisher,
	cfg *config.Config,
	log logger.Logger,
) *AccessManagerListener {
	return &AccessManagerListener{
		client:    client,
		parser:    NewAccessManagerLogParser(client, contract, addr),
		stateRepo: stateRepo,
		publisher: publisher,
		cfg:       cfg,
		log:       log,
	}
}

// Run processes one tick at a time from tickC until ctx is cancelled.
func (l *AccessManagerListener) Run(ctx context.Context, tickC <-chan time.Time) error {
	l.log.Info("AccessManagerListener started")
	for {
		select {
		case <-ctx.Done():
			l.log.Info("AccessManagerListener stopped")
			return nil
		case <-tickC:
			if err := l.tick(ctx); err != nil {
				l.log.Warn("AccessManagerListener tick error", "error", err)
			}
		}
	}
}

func (l *AccessManagerListener) tick(ctx context.Context) error {
	fromBlock, err := l.fromBlock(ctx)
	if err != nil {
		return fmt.Errorf("read last block: %w", err)
	}

	latest, err := l.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("get latest block: %w", err)
	}

	if fromBlock > latest {
		return nil
	}

	batchSize := l.cfg.Blockchain.BlockBatchSize
	if batchSize <= 0 {
		batchSize = defaultBlockBatchSize
	}

	toBlock := fromBlock + uint64(batchSize) - 1
	if toBlock > latest {
		toBlock = latest
	}

	logs, err := l.parser.ParseRange(ctx, fromBlock, toBlock)
	if err != nil {
		return fmt.Errorf("parse range [%d, %d]: %w", fromBlock, toBlock, err)
	}

	if l.publisher != nil {
		subject := InstanceSubject(l.cfg.InstanceName, subjectAccessManagerEvents)
		for _, cl := range logs {
			if pubErr := l.publisher.Publish(ctx, subject, cl); pubErr != nil {
				return fmt.Errorf("publish event %s tx %s: %w", cl.EventName, cl.TransactionHash, pubErr)
			}
		}
	}

	if len(logs) > 0 {
		l.log.Info("AccessManagerListener batch processed",
			"from", fromBlock, "to", toBlock, "events", len(logs))
	}

	return l.stateRepo.Set(ctx, core.CursorAccessManagerLastBlock, strconv.FormatUint(toBlock+1, 10))
}

func (l *AccessManagerListener) fromBlock(ctx context.Context) (uint64, error) {
	val, err := l.stateRepo.Get(ctx, core.CursorAccessManagerLastBlock)
	if errors.Is(err, core.ErrRecordNotFound) {
		return l.cfg.Blockchain.StartingBlock, nil
	}
	if err != nil {
		return 0, err
	}
	n, parseErr := strconv.ParseUint(val, 10, 64)
	if parseErr != nil {
		return l.cfg.Blockchain.StartingBlock, nil
	}
	return n, nil
}

const (
	StreamAccessManager        = "AM_EVENTS"
	subjectAccessManagerEvents = "ops.access_manager.events"
)
