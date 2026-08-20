package database

import (
	"context"

	"gorm.io/gorm"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
)

type txKey struct{}

// WithTx stores a transactional *gorm.DB in the context so repositories can pick it up.
func WithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// TxFromCtx returns the transactional *gorm.DB stored by WithTx, or the fallback DB with
// the context applied. Repositories should call this instead of using their db field directly.
func TxFromCtx(ctx context.Context, fallback *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return fallback.WithContext(ctx)
}

var _ core.Transactor = (*GormTransactor)(nil)

// GormTransactor implements core.Transactor using GORM's transaction support.
type GormTransactor struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) *GormTransactor {
	return &GormTransactor{db: db}
}

// WithTransaction runs fn inside a single DB transaction. The context passed to fn carries
// the active *gorm.DB so all repository calls within fn share the same transaction.
func (t *GormTransactor) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(WithTx(ctx, tx))
	})
}
