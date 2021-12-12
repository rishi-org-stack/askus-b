package router

import (
	"askUs/v1/package/advice"
	"askUs/v1/util"
	"askUs/v1/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

//TODO
//1. /advice/g post to create a advice global
//2. /advice/:patientID post to send a advice to a particular patient in cnnections
//3. /advice/ Get to get all global advice
//4. /advice/g/:adviceID GET to get a global advice
//5./advice/:adviceID GET patient if global get tht advice if personnel see if patient is or not
// 6. /advice/:adviceID Get doctor if advice  is personnel then if given by is curnt doctor
//return advice
//7. /advice/ GET doctor get al advice of doctor and advices to patient
//8. /advice/ Get patient  get all global advice and advices frm doctor
//9. /advice/:adviceID PUT if advce is global create llike onject

type HTTP struct {
	svc advice.Service
}

func Route(g *echo.Group, serv advice.Service, m ...echo.MiddlewareFunc) {
	h := &HTTP{
		svc: serv,
	}
	adviceGrp := g.Group("/advice", m...)
	//TODO: i need to have a / page which should not just spit out all global advices rather than with some sense
	//Doc
	adviceGrp.POST("/g", h.create)
	adviceGrp.POST("/:id", h.createPersonel)
	adviceGrp.GET("/p/my", h.getPersonelAdvicePostedByMe)
	adviceGrp.GET("/p/:id", h.getPatientAdviceGrp)
	adviceGrp.GET("/g/my", h.getPersonels)
	//Patient
	adviceGrp.GET("/p/forme", h.getPersonels)
	//common
	adviceGrp.GET("/", h.getGlobals)
	adviceGrp.GET("/g/:adviceID", h.getGlobal)
	adviceGrp.GET("/:adviceID", h.getPersonelAdvice)
	adviceGrp.GET("/g/:adviceID/like", h.likeGlobal)
	adviceGrp.GET("/all", h.getAll)
	// adviceGrp.GET("/:id", h.ok)
	//TODO: we need a route for doc to get all his advices
	//Mostly: Personnels

}

func (h *HTTP) create(c echo.Context) error {
	ad := &advice.Advice{}
	if err := c.Bind(ad); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := h.svc.CreateAdvice(util.ToContextService(c), ad)
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *HTTP) createPersonel(c echo.Context) error {
	ad := &advice.Advice{}
	if err := c.Bind(ad); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := h.svc.CreatePersonelAdvice(util.ToContextService(c), ad, c.ParamValues()[0])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *HTTP) getGlobals(c echo.Context) error {
	return c.String(200, "ok")
}

func (h *HTTP) getGlobal(c echo.Context) error {
	res, err := h.svc.GetGlobalAdvice(util.ToContextService(c), c.ParamValues()[0])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}
func (h *HTTP) getPersonelAdvice(c echo.Context) error {
	res, err := h.svc.GetPersonelAdvice(util.ToContextService(c), c.ParamValues()[0])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *HTTP) getPersonels(c echo.Context) error {
	res, err := h.svc.GetPersonelAdvices(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *HTTP) getPersonelAdvicePostedByMe(c echo.Context) error {
	res, err := h.svc.GetPersonelAdvicesPostedByMe(util.ToContextService(c))
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

func (h *HTTP) getPatientAdviceGrp(c echo.Context) error {
	res, err := h.svc.GetPatientAndMyAdvices(util.ToContextService(c), c.ParamValues()[0])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
	// return c.String(200, "ok")
}
func (h *HTTP) getAll(c echo.Context) error {
	return c.String(200, "ok")
}

func (h *HTTP) likeGlobal(c echo.Context) error {
	// return c.String(200, "ok")
	res, err := h.svc.LikeAdvice(util.ToContextService(c), c.ParamValues()[0])
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, res)
}

// func (h *HTTP) ok(c echo.Context) error {
// 	return c.String(http.StatusAccepted, strings.Join(c.ParamValues(), ""))
// }
