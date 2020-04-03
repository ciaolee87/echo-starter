package main

import (
	"fmt"
	"github.com/ciaolee87/echo-starter/src/echo/wEcho"
	"github.com/ciaolee87/echo-starter/src/echo/wError"
	"github.com/ciaolee87/echo-starter/src/router/account"
	"github.com/ciaolee87/echo-starter/src/utils/env"
	"github.com/ciaolee87/echo-starter/src/utils/wGorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

var Server *wEcho.BizEcho

func init() {
	// env 로딩
	env.LoadEnv()

	// DB 접속
	wGorm.Connect()
}

func main() {

	echoSever := echo.New()
	Server = &wEcho.BizEcho{echoSever}

	// 로거 등록
	Server.Use(middleware.RequestID())

	// 에러핸들러
	Server.HTTPErrorHandler = wError.ErrorHandler

	// 라우터 등록
	account.Router(Server)

	_ = Server.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
