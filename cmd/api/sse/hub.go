// Package sse provides an in-memory fan-out hub for Server-Sent Events.
// A single producer (the NATS subscriber) calls Broadcast; each connected SSE
// handler registers a channel and streams the messages it receives.
package sse

import "sync"

// clientBuffer is the per-connection channel buffer. A client that falls this far
// behind has messages dropped rather than blocking the broadcaster.
const clientBuffer = 16

// Hub fans out messages to all registered SSE clients. Safe for concurrent use.
type Hub struct {
	mu      sync.RWMutex
	clients map[chan []byte]struct{}
}

// NewHub creates an empty Hub.
func NewHub() *Hub {
	return &Hub{clients: make(map[chan []byte]struct{})}
}

// Register returns a new buffered channel subscribed to broadcasts.
// The caller must Unregister it when the connection ends.
func (h *Hub) Register() chan []byte {
	ch := make(chan []byte, clientBuffer)
	h.mu.Lock()
	h.clients[ch] = struct{}{}
	h.mu.Unlock()
	return ch
}

// Unregister removes ch from the hub and closes it. Safe to call once per channel.
func (h *Hub) Unregister(ch chan []byte) {
	h.mu.Lock()
	if _, ok := h.clients[ch]; ok {
		delete(h.clients, ch)
		close(ch)
	}
	h.mu.Unlock()
}

// Broadcast delivers msg to every registered client without blocking. A client whose
// buffer is full has this message dropped (it will catch up on the next delivery).
func (h *Hub) Broadcast(msg []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for ch := range h.clients {
		select {
		case ch <- msg:
		default:
		}
	}
}

// ClientCount returns the number of connected clients (used in tests and metrics).
func (h *Hub) ClientCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}
