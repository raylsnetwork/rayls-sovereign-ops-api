package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// emailSignupRequest is the body of POST /auth/signup. Only email is required;
// name is optional (defaults to the local part of the email).
//
// The remaining fields are the sign-up form's profile answers. They are optional here
// even though the form marks them required — the endpoint's job is to authenticate, and
// a missing answer must not cost the caller a session. They are persisted to
// user_signup_details; see EmailSignup.
type emailSignupRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`

	Company    string `json:"company"`
	Employees  string `json:"employees"`
	HeardAbout string `json:"heardAbout"`
	Goals      string `json:"goals"`
}

// EmailSignup handles email self sign-up. It reuses the same find-or-create +
// login-decision-tree + IssueToken machinery as the OAuth callback, treating
// email as a provider whose oauthID is the email address itself.
//
// DISABLED BY DEFAULT (config.Auth.EmailSignupEnabled). There is no verification code
// step yet, so this endpoint issues a valid session to whoever submits an address — it
// proves nothing about who is calling, and an attacker can simply submit the admin's
// address to become the admin. That is tolerable on a local machine and nowhere else,
// which matters even more now the identity service is shared by every chain: one forged
// session would be honoured across all of them.
//
// TODO: require a verified signup code, then drop the flag and enable this by default.
func (h *OAuthHandler) EmailSignup(w http.ResponseWriter, r *http.Request) {
	if !h.emailSignupEnabled {
		http.Error(w, `{"code":"NOT_FOUND","message":"Email signup is not enabled."}`, http.StatusNotFound)
		return
	}

	var req emailSignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"code":"VALIDATION_ERROR","message":"Invalid request body."}`, http.StatusBadRequest)
		return
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))
	if email == "" {
		http.Error(w, `{"code":"VALIDATION_ERROR","message":"email is required."}`, http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		name = email
		if at := strings.IndexByte(email, '@'); at > 0 {
			name = email[:at]
		}
	}

	// oauthID = email: a stable per-user key for the email provider. emailVerified=true
	// is the stand-in for the skipped code step (see TODO above).
	user, roles, authErr := h.authSvc.FindOrCreateOAuthUser(
		r.Context(),
		domain.OAuthProviderEmail,
		email,
		name,
		email,
		true,
	)
	if authErr != nil {
		h.log.Debug("Email signup blocked", "email", email, "error", authErr.Error())
		status, body := oauthErrorResponse(authErr)
		w.Header().Set("Content-Type", "application/json")
		if status == http.StatusServiceUnavailable {
			w.Header().Set("Retry-After", "30")
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
		return
	}

	h.storeSignupDetails(r.Context(), user.ID, &req)

	if err := h.tokenWrapper.IssueToken(r.Context(), w, user, "email", roles); err != nil {
		h.log.Error("Failed to issue token", "provider", "email", "error", err)
		http.Error(w, `{"error":"failed to issue token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, _, err := h.tokenWrapper.CreateRefreshTokenString(user.ID.String())
	if err != nil {
		h.log.Error("Failed to create refresh token", "provider", "email", "error", err)
		http.Error(w, `{"error":"failed to create refresh token"}`, http.StatusInternalServerError)
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

// storeSignupDetails persists the sign-up form's profile answers for userID.
//
// Best-effort by design: these answers are lead information, not credentials, so a write
// failure is logged and swallowed rather than costing the caller the session they just
// authenticated for. Skipped entirely when the request carries no answers (an API client
// posting only an email) so a bare signup cannot blank out details already on file.
func (h *OAuthHandler) storeSignupDetails(ctx context.Context, userID uuid.UUID, req *emailSignupRequest) {
	if h.signupDetailsRepo == nil {
		return
	}

	company := strings.TrimSpace(req.Company)
	employees := strings.TrimSpace(req.Employees)
	heardAbout := strings.TrimSpace(req.HeardAbout)
	goals := strings.TrimSpace(req.Goals)
	if company == "" && employees == "" && heardAbout == "" && goals == "" {
		return
	}

	details := &domain.UserSignupDetails{
		UserID:     userID,
		Company:    company,
		Employees:  employees,
		HeardAbout: heardAbout,
		Goals:      goals,
	}
	if err := h.signupDetailsRepo.Upsert(ctx, details); err != nil {
		h.log.Error("Failed to store signup details", "userId", userID.String(), "error", err)
		return
	}
	h.log.Info("Stored signup details", "userId", userID.String(), "company", company)
}
