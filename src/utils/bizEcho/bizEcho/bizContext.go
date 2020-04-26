package bizEcho

import (
	"github.com/ciaolee87/echo-starter/src/utils/bizLogger"
	"github.com/labstack/echo"
	"strings"
)

// echo.Context 확장
type BizContext struct {
	ctx *echo.Context
}

func NewBizContext(c *echo.Context) *BizContext {
	context := BizContext{c}
	return &context
}

// 일정 규칙의 JSON 방식으로 데이터 전송
func (c *BizContext) BizJson(body *BizJSON) error {
	if body == nil {
		body = NewJSON()
	}
	return (*c.ctx).JSON(body.Code, body)
}

// 로거 등록
func (c *BizContext) BizLog(title string, args ...string) {
	requestId := c.BizRequestID()
	bizLogger.Log(requestId, title, strings.Join(args, " "))
}

// 로그 푸시
func (c *BizContext) BizLogFlush() {
	bizLogger.Flush(c.BizRequestID())
}

// requestId 가저오는 메소드
func (c *BizContext) BizRequestID() string {
	return (*c.ctx).Request().Header.Get(echo.HeaderXRequestID)
}

// 즉시 출력되는 로거
func (c *BizContext) BizLineLogger(title string, args ...string) {
	bizLogger.LineLogger(title, strings.Join(args, " "))
}
