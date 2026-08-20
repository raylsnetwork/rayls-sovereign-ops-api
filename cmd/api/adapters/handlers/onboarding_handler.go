package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/middleware"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// OnboardingHandler handles self-service onboarding of a user's address pairs into RNUserGovernance.
type OnboardingHandler struct {
	svc core.OnboardingService
	log logger.Logger
}

func NewOnboardingHandler(svc core.OnboardingService, log logger.Logger) *OnboardingHandler {
	return &OnboardingHandler{svc: svc, log: log}
}

// addressPairResponse is the 201 UserAddressPair response shape.
type addressPairResponse struct {
	PublicChainAddress  string    `json:"public_chain_address"`
	PrivateChainAddress string    `json:"private_chain_address"`
	Status              uint8     `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
}

// AddressPairListFilter declares the query parameters accepted by ListMine, consumed by the
// query-validation middleware. Only `status=0` (pending) is accepted; omitting it returns all pairs.
type AddressPairListFilter struct {
	Status string `form:"status" enums:"0"`
}

// pendingUserAddressPairsResponse groups a user's pending pairs under the resolved ops-api UUID,
// reusing the addressPairResponse shape for each pair.
type pendingUserAddressPairsResponse struct {
	UserID       string                `json:"user_id"`
	AddressPairs []addressPairResponse `json:"address_pairs"`
}

// setApprovalStatusRequest is the PATCH body. The user identity is taken from the path (:id), never
// from the body. Allowed status values are 0 (pending), 1 (approved), and 2 (rejected).
type setApprovalStatusRequest struct {
	PublicAddress  string `json:"public_address"`
	PrivateAddress string `json:"private_address"`
	Status         uint8  `json:"status"`
}

func toAddressPairResponse(p core.OnChainAddressPair) addressPairResponse {
	return addressPairResponse{
		PublicChainAddress:  domain.ChecksumAddress(p.PublicChainAddress),
		PrivateChainAddress: domain.ChecksumAddress(p.PrivateChainAddress),
		Status:              uint8(p.Status),
		CreatedAt:           p.CreatedAt,
	}
}

// AddAddressPair creates a fresh HSM wallet pair and registers it in RNUserGovernance.
// @Summary Add an address pair
// @Description Creates a fresh HSM wallet pair (private + public chain) and registers it on-chain in
// @Description RNUserGovernance as a pending pair. The on-chain write is operator-signed. Takes no body;
// @Description each call creates a new pair, so a user accumulates many pairs over repeated calls.
// @Tags onboarding
// @Produce json
// @Success 201 {object} addressPairResponse
// @Failure 401 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/me/address-pairs [post]
func (h *OnboardingHandler) AddAddressPair(c *gin.Context) {
	jwtUser, ok := middleware.GetAuthUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	userID, err := uuid.Parse(jwtUser.ID)
	if err != nil {
		h.log.Error("Invalid user ID in JWT", "id", jwtUser.ID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user session"})
		return
	}

	pair, err := h.svc.AddAddressPair(c.Request.Context(), userID)
	if err != nil {
		if RespondIfReverted(c, h.log, err) {
			return
		}
		HandleError(c, h.log, err)
		return
	}

	c.JSON(http.StatusCreated, toAddressPairResponse(*pair))
}

// ListMine returns the authenticated caller's on-chain address pairs.
// @Summary List own address pairs
// @Description Returns the authenticated user's on-chain address pairs from RNUserGovernance, derived
// @Description from the JWT (keccak256 of the user ID). Omitting the status filter returns all pairs;
// @Description status=0 returns only pending pairs.
// @Tags onboarding
// @Produce json
// @Param status query int false "Filter by approval status" Enums(0)
// @Success 200 {array} addressPairResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/me/address-pairs [get]
func (h *OnboardingHandler) ListMine(c *gin.Context) {
	jwtUser, ok := middleware.GetAuthUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	userID, err := uuid.Parse(jwtUser.ID)
	if err != nil {
		h.log.Error("Invalid user ID in JWT", "id", jwtUser.ID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user session"})
		return
	}

	var status *domain.ApprovalStatus
	if raw := c.Query("status"); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil || domain.ApprovalStatus(v) != domain.ApprovalStatusPending {
			c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 0 (pending)"})
			return
		}
		s := domain.ApprovalStatus(v)
		status = &s
	}

	pairs, err := h.svc.ListMine(c.Request.Context(), userID, status)
	if err != nil {
		HandleError(c, h.log, err)
		return
	}

	items := make([]addressPairResponse, 0, len(pairs))
	for _, p := range pairs {
		items = append(items, toAddressPairResponse(p))
	}
	c.JSON(http.StatusOK, items)
}

// ListAllPending returns every pending address pair across all users, owner resolved to UUID.
// @Summary List all pending address pairs (operator)
// @Description Returns every pending address pair across all users, read on-chain via
// @Description getAllPendingAddressPairs and grouped under the resolved ops-api user UUID. Hashes that
// @Description do not resolve to a known user are skipped. Requires the PRIVACY_NODE_OPERATOR role.
// @Tags onboarding
// @Produce json
// @Success 200 {array} pendingUserAddressPairsResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/v1/admin/address-pairs/pending [get]
func (h *OnboardingHandler) ListAllPending(c *gin.Context) {
	groups, err := h.svc.ListAllPending(c.Request.Context())
	if err != nil {
		HandleError(c, h.log, err)
		return
	}

	items := make([]pendingUserAddressPairsResponse, 0, len(groups))
	for _, g := range groups {
		pairs := make([]addressPairResponse, 0, len(g.AddressPairs))
		for _, p := range g.AddressPairs {
			pairs = append(pairs, toAddressPairResponse(p))
		}
		items = append(items, pendingUserAddressPairsResponse{
			UserID:       g.UserID.String(),
			AddressPairs: pairs,
		})
	}
	c.JSON(http.StatusOK, items)
}

// SetApprovalStatus sets the approval status of an address pair for the user identified by the path :id.
// @Summary Set address pair approval status (operator)
// @Description Sets an address pair's approval status: 0 (pending), 1 (approved), or 2 (rejected). This
// @Description includes reverting an already-decided pair back to pending. The target user is taken from
// @Description the path :id, never the body; the on-chain write is operator-signed.
// @Description Requires the PRIVACY_NODE_OPERATOR role.
// @Tags onboarding
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body setApprovalStatusRequest true "Approval decision"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /api/v1/admin/users/{id}/address-pairs/status [patch]
func (h *OnboardingHandler) SetApprovalStatus(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a valid UUID"})
		return
	}

	var req setApprovalStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	status := domain.ApprovalStatus(req.Status)
	if status > domain.ApprovalStatusRejected {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 0 (pending), 1 (approved), or 2 (rejected)"})
		return
	}

	if !common.IsHexAddress(req.PublicAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "public_address must be a valid hex address"})
		return
	}
	if !common.IsHexAddress(req.PrivateAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "private_address must be a valid hex address"})
		return
	}

	if err := h.svc.SetApprovalStatus(
		c.Request.Context(),
		userID,
		req.PublicAddress,
		req.PrivateAddress,
		status,
	); err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		if RespondIfReverted(c, h.log, err) {
			return
		}
		HandleError(c, h.log, err)
		return
	}

	c.Status(http.StatusOK)
}
