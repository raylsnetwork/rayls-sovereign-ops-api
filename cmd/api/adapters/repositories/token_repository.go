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

var _ core.TokenRepository = (*TokenRepository)(nil)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Upsert(ctx context.Context, token *domain.Token) error {
	record := models.TokenFromDomain(token)
	record.ContractAddress = domain.NormalizeAddress(record.ContractAddress)
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "contract_address"}},
			// erc_standard, issuer_id, total_supply and holder_count are guarded so neither write
			// path clobbers a value the other owns, regardless of which wins the race.
			//
			// The deploy path is authoritative about a token's standard; the indexer only knows
			// what Blockscout reports, which transiently mislabels a freshly deployed stablecoin
			// (erc_standard 7) as ERC-20 (1) before its Rayls-specific discovery corrects it.
			// A deployed token always carries a non-empty issuer_id (the indexer leaves it empty),
			// so we keep the existing erc_standard whenever the row was deploy-recorded. We also
			// keep it when the incoming value is Custom=0 (an indexer discovery that lacks a type).
			//
			// total_supply and holder_count are owned by the indexer. The deploy path records a
			// token with empty supply and a 0 holder count, so if the deploy write lands after the
			// indexer has already populated those fields it must not reset them — we keep the
			// existing value whenever the incoming one is empty/zero.
			DoUpdates: append(
				clause.AssignmentColumns([]string{"name", "symbol", "updated_at"}),
				clause.Assignment{
					Column: clause.Column{Name: "erc_standard"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.erc_standard = 0 OR (tokens.issuer_id IS NOT NULL AND tokens.issuer_id != '') " +
							"THEN tokens.erc_standard ELSE EXCLUDED.erc_standard END",
					),
				},
				clause.Assignment{
					Column: clause.Column{Name: "issuer_id"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.issuer_id = '' OR EXCLUDED.issuer_id IS NULL THEN tokens.issuer_id ELSE EXCLUDED.issuer_id END",
					),
				},
				clause.Assignment{
					Column: clause.Column{Name: "total_supply"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.total_supply = '' OR EXCLUDED.total_supply IS NULL THEN tokens.total_supply ELSE EXCLUDED.total_supply END",
					),
				},
				clause.Assignment{
					// decimals is owned by whichever path knows it. The indexer sends 0 until
					// Blockscout catalogs the token, and the deploy path records the real value,
					// so keep the existing decimals whenever the incoming one is 0 (unknown).
					Column: clause.Column{Name: "decimals"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.decimals = 0 THEN tokens.decimals ELSE EXCLUDED.decimals END",
					),
				},
				clause.Assignment{
					Column: clause.Column{Name: "holder_count"},
					Value: gorm.Expr(
						"CASE WHEN EXCLUDED.holder_count = 0 THEN tokens.holder_count ELSE EXCLUDED.holder_count END",
					),
				},
			),
		}).
		Create(record).Error
}

func (r *TokenRepository) UpdateSupplyAndHolders(
	ctx context.Context,
	address, totalSupply string,
	holderCount int,
) error {
	return r.db.WithContext(ctx).
		Model(&models.Token{}).
		Where("contract_address = ?", domain.NormalizeAddress(address)).
		Updates(map[string]any{
			"total_supply": totalSupply,
			"holder_count": holderCount,
			"updated_at":   time.Now().UTC(),
		}).Error
}

func (r *TokenRepository) FindByAddress(ctx context.Context, address string) (*domain.Token, error) {
	var record models.Token
	if err := r.db.WithContext(ctx).
		Where("contract_address = ?", domain.NormalizeAddress(address)).
		First(&record).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, err
	}
	return record.ToDomain(), nil
}

func ercStandardFromString(t string) domain.ErcStandard {
	return domain.ParseErcStandard(t)
}

func (r *TokenRepository) List(ctx context.Context, filter core.TokenFilter) ([]*domain.Token, int64, error) {
	q := r.db.WithContext(ctx).Model(&models.Token{})
	if filter.TokenClass != "" {
		q = q.Where("token_class = ?", filter.TokenClass)
	}
	if filter.Status != nil {
		q = q.Where("status = ?", *filter.Status)
	}
	if filter.Name != "" {
		q = q.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Symbol != "" {
		q = q.Where("symbol ILIKE ?", "%"+filter.Symbol+"%")
	}
	if filter.TokenType != "" {
		q = q.Where("erc_standard = ?", ercStandardFromString(filter.TokenType))
	}
	if filter.IssuerID != "" {
		q = q.Where("issuer_id = ?", filter.IssuerID)
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	limit := filter.Limit
	if limit < 1 {
		limit = 20
	}

	var records []models.Token
	if err := q.Offset((page - 1) * limit).Limit(limit).Order("created_at DESC").Find(&records).Error; err != nil {
		return nil, 0, err
	}

	tokens := make([]*domain.Token, len(records))
	for i := range records {
		tokens[i] = records[i].ToDomain()
	}
	return tokens, count, nil
}
