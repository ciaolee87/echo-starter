package main

import (
	"fmt"
	"github.com/ciaolee87/echo-starter/src/utils/bizEcho/bizEcho"
	"github.com/ciaolee87/echo-starter/src/utils/bizEcho/bizMiddleware"
	"github.com/ciaolee87/echo-starter/src/utils/bizEnv"
	"github.com/ciaolee87/echo-starter/src/utils/bizMq"
	"github.com/ciaolee87/echo-starter/src/utils/bizMq/bizMqLogger"
)

var Server *bizEcho.BizEcho

func init() {

}

func main() {
	// 로거에 접속한다
	loggerMq := bizMq.NewConnection(bizEnv.Get("LOGGER_SERVER"))
	bizMqLogger.InitMqLogger(loggerMq, bizEnv.Get("LOGGER_QUEUE"), bizEnv.Get("LOGGER_ID"))

	Server = bizEcho.NewEcho()

	// 모든 요청에 고유 ID 값 등록
	Server.Pre(bizMiddleware.LogIdMiddleware())

	api := Server.BizGroup("/api", bizMiddleware.LogRequestInformation(bizMqLogger.ORDER_STACK))
	api.BizGET("/greeting", func(ctx *bizEcho.BizContext) error {

		ctx.BizLogFlush("안녕 로깅중")

		return ctx.BizJson(bizEcho.NewJSON())
	})

	api.BizGET("/error", func(c *bizEcho.BizContext) error {
		panic(bizEcho.Error2000)
	})

	// Page Not Found
	Server.BizGET("/*", bizEcho.PageNotFoundHandler)

	Server.Start(fmt.Sprintf(":%s", bizEnv.Get("PORT")))
}
