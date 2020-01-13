package context

import "github.com/labstack/echo"

type BizEcho struct {
	*echo.Echo
}

func (e *BizEcho) BizGET(path string, handler func(c BizContext) error) {
	e.GET(path, BizHandler(handler))
}
