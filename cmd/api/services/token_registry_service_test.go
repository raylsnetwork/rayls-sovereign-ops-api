package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/services/testutil"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

// fakeTokenRegistry models the on-chain TokenRegistry adapter, recording write calls and serving a
// canned read-back entry from GetByAddress.
type fakeTokenRegistry struct {
	registerCalls    int
	lastOperator     string
	lastRegisterIn   core.RegisterTokenInput
	registerErr      error
	byAddress        *core.RegisteredToken
	getByAddressErr  error
	tokens           []core.RegisteredToken // served by List/ListByStatus
	setStatusCalls   int
	lastStatusToken  string
	lastStatus       domain.PrivacyNodeStatus
	setStatusErr     error
	freezeCalls      int
	unfreezeCalls    int
	lastFreezeToken  string
	lastFreezeLayer  domain.FreezeLayer
	freezeErr        error
	unfreezeErr      error
	submitCalls      int
	lastSubmitToken  string
	lastSubmitTarget domain.SubmitTarget
	submitErr        error
}

func (f *fakeTokenRegistry) Register(_ context.Context, operator string, in core.RegisterTokenInput) (string, error) {
	f.registerCalls++
	f.lastOperator = operator
	f.lastRegisterIn = in
	return "0xhash", f.registerErr
}

func (f *fakeTokenRegistry) SetStatus(
	_ context.Context,
	operator, tokenAddress string,
	status domain.PrivacyNodeStatus,
) (string, error) {
	f.setStatusCalls++
	f.lastOperator = operator
	f.lastStatusToken = tokenAddress
	f.lastStatus = status
	return "0xhash", f.setStatusErr
}

func (f *fakeTokenRegistry) Freeze(
	_ context.Context,
	operator, tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	f.freezeCalls++
	f.lastOperator = operator
	f.lastFreezeToken = tokenAddress
	f.lastFreezeLayer = layer
	return "0xhash", f.freezeErr
}

func (f *fakeTokenRegistry) Unfreeze(
	_ context.Context,
	operator, tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	f.unfreezeCalls++
	f.lastOperator = operator
	f.lastFreezeToken = tokenAddress
	f.lastFreezeLayer = layer
	return "0xhash", f.unfreezeErr
}

func (f *fakeTokenRegistry) Submit(
	_ context.Context,
	operator, tokenAddress string,
	target domain.SubmitTarget,
) (string, error) {
	f.submitCalls++
	f.lastOperator = operator
	f.lastSubmitToken = tokenAddress
	f.lastSubmitTarget = target
	return "0xhash", f.submitErr
}

func (f *fakeTokenRegistry) List(_ context.Context) ([]core.RegisteredToken, error) {
	return f.tokens, nil
}

func (f *fakeTokenRegistry) ListByStatus(
	_ context.Context,
	_ domain.PrivacyNodeStatus,
) ([]core.RegisteredToken, error) {
	return f.tokens, nil
}

func (f *fakeTokenRegistry) GetByAddress(_ context.Context, _ string) (*core.RegisteredToken, error) {
	return f.byAddress, f.getByAddressErr
}

func (f *fakeTokenRegistry) GetBySymbol(_ context.Context, _ string) (*core.RegisteredToken, error) {
	return f.byAddress, f.getByAddressErr
}

func (f *fakeTokenRegistry) Exists(_ context.Context, _ string) (bool, error) { return false, nil }

func TestTokenRegistryService_Register_ResolvesOperatorAndReadsBack(t *testing.T) {
	// Register resolves the operator, writes via the adapter, then returns the read-back entry.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{byAddress: &core.RegisteredToken{
		Symbol: "MYT", TokenAddress: "0xtoken", Status: domain.PrivacyNodeStatusWaitingApproval,
	}}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	in := core.RegisterTokenInput{TokenAddress: "0xtoken"}
	token, err := svc.Register(context.Background(), in)

	require.NoError(t, err)
	assert.Equal(t, 1, registry.registerCalls)
	assert.Equal(t, "0xOPERATOR", registry.lastOperator)
	assert.Equal(t, in, registry.lastRegisterIn)
	assert.Equal(t, "MYT", token.Symbol)
	assert.Equal(t, domain.PrivacyNodeStatusWaitingApproval, token.Status)
}

func TestTokenRegistryService_Register_ResolverFailurePropagates(t *testing.T) {
	// When the operator cannot be resolved, the adapter is never called and the error propagates.
	resolver := &fakeOperatorResolver{err: &core.NoOperatorSignerError{}}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Register(context.Background(), core.RegisterTokenInput{TokenAddress: "0xtoken"})

	require.Error(t, err)
	assert.True(t, errors.As(err, new(*core.NoOperatorSignerError)))
	assert.Equal(t, 0, registry.registerCalls)
}

func TestTokenRegistryService_SetStatus_ResolvesOperator(t *testing.T) {
	// SetStatus resolves the operator and forwards the status write to the adapter.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.SetStatus(context.Background(), "0xtoken", domain.PrivacyNodeStatusAuthorized)

	require.NoError(t, err)
	assert.Equal(t, 1, registry.setStatusCalls)
	assert.Equal(t, "0xOPERATOR", registry.lastOperator)
	assert.Equal(t, "0xtoken", registry.lastStatusToken)
	assert.Equal(t, domain.PrivacyNodeStatusAuthorized, registry.lastStatus)
}

func TestTokenRegistryService_Register_AdapterErrorPropagates(t *testing.T) {
	// When the on-chain register write fails, the error propagates and no read-back is attempted.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{registerErr: errors.New("reverted: already registered")}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Register(context.Background(), core.RegisterTokenInput{TokenAddress: "0xtoken"})

	require.Error(t, err)
	assert.Equal(t, 1, registry.registerCalls)
}

func TestTokenRegistryService_Register_ReadBackErrorPropagates(t *testing.T) {
	// A failure reading the token back after a successful register propagates to the caller.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{getByAddressErr: errors.New("not found")}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Register(context.Background(), core.RegisterTokenInput{TokenAddress: "0xtoken"})

	require.Error(t, err)
	assert.Equal(t, 1, registry.registerCalls)
}

func TestTokenRegistryService_Register_ReadBackErrorCarriesTheTxHash(t *testing.T) {
	// The registration is already committed on-chain, so the error must name the tx and say the
	// write succeeded — otherwise the caller retries and gets a misleading 422 from the contract.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{getByAddressErr: errors.New("rpc hiccup")}
	log := &testutil.RecordingLogger{}
	svc := NewTokenRegistryService(resolver, registry, log)

	_, err := svc.Register(context.Background(), core.RegisterTokenInput{TokenAddress: "0xtoken"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "0xhash", "the tx hash is how the caller checks the receipt")
	assert.Contains(t, err.Error(), "registered on-chain")
	assert.Contains(t, err.Error(), "rather than retrying")
	assert.NotEmpty(t, log.Errors, "an on-chain write with a failed read-back must be logged")
}

func TestTokenRegistryService_List_WarnsOnUnknownStatus(t *testing.T) {
	// A status this build does not know means the contract is ahead of the API; it is logged
	// rather than silently rendered as the "unknown" label.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{tokens: []core.RegisteredToken{
		{TokenAddress: "0xtoken", Status: domain.PrivacyNodeStatus(9)},
	}}
	log := &testutil.RecordingLogger{}
	svc := NewTokenRegistryService(resolver, registry, log)

	tokens, err := svc.List(context.Background())

	require.NoError(t, err)
	assert.Len(t, tokens, 1, "an unknown status must not drop the token from the response")
	assert.NotEmpty(t, log.Warns)
}

func TestTokenRegistryService_List_QuietForKnownStatuses(t *testing.T) {
	// Every status in the enum is recognized, so a normal listing logs nothing.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{tokens: []core.RegisteredToken{
		{TokenAddress: "0xa", Status: domain.PrivacyNodeStatusWaitingApproval},
		{TokenAddress: "0xb", Status: domain.PrivacyNodeStatusAuthorized},
		{TokenAddress: "0xc", Status: domain.PrivacyNodeStatusUnauthorized},
		{TokenAddress: "0xd", Status: domain.PrivacyNodeStatusFrozen},
	}}
	log := &testutil.RecordingLogger{}
	svc := NewTokenRegistryService(resolver, registry, log)

	_, err := svc.List(context.Background())

	require.NoError(t, err)
	assert.Empty(t, log.Warns)
}

func TestTokenRegistryService_GetByAddress_WarnsOnUnknownStatus(t *testing.T) {
	// The single-token read surfaces the same drift as the list endpoints.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{
		byAddress: &core.RegisteredToken{TokenAddress: "0xtoken", Status: domain.PrivacyNodeStatus(9)},
	}
	log := &testutil.RecordingLogger{}
	svc := NewTokenRegistryService(resolver, registry, log)

	token, err := svc.GetByAddress(context.Background(), "0xtoken")

	require.NoError(t, err)
	require.NotNil(t, token)
	assert.NotEmpty(t, log.Warns)
}

func TestTokenRegistryService_SetStatus_ResolverFailurePropagates(t *testing.T) {
	// When the operator cannot be resolved, the status write is never attempted.
	resolver := &fakeOperatorResolver{err: &core.NoOperatorSignerError{}}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.SetStatus(context.Background(), "0xtoken", domain.PrivacyNodeStatusAuthorized)

	require.Error(t, err)
	assert.True(t, errors.As(err, new(*core.NoOperatorSignerError)))
	assert.Equal(t, 0, registry.setStatusCalls)
}

func TestTokenRegistryService_SetStatus_AdapterErrorPropagates(t *testing.T) {
	// A failure in the on-chain status write propagates to the caller.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{setStatusErr: errors.New("reverted: not permitted")}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.SetStatus(context.Background(), "0xtoken", domain.PrivacyNodeStatusAuthorized)

	require.Error(t, err)
	assert.Equal(t, 1, registry.setStatusCalls)
}

func TestTokenRegistryService_Freeze_ResolvesOperatorAndForwardsLayer(t *testing.T) {
	// Freeze resolves the operator and forwards the token address + layer to the adapter.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Freeze(context.Background(), "0xtoken", domain.FreezeLayerPublicChain)

	require.NoError(t, err)
	assert.Equal(t, 1, registry.freezeCalls)
	assert.Equal(t, "0xOPERATOR", registry.lastOperator)
	assert.Equal(t, "0xtoken", registry.lastFreezeToken)
	assert.Equal(t, domain.FreezeLayerPublicChain, registry.lastFreezeLayer)
}

func TestTokenRegistryService_Freeze_ResolverFailurePropagates(t *testing.T) {
	// When the operator cannot be resolved, the freeze write is never attempted.
	resolver := &fakeOperatorResolver{err: &core.NoOperatorSignerError{}}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Freeze(context.Background(), "0xtoken", domain.FreezeLayerPrivacyNode)

	require.Error(t, err)
	assert.True(t, errors.As(err, new(*core.NoOperatorSignerError)))
	assert.Equal(t, 0, registry.freezeCalls)
}

func TestTokenRegistryService_Freeze_AdapterErrorPropagates(t *testing.T) {
	// A failure in the on-chain freeze write propagates to the caller.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{freezeErr: errors.New("reverted: not permitted")}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Freeze(context.Background(), "0xtoken", domain.FreezeLayerPrivacyNode)

	require.Error(t, err)
	assert.Equal(t, 1, registry.freezeCalls)
}

func TestTokenRegistryService_Unfreeze_ResolvesOperatorAndForwardsLayer(t *testing.T) {
	// Unfreeze resolves the operator and forwards the token address + layer to the adapter.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Unfreeze(context.Background(), "0xtoken", domain.FreezeLayerPrivacyNode)

	require.NoError(t, err)
	assert.Equal(t, 1, registry.unfreezeCalls)
	assert.Equal(t, "0xOPERATOR", registry.lastOperator)
	assert.Equal(t, "0xtoken", registry.lastFreezeToken)
	assert.Equal(t, domain.FreezeLayerPrivacyNode, registry.lastFreezeLayer)
}

func TestTokenRegistryService_Submit_ResolvesOperatorAndForwardsTarget(t *testing.T) {
	// Submit resolves the operator and forwards the token address + target to the adapter.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Submit(context.Background(), "0xtoken", domain.SubmitTargetHub)

	require.NoError(t, err)
	assert.Equal(t, 1, registry.submitCalls)
	assert.Equal(t, "0xOPERATOR", registry.lastOperator)
	assert.Equal(t, "0xtoken", registry.lastSubmitToken)
	assert.Equal(t, domain.SubmitTargetHub, registry.lastSubmitTarget)
}

func TestTokenRegistryService_Submit_ResolverFailurePropagates(t *testing.T) {
	// When the operator cannot be resolved, the submit write is never attempted.
	resolver := &fakeOperatorResolver{err: &core.NoOperatorSignerError{}}
	registry := &fakeTokenRegistry{}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Submit(context.Background(), "0xtoken", domain.SubmitTargetPublicChain)

	require.Error(t, err)
	assert.True(t, errors.As(err, new(*core.NoOperatorSignerError)))
	assert.Equal(t, 0, registry.submitCalls)
}

func TestTokenRegistryService_Submit_AdapterErrorPropagates(t *testing.T) {
	// A failure in the on-chain submit write propagates to the caller.
	resolver := &fakeOperatorResolver{addr: "0xOPERATOR"}
	registry := &fakeTokenRegistry{submitErr: errors.New("reverted: not authorized")}
	svc := NewTokenRegistryService(resolver, registry, &testutil.StubLogger{})

	_, err := svc.Submit(context.Background(), "0xtoken", domain.SubmitTargetHub)

	require.Error(t, err)
	assert.Equal(t, 1, registry.submitCalls)
}
