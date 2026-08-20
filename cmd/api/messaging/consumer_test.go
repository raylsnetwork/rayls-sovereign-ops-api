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

func TestConsumer_Next_ReceivesPublishedMessage(t *testing.T) {
	// Next blocks until a message is available and returns its payload
	mgr := setupNATS(t)
	pub := mgr.NewPublisher()
	cons, err := mgr.NewConsumer(context.Background(), "test_consumer", "ops.token.status")
	require.NoError(t, err)
	t.Cleanup(cons.Stop)

	type event struct {
		Status string `json:"status"`
	}
	sent := event{Status: "ACTIVE"}
	require.NoError(t, pub.Publish(context.Background(), "ops.token.status", sent))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg, err := cons.Next(ctx)
	require.NoError(t, err)
	require.NoError(t, msg.Ack(ctx))

	var got event
	require.NoError(t, json.Unmarshal(msg.Data, &got))
	assert.Equal(t, sent.Status, got.Status)
}

func TestConsumer_Next_ReturnsErrorOnContextCancellation(t *testing.T) {
	// Next returns a context error when the context is cancelled before a message arrives
	mgr := setupNATS(t)
	cons, err := mgr.NewConsumer(context.Background(), "test_cancel", "ops.token.cancel")
	require.NoError(t, err)
	t.Cleanup(cons.Stop)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = cons.Next(ctx)

	assert.ErrorIs(t, err, context.Canceled)
}
