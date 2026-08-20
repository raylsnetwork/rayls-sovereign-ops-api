package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/indexer"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/di"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

func RunWorker(configPath string) error {
	container, err := di.New(configPath)
	if err != nil {
		return fmt.Errorf("failed to create DI container: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func() {
		if cleanupErr := container.Close(); cleanupErr != nil {
			logger.Error("Error cleaning up container", "error", cleanupErr)
		}
	}()

	if container.BlockscoutBackfiller != nil {
		if backfillErr := container.BlockscoutBackfiller.Run(ctx); backfillErr != nil {
			logger.Warn("Blockscout backfill failed", "error", backfillErr)
		}
	}

	if container.BlockscoutListener != nil {
		go container.BlockscoutListener.Start(ctx)
	}

	if container.BlockscoutBalancesBackfiller != nil {
		if backfillErr := container.BlockscoutBalancesBackfiller.Run(ctx); backfillErr != nil {
			logger.Warn("Blockscout balances backfill failed", "error", backfillErr)
		}
	}

	if container.BlockscoutBalancesListener != nil {
		go container.BlockscoutBalancesListener.Start(ctx)
	}

	// Start the consumer before the backfiller so events are processed in real-time
	// as the backfill publishes them, instead of piling up until backfill completes.
	if container.AccessManagerEventHandler != nil && container.NATSManager != nil {
		instance := container.Config.InstanceName
		consumer, consumerErr := container.NATSManager.NewConsumer(ctx,
			indexer.InstanceJSName(instance, indexer.StreamAccessManager),
			indexer.InstanceJSName(instance, "am_events_processor"),
			indexer.InstanceSubject(instance, "ops.access_manager.events"))
		if consumerErr != nil {
			logger.Warn("Failed to create access manager event consumer", "error", consumerErr)
		} else {
			go func() {
				for {
					msg, nextErr := consumer.Next(ctx)
					if nextErr != nil {
						if ctx.Err() != nil {
							return
						}
						logger.Error("Access manager consumer error", "error", nextErr)
						continue
					}
					var cl indexer.ContractLog
					if unmarshalErr := json.Unmarshal(msg.Data, &cl); unmarshalErr != nil {
						_ = msg.Nak()
						continue
					}
					if handleErr := container.AccessManagerEventHandler.Handle(ctx, cl); handleErr != nil {
						logger.Error("Access manager event handler failed", "error", handleErr)
					}
					_ = msg.Ack(ctx)
				}
			}()
		}
	}

	backfillDone := make(chan struct{})
	if container.AccessManagerBackfiller != nil {
		go func() {
			defer close(backfillDone)
			if backfillErr := container.AccessManagerBackfiller.Run(ctx); backfillErr != nil {
				logger.Warn("AccessManagerBackfiller failed", "error", backfillErr)
			}
		}()
	} else {
		close(backfillDone)
	}

	if container.AccessManagerListener != nil {
		ticker := time.NewTicker(time.Second)
		go func() {
			defer ticker.Stop()
			select {
			case <-backfillDone:
			case <-ctx.Done():
				return
			}
			if err := container.AccessManagerListener.Run(ctx, ticker.C); err != nil {
				logger.Error("AccessManagerListener stopped", "error", err)
			}
		}()
	}

	logger.Info("Worker started")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-sigCh:
		logger.Info("Worker shutting down", "signal", sig.String())
		cancel()
	case <-ctx.Done():
	}

	return nil
}
