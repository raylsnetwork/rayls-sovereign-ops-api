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

type fakeTokenDeployService struct {
	addr, hash string
	chainID    string
	err        error
	gotSigner  string
	gotSpec    core.TokenDeploySpec
}

func (f *fakeTokenDeployService) Deploy(
	_ context.Context,
	signer string,
	spec core.TokenDeploySpec,
) (string, string, error) {
	f.gotSigner = signer
	f.gotSpec = spec
	return f.addr, f.hash, f.err
}

func (f *fakeTokenDeployService) EstimateDeploy(
	_ context.Context,
	signer string,
	spec core.TokenDeploySpec,
) (core.TokenDeployEstimate, error) {
	f.gotSigner = signer
	f.gotSpec = spec
	if f.err != nil {
		return core.TokenDeployEstimate{}, f.err
	}
	return core.TokenDeployEstimate{
		GasLimit:    1_000_000,
		GasPriceWei: "100000000000",
		TotalFeeWei: "100000000000000000",
	}, nil
}

func (f *fakeTokenDeployService) ChainID() string { return f.chainID }

type fakeTokenRepo struct {
	upserted  []*domain.Token
	byAddress map[string]*domain.Token
}

func (f *fakeTokenRepo) Upsert(_ context.Context, t *domain.Token) error {
	f.upserted = append(f.upserted, t)
	return nil
}

func (f *fakeTokenRepo) UpdateSupplyAndHolders(_ context.Context, _, _ string, _ int) error {
	return nil
}

func (f *fakeTokenRepo) FindByAddress(_ context.Context, address string) (*domain.Token, error) {
	if t, ok := f.byAddress[address]; ok {
		return t, nil
	}
	return nil, core.ErrRecordNotFound
}

func (f *fakeTokenRepo) List(_ context.Context, _ core.TokenFilter) ([]*domain.Token, int64, error) {
	return nil, 0, nil
}

// fakeTokenRegistry records Register/SetStatus calls; registerErr/setStatusErr force failures.
type fakeTokenRegistry struct {
	registered   []string
	authorized   map[string]domain.PrivacyNodeStatus
	registerErr  error
	setStatusErr error
}

func (f *fakeTokenRegistry) Register(_ context.Context, in core.RegisterTokenInput) (*core.RegisteredToken, error) {
	if f.registerErr != nil {
		return nil, f.registerErr
	}
	f.registered = append(f.registered, in.TokenAddress)
	return &core.RegisteredToken{TokenAddress: in.TokenAddress, Status: domain.PrivacyNodeStatusWaitingApproval}, nil
}

func (f *fakeTokenRegistry) SetStatus(_ context.Context, addr string, status domain.PrivacyNodeStatus) (string, error) {
	if f.setStatusErr != nil {
		return "", f.setStatusErr
	}
	if f.authorized == nil {
		f.authorized = map[string]domain.PrivacyNodeStatus{}
	}
	f.authorized[addr] = status
	return "0xreg", nil
}

func (f *fakeTokenRegistry) Freeze(context.Context, string, domain.FreezeLayer) (string, error)   { return "", nil }
func (f *fakeTokenRegistry) Unfreeze(context.Context, string, domain.FreezeLayer) (string, error) { return "", nil }
func (f *fakeTokenRegistry) Submit(context.Context, string, domain.SubmitTarget) (string, error)  { return "", nil }
func (f *fakeTokenRegistry) List(context.Context) ([]core.RegisteredToken, error)                 { return nil, nil }
func (f *fakeTokenRegistry) ListByStatus(context.Context, domain.PrivacyNodeStatus) ([]core.RegisteredToken, error) {
	return nil, nil
}
func (f *fakeTokenRegistry) GetByAddress(context.Context, string) (*core.RegisteredToken, error) { return nil, nil }
func (f *fakeTokenRegistry) GetBySymbol(context.Context, string) (*core.RegisteredToken, error)  { return nil, nil }
func (f *fakeTokenRegistry) Exists(context.Context, string) (bool, error)                        { return false, nil }

func newDeployTestContext(t *testing.T, body string, authUserID string) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/tokens", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	if authUserID != "" {
		c.Set("auth_user", &token.User{ID: authUserID})
	}
	return c, w
}

func walletFor(userID uuid.UUID, addr string) testutil.FakeUserWalletRepository {
	return testutil.FakeUserWalletRepository{
		Wallets: []domain.UserWallet{{
			UserID:          userID,
			RaylsAddress:    addr,
			CustodyProvider: domain.CustodyProviderRaylsHSM, // signer lookups require an HSM wallet
			Chain:           domain.WalletChainPrivate,
			IsActive:        true,
		}},
	}
}

func TestTokenDeployHandler_Deploy_SuccessReturns201AndRecordsToken(t *testing.T) {
	// A valid request deploys the token, signs with the user's wallet, and persists it.
	userID := uuid.New()
	// The service returns a checksummed (mixed-case) address; the handler must canonicalize it.
	svc := &fakeTokenDeployService{addr: "0xAbCdEf0000000000000000000000000000000001", hash: "0xdef", chainID: "1337"}
	walletRepo := walletFor(userID, "0xWALLET")
	tokenRepo := &fakeTokenRepo{}
	h := NewTokenDeployHandler(svc, &walletRepo, tokenRepo, &testutil.StubLogger{})

	body := `{"standard":"ERC20","name":"Token","symbol":"TKN","decimals":18}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp deployTokenResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	// Response and persisted row use the canonical lowercase form.
	assert.Equal(t, "0xabcdef0000000000000000000000000000000001", resp.DeployedAddress)
	assert.Equal(t, "0xdef", resp.TxHash)
	assert.Equal(t, "0xWALLET", svc.gotSigner)
	assert.Equal(t, domain.ErcStandardERC20, svc.gotSpec.ErcStandard)
	require.Len(t, tokenRepo.upserted, 1)
	assert.Equal(t, "0xabcdef0000000000000000000000000000000001", tokenRepo.upserted[0].ContractAddress)
	assert.Equal(t, domain.TokenStatusInternal, tokenRepo.upserted[0].Status)
	assert.Equal(t, "1337", tokenRepo.upserted[0].IssuerID)
}

func TestTokenDeployHandler_Deploy_RegistersAndAuthorizesToken(t *testing.T) {
	// With a registry wired, a successful deploy registers the token and authorizes it.
	userID := uuid.New()
	svc := &fakeTokenDeployService{addr: "0xAbCdEf0000000000000000000000000000000003", hash: "0xhash", chainID: "1337"}
	walletRepo := walletFor(userID, "0xWALLET")
	registry := &fakeTokenRegistry{}
	h := NewTokenDeployHandler(svc, &walletRepo, &fakeTokenRepo{}, &testutil.StubLogger{})
	h.SetTokenRegistry(registry)

	body := `{"standard":"STABLECOIN","name":"USD Coin","symbol":"USDC","decimals":6}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	require.Equal(t, http.StatusCreated, w.Code)
	const addr = "0xabcdef0000000000000000000000000000000003"
	assert.Equal(t, []string{addr}, registry.registered)
	assert.Equal(t, domain.PrivacyNodeStatusAuthorized, registry.authorized[addr])
}

func TestTokenDeployHandler_Deploy_RegistrationFailureReturns502(t *testing.T) {
	// A registration failure after a successful deploy surfaces as 502.
	userID := uuid.New()
	svc := &fakeTokenDeployService{addr: "0xAbCdEf0000000000000000000000000000000004", hash: "0xhash", chainID: "1337"}
	walletRepo := walletFor(userID, "0xWALLET")
	tokenRepo := &fakeTokenRepo{}
	registry := &fakeTokenRegistry{registerErr: fmt.Errorf("boom")}
	h := NewTokenDeployHandler(svc, &walletRepo, tokenRepo, &testutil.StubLogger{})
	h.SetTokenRegistry(registry)

	body := `{"standard":"ERC20","name":"Token","symbol":"TKN","decimals":18}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	assert.Equal(t, http.StatusBadGateway, w.Code)
	require.Len(t, tokenRepo.upserted, 1) // still tracked despite the registration failure
}

func TestTokenDeployHandler_Deploy_StableCoinReachesServiceAsStableCoinStandard(t *testing.T) {
	// A STABLECOIN request parses to ErcStandardStableCoin and reaches the deploy service with it.
	userID := uuid.New()
	svc := &fakeTokenDeployService{addr: "0xAbCdEf0000000000000000000000000000000002", hash: "0xfee", chainID: "1337"}
	walletRepo := walletFor(userID, "0xWALLET")
	tokenRepo := &fakeTokenRepo{}
	h := NewTokenDeployHandler(svc, &walletRepo, tokenRepo, &testutil.StubLogger{})

	body := `{"standard":"STABLECOIN","name":"USD Coin","symbol":"USDC","decimals":6}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	require.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, domain.ErcStandardStableCoin, svc.gotSpec.ErcStandard)
	require.Len(t, tokenRepo.upserted, 1)
	assert.Equal(t, domain.ErcStandardStableCoin, tokenRepo.upserted[0].ErcStandard)
}

func TestTokenDeployHandler_ParseErcStandard_StableCoinAliases(t *testing.T) {
	// All documented stablecoin aliases resolve to ErcStandardStableCoin.
	for _, in := range []string{"STABLECOIN", "stablecoin", "STABLE_COIN", "RAYLS_STABLECOIN"} {
		std, ok := parseErcStandard(in)
		require.True(t, ok, "alias %q should parse", in)
		assert.Equal(t, domain.ErcStandardStableCoin, std)
	}
	assert.Contains(t, supportedStandards(), "STABLECOIN")
}

func TestTokenDeployHandler_Deploy_NoAuthReturns401(t *testing.T) {
	// A request without an authenticated user is rejected with 401.
	svc := &fakeTokenDeployService{}
	walletRepo := testutil.FakeUserWalletRepository{}
	h := NewTokenDeployHandler(svc, &walletRepo, &fakeTokenRepo{}, &testutil.StubLogger{})

	c, w := newDeployTestContext(t, `{"standard":"ERC20","name":"T","symbol":"T"}`, "")

	h.Deploy(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestTokenDeployHandler_Deploy_InvalidStandardReturns400(t *testing.T) {
	// An unknown token standard is rejected with 400 before any deploy.
	userID := uuid.New()
	svc := &fakeTokenDeployService{}
	walletRepo := walletFor(userID, "0xWALLET")
	h := NewTokenDeployHandler(svc, &walletRepo, &fakeTokenRepo{}, &testutil.StubLogger{})

	body := `{"standard":"ERC999","name":"T","symbol":"T"}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTokenDeployHandler_Deploy_MissingRequiredFieldReturns400(t *testing.T) {
	// An ERC20 request without a symbol fails validation with 400.
	userID := uuid.New()
	svc := &fakeTokenDeployService{}
	walletRepo := walletFor(userID, "0xWALLET")
	h := NewTokenDeployHandler(svc, &walletRepo, &fakeTokenRepo{}, &testutil.StubLogger{})

	body := `{"standard":"ERC20","name":"Token"}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTokenDeployHandler_Deploy_NoWalletReturns400(t *testing.T) {
	// A user without a custody wallet cannot deploy.
	userID := uuid.New()
	svc := &fakeTokenDeployService{}
	walletRepo := testutil.FakeUserWalletRepository{}
	h := NewTokenDeployHandler(svc, &walletRepo, &fakeTokenRepo{}, &testutil.StubLogger{})

	body := `{"standard":"ERC20","name":"Token","symbol":"TKN"}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTokenDeployHandler_Deploy_RevertReturns422(t *testing.T) {
	// An on-chain revert (e.g. missing deployer role) maps to 422.
	userID := uuid.New()
	svc := &fakeTokenDeployService{err: fmt.Errorf("token deploy %w (tx 0x1)", core.ErrTxReverted)}
	walletRepo := walletFor(userID, "0xWALLET")
	h := NewTokenDeployHandler(svc, &walletRepo, &fakeTokenRepo{}, &testutil.StubLogger{})

	body := `{"standard":"ERC20","name":"Token","symbol":"TKN"}`
	c, w := newDeployTestContext(t, body, userID.String())

	h.Deploy(c)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
