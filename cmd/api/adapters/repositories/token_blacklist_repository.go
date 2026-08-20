package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

var _ core.TokenBlacklistRepository = (*TokenBlacklistRepository)(nil)

type TokenBlacklistRepository struct {
	db *gorm.DB
}

func NewTokenBlacklistRepository(db *gorm.DB) *TokenBlacklistRepository {
	return &TokenBlacklistRepository{db: db}
}

func (r *TokenBlacklistRepository) Create(ctx context.Context, entry *domain.TokenBlacklist) error {
	record := models.TokenBlacklistFromDomain(entry)
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return err
	}
	entry.ID = record.ID
	entry.CreatedAt = record.CreatedAt
	return nil
}

func (r *TokenBlacklistRepository) ExistsByJTI(ctx context.Context, jti string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.TokenBlacklist{}).
		Where("jti = ?", jti).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *TokenBlacklistRepository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expires_at < ?", time.Now().UTC()).
		Delete(&models.TokenBlacklist{}).Error
}
