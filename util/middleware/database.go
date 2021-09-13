package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ConnectionMDB(client *gorm.DB) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("pgClient", client)
			return hf(c)
		}
	}
}
