package main

import (
	"fmt"
	bizEcho2 "github.com/ciaolee87/echo-starter/src/utils/bizEcho/bizEcho"
	"github.com/ciaolee87/echo-starter/src/utils/bizEnv"
	"github.com/labstack/echo/middleware"
)

var Server *bizEcho2.BizEcho

func init() {

}

func main() {
	Server = bizEcho2.NewEcho()

	// 각각의 요청에 identifier 입력
	Server.Echo.Use(middleware.RequestID())

	// 로거
	Server.Echo.Use(bizEcho2.NewLoggerMiddleware())

	// 중앙 에러 헨들러 작성
	Server.Echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))

	Server.BizGet("/cool", func(ctx *bizEcho2.BizContext) error {
		return ctx.BizJson(bizEcho2.NewErrorJSON())
	})

	helloGroup := Server.BizGroup("/hello")
	helloGroup.BizGET("/greeting", func(ctx *bizEcho2.BizContext) error {
		return ctx.BizJson(bizEcho2.NewJSON())
	})

	Server.Echo.Start(fmt.Sprintf(":%s", bizEnv.Get("PORT")))
}
