package main

import (
	"fmt"
	"github.com/ciaolee87/echo-starter/src/utils/bizEcho/bizEcho"
	"github.com/ciaolee87/echo-starter/src/utils/bizEcho/bizMiddleware"
	"github.com/ciaolee87/echo-starter/src/utils/bizEnv"
	"github.com/ciaolee87/echo-starter/src/utils/bizRabbitMq/bizMqLogSender"
)

var Server *bizEcho.BizEcho

func init() {

}

func main() {
	// 로거에 접속한다
	bizMqLogSender.InitLogger(
		bizEnv.Get(""),
		bizEnv.Get(""),
		bizEnv.Get(""),
	)

	Server = bizEcho.NewEcho()

	// 모든 요청에 고유 ID 값 등록
	Server.Pre(bizMiddleware.LogIdMiddleware())

	api := Server.BizGroup("/api")
	api.BizGET("/greeting", func(ctx *bizEcho.BizContext) error {

		ctx.BizLog("안녕하세요", "로거 작동 잘해요!!")

		return ctx.BizJson(bizEcho.NewJSON())
	})

	// Page Not Found
	api.BizGET("/*", bizEcho.PageNotFoundHandler)

	// Page Not Found
	Server.BizGET("/*", bizEcho.PageNotFoundHandler)

	Server.Start(fmt.Sprintf(":%s", bizEnv.Get("PORT")))
}
