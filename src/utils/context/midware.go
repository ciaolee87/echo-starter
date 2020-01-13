package context

import "github.com/labstack/echo"

func ExtendContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bizContext := BizContext{c}
		return next(bizContext)
	}
}
