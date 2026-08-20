package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

// fakeOnboardingService captures the arguments ListMine is called with and returns canned pairs.
type fakeOnboardingService struct {
	gotStatus     *domain.ApprovalStatus
	statusSet     bool
	pairs         []core.OnChainAddressPair
	allPending    []core.PendingUserAddressPairs
	setStatusCall bool
	setStatusErr  error // if set, returned by SetApprovalStatus
	gotUserID     uuid.UUID
	gotApproval   domain.ApprovalStatus
}

func (f *fakeOnboardingService) AddAddressPair(_ context.Context, _ uuid.UUID) (*core.OnChainAddressPair, error) {
	return nil, nil
}

func (f *fakeOnboardingService) ListMine(
	_ context.Context,
	_ uuid.UUID,
	status *domain.ApprovalStatus,
) ([]core.OnChainAddressPair, error) {
	f.gotStatus = status
	f.statusSet = true
	return f.pairs, nil
}

func (f *fakeOnboardingService) ListAllPending(_ context.Context) ([]core.PendingUserAddressPairs, error) {
	return f.allPending, nil
}

func (f *fakeOnboardingService) SetApprovalStatus(
	_ context.Context,
	userID uuid.UUID,
	_, _ string,
	status domain.ApprovalStatus,
) error {
	f.setStatusCall = true
	f.gotUserID = userID
	f.gotApproval = status
	return f.setStatusErr
}

func newListContext(t *testing.T, rawQuery, authUserID string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/me/address-pairs?"+rawQuery, nil)
	if authUserID != "" {
		c.Set("auth_user", &token.User{ID: authUserID})
	}
	return c, w
}

func TestOnboardingHandler_ListMine_NoFilterPassesNilStatus(t *testing.T) {
	// With no status query, the handler calls the service with a nil status filter and returns the pairs.
	svc := &fakeOnboardingService{pairs: []core.OnChainAddressPair{
		{PublicChainAddress: "0xpub", PrivateChainAddress: "0xpriv", Status: domain.ApprovalStatusPending},
	}}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newListContext(t, "", uuid.New().String())
	h.ListMine(c)

	require.Equal(t, http.StatusOK, w.Code)
	require.True(t, svc.statusSet)
	assert.Nil(t, svc.gotStatus)
	var resp []addressPairResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	require.Len(t, resp, 1)
	assert.Equal(t, uint8(domain.ApprovalStatusPending), resp[0].Status)
}

func TestOnboardingHandler_ListMine_PendingFilterPassesPendingStatus(t *testing.T) {
	// status=0 passes a pending filter to the service.
	svc := &fakeOnboardingService{}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newListContext(t, "status=0", uuid.New().String())
	h.ListMine(c)

	require.Equal(t, http.StatusOK, w.Code)
	require.NotNil(t, svc.gotStatus)
	assert.Equal(t, domain.ApprovalStatusPending, *svc.gotStatus)
}

func TestOnboardingHandler_ListMine_MissingAuthReturns401(t *testing.T) {
	// Without an authenticated user in context the handler returns 401 and never calls the service.
	svc := &fakeOnboardingService{}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newListContext(t, "", "")
	h.ListMine(c)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, svc.statusSet)
}

func TestOnboardingHandler_ListAllPending_ReturnsGroupsKeyedByUserID(t *testing.T) {
	// The handler returns each group's resolved user_id alongside its address pairs.
	userID := uuid.New()
	svc := &fakeOnboardingService{allPending: []core.PendingUserAddressPairs{
		{
			UserID: userID,
			AddressPairs: []core.OnChainAddressPair{
				{PublicChainAddress: "0xpub", Status: domain.ApprovalStatusPending},
			},
		},
	}}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/admin/address-pairs/pending", nil)
	h.ListAllPending(c)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []pendingUserAddressPairsResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	require.Len(t, resp, 1)
	assert.Equal(t, userID.String(), resp[0].UserID)
	require.Len(t, resp[0].AddressPairs, 1)
	assert.Equal(t, uint8(domain.ApprovalStatusPending), resp[0].AddressPairs[0].Status)
}

func newPatchContext(t *testing.T, id, body string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(
		http.MethodPatch,
		"/api/v1/admin/users/"+id+"/address-pairs/status",
		strings.NewReader(body),
	)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: id}}
	return c, w
}

func TestOnboardingHandler_SetApprovalStatus_Approve(t *testing.T) {
	// A valid approve request returns 200 and forwards the user id and status to the service.
	userID := uuid.New()
	svc := &fakeOnboardingService{}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newPatchContext(
		t,
		userID.String(),
		`{"public_address":"0x1111111111111111111111111111111111111111","private_address":"0x2222222222222222222222222222222222222222","status":1}`,
	)
	h.SetApprovalStatus(c)

	require.Equal(t, http.StatusOK, w.Code)
	require.True(t, svc.setStatusCall)
	assert.Equal(t, userID, svc.gotUserID)
	assert.Equal(t, domain.ApprovalStatusApproved, svc.gotApproval)
}

func TestOnboardingHandler_SetApprovalStatus_InvalidAddressRejected(t *testing.T) {
	// A non-hex address is rejected with 400 and never reaches the service.
	svc := &fakeOnboardingService{}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newPatchContext(
		t,
		uuid.New().String(),
		`{"public_address":"0xnothex","private_address":"0x2222222222222222222222222222222222222222","status":1}`,
	)
	h.SetApprovalStatus(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCall)
}

func TestOnboardingHandler_SetApprovalStatus_UnknownUserReturns404(t *testing.T) {
	// When the service reports the user does not exist, the handler returns 404.
	svc := &fakeOnboardingService{setStatusErr: core.ErrRecordNotFound}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newPatchContext(
		t,
		uuid.New().String(),
		`{"public_address":"0x1111111111111111111111111111111111111111","private_address":"0x2222222222222222222222222222222222222222","status":1}`,
	)
	h.SetApprovalStatus(c)

	require.Equal(t, http.StatusNotFound, w.Code)
	assert.True(t, svc.setStatusCall)
}

func TestOnboardingHandler_SetApprovalStatus_OnChainRevertReturns422(t *testing.T) {
	// An on-chain revert from the service maps to 422, not a generic 500.
	svc := &fakeOnboardingService{setStatusErr: fmt.Errorf("%w (tx 0xabc): Pair not pending", core.ErrTxReverted)}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newPatchContext(
		t,
		uuid.New().String(),
		`{"public_address":"0x1111111111111111111111111111111111111111","private_address":"0x2222222222222222222222222222222222222222","status":1}`,
	)
	h.SetApprovalStatus(c)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.True(t, svc.setStatusCall)
}

func TestOnboardingHandler_SetApprovalStatus_InvalidStatusRejected(t *testing.T) {
	// A status above rejected (2) is rejected with 400 and never reaches the service.
	svc := &fakeOnboardingService{}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newPatchContext(t, uuid.New().String(), `{"public_address":"0xpub","private_address":"0xpriv","status":3}`)
	h.SetApprovalStatus(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCall)
}

func TestOnboardingHandler_SetApprovalStatus_RevertToPending(t *testing.T) {
	// A status of 0 (pending) is accepted and forwarded to the service, reverting a decided pair.
	userID := uuid.New()
	svc := &fakeOnboardingService{}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newPatchContext(
		t,
		userID.String(),
		`{"public_address":"0x1111111111111111111111111111111111111111","private_address":"0x2222222222222222222222222222222222222222","status":0}`,
	)
	h.SetApprovalStatus(c)

	require.Equal(t, http.StatusOK, w.Code)
	require.True(t, svc.setStatusCall)
	assert.Equal(t, userID, svc.gotUserID)
	assert.Equal(t, domain.ApprovalStatusPending, svc.gotApproval)
}

func TestOnboardingHandler_SetApprovalStatus_InvalidIDRejected(t *testing.T) {
	// A non-UUID path id is rejected with 400 before binding the body.
	svc := &fakeOnboardingService{}
	h := NewOnboardingHandler(svc, &testutil.StubLogger{})

	c, w := newPatchContext(t, "not-a-uuid", `{"public_address":"0xpub","private_address":"0xpriv","status":1}`)
	h.SetApprovalStatus(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCall)
}
