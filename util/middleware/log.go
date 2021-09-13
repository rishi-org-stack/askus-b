package middleware

import (
	log "askUs/v1/util/log"

	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reques := map[string]interface{}{
				"method": c.Request().Method,
				"host":   c.Request().Host,
				"path":   c.Request().URL.Path,
				"body":   c.Request().Body,
			}
			log.Log.Info(reques)
			return hf(c)
		}
	}
}
