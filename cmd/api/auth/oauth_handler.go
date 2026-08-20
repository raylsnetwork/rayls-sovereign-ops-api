package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	microsoftOAuth "golang.org/x/oauth2/microsoft"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

const (
	oauthStateCookieTTL  = 10 * time.Minute
	oauthUserInfoTimeout = 10 * time.Second
)

// oauthState is stored in the state cookie during the OAuth flow.
type oauthState struct {
	State      string `json:"state"`
	RedirectTo string `json:"redirect_to,omitempty"`
}

// relayState is the (base64url) OAuth `state` parameter. It carries the CSRF
// token plus this instance's own callback URL, so a fixed relay host can bounce
// the OAuth callback back to the right instance.
type relayState struct {
	C string `json:"c"` // CSRF token
	R string `json:"r"` // return-to: the instance's own /auth/<provider>/callback
}

type googleUserInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	EmailVerified bool   `json:"email_verified"`
}

type microsoftUserInfo struct {
	ID                string `json:"id"`
	DisplayName       string `json:"displayName"`
	Mail              string `json:"mail"`
	UserPrincipalName string `json:"userPrincipalName"`
}

// OAuthHandler handles Google and Microsoft OAuth2 login flows.
// It replaces go-pkgz/auth's OAuth providers with a custom implementation
// that fully applies the v1.4 login decision tree before issuing tokens.
type OAuthHandler struct {
	authSvc              core.AuthService
	tokenWrapper         *TokenWrapper
	googleCfg            *oauth2.Config
	microsoftCfg         *oauth2.Config
	postLoginRedirectURL string
	// baseURL is THIS instance's own host; carried in the OAuth state so a fixed
	// relay can bounce the callback back here. Empty = no relay (single-host).
	baseURL string
	// emailSignupEnabled gates POST /auth/signup. Off by default — see
	// config.Auth.EmailSignupEnabled and EmailSignup.
	emailSignupEnabled bool
	// signupDetailsRepo persists the sign-up form's profile answers. Only the email
	// path collects them; nil on the per-chain ops-api, whose database has no
	// user_signup_details table (identity tables were dropped in migration 000009).
	signupDetailsRepo core.UserSignupDetailsRepository
	log               logger.Logger
}

// NewOAuthHandler creates an OAuthHandler. Configs with nil value are treated as disabled.
// postLoginRedirectURL is the frontend URL to redirect to after successful login (empty = return JSON).
// baseURL is this instance's own host (for the relay bounce target).
// emailSignupEnabled turns on the unverified email signup endpoint (dev only).
// signupDetailsRepo stores the sign-up form's profile answers; nil disables that (the
// session is still issued — see EmailSignup).
func NewOAuthHandler(
	authSvc core.AuthService,
	tokenWrapper *TokenWrapper,
	googleCfg *oauth2.Config,
	microsoftCfg *oauth2.Config,
	postLoginRedirectURL string,
	baseURL string,
	emailSignupEnabled bool,
	signupDetailsRepo core.UserSignupDetailsRepository,
	log logger.Logger,
) *OAuthHandler {
	return &OAuthHandler{
		authSvc:              authSvc,
		tokenWrapper:         tokenWrapper,
		googleCfg:            googleCfg,
		microsoftCfg:         microsoftCfg,
		postLoginRedirectURL: postLoginRedirectURL,
		baseURL:              baseURL,
		emailSignupEnabled:   emailSignupEnabled,
		signupDetailsRepo:    signupDetailsRepo,
		log:                  log,
	}
}

// NewGoogleOAuthConfig builds the oauth2.Config for Google. redirectBase is the
// host the redirect_uri is registered under (a fixed relay host, or BASE_URL).
func NewGoogleOAuthConfig(clientID, clientSecret, redirectBase string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectBase + "/auth/google/callback",
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

// NewMicrosoftOAuthConfig builds the oauth2.Config for Microsoft (common tenant).
func NewMicrosoftOAuthConfig(clientID, clientSecret, redirectBase string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectBase + "/auth/microsoft/callback",
		Scopes:       []string{"openid", "email", "profile", "User.Read"},
		Endpoint:     microsoftOAuth.AzureADEndpoint("common"),
	}
}

// GoogleLogin redirects the user to the Google consent page.
func (h *OAuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	h.startLogin(w, r, h.googleCfg, "google")
}

// GoogleCallback handles the callback from Google.
func (h *OAuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	h.handleCallback(w, r, h.googleCfg, "google", h.fetchGoogleUserInfo)
}

// MicrosoftLogin redirects the user to the Microsoft consent page.
func (h *OAuthHandler) MicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	h.startLogin(w, r, h.microsoftCfg, "microsoft")
}

// MicrosoftCallback handles the callback from Microsoft.
func (h *OAuthHandler) MicrosoftCallback(w http.ResponseWriter, r *http.Request) {
	h.handleCallback(w, r, h.microsoftCfg, "microsoft", h.fetchMicrosoftUserInfo)
}

func (h *OAuthHandler) startLogin(w http.ResponseWriter, r *http.Request, cfg *oauth2.Config, provider string) {
	if cfg == nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s OAuth not configured"}`, provider), http.StatusNotFound)
		return
	}

	stateBytes := make([]byte, 16)
	if _, err := rand.Read(stateBytes); err != nil {
		h.log.Error("Failed to generate OAuth state", "provider", provider, "error", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	redirectTo := r.URL.Query().Get("redirect_to")
	if redirectTo == "" {
		redirectTo = h.postLoginRedirectURL
	}

	// The OAuth `state` carries the CSRF token + this instance's own callback URL
	// so a fixed relay can bounce the callback back here. The instance later
	// validates state == the cookie (CSRF), so an attacker can't forge it.
	relayJSON, err := json.Marshal(relayState{
		C: hex.EncodeToString(stateBytes),
		R: h.baseURL + "/auth/" + provider + "/callback",
	})
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}
	stateToken := base64.RawURLEncoding.EncodeToString(relayJSON)

	jsonVal, err := json.Marshal(oauthState{State: stateToken, RedirectTo: redirectTo})
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state_" + provider,
		Value:    base64.RawURLEncoding.EncodeToString(jsonVal),
		Path:     "/",
		MaxAge:   int(oauthStateCookieTTL.Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, cfg.AuthCodeURL(stateToken, oauth2.AccessTypeOnline), http.StatusTemporaryRedirect)
}

type userInfoFetcher func(ctx context.Context, token *oauth2.Token, cfg *oauth2.Config) (oauthID, name, email string, emailVerified bool, err error)

func (h *OAuthHandler) handleCallback(
	w http.ResponseWriter,
	r *http.Request,
	cfg *oauth2.Config,
	provider string,
	fetchUserInfo userInfoFetcher,
) {
	if cfg == nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s OAuth not configured"}`, provider), http.StatusNotFound)
		return
	}

	stateCookie, err := r.Cookie("oauth_state_" + provider)
	if err != nil {
		http.Error(
			w,
			`{"code":"INVALID_CREDENTIAL","message":"Invalid OAuth state. Please try again."}`,
			http.StatusUnauthorized,
		)
		return
	}

	decoded, decodeErr := base64.RawURLEncoding.DecodeString(stateCookie.Value)
	var savedState oauthState
	if decodeErr != nil || json.Unmarshal(decoded, &savedState) != nil ||
		savedState.State != r.URL.Query().Get("state") {
		http.Error(
			w,
			`{"code":"INVALID_CREDENTIAL","message":"Invalid OAuth state. Please try again."}`,
			http.StatusUnauthorized,
		)
		return
	}
	redirectTo := savedState.RedirectTo

	// Clear state cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state_" + provider,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, `{"code":"INVALID_CREDENTIAL","message":"Missing OAuth code."}`, http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), oauthUserInfoTimeout)
	defer cancel()

	oauthToken, err := cfg.Exchange(ctx, code)
	if err != nil {
		h.log.Error("OAuth code exchange failed", "provider", provider, "error", err)
		http.Error(
			w,
			`{"code":"INVALID_CREDENTIAL","message":"Authentication failed. Please try again."}`,
			http.StatusUnauthorized,
		)
		return
	}

	oauthID, name, email, emailVerified, err := fetchUserInfo(ctx, oauthToken, cfg)
	if err != nil {
		h.log.Error("Failed to fetch user info", "provider", provider, "error", err)
		http.Error(
			w,
			`{"code":"INTERNAL_ERROR","message":"Failed to retrieve user information."}`,
			http.StatusInternalServerError,
		)
		return
	}

	providerEnum, parseErr := domain.ParseOAuthProvider(provider)
	if parseErr != nil {
		h.log.Error("Unknown OAuth provider", "provider", provider)
		http.Error(w, `{"code":"INTERNAL_ERROR","message":"Unknown OAuth provider."}`, http.StatusInternalServerError)
		return
	}

	user, roles, authErr := h.authSvc.FindOrCreateOAuthUser(
		r.Context(),
		providerEnum,
		oauthID,
		name,
		email,
		emailVerified,
	)
	if authErr != nil {
		h.log.Debug("OAuth login blocked", "provider", provider, "error", authErr.Error())
		if redirectTo != "" {
			errCode, _ := oauthErrorCode(authErr)
			http.Redirect(w, r, redirectTo+"?error="+errCode, http.StatusTemporaryRedirect)
			return
		}
		status, body := oauthErrorResponse(authErr)
		w.Header().Set("Content-Type", "application/json")
		if status == http.StatusServiceUnavailable {
			w.Header().Set("Retry-After", "30")
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
		return
	}

	if err := h.tokenWrapper.IssueToken(r.Context(), w, user, provider, roles); err != nil {
		h.log.Error("Failed to issue token", "provider", provider, "error", err)
		http.Error(w, `{"error":"failed to issue token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, _, err := h.tokenWrapper.CreateRefreshTokenString(user.ID.String())
	if err != nil {
		h.log.Error("Failed to create refresh token", "provider", provider, "error", err)
		http.Error(w, `{"error":"failed to create refresh token"}`, http.StatusInternalServerError)
		return
	}

	if redirectTo != "" {
		http.Redirect(w, r, redirectTo+"#refresh_token="+refreshToken, http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "authenticated",
		"refresh_token": refreshToken,
		"user": map[string]interface{}{
			"id":   user.ID.String(),
			"name": user.Name,
		},
	})
}

func (h *OAuthHandler) fetchGoogleUserInfo(
	ctx context.Context,
	token *oauth2.Token,
	cfg *oauth2.Config,
) (oauthID, name, email string, emailVerified bool, err error) {
	client := cfg.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return "", "", "", false, fmt.Errorf("get Google userinfo: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<16))
	if err != nil {
		return "", "", "", false, fmt.Errorf("read Google userinfo response: %w", err)
	}

	var info googleUserInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return "", "", "", false, fmt.Errorf("decode Google userinfo: %w", err)
	}
	if info.Sub == "" {
		return "", "", "", false, fmt.Errorf("Google userinfo missing sub")
	}
	return info.Sub, info.Name, info.Email, info.EmailVerified, nil
}

func (h *OAuthHandler) fetchMicrosoftUserInfo(
	ctx context.Context,
	token *oauth2.Token,
	cfg *oauth2.Config,
) (oauthID, name, email string, emailVerified bool, err error) {
	client := cfg.Client(ctx, token)
	resp, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		return "", "", "", false, fmt.Errorf("get Microsoft userinfo: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<16))
	if err != nil {
		return "", "", "", false, fmt.Errorf("read Microsoft userinfo response: %w", err)
	}

	var info microsoftUserInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return "", "", "", false, fmt.Errorf("decode Microsoft userinfo: %w", err)
	}
	if info.ID == "" {
		return "", "", "", false, fmt.Errorf("Microsoft userinfo missing id")
	}

	email = info.Mail
	if email == "" {
		email = info.UserPrincipalName
	}
	// Microsoft's Graph API only returns verified emails.
	return info.ID, info.DisplayName, email, true, nil
}

func oauthErrorCode(err error) (code string, status int) {
	var linked *core.EmailAlreadyLinkedError
	if errors.As(err, &linked) {
		return "EMAIL_ALREADY_LINKED", http.StatusConflict
	}
	var pending *core.RoleAssignmentPendingError
	if errors.As(err, &pending) {
		return "ROLE_ASSIGNMENT_PENDING", http.StatusForbidden
	}
	var walletRegistered *core.WalletRegisteredError
	if errors.As(err, &walletRegistered) {
		return "WALLET_REGISTERED", http.StatusForbidden
	}
	var suspended *core.AccountSuspendedError
	if errors.As(err, &suspended) {
		return "ACCOUNT_SUSPENDED", http.StatusForbidden
	}
	var svcUnavail *core.ServiceUnavailableError
	if errors.As(err, &svcUnavail) {
		return "SERVICE_UNAVAILABLE", http.StatusServiceUnavailable
	}
	return "INTERNAL_ERROR", http.StatusInternalServerError
}

func oauthErrorResponse(err error) (int, string) {
	jsonBody := func(code, msg string) string {
		b, _ := json.Marshal(map[string]string{"code": code, "message": msg})
		return string(b)
	}

	var linked *core.EmailAlreadyLinkedError
	if errors.As(err, &linked) {
		return http.StatusConflict, jsonBody("EMAIL_ALREADY_LINKED", linked.Message())
	}

	var pending *core.RoleAssignmentPendingError
	if errors.As(err, &pending) {
		return http.StatusForbidden, jsonBody(
			"ROLE_ASSIGNMENT_PENDING",
			"Your account has no role assigned yet. Contact your operator.",
		)
	}

	var walletRegistered *core.WalletRegisteredError
	if errors.As(err, &walletRegistered) {
		return http.StatusForbidden, jsonBody(
			"WALLET_REGISTERED",
			"This wallet is now registered but has no role yet. Contact your operator.",
		)
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
