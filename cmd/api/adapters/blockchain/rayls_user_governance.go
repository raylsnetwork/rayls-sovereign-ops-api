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
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts/RNUserGovernanceV1"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

var _ core.UserGovernanceService = (*RaylsUserGovernanceService)(nil)

// RaylsUserGovernanceService manages on-chain user identity and address pairs in RNUserGovernance.
// All writes are operator-authority writes signed via custody by the resolved operator wallet —
// the operator key never leaves the HSM.
type RaylsUserGovernanceService struct {
	client  *ethclient.Client
	gov     *RNUserGovernanceV1.RNUserGovernanceV1
	govAddr common.Address
	custody core.CustodyService
	chainID *big.Int
}

// NewRaylsUserGovernanceService reuses an existing RPC client (e.g. from RaylsAccessManagerService)
// and resolves the RNUserGovernance address via DeploymentProxyRegistry. A 30-second timeout covers
// the sequential startup RPC calls (ChainID, getContract).
func NewRaylsUserGovernanceService(
	client *ethclient.Client,
	registryAddr string,
	custody core.CustodyService,
) (*RaylsUserGovernanceService, error) {
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
		Data: registry.PackGetContract("RNUserGovernance"),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("resolve RNUserGovernance address: %w", err)
	}
	govAddr, err := registry.UnpackGetContract(regOut)
	if err != nil {
		return nil, fmt.Errorf("decode RNUserGovernance address: %w", err)
	}
	if govAddr == (common.Address{}) {
		return nil, fmt.Errorf("DeploymentProxyRegistry returned zero address for RNUserGovernance")
	}

	if err := contracts.EnsureCode(startupCtx, client, govAddr); err != nil {
		return nil, fmt.Errorf("instantiate RNUserGovernance: %w", err)
	}

	return &RaylsUserGovernanceService{
		client:  client,
		gov:     RNUserGovernanceV1.NewRNUserGovernanceV1(),
		govAddr: govAddr,
		custody: custody,
		chainID: chainID,
	}, nil
}

// EnsureUser idempotently creates the on-chain user: it reads HasUser and only signs a CreateUser
// transaction when the user does not already exist on-chain.
func (s *RaylsUserGovernanceService) EnsureUser(
	ctx context.Context,
	operatorAddress string,
	onChainUserID [32]byte,
) error {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.govAddr,
		Data: s.gov.PackHasUser(onChainUserID),
	}, nil)
	if err != nil {
		return fmt.Errorf("check on-chain user exists: %w", err)
	}
	exists, err := s.gov.UnpackHasUser(out)
	if err != nil {
		return fmt.Errorf("decode HasUser result: %w", err)
	}
	if exists {
		return nil
	}

	if _, err := s.signAndSend(ctx, operatorAddress, s.gov.PackCreateUser(onChainUserID)); err != nil {
		return fmt.Errorf("create on-chain user: %w", err)
	}
	return nil
}

func (s *RaylsUserGovernanceService) AddAddressPair(
	ctx context.Context,
	operatorAddress string,
	onChainUserID [32]byte,
	publicAddr, privateAddr string,
) (string, error) {
	if !common.IsHexAddress(publicAddr) {
		return "", fmt.Errorf("invalid public address: %q", publicAddr)
	}
	if !common.IsHexAddress(privateAddr) {
		return "", fmt.Errorf("invalid private address: %q", privateAddr)
	}
	calldata := s.gov.PackAddAddressPair(
		onChainUserID,
		common.HexToAddress(publicAddr),
		common.HexToAddress(privateAddr),
	)
	return s.signAndSend(ctx, operatorAddress, calldata)
}

func (s *RaylsUserGovernanceService) SetApprovalStatus(
	ctx context.Context,
	operatorAddress string,
	onChainUserID [32]byte,
	publicAddr, privateAddr string,
	status domain.ApprovalStatus,
) (string, error) {
	if !common.IsHexAddress(publicAddr) {
		return "", fmt.Errorf("invalid public address: %q", publicAddr)
	}
	if !common.IsHexAddress(privateAddr) {
		return "", fmt.Errorf("invalid private address: %q", privateAddr)
	}
	calldata := s.gov.PackSetAddressPairApprovalStatus(
		onChainUserID,
		common.HexToAddress(publicAddr),
		common.HexToAddress(privateAddr),
		uint8(status),
	)
	return s.signAndSend(ctx, operatorAddress, calldata)
}

// ListAllPending reads every pending address pair across all on-chain users, grouped by the raw
// keccak256 user ID. It backs the admin discovery endpoint.
func (s *RaylsUserGovernanceService) ListAllPending(ctx context.Context) ([]core.OnChainPendingGroup, error) {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.govAddr,
		Data: s.gov.PackGetAllPendingAddressPairs(),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("get all pending address pairs: %w", err)
	}
	decoded, err := s.gov.UnpackGetAllPendingAddressPairs(out)
	if err != nil {
		return nil, fmt.Errorf("decode all pending address pairs: %w", err)
	}

	groups := make([]core.OnChainPendingGroup, 0, len(decoded.Arg0))
	for i, id := range decoded.Arg0 {
		groups = append(groups, core.OnChainPendingGroup{
			OnChainUserID: id,
			AddressPairs:  toOnChainAddressPairs(decoded.Arg1[i]),
		})
	}
	return groups, nil
}

func (s *RaylsUserGovernanceService) ListPending(
	ctx context.Context,
	onChainUserID [32]byte,
) ([]core.OnChainAddressPair, error) {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.govAddr,
		Data: s.gov.PackGetPendingAddressPairs(onChainUserID),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("get pending address pairs: %w", err)
	}
	pairs, err := s.gov.UnpackGetPendingAddressPairs(out)
	if err != nil {
		return nil, fmt.Errorf("decode pending address pairs: %w", err)
	}
	return toOnChainAddressPairs(pairs), nil
}

func (s *RaylsUserGovernanceService) ListApproved(
	ctx context.Context,
	onChainUserID [32]byte,
) ([]core.OnChainAddressPair, error) {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.govAddr,
		Data: s.gov.PackGetApprovedAddressPairs(onChainUserID),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("get approved address pairs: %w", err)
	}
	pairs, err := s.gov.UnpackGetApprovedAddressPairs(out)
	if err != nil {
		return nil, fmt.Errorf("decode approved address pairs: %w", err)
	}
	return toOnChainAddressPairs(pairs), nil
}

// toOnChainAddressPairs maps the on-chain address-pair structs to the read model, normalizing
// addresses and converting the on-chain timestamp to UTC.
func toOnChainAddressPairs(pairs []RNUserGovernanceV1.IUserGovernanceAddressPair) []core.OnChainAddressPair {
	res := make([]core.OnChainAddressPair, 0, len(pairs))
	for _, p := range pairs {
		res = append(res, core.OnChainAddressPair{
			PublicChainAddress:  domain.NormalizeAddress(p.PublicAddress.Hex()),
			PrivateChainAddress: domain.NormalizeAddress(p.PrivateAddress.Hex()),
			Status:              domain.ApprovalStatus(p.ApprovalStatus),
			CreatedAt:           time.Unix(p.CreatedAt.Int64(), 0).UTC(),
		})
	}
	return res
}

// signAndSend builds an unsigned tx to the RNUserGovernance contract with the calldata, delegates
// signing to custody using the operator wallet, and polls for the receipt.
func (s *RaylsUserGovernanceService) signAndSend(
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
	gasLimit, err := s.client.EstimateGas(ctx, ethereum.CallMsg{From: from, To: &s.govAddr, Data: calldata})
	if err != nil {
		// eth_estimateGas re-executes the call, so a call that will revert (e.g. approving a pair whose
		// address is not mapped to the user) fails here before any tx is sent — no receipt to inspect.
		// Classify it as a revert so it surfaces as a client-correctable 422, consistent with the
		// receipt-status-0 path.
		if contracts.IsRevertError(err) {
			if reason := contracts.SanitizeRevertForClient(err); reason != "" {
				return "", fmt.Errorf("%w: %s", core.ErrTxReverted, reason)
			}
			return "", fmt.Errorf("estimate gas reverted: %w", core.ErrTxReverted)
		}
		return "", fmt.Errorf("estimate gas: %w", err)
	}

	tx := types.NewTransaction(nonce, s.govAddr, big.NewInt(0), gasLimit, gasPrice, calldata)
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
				ethereum.CallMsg{From: from, To: &s.govAddr, Data: calldata}, receipt.BlockNumber); reason != "" {
				return "", fmt.Errorf("%w (tx %s): %s", core.ErrTxReverted, txHash, reason)
			}
			return "", fmt.Errorf("%w (tx %s)", core.ErrTxReverted, txHash)
		}
		return txHash, nil
	}
}
