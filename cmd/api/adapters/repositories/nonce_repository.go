package repositories

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

var _ core.NonceRepository = (*NonceRepository)(nil)

type NonceRepository struct {
	db *gorm.DB
}

func NewNonceRepository(db *gorm.DB) *NonceRepository {
	return &NonceRepository{db: db}
}

func (r *NonceRepository) Create(ctx context.Context, nonce *domain.Nonce) error {
	record := models.NonceFromDomain(nonce)
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return err
	}
	nonce.ID = record.ID
	nonce.CreatedAt = record.CreatedAt
	nonce.UpdatedAt = record.UpdatedAt
	return nil
}

// FindValidAndMarkUsed atomically finds a valid nonce and marks it as used within a single
// transaction using SELECT FOR UPDATE. This prevents race conditions where two concurrent
// requests could both validate the same nonce before either marks it used.
func (r *NonceRepository) FindValidAndMarkUsed(ctx context.Context, addr, nonce string) (*domain.Nonce, error) {
	var n models.Nonce

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("wallet_address = ? AND nonce = ? AND used = false AND expires_at > ?", addr, nonce, time.Now().UTC()).
			First(&n).Error; err != nil {
			return err
		}

		return tx.Model(&n).Update("used", true).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}

	return n.ToDomain(), nil
}

func (r *NonceRepository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at < ?", time.Now().UTC()).Delete(&models.Nonce{}).Error
}
