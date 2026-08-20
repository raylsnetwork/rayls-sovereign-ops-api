package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-pkgz/auth/v2/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

type fakeTokenActionService struct {
	txHash      string
	mintErr     error
	mintCalled  bool
	gotStandard domain.ErcStandard
	gotMint     core.MintInput

	pauser        string
	pauserErr     error
	isPaused      bool
	isPausedErr   error
	setPausedErr  error
	setPausedCall bool
	gotPaused     bool
}

func (f *fakeTokenActionService) Mint(
	_ context.Context,
	_, _ string,
	std domain.ErcStandard,
	in core.MintInput,
) (string, error) {
	f.mintCalled = true
	f.gotStandard = std
	f.gotMint = in
	return f.txHash, f.mintErr
}

func (f *fakeTokenActionService) Burn(
	_ context.Context,
	_, _ string,
	_ domain.ErcStandard,
	_ core.BurnInput,
) (string, error) {
	return f.txHash, nil
}

func (f *fakeTokenActionService) SetPaused(_ context.Context, _, _ string, paused bool) (string, error) {
	f.setPausedCall = true
	f.gotPaused = paused
	return f.txHash, f.setPausedErr
}

func (f *fakeTokenActionService) Pauser(_ context.Context, _ string) (string, error) {
	return f.pauser, f.pauserErr
}

func (f *fakeTokenActionService) IsPaused(_ context.Context, _ string) (bool, error) {
	return f.isPaused, f.isPausedErr
}

type fakeTeleportService struct {
	txHash         string
	err            error
	teleportCalled bool
	gotStandard    domain.ErcStandard
	gotTeleport    core.TeleportInput
}

func (f *fakeTeleportService) Teleport(
	_ context.Context,
	_ string,
	std domain.ErcStandard,
	in core.TeleportInput,
) (string, error) {
	f.teleportCalled = true
	f.gotStandard = std
	f.gotTeleport = in
	return f.txHash, f.err
}

func newActionContext(t *testing.T, address, body, authUserID string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/tokens/"+address+"/mint", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "address", Value: address}}
	if authUserID != "" {
		c.Set("auth_user", &token.User{ID: authUserID})
	}
	return c, w
}

func erc20Repo(addr string) *fakeTokenRepo {
	return &fakeTokenRepo{byAddress: map[string]*domain.Token{
		addr: {ContractAddress: addr, ErcStandard: domain.ErcStandardERC20, Decimals: 18},
	}}
}

func TestTokenActionHandler_Mint_SuccessScalesAmount(t *testing.T) {
	// A permitted mint scales the human amount by the token decimals and returns the tx hash.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0xWALLET")
	svc := &fakeTokenActionService{txHash: "0xtx"}
	perms := &fakeTokenPermissionService{res: &core.TokenPermissions{CanMint: true}}
	h := NewTokenActionHandler(
		svc,
		&fakeTeleportService{},
		perms,
		erc20Repo("0xtoken"),
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"to":"0x0000000000000000000000000000000000000001","amount":"1.5"}`,
		userID.String(),
	)
	h.Mint(c)

	require.Equal(t, http.StatusOK, w.Code)
	var resp txResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "0xtx", resp.TxHash)
	assert.True(t, svc.mintCalled)
	// 1.5 * 10^18
	expected, _ := new(big.Int).SetString("1500000000000000000", 10)
	assert.Equal(t, expected, svc.gotMint.Amount)
}

func TestTokenActionHandler_Mint_NotPermittedReturns403(t *testing.T) {
	// Without the mint permission the request is rejected before any signing.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0xWALLET")
	svc := &fakeTokenActionService{txHash: "0xtx"}
	perms := &fakeTokenPermissionService{res: &core.TokenPermissions{CanMint: false}}
	h := NewTokenActionHandler(
		svc,
		&fakeTeleportService{},
		perms,
		erc20Repo("0xtoken"),
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"to":"0x0000000000000000000000000000000000000001","amount":"1"}`,
		userID.String(),
	)
	h.Mint(c)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.False(t, svc.mintCalled)
}

func TestTokenActionHandler_Mint_TokenNotFoundReturns404(t *testing.T) {
	// Minting an unknown token yields 404.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0xWALLET")
	perms := &fakeTokenPermissionService{res: &core.TokenPermissions{CanMint: true}}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		&fakeTeleportService{},
		perms,
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xunknown",
		`{"to":"0x0000000000000000000000000000000000000001","amount":"1"}`,
		userID.String(),
	)
	h.Mint(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestTokenActionHandler_Teleport_SuccessERC20(t *testing.T) {
	// A valid ERC20 teleport signs with the caller's wallet and passes raw base units to the service.
	userID := uuid.New()
	signer := "0x1111111111111111111111111111111111111111"
	walletRepo := walletFor(userID, signer)
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x1111111111111111111111111111111111111111","to":"0x0000000000000000000000000000000000000002","amount":"1000","standard":1}`,
		userID.String(),
	)
	h.Teleport(c)

	require.Equal(t, http.StatusOK, w.Code)
	var resp teleportResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "0xtx", resp.TxHash)
	assert.True(t, tele.teleportCalled)
	assert.Equal(t, signer, tele.gotTeleport.From)
	assert.Equal(t, big.NewInt(1000), tele.gotTeleport.Amount)
}

func TestTokenActionHandler_Teleport_UnsupportedStandardReturns400(t *testing.T) {
	// Enygma (4) and other non ERC20/721/1155 standards are rejected before any signing.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0x1111111111111111111111111111111111111111")
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"to":"0x0000000000000000000000000000000000000002","amount":"1000","standard":4}`,
		userID.String(),
	)
	h.Teleport(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, tele.teleportCalled)
}

func TestTokenActionHandler_Teleport_FromMismatchReturns400(t *testing.T) {
	// A from that is not the caller's own custody wallet is rejected.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0x1111111111111111111111111111111111111111")
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x2222222222222222222222222222222222222222","to":"0x0000000000000000000000000000000000000003","amount":"1","standard":1}`,
		userID.String(),
	)
	h.Teleport(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, tele.teleportCalled)
}

func TestTokenActionHandler_Teleport_MissingFromReturns400(t *testing.T) {
	// from is required: it selects which wallet signs, so an empty from is rejected before signing.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0x1111111111111111111111111111111111111111")
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"to":"0x0000000000000000000000000000000000000002","amount":"1000","standard":1}`,
		userID.String(),
	)
	h.Teleport(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, tele.teleportCalled)
}

func TestTokenActionHandler_Teleport_PublicChainWalletRejected(t *testing.T) {
	// A wallet the user owns but tagged for the public chain cannot sign a teleport (source is private).
	userID := uuid.New()
	signer := "0x1111111111111111111111111111111111111111"
	walletRepo := testutil.FakeUserWalletRepository{Wallets: []domain.UserWallet{{
		UserID:          userID,
		RaylsAddress:    signer,
		CustodyProvider: domain.CustodyProviderRaylsHSM,
		Chain:           domain.WalletChainPublic,
		IsActive:        true,
	}}}
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x1111111111111111111111111111111111111111","to":"0x0000000000000000000000000000000000000002","amount":"1000","standard":1}`,
		userID.String(),
	)
	h.Teleport(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, tele.teleportCalled)
}

func TestTokenActionHandler_Teleport_MissingAmountReturnsLegacyMessage(t *testing.T) {
	// An ERC20 teleport without amount is rejected with the legacy backend's message text.
	userID := uuid.New()
	signer := "0x1111111111111111111111111111111111111111"
	walletRepo := walletFor(userID, signer)
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x1111111111111111111111111111111111111111","to":"0x0000000000000000000000000000000000000002","standard":1}`,
		userID.String(),
	)
	h.Teleport(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "amount is required for standard 1 (RAYLS_ERC20)", resp["error"])
	assert.False(t, tele.teleportCalled)
}

func TestTokenActionHandler_Teleport_MissingTokenIdReturnsLegacyMessage(t *testing.T) {
	// An ERC721 teleport without tokenId is rejected with the legacy backend's message text.
	userID := uuid.New()
	signer := "0x1111111111111111111111111111111111111111"
	walletRepo := walletFor(userID, signer)
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x1111111111111111111111111111111111111111","to":"0x0000000000000000000000000000000000000002","standard":2}`,
		userID.String(),
	)
	h.Teleport(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "tokenId is required for standard 2 (RAYLS_ERC721)", resp["error"])
	assert.False(t, tele.teleportCalled)
}

func TestTokenActionHandler_Teleport_InvalidStandardReturnsLegacyMessage(t *testing.T) {
	// An out-of-range standard is rejected with the legacy backend's message text.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0x1111111111111111111111111111111111111111")
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"to":"0x0000000000000000000000000000000000000002","amount":"1000","standard":99}`,
		userID.String(),
	)
	h.Teleport(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Contains(t, resp["error"], "invalid token standard (expected")
	assert.False(t, tele.teleportCalled)
}

func TestTokenActionHandler_Mint_OnChainRevertReturns422(t *testing.T) {
	// A permitted mint whose on-chain tx reverts maps to 422, not a generic gateway error.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0xWALLET")
	svc := &fakeTokenActionService{mintErr: fmt.Errorf("%w (tx 0xabc)", core.ErrTxReverted)}
	perms := &fakeTokenPermissionService{res: &core.TokenPermissions{CanMint: true}}
	h := NewTokenActionHandler(
		svc,
		&fakeTeleportService{},
		perms,
		erc20Repo("0xtoken"),
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"to":"0x0000000000000000000000000000000000000001","amount":"1"}`,
		userID.String(),
	)
	h.Mint(c)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.True(t, svc.mintCalled)
}

func TestTokenActionHandler_Mint_ServiceErrorReturnsBadGateway(t *testing.T) {
	// A non-revert on-chain failure (e.g. RPC error) maps to 502.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0xWALLET")
	svc := &fakeTokenActionService{mintErr: errors.New("rpc dial failed")}
	perms := &fakeTokenPermissionService{res: &core.TokenPermissions{CanMint: true}}
	h := NewTokenActionHandler(
		svc,
		&fakeTeleportService{},
		perms,
		erc20Repo("0xtoken"),
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"to":"0x0000000000000000000000000000000000000001","amount":"1"}`,
		userID.String(),
	)
	h.Mint(c)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestTokenActionHandler_Teleport_OnChainRevertReturns422(t *testing.T) {
	// A teleport whose on-chain tx reverts maps to 422 rather than a generic 500.
	userID := uuid.New()
	signer := "0x1111111111111111111111111111111111111111"
	walletRepo := walletFor(userID, signer)
	tele := &fakeTeleportService{err: fmt.Errorf("%w (tx 0xabc)", core.ErrTxReverted)}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x1111111111111111111111111111111111111111","to":"0x0000000000000000000000000000000000000002","amount":"1000","standard":1}`,
		userID.String(),
	)
	h.Teleport(c)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.True(t, tele.teleportCalled)
}

func TestTokenActionHandler_Teleport_DisabledReturns501(t *testing.T) {
	// With teleport disabled (no teleport service wired), the endpoint returns 501.
	userID := uuid.New()
	signer := "0x1111111111111111111111111111111111111111"
	walletRepo := walletFor(userID, signer)
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		nil,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x1111111111111111111111111111111111111111","to":"0x0000000000000000000000000000000000000002","amount":"1000","standard":1}`,
		userID.String(),
	)
	h.Teleport(c)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestTokenActionHandler_Teleport_ERC1155InvalidDataReturns400(t *testing.T) {
	// An ERC1155 teleport with non-hex data is rejected at the handler before any signing.
	userID := uuid.New()
	signer := "0x1111111111111111111111111111111111111111"
	walletRepo := walletFor(userID, signer)
	tele := &fakeTeleportService{txHash: "0xtx"}
	h := NewTokenActionHandler(
		&fakeTokenActionService{},
		tele,
		&fakeTokenPermissionService{},
		&fakeTokenRepo{},
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(
		t,
		"0xtoken",
		`{"from":"0x1111111111111111111111111111111111111111","to":"0x0000000000000000000000000000000000000002","tokenId":"7","amount":"5","data":"nothex","standard":3}`,
		userID.String(),
	)
	h.Teleport(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, tele.teleportCalled)
}

func TestScaleDecimal(t *testing.T) {
	// Human decimals scale to base units with integer math; invalid inputs error.
	v, err := scaleDecimal("1.5", 18)
	require.NoError(t, err)
	expected, _ := new(big.Int).SetString("1500000000000000000", 10)
	assert.Equal(t, expected, v)

	v, err = scaleDecimal("2", 0)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(2), v)

	_, err = scaleDecimal("1.123", 2)
	assert.Error(t, err) // more decimals than allowed

	_, err = scaleDecimal("-1", 18)
	assert.Error(t, err)
}

func stablecoinRepo(addr string) *fakeTokenRepo {
	return &fakeTokenRepo{byAddress: map[string]*domain.Token{
		addr: {ContractAddress: addr, ErcStandard: domain.ErcStandardStableCoin, Decimals: 6},
	}}
}

func newPauseHandler(t *testing.T, svc *fakeTokenActionService, userID uuid.UUID) *TokenActionHandler {
	t.Helper()
	walletRepo := walletFor(userID, "0xWALLET")
	return NewTokenActionHandler(
		svc,
		&fakeTeleportService{},
		// Empty permissions on purpose: pause must NOT consult the Access Manager.
		&fakeTokenPermissionService{res: &core.TokenPermissions{}},
		stablecoinRepo("0xtoken"),
		&walletRepo,
		&testutil.StubLogger{},
	)
}

func TestTokenActionHandler_Pause_SucceedsForThePauser(t *testing.T) {
	// The contract's pauser can pause a running stablecoin, with no AM permission involved.
	userID := uuid.New()
	svc := &fakeTokenActionService{txHash: "0xtx", pauser: "0xWALLET", isPaused: false}
	h := newPauseHandler(t, svc, userID)

	c, w := newActionContext(t, "0xtoken", `{"paused":true}`, userID.String())
	h.Pause(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.setPausedCall)
	assert.True(t, svc.gotPaused)
}

func TestTokenActionHandler_Pause_NonPauserReturns403(t *testing.T) {
	// onlyPauser is a msg.sender check — a wallet that is not the pauser is refused before signing.
	userID := uuid.New()
	svc := &fakeTokenActionService{txHash: "0xtx", pauser: "0xSOMEONEELSE"}
	h := newPauseHandler(t, svc, userID)

	c, w := newActionContext(t, "0xtoken", `{"paused":true}`, userID.String())
	h.Pause(c)

	require.Equal(t, http.StatusForbidden, w.Code)
	assert.False(t, svc.setPausedCall)
}

func TestTokenActionHandler_Pause_AlreadyPausedReturns409(t *testing.T) {
	// Pausing an already-paused token would revert on-chain; refuse it before spending gas.
	userID := uuid.New()
	svc := &fakeTokenActionService{txHash: "0xtx", pauser: "0xWALLET", isPaused: true}
	h := newPauseHandler(t, svc, userID)

	c, w := newActionContext(t, "0xtoken", `{"paused":true}`, userID.String())
	h.Pause(c)

	require.Equal(t, http.StatusConflict, w.Code)
	assert.False(t, svc.setPausedCall)
}

func TestTokenActionHandler_Pause_UnpauseSendsFalse(t *testing.T) {
	// Resuming a paused token calls unpause().
	userID := uuid.New()
	svc := &fakeTokenActionService{txHash: "0xtx", pauser: "0xWALLET", isPaused: true}
	h := newPauseHandler(t, svc, userID)

	c, w := newActionContext(t, "0xtoken", `{"paused":false}`, userID.String())
	h.Pause(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, svc.setPausedCall)
	assert.False(t, svc.gotPaused)
}

func TestTokenActionHandler_Pause_MissingPausedFieldReturns400(t *testing.T) {
	// An omitted `paused` must not be read as false — that would silently unpause.
	userID := uuid.New()
	svc := &fakeTokenActionService{txHash: "0xtx", pauser: "0xWALLET"}
	h := newPauseHandler(t, svc, userID)

	c, w := newActionContext(t, "0xtoken", `{}`, userID.String())
	h.Pause(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setPausedCall)
}

func TestTokenActionHandler_Pause_NonStablecoinReturns400(t *testing.T) {
	// Only the stablecoin has pause()/unpause(); other standards would revert with no selector.
	userID := uuid.New()
	walletRepo := walletFor(userID, "0xWALLET")
	svc := &fakeTokenActionService{txHash: "0xtx", pauser: "0xWALLET"}
	h := NewTokenActionHandler(
		svc,
		&fakeTeleportService{},
		&fakeTokenPermissionService{res: &core.TokenPermissions{}},
		erc20Repo("0xtoken"),
		&walletRepo,
		&testutil.StubLogger{},
	)

	c, w := newActionContext(t, "0xtoken", `{"paused":true}`, userID.String())
	h.Pause(c)

	require.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, svc.setPausedCall)
}
