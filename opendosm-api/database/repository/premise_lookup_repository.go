package repository

import (
	"context"
	"github.com/muazwzxv/opendosm-api/database/model"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/slog"
)

type PremiseLookupRepository interface {
	GetByPremiseCode(ctx context.Context, premiseCode string) (*model.PremiseLookup, error)
}

type premise struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewPremiseLookup(db *gorm.DB, log *slog.Logger) PremiseLookupRepository {
	return &premise{
		db:  db,
		log: log,
	}
}

func (r *premise) GetByPremiseCode(ctx context.Context, premiseCode string) (*model.PremiseLookup, error) {
	var premiseModel *model.PremiseLookup
	if err := r.db.WithContext(ctx).Where("premise_code", premiseCode).First(&premiseModel).Error; err != nil {
		r.log.ErrorContext(ctx, "Error querying item lookup, error: %v", err)
		return nil, err
	}
	return premiseModel, nil
}
