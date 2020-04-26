package bizEcho

import (
	"github.com/labstack/echo"
	"net/http"
)

type BizGroup struct {
	*echo.Group
}

func (b *BizGroup) BizGET(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return b.Add(http.MethodGet, path, getDefaultHandler(h), m...)
}

func (b *BizGroup) BizPOST(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return b.Add(http.MethodPost, path, getDefaultHandler(h), m...)
}
