package server

import (
	"askUs/v1/util/config"

	"github.com/labstack/echo/v4"
)

type (
	Server struct {
		Port string
	}
)

func Init(env *config.Env) *Server {
	return &Server{
		Port: env.Port,
	}
}
func (serv *Server) Start() *echo.Echo {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e = e.Group("/api/v1")
	return e
}
