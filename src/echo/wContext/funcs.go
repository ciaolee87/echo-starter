package wContext

import "github.com/labstack/echo"

func CvtHandler(handler func(c Context) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		bizContext := Context{c}
		return handler(bizContext)
	}
}
