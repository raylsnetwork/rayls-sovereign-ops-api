package auth

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

type siweVerifyRequest struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
	Nonce     string `json:"nonce"`
}

// SIWEProvider implements the go-pkgz/auth provider.Provider interface for SIWE authentication
type SIWEProvider struct {
	authSvc      core.AuthService
	tokenWrapper *TokenWrapper
	log          logger.Logger
}

// NewSIWEProvider creates a new SIWE provider
func NewSIWEProvider(authSvc core.AuthService, tokenWrapper *TokenWrapper, log logger.Logger) *SIWEProvider {
	return &SIWEProvider{
		authSvc:      authSvc,
		tokenWrapper: tokenWrapper,
		log:          log,
	}
}

// Name returns the provider name
func (p *SIWEProvider) Name() string { return "siwe" }

// LoginHandler generates a SIWE challenge (nonce + message) for the given wallet address
func (p *SIWEProvider) LoginHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if !isValidEthAddress(address) {
		http.Error(w, `{"error":"address must be a valid Ethereum address (0x + 40 hex chars)"}`, http.StatusBadRequest)
		return
	}

	message, nonce, err := p.authSvc.GenerateChallenge(r.Context(), address)
	if err != nil {
		p.log.Error("Failed to generate SIWE challenge", "error", err)
		http.Error(w, `{"error":"failed to generate challenge"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": message,
		"nonce":   nonce,
	})
}

// AuthHandler verifies a SIWE signature, applies the v1.4 decision tree, and issues tokens.
// Expects a POST with a JSON body containing address, signature, and nonce.
func (p *SIWEProvider) AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed, use POST"}`, http.StatusMethodNotAllowed)
		return
	}

	var req siweVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid JSON body"}`, http.StatusBadRequest)
		return
	}

	if req.Signature == "" || req.Nonce == "" {
		http.Error(w, `{"error":"address, signature, and nonce are required"}`, http.StatusBadRequest)
		return
	}
	if !isValidEthAddress(req.Address) {
		http.Error(w, `{"error":"address must be a valid Ethereum address (0x + 40 hex chars)"}`, http.StatusBadRequest)
		return
	}

	user, roles, err := p.authSvc.VerifySIWE(r.Context(), req.Address, req.Signature, req.Nonce)
	if err != nil {
		p.log.Debug("SIWE login blocked", "address", req.Address, "reason", err.Error())
		status, body := p.errorResponse(err)
		w.Header().Set("Content-Type", "application/json")
		if status == http.StatusServiceUnavailable {
			w.Header().Set("Retry-After", "30")
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := p.tokenWrapper.IssueToken(r.Context(), w, user, "siwe", roles); err != nil {
		p.log.Error("Failed to issue token", "error", err)
		http.Error(w, `{"error":"failed to issue token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, _, err := p.tokenWrapper.CreateRefreshTokenString(user.ID.String())
	if err != nil {
		p.log.Error("Failed to create refresh token", "error", err)
		http.Error(w, `{"error":"failed to create refresh token"}`, http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "authenticated",
		"refresh_token": refreshToken,
		"user": map[string]interface{}{
			"id":   user.ID.String(),
			"name": user.Name,
		},
	})
}

// isValidEthAddress checks that the address is 0x followed by 40 hex characters.
func isValidEthAddress(addr string) bool {
	if len(addr) != 42 || !strings.HasPrefix(addr, "0x") {
		return false
	}
	_, err := hex.DecodeString(addr[2:])
	return err == nil
}

// errorResponse maps a domain error to an HTTP status code and JSON error body.
func (p *SIWEProvider) errorResponse(err error) (int, string) {
	jsonBody := func(code, msg string) string {
		b, _ := json.Marshal(map[string]string{"code": code, "message": msg})
		return string(b)
	}

	var unauth *core.UnauthorizedError
	if errors.As(err, &unauth) {
		return http.StatusUnauthorized, jsonBody("INVALID_CREDENTIAL", "Authentication failed. Please try again.")
	}

	var walletRegistered *core.WalletRegisteredError
	if errors.As(err, &walletRegistered) {
		return http.StatusForbidden, jsonBody(
			"WALLET_REGISTERED",
			"This wallet is now registered but has no role yet. Contact your operator.",
		)
	}

	var pending *core.RoleAssignmentPendingError
	if errors.As(err, &pending) {
		return http.StatusForbidden, jsonBody(
			"ROLE_ASSIGNMENT_PENDING",
			"Your account has no role assigned yet. Contact your operator.",
		)
	}

	var linked *core.EmailAlreadyLinkedError
	if errors.As(err, &linked) {
		return http.StatusConflict, jsonBody("EMAIL_ALREADY_LINKED", linked.Message())
	}

	var suspended *core.AccountSuspendedError
	if errors.As(err, &suspended) {
		return http.StatusForbidden, jsonBody(
			"ACCOUNT_SUSPENDED",
			"Your account has been suspended. Contact your operator.",
		)
	}

	var svcUnavail *core.ServiceUnavailableError
	if errors.As(err, &svcUnavail) {
		return http.StatusServiceUnavailable, jsonBody(
			"SERVICE_UNAVAILABLE",
			"Service temporarily unavailable. Please try again shortly.",
		)
	}

	return http.StatusInternalServerError, jsonBody(
		"INTERNAL_ERROR",
		"An unexpected error occurred. Please try again or contact support.",
	)
}

// LogoutHandler clears the auth token
func (p *SIWEProvider) LogoutHandler(w http.ResponseWriter, _ *http.Request) {
	p.tokenWrapper.ClearCookie(w)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "logged out"})
}
