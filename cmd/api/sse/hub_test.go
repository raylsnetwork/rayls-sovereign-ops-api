package sse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHub_Broadcast_DeliversToAllRegisteredClients(t *testing.T) {
	// A broadcast reaches every registered client channel.
	h := NewHub()
	a := h.Register()
	b := h.Register()

	h.Broadcast([]byte("hello"))

	assert.Equal(t, []byte("hello"), <-a)
	assert.Equal(t, []byte("hello"), <-b)
}

func TestHub_Unregister_ClosesChannelAndStopsDelivery(t *testing.T) {
	// After Unregister the channel is closed and no longer receives broadcasts.
	h := NewHub()
	ch := h.Register()

	h.Unregister(ch)
	h.Broadcast([]byte("after"))

	_, open := <-ch
	assert.False(t, open)
	assert.Equal(t, 0, h.ClientCount())
}

func TestHub_Broadcast_DoesNotBlockOnFullClient(t *testing.T) {
	// A slow client whose buffer is full must not block the broadcaster; extra messages drop.
	h := NewHub()
	ch := h.Register()

	// Fill the buffer plus extra without ever reading.
	for i := 0; i < clientBuffer+5; i++ {
		h.Broadcast([]byte("x"))
	}

	// The buffer holds exactly clientBuffer messages; the rest were dropped (non-blocking).
	require.Len(t, ch, clientBuffer)
}
