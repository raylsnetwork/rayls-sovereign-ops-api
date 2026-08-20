package models

import (
	"github.com/google/uuid"

	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

type UserWallet struct {
	Model
	UserID            uuid.UUID `gorm:"not null"`
	RaylsAddress      string    `gorm:"not null"`
	CustodyProvider   int       `gorm:"not null"`
	CustodyExternalID string    `gorm:"not null;column:custody_external_id"`
	Chain             int       `gorm:"not null;default:1"`
	IsActive          bool      `gorm:"not null;default:true"`
}

func (UserWallet) TableName() string { return "user_wallets" }

func (u *UserWallet) ToDomain() *domain.UserWallet {
	return &domain.UserWallet{
		Model:             u.ToDomainModel(),
		UserID:            u.UserID,
		RaylsAddress:      u.RaylsAddress,
		CustodyProvider:   domain.CustodyProviderType(u.CustodyProvider),
		CustodyExternalID: u.CustodyExternalID,
		Chain:             domain.WalletChain(u.Chain),
		IsActive:          u.IsActive,
	}
}

func UserWalletFromDomain(d *domain.UserWallet) *UserWallet {
	return &UserWallet{
		Model:             ModelFromDomain(d.Model),
		UserID:            d.UserID,
		RaylsAddress:      d.RaylsAddress,
		CustodyProvider:   int(d.CustodyProvider),
		CustodyExternalID: d.CustodyExternalID,
		Chain:             int(d.Chain),
		IsActive:          d.IsActive,
	}
}
