package server

import (
	"fmt"
	"github.com/ciaolee87/echo-starter/src/echo/wContext"
	"github.com/ciaolee87/echo-starter/src/echo/wEcho"
	"github.com/ciaolee87/echo-starter/src/echo/wError"
	"github.com/ciaolee87/echo-starter/src/utils/env"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Server *wEcho.BizEcho

func init() {
	env.LoadEnv()
	echoSever := echo.New()
	Server = &wEcho.BizEcho{echoSever}

	// 로거 등록
	Server.Use(middleware.RequestID())

	// 에러핸들러
	Server.HTTPErrorHandler = wError.ErrorHandler

	// Context 확장
	Server.Use(wContext.ExtendContext)
}

func Run(port int) {
	_ = Server.Start(fmt.Sprintf(":%d", port))
}
