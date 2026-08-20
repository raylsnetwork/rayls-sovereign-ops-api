package testutil

import (
	"context"
	"sort"
	"strings"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

var _ core.AccessManagerRoleMemberRepository = (*FakeAccessManagerRoleMemberRepository)(nil)

// FakeAccessManagerRoleMemberRepository is an in-memory AccessManagerRoleMemberRepository for unit tests.
type FakeAccessManagerRoleMemberRepository struct {
	Members        []*domain.AccessManagerRoleMember
	ListByRoleErr  error // if set, returned by ListByRole instead of searching
	FindAccountErr error // if set, returned by FindActiveAccountWithAllRoles instead of searching
}

func (r *FakeAccessManagerRoleMemberRepository) Upsert(
	_ context.Context,
	member *domain.AccessManagerRoleMember,
) error {
	for i := range r.Members {
		if r.Members[i].RoleID == member.RoleID && strings.EqualFold(r.Members[i].Account, member.Account) {
			r.Members[i] = member
			return nil
		}
	}
	r.Members = append(r.Members, member)
	return nil
}

func (r *FakeAccessManagerRoleMemberRepository) Revoke(_ context.Context, roleID uint64, account string) error {
	for i := range r.Members {
		if r.Members[i].RoleID == roleID && strings.EqualFold(r.Members[i].Account, account) {
			r.Members[i].IsActive = false
		}
	}
	return nil
}

func (r *FakeAccessManagerRoleMemberRepository) ListByRole(
	_ context.Context,
	roleID uint64,
	activeOnly bool,
) ([]*domain.AccessManagerRoleMember, error) {
	if r.ListByRoleErr != nil {
		return nil, r.ListByRoleErr
	}
	var out []*domain.AccessManagerRoleMember
	for i := range r.Members {
		if r.Members[i].RoleID != roleID {
			continue
		}
		if activeOnly && !r.Members[i].IsActive {
			continue
		}
		out = append(out, r.Members[i])
	}
	return out, nil
}

func (r *FakeAccessManagerRoleMemberRepository) FindActiveAccountWithAllRoles(
	_ context.Context,
	roleIDs []uint64,
) (string, error) {
	if r.FindAccountErr != nil {
		return "", r.FindAccountErr
	}
	if len(roleIDs) == 0 {
		return "", core.ErrRecordNotFound
	}

	wanted := make(map[uint64]struct{}, len(roleIDs))
	for _, id := range roleIDs {
		wanted[id] = struct{}{}
	}

	// Tally the distinct wanted roles held by each active account (case-insensitive account key).
	rolesByAccount := make(map[string]map[uint64]struct{})
	displayAccount := make(map[string]string)
	for _, m := range r.Members {
		if !m.IsActive {
			continue
		}
		if _, ok := wanted[m.RoleID]; !ok {
			continue
		}
		key := strings.ToLower(m.Account)
		if rolesByAccount[key] == nil {
			rolesByAccount[key] = make(map[uint64]struct{})
			displayAccount[key] = m.Account
		}
		rolesByAccount[key][m.RoleID] = struct{}{}
	}

	var matches []string
	for key, held := range rolesByAccount {
		if len(held) == len(wanted) {
			matches = append(matches, displayAccount[key])
		}
	}
	if len(matches) == 0 {
		return "", core.ErrRecordNotFound
	}
	sort.Strings(matches)
	return matches[0], nil
}

func (r *FakeAccessManagerRoleMemberRepository) ListByAccount(
	_ context.Context,
	account string,
) ([]*domain.AccessManagerRoleMember, error) {
	var out []*domain.AccessManagerRoleMember
	for i := range r.Members {
		if strings.EqualFold(r.Members[i].Account, account) {
			out = append(out, r.Members[i])
		}
	}
	return out, nil
}
