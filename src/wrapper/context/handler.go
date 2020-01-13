package context

import "github.com/labstack/echo"

func BizHandler(handler func(c BizContext) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		bizContext := BizContext{c}
		return handler(bizContext)
	}
}
