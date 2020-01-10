package account

import "github.com/labstack/echo"

func Router(e *echo.Echo) {

	e.GET("/user", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return c.String(500, "parsing error")
		}
		return c.JSON(200, u)
	})
}
