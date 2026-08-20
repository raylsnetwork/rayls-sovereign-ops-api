package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AccessManagerRole struct {
	RoleID         uint64
	Name           string
	Label          string
	AdminRoleID    *uint64
	GuardianRoleID *uint64
	GrantDelaySecs int
	UpdatedAt      time.Time
}

type AccessManagerRoleMember struct {
	RoleID         uint64
	Account        string
	ExecutionDelay int
	ActiveSince    *time.Time
	IsActive       bool
	UpdatedAt      time.Time
}

type AccessManagerManagedContract struct {
	ContractAddress string
	Deployer        string
	IsPaused        bool
	RegisteredAt    time.Time
	UpdatedAt       time.Time
}

type AccessManagerFunctionPermission struct {
	ContractAddress string
	Selector        string
	RoleID          uint64
	IsActive        bool
	UpdatedAt       time.Time
}

type AccessManagerScheduledOperation struct {
	OperationID     string
	Caller          string
	ManagedContract string
	ExecuteAfter    *time.Time
	Status          string
	UpdatedAt       time.Time
}

type AccessManagerContractScopedRoleMember struct {
	RoleID          uint64
	Account         string
	ContractAddress string
	ExecutionDelay  int
	ActiveSince     *time.Time
	IsActive        bool
	UpdatedAt       time.Time
}

type AccessManagerEventLog struct {
	ID              uuid.UUID
	EventName       string
	ContractAddress string
	BlockNumber     uint64
	TxHash          string
	LogIndex        uint64
	BlockTime       time.Time
	Payload         json.RawMessage
	CreatedAt       time.Time
}

const (
	ScheduledOperationStatusScheduled = "scheduled"
	ScheduledOperationStatusExecuted  = "executed"
	ScheduledOperationStatusCanceled  = "canceled"
)
