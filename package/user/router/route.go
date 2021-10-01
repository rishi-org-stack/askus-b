package router

import (
	user "askUs/v1/package/user"
	"askUs/v1/util"
	"askUs/v1/util/response"

	"github.com/labstack/echo/v4"
)

type Http struct {
	uSer user.Service
}

func Route(g *echo.Group, userService user.Service, m ...echo.MiddlewareFunc) {
	h := &Http{
		uSer: userService,
	}
	grpUser := g.Group("/user", m...)
	grpUser.GET("/", h.getById)
	grpUser.GET("/:name", h.getByName)
	grpUser.PUT("/d", h.updateById)
	grpUser.PUT("/p", h.updateByIdP)
}

func (h *Http) getById(c echo.Context) error {

	user, err := h.uSer.GetUserByID(util.ToContextService(c))

	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}

func (h *Http) getByName(c echo.Context) error {
	name := c.Param("name")
	user, err := h.uSer.GetDoctorByName(util.ToContextService(c), name)

	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}

func (h *Http) updateById(c echo.Context) error {
	US := &user.Doctor{}
	if err := c.Bind(US); err != nil {
		return response.RespondError(c, err)
	}
	user, err := h.uSer.UpdateDoctortByID(util.ToContextService(c), US)

	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}

func (h *Http) updateByIdP(c echo.Context) error {
	US := &user.Patient{}
	if err := c.Bind(US); err != nil {
		return response.RespondError(c, err)
	}
	user, err := h.uSer.UpdatePatientByID(util.ToContextService(c), US)

	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}
