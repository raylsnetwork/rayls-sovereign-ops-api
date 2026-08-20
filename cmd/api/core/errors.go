package core

import (
	"fmt"

	"github.com/raylsnetwork/rayls-privacy-ops-api/withstack"
)

// ============================================================================
// Domain Errors
// These errors express business-level failures
// HTTP adapters map these to appropriate status codes (404, 400, etc.)
// ============================================================================

// NotFoundError indicates a requested resource does not exist
type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: %s", e.Resource, e.ID)
}

// NewNotFoundError creates a NotFoundError
func NewNotFoundError(resource, id string) *NotFoundError {
	return &NotFoundError{
		Resource: resource,
		ID:       id,
	}
}

// ValidationError indicates invalid input parameters
type ValidationError struct {
	Field   string // The field that failed validation
	Message string
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// NewValidationError creates a ValidationError
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// InternalError wraps unexpected errors that should be logged and returned as 500
type InternalError struct {
	Operation string
	Err       error
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("internal error during %s: %v", e.Operation, e.Err)
}

func (e *InternalError) Unwrap() error {
	return e.Err
}

// NewInternalError creates an InternalError with stack trace
func NewInternalError(operation string, err error) *InternalError {
	return &InternalError{
		Operation: operation,
		Err:       withstack.Wrap(err),
	}
}

// UnauthorizedError indicates missing or invalid authentication
type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized: %s", e.Message)
}

// NewUnauthorizedError creates an UnauthorizedError
func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{Message: message}
}

// ForbiddenError indicates insufficient permissions
type ForbiddenError struct {
	Message string
}

func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("forbidden: %s", e.Message)
}

// NewForbiddenError creates a ForbiddenError
func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{Message: message}
}

// EmailAlreadyLinkedError is returned when a login attempt presents an email that already
// belongs to an account authenticated by a DIFFERENT provider. Maps to HTTP 409 with code
// EMAIL_ALREADY_LINKED.
//
// This is not a permission problem: the account exists and may well be fully approved. We
// refuse because the incoming identity is unproven — the email sign-up path has no
// verification step, so honouring it would hand an established account to whoever types its
// address. Provider names the one the account is already using, so the caller can tell the
// user which button to press instead of guessing.
type EmailAlreadyLinkedError struct {
	Provider string
}

func (e *EmailAlreadyLinkedError) Error() string { return "EMAIL_ALREADY_LINKED" }

// Message explains that the address belongs to an account using another login method, naming
// that method when it is known. Provider is "" when the lookup failed, in which case the text
// stays generic rather than pointing the user at a method we cannot name. It lives here rather
// than in a handler because both the auth package and the Gin error handler render it.
func (e *EmailAlreadyLinkedError) Message() string {
	switch e.Provider {
	case "":
		return "This email is already registered with a different sign-in method. " +
			"Please use the method you signed up with."
	case "siwe":
		return "This email is already registered to an account that signs in with a wallet. " +
			"Please connect your wallet instead."
	default:
		return fmt.Sprintf(
			"This email is already registered with %s sign-in. Please continue with %s instead.",
			e.Provider, e.Provider,
		)
	}
}

// RoleAssignmentPendingError is returned when an account exists and is authenticated but has
// no role yet: either it is still waiting_role_assignment with no provisioner wired to
// activate it, or it is role_assigned and the AccessManager reports no on-chain roles.
// Maps to HTTP 403 with code ROLE_ASSIGNMENT_PENDING.
type RoleAssignmentPendingError struct{}

func (e *RoleAssignmentPendingError) Error() string { return "ROLE_ASSIGNMENT_PENDING" }

// WalletRegisteredError is returned on a first SIWE login: the wallet was just auto-registered
// but no role has been granted to it yet, so there is nothing to issue a session against.
// Maps to HTTP 403 with code WALLET_REGISTERED.
//
// Distinct from RoleAssignmentPendingError because the remedy differs — the wallet is new and
// the operator has to grant it a role, rather than an existing account waiting on activation.
type WalletRegisteredError struct{}

func (e *WalletRegisteredError) Error() string { return "WALLET_REGISTERED" }

// AccountSuspendedError is returned when a user's account has is_active=false.
// Maps to HTTP 403 with code ACCOUNT_SUSPENDED.
type AccountSuspendedError struct{}

func (e *AccountSuspendedError) Error() string { return "ACCOUNT_SUSPENDED" }

// ServiceUnavailableError is returned when an on-chain RPC call fails at login.
// Maps to HTTP 503 with code SERVICE_UNAVAILABLE and Retry-After: 30.
type ServiceUnavailableError struct{}

func (e *ServiceUnavailableError) Error() string { return "SERVICE_UNAVAILABLE" }

// BootstrapAlreadyCompletedError is returned when POST /admin/bootstrap is called but a user already exists.
// Maps to HTTP 409 with code BOOTSTRAP_ALREADY_COMPLETED.
//
// Address carries the existing admin's custody wallet when it could be resolved. Accounts
// live in the SHARED identity database, so every chain after the first sees this 409 — but
// the ON-CHAIN roles are per-chain and still have to be granted on each one. Returning the
// address makes the conflict actionable: the caller provisioning a new chain can grant to
// the wallet that already exists instead of treating 409 as "nothing to do". Empty when no
// wallet is on file (the user exists but has none).
type BootstrapAlreadyCompletedError struct {
	Address string
}

func (e *BootstrapAlreadyCompletedError) Error() string { return "BOOTSTRAP_ALREADY_COMPLETED" }

// NoOperatorSignerError is returned when no eligible HSM-custodied PRIVACY_NODE_OPERATOR member
// exists to sign an operator-authority governance write. Maps to HTTP 503 with code
// NO_OPERATOR_SIGNER and Retry-After: 30. Failing closed prevents signing with an unintended key.
type NoOperatorSignerError struct{}

func (e *NoOperatorSignerError) Error() string { return "NO_OPERATOR_SIGNER" }
