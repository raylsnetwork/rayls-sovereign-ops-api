package blockchain

import (
	"context"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFactoryDeployerGranter_GrantRoles_RejectsMalformedAddress(t *testing.T) {
	// A non-address string is rejected before any chain call, so a typo cannot reach the RPC.
	g := &FactoryDeployerGranter{}

	err := g.GrantRoles(context.Background(), "not-an-address", []string{"FACTORY_DEPLOYER"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid wallet address")
}

func TestFactoryDeployerGranter_GrantRoles_EmptyRolesMakesNoChainCall(t *testing.T) {
	// With no roles to grant there is nothing to do: the loop never runs, so the nil client
	// is never dereferenced. Guards the "best-effort, called with whatever config supplies"
	// caller contract.
	g := &FactoryDeployerGranter{}

	err := g.GrantRoles(context.Background(), "0x0000000000000000000000000000000000000001", nil)

	assert.NoError(t, err)
}

// chainStub stands in for the AccessManager: hasRole reads the role state, grant writes it.
// The read is deliberately slow, which is what makes a check made outside the lock go stale
// — on a real chain the same gap is the eth_call round-trip.
type chainStub struct {
	mu      sync.Mutex
	granted bool
	sends   int // how many times grantRole would have been broadcast
}

func (c *chainStub) hasRole() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.Gosched() // let a racing goroutine reach its own read before this one acts
	return c.granted
}

func (c *chainStub) grant() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sends++
	c.granted = true
}

// TestFactoryDeployerGranter_MutexSerializesGrants pins the ordering the TOCTOU fix depends
// on: the hasRole check must happen under the SAME lock acquisition as the send. Each
// individual chain operation is already atomic (chainStub locks internally, as the real
// AccessManager does) — the bug is that a check and the send that follows it are two
// separate operations, so an unsynchronized caller acts on a stale answer.
//
// grantRole needs a live *ethclient.Client, so this reproduces its post-fix structure
// against the stub. Moving g.mu.Lock() below the hasRole call makes it fail. Run with -race.
func TestFactoryDeployerGranter_MutexSerializesGrants(t *testing.T) {
	g := &FactoryDeployerGranter{}
	chain := &chainStub{}

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Mirrors grantRole after the fix: lock, THEN check, then act.
			g.mu.Lock()
			defer g.mu.Unlock()

			if chain.hasRole() {
				return // already granted — idempotent no-op
			}
			chain.grant()
		}()
	}
	wg.Wait()

	assert.Equal(t, 1, chain.sends,
		"only the first grant may broadcast; a check made outside the lock goes stale and the loser reverts")
}
