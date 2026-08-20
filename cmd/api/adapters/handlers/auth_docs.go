package handlers

// This file contains Swagger-only documentation for auth endpoints.
// These are not real Gin handlers — they exist solely to generate OpenAPI specs.

type siweVerifyRequest struct {
	Address   string `json:"address" example:"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"`
	Signature string `json:"signature" example:"0xabc123..."`
	Nonce     string `json:"nonce" example:"abc123xyz"`
}

// googleLogin godoc
// @Summary Google OAuth login
// @Description Redirects the user to Google's OAuth consent page. On success redirects to callback.
// @Tags auth
// @Produce json
// @Success 302 {string} string "Redirect to Google"
// @Router /auth/google/login [get]
func googleLogin() {}

// googleCallback godoc
// @Summary Google OAuth callback
// @Description Handles the OAuth callback from Google. Issues JWT cookie + refresh token, or returns 403 if approval pending.
// @Tags auth
// @Produce json
// @Param code query string true "Authorization code from Google"
// @Param state query string true "CSRF state"
// @Success 200 {object} map[string]interface{} "authenticated"
// @Failure 401 {object} map[string]string "INVALID_CREDENTIAL"
// @Failure 403 {object} map[string]string "ROLE_ASSIGNMENT_PENDING, WALLET_REGISTERED or ACCOUNT_SUSPENDED"
// @Failure 409 {object} map[string]string "EMAIL_ALREADY_LINKED"
// @Failure 503 {object} map[string]string "SERVICE_UNAVAILABLE"
// @Router /auth/google/callback [get]
func googleCallback() {}

// microsoftLogin godoc
// @Summary Microsoft OAuth login
// @Description Redirects the user to Microsoft's OAuth consent page.
// @Tags auth
// @Produce json
// @Success 302 {string} string "Redirect to Microsoft"
// @Router /auth/microsoft/login [get]
func microsoftLogin() {}

// microsoftCallback godoc
// @Summary Microsoft OAuth callback
// @Description Handles the OAuth callback from Microsoft. Issues JWT cookie + refresh token, or returns 403 if approval pending.
// @Tags auth
// @Produce json
// @Param code query string true "Authorization code from Microsoft"
// @Param state query string true "CSRF state"
// @Success 200 {object} map[string]interface{} "authenticated"
// @Failure 401 {object} map[string]string "INVALID_CREDENTIAL"
// @Failure 403 {object} map[string]string "ROLE_ASSIGNMENT_PENDING, WALLET_REGISTERED or ACCOUNT_SUSPENDED"
// @Failure 409 {object} map[string]string "EMAIL_ALREADY_LINKED"
// @Router /auth/microsoft/callback [get]
func microsoftCallback() {}

// siweLogin godoc
// @Summary SIWE challenge
// @Description Generates a Sign-In with Ethereum challenge (EIP-4361 message + nonce) for the given wallet address.
// @Tags auth
// @Produce json
// @Param address query string true "Ethereum wallet address (0x + 40 hex chars)"
// @Success 200 {object} map[string]string "message and nonce"
// @Failure 400 {object} map[string]string
// @Router /auth/siwe/login [get]
func siweLogin() {}

// siweCallback godoc
// @Summary SIWE verify signature
// @Description Verifies the signed SIWE message. Auto-registers new wallets and applies the v1.4 decision tree. Issues JWT cookie + refresh_token on success, or returns 403/503 based on account state.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body siweVerifyRequest true "Signed SIWE payload"
// @Success 200 {object} map[string]interface{} "authenticated with refresh_token"
// @Failure 401 {object} map[string]string "INVALID_CREDENTIAL"
// @Failure 403 {object} map[string]string "ROLE_ASSIGNMENT_PENDING, WALLET_REGISTERED or ACCOUNT_SUSPENDED"
// @Failure 409 {object} map[string]string "EMAIL_ALREADY_LINKED"
// @Failure 503 {object} map[string]string "SERVICE_UNAVAILABLE"
// @Router /auth/siwe/callback [post]
func siweCallback() {}
