package bizEcho

import (
	"github.com/labstack/echo"
)

// echo.Context 확장
type BizContext struct {
	echo.Context
	BizLogger
}

func NewBizContext(c *echo.Context) *BizContext {
	logId := (*c).Request().Header.Get("LogId")
	context := BizContext{*c, BizLogger{
		LogId: logId,
	}}
	return &context
}

// 일정 규칙의 JSON 방식으로 데이터 전송
func (c *BizContext) BizJson(body *BizJSON) error {
	if body == nil {
		body = NewJSON()
	}

	if body.Code < 1000 {
		return c.JSON(body.Code, body)
	} else {
		return c.JSON(200, body)
	}

}

// requestId 가저오는 메소드
func (c *BizContext) BizRequestID() string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
