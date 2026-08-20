package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/infrastructure/database"
)

var _ core.UserOAuthProviderRepository = (*UserOAuthProviderRepository)(nil)

const constraintOAuthProvider = "uq_oauth_provider"

type UserOAuthProviderRepository struct {
	db *gorm.DB
}

func NewUserOAuthProviderRepository(db *gorm.DB) *UserOAuthProviderRepository {
	return &UserOAuthProviderRepository{db: db}
}

func (r *UserOAuthProviderRepository) Create(ctx context.Context, provider *domain.UserOAuthProvider) error {
	record := models.UserOAuthProviderFromDomain(provider)
	if err := database.TxFromCtx(ctx, r.db).Create(record).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == constraintOAuthProvider {
			return core.ErrDuplicateOAuthLink
		}
		return err
	}
	provider.ID = record.ID
	provider.CreatedAt = record.CreatedAt
	return nil
}

func (r *UserOAuthProviderRepository) FindByProviderAndID(
	ctx context.Context,
	provider domain.OAuthProvider,
	oauthID string,
) (*domain.UserOAuthProvider, error) {
	var record models.UserOAuthProvider
	if err := database.TxFromCtx(ctx, r.db).
		Where(&models.UserOAuthProvider{Provider: int(provider), OAuthID: oauthID}).
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}

func (r *UserOAuthProviderRepository) FindByProviderAndUserID(
	ctx context.Context,
	provider domain.OAuthProvider,
	userID uuid.UUID,
) (*domain.UserOAuthProvider, error) {
	var record models.UserOAuthProvider
	if err := database.TxFromCtx(ctx, r.db).
		Where(&models.UserOAuthProvider{Provider: int(provider), UserID: userID}).
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}

func (r *UserOAuthProviderRepository) UpdateOAuthID(ctx context.Context, id uuid.UUID, oauthID string) error {
	return database.TxFromCtx(ctx, r.db).
		Model(&models.UserOAuthProvider{}).
		Where("id = ?", id).
		Update("oauth_id", oauthID).Error
}

func (r *UserOAuthProviderRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	if err := database.TxFromCtx(ctx, r.db).
		Model(&models.UserOAuthProvider{}).
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
