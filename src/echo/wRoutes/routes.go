package wRoutes

import (
	"github.com/ciaolee87/echo-starter/src/echo/wEcho"
)

type Routes interface {
	Router(e *wEcho.BizEcho)
}
