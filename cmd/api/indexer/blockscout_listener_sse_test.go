package indexer

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

type fakeLivePublisher struct {
	subject string
	data    []byte
	calls   int
}

func (f *fakeLivePublisher) PublishLive(subject string, data []byte) error {
	f.subject = subject
	f.data = data
	f.calls++
	return nil
}

func TestBlockscoutListener_PublishSSE_EmitsCuratedEventOnInstanceSubject(t *testing.T) {
	// publishSSE marshals a curated TokenSSEEvent and publishes it on the instance-scoped subject.
	lp := &fakeLivePublisher{}
	l := NewBlockscoutListener("", nil, nil, lp, &testutil.StubLogger{}, "inst", "")

	l.publishSSE("discovered", domain.TokenStatusInternal, tokenChangePayload{
		Address: "0xABCDEF", Name: "Token", Symbol: "TKN", Type: "ERC-20", TotalSupply: "100", HolderCount: 3,
	})

	require.Equal(t, 1, lp.calls)
	assert.Equal(t, "ops.inst.sse.tokens", lp.subject)

	var evt TokenSSEEvent
	require.NoError(t, json.Unmarshal(lp.data, &evt))
	assert.Equal(t, "discovered", evt.Type)
	assert.Equal(t, "internal", evt.Status)
	assert.Equal(t, domain.NormalizeAddress("0xABCDEF"), evt.Address)
	assert.Equal(t, "TKN", evt.Symbol)
	assert.Equal(t, "ERC-20", evt.ErcStandard)
	assert.Equal(t, 3, evt.HolderCount)
}

func TestBlockscoutListener_PublishSSE_NilPublisherIsNoop(t *testing.T) {
	// With no live publisher (NATS off) publishSSE must not publish or panic.
	l := NewBlockscoutListener("", nil, nil, nil, &testutil.StubLogger{}, "", "")

	l.publishSSE("supply_updated", domain.TokenStatusActive, tokenChangePayload{Address: "0x1"})
}
