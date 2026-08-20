//go:build ignore

package messaging

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
)

// setupNATS starts a NATS container with JetStream enabled and returns a connected Manager.
// The container and connection are torn down via t.Cleanup.
func setupNATS(t *testing.T) *Manager {
	t.Helper()

	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "failed to connect to Docker")

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "nats",
		Tag:        "2",
		Cmd:        []string{"-js"},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	require.NoError(t, err, "failed to start NATS container")

	t.Cleanup(func() {
		_ = pool.Purge(resource)
	})

	natsURL := fmt.Sprintf("nats://localhost:%s", resource.GetPort("4222/tcp"))

	var mgr *Manager
	err = pool.Retry(func() error {
		var connErr error
		mgr, connErr = NewManager(context.Background(), natsURL)
		return connErr
	})
	require.NoError(t, err, "NATS did not become ready in time")

	t.Cleanup(func() {
		mgr.Close()
	})

	return mgr
}
