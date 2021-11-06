package router

import (
	"askUs/v1/package/asset"
	"askUs/v1/package/report"
	"askUs/v1/util"
	"askUs/v1/util/response"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	s report.Service
}

func Router(g *echo.Group, s report.Service, m ...echo.MiddlewareFunc) {
	h := &HTTP{s: s}
	reportGroup := g.Group("/report", m...)
	reportGroup.POST("/add", h.create)
	reportGroup.PUT("/update/:id", h.update)
	reportGroup.DELETE("/delete/:id", h.delete)
}

func (h *HTTP) create(c echo.Context) error {
	file, handler, err := c.Request().FormFile("file")
	header := c.Request().FormValue("header")
	if err != nil {
		return err
	}
	res, err := h.s.Create(util.ToContextService(c),
		&asset.UploadRequest{Reader: file, FileName: handler.Filename}, header)
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *HTTP) update(c echo.Context) error {
	file, handler, err := c.Request().FormFile("file")
	header := c.Request().FormValue("header")
	if err != nil {
		return err
	}

	values := c.ParamValues()
	res, err := h.s.Update(util.ToContextService(c), values[0],
		&asset.UploadRequest{Reader: file, FileName: handler.Filename}, header)
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *HTTP) delete(c echo.Context) error {

	values := c.ParamValues()
	res, err := h.s.Delete(util.ToContextService(c), values[0])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}
