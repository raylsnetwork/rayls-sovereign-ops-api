package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/infrastructure/database"
)

var _ core.UserSignupDetailsRepository = (*UserSignupDetailsRepository)(nil)

type UserSignupDetailsRepository struct {
	db *gorm.DB
}

func NewUserSignupDetailsRepository(db *gorm.DB) *UserSignupDetailsRepository {
	return &UserSignupDetailsRepository{db: db}
}

// Upsert stores the sign-up answers for a user, overwriting any previous submission.
// One row per user (uq on user_id), so signing up again with the same address refreshes
// the details instead of accumulating rows.
func (r *UserSignupDetailsRepository) Upsert(ctx context.Context, details *domain.UserSignupDetails) error {
	record := models.UserSignupDetailsFromDomain(details)
	if err := database.TxFromCtx(ctx, r.db).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns(
				[]string{"company", "employees", "heard_about", "goals", "updated_at"},
			),
		}).
		Create(record).Error; err != nil {
		return err
	}
	details.ID = record.ID
	details.CreatedAt = record.CreatedAt
	details.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *UserSignupDetailsRepository) FindByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*domain.UserSignupDetails, error) {
	var record models.UserSignupDetails
	if err := database.TxFromCtx(ctx, r.db).
		Where("user_id = ?", userID).
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}
