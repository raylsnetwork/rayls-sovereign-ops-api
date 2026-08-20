package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/raylsnetwork/rayls-privacy-ops-api/withstack"
)

// TLSConfig bundles the optional mutual-TLS files NewManager uses when
// connecting to NATS. All three paths must be non-empty for TLS to be applied;
// leaving them empty keeps the previous plain-text behaviour.
type TLSConfig struct {
	CAFile   string
	CertFile string
	KeyFile  string
}

// Manager owns the NATS connection and JetStream context.
// Each domain registers its own stream via EnsureStream.
type Manager struct {
	nc *nats.Conn
	js jetstream.JetStream
}

func NewManager(ctx context.Context, natsURL string, tlsCfg TLSConfig) (*Manager, error) {
	opts := []nats.Option{}
	if tlsCfg.CAFile != "" {
		opts = append(opts, nats.RootCAs(tlsCfg.CAFile))
	}
	if tlsCfg.CertFile != "" && tlsCfg.KeyFile != "" {
		opts = append(opts, nats.ClientCert(tlsCfg.CertFile, tlsCfg.KeyFile))
	}
	nc, err := nats.Connect(natsURL, opts...)
	if err != nil {
		return nil, withstack.Wrap(fmt.Errorf("failed to connect to NATS at %s: %w", natsURL, err))
	}

	js, err := jetstream.New(nc)
	if err != nil {
		nc.Close()
		return nil, withstack.Wrap(fmt.Errorf("failed to create JetStream context: %w", err))
	}

	return &Manager{nc: nc, js: js}, nil
}

// EnsureStream creates or updates a WorkQueue JetStream stream for the given subjects.
// Call once per domain at startup before creating publishers or consumers.
func (m *Manager) EnsureStream(ctx context.Context, name string, subjects []string) error {
	cfg := jetstream.StreamConfig{
		Name:       name,
		Retention:  jetstream.WorkQueuePolicy,
		Subjects:   subjects,
		Duplicates: time.Hour,
	}
	if _, err := m.js.CreateOrUpdateStream(ctx, cfg); err != nil {
		return withstack.Wrap(fmt.Errorf("failed to ensure stream %s: %w", name, err))
	}
	return nil
}

// Close drains in-flight messages and closes the underlying NATS connection.
// Must be called on application shutdown.
func (m *Manager) Close() {
	m.nc.Drain()
}

// NewPublisher returns a Publisher that fulfils core.EventPublisher.
func (m *Manager) NewPublisher() *Publisher {
	return newPublisher(m.js)
}

// PublishLive publishes an ephemeral (non-JetStream) message over core NATS, for live
// fan-out such as SSE. The message is NOT persisted in any stream and is delivered
// at-most-once to current subscribers. Fulfils core.LiveEventPublisher.
func (m *Manager) PublishLive(subject string, data []byte) error {
	if err := m.nc.Publish(subject, data); err != nil {
		return withstack.Wrap(fmt.Errorf("failed to publish live message to %s: %w", subject, err))
	}
	return nil
}

// Subscribe registers a core NATS subscriber on subject for live fan-out (e.g. SSE).
// The returned subscription must be Unsubscribe()'d on shutdown.
func (m *Manager) Subscribe(subject string, handler func(data []byte)) (*nats.Subscription, error) {
	sub, err := m.nc.Subscribe(subject, func(msg *nats.Msg) { handler(msg.Data) })
	if err != nil {
		return nil, withstack.Wrap(fmt.Errorf("failed to subscribe to %s: %w", subject, err))
	}
	return sub, nil
}

// NewConsumer creates a durable pull consumer on streamName filtered to filterSubject.
// name must be unique per consumer (e.g. "am_events_processor").
// Messages are redelivered every 2 minutes on failure with no maximum retry limit.
func (m *Manager) NewConsumer(ctx context.Context, streamName, name, filterSubject string) (*Consumer, error) {
	cons, err := m.js.CreateOrUpdateConsumer(ctx, streamName, jetstream.ConsumerConfig{
		Durable:       name,
		MaxDeliver:    -1,
		AckWait:       2 * time.Minute,
		FilterSubject: filterSubject,
	})
	if err != nil {
		return nil, withstack.Wrap(fmt.Errorf("failed to create consumer %s on stream %s: %w", name, streamName, err))
	}
	return newConsumer(cons)
}
