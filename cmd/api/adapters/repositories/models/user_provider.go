package models

import (
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

type UserOAuthProvider struct {
	ImmutableModel
	UserID        uuid.UUID `gorm:"not null"`
	Provider      int       `gorm:"not null"`
	OAuthID       string    `gorm:"not null;column:oauth_id"`
	Email         string
	WalletAddress string `gorm:"not null;default:'';column:wallet_address"`
}

func (UserOAuthProvider) TableName() string { return "user_providers" }

func (u *UserOAuthProvider) ToDomain() *domain.UserOAuthProvider {
	return &domain.UserOAuthProvider{
		ImmutableModel: u.ToDomainImmutableModel(),
		UserID:         u.UserID,
		Provider:       domain.OAuthProvider(u.Provider),
		OAuthID:        u.OAuthID,
		Email:          u.Email,
		WalletAddress:  u.WalletAddress,
	}
}

func UserOAuthProviderFromDomain(d *domain.UserOAuthProvider) *UserOAuthProvider {
	return &UserOAuthProvider{
		ImmutableModel: ImmutableModelFromDomain(d.ImmutableModel),
		UserID:         d.UserID,
		Provider:       int(d.Provider),
		OAuthID:        d.OAuthID,
		Email:          d.Email,
		WalletAddress:  d.WalletAddress,
	}
}
