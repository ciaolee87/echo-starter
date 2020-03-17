package account

import (
	"github.com/ciaolee87/echo-starter/src/echo/wContext"
	"github.com/ciaolee87/echo-starter/src/echo/wEcho"
	"github.com/ciaolee87/echo-starter/src/echo/wJSON"
)

func Router(e wEcho.BizEcho) {
	e.BizGET("/user", func(c wContext.Context) error {
		return c.BizJson(wJSON.DefaultJson("Hello world"))
	})
}
