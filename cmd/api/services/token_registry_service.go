package services

import (
	"context"
	"fmt"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

var _ core.TokenRegistryService = (*tokenRegistryService)(nil)

// tokenRegistryService orchestrates the token-registry capability. Reads pass straight through to
// the adapter; writes are operator-authority signed by the resolved operator wallet.
type tokenRegistryService struct {
	operatorResolver core.OperatorSignerResolver
	registry         core.TokenRegistryAdapter
	log              logger.Logger
}

func NewTokenRegistryService(
	operatorResolver core.OperatorSignerResolver,
	registry core.TokenRegistryAdapter,
	log logger.Logger,
) core.TokenRegistryService {
	return &tokenRegistryService{
		operatorResolver: operatorResolver,
		registry:         registry,
		log:              log,
	}
}

// Register resolves the operator wallet, registers the token on-chain, then reads the entry back via
// GetByAddress and returns it. Newly registered tokens start in WAITING_APPROVAL status.
func (s *tokenRegistryService) Register(
	ctx context.Context,
	in core.RegisterTokenInput,
) (*core.RegisteredToken, error) {
	operatorAddress, err := s.operatorResolver.Resolve(ctx)
	if err != nil {
		return nil, fmt.Errorf("resolve operator signer: %w", err)
	}

	txHash, err := s.registry.Register(ctx, operatorAddress, in)
	if err != nil {
		return nil, fmt.Errorf("register token: %w", err)
	}

	// Past this point the registration is COMMITTED on-chain (the tx mined with status 1). A
	// failed read-back is therefore a reporting failure, not a registration failure — so the
	// error carries the tx hash, and says so. Without it the caller sees a 500, retries, and
	// gets a 422 from the contract refusing to re-register: a misleading signal for what is
	// actually a success. The retry is not the fix; re-reading the token is.
	token, err := s.registry.GetByAddress(ctx, in.TokenAddress)
	if err != nil {
		s.log.Error("token registered on-chain but the read-back failed — do NOT retry the registration",
			"address", in.TokenAddress, "txHash", txHash, "error", err)
		return nil, fmt.Errorf(
			"token %s was registered on-chain (tx %s) but reading it back failed; "+
				"the registration succeeded — re-read the token rather than retrying: %w",
			in.TokenAddress, txHash, err)
	}

	s.log.Info("token registered", "address", in.TokenAddress, "status", token.Status.Label())
	return token, nil
}

// SetStatus resolves the operator wallet and submits the operator-signed status update.
func (s *tokenRegistryService) SetStatus(
	ctx context.Context,
	tokenAddress string,
	status domain.PrivacyNodeStatus,
) (string, error) {
	operatorAddress, err := s.operatorResolver.Resolve(ctx)
	if err != nil {
		return "", fmt.Errorf("resolve operator signer: %w", err)
	}

	txHash, err := s.registry.SetStatus(ctx, operatorAddress, tokenAddress, status)
	if err != nil {
		return "", fmt.Errorf("set token status: %w", err)
	}

	s.log.Info("token status set", "address", tokenAddress, "status", status.Label())
	return txHash, nil
}

// Freeze resolves the operator wallet and submits the operator-signed freeze at the given layer.
func (s *tokenRegistryService) Freeze(
	ctx context.Context,
	tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	operatorAddress, err := s.operatorResolver.Resolve(ctx)
	if err != nil {
		return "", fmt.Errorf("resolve operator signer: %w", err)
	}

	txHash, err := s.registry.Freeze(ctx, operatorAddress, tokenAddress, layer)
	if err != nil {
		return "", fmt.Errorf("freeze token: %w", err)
	}

	s.log.Info("token frozen", "address", tokenAddress, "layer", string(layer))
	return txHash, nil
}

// Unfreeze resolves the operator wallet and submits the operator-signed unfreeze at the given layer.
func (s *tokenRegistryService) Unfreeze(
	ctx context.Context,
	tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	operatorAddress, err := s.operatorResolver.Resolve(ctx)
	if err != nil {
		return "", fmt.Errorf("resolve operator signer: %w", err)
	}

	txHash, err := s.registry.Unfreeze(ctx, operatorAddress, tokenAddress, layer)
	if err != nil {
		return "", fmt.Errorf("unfreeze token: %w", err)
	}

	s.log.Info("token unfrozen", "address", tokenAddress, "layer", string(layer))
	return txHash, nil
}

// Submit resolves the operator wallet and submits the operator-signed submitToHub /
// submitToPublicChain for the given target. Submitting only initiates the flow; activation on the Hub
// and deployment on the Public Chain complete later via cross-chain PNH / relayer callbacks.
func (s *tokenRegistryService) Submit(
	ctx context.Context,
	tokenAddress string,
	target domain.SubmitTarget,
) (string, error) {
	operatorAddress, err := s.operatorResolver.Resolve(ctx)
	if err != nil {
		return "", fmt.Errorf("resolve operator signer: %w", err)
	}

	txHash, err := s.registry.Submit(ctx, operatorAddress, tokenAddress, target)
	if err != nil {
		return "", fmt.Errorf("submit token: %w", err)
	}

	s.log.Info("token submitted", "address", tokenAddress, "target", string(target))
	return txHash, nil
}

// warnOnUnknownStatus logs any status this build does not recognize. The TokenRegistry contract is
// upgradeable, so a status added on-chain after this API was built would otherwise surface only as
// the "unknown" label in a response — no log, no error, and nothing to point at during a debug
// session. Read-only: the response still carries the raw value, so new statuses do not break reads.
func (s *tokenRegistryService) warnOnUnknownStatus(tokens []core.RegisteredToken) {
	for i := range tokens {
		if !tokens[i].Status.IsKnown() {
			s.log.Warn("unknown PrivacyNodeStatus from the TokenRegistry — the contract is likely ahead of this build",
				"value", uint8(tokens[i].Status), "address", tokens[i].TokenAddress)
		}
	}
}

func (s *tokenRegistryService) List(ctx context.Context) ([]core.RegisteredToken, error) {
	tokens, err := s.registry.List(ctx)
	if err != nil {
		return nil, err
	}
	s.warnOnUnknownStatus(tokens)
	return tokens, nil
}

func (s *tokenRegistryService) ListByStatus(
	ctx context.Context,
	status domain.PrivacyNodeStatus,
) ([]core.RegisteredToken, error) {
	tokens, err := s.registry.ListByStatus(ctx, status)
	if err != nil {
		return nil, err
	}
	s.warnOnUnknownStatus(tokens)
	return tokens, nil
}

func (s *tokenRegistryService) GetByAddress(ctx context.Context, tokenAddress string) (*core.RegisteredToken, error) {
	token, err := s.registry.GetByAddress(ctx, tokenAddress)
	if err != nil {
		return nil, err
	}
	s.warnOnUnknownStatus([]core.RegisteredToken{*token})
	return token, nil
}

func (s *tokenRegistryService) GetBySymbol(ctx context.Context, symbol string) (*core.RegisteredToken, error) {
	token, err := s.registry.GetBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}
	s.warnOnUnknownStatus([]core.RegisteredToken{*token})
	return token, nil
}

func (s *tokenRegistryService) Exists(ctx context.Context, tokenAddress string) (bool, error) {
	return s.registry.Exists(ctx, tokenAddress)
}
