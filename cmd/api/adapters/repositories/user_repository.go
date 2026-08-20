package repositories

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/infrastructure/database"
)

var _ core.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

const constraintUsersEmail = "idx_users_email"

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user models.User
	if err := database.TxFromCtx(ctx, r.db).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return user.ToDomain(), nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	record := models.UserFromDomain(user)
	if err := database.TxFromCtx(ctx, r.db).Create(record).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == constraintUsersEmail {
			return core.ErrDuplicateEmail
		}
		return err
	}
	user.ID = record.ID
	user.CreatedAt = record.CreatedAt
	user.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	record := models.UserFromDomain(user)
	if err := database.TxFromCtx(ctx, r.db).Save(record).Error; err != nil {
		return err
	}
	user.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user models.User
	if err := database.TxFromCtx(ctx, r.db).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return user.ToDomain(), nil
}

// SetOnChainUserID persists the keccak256(user.ID) hash on the user row. Idempotent: the hash is
// deterministic, so the WHERE ... IS NULL guard makes repeat calls a no-op.
func (r *UserRepository) SetOnChainUserID(ctx context.Context, userID uuid.UUID, onChainUserID []byte) error {
	return database.TxFromCtx(ctx, r.db).
		Model(&models.User{}).
		Where("id = ? AND on_chain_user_id IS NULL", userID).
		Update("on_chain_user_id", onChainUserID).Error
}

// FindByOnChainUserIDs reverse-maps on-chain keccak256 hashes to ops-api user UUIDs in one query.
// The result is keyed by the lowercase hex encoding of the hash; hashes with no matching user are
// simply absent from the map. An empty input returns an empty map without querying.
func (r *UserRepository) FindByOnChainUserIDs(ctx context.Context, hashes [][]byte) (map[string]uuid.UUID, error) {
	result := make(map[string]uuid.UUID, len(hashes))
	if len(hashes) == 0 {
		return result, nil
	}

	var users []models.User
	if err := database.TxFromCtx(ctx, r.db).
		Select("id", "on_chain_user_id").
		Where("on_chain_user_id IN ?", hashes).
		Find(&users).Error; err != nil {
		return nil, err
	}

	for _, u := range users {
		result[hex.EncodeToString(u.OnChainUserID)] = u.ID
	}
	return result, nil
}

func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := database.TxFromCtx(ctx, r.db).Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
