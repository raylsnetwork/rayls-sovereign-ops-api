package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/utils"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

// RespondIfReverted maps an on-chain revert error to 422 Unprocessable Entity (the request was
// well-formed but the contract rejected it), surfacing the revert detail as a hint. It returns true
// when it handled the error so the caller can return; otherwise the caller falls through to
// HandleError. Used by operator-signed write handlers (token register, onboarding) to keep revert
// responses consistent with the mint/burn and deploy handlers.
func RespondIfReverted(c *gin.Context, log logger.Logger, err error) bool {
	if !errors.Is(err, core.ErrTxReverted) {
		return false
	}
	log.Warn("On-chain transaction reverted", "error", err)
	c.JSON(
		http.StatusUnprocessableEntity,
		utils.ErrorResponse{Error: "transaction reverted on-chain", Hint: err.Error()},
	)
	return true
}

// HandleError maps domain errors to HTTP status codes
func HandleError(c *gin.Context, log logger.Logger, err error) {
	var notFoundErr *core.NotFoundError
	var validationErr *core.ValidationError
	var internalErr *core.InternalError
	var unauthorizedErr *core.UnauthorizedError
	var forbiddenErr *core.ForbiddenError
	var roleAssignmentPendingErr *core.RoleAssignmentPendingError
	var walletRegisteredErr *core.WalletRegisteredError
	var emailAlreadyLinkedErr *core.EmailAlreadyLinkedError
	var accountSuspendedErr *core.AccountSuspendedError
	var serviceUnavailableErr *core.ServiceUnavailableError
	var noOperatorSignerErr *core.NoOperatorSignerError
	var bootstrapAlreadyCompletedErr *core.BootstrapAlreadyCompletedError

	switch {
	case errors.As(err, &roleAssignmentPendingErr):
		log.Debug("Role assignment pending")
		c.JSON(http.StatusForbidden, gin.H{
			"code":    "ROLE_ASSIGNMENT_PENDING",
			"message": "Your account has no role assigned yet. Contact your operator.",
		})

	case errors.As(err, &walletRegisteredErr):
		log.Debug("Wallet registered without a role")
		c.JSON(http.StatusForbidden, gin.H{
			"code":    "WALLET_REGISTERED",
			"message": "This wallet is now registered but has no role yet. Contact your operator.",
		})

	case errors.As(err, &emailAlreadyLinkedErr):
		log.Debug("Email already linked to another provider", "provider", emailAlreadyLinkedErr.Provider)
		c.JSON(http.StatusConflict, gin.H{
			"code":    "EMAIL_ALREADY_LINKED",
			"message": emailAlreadyLinkedErr.Message(),
		})

	case errors.As(err, &accountSuspendedErr):
		log.Debug("Account suspended")
		c.JSON(http.StatusForbidden, gin.H{
			"code":    "ACCOUNT_SUSPENDED",
			"message": "Your account has been suspended. Contact your operator.",
		})

	case errors.As(err, &serviceUnavailableErr):
		log.Warn("Service unavailable (RPC error at login)", "error", err)
		c.Header("Retry-After", "30")
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    "SERVICE_UNAVAILABLE",
			"message": "Service temporarily unavailable. Please try again shortly.",
		})

	case errors.As(err, &noOperatorSignerErr):
		log.Warn("No operator signer available", "error", err)
		c.Header("Retry-After", "30")
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    "NO_OPERATOR_SIGNER",
			"message": "No operator signer available. Please try again shortly.",
		})

	case errors.As(err, &bootstrapAlreadyCompletedErr):
		log.Debug("Bootstrap already completed")
		body := gin.H{
			"code":    "BOOTSTRAP_ALREADY_COMPLETED",
			"message": "Bootstrap has already been completed.",
		}
		// The existing admin's wallet, when known. Accounts are shared across chains but
		// on-chain roles are not, so a caller provisioning a new chain needs this address
		// to grant roles there — 409 means "already exists", not "nothing to do".
		if bootstrapAlreadyCompletedErr.Address != "" {
			body["address"] = bootstrapAlreadyCompletedErr.Address
		}
		c.JSON(http.StatusConflict, body)

	case errors.As(err, &unauthorizedErr):
		log.Debug("Unauthorized", "message", unauthorizedErr.Message)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  "INVALID_CREDENTIAL",
			"error": unauthorizedErr.Message,
		})

	case errors.As(err, &forbiddenErr):
		log.Debug("Forbidden", "message", forbiddenErr.Message)
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})

	case errors.As(err, &notFoundErr):
		log.Debug("Resource not found", "resource", notFoundErr.Resource, "id", notFoundErr.ID)
		c.JSON(http.StatusNotFound, gin.H{"error": notFoundErr.Error()})

	case errors.Is(err, core.ErrWalletNotFound):
		log.Debug("Wallet not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})

	case errors.As(err, &validationErr):
		log.Debug("Validation error", "field", validationErr.Field, "message", validationErr.Message)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})

	case errors.As(err, &internalErr):
		log.Error("Internal error", "error", internalErr.Err, "operation", internalErr.Operation)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "INTERNAL_ERROR",
			"error": "An unexpected error occurred. Please try again or contact support.",
		})

	default:
		log.Error("Unknown error", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "INTERNAL_ERROR",
			"error": "An unexpected error occurred. Please try again or contact support.",
		})
	}
}
