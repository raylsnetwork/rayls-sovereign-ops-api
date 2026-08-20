package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// AccessManagerEventHandler routes decoded ContractLogs to the appropriate
// repository upserts, maintaining both materialized state and the audit trail.
type AccessManagerEventHandler struct {
	roleRepo         core.AccessManagerRoleRepository
	memberRepo       core.AccessManagerRoleMemberRepository
	contractRepo     core.AccessManagerManagedContractRepository
	fnPermRepo       core.AccessManagerFunctionPermissionRepository
	scheduledOpRepo  core.AccessManagerScheduledOperationRepository
	scopedMemberRepo core.AccessManagerContractScopedRoleMemberRepository
	eventLogRepo     core.AccessManagerEventLogRepository
	log              logger.Logger
}

func NewAccessManagerEventHandler(
	roleRepo core.AccessManagerRoleRepository,
	memberRepo core.AccessManagerRoleMemberRepository,
	contractRepo core.AccessManagerManagedContractRepository,
	fnPermRepo core.AccessManagerFunctionPermissionRepository,
	scheduledOpRepo core.AccessManagerScheduledOperationRepository,
	scopedMemberRepo core.AccessManagerContractScopedRoleMemberRepository,
	eventLogRepo core.AccessManagerEventLogRepository,
	log logger.Logger,
) *AccessManagerEventHandler {
	return &AccessManagerEventHandler{
		roleRepo:         roleRepo,
		memberRepo:       memberRepo,
		contractRepo:     contractRepo,
		fnPermRepo:       fnPermRepo,
		scheduledOpRepo:  scheduledOpRepo,
		scopedMemberRepo: scopedMemberRepo,
		eventLogRepo:     eventLogRepo,
		log:              log,
	}
}

// Handle processes a single ContractLog: always appends to the audit trail,
// then updates the materialized state table for the specific event type.
func (h *AccessManagerEventHandler) Handle(ctx context.Context, cl ContractLog) error {
	if err := h.eventLogRepo.Insert(ctx, &domain.AccessManagerEventLog{
		EventName:       cl.EventName,
		ContractAddress: cl.ContractAddress,
		BlockNumber:     cl.BlockNumber,
		TxHash:          cl.TransactionHash,
		LogIndex:        cl.LogIndex,
		BlockTime:       cl.BlockTimestamp,
		Payload:         cl.RawEventData,
	}); err != nil {
		h.log.Warn("Failed to insert event log", "event", cl.EventName, "tx", cl.TransactionHash, "error", err)
	}

	switch cl.EventName {
	case "RoleRegistered":
		return h.handleRoleRegistered(ctx, cl)
	case "RoleLabelSet":
		return h.handleRoleLabelSet(ctx, cl)
	case "RoleAdminChanged":
		return h.handleRoleAdminChanged(ctx, cl)
	case "RoleGuardianChanged":
		return h.handleRoleGuardianChanged(ctx, cl)
	case "RoleGrantDelayChanged":
		return h.handleRoleGrantDelayChanged(ctx, cl)
	case "RoleGranted":
		return h.handleRoleGranted(ctx, cl)
	case "RoleRevoked":
		return h.handleRoleRevoked(ctx, cl)
	case "ManagedContractRegistered":
		return h.handleManagedContractRegistered(ctx, cl)
	case "ContractPauseUpdated":
		return h.handleContractPauseUpdated(ctx, cl)
	case "FunctionAllowedRoleAdded":
		return h.handleFunctionAllowedRoleAdded(ctx, cl)
	case "FunctionAllowedRoleRemoved":
		return h.handleFunctionAllowedRoleRemoved(ctx, cl)
	case "OperationScheduled":
		return h.handleOperationScheduled(ctx, cl)
	case "OperationExecuted":
		return h.handleOperationExecuted(ctx, cl)
	case "OperationCanceled":
		return h.handleOperationCanceled(ctx, cl)
	case "ContractScopedRoleGranted":
		return h.handleContractScopedRoleGranted(ctx, cl)
	case "ContractScopedRoleRevoked":
		return h.handleContractScopedRoleRevoked(ctx, cl)
	}
	return nil
}

// ---------------------------------------------------------------------------
// Role events
// ---------------------------------------------------------------------------

type jsonRoleRegistered struct {
	RoleId uint64 `json:"RoleId"`
	Name   string `json:"Name"`
}

func (h *AccessManagerEventHandler) handleRoleRegistered(ctx context.Context, cl ContractLog) error {
	var e jsonRoleRegistered
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal RoleRegistered: %w", err)
	}
	return h.roleRepo.Upsert(ctx, &domain.AccessManagerRole{
		RoleID: e.RoleId,
		Name:   e.Name,
	})
}

type jsonRoleLabelSet struct {
	RoleId uint64 `json:"RoleId"`
	Label  string `json:"Label"`
}

func (h *AccessManagerEventHandler) handleRoleLabelSet(ctx context.Context, cl ContractLog) error {
	var e jsonRoleLabelSet
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal RoleLabelSet: %w", err)
	}
	existing, err := h.roleRepo.FindByID(ctx, e.RoleId)
	if err != nil {
		existing = &domain.AccessManagerRole{RoleID: e.RoleId}
	}
	existing.Label = e.Label
	return h.roleRepo.Upsert(ctx, existing)
}

type jsonRoleAdminChanged struct {
	RoleId   uint64 `json:"RoleId"`
	NewAdmin uint64 `json:"NewAdmin"`
}

func (h *AccessManagerEventHandler) handleRoleAdminChanged(ctx context.Context, cl ContractLog) error {
	var e jsonRoleAdminChanged
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal RoleAdminChanged: %w", err)
	}
	existing, err := h.roleRepo.FindByID(ctx, e.RoleId)
	if err != nil {
		existing = &domain.AccessManagerRole{RoleID: e.RoleId}
	}
	existing.AdminRoleID = &e.NewAdmin
	return h.roleRepo.Upsert(ctx, existing)
}

type jsonRoleGuardianChanged struct {
	RoleId      uint64 `json:"RoleId"`
	NewGuardian uint64 `json:"NewGuardian"`
}

func (h *AccessManagerEventHandler) handleRoleGuardianChanged(ctx context.Context, cl ContractLog) error {
	var e jsonRoleGuardianChanged
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal RoleGuardianChanged: %w", err)
	}
	existing, err := h.roleRepo.FindByID(ctx, e.RoleId)
	if err != nil {
		existing = &domain.AccessManagerRole{RoleID: e.RoleId}
	}
	existing.GuardianRoleID = &e.NewGuardian
	return h.roleRepo.Upsert(ctx, existing)
}

type jsonRoleGrantDelayChanged struct {
	RoleId   uint64 `json:"RoleId"`
	NewDelay uint32 `json:"NewDelay"`
}

func (h *AccessManagerEventHandler) handleRoleGrantDelayChanged(ctx context.Context, cl ContractLog) error {
	var e jsonRoleGrantDelayChanged
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal RoleGrantDelayChanged: %w", err)
	}
	existing, err := h.roleRepo.FindByID(ctx, e.RoleId)
	if err != nil {
		existing = &domain.AccessManagerRole{RoleID: e.RoleId}
	}
	existing.GrantDelaySecs = int(e.NewDelay)
	return h.roleRepo.Upsert(ctx, existing)
}

type jsonRoleGranted struct {
	RoleId         uint64      `json:"RoleId"`
	Account        string      `json:"Account"`
	ExecutionDelay uint32      `json:"ExecutionDelay"`
	ActiveSince    json.Number `json:"ActiveSince"` // big.Int marshals as number in Go 1.22+
}

func (h *AccessManagerEventHandler) handleRoleGranted(ctx context.Context, cl ContractLog) error {
	var e jsonRoleGranted
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal RoleGranted: %w", err)
	}
	h.ensureRoleExists(ctx, e.RoleId)
	activeSince := parseUnixTimestamp(e.ActiveSince.String())
	return h.memberRepo.Upsert(ctx, &domain.AccessManagerRoleMember{
		RoleID:         e.RoleId,
		Account:        e.Account,
		ExecutionDelay: int(e.ExecutionDelay),
		ActiveSince:    activeSince,
		IsActive:       true,
	})
}

type jsonRoleRevoked struct {
	RoleId  uint64 `json:"RoleId"`
	Account string `json:"Account"`
}

func (h *AccessManagerEventHandler) handleRoleRevoked(ctx context.Context, cl ContractLog) error {
	var e jsonRoleRevoked
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal RoleRevoked: %w", err)
	}
	return h.memberRepo.Revoke(ctx, e.RoleId, e.Account)
}

// ---------------------------------------------------------------------------
// Managed contract events
// ---------------------------------------------------------------------------

type jsonManagedContractRegistered struct {
	ManagedContract   string `json:"ManagedContract"`
	ContractAuthority string `json:"ContractAuthority"`
}

func (h *AccessManagerEventHandler) handleManagedContractRegistered(ctx context.Context, cl ContractLog) error {
	var e jsonManagedContractRegistered
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal ManagedContractRegistered: %w", err)
	}
	return h.contractRepo.Upsert(ctx, &domain.AccessManagerManagedContract{
		ContractAddress: e.ManagedContract,
		Deployer:        e.ContractAuthority,
		RegisteredAt:    cl.BlockTimestamp,
	})
}

type jsonContractPauseUpdated struct {
	ManagedContract string `json:"ManagedContract"`
	Paused          bool   `json:"Paused"`
}

func (h *AccessManagerEventHandler) handleContractPauseUpdated(ctx context.Context, cl ContractLog) error {
	var e jsonContractPauseUpdated
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal ContractPauseUpdated: %w", err)
	}
	return h.contractRepo.SetPaused(ctx, e.ManagedContract, e.Paused)
}

// ---------------------------------------------------------------------------
// Function permission events
// ---------------------------------------------------------------------------

// Selector is serialised as a JSON array of 4 numbers (Go [4]byte → []uint8).
type jsonFunctionAllowedRole struct {
	ManagedContract string   `json:"ManagedContract"`
	Selector        [4]uint8 `json:"Selector"`
	RoleId          uint64   `json:"RoleId"`
}

func (h *AccessManagerEventHandler) handleFunctionAllowedRoleAdded(ctx context.Context, cl ContractLog) error {
	var e jsonFunctionAllowedRole
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal FunctionAllowedRoleAdded: %w", err)
	}
	return h.fnPermRepo.Upsert(ctx, &domain.AccessManagerFunctionPermission{
		ContractAddress: e.ManagedContract,
		Selector:        selectorHex(e.Selector),
		RoleID:          e.RoleId,
		IsActive:        true,
	})
}

func (h *AccessManagerEventHandler) handleFunctionAllowedRoleRemoved(ctx context.Context, cl ContractLog) error {
	var e jsonFunctionAllowedRole
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal FunctionAllowedRoleRemoved: %w", err)
	}
	return h.fnPermRepo.Remove(ctx, e.ManagedContract, selectorHex(e.Selector), e.RoleId)
}

// ---------------------------------------------------------------------------
// Scheduled operation events
// ---------------------------------------------------------------------------

type jsonOperationScheduled struct {
	OperationId     [32]uint8   `json:"OperationId"`
	Caller          string      `json:"Caller"`
	ManagedContract string      `json:"ManagedContract"`
	ExecuteAfter    json.Number `json:"ExecuteAfter"` // big.Int marshals as number in Go 1.22+
}

func (h *AccessManagerEventHandler) handleOperationScheduled(ctx context.Context, cl ContractLog) error {
	var e jsonOperationScheduled
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal OperationScheduled: %w", err)
	}
	executeAfter := parseUnixTimestamp(e.ExecuteAfter.String())
	return h.scheduledOpRepo.Upsert(ctx, &domain.AccessManagerScheduledOperation{
		OperationID:     bytes32Hex(e.OperationId),
		Caller:          e.Caller,
		ManagedContract: e.ManagedContract,
		ExecuteAfter:    executeAfter,
		Status:          domain.ScheduledOperationStatusScheduled,
	})
}

type jsonOperationIDOnly struct {
	OperationId [32]uint8 `json:"OperationId"`
}

func (h *AccessManagerEventHandler) handleOperationExecuted(ctx context.Context, cl ContractLog) error {
	var e jsonOperationIDOnly
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal OperationExecuted: %w", err)
	}
	return h.scheduledOpRepo.UpdateStatus(ctx, bytes32Hex(e.OperationId), domain.ScheduledOperationStatusExecuted)
}

func (h *AccessManagerEventHandler) handleOperationCanceled(ctx context.Context, cl ContractLog) error {
	var e jsonOperationIDOnly
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal OperationCanceled: %w", err)
	}
	return h.scheduledOpRepo.UpdateStatus(ctx, bytes32Hex(e.OperationId), domain.ScheduledOperationStatusCanceled)
}

// ---------------------------------------------------------------------------
// Contract-scoped role events
// ---------------------------------------------------------------------------

type jsonContractScopedRoleGranted struct {
	RoleId          uint64      `json:"RoleId"`
	Account         string      `json:"Account"`
	ManagedContract string      `json:"ManagedContract"`
	ExecutionDelay  uint32      `json:"ExecutionDelay"`
	ActiveSince     json.Number `json:"ActiveSince"` // big.Int marshals as number in Go 1.22+
}

func (h *AccessManagerEventHandler) handleContractScopedRoleGranted(ctx context.Context, cl ContractLog) error {
	var e jsonContractScopedRoleGranted
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal ContractScopedRoleGranted: %w", err)
	}
	h.ensureRoleExists(ctx, e.RoleId)
	activeSince := parseUnixTimestamp(e.ActiveSince.String())
	return h.scopedMemberRepo.Upsert(ctx, &domain.AccessManagerContractScopedRoleMember{
		RoleID:          e.RoleId,
		Account:         e.Account,
		ContractAddress: e.ManagedContract,
		ExecutionDelay:  int(e.ExecutionDelay),
		ActiveSince:     activeSince,
		IsActive:        true,
	})
}

type jsonContractScopedRoleRevoked struct {
	RoleId          uint64 `json:"RoleId"`
	Account         string `json:"Account"`
	ManagedContract string `json:"ManagedContract"`
}

func (h *AccessManagerEventHandler) handleContractScopedRoleRevoked(ctx context.Context, cl ContractLog) error {
	var e jsonContractScopedRoleRevoked
	if err := json.Unmarshal(cl.RawEventData, &e); err != nil {
		return fmt.Errorf("unmarshal ContractScopedRoleRevoked: %w", err)
	}
	return h.scopedMemberRepo.Revoke(ctx, e.RoleId, e.Account, e.ManagedContract)
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func selectorHex(b [4]uint8) string {
	return fmt.Sprintf("0x%02x%02x%02x%02x", b[0], b[1], b[2], b[3])
}

func bytes32Hex(b [32]uint8) string {
	return fmt.Sprintf(
		"0x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		b[0],
		b[1],
		b[2],
		b[3],
		b[4],
		b[5],
		b[6],
		b[7],
		b[8],
		b[9],
		b[10],
		b[11],
		b[12],
		b[13],
		b[14],
		b[15],
		b[16],
		b[17],
		b[18],
		b[19],
		b[20],
		b[21],
		b[22],
		b[23],
		b[24],
		b[25],
		b[26],
		b[27],
		b[28],
		b[29],
		b[30],
		b[31],
	)
}

// ensureRoleExists creates a stub role row if the role is not yet known.
// This handles roles registered before BLOCKCHAIN_STARTING_BLOCK whose RoleRegistered
// events were never scanned, but whose RoleGranted events are in range.
func (h *AccessManagerEventHandler) ensureRoleExists(ctx context.Context, roleID uint64) {
	if _, err := h.roleRepo.FindByID(ctx, roleID); err != nil {
		_ = h.roleRepo.Upsert(ctx, &domain.AccessManagerRole{RoleID: roleID})
	}
}

// parseUnixTimestamp parses a uint48 unix timestamp from its decimal string representation.
func parseUnixTimestamp(s string) *time.Time {
	if s == "" || s == "0" {
		return nil
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}
	t := time.Unix(n, 0).UTC()
	return &t
}
