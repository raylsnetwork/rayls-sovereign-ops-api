package domain

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOnChainUserID_HashesUUIDString(t *testing.T) {
	// OnChainUserID is keccak256 of the canonical UUID string, matching the legacy backend.
	userID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	got := OnChainUserID(userID)

	var want [32]byte
	copy(want[:], crypto.Keccak256([]byte(userID.String())))
	assert.Equal(t, want, got)
}

func TestApprovalStatus_Label(t *testing.T) {
	// Labels map the on-chain enum values to lowercase strings.
	assert.Equal(t, "pending", ApprovalStatusPending.Label())
	assert.Equal(t, "approved", ApprovalStatusApproved.Label())
	assert.Equal(t, "rejected", ApprovalStatusRejected.Label())
	assert.Equal(t, "unknown", ApprovalStatus(99).Label())
}
