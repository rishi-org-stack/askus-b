package response

import (
	utilError "askUs/v1/util/error"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Respond(c echo.Context, res *Response) (err error) {
	// contentType := c.Request().Header.Get("Content-Type")

	// switch contentType {
	// case "application/json":
	err = c.JSON(res.Status, res)
	// case "application/xml":
	// 	err = c.XML(res.Status, res)
	// }
	return
}
func RespondError(c echo.Context, res utilError.ApiErrorInterface) (err error) {
	cres := res.(utilError.ApiError)
	err = c.JSON(cres.Status, cres)
	return
}
