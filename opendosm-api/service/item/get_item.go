package item

import (
	"context"
	"errors"
	"net/http"

	"github.com/muazwzxv/opendosm-api/dto"
	"github.com/muazwzxv/opendosm-api/util"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/util/typeutil"
)

func (s *item) GetItem(ctx context.Context, itemCode string) (*dto.ItemDto, error) {
	s.log.InfoContext(ctx, "get item by item code", "itemCode", itemCode)

	itemModel, err := s.repository.GetByItemCode(ctx, itemCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util.BuildError(http.StatusNotFound, "NOT_FOUND")
		}
		return nil, util.BuildError(http.StatusInternalServerError, "INTERNAL_SERVER")
	}

	return typeutil.MustConvert[*dto.ItemDto](itemModel), nil
}
