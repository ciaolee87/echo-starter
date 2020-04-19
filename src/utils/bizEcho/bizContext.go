package bizEcho

import (
	"errors"
	"fmt"
	"github.com/ciaolee87/echo-starter/src/utils/bizLogger"
	"github.com/labstack/echo"
	"time"
)

type BizContext struct {
	echo.Context
	Logger bizLogger.StackLogger
}

func NewBizContext(c *echo.Context) *BizContext {
	context := BizContext{*c, *bizLogger.NewStackLogger()}

	// 로거에 기본 데이터 입력
	now := time.Now()
	context.Logger.Log("time",  fmt.Sprintf("%2d-%02d-%02d", now.Year(), now.Month(), now.Day()))
	context.Logger.Log("ip", context.RealIP())
	context.Logger.Log("path",context.Path())
	context.Logger.Log("",context.)

	return &context
}

// 일정 규칙의 JSON 방식으로 데이터 전송
func (c *BizContext) BizJson(body *BizJSON) error {

	if body == nil {
		body = NewJSON()
	}

	if c.Context == nil {
		return errors.New("context 생성 실패")
	}

	return c.Context.JSON(body.Code, body)
}
