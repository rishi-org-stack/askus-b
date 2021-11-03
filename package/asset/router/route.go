package router

import (
	"askUs/v1/package/asset"
	"askUs/v1/util"
	"askUs/v1/util/response"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	assetS asset.Service
}

func Route(g *echo.Group, assetS asset.Service, m ...echo.MiddlewareFunc) {
	h := &HTTP{
		assetS: assetS,
	}
	assetGrp := g.Group("/asset", m...)
	assetGrp.POST("/up", h.upload)
	assetGrp.GET("/down/:kind/:file", h.download)
}

func (h *HTTP) upload(c echo.Context) error {
	file, handler, err := c.Request().FormFile("file")
	kind := c.Request().FormValue("kind")
	if err != nil {
		return err
	}
	res, err := h.assetS.Upload(util.ToContextService(c),
		&asset.UploadRequest{Reader: file, Kind: kind, FileName: handler.Filename})
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
	// return c.JSON(http.StatusAccepted, res)
}

func (h *HTTP) download(c echo.Context) error {
	values := c.ParamValues()
	kind := values[0]
	file := values[1]
	url := kind + "/" + file
	_, err := h.assetS.Download(
		util.ToContextService(c),
		url,
		c.Response().Writer)
	if err != nil {
		return response.RespondError(c, err)
	}

	return nil
	// return c.JSON(http.StatusAccepted, res)
}
