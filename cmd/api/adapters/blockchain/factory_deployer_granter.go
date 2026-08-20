package blockchain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/DeploymentProxyRegistryV1"
	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/RaylsAccessManagerV1"
)

var _ core.RoleGranter = (*FactoryDeployerGranter)(nil)

// grantExecutionDelay is the executionDelay passed to grantRole: 0 = the role takes effect
// immediately, with no timelock. Matches how the deploy/ops tooling grants FACTORY_DEPLOYER.
const grantExecutionDelay uint32 = 0

// FactoryDeployerGranter grants on-chain roles (FACTORY_DEPLOYER, PRIVACY_NODE_OPERATOR) to
// freshly-provisioned custody wallets so they can deploy tokens via RNContractFactory and pass
// the login role check. Each grant is signed by a privileged
// grantor EOA (the role admin) whose key lives in config — DEV-ONLY, mirroring FaucetFunder: in
// production, role assignment stays out of band and this granter is left nil.
//
// A mutex serializes grants so concurrent logins don't reuse the grantor's nonce.
type FactoryDeployerGranter struct {
	client  *ethclient.Client
	ram     *RaylsAccessManagerV1.RaylsAccessManagerV1
	ramAddr common.Address
	key     *ecdsa.PrivateKey
	from    common.Address
	chainID *big.Int
	mu      sync.Mutex
}

// NewFactoryDeployerGranter builds a granter from a hex private key (with or without 0x prefix),
// resolving the RaylsAccessManager address via DeploymentProxyRegistry. Reuses an existing RPC
// client (e.g. the one from RaylsAccessManagerService) to avoid a second dial.
func NewFactoryDeployerGranter(
	client *ethclient.Client,
	registryAddr, grantorKeyHex string,
) (*FactoryDeployerGranter, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(grantorKeyHex), "0x"))
	if err != nil {
		return nil, fmt.Errorf("parse role grantor private key: %w", err)
	}

	startupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	chainID, err := client.ChainID(startupCtx)
	if err != nil {
		return nil, fmt.Errorf("get chain ID for role granter: %w", err)
	}

	regAddr := common.HexToAddress(registryAddr)
	if codeErr := contracts.EnsureCode(startupCtx, client, regAddr); codeErr != nil {
		return nil, fmt.Errorf("instantiate DeploymentProxyRegistry: %w", codeErr)
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

	return &FactoryDeployerGranter{
		client:  client,
		ram:     RaylsAccessManagerV1.NewRaylsAccessManagerV1(),
		ramAddr: ramAddr,
		key:     key,
		from:    crypto.PubkeyToAddress(key.PublicKey),
		chainID: chainID,
	}, nil
}

// GrantRoles grants each named role to address, skipping any it already holds. Roles are granted
// in order; the first failure aborts and is returned (the caller treats it as best-effort).
func (g *FactoryDeployerGranter) GrantRoles(ctx context.Context, address string, roles []string) error {
	if !common.IsHexAddress(address) {
		return fmt.Errorf("invalid wallet address: %q", address)
	}
	account := common.HexToAddress(address)

	for _, role := range roles {
		if err := g.grantRole(ctx, role, account); err != nil {
			return err
		}
	}
	return nil
}

// grantRole grants a single named role to account, unless it already holds it.
func (g *FactoryDeployerGranter) grantRole(ctx context.Context, roleName string, account common.Address) error {
	roleID, err := g.roleIDByName(ctx, roleName)
	if err != nil {
		return err
	}

	// The lock is taken BEFORE the hasRole read, not just around the send: it serializes the
	// grantor's nonce, and a check made outside it can be invalidated by a concurrent grant
	// that lands between the read and the send. Two logins racing on the same wallet would
	// both see "no role" and both submit — the second reverts, failing a login that had
	// nothing wrong with it.
	g.mu.Lock()
	defer g.mu.Unlock()

	hasRole, err := g.hasRole(ctx, roleID, account)
	if err != nil {
		return err
	}
	if hasRole {
		return nil // already granted — idempotent no-op
	}

	calldata := g.ram.PackGrantRole(roleID, account, grantExecutionDelay)

	nonce, err := g.client.PendingNonceAt(ctx, g.from)
	if err != nil {
		return fmt.Errorf("get grantor nonce: %w", err)
	}
	gasPrice, err := g.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("suggest gas price: %w", err)
	}
	gasLimit, err := g.client.EstimateGas(ctx, ethereum.CallMsg{From: g.from, To: &g.ramAddr, Data: calldata})
	if err != nil {
		return fmt.Errorf("estimate gas for grantRole: %w", err)
	}

	tx := types.NewTransaction(nonce, g.ramAddr, big.NewInt(0), gasLimit, gasPrice, calldata)
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(g.chainID), g.key)
	if err != nil {
		return fmt.Errorf("sign grantRole tx: %w", err)
	}
	if err := g.client.SendTransaction(ctx, signed); err != nil {
		return fmt.Errorf("send grantRole tx: %w", err)
	}

	// Wait briefly for inclusion so the role is effective on the immediately-following deploy.
	receiptCtx, cancel := context.WithTimeout(ctx, 90*time.Second)
	defer cancel()
	for {
		receipt, err := g.client.TransactionReceipt(receiptCtx, signed.Hash())
		if errors.Is(err, ethereum.NotFound) {
			select {
			case <-receiptCtx.Done():
				return fmt.Errorf("timed out waiting for grantRole tx %s: %w", signed.Hash().Hex(), receiptCtx.Err())
			case <-time.After(2 * time.Second):
				continue
			}
		}
		if err != nil {
			return fmt.Errorf("get grantRole receipt: %w", err)
		}
		if receipt.Status == 0 {
			return fmt.Errorf("grantRole %w (tx %s) — grantor %s may lack admin authority over %s",
				core.ErrTxReverted, signed.Hash().Hex(), g.from.Hex(), roleName)
		}
		return nil
	}
}

// roleIDByName resolves the numeric role ID for a role name via getRoleIdByName.
func (g *FactoryDeployerGranter) roleIDByName(ctx context.Context, name string) (uint64, error) {
	out, err := g.client.CallContract(ctx, ethereum.CallMsg{
		To:   &g.ramAddr,
		Data: g.ram.PackGetRoleIdByName(name),
	}, nil)
	if err != nil {
		return 0, fmt.Errorf("getRoleIdByName(%s): %w", name, err)
	}
	roleID, err := g.ram.UnpackGetRoleIdByName(out)
	if err != nil {
		return 0, fmt.Errorf("decode getRoleIdByName(%s): %w", name, err)
	}
	return roleID, nil
}

// hasRole returns whether account already holds the given role (global scope).
func (g *FactoryDeployerGranter) hasRole(ctx context.Context, roleID uint64, account common.Address) (bool, error) {
	out, err := g.client.CallContract(ctx, ethereum.CallMsg{
		To:   &g.ramAddr,
		Data: g.ram.PackHasRole(roleID, account),
	}, nil)
	if err != nil {
		return false, fmt.Errorf("hasRole: %w", err)
	}
	res, err := g.ram.UnpackHasRole(out)
	if err != nil {
		return false, fmt.Errorf("decode hasRole: %w", err)
	}
	return res.IsMember, nil
}
