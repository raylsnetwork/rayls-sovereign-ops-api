package models

import (
	"time"

	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

type TokenBlacklist struct {
	ImmutableModel
	JTI       string    `gorm:"not null;unique"`
	ExpiresAt time.Time `gorm:"not null"`
}

func (TokenBlacklist) TableName() string { return "token_blacklist" }

func (t *TokenBlacklist) ToDomain() *domain.TokenBlacklist {
	return &domain.TokenBlacklist{
		ImmutableModel: t.ToDomainImmutableModel(),
		JTI:            t.JTI,
		ExpiresAt:      t.ExpiresAt,
	}
}

func TokenBlacklistFromDomain(d *domain.TokenBlacklist) *TokenBlacklist {
	return &TokenBlacklist{
		ImmutableModel: ImmutableModelFromDomain(d.ImmutableModel),
		JTI:            d.JTI,
		ExpiresAt:      d.ExpiresAt,
	}
}
