package models

import "github.com/raylsnetwork/rayls-sovereign-ops-api/domain"

type User struct {
	Model
	Name     string `gorm:"not null;default:''"`
	Email    string `gorm:"not null;default:''"`
	IsActive bool   `gorm:"not null;default:true"`
	Status   int    `gorm:"not null;default:1"`
	// Uniqueness is enforced by the partial index idx_users_on_chain_user_id (migration 000007), not a
	// GORM unique tag — a plain unique constraint would reject multiple NULLs (non-onboarded users).
	OnChainUserID []byte `gorm:"column:on_chain_user_id;type:bytea"`
}

func (User) TableName() string { return "users" }

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		Model:         u.ToDomainModel(),
		Name:          u.Name,
		Email:         u.Email,
		IsActive:      u.IsActive,
		Status:        domain.UserStatus(u.Status),
		OnChainUserID: u.OnChainUserID,
	}
}

func UserFromDomain(d *domain.User) *User {
	return &User{
		Model:         ModelFromDomain(d.Model),
		Name:          d.Name,
		Email:         d.Email,
		IsActive:      d.IsActive,
		Status:        int(d.Status),
		OnChainUserID: d.OnChainUserID,
	}
}
