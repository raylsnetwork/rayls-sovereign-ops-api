package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

var _ core.TokenPermissionService = (*TokenPermissionService)(nil)

// TokenPermissionService derives, from the Access Manager tables, which functions a wallet may
// call on a token contract: the intersection of the wallet's active roles (global + contract-scoped)
// with the contract's function→role permissions.
type TokenPermissionService struct {
	members   core.AccessManagerRoleMemberRepository
	scoped    core.AccessManagerContractScopedRoleMemberRepository
	fnPerms   core.AccessManagerFunctionPermissionRepository
	contracts core.AccessManagerManagedContractRepository
	// tokens and actions resolve the stablecoin-only pause capability, which no Access Manager
	// table can answer (pause() is gated on the contract's own `pauser` address). Both nil on a
	// chainless deployment — CanPause/IsTokenPaused then stay false rather than erroring.
	tokens  core.TokenRepository
	actions core.TokenActionService
}

func NewTokenPermissionService(
	members core.AccessManagerRoleMemberRepository,
	scoped core.AccessManagerContractScopedRoleMemberRepository,
	fnPerms core.AccessManagerFunctionPermissionRepository,
	contracts core.AccessManagerManagedContractRepository,
	tokens core.TokenRepository,
	actions core.TokenActionService,
) *TokenPermissionService {
	return &TokenPermissionService{
		members:   members,
		scoped:    scoped,
		fnPerms:   fnPerms,
		contracts: contracts,
		tokens:    tokens,
		actions:   actions,
	}
}

// SetTokenActionService supplies the on-chain reader used for the stablecoin pause capability.
//
// Injected after construction because the token action service is built later in the container
// (it needs the AccessManager's RPC client), while this service is needed before that point.
func (s *TokenPermissionService) SetTokenActionService(actions core.TokenActionService) {
	s.actions = actions
}

func (s *TokenPermissionService) GetTokenPermissions(
	ctx context.Context,
	contractAddress, walletAddress string,
) (*core.TokenPermissions, error) {
	contract := domain.NormalizeAddress(contractAddress)
	wallet := domain.NormalizeAddress(walletAddress)

	roleIDs, err := s.activeRoleIDs(ctx, wallet, contract)
	if err != nil {
		return nil, err
	}

	perms, err := s.fnPerms.ListByContract(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("list function permissions: %w", err)
	}

	result := &core.TokenPermissions{
		ContractAddress: contract,
		WalletAddress:   wallet,
		Functions:       []core.TokenFunction{},
	}

	seen := make(map[string]struct{})
	for _, p := range perms {
		if !p.IsActive {
			continue
		}
		if _, ok := roleIDs[p.RoleID]; !ok {
			continue
		}
		sel := strings.ToLower(p.Selector)
		if _, dup := seen[sel]; dup {
			continue
		}
		seen[sel] = struct{}{}

		name := tokenFunctionName(sel)
		result.Functions = append(result.Functions, core.TokenFunction{Selector: sel, Name: name})
		switch name {
		case fnMint:
			result.CanMint = true
		case fnBurn:
			result.CanBurn = true
		case fnSubmitTokenUpdate:
			result.CanSubmitTokenUpdate = true
		}
	}

	mc, err := s.contracts.FindByAddress(ctx, contract)
	if err != nil && !errors.Is(err, core.ErrRecordNotFound) {
		return nil, fmt.Errorf("find managed contract: %w", err)
	}
	if mc != nil {
		result.IsPaused = mc.IsPaused
	}

	s.resolvePause(ctx, contract, wallet, result)

	return result, nil
}

// resolvePause fills CanPause/IsTokenPaused for stablecoins by reading the contract directly.
//
// Best-effort: a failed read leaves both false rather than failing the whole permissions call,
// which the UI needs to render mint/burn regardless. Skipped entirely for non-stablecoins (no
// pause function) and on chainless deployments (no client to read with).
func (s *TokenPermissionService) resolvePause(
	ctx context.Context,
	contract, wallet string,
	result *core.TokenPermissions,
) {
	if s.tokens == nil || s.actions == nil {
		return
	}

	token, err := s.tokens.FindByAddress(ctx, contract)
	if err != nil || token == nil || token.ErcStandard != domain.ErcStandardStableCoin {
		return
	}

	if pauser, err := s.actions.Pauser(ctx, contract); err == nil {
		result.CanPause = strings.EqualFold(domain.NormalizeAddress(pauser), wallet)
	}
	if paused, err := s.actions.IsPaused(ctx, contract); err == nil {
		result.IsTokenPaused = paused
	}
}

// activeRoleIDs returns the set of role IDs the wallet holds that apply to the contract:
// active global roles plus active roles scoped to this contract.
func (s *TokenPermissionService) activeRoleIDs(
	ctx context.Context,
	wallet, contract string,
) (map[uint64]struct{}, error) {
	roleIDs := make(map[uint64]struct{})

	globals, err := s.members.ListByAccount(ctx, wallet)
	if err != nil {
		return nil, fmt.Errorf("list global roles: %w", err)
	}
	for _, m := range globals {
		if m.IsActive {
			roleIDs[m.RoleID] = struct{}{}
		}
	}

	scopedMembers, err := s.scoped.ListByAccount(ctx, wallet)
	if err != nil {
		return nil, fmt.Errorf("list contract-scoped roles: %w", err)
	}
	for _, m := range scopedMembers {
		if m.IsActive && domain.NormalizeAddress(m.ContractAddress) == contract {
			roleIDs[m.RoleID] = struct{}{}
		}
	}

	return roleIDs, nil
}
