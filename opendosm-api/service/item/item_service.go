package item

import (
	"context"

	"github.com/muazwzxv/opendosm-api/database/repository"
	"github.com/muazwzxv/opendosm-api/dto"
	"github.com/muazwzxv/opendosm-api/service"
	"goyave.dev/goyave/v5/slog"
)

type Service interface {
	GetItem(ctx context.Context, itemCode string) (*dto.ItemDto, error)
	Name() string
}

type item struct {
	repository repository.ItemLookupRepository
	log        *slog.Logger
}

func NewItemService(repo repository.ItemLookupRepository, log *slog.Logger) Service {
	return &item{
		repository: repo,
		log:        log,
	}
}

func (s *item) Name() string {
	return service.Item
}
