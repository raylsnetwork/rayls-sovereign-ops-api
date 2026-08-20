package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

type fakeTokenRegistryService struct {
	registerCalled bool
	gotInput       core.RegisterTokenInput
	result         *core.RegisteredToken
	registerErr    error

	listCalled bool
	listResult []core.RegisteredToken
	listErr    error

	listByStatusCalled bool
	gotStatus          domain.PrivacyNodeStatus
	listByStatusResult []core.RegisteredToken
	listByStatusErr    error

	setStatusCalled bool
	gotStatusAddr   string
	gotSetStatus    domain.PrivacyNodeStatus
	setStatusErr    error

	freezeCalled   bool
	unfreezeCalled bool
	gotFreezeAddr  string
	gotFreezeLayer domain.FreezeLayer
	freezeErr      error
	unfreezeErr    error

	submitCalled    bool
	gotSubmitAddr   string
	gotSubmitTarget domain.SubmitTarget
	submitErr       error
}

func (f *fakeTokenRegistryService) Register(
	_ context.Context,
	in core.RegisterTokenInput,
) (*core.RegisteredToken, error) {
	f.registerCalled = true
	f.gotInput = in
	return f.result, f.registerErr
}

func (f *fakeTokenRegistryService) SetStatus(
	_ context.Context,
	tokenAddress string,
	status domain.PrivacyNodeStatus,
) (string, error) {
	f.setStatusCalled = true
	f.gotStatusAddr = tokenAddress
	f.gotSetStatus = status
	return "0xhash", f.setStatusErr
}

func (f *fakeTokenRegistryService) Freeze(
	_ context.Context,
	tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	f.freezeCalled = true
	f.gotFreezeAddr = tokenAddress
	f.gotFreezeLayer = layer
	return "0xhash", f.freezeErr
}

func (f *fakeTokenRegistryService) Unfreeze(
	_ context.Context,
	tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	f.unfreezeCalled = true
	f.gotFreezeAddr = tokenAddress
	f.gotFreezeLayer = layer
	return "0xhash", f.unfreezeErr
}

func (f *fakeTokenRegistryService) Submit(
	_ context.Context,
	tokenAddress string,
	target domain.SubmitTarget,
) (string, error) {
	f.submitCalled = true
	f.gotSubmitAddr = tokenAddress
	f.gotSubmitTarget = target
	return "0xhash", f.submitErr
}

func (f *fakeTokenRegistryService) List(_ context.Context) ([]core.RegisteredToken, error) {
	f.listCalled = true
	return f.listResult, f.listErr
}

func (f *fakeTokenRegistryService) ListByStatus(
	_ context.Context,
	status domain.PrivacyNodeStatus,
) ([]core.RegisteredToken, error) {
	f.listByStatusCalled = true
	f.gotStatus = status
	return f.listByStatusResult, f.listByStatusErr
}

func (f *fakeTokenRegistryService) GetByAddress(_ context.Context, _ string) (*core.RegisteredToken, error) {
	return nil, nil
}

func (f *fakeTokenRegistryService) GetBySymbol(_ context.Context, _ string) (*core.RegisteredToken, error) {
	return nil, nil
}

func (f *fakeTokenRegistryService) Exists(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func newRegisterContext(t *testing.T, address, authUserID string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/tokens/"+address+"/register", nil)
	c.Params = gin.Params{{Key: "address", Value: address}}
	if authUserID != "" {
		c.Set("auth_user", &token.User{ID: authUserID})
	}
	return c, w
}

const validTokenAddr = "0x1111111111111111111111111111111111111111"

func TestTokenRegistryHandler_Register_Success(t *testing.T) {
	// A valid address-only registration returns 201 with the read-back token (status WAITING_APPROVAL).
	svc := &fakeTokenRegistryService{result: &core.RegisteredToken{
		TokenAddress: validTokenAddr, Name: "My Token", Symbol: "MYT",
		Standard: domain.ErcStandardERC20, Status: domain.PrivacyNodeStatusWaitingApproval,
	}}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegisterContext(t, validTokenAddr, "user-1")
	h.Register(c)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp registeredTokenResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "MYT", resp.Symbol)
	assert.Equal(t, uint8(domain.PrivacyNodeStatusWaitingApproval), resp.Status)
	assert.True(t, svc.registerCalled)
	assert.Equal(t, domain.NormalizeAddress(validTokenAddr), svc.gotInput.TokenAddress)
}

func TestTokenRegistryHandler_Register_InvalidAddressReturns400(t *testing.T) {
	// A non-hex path address is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegisterContext(t, "not-an-address", "user-1")
	h.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.registerCalled)
}

func TestTokenRegistryHandler_Register_NoOperatorReturns503(t *testing.T) {
	// When no operator signer is available, the handler surfaces 503.
	svc := &fakeTokenRegistryService{registerErr: &core.NoOperatorSignerError{}}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegisterContext(t, validTokenAddr, "user-1")
	h.Register(c)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestTokenRegistryHandler_Register_OnChainRevertReturns422(t *testing.T) {
	// An on-chain revert (e.g. token/symbol already registered) maps to 422, not a generic 500.
	svc := &fakeTokenRegistryService{registerErr: fmt.Errorf("%w (tx 0xabc)", core.ErrTxReverted)}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegisterContext(t, validTokenAddr, "user-1")
	h.Register(c)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestTokenRegistryHandler_Register_ServiceErrorReturns500(t *testing.T) {
	// A non-revert service failure (e.g. RPC dial) maps to a generic 500.
	svc := &fakeTokenRegistryService{registerErr: errors.New("rpc dial failed")}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegisterContext(t, validTokenAddr, "user-1")
	h.Register(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestTokenRegistryHandler_Register_UnauthenticatedReturns401(t *testing.T) {
	// Without an authenticated user the request is rejected.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegisterContext(t, validTokenAddr, "")
	h.Register(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, svc.registerCalled)
}

func newRegistryGetContext(t *testing.T, path, authUserID string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, path, nil)
	if authUserID != "" {
		c.Set("auth_user", &token.User{ID: authUserID})
	}
	return c, w
}

func TestTokenRegistryHandler_List_ReturnsAllTokens(t *testing.T) {
	// List returns every catalog entry via TokenRegistryService.List.
	svc := &fakeTokenRegistryService{listResult: []core.RegisteredToken{
		{
			TokenAddress: validTokenAddr,
			Symbol:       "MYT",
			Standard:     domain.ErcStandardERC20,
			Status:       domain.PrivacyNodeStatusAuthorized,
		},
		{
			TokenAddress: "0x2222222222222222222222222222222222222222",
			Symbol:       "PEND",
			Standard:     domain.ErcStandardERC721,
			Status:       domain.PrivacyNodeStatusWaitingApproval,
		},
	}}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegistryGetContext(t, "/api/tokens/registry", "user-1")
	h.List(c)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []registeredTokenResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, svc.listCalled)
	assert.False(t, svc.listByStatusCalled)
	require.Len(t, resp, 2)
	assert.Equal(t, "MYT", resp[0].Symbol)
	assert.Equal(t, uint8(domain.PrivacyNodeStatusAuthorized), resp[0].Status)
}

func TestTokenRegistryHandler_List_EmptySerializesAsArray(t *testing.T) {
	// An empty catalog serializes as [] rather than null.
	svc := &fakeTokenRegistryService{listResult: nil}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegistryGetContext(t, "/api/tokens/registry", "user-1")
	h.List(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", strings.TrimSpace(w.Body.String()))
}

func TestTokenRegistryHandler_List_UnauthenticatedReturns401(t *testing.T) {
	// Without an authenticated user the list request is rejected.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegistryGetContext(t, "/api/tokens/registry", "")
	h.List(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, svc.listCalled)
}

func TestTokenRegistryHandler_List_ServiceErrorReturns500(t *testing.T) {
	// A read failure is mapped to a 500 via HandleError.
	svc := &fakeTokenRegistryService{listErr: errors.New("rpc dial failed")}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegistryGetContext(t, "/api/tokens/registry", "user-1")
	h.List(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestTokenRegistryHandler_ListPending_FiltersByWaitingApproval(t *testing.T) {
	// ListPending returns only pending tokens via ListByStatus(WAITING_APPROVAL).
	svc := &fakeTokenRegistryService{listByStatusResult: []core.RegisteredToken{
		{
			TokenAddress: validTokenAddr,
			Symbol:       "PEND",
			Standard:     domain.ErcStandardERC20,
			Status:       domain.PrivacyNodeStatusWaitingApproval,
		},
	}}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegistryGetContext(t, "/api/tokens/registry/pending", "user-1")
	h.ListPending(c)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []registeredTokenResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, svc.listByStatusCalled)
	assert.False(t, svc.listCalled)
	assert.Equal(t, domain.PrivacyNodeStatusWaitingApproval, svc.gotStatus)
	require.Len(t, resp, 1)
	assert.Equal(t, uint8(domain.PrivacyNodeStatusWaitingApproval), resp[0].Status)
}

func TestTokenRegistryHandler_ListPending_UnauthenticatedReturns401(t *testing.T) {
	// Without an authenticated user the pending list request is rejected.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newRegistryGetContext(t, "/api/tokens/registry/pending", "")
	h.ListPending(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, svc.listByStatusCalled)
}

func newSetStatusContext(t *testing.T, address, body string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(
		http.MethodPatch,
		"/api/v1/admin/tokens/"+address+"/status",
		strings.NewReader(body),
	)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "address", Value: address}}
	return c, w
}

func TestTokenRegistryHandler_SetStatus_Authorized(t *testing.T) {
	// "authorized" returns 200 and forwards the normalized address and status to the service.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"authorized"}`)
	h.SetStatus(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.setStatusCalled)
	assert.Equal(t, domain.NormalizeAddress(validTokenAddr), svc.gotStatusAddr)
	assert.Equal(t, domain.PrivacyNodeStatusAuthorized, svc.gotSetStatus)
}

func TestTokenRegistryHandler_SetStatus_Unauthorized(t *testing.T) {
	// "unauthorized" is accepted and forwarded.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"unauthorized"}`)
	h.SetStatus(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.setStatusCalled)
	assert.Equal(t, domain.PrivacyNodeStatusUnauthorized, svc.gotSetStatus)
}

func TestTokenRegistryHandler_SetStatus_RejectsWaitingApproval(t *testing.T) {
	// "waiting_approval" is the initial state set by registration and is rejected here.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"waiting_approval"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCalled)
}

func TestTokenRegistryHandler_SetStatus_RejectsFrozen(t *testing.T) {
	// "frozen" has dedicated freeze contract methods and is rejected on this endpoint.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"frozen"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCalled)
}

func TestTokenRegistryHandler_SetStatus_RejectsUndefined(t *testing.T) {
	// "undefined" is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"undefined"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCalled)
}

func TestTokenRegistryHandler_SetStatus_RejectsOutOfRange(t *testing.T) {
	// An unrecognized status label is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"bogus_status"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCalled)
}

func TestTokenRegistryHandler_SetStatus_InvalidAddressReturns400(t *testing.T) {
	// A non-hex path address is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, "not-an-address", `{"status":"authorized"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setStatusCalled)
}

func TestTokenRegistryHandler_SetStatus_OnChainRevertReturns422(t *testing.T) {
	// An on-chain revert (e.g. the token is not registered) maps to 422, not a generic 500.
	svc := &fakeTokenRegistryService{setStatusErr: fmt.Errorf("%w (tx 0xabc)", core.ErrTxReverted)}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"authorized"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestTokenRegistryHandler_SetStatus_ServiceErrorReturns500(t *testing.T) {
	// A non-revert service failure on the status write maps to a generic 500.
	svc := &fakeTokenRegistryService{setStatusErr: errors.New("rpc dial failed")}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"authorized"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestTokenRegistryHandler_SetStatus_NoOperatorReturns503(t *testing.T) {
	// When no operator signer is available, the handler surfaces 503.
	svc := &fakeTokenRegistryService{setStatusErr: &core.NoOperatorSignerError{}}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newSetStatusContext(t, validTokenAddr, `{"status":"authorized"}`)
	h.SetStatus(c)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func newFreezeContext(t *testing.T, address, action, body string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(
		http.MethodPost,
		"/api/v1/admin/tokens/"+address+"/"+action,
		strings.NewReader(body),
	)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "address", Value: address}}
	return c, w
}

func TestTokenRegistryHandler_Freeze_PrivacyNode(t *testing.T) {
	// A privacy_node freeze returns 200 and forwards the normalized address and layer to the service.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "freeze", `{"layer":"privacy_node"}`)
	h.Freeze(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.freezeCalled)
	assert.False(t, svc.unfreezeCalled)
	assert.Equal(t, domain.NormalizeAddress(validTokenAddr), svc.gotFreezeAddr)
	assert.Equal(t, domain.FreezeLayerPrivacyNode, svc.gotFreezeLayer)
}

func TestTokenRegistryHandler_Freeze_PublicChain(t *testing.T) {
	// A public_chain freeze is accepted and forwarded with the public_chain layer.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "freeze", `{"layer":"public_chain"}`)
	h.Freeze(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.freezeCalled)
	assert.Equal(t, domain.FreezeLayerPublicChain, svc.gotFreezeLayer)
}

func TestTokenRegistryHandler_Freeze_RejectsHubLayer(t *testing.T) {
	// The hub (PNH) layer is not supported on this endpoint and is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "freeze", `{"layer":"hub"}`)
	h.Freeze(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.freezeCalled)
}

func TestTokenRegistryHandler_Freeze_RejectsUnknownLayer(t *testing.T) {
	// An unrecognized layer is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "freeze", `{"layer":"nonsense"}`)
	h.Freeze(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.freezeCalled)
}

func TestTokenRegistryHandler_Freeze_InvalidAddressReturns400(t *testing.T) {
	// A non-hex path address is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, "not-an-address", "freeze", `{"layer":"privacy_node"}`)
	h.Freeze(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.freezeCalled)
}

func TestTokenRegistryHandler_Freeze_OnChainRevertReturns422(t *testing.T) {
	// An on-chain revert (e.g. token not registered or operator not authorized) maps to 422.
	svc := &fakeTokenRegistryService{freezeErr: fmt.Errorf("%w (tx 0xabc)", core.ErrTxReverted)}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "freeze", `{"layer":"privacy_node"}`)
	h.Freeze(c)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestTokenRegistryHandler_Freeze_NoOperatorReturns503(t *testing.T) {
	// When no operator signer is available, the handler surfaces 503.
	svc := &fakeTokenRegistryService{freezeErr: &core.NoOperatorSignerError{}}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "freeze", `{"layer":"privacy_node"}`)
	h.Freeze(c)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestTokenRegistryHandler_Unfreeze_PrivacyNode(t *testing.T) {
	// An unfreeze returns 200 and forwards the normalized address and layer to the service.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "unfreeze", `{"layer":"privacy_node"}`)
	h.Unfreeze(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.unfreezeCalled)
	assert.False(t, svc.freezeCalled)
	assert.Equal(t, domain.NormalizeAddress(validTokenAddr), svc.gotFreezeAddr)
	assert.Equal(t, domain.FreezeLayerPrivacyNode, svc.gotFreezeLayer)
}

func TestTokenRegistryHandler_Unfreeze_RejectsUnknownLayer(t *testing.T) {
	// An unrecognized layer is rejected on unfreeze before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "unfreeze", `{"layer":""}`)
	h.Unfreeze(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.unfreezeCalled)
}

func TestTokenRegistryHandler_Submit_Hub(t *testing.T) {
	// A hub submit returns 200 and forwards the normalized address and target to the service.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "submit", `{"target":"hub"}`)
	h.Submit(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.submitCalled)
	assert.Equal(t, domain.NormalizeAddress(validTokenAddr), svc.gotSubmitAddr)
	assert.Equal(t, domain.SubmitTargetHub, svc.gotSubmitTarget)
}

func TestTokenRegistryHandler_Submit_PublicChain(t *testing.T) {
	// A public_chain submit is accepted and forwarded with the public_chain target.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "submit", `{"target":"public_chain"}`)
	h.Submit(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.submitCalled)
	assert.Equal(t, domain.SubmitTargetPublicChain, svc.gotSubmitTarget)
}

func TestTokenRegistryHandler_Submit_RejectsUnknownTarget(t *testing.T) {
	// An unrecognized target is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "submit", `{"target":"nonsense"}`)
	h.Submit(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.submitCalled)
}

func TestTokenRegistryHandler_Submit_InvalidAddressReturns400(t *testing.T) {
	// A non-hex path address is rejected before any service call.
	svc := &fakeTokenRegistryService{}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, "not-an-address", "submit", `{"target":"hub"}`)
	h.Submit(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.submitCalled)
}

func TestTokenRegistryHandler_Submit_OnChainRevertReturns422(t *testing.T) {
	// An on-chain revert (e.g. token not yet AUTHORIZED on the PN) maps to 422.
	svc := &fakeTokenRegistryService{submitErr: fmt.Errorf("%w (tx 0xabc)", core.ErrTxReverted)}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "submit", `{"target":"hub"}`)
	h.Submit(c)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestTokenRegistryHandler_Submit_NoOperatorReturns503(t *testing.T) {
	// When no operator signer is available, the handler surfaces 503.
	svc := &fakeTokenRegistryService{submitErr: &core.NoOperatorSignerError{}}
	h := NewTokenRegistryHandler(svc, &testutil.StubLogger{})

	c, w := newFreezeContext(t, validTokenAddr, "submit", `{"target":"public_chain"}`)
	h.Submit(c)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}
