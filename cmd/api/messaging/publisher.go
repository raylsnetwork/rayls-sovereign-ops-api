package messaging

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/withstack"
)

var _ core.EventPublisher = (*Publisher)(nil)

// Publisher fulfils core.EventPublisher by publishing JSON-encoded messages to JetStream.
// A deterministic Msg-Id (SHA-256 of subject + payload) is set on every publish so that
// the stream's 1-hour deduplication window prevents double-processing when the poller
// re-polls an overlapping block range after a restart.
type Publisher struct {
	js jetstream.JetStream
}

func newPublisher(js jetstream.JetStream) *Publisher {
	return &Publisher{js: js}
}

func (p *Publisher) Publish(ctx context.Context, subject string, msg any) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return withstack.Wrap(fmt.Errorf("failed to marshal event for subject %s: %w", subject, err))
	}

	h := sha256.New()
	h.Write([]byte(subject))
	h.Write([]byte(":"))
	h.Write(data)
	msgID := hex.EncodeToString(h.Sum(nil))

	if _, err := p.js.Publish(ctx, subject, data, jetstream.WithMsgID(msgID)); err != nil {
		return withstack.Wrap(fmt.Errorf("failed to publish to %s: %w", subject, err))
	}

	return nil
}
