package domain

type UserStatus int

const (
	UserStatusUnknown               UserStatus = iota // 0 — reserved/unset
	UserStatusWaitingRoleAssignment                   // 1
	UserStatusRoleAssigned                            // 2
)

func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusWaitingRoleAssignment, UserStatusRoleAssigned:
		return true
	}
	return false
}

type User struct {
	Model
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	IsActive bool       `json:"isActive"`
	Status   UserStatus `json:"status"`
	// OnChainUserID is keccak256(user.ID) — the identity the governance contract knows the user by.
	// Persisted at onboarding for reverse-lookup (bytes32 -> UUID); nil for users who never onboarded.
	OnChainUserID []byte `json:"onChainUserId,omitempty"`
}
