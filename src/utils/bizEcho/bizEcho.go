package bizEcho

import (
	"github.com/labstack/echo"
)

type BizEcho struct {
	Echo *echo.Echo
}

func NewEcho() *BizEcho {
	server := BizEcho{Echo: echo.New()}
	return &server
}

func (e *BizEcho) GET(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Echo.GET(path, func(context echo.Context) error {
		bizContext := NewBizContext(&context)
		return h(bizContext)
	}, m...)
}

func (e *BizEcho) POST(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Echo.POST(path, func(context echo.Context) error {
		bizContext := NewBizContext(&context)
		return h(bizContext)
	}, m...)
}
