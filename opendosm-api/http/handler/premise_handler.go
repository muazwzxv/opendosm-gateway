package handler

import (
	"net/http"

	"github.com/muazwzxv/opendosm-api/service"
	"github.com/muazwzxv/opendosm-api/service/premise"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/slog"
)

type PremiseHandler struct {
	goyave.Component
	PremiseService premise.Service
	log            *slog.Logger
}

func (c *PremiseHandler) Init(server *goyave.Server) {
	c.PremiseService = server.Service(service.Premise).(premise.Service)
	c.log = server.Logger
	c.Component.Init(server)
}

func (c *PremiseHandler) RegisterRoutes(router *goyave.Router) {
	v1 := router.Subrouter("/v1")
	_ = v1
}

func (c *PremiseHandler) GetByPremiseCode(resp *goyave.Response, req *goyave.Request) {
	resp.JSON(http.StatusOK, map[string]any{
		"items": nil,
	})
}
