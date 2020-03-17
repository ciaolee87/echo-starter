package wContext

import "github.com/ciaolee87/echo-starter/src/echo/wJSON"

func (c *Context) BizJson(body *wJSON.JSON) error {
	if body == nil {
		body = wJSON.Make(200, "success", nil)
	} else {
		body.FillDefault()
	}

	return c.JSON(*body.Code, body)
}
