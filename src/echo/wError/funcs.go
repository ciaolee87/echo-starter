package wError

import (
	"errors"
	"github.com/ciaolee87/echo-starter/src/echo/wContext"
	"github.com/ciaolee87/echo-starter/src/echo/wJSON"
	"github.com/labstack/echo"
)

func ErrorHandler(err error, c echo.Context) {
	bizContext := wContext.Context{c}
	bizError, e := err.(Error)

	// todo 로그 처리 하기
	if e {
		// 사용자가 낸 bizError 에러라면
		bizError.JSON.FillErrorDefault()
		_ = bizContext.BizJson(&bizError.JSON)
	} else {
		// 우발적인 에러라면
		_ = bizContext.BizJson(wJSON.ErrorJson(nil))
	}
}

func MakeError(code int, msg string, body *interface{}, cb *func()) Error {
	return Error{
		error: errors.New("system-error"),
		JSON:  wJSON.JSON{Code: &code, Msg: &msg, Body: body},
		Cb:    cb,
	}
}
