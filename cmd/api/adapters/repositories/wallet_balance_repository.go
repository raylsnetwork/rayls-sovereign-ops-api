package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/adapters/repositories/models"
	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
)

var _ core.WalletBalanceRepository = (*WalletBalanceRepository)(nil)

type WalletBalanceRepository struct {
	db *gorm.DB
}

func NewWalletBalanceRepository(db *gorm.DB) *WalletBalanceRepository {
	return &WalletBalanceRepository{db: db}
}

// Upsert writes a wallet balance row. If a row already exists for
// (wallet_address, token_address), the incoming balance and block_number
// overwrite the stored values only when the incoming block_number is greater
// than or equal to the stored one. This prevents a stale backfill from
// regressing a fresher value pushed by the listener.
func (r *WalletBalanceRepository) Upsert(ctx context.Context, balance *domain.WalletBalance) error {
	record := models.WalletBalanceFromDomain(balance)
	record.WalletAddress = domain.NormalizeAddress(record.WalletAddress)
	record.TokenAddress = domain.NormalizeAddress(record.TokenAddress)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "wallet_address"}, {Name: "token_address"}},
			DoUpdates: []clause.Assignment{
				{
					Column: clause.Column{Name: "balance"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.block_number >= wallet_balances.block_number THEN EXCLUDED.balance ELSE wallet_balances.balance END",
					),
				},
				{
					Column: clause.Column{Name: "block_number"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.block_number >= wallet_balances.block_number THEN EXCLUDED.block_number ELSE wallet_balances.block_number END",
					),
				},
				{
					Column: clause.Column{Name: "updated_at"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.block_number >= wallet_balances.block_number THEN EXCLUDED.updated_at ELSE wallet_balances.updated_at END",
					),
				},
			},
		}).
		Create(record).Error
}

func (r *WalletBalanceRepository) ListByWallet(
	ctx context.Context,
	walletAddress string,
) ([]*domain.WalletBalance, error) {
	addr := domain.NormalizeAddress(walletAddress)
	var records []models.WalletBalance
	if err := r.db.WithContext(ctx).
		Where("wallet_address = ?", addr).
		Order("token_address ASC").
		Find(&records).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.WalletBalance, len(records))
	for i := range records {
		out[i] = records[i].ToDomain()
	}
	return out, nil
}

func (r *WalletBalanceRepository) GetByWalletAndToken(
	ctx context.Context,
	walletAddress, tokenAddress string,
) (*domain.WalletBalance, error) {
	var record models.WalletBalance
	err := r.db.WithContext(ctx).
		Where("wallet_address = ? AND token_address = ?",
			domain.NormalizeAddress(walletAddress),
			domain.NormalizeAddress(tokenAddress),
		).
		First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}
