package item

import (
	"context"
	"errors"
	"github.com/muazwzxv/opendosm-api/dto"
	"github.com/muazwzxv/opendosm-api/util"
	"gorm.io/gorm"
	"net/http"
)

func (s *item) GetItem(ctx context.Context, itemCode string) (*dto.ItemDto, error) {
	s.log.InfoContext(ctx, "get item by item code: %s", itemCode)

	itemModel, err := s.repository.GetByItemCode(ctx, itemCode)
	_ = itemModel
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util.BuildError(http.StatusNotFound, "NOT_FOUND")
		}
		return nil, util.BuildError(http.StatusInternalServerError, "INTERNAL_SERVER")
	}

	// TODO: convert to dto before returning
	return nil, nil
}
