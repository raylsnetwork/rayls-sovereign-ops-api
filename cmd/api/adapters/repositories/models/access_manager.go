package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

type AMRole struct {
	RoleID         uint64 `gorm:"primaryKey;autoIncrement:false"`
	Name           string `gorm:"type:text;not null;default:''"`
	Label          string `gorm:"type:text;not null;default:''"`
	AdminRoleID    *uint64
	GuardianRoleID *uint64
	GrantDelaySecs int       `gorm:"not null;default:0"`
	UpdatedAt      time.Time `gorm:"not null;default:now()"`
}

func (AMRole) TableName() string { return "am_roles" }

func (r *AMRole) ToDomain() *domain.AccessManagerRole {
	return &domain.AccessManagerRole{
		RoleID:         r.RoleID,
		Name:           r.Name,
		Label:          r.Label,
		AdminRoleID:    r.AdminRoleID,
		GuardianRoleID: r.GuardianRoleID,
		GrantDelaySecs: r.GrantDelaySecs,
		UpdatedAt:      r.UpdatedAt,
	}
}

func AMRoleFromDomain(d *domain.AccessManagerRole) *AMRole {
	return &AMRole{
		RoleID:         d.RoleID,
		Name:           d.Name,
		Label:          d.Label,
		AdminRoleID:    d.AdminRoleID,
		GuardianRoleID: d.GuardianRoleID,
		GrantDelaySecs: d.GrantDelaySecs,
		UpdatedAt:      d.UpdatedAt,
	}
}

// ---------------------------------------------------------------------------

type AMRoleMember struct {
	RoleID         uint64 `gorm:"primaryKey;autoIncrement:false"`
	Account        string `gorm:"primaryKey;size:42;not null"`
	ExecutionDelay int    `gorm:"not null;default:0"`
	ActiveSince    *time.Time
	IsActive       bool      `gorm:"not null;default:true"`
	UpdatedAt      time.Time `gorm:"not null;default:now()"`
}

func (AMRoleMember) TableName() string { return "am_role_members" }

func (m *AMRoleMember) ToDomain() *domain.AccessManagerRoleMember {
	return &domain.AccessManagerRoleMember{
		RoleID:         m.RoleID,
		Account:        m.Account,
		ExecutionDelay: m.ExecutionDelay,
		ActiveSince:    m.ActiveSince,
		IsActive:       m.IsActive,
		UpdatedAt:      m.UpdatedAt,
	}
}

func AMRoleMemberFromDomain(d *domain.AccessManagerRoleMember) *AMRoleMember {
	return &AMRoleMember{
		RoleID:         d.RoleID,
		Account:        d.Account,
		ExecutionDelay: d.ExecutionDelay,
		ActiveSince:    d.ActiveSince,
		IsActive:       d.IsActive,
		UpdatedAt:      d.UpdatedAt,
	}
}

// ---------------------------------------------------------------------------

type AMManagedContract struct {
	ContractAddress string    `gorm:"primaryKey;size:42;not null"`
	Deployer        string    `gorm:"size:42;not null"`
	IsPaused        bool      `gorm:"not null;default:false"`
	RegisteredAt    time.Time `gorm:"not null;default:now()"`
	UpdatedAt       time.Time `gorm:"not null;default:now()"`
}

func (AMManagedContract) TableName() string { return "am_managed_contracts" }

func (c *AMManagedContract) ToDomain() *domain.AccessManagerManagedContract {
	return &domain.AccessManagerManagedContract{
		ContractAddress: c.ContractAddress,
		Deployer:        c.Deployer,
		IsPaused:        c.IsPaused,
		RegisteredAt:    c.RegisteredAt,
		UpdatedAt:       c.UpdatedAt,
	}
}

func AMManagedContractFromDomain(d *domain.AccessManagerManagedContract) *AMManagedContract {
	return &AMManagedContract{
		ContractAddress: d.ContractAddress,
		Deployer:        d.Deployer,
		IsPaused:        d.IsPaused,
		RegisteredAt:    d.RegisteredAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

// ---------------------------------------------------------------------------

type AMFunctionPermission struct {
	ContractAddress string    `gorm:"primaryKey;size:42;not null"`
	Selector        string    `gorm:"primaryKey;size:10;not null"`
	RoleID          uint64    `gorm:"primaryKey;autoIncrement:false"`
	IsActive        bool      `gorm:"not null;default:true"`
	UpdatedAt       time.Time `gorm:"not null;default:now()"`
}

func (AMFunctionPermission) TableName() string { return "am_function_permissions" }

func (p *AMFunctionPermission) ToDomain() *domain.AccessManagerFunctionPermission {
	return &domain.AccessManagerFunctionPermission{
		ContractAddress: p.ContractAddress,
		Selector:        p.Selector,
		RoleID:          p.RoleID,
		IsActive:        p.IsActive,
		UpdatedAt:       p.UpdatedAt,
	}
}

func AMFunctionPermissionFromDomain(d *domain.AccessManagerFunctionPermission) *AMFunctionPermission {
	return &AMFunctionPermission{
		ContractAddress: d.ContractAddress,
		Selector:        d.Selector,
		RoleID:          d.RoleID,
		IsActive:        d.IsActive,
		UpdatedAt:       d.UpdatedAt,
	}
}

// ---------------------------------------------------------------------------

type AMScheduledOperation struct {
	OperationID     string `gorm:"primaryKey;size:66;not null"`
	Caller          string `gorm:"size:42;not null"`
	ManagedContract string `gorm:"size:42;not null"`
	ExecuteAfter    *time.Time
	Status          string    `gorm:"size:20;not null;default:scheduled"`
	UpdatedAt       time.Time `gorm:"not null;default:now()"`
}

func (AMScheduledOperation) TableName() string { return "am_scheduled_operations" }

func (o *AMScheduledOperation) ToDomain() *domain.AccessManagerScheduledOperation {
	return &domain.AccessManagerScheduledOperation{
		OperationID:     o.OperationID,
		Caller:          o.Caller,
		ManagedContract: o.ManagedContract,
		ExecuteAfter:    o.ExecuteAfter,
		Status:          o.Status,
		UpdatedAt:       o.UpdatedAt,
	}
}

func AMScheduledOperationFromDomain(d *domain.AccessManagerScheduledOperation) *AMScheduledOperation {
	return &AMScheduledOperation{
		OperationID:     d.OperationID,
		Caller:          d.Caller,
		ManagedContract: d.ManagedContract,
		ExecuteAfter:    d.ExecuteAfter,
		Status:          d.Status,
		UpdatedAt:       d.UpdatedAt,
	}
}

// ---------------------------------------------------------------------------

type AMContractScopedRoleMember struct {
	RoleID          uint64 `gorm:"primaryKey;autoIncrement:false"`
	Account         string `gorm:"primaryKey;size:42;not null"`
	ContractAddress string `gorm:"primaryKey;size:42;not null"`
	ExecutionDelay  int    `gorm:"not null;default:0"`
	ActiveSince     *time.Time
	IsActive        bool      `gorm:"not null;default:true"`
	UpdatedAt       time.Time `gorm:"not null;default:now()"`
}

func (AMContractScopedRoleMember) TableName() string { return "am_contract_scoped_role_members" }

func (m *AMContractScopedRoleMember) ToDomain() *domain.AccessManagerContractScopedRoleMember {
	return &domain.AccessManagerContractScopedRoleMember{
		RoleID:          m.RoleID,
		Account:         m.Account,
		ContractAddress: m.ContractAddress,
		ExecutionDelay:  m.ExecutionDelay,
		ActiveSince:     m.ActiveSince,
		IsActive:        m.IsActive,
		UpdatedAt:       m.UpdatedAt,
	}
}

func AMContractScopedRoleMemberFromDomain(d *domain.AccessManagerContractScopedRoleMember) *AMContractScopedRoleMember {
	return &AMContractScopedRoleMember{
		RoleID:          d.RoleID,
		Account:         d.Account,
		ContractAddress: d.ContractAddress,
		ExecutionDelay:  d.ExecutionDelay,
		ActiveSince:     d.ActiveSince,
		IsActive:        d.IsActive,
		UpdatedAt:       d.UpdatedAt,
	}
}

// ---------------------------------------------------------------------------

type AMEventLog struct {
	ID              uuid.UUID       `gorm:"primaryKey;type:uuid"`
	EventName       string          `gorm:"type:text;not null"`
	ContractAddress string          `gorm:"size:42;not null"`
	BlockNumber     uint64          `gorm:"not null"`
	TxHash          string          `gorm:"size:66;not null;uniqueIndex:idx_am_event_log_tx_log"`
	LogIndex        uint64          `gorm:"not null;uniqueIndex:idx_am_event_log_tx_log"`
	BlockTime       time.Time       `gorm:"not null"`
	Payload         json.RawMessage `gorm:"type:jsonb;not null"`
	CreatedAt       time.Time       `gorm:"not null;default:now()"`
}

func (AMEventLog) TableName() string { return "am_event_log" }

func (e *AMEventLog) ToDomain() *domain.AccessManagerEventLog {
	return &domain.AccessManagerEventLog{
		ID:              e.ID,
		EventName:       e.EventName,
		ContractAddress: e.ContractAddress,
		BlockNumber:     e.BlockNumber,
		TxHash:          e.TxHash,
		LogIndex:        e.LogIndex,
		BlockTime:       e.BlockTime,
		Payload:         e.Payload,
		CreatedAt:       e.CreatedAt,
	}
}

func AMEventLogFromDomain(d *domain.AccessManagerEventLog) *AMEventLog {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return &AMEventLog{
		ID:              d.ID,
		EventName:       d.EventName,
		ContractAddress: d.ContractAddress,
		BlockNumber:     d.BlockNumber,
		TxHash:          d.TxHash,
		LogIndex:        d.LogIndex,
		BlockTime:       d.BlockTime,
		Payload:         d.Payload,
		CreatedAt:       d.CreatedAt,
	}
}
