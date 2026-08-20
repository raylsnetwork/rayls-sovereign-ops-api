//go:build ignore

package messaging

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPublisher_Publish_DeliversMessageToStream(t *testing.T) {
	// Publish serialises the payload as JSON and the consumer can read it back
	mgr := setupNATS(t)
	pub := mgr.NewPublisher()

	type payload struct {
		Address string `json:"address"`
	}
	msg := payload{Address: "0xdeadbeef"}

	err := pub.Publish(context.Background(), "ops.token.registered", msg)

	require.NoError(t, err)

	// verify the message arrived by consuming it
	cons, err := mgr.NewConsumer(context.Background(), "test_deliver", "ops.token.registered")
	require.NoError(t, err)
	t.Cleanup(cons.Stop)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	received, err := cons.Next(ctx)
	require.NoError(t, err)
	require.NoError(t, received.Ack(ctx))

	var got payload
	require.NoError(t, json.Unmarshal(received.Data, &got))
	assert.Equal(t, msg.Address, got.Address)
}

func TestPublisher_Publish_ReturnsErrorAfterConnectionClosed(t *testing.T) {
	// Publish returns an error when the underlying NATS connection has been closed
	mgr := setupNATS(t)
	pub := mgr.NewPublisher()

	mgr.Close()

	err := pub.Publish(context.Background(), "ops.token.registered", map[string]string{"k": "v"})

	assert.Error(t, err)
}
