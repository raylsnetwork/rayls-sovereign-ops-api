package models

import (
	"time"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

type Nonce struct {
	Model
	WalletAddress string    `gorm:"not null"`
	Nonce         string    `gorm:"not null"`
	Message       string    `gorm:"not null"`
	ExpiresAt     time.Time `gorm:"not null"`
	Used          bool      `gorm:"not null;default:false"`
}

func (n *Nonce) ToDomain() *domain.Nonce {
	return &domain.Nonce{
		Model:         n.ToDomainModel(),
		WalletAddress: n.WalletAddress,
		Nonce:         n.Nonce,
		Message:       n.Message,
		ExpiresAt:     n.ExpiresAt,
		Used:          n.Used,
	}
}

func NonceFromDomain(d *domain.Nonce) *Nonce {
	return &Nonce{
		Model:         ModelFromDomain(d.Model),
		WalletAddress: d.WalletAddress,
		Nonce:         d.Nonce,
		Message:       d.Message,
		ExpiresAt:     d.ExpiresAt,
		Used:          d.Used,
	}
}
