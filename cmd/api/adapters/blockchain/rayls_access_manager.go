package blockchain

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts/DeploymentProxyRegistryV1"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/contracts/RaylsAccessManagerV1"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

var (
	_ core.RoleService              = (*RaylsAccessManagerService)(nil)
	_ core.RaylsAccessManagerClient = (*RaylsAccessManagerService)(nil)
)

// RaylsAccessManagerService uses RaylsAccessManager (resolved via DeploymentProxyRegistry)
// for on-chain role lookup (getAccountRolesWithInfo). Role assignment is performed out of
// band by the deploy/ops tooling, so this service is read-only.
type RaylsAccessManagerService struct {
	client  *ethclient.Client
	ram     *RaylsAccessManagerV1.RaylsAccessManagerV1
	ramAddr common.Address
}

// NewRaylsAccessManagerService connects to the RPC and resolves the RaylsAccessManager
// address via DeploymentProxyRegistry. A 30-second timeout covers the sequential startup
// RPC calls (getContract + bytecode checks).
func NewRaylsAccessManagerService(rpcURL, registryAddr string) (*RaylsAccessManagerService, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RPC %s: %w", rpcURL, err)
	}

	startupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	regAddr := common.HexToAddress(registryAddr)
	if err := contracts.EnsureCode(startupCtx, client, regAddr); err != nil {
		return nil, fmt.Errorf("instantiate DeploymentProxyRegistry: %w", err)
	}
	registry := DeploymentProxyRegistryV1.NewDeploymentProxyRegistryV1()

	regOut, err := client.CallContract(startupCtx, ethereum.CallMsg{
		To:   &regAddr,
		Data: registry.PackGetContract("RaylsAccessManager"),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("resolve RaylsAccessManager address: %w", err)
	}
	ramAddr, err := registry.UnpackGetContract(regOut)
	if err != nil {
		return nil, fmt.Errorf("decode RaylsAccessManager address: %w", err)
	}
	if ramAddr == (common.Address{}) {
		return nil, fmt.Errorf("DeploymentProxyRegistry returned zero address for RaylsAccessManager")
	}

	if err := contracts.EnsureCode(startupCtx, client, ramAddr); err != nil {
		return nil, fmt.Errorf("instantiate RaylsAccessManager: %w", err)
	}
	ram := RaylsAccessManagerV1.NewRaylsAccessManagerV1()

	return &RaylsAccessManagerService{
		client:  client,
		ram:     ram,
		ramAddr: ramAddr,
	}, nil
}

// GetRoles calls getAccountRolesWithInfo and returns the role labels for the given wallet address.
// Returns an empty slice (not an error) when the wallet has no roles assigned.
func (s *RaylsAccessManagerService) GetRoles(ctx context.Context, walletAddress string) ([]string, error) {
	if !common.IsHexAddress(walletAddress) {
		return nil, fmt.Errorf("invalid wallet address: %q", walletAddress)
	}

	infos, err := s.accountRolesWithInfo(ctx, common.HexToAddress(walletAddress))
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(infos))
	for _, r := range infos {
		names = append(names, domain.NormalizeRoleLabel(r.Label))
	}
	return names, nil
}

// GetUserRole calls getAccountRolesWithInfo on RaylsAccessManager for the given wallet address.
// Returns role names joined by commas, or an empty string if the wallet has no roles assigned.
// userID is unused — RAM only needs the wallet address.
func (s *RaylsAccessManagerService) GetUserRole(
	ctx context.Context,
	_ uuid.UUID,
	walletAddress string,
) (string, error) {
	if !common.IsHexAddress(walletAddress) {
		return "", fmt.Errorf("invalid wallet address: %q", walletAddress)
	}

	infos, err := s.accountRolesWithInfo(ctx, common.HexToAddress(walletAddress))
	if err != nil {
		return "", err
	}

	names := make([]string, 0, len(infos))
	for _, r := range infos {
		names = append(names, domain.NormalizeRoleLabel(r.Label))
	}

	return strings.Join(names, ","), nil
}

// accountRolesWithInfo performs the getAccountRolesWithInfo view call and decodes the result.
func (s *RaylsAccessManagerService) accountRolesWithInfo(
	ctx context.Context,
	account common.Address,
) ([]RaylsAccessManagerV1.IRaylsAccessManagerRoleInfo, error) {
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &s.ramAddr,
		Data: s.ram.PackGetAccountRolesWithInfo(account),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("getAccountRolesWithInfo: %w", err)
	}
	infos, err := s.ram.UnpackGetAccountRolesWithInfo(out)
	if err != nil {
		return nil, fmt.Errorf("decode getAccountRolesWithInfo: %w", err)
	}
	return infos, nil
}

// EthClient returns the underlying Ethereum RPC client.
func (s *RaylsAccessManagerService) EthClient() *ethclient.Client { return s.client }

// Contract returns the bound RaylsAccessManagerV1 contract instance.
func (s *RaylsAccessManagerService) Contract() *RaylsAccessManagerV1.RaylsAccessManagerV1 {
	return s.ram
}

// ContractAddress returns the resolved on-chain address of RaylsAccessManager.
func (s *RaylsAccessManagerService) ContractAddress() common.Address { return s.ramAddr }

// Close releases the underlying WebSocket/HTTP connection to the RPC node.
func (s *RaylsAccessManagerService) Close() { s.client.Close() }
