package auth_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/auth"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// newSignupHandler builds an enabled email-signup handler backed by the given details repo,
// returning the user the fake auth service resolves every signup to.
func newSignupHandler(
	t *testing.T,
	detailsRepo *testutil.FakeUserSignupDetailsRepository,
) (*auth.OAuthHandler, *domain.User) {
	t.Helper()
	user := buildUser()
	authSvc := &testutil.FakeAuthService{User: user}
	wrapper := auth.NewIdentityTokenWrapper(
		newTokenService(t), testJWTSecret, &testutil.FakeUserWalletRepository{}, "",
	)

	h := auth.NewOAuthHandler(
		authSvc, wrapper, nil, nil, "", "", true, detailsRepo, &testutil.StubLogger{},
	)
	return h, user
}

const fullSignupBody = `{
	"email":"alice@example.com",
	"name":"Alice",
	"company":"Acme Bank",
	"employees":"51-200",
	"heardAbout":"Conference",
	"goals":"Issue a tokenised deposit."
}`

func TestOAuthHandler_EmailSignup_DisabledByDefault(t *testing.T) {
	// Signup issues a session from an unverified email address, so it must stay off
	// unless explicitly enabled — otherwise anyone can log in as the admin.
	h := auth.NewOAuthHandler(nil, nil, nil, nil, "", "", false, nil, &testutil.StubLogger{})

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/signup",
		bytes.NewBufferString(`{"email":"attacker@example.com"}`))

	h.EmailSignup(w, r)

	require.Equal(t, http.StatusNotFound, w.Code)
	assert.NotContains(t, w.Header().Get("Set-Cookie"), "JWT")
}

func TestOAuthHandler_EmailSignup_StoresSignupDetails(t *testing.T) {
	// Every field the sign-up form collects is persisted against the authenticated user.
	repo := &testutil.FakeUserSignupDetailsRepository{}
	h, user := newSignupHandler(t, repo)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBufferString(fullSignupBody))

	h.EmailSignup(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	require.Len(t, repo.Details, 1)
	assert.Equal(t, user.ID, repo.Details[0].UserID)
	assert.Equal(t, "Acme Bank", repo.Details[0].Company)
	assert.Equal(t, "51-200", repo.Details[0].Employees)
	assert.Equal(t, "Conference", repo.Details[0].HeardAbout)
	assert.Equal(t, "Issue a tokenised deposit.", repo.Details[0].Goals)
}

func TestOAuthHandler_EmailSignup_SkipsStoreWhenNoDetailsSubmitted(t *testing.T) {
	// An email-only signup must not write a blank row that would wipe details already on file.
	repo := &testutil.FakeUserSignupDetailsRepository{}
	h, _ := newSignupHandler(t, repo)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/signup",
		bytes.NewBufferString(`{"email":"alice@example.com"}`))

	h.EmailSignup(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Empty(t, repo.Details)
}

func TestOAuthHandler_EmailSignup_IssuesSessionWhenDetailsStoreFails(t *testing.T) {
	// The answers are lead data, not credentials: losing them must not cost the caller
	// the session they just authenticated for.
	repo := &testutil.FakeUserSignupDetailsRepository{UpsertErr: errors.New("db down")}
	h, _ := newSignupHandler(t, repo)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBufferString(fullSignupBody))

	h.EmailSignup(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Set-Cookie"), "JWT")
}

func TestOAuthHandler_EmailSignup_ReSignupOverwritesDetails(t *testing.T) {
	// One row per user: filling the form again updates it rather than adding a second row.
	repo := &testutil.FakeUserSignupDetailsRepository{}
	h, user := newSignupHandler(t, repo)

	first := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBufferString(fullSignupBody))
	h.EmailSignup(httptest.NewRecorder(), first)

	second := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBufferString(
		`{"email":"alice@example.com","company":"Globex","employees":"1-50","heardAbout":"Referral","goals":"Settle privately."}`,
	))
	h.EmailSignup(httptest.NewRecorder(), second)

	require.Len(t, repo.Details, 1)
	assert.Equal(t, user.ID, repo.Details[0].UserID)
	assert.Equal(t, "Globex", repo.Details[0].Company)
}
