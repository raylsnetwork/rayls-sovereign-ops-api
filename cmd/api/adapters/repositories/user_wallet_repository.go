package repositories

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/infrastructure/database"
)

var _ core.UserWalletRepository = (*UserWalletRepository)(nil)

const constraintRaylsAddress = "uq_rayls_address"

type UserWalletRepository struct {
	db *gorm.DB
}

func NewUserWalletRepository(db *gorm.DB) *UserWalletRepository {
	return &UserWalletRepository{db: db}
}

func (r *UserWalletRepository) Create(ctx context.Context, wallet *domain.UserWallet) error {
	record := models.UserWalletFromDomain(wallet)
	if err := database.TxFromCtx(ctx, r.db).Create(record).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == constraintRaylsAddress {
			return core.ErrDuplicateWalletAddress
		}
		return err
	}
	wallet.ID = record.ID
	wallet.CreatedAt = record.CreatedAt
	wallet.UpdatedAt = record.UpdatedAt
	return nil
}

// FindByUserID returns the user's identity/login wallet — the earliest active wallet
// (provider-agnostic). Ordered by created_at so the result is deterministic when a user has
// accumulated multiple active wallets (e.g. via onboarding).
func (r *UserWalletRepository) FindByUserID(ctx context.Context, userID uuid.UUID) (*domain.UserWallet, error) {
	var record models.UserWallet
	if err := database.TxFromCtx(ctx, r.db).
		Where(&models.UserWallet{UserID: userID, IsActive: true}).
		Order("created_at ASC, id ASC").
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}

// GetSignerWalletForChain returns the wallet used for on-chain signing on the given chain: the
// earliest active HSM-custodied wallet. Custody can only sign HSM wallets, so a self-custody wallet
// is never returned here. Ordered by created_at for deterministic resolution.
func (r *UserWalletRepository) GetSignerWalletForChain(
	ctx context.Context,
	userID uuid.UUID,
	chain domain.WalletChain,
) (*domain.UserWallet, error) {
	var record models.UserWallet
	if err := database.TxFromCtx(ctx, r.db).
		Where("user_id = ? AND is_active = ? AND custody_provider = ? AND chain = ?",
			userID, true, domain.CustodyProviderRaylsHSM, chain).
		Order("created_at ASC, id ASC").
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}

// GetSignerWalletByAddress returns the user's active HSM-custodied wallet whose address matches
// (case-insensitive). Custody can only sign HSM wallets, so a self-custody wallet is never returned
// here; the chain rule is left to the caller. Used to verify a caller-supplied signer is one of
// their own signer wallets before signing.
func (r *UserWalletRepository) GetSignerWalletByAddress(
	ctx context.Context,
	userID uuid.UUID,
	address string,
) (*domain.UserWallet, error) {
	var record models.UserWallet
	if err := database.TxFromCtx(ctx, r.db).
		Where("user_id = ? AND is_active = ? AND custody_provider = ? AND LOWER(rayls_address) = ?",
			userID, true, domain.CustodyProviderRaylsHSM, strings.ToLower(address)).
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}

// CompletePending fills in the address a mint produced and flips the intent row active.
// Scoped to rows that are still pending so a retry cannot overwrite a completed wallet.
func (r *UserWalletRepository) CompletePending(
	ctx context.Context,
	id uuid.UUID,
	address, externalID string,
) error {
	result := database.TxFromCtx(ctx, r.db).
		Model(&models.UserWallet{}).
		Where("id = ? AND rayls_address LIKE ?", id, domain.PendingAddressPattern).
		Updates(map[string]any{
			"rayls_address":       address,
			"custody_external_id": externalID,
			"is_active":           true,
		})
	if err := result.Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == constraintRaylsAddress {
			return core.ErrDuplicateWalletAddress
		}
		return err
	}
	if result.RowsAffected == 0 {
		return core.ErrRecordNotFound
	}
	return nil
}

// DeletePending removes an intent row whose mint never happened. Restricted to pending rows:
// a real wallet must never be deleted by this path.
func (r *UserWalletRepository) DeletePending(ctx context.Context, id uuid.UUID) error {
	return database.TxFromCtx(ctx, r.db).
		Where("id = ? AND rayls_address LIKE ?", id, domain.PendingAddressPattern).
		Delete(&models.UserWallet{}).Error
}

// FindPendingByUserID returns unfulfilled mint intents for the user, oldest first.
func (r *UserWalletRepository) FindPendingByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]domain.UserWallet, error) {
	var records []models.UserWallet
	if err := database.TxFromCtx(ctx, r.db).
		Where("user_id = ? AND rayls_address LIKE ?", userID, domain.PendingAddressPattern).
		Order("created_at ASC, id ASC").
		Find(&records).Error; err != nil {
		return nil, err
	}
	wallets := make([]domain.UserWallet, 0, len(records))
	for i := range records {
		wallets = append(wallets, *records[i].ToDomain())
	}
	return wallets, nil
}

func (r *UserWalletRepository) FindByRaylsAddress(ctx context.Context, address string) (*domain.UserWallet, error) {
	var record models.UserWallet
	if err := database.TxFromCtx(ctx, r.db).
		Where("LOWER(rayls_address) = ?", strings.ToLower(address)).
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}
