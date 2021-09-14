package router

import (
	user "askUs/v1/package/user"
	"askUs/v1/util"
	utilResponse "askUs/v1/util/response"

	"github.com/labstack/echo/v4"
)

type Http struct {
	uSer user.Service
}
type Res struct {
	Data interface{}
	Msg  string
}

func Route(g *echo.Group, userService user.Service, m ...echo.MiddlewareFunc) {
	h := &Http{
		uSer: userService,
	}
	grpUser := g.Group("/user", m...)
	grpUser.GET("/", h.getById)
	grpUser.PUT("/", h.updateById)
}
func (h *Http) getById(c echo.Context) error {

	user, err := h.uSer.GetUser(util.ToContextService(c))

	if err != nil {
		return utilResponse.RespondError(c, err)
	}

	return utilResponse.Respond(c, user)
}

func (h *Http) updateById(c echo.Context) error {
	US := &user.User{}
	if err := c.Bind(US); err != nil {
		return utilResponse.RespondError(c, err)
	}
	user, err := h.uSer.UpdateUser(util.ToContextService(c), US)

	if err != nil {
		return utilResponse.RespondError(c, err)
	}
	return utilResponse.Respond(c, user)
}
