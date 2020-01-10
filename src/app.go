package main

import (
	"./route/account"
	"./utils/env"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

var Echo *echo.Echo

func init() {
	env.LoadEnv()
	Echo = echo.New()
}

func main() {
	Echo.Use(middleware.RequestID())

	account.Router(Echo)
	Echo.Start(":" + os.Getenv("PORT"))
}
