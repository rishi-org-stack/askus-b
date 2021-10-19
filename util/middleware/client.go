package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ClientCheck() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			for _, route := range routes {
				if strings.Compare(path, route) == 0 {
					client := c.Request().Header.Get("X-Client")
					if client != DoctorClient && client != PatientClient {
						return c.JSON(http.StatusBadRequest, "unknown user")
					}
					c.Set("client", client)
				}
			}

			return hf(c)
		}
	}
}

func ClientTypeCheck(clientType string) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			client := c.Request().Header.Get("X-Client")
			if client != clientType {
				return c.JSON(http.StatusBadRequest, "unauthorized user")
			}
			c.Set("client", client)

			return hf(c)
		}
	}
}

var routes = []string{
	"/api/v1/auth/",
	"/api/v1/auth/verify",
}

const (
	DoctorClient  = "doctor"
	PatientClient = "patient"
)
