package repository

import (
	"context"
	"github.com/muazwzxv/opendosm-api/database/model"
	"gorm.io/gorm"
	"log/slog"
)

type ItemLookupRepository interface {
	GetByItemCode(ctx context.Context, itemCode string) (*model.ItemLookup, error)
}

type item struct {
	DB *gorm.DB
}

func NewItemLookup(db *gorm.DB) ItemLookupRepository {
	return &item{
		DB: db,
	}
}

func (r *item) GetByItemCode(ctx context.Context, itemCode string) (*model.ItemLookup, error) {
	var item *model.ItemLookup
	if err := r.DB.Where("item_code", itemCode).First(&item).Error; err != nil {
		slog.Error("Error: %v", err)
		return nil, err
	}
	return item, nil
}
