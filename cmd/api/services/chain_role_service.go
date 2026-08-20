package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
)

// ChainRoleService answers "what may this user do ON THIS CHAIN?".
//
// Roles used to be resolved once at login and frozen into the JWT. That stopped working
// when identity became shared: one token is presented to every chain, and a role granted
// on chain A says nothing about chain B. So each chain's ops-api now resolves roles for
// itself, per request, from the am_* tables its own AccessManager indexer maintains.
//
// A local database read, not an RPC — the indexer has already mirrored the chain state.
// It also fixes a long-standing wart: roles were a login-time snapshot, so a revoked role
// kept working until the user logged in again. Now revocation takes effect as soon as the
// indexer sees it.
type ChainRoleService struct {
	wallets core.UserWalletRepository
	// users tells "this account has no wallet yet" apart from "this account does not
	// exist" — a token minted against an identity database that has since been recreated.
	// Without it both look like "no wallet", and a stale session degrades into a working
	// login with no permissions instead of a prompt to sign in again.
	users   core.UserRepository
	members core.AccessManagerRoleMemberRepository
	roles   core.AccessManagerRoleRepository
	log     logger.Logger
}

// NewChainRoleService creates the per-chain role resolver.
func NewChainRoleService(
	wallets core.UserWalletRepository,
	users core.UserRepository,
	members core.AccessManagerRoleMemberRepository,
	roles core.AccessManagerRoleRepository,
	log logger.Logger,
) *ChainRoleService {
	return &ChainRoleService{
		wallets: wallets,
		users:   users,
		members: members,
		roles:   roles,
		log:     log,
	}
}

// RolesForUser returns the role NAMES (e.g. "PRIVACY_NODE_OPERATOR") the user holds on
// this chain, via their custody wallet.
//
// The wallet comes from the shared identity service — one EVM keypair per user, the same
// address on every chain. This service never mints one: if a user has no wallet, that is an
// identity-side gap, not something a chain can fix.
//
// An empty slice — not an error — is the correct answer for a user with no wallet or no
// grants yet: they are a valid account that may do nothing on this chain. Callers turn
// that into a 403.
func (s *ChainRoleService) RolesForUser(ctx context.Context, userID uuid.UUID) ([]string, error) {
	wallet, err := s.wallets.GetSignerWalletForChain(ctx, userID, domain.WalletChainPrivate)
	if err != nil && !errors.Is(err, core.ErrRecordNotFound) {
		return nil, fmt.Errorf("find signer wallet: %w", err)
	}

	if wallet == nil || wallet.RaylsAddress == "" {
		// No wallet has two very different causes. Separate them, because one is the user's
		// problem to wait out and the other is a session they must replace.
		if _, userErr := s.users.FindByID(ctx, userID); userErr != nil {
			if errors.Is(userErr, core.ErrRecordNotFound) {
				// The token names a user identity has never heard of — almost always a
				// session minted against an identity database that was since recreated
				// (a --clean). Returning "no roles" would present a working login with no
				// permissions and no way out; 401 makes the UI re-authenticate.
				s.log.Warn("token names an unknown user — stale session, forcing re-login",
					"userID", userID)
				return nil, core.NewUnauthorizedError("session is no longer valid, please sign in again")
			}
			return nil, fmt.Errorf("find user: %w", userErr)
		}
		s.log.Warn("user has no custody wallet — identity has not provisioned one", "userID", userID)
		return nil, nil
	}

	return s.RolesForWallet(ctx, wallet.RaylsAddress)
}

// RolesForWallet returns the role names held by an address on this chain.
func (s *ChainRoleService) RolesForWallet(ctx context.Context, walletAddress string) ([]string, error) {
	address := domain.NormalizeAddress(walletAddress)

	members, err := s.members.ListByAccount(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("list role members: %w", err)
	}

	names := make([]string, 0, len(members))
	for _, m := range members {
		if !m.IsActive {
			continue
		}
		role, roleErr := s.roles.FindByID(ctx, m.RoleID)
		if roleErr != nil {
			// A membership whose role the indexer has not seen yet: skip it rather than
			// failing the whole request, or one unindexed role would lock the user out.
			s.log.Warn("role member references an unknown role", "roleID", m.RoleID, "account", address)
			continue
		}
		if role.Name != "" {
			names = append(names, role.Name)
		}
	}
	return names, nil
}
