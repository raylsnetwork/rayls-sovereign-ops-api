package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// AccessManagerHandler serves read-only endpoints backed by the indexed AccessManager state.
type AccessManagerHandler struct {
	roleRepo        core.AccessManagerRoleRepository
	memberRepo      core.AccessManagerRoleMemberRepository
	contractRepo    core.AccessManagerManagedContractRepository
	fnPermRepo      core.AccessManagerFunctionPermissionRepository
	scheduledOpRepo core.AccessManagerScheduledOperationRepository
	eventLogRepo    core.AccessManagerEventLogRepository
	log             logger.Logger
}

func NewAccessManagerHandler(
	roleRepo core.AccessManagerRoleRepository,
	memberRepo core.AccessManagerRoleMemberRepository,
	contractRepo core.AccessManagerManagedContractRepository,
	fnPermRepo core.AccessManagerFunctionPermissionRepository,
	scheduledOpRepo core.AccessManagerScheduledOperationRepository,
	eventLogRepo core.AccessManagerEventLogRepository,
	log logger.Logger,
) *AccessManagerHandler {
	return &AccessManagerHandler{
		roleRepo:        roleRepo,
		memberRepo:      memberRepo,
		contractRepo:    contractRepo,
		fnPermRepo:      fnPermRepo,
		scheduledOpRepo: scheduledOpRepo,
		eventLogRepo:    eventLogRepo,
		log:             log,
	}
}

// ListRoles returns all indexed roles.
// @Summary List access manager roles
// @Tags access-manager
// @Produce json
// @Success 200 {array} map[string]any
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/roles [get]
func (h *AccessManagerHandler) ListRoles(c *gin.Context) {
	roles, err := h.roleRepo.List(c.Request.Context())
	if err != nil {
		h.log.Error("Failed to list roles", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list roles"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// GetRole returns a single role by its uint64 ID.
// @Summary Get role by ID
// @Tags access-manager
// @Produce json
// @Param roleId path int true "Role ID"
// @Success 200 {object} map[string]any
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/roles/{roleId} [get]
func (h *AccessManagerHandler) GetRole(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roleId must be a positive integer"})
		return
	}
	role, err := h.roleRepo.FindByID(c.Request.Context(), roleID)
	if err != nil {
		HandleError(c, h.log, err)
		return
	}
	c.JSON(http.StatusOK, role)
}

// ListRoleMembers returns members of a role.
// @Summary List role members
// @Tags access-manager
// @Produce json
// @Param roleId path  int  true  "Role ID"
// @Param active query bool false "Filter active-only members (default true)"
// @Success 200 {array} map[string]any
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/roles/{roleId}/members [get]
func (h *AccessManagerHandler) ListRoleMembers(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roleId must be a positive integer"})
		return
	}
	activeOnly := c.DefaultQuery("active", "true") == "true"
	members, err := h.memberRepo.ListByRole(c.Request.Context(), roleID, activeOnly)
	if err != nil {
		h.log.Error("Failed to list role members", "roleId", roleID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list role members"})
		return
	}
	c.JSON(http.StatusOK, members)
}

// ListAccountRoles returns all roles assigned to a wallet address.
// @Summary List roles for a wallet address
// @Tags access-manager
// @Produce json
// @Param address path string true "Wallet address (0x...)"
// @Success 200 {array} map[string]any
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/accounts/{address}/roles [get]
func (h *AccessManagerHandler) ListAccountRoles(c *gin.Context) {
	address := c.Param("address")
	members, err := h.memberRepo.ListByAccount(c.Request.Context(), address)
	if err != nil {
		h.log.Error("Failed to list account roles", "address", address, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list account roles"})
		return
	}
	c.JSON(http.StatusOK, members)
}

// ListManagedContracts returns all registered managed contracts.
// @Summary List managed contracts
// @Tags access-manager
// @Produce json
// @Success 200 {array} map[string]any
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/managed-contracts [get]
func (h *AccessManagerHandler) ListManagedContracts(c *gin.Context) {
	contracts, err := h.contractRepo.List(c.Request.Context())
	if err != nil {
		h.log.Error("Failed to list managed contracts", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list managed contracts"})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

// GetManagedContract returns a single managed contract by address.
// @Summary Get managed contract
// @Tags access-manager
// @Produce json
// @Param address path string true "Contract address (0x...)"
// @Success 200 {object} map[string]any
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/managed-contracts/{address} [get]
func (h *AccessManagerHandler) GetManagedContract(c *gin.Context) {
	address := c.Param("address")
	contract, err := h.contractRepo.FindByAddress(c.Request.Context(), address)
	if err != nil {
		HandleError(c, h.log, err)
		return
	}
	c.JSON(http.StatusOK, contract)
}

// ListFunctionPermissions returns function-level permissions for a contract.
// @Summary List function permissions
// @Tags access-manager
// @Produce json
// @Param address path string true "Contract address (0x...)"
// @Success 200 {array} map[string]any
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/managed-contracts/{address}/permissions [get]
func (h *AccessManagerHandler) ListFunctionPermissions(c *gin.Context) {
	address := c.Param("address")
	perms, err := h.fnPermRepo.ListByContract(c.Request.Context(), address)
	if err != nil {
		h.log.Error("Failed to list function permissions", "address", address, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list function permissions"})
		return
	}
	c.JSON(http.StatusOK, perms)
}

// ListScheduledOperations returns scheduled timelock operations.
// @Summary List scheduled operations
// @Tags access-manager
// @Produce json
// @Param status query string false "Filter by status: scheduled, executed, canceled"
// @Success 200 {array} map[string]any
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/scheduled-operations [get]
func (h *AccessManagerHandler) ListScheduledOperations(c *gin.Context) {
	status := c.Query("status")
	ops, err := h.scheduledOpRepo.List(c.Request.Context(), status)
	if err != nil {
		h.log.Error("Failed to list scheduled operations", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list scheduled operations"})
		return
	}
	c.JSON(http.StatusOK, ops)
}

// ListEvents returns the paginated raw event audit log.
// @Summary List access manager events
// @Tags access-manager
// @Produce json
// @Param event query string false "Filter by event name (e.g. RoleGranted)"
// @Param page  query int    false "Page number (default 1)"
// @Param limit query int    false "Items per page (default 20)"
// @Success 200 {object} map[string]any
// @Failure 500 {object} map[string]string
// @Router /api/v1/admin/events [get]
func (h *AccessManagerHandler) ListEvents(c *gin.Context) {
	eventName := c.Query("event")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}

	events, total, err := h.eventLogRepo.List(c.Request.Context(), eventName, page, limit)
	if err != nil {
		h.log.Error("Failed to list access manager events", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": events,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
