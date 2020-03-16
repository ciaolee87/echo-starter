package main

import (
	"github.com/ciaolee87/echo-starter/src/route/account"
	"github.com/ciaolee87/echo-starter/src/utils/env"
	"github.com/ciaolee87/echo-starter/src/wrapper/context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

var Echo context.BizEcho

func init() {
	env.LoadEnv()
	Echo = context.BizEcho{echo.New()}
}

func main() {
	// 로거 등록
	Echo.Use(middleware.RequestID())

	// 에러핸들러
	Echo.HTTPErrorHandler = context.BizErrorHandler

	// Context 확장
	Echo.Use(context.ExtendContext)

	// 라우터
	account.Router(Echo)

	// 서버 시작
	Echo.Start(":" + os.Getenv("PORT"))
}
