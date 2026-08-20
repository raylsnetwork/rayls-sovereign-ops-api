package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts/DeploymentProxyRegistryV1"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts/RNContractFactoryV1"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

var _ core.TokenDeployService = (*RaylsContractFactoryService)(nil)

// RaylsContractFactoryService deploys protocol tokens through the RNContractFactory
// (resolved via DeploymentProxyRegistry under the key "RNContractFactory").
// Each deploy is signed by the caller's custody wallet — the key never leaves the HSM.
type RaylsContractFactoryService struct {
	client      *ethclient.Client
	factory     *RNContractFactoryV1.RNContractFactoryV1
	factoryAddr common.Address
	custody     core.CustodyService
	chainID     *big.Int
}

// NewRaylsContractFactoryService resolves the RNContractFactory address via DeploymentProxyRegistry,
// reusing an existing RPC client (e.g. the one created by RaylsAccessManagerService) to avoid a second dial.
func NewRaylsContractFactoryService(
	client *ethclient.Client,
	registryAddr string,
	custody core.CustodyService,
) (*RaylsContractFactoryService, error) {
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
	registry := DeploymentProxyRegistryV1.NewDeploymentProxyRegistryV1()

	regOut, err := client.CallContract(startupCtx, ethereum.CallMsg{
		To:   &regAddr,
		Data: registry.PackGetContract("RNContractFactory"),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("resolve RNContractFactory address: %w", err)
	}
	factoryAddr, err := registry.UnpackGetContract(regOut)
	if err != nil {
		return nil, fmt.Errorf("decode RNContractFactory address: %w", err)
	}
	if factoryAddr == (common.Address{}) {
		return nil, fmt.Errorf("DeploymentProxyRegistry returned zero address for RNContractFactory")
	}

	if err := contracts.EnsureCode(startupCtx, client, factoryAddr); err != nil {
		return nil, fmt.Errorf("instantiate RNContractFactory: %w", err)
	}

	return &RaylsContractFactoryService{
		client:      client,
		factory:     RNContractFactoryV1.NewRNContractFactoryV1(),
		factoryAddr: factoryAddr,
		custody:     custody,
		chainID:     chainID,
	}, nil
}

// ChainID returns the chain ID the factory is deployed on (decimal string).
func (s *RaylsContractFactoryService) ChainID() string {
	return s.chainID.String()
}

// EstimateDeploy returns the real on-chain gas estimate for deploying spec from signerAddress,
// without executing the deploy (eth_estimateGas + eth_gasPrice). Same calldata the Deploy path
// builds, so the estimate reflects the exact transaction that would be signed.
func (s *RaylsContractFactoryService) EstimateDeploy(
	ctx context.Context,
	signerAddress string,
	spec core.TokenDeploySpec,
) (core.TokenDeployEstimate, error) {
	if !common.IsHexAddress(signerAddress) {
		return core.TokenDeployEstimate{}, fmt.Errorf("invalid signer address: %q", signerAddress)
	}

	calldata, err := s.deployCalldata(spec)
	if err != nil {
		return core.TokenDeployEstimate{}, err
	}

	from := common.HexToAddress(signerAddress)

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return core.TokenDeployEstimate{}, fmt.Errorf("suggest gas price: %w", err)
	}

	gasLimit, err := s.client.EstimateGas(ctx, ethereum.CallMsg{
		From: from,
		To:   &s.factoryAddr,
		Data: calldata,
	})
	if err != nil {
		return core.TokenDeployEstimate{}, fmt.Errorf("estimate gas: %w", err)
	}

	totalFee := new(big.Int).Mul(new(big.Int).SetUint64(gasLimit), gasPrice)
	return core.TokenDeployEstimate{
		GasLimit:    gasLimit,
		GasPriceWei: gasPrice.String(),
		TotalFeeWei: totalFee.String(),
	}, nil
}

// Deploy builds the typed factory deploy calldata for spec, signs it via the custody wallet at
// signerAddress, polls for the receipt, and returns the deployed token address and tx hash.
func (s *RaylsContractFactoryService) Deploy(
	ctx context.Context,
	signerAddress string,
	spec core.TokenDeploySpec,
) (string, string, error) {
	if !common.IsHexAddress(signerAddress) {
		return "", "", fmt.Errorf("invalid signer address: %q", signerAddress)
	}

	calldata, err := s.deployCalldata(spec)
	if err != nil {
		return "", "", err
	}

	from := common.HexToAddress(signerAddress)

	nonce, err := s.client.PendingNonceAt(ctx, from)
	if err != nil {
		return "", "", fmt.Errorf("get pending nonce: %w", err)
	}

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", "", fmt.Errorf("suggest gas price: %w", err)
	}

	gasLimit, err := s.client.EstimateGas(ctx, ethereum.CallMsg{
		From: from,
		To:   &s.factoryAddr,
		Data: calldata,
	})
	if err != nil {
		return "", "", fmt.Errorf("estimate gas: %w", err)
	}

	tx := types.NewTransaction(nonce, s.factoryAddr, big.NewInt(0), gasLimit, gasPrice, calldata)
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return "", "", fmt.Errorf("encode transaction: %w", err)
	}

	txHash, err := s.custody.SignAndTransact(ctx, txBytes, signerAddress, s.chainID.String())
	if err != nil {
		return "", "", fmt.Errorf("sign and transact deploy: %w", err)
	}

	receipt, err := s.waitForReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return "", "", err
	}
	if receipt.Status == 0 {
		return "", "", fmt.Errorf("token deploy %w (tx %s)", core.ErrTxReverted, txHash)
	}

	deployedAddr, err := s.extractDeployedAddress(receipt)
	if err != nil {
		return "", "", err
	}
	return deployedAddr.Hex(), txHash, nil
}

// deployCalldata maps the token standard to the matching typed factory deploy method.
// The resourceId is intentionally left as bytes32(0) — it is not assigned at deploy time.
func (s *RaylsContractFactoryService) deployCalldata(spec core.TokenDeploySpec) ([]byte, error) {
	var resourceID [32]byte // bytes32(0)

	switch spec.ErcStandard {
	case domain.ErcStandardERC20:
		return s.factory.PackDeployErc20(spec.Name, spec.Symbol, spec.Decimals, resourceID), nil
	case domain.ErcStandardEnygma:
		return s.factory.PackDeployEnygma(spec.Name, spec.Symbol, spec.Decimals, resourceID), nil
	case domain.ErcStandardERC721:
		return s.factory.PackDeployErc721(spec.URI, spec.Name, spec.Symbol, resourceID), nil
	case domain.ErcStandardERC1155:
		return s.factory.PackDeployErc1155(spec.URI, spec.Name, resourceID), nil
	case domain.ErcStandardZkDvpERC721:
		return s.factory.PackDeployErc721Dvp(spec.URI, spec.Name, spec.Symbol, resourceID), nil
	case domain.ErcStandardZkDvpERC1155:
		return s.factory.PackDeployErc1155Dvp(spec.URI, spec.Name, resourceID), nil
	case domain.ErcStandardStableCoin:
		// The *AsUser* wrapper, so the DEPLOYING WALLET becomes trusted.owner rather than the
		// factory's own factoryOwner. That matters uniquely for the stablecoin: its handler
		// seeds masterMinter, pauser and blacklister from trusted.owner, so the plain
		// deployRegistered path left every token's compliance roles on a single protocol
		// address and no user could pause the token they had just deployed.
		//
		// The handler decodes the same (name, symbol, decimals) tuple as the ERC20 handler.
		// No resourceId argument: the *AsUser wrappers hardcode bytes32(0) on-chain, since the
		// id is only assigned later by the PNH TokenRegistry via setResourceId.
		return s.factory.PackDeployStableCoinAsUser(spec.Name, spec.Symbol, spec.Decimals), nil
	default:
		return nil, fmt.Errorf("unsupported token standard for factory deploy: %d", spec.ErcStandard)
	}
}

// extractDeployedAddress reads the deployed token address from the factory's deployment event.
// Typed deploys emit RegisteredContractDeployed; the generic deploy emits ContractDeployed.
func (s *RaylsContractFactoryService) extractDeployedAddress(receipt *types.Receipt) (common.Address, error) {
	for _, log := range receipt.Logs {
		if log.Address != s.factoryAddr {
			continue
		}
		if ev, err := s.factory.UnpackRegisteredContractDeployedEvent(log); err == nil {
			return ev.DeployedAddress, nil
		}
		if ev, err := s.factory.UnpackContractDeployedEvent(log); err == nil {
			return ev.DeployedAddress, nil
		}
	}
	return common.Address{}, fmt.Errorf("deployment event not found in receipt %s", receipt.TxHash.Hex())
}

// waitForReceipt polls until the transaction receipt is available or ctx is cancelled.
func (s *RaylsContractFactoryService) waitForReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := s.client.TransactionReceipt(ctx, hash)
		if errors.Is(err, ethereum.NotFound) {
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("context cancelled waiting for deploy receipt: %w", ctx.Err())
			case <-time.After(2 * time.Second):
				continue
			}
		}
		if err != nil {
			return nil, fmt.Errorf("get deploy receipt: %w", err)
		}
		return receipt, nil
	}
}
