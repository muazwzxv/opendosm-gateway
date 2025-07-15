package repository

import (
	"context"
	"github.com/muazwzxv/opendosm-api/database/model"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/slog"
)

type ItemLookupRepository interface {
	GetByItemCode(ctx context.Context, itemCode string) (*model.ItemLookup, error)
}

type item struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewItemLookup(db *gorm.DB, log *slog.Logger) ItemLookupRepository {
	return &item{
		db:  db,
		log: log,
	}
}

func (r *item) GetByItemCode(ctx context.Context, itemCode string) (*model.ItemLookup, error) {
	var itemModel *model.ItemLookup
	if err := r.db.WithContext(ctx).Where("item_code", itemCode).First(&itemModel).Error; err != nil {
		r.log.ErrorContext(ctx, "Error querying item lookup, error: %v", err)
		return nil, err
	}
	return itemModel, nil
}
