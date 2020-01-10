package error

import "github.com/labstack/echo"

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.String(200, "error")
	}
}

type ErrorJson struct {
	Code int
	Msg  string
}
