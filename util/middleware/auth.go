package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TokenParser interface {
	ParseToken(string) (*jwt.Token, error)
}

func JwtAuth(tk TokenParser) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := tk.ParseToken(c.Request().Header.Get("Authorization"))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token re login"})
			}
			claims := token.Claims.(jwt.MapClaims)
			id := claims["id"].(float64)
			client := claims["client"].(string)
			c.Set("id", id)
			c.Set("client", client)
			return hf(c)
		}
	}
}
