package context

import "github.com/labstack/echo"

type BizError struct {
	error
	Code     int
	Msg      string
	Body     interface{}
	Callback func()
}

func BizErrorHandler(err error, c echo.Context) {
	bizContext := c.(BizContext)
	bizError, error := err.(BizError)

	if !error {
		bizContext.BizSendJson(&BizJSON{Code: 200, Msg: "Server Error"})
	} else {
		bizContext.BizSendJson(&BizJSON{Code: bizError.Code, Msg: bizError.Msg, Body: bizError.Body})
	}
}
