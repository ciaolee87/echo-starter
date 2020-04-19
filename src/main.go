package main

import (
	"github.com/ciaolee87/echo-starter/src/utils/bizEcho"
	"github.com/labstack/echo/middleware"
)

var Server *bizEcho.BizEcho

func init() {

}

func main() {
	Server = bizEcho.NewEcho()

	// 각각의 요청에 identifier 입력
	Server.Echo.Use(middleware.RequestID())

	// 중앙 에러 헨들러 작성
	Server.Echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))
	Server.Echo.HTTPErrorHandler = bizEcho.ErrorHandler

}
