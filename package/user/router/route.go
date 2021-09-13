package router

import (
	// "context"
	// "encoding/json"
	"logit/v1/package/idea"
	user "logit/v1/package/user"
	"logit/v1/util"
	utilResponse "logit/v1/util/response"
	"net/http"

	// utilError "logit/v1/util/error"
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
	grpIdea := g.Group("/idea", m...)
	grpUser.GET("/", h.getById)
	grpUser.GET("/allIdea", h.getById)
	grpUser.PUT("/", h.updateById)
	grpIdea.GET("/", h.getIdeaById)
	grpIdea.POST("/", h.addIdeaById)
	grpUser.PUT("/:id/:mark", h.updateStatusOfIdeaById)
	grpUser.PUT("/:id/extend", h.extendDeadline)
}
func (h *Http) getById(c echo.Context) error {

	user, err := h.uSer.GetUser(util.ToContextService(c))

	if err != nil {
		return utilResponse.RespondError(c, err)
	}

	return utilResponse.Respond(c, user)
}

func (h *Http) getIdeaById(c echo.Context) error {
	id := c.QueryParam("id")
	pass := c.QueryParam("pass")
	idea, err := h.uSer.GetIdea(util.ToContextService(c), id, pass)

	if err != nil {
		return utilResponse.RespondError(c, err)
	}
	return utilResponse.Respond(c, idea)
}

func (h *Http) addIdeaById(c echo.Context) error {
	US := &idea.Idea{}
	if err := c.Bind(US); err != nil {
		return handleError(err, c)
	}
	user, err := h.uSer.AddIdeaByID(util.ToContextService(c), US)
	if err != nil {
		return utilResponse.RespondError(c, err)
	}
	return utilResponse.Respond(c, user)
}

func (h *Http) updateById(c echo.Context) error {
	US := &user.User{}
	if err := c.Bind(US); err != nil {
		return handleError(err, c)
	}
	user, err := h.uSer.UpdateUser(util.ToContextService(c), US)

	if err != nil {
		return utilResponse.RespondError(c, err)
	}
	return utilResponse.Respond(c, user)
}

func (h *Http) updateIdeaById(c echo.Context) error {
	US := &idea.Idea{}
	id := c.QueryParam("id")
	pass := c.QueryParam("pass")
	if err := c.Bind(US); err != nil {
		return handleError(err, c)
	}
	user, err := h.uSer.UpdateIdea(util.ToContextService(c), US, id, pass)

	if err != nil {
		return utilResponse.RespondError(c, err)

	}
	return utilResponse.Respond(c, user)
}

func (h *Http) updateStatusOfIdeaById(c echo.Context) error {
	id := c.Param("id")
	mark := c.Param("mark")
	user, err := h.uSer.UpdateStatus(util.ToContextService(c), id, mark)

	if err != nil {
		return utilResponse.RespondError(c, err)
	}
	return utilResponse.Respond(c, user)
}

func (h *Http) extendDeadline(c echo.Context) error {
	yr := c.QueryParam("yr")
	mon := c.QueryParam("mon")
	hr := c.QueryParam("hr")
	id := c.Param("id")
	user, err := h.uSer.ExtendDeadline(
		util.ToContextService(c),
		id,
		yr,
		mon,
		hr)

	if err != nil {
		return utilResponse.RespondError(c, err)
	}
	return utilResponse.Respond(c, user)
}
func handleError(e error, c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, e)
}
