package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivacyNodeStatus_IsKnown_AcceptsEveryEnumValue(t *testing.T) {
	// Every status the build defines is recognized, so a normal read never warns.
	for _, s := range []PrivacyNodeStatus{
		PrivacyNodeStatusUndefined,
		PrivacyNodeStatusWaitingApproval,
		PrivacyNodeStatusAuthorized,
		PrivacyNodeStatusUnauthorized,
		PrivacyNodeStatusFrozen,
	} {
		assert.True(t, s.IsKnown(), "status %d (%s) must be known", uint8(s), s.Label())
	}
}

func TestPrivacyNodeStatus_IsKnown_RejectsValueAddedOnChain(t *testing.T) {
	// A status added by a contract upgrade this build predates is reported as unknown — the
	// signal the read paths log on.
	assert.False(t, PrivacyNodeStatus(5).IsKnown())
	assert.Equal(t, "unknown", PrivacyNodeStatus(5).Label())
}

func TestParseSettableStatus_AcceptsTheOperatorSettableLabels(t *testing.T) {
	// The two statuses an operator may set map to their enum values.
	authorized, ok := ParseSettableStatus("authorized")
	assert.True(t, ok)
	assert.Equal(t, PrivacyNodeStatusAuthorized, authorized)

	unauthorized, ok := ParseSettableStatus("unauthorized")
	assert.True(t, ok)
	assert.Equal(t, PrivacyNodeStatusUnauthorized, unauthorized)
}

func TestParseSettableStatus_RejectsNonOperatorStatuses(t *testing.T) {
	// waiting_approval/undefined are registration-assigned and frozen has its own endpoints, so
	// none of them may be set through this path.
	for _, label := range []string{"undefined", "waiting_approval", "frozen"} {
		_, ok := ParseSettableStatus(label)
		assert.False(t, ok, "%q must not be settable via updatePrivacyNodeStatus", label)
	}
}

func TestParseSettableStatus_RejectsUnknownAndNumericInput(t *testing.T) {
	// Unrecognized labels are rejected, including the raw enum numbers the endpoint used to take.
	for _, label := range []string{"", "AUTHORIZED", "2", "bogus"} {
		_, ok := ParseSettableStatus(label)
		assert.False(t, ok, "%q must be rejected", label)
	}
}

func TestParseSettableStatus_RoundTripsWithLabel(t *testing.T) {
	// The parser is keyed on Label(), so the two stay in sync if either is edited.
	for _, s := range []PrivacyNodeStatus{PrivacyNodeStatusAuthorized, PrivacyNodeStatusUnauthorized} {
		parsed, ok := ParseSettableStatus(s.Label())
		assert.True(t, ok)
		assert.Equal(t, s, parsed)
	}
}
