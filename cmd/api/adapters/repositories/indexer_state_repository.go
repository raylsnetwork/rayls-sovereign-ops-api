package repositories

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
)

var _ core.IndexerStateRepository = (*IndexerStateRepository)(nil)

type indexerStateRecord struct {
	Key       string `gorm:"primaryKey"`
	Value     string
	UpdatedAt time.Time
}

func (indexerStateRecord) TableName() string { return "indexer_state" }

type IndexerStateRepository struct{ db *gorm.DB }

func NewIndexerStateRepository(db *gorm.DB) *IndexerStateRepository {
	return &IndexerStateRepository{db: db}
}

func (r *IndexerStateRepository) Get(ctx context.Context, key string) (string, error) {
	var rec indexerStateRecord
	if err := r.db.WithContext(ctx).First(&rec, "key = ?", key).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", core.ErrRecordNotFound
		}
		return "", err
	}
	return rec.Value, nil
}

func (r *IndexerStateRepository) Set(ctx context.Context, key, value string) error {
	rec := indexerStateRecord{Key: key, Value: value, UpdatedAt: time.Now().UTC()}
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
		}).
		Create(&rec).Error
}
