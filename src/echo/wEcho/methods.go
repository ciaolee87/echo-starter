package wEcho

import (
	"github.com/ciaolee87/echo-starter/src/echo/wContext"
	"github.com/labstack/echo"
)

func (e *BizEcho) BizGET(path string, handler func(c wContext.Context) error, mid ...echo.MiddlewareFunc) {
	e.GET(path, wContext.CvtHandler(handler), mid...)
}
