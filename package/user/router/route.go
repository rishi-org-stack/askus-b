package router

import (
	user "askUs/v1/package/user"
	"askUs/v1/util"
	"askUs/v1/util/middleware"
	"askUs/v1/util/response"
	"log"

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
	grpUser.PUT("/d", h.updateById, middleware.ClientTypeCheck(user.DoctorClient))
	grpUser.PUT("/p", h.updateByIdP, middleware.ClientTypeCheck(user.PatientClient))
	grpUser.GET("/fbd", h.followedByDoctors, middleware.ClientTypeCheck(user.DoctorClient))
	grpUser.GET("/fd", h.followingDoctors, middleware.ClientTypeCheck(user.DoctorClient))
	grpUser.GET("/fbp", h.followedByPatients, middleware.ClientTypeCheck(user.DoctorClient))
	grpUser.GET("/fdbp", h.followingDoctorsByPatients, middleware.ClientTypeCheck(user.PatientClient))
	grpReq := g.Group("/req", m...)
	grpReq.POST("/:userID", h.createReq)
	grpReq.GET("/my", h.getMyReqs)
	grpReq.GET("/forme", h.getReqs)                  //, middleware.ClientTypeCheck(user.DoctorClient))
	grpReq.PUT("/:reqID/:status", h.updateStatusReq) //, middleware.ClientTypeCheck(user.DoctorClient))

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
	log.Println(US)
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

func (h *Http) createReq(c echo.Context) error {
	req, err := h.uSer.CreateReq(util.ToContextService(c), c.ParamValues()[0])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, req)
}

func (h *Http) updateStatusReq(c echo.Context) error {
	req, err := h.uSer.UpdateStatusOfReq(util.ToContextService(c), c.ParamValues()[0], c.ParamValues()[1])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, req)
}

func (h *Http) getMyReqs(c echo.Context) error {
	res, err := h.uSer.GetMyRequests(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *Http) getReqs(c echo.Context) error {
	res, err := h.uSer.GetRequestForMe(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *Http) followedByDoctors(c echo.Context) error {
	res, err := h.uSer.GetFBD(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}
func (h *Http) followedByPatients(c echo.Context) error {
	res, err := h.uSer.GetFBP(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *Http) followingDoctors(c echo.Context) error {
	res, err := h.uSer.GetFD(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *Http) followingDoctorsByPatients(c echo.Context) error {
	res, err := h.uSer.GetFDBP(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}
