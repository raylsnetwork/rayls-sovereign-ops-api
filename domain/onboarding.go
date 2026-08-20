package domain

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

// ApprovalStatus is the on-chain approval state of a user's address pair in RNUserGovernance.
// The integer values mirror the contract's IUserGovernance.ApprovalStatus enum.
type ApprovalStatus uint8

const (
	ApprovalStatusPending  ApprovalStatus = 0
	ApprovalStatusApproved ApprovalStatus = 1
	ApprovalStatusRejected ApprovalStatus = 2
)

// Label returns a lowercase string label for the status (used in API payloads).
func (s ApprovalStatus) Label() string {
	switch s {
	case ApprovalStatusPending:
		return "pending"
	case ApprovalStatusApproved:
		return "approved"
	case ApprovalStatusRejected:
		return "rejected"
	default:
		return "unknown"
	}
}

// OnChainUserID derives the on-chain user identity from an ops-api user ID as keccak256 of the
// canonical UUID string: keccak256([]byte(userID.String())).
func OnChainUserID(userID uuid.UUID) [32]byte {
	var id [32]byte
	copy(id[:], crypto.Keccak256([]byte(userID.String())))
	return id
}
