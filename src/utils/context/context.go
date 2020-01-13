package context

import "github.com/labstack/echo"

type BizContext struct {
	echo.Context
}

type BizJSON struct {
	Code int         `json:code`
	Msg  string      `json:msg`
	Body interface{} `json:body`
}

func (c *BizContext) BizSendJson(body interface{}) error {
	return c.JSON(200, body)
}
