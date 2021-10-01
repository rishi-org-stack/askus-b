package router

import (
	"askUs/v1/package/advice"

	"github.com/labstack/echo/v4"
)

//TODO
//1. /advice/g post to create a advice global
//2. /advice/:patientID post to send a advice to a particular patient in cnnections
//3. /advice/ Get to get all global advice
//4. /advice/:adviceID GET to get a global advice
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
	adviceGrp := g.Group("/advice")
	adviceGrp.GET("/", h.ok)
}

func (h *HTTP) ok(c echo.Context) error {
	return c.String(200, "ok")
}
