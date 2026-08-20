package models

import "github.com/raylsnetwork/rayls-sovereign-ops-api/domain"

type WalletBalance struct {
	Model
	WalletAddress string `gorm:"size:42;not null"`
	TokenAddress  string `gorm:"size:42;not null"`
	Balance       string `gorm:"type:text;not null"`
	BlockNumber   int64  `gorm:"not null;default:0"`
}

func (WalletBalance) TableName() string { return "wallet_balances" }

func (w *WalletBalance) ToDomain() *domain.WalletBalance {
	return &domain.WalletBalance{
		Model:         w.ToDomainModel(),
		WalletAddress: w.WalletAddress,
		TokenAddress:  w.TokenAddress,
		Balance:       w.Balance,
		BlockNumber:   uint64(w.BlockNumber),
	}
}

func WalletBalanceFromDomain(d *domain.WalletBalance) *WalletBalance {
	return &WalletBalance{
		Model:         ModelFromDomain(d.Model),
		WalletAddress: d.WalletAddress,
		TokenAddress:  d.TokenAddress,
		Balance:       d.Balance,
		BlockNumber:   int64(d.BlockNumber),
	}
}
