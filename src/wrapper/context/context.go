package context

import "github.com/labstack/echo"

type BizContext struct {
	echo.Context
}

type BizJSON struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body"`
}

func (c *BizContext) BizSendJson(bizJson *BizJSON) error {

	var resCode = 200
	switch {
	case bizJson.Code == 0:
		bizJson.Code = 200
	case 0 < bizJson.Code && bizJson.Code < 1000:
		resCode = bizJson.Code
	}

	if bizJson.Msg == "" {
		bizJson.Msg = "success"
	}

	return c.JSON(resCode, bizJson)
}
