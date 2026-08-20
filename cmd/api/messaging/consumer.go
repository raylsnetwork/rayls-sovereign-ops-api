package messaging

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/withstack"
)

// Message wraps a raw JetStream message with typed Ack/Nak helpers.
type Message struct {
	Data []byte
	Ack  func(context.Context) error
	Nak  func() error
}

// Consumer pulls messages one at a time from a durable JetStream consumer.
// Call Next in a loop inside a goroutine. The loop exits when ctx is cancelled.
type Consumer struct {
	iter jetstream.MessagesContext
}

func newConsumer(jsCons jetstream.Consumer) (*Consumer, error) {
	iter, err := jsCons.Messages()
	if err != nil {
		return nil, withstack.Wrap(fmt.Errorf("failed to create JetStream messages iterator: %w", err))
	}
	return &Consumer{iter: iter}, nil
}

// Stop releases the server-side subscription and internal goroutine held by the
// messages iterator. Must be called when the consumer is no longer needed.
func (c *Consumer) Stop() {
	c.iter.Stop()
}

// Next blocks until the next message arrives or ctx is done.
// It transparently retries on internal MaxWait timeouts so callers only see
// real messages or a context-cancellation error.
func (c *Consumer) Next(ctx context.Context) (*Message, error) {
	for {
		jsMsg, err := c.iter.Next(jetstream.NextContext(ctx))
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				continue
			}
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			return nil, withstack.Wrap(fmt.Errorf("failed to get next JetStream message: %w", err))
		}

		return &Message{
			Data: jsMsg.Data(),
			Ack:  jsMsg.DoubleAck,
			Nak:  jsMsg.Nak,
		}, nil
	}
}
