package premise

import (
	"context"
	"errors"
	"net/http"

	"github.com/muazwzxv/opendosm-api/dto"
	"github.com/muazwzxv/opendosm-api/util"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/util/typeutil"
)

func (s *premise) GetPremise(ctx context.Context, premiseCode string) (*dto.PremiseDto, error) {
	s.log.InfoContext(ctx, "get premise by premise code", "premiseCode", premiseCode)

	premiseModel, err := s.repository.GetByPremiseCode(ctx, premiseCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util.BuildError(http.StatusNotFound, "NOT_FOUND")
		}
		s.log.ErrorContext(ctx, "error querying premise",
			"premiseCode", premiseCode,
			"error", err)
		return nil, util.BuildError(http.StatusInternalServerError, "INTERNAL_SERVER")
	}

	return typeutil.MustConvert[*dto.PremiseDto](premiseModel), nil
}
