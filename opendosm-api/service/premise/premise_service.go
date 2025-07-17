package premise

import (
	"log/slog"

	"github.com/muazwzxv/opendosm-api/database/repository"
	"github.com/muazwzxv/opendosm-api/service"
)

type Service interface {
}

type premise struct {
	repository repository.PremiseLookupRepository
	log        *slog.Logger
}

func NewPremiseService(repo repository.PremiseLookupRepository, log *slog.Logger) Service {
	return &premise{
		repository: repo,
		log:        log,
	}
}

func (s *premise) Name() string {
	return service.Premise
}
