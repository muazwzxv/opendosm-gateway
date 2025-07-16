package handler

import (
	"net/http"

	"github.com/muazwzxv/opendosm-api/service"
	"github.com/muazwzxv/opendosm-api/service/item"
	"github.com/muazwzxv/opendosm-api/util"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/slog"
)

type ItemHandler struct {
	goyave.Component
	ItemService item.Service
	log         *slog.Logger
}

func (c *ItemHandler) Init(server *goyave.Server) {
	c.ItemService = server.Service(service.Item).(item.Service)
	c.log = server.Logger
	c.Component.Init(server)
}

func (c *ItemHandler) RegisterRoutes(router *goyave.Router) {
	v1 := router.Subrouter("/v1")

	v1.Get("/item/{itemCode}", c.GetByProductCode).ValidateBody(nil) // TODO: implement body validator
	v1.Get("/item", c.ListItem).ValidateQuery(nil)                   // TODO: implement query validator
}

func (c *ItemHandler) GetByProductCode(resp *goyave.Response, req *goyave.Request) {
	productCode := req.RouteParams["itemCode"]
	if productCode == "" {
		util.HandleError(resp, c.log, util.BadRequest)
		return
	}
	item, err := c.ItemService.GetItem(req.Context(), productCode)
	if err != nil {
		util.HandleError(resp, c.log, err)
		return
	}

	resp.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (c *ItemHandler) ListItem(resp *goyave.Response, req *goyave.Request) {
	// TODO: implement the method
}
