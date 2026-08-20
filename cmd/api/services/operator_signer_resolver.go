package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

var _ core.OperatorSignerResolver = (*operatorSignerResolver)(nil)

// operatorSignerResolver determines, at request time, which HSM operator wallet signs
// operator-authority governance writes. Governance writes are split across two on-chain roles
// (PRIVACY_NODE_OPERATOR gates updateTokenStatus/createUser/setAddressPairApprovalStatus;
// BANK_EMPLOYEE gates addToken/addAddressPair), so a single signer must hold both. The resolver asks
// the member repository for the one account holding both roles and returns its HSM wallet.
type operatorSignerResolver struct {
	roles   core.AccessManagerRoleRepository
	members core.AccessManagerRoleMemberRepository
	wallets core.UserWalletRepository
	log     logger.Logger
}

// NewOperatorSignerResolver builds an OperatorSignerResolver. It resolves the
// domain.RolePrivacyNodeOperator and domain.RoleBankEmployee roles from the Access Manager tables;
// on-chain labels are normalized before comparison so casing/spacing does not affect matching.
func NewOperatorSignerResolver(
	roles core.AccessManagerRoleRepository,
	members core.AccessManagerRoleMemberRepository,
	wallets core.UserWalletRepository,
	log logger.Logger,
) core.OperatorSignerResolver {
	return &operatorSignerResolver{
		roles:   roles,
		members: members,
		wallets: wallets,
		log:     log,
	}
}

// Resolve returns the RaylsAddress of the first active member holding BOTH the PRIVACY_NODE_OPERATOR
// and BANK_EMPLOYEE roles whose UserWallet is HSM-custodied — the two roles that together gate every
// governance write. It fails closed with *core.NoOperatorSignerError rather than signing with an
// unintended key (or one that would revert on-chain) when either role is missing or no member holding
// both has an eligible HSM wallet.
func (r *operatorSignerResolver) Resolve(ctx context.Context) (string, error) {
	operatorRoleID, err := r.roleIDByLabel(ctx, domain.RolePrivacyNodeOperator)
	if err != nil {
		return "", err
	}
	bankRoleID, err := r.roleIDByLabel(ctx, domain.RoleBankEmployee)
	if err != nil {
		return "", err
	}

	account, err := r.members.FindActiveAccountWithAllRoles(ctx, []uint64{operatorRoleID, bankRoleID})
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			r.log.Warn("No active member holds required roles",
				"roles", []string{domain.RolePrivacyNodeOperator, domain.RoleBankEmployee})
			return "", &core.NoOperatorSignerError{}
		}
		return "", fmt.Errorf("find dual-role signer: %w", err)
	}

	wallet, err := r.wallets.FindByRaylsAddress(ctx, account)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			r.log.Warn("Dual-role member has no signable wallet", "account", account)
			return "", &core.NoOperatorSignerError{}
		}
		return "", fmt.Errorf("look up operator wallet %s: %w", account, err)
	}
	if wallet.CustodyProvider != domain.CustodyProviderRaylsHSM {
		r.log.Warn("Operator wallet is not HSM-custodied", "account", account)
		return "", &core.NoOperatorSignerError{} // custody can only sign HSM wallets
	}
	return wallet.RaylsAddress, nil
}

// roleIDByLabel finds the Access Manager role whose label matches the given role. Labels are
// normalized before comparison. Matching is on Label only — the on-chain RoleInfo exposes just the
// label, and the role-gating middleware relies on the same field, so a correctly configured role
// always has its Label set.
func (r *operatorSignerResolver) roleIDByLabel(ctx context.Context, label string) (uint64, error) {
	roles, err := r.roles.List(ctx)
	if err != nil {
		return 0, fmt.Errorf("list access manager roles: %w", err)
	}

	for _, role := range roles {
		if domain.NormalizeRoleLabel(role.Label) == label {
			return role.RoleID, nil
		}
	}

	r.log.Warn("Role not found in access manager", "role", label)
	return 0, &core.NoOperatorSignerError{}
}
