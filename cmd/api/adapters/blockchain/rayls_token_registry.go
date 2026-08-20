package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/DeploymentProxyRegistryV1"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/PNTokenCoreV1"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/PNTokenFreezeManagerV1"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/PNTokenRegistryV1"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

var _ core.TokenRegistryAdapter = (*RaylsTokenRegistryService)(nil)

// RaylsTokenRegistryService manages the on-chain TokenRegistry (PNTokenCoreV1) catalog. All writes are
// operator-authority writes signed via custody by the operator wallet passed in — the operator key
// never leaves the HSM.
type RaylsTokenRegistryService struct {
	client       *ethclient.Client
	registry     *PNTokenRegistryV1.PNTokenRegistryV1
	registryAddr common.Address
	custody      core.CustodyService
	chainID      *big.Int
}

// NewRaylsTokenRegistryService reuses an existing RPC client (e.g. from RaylsAccessManagerService)
// and resolves the TokenRegistry address via DeploymentProxyRegistry. A 30-second timeout covers
// the sequential startup RPC calls (ChainID, getContract).
func NewRaylsTokenRegistryService(
	client *ethclient.Client,
	registryAddr string,
	custody core.CustodyService,
) (*RaylsTokenRegistryService, error) {
	startupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	chainID, err := client.ChainID(startupCtx)
	if err != nil {
		return nil, fmt.Errorf("get chain ID: %w", err)
	}

	regAddr := common.HexToAddress(registryAddr)
	if err := contracts.EnsureCode(startupCtx, client, regAddr); err != nil {
		return nil, fmt.Errorf("instantiate DeploymentProxyRegistry: %w", err)
	}
	proxyRegistry := DeploymentProxyRegistryV1.NewDeploymentProxyRegistryV1()

	regOut, err := client.CallContract(startupCtx, ethereum.CallMsg{
		To:   &regAddr,
		Data: proxyRegistry.PackGetContract("TokenRegistry"),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("resolve TokenRegistry address: %w", err)
	}
	tokenRegistryAddr, err := proxyRegistry.UnpackGetContract(regOut)
	if err != nil {
		return nil, fmt.Errorf("decode TokenRegistry address: %w", err)
	}
	if tokenRegistryAddr == (common.Address{}) {
		return nil, fmt.Errorf("DeploymentProxyRegistry returned zero address for TokenRegistry")
	}

	if err := contracts.EnsureCode(startupCtx, client, tokenRegistryAddr); err != nil {
		return nil, fmt.Errorf("instantiate TokenRegistry: %w", err)
	}

	// Register the facade + delegated module error ABIs so on-chain reverts surface as readable
	// "ErrorName(args…)" instead of a raw selector (reverts from TokenCore / TokenFreezeManager
	// bubble up through the facade).
	for _, m := range []*bind.MetaData{
		&PNTokenRegistryV1.PNTokenRegistryV1MetaData,
		&PNTokenCoreV1.PNTokenCoreV1MetaData,
		&PNTokenFreezeManagerV1.PNTokenFreezeManagerV1MetaData,
	} {
		if err := contracts.RegisterErrorsFromMetaData(m); err != nil {
			return nil, fmt.Errorf("register error ABI: %w", err)
		}
	}

	return &RaylsTokenRegistryService{
		client:       client,
		registry:     PNTokenRegistryV1.NewPNTokenRegistryV1(),
		registryAddr: tokenRegistryAddr,
		custody:      custody,
		chainID:      chainID,
	}, nil
}

// Register adds an already-deployed token contract to the catalog. Registration is address-only —
// the contract reads the token's metadata (name/symbol/standard/supply) on-chain. It rejects
// externally owned accounts and empty addresses by requiring deployed code at the token address
// before encoding registerToken. The token starts in privacyNodeStatus WAITING_APPROVAL.
func (s *RaylsTokenRegistryService) Register(
	ctx context.Context,
	operatorAddress string,
	in core.RegisterTokenInput,
) (string, error) {
	if !common.IsHexAddress(in.TokenAddress) {
		return "", fmt.Errorf("invalid token address: %q", in.TokenAddress)
	}
	tokenAddr := common.HexToAddress(in.TokenAddress)
	if err := contracts.EnsureCode(ctx, s.client, tokenAddr); err != nil {
		return "", fmt.Errorf("validate token contract: %w", err)
	}

	calldata := s.registry.PackRegisterToken(tokenAddr)
	return s.signAndSend(ctx, operatorAddress, calldata)
}

func (s *RaylsTokenRegistryService) SetStatus(
	ctx context.Context,
	operatorAddress, tokenAddress string,
	status domain.PrivacyNodeStatus,
) (string, error) {
	if !common.IsHexAddress(tokenAddress) {
		return "", fmt.Errorf("invalid token address: %q", tokenAddress)
	}
	calldata := s.registry.PackUpdatePrivacyNodeStatus(common.HexToAddress(tokenAddress), uint8(status))
	return s.signAndSend(ctx, operatorAddress, calldata)
}

// Freeze freezes a registered token at the given layer using the facade's dedicated freeze methods
// (freezeOnPrivacyNode / freezeOnPublicChain), operator-signed.
func (s *RaylsTokenRegistryService) Freeze(
	ctx context.Context,
	operatorAddress, tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	return s.setFreeze(ctx, operatorAddress, tokenAddress, layer, true)
}

// Unfreeze unfreezes a registered token at the given layer (unfreezeOnPrivacyNode /
// unfreezeOnPublicChain), operator-signed.
func (s *RaylsTokenRegistryService) Unfreeze(
	ctx context.Context,
	operatorAddress, tokenAddress string,
	layer domain.FreezeLayer,
) (string, error) {
	return s.setFreeze(ctx, operatorAddress, tokenAddress, layer, false)
}

// setFreeze selects the freeze/unfreeze calldata for the requested layer and submits it. The Hub
// (PNH) layer is not supported here — it is a cross-chain PNH callback, not an operator action.
func (s *RaylsTokenRegistryService) setFreeze(
	ctx context.Context,
	operatorAddress, tokenAddress string,
	layer domain.FreezeLayer,
	frozen bool,
) (string, error) {
	if !common.IsHexAddress(tokenAddress) {
		return "", fmt.Errorf("invalid token address: %q", tokenAddress)
	}
	tokenAddr := common.HexToAddress(tokenAddress)

	var calldata []byte
	switch layer {
	case domain.FreezeLayerPrivacyNode:
		if frozen {
			calldata = s.registry.PackFreezeOnPrivacyNode(tokenAddr)
		} else {
			calldata = s.registry.PackUnfreezeOnPrivacyNode(tokenAddr)
		}
	case domain.FreezeLayerPublicChain:
		if frozen {
			calldata = s.registry.PackFreezeOnPublicChain(tokenAddr)
		} else {
			calldata = s.registry.PackUnfreezeOnPublicChain(tokenAddr)
		}
	default:
		return "", fmt.Errorf("unsupported freeze layer: %q", layer)
	}

	return s.signAndSend(ctx, operatorAddress, calldata)
}

// Submit submits an AUTHORIZED token to the given target using the facade's submitToHub /
// submitToPublicChain methods, operator-signed. It only initiates the flow — Hub activation and
// public-chain deployment complete later via cross-chain PNH / relayer callbacks. The contract
// enforces the privacyNodeStatus == AUTHORIZED precondition and reverts otherwise (surfaced as 422).
func (s *RaylsTokenRegistryService) Submit(
	ctx context.Context,
	operatorAddress, tokenAddress string,
	target domain.SubmitTarget,
) (string, error) {
	if !common.IsHexAddress(tokenAddress) {
		return "", fmt.Errorf("invalid token address: %q", tokenAddress)
	}
	tokenAddr := common.HexToAddress(tokenAddress)

	var calldata []byte
	switch target {
	case domain.SubmitTargetHub:
		calldata = s.registry.PackSubmitToHub(tokenAddr)
	case domain.SubmitTargetPublicChain:
		calldata = s.registry.PackSubmitToPublicChain(tokenAddr)
	default:
		return "", fmt.Errorf("unsupported submit target: %q", target)
	}

	return s.signAndSend(ctx, operatorAddress, calldata)
}

func (s *RaylsTokenRegistryService) List(ctx context.Context) ([]core.RegisteredToken, error) {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.registryAddr,
		Data: s.registry.PackGetAllTokens(),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("get all tokens: %w", err)
	}
	tokens, err := s.registry.UnpackGetAllTokens(out)
	if err != nil {
		return nil, fmt.Errorf("decode all tokens: %w", err)
	}
	return toRegisteredTokens(tokens), nil
}

func (s *RaylsTokenRegistryService) ListByStatus(
	ctx context.Context,
	status domain.PrivacyNodeStatus,
) ([]core.RegisteredToken, error) {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.registryAddr,
		Data: s.registry.PackGetTokensByPrivacyNodeStatus(uint8(status)),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("get tokens by status: %w", err)
	}
	tokens, err := s.registry.UnpackGetTokensByPrivacyNodeStatus(out)
	if err != nil {
		return nil, fmt.Errorf("decode tokens by status: %w", err)
	}
	return toRegisteredTokens(tokens), nil
}

func (s *RaylsTokenRegistryService) GetByAddress(
	ctx context.Context,
	tokenAddress string,
) (*core.RegisteredToken, error) {
	if !common.IsHexAddress(tokenAddress) {
		return nil, fmt.Errorf("invalid token address: %q", tokenAddress)
	}
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.registryAddr,
		Data: s.registry.PackGetTokenByAddress(common.HexToAddress(tokenAddress)),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("get token by address: %w", err)
	}
	token, err := s.registry.UnpackGetTokenByAddress(out)
	if err != nil {
		return nil, fmt.Errorf("decode token by address: %w", err)
	}
	registered := toRegisteredToken(token)
	return &registered, nil
}

func (s *RaylsTokenRegistryService) GetBySymbol(ctx context.Context, symbol string) (*core.RegisteredToken, error) {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.registryAddr,
		Data: s.registry.PackGetTokenBySymbol(symbol),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("get token by symbol: %w", err)
	}
	token, err := s.registry.UnpackGetTokenBySymbol(out)
	if err != nil {
		return nil, fmt.Errorf("decode token by symbol: %w", err)
	}
	registered := toRegisteredToken(token)
	return &registered, nil
}

func (s *RaylsTokenRegistryService) Exists(ctx context.Context, tokenAddress string) (bool, error) {
	if !common.IsHexAddress(tokenAddress) {
		return false, fmt.Errorf("invalid token address: %q", tokenAddress)
	}
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.registryAddr,
		Data: s.registry.PackTokenExists(common.HexToAddress(tokenAddress)),
	}, nil)
	if err != nil {
		return false, fmt.Errorf("check token exists: %w", err)
	}
	exists, err := s.registry.UnpackTokenExists(out)
	if err != nil {
		return false, fmt.Errorf("decode token exists result: %w", err)
	}
	return exists, nil
}

func toRegisteredTokens(tokens []PNTokenRegistryV1.TokenStructsToken) []core.RegisteredToken {
	res := make([]core.RegisteredToken, 0, len(tokens))
	for _, t := range tokens {
		res = append(res, toRegisteredToken(t))
	}
	return res
}

// toRegisteredToken maps an on-chain token struct to the read model, normalizing the address and
// converting the on-chain timestamp to UTC. Status is the PN-controlled privacyNodeStatus.
func toRegisteredToken(t PNTokenRegistryV1.TokenStructsToken) core.RegisteredToken {
	return core.RegisteredToken{
		Standard:     domain.ErcStandard(t.ErcStandard),
		Name:         t.Name,
		Symbol:       t.Symbol,
		URI:          t.Uri,
		TokenAddress: domain.NormalizeAddress(t.TokenAddress.Hex()),
		Status:       domain.PrivacyNodeStatus(t.PrivacyNodeStatus),
		LastUpdated:  time.Unix(t.UpdatedAt.Int64(), 0).UTC(),
	}
}

// signAndSend builds an unsigned tx to the TokenRegistry contract with the calldata, delegates
// signing to custody using the operator wallet, and polls for the receipt.
func (s *RaylsTokenRegistryService) signAndSend(
	ctx context.Context,
	operatorAddress string,
	calldata []byte,
) (string, error) {
	if !common.IsHexAddress(operatorAddress) {
		return "", fmt.Errorf("invalid operator address: %q", operatorAddress)
	}
	from := common.HexToAddress(operatorAddress)

	nonce, err := s.client.PendingNonceAt(ctx, from)
	if err != nil {
		return "", fmt.Errorf("get pending nonce: %w", err)
	}
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("suggest gas price: %w", err)
	}
	gasLimit, err := s.client.EstimateGas(ctx, ethereum.CallMsg{From: from, To: &s.registryAddr, Data: calldata})
	if err != nil {
		// eth_estimateGas re-executes the call, so a call that will revert (e.g. set-status on an
		// unregistered token) fails here before any tx is sent — no receipt to inspect. Classify it as
		// a revert so it surfaces as a client-correctable 422, consistent with the receipt-status-0 path.
		if contracts.IsRevertError(err) {
			if reason := contracts.SanitizeRevertForClient(err); reason != "" {
				return "", fmt.Errorf("%w: %s", core.ErrTxReverted, reason)
			}
			return "", fmt.Errorf("estimate gas reverted: %w", core.ErrTxReverted)
		}
		return "", fmt.Errorf("estimate gas: %w", err)
	}

	tx := types.NewTransaction(nonce, s.registryAddr, big.NewInt(0), gasLimit, gasPrice, calldata)
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("encode transaction: %w", err)
	}

	txHash, err := s.custody.SignAndTransact(ctx, txBytes, operatorAddress, s.chainID.String())
	if err != nil {
		return "", fmt.Errorf("sign and transact: %w", err)
	}

	// Bound the receipt wait so a dropped or replaced transaction can't hang the request forever.
	waitCtx, cancel := context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	hash := common.HexToHash(txHash)
	for {
		receipt, err := s.client.TransactionReceipt(waitCtx, hash)
		if errors.Is(err, ethereum.NotFound) {
			select {
			case <-waitCtx.Done():
				return "", fmt.Errorf("timed out waiting for receipt (tx %s): %w", txHash, waitCtx.Err())
			case <-time.After(2 * time.Second):
				continue
			}
		}
		if err != nil {
			return "", fmt.Errorf("get receipt: %w", err)
		}
		if receipt.Status == 0 {
			if reason := contracts.SimulateRevertReason(ctx, s.client,
				ethereum.CallMsg{From: from, To: &s.registryAddr, Data: calldata}, receipt.BlockNumber); reason != "" {
				return "", fmt.Errorf("%w (tx %s): %s", core.ErrTxReverted, txHash, reason)
			}
			return "", fmt.Errorf("%w (tx %s)", core.ErrTxReverted, txHash)
		}
		return txHash, nil
	}
}
