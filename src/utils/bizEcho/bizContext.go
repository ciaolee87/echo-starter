package bizEcho

import (
	"errors"
	"fmt"
	"github.com/ciaolee87/echo-starter/src/utils/bizLogger"
	"github.com/labstack/echo"
)

type BizContext struct {
	echo.Context
	Logger bizLogger.StackLogger
}

func NewBizContext(c *echo.Context) *BizContext {
	context := BizContext{*c, *bizLogger.NewStackLogger()}
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

// 스트링 배열 합치기
func parseStringArrayToString(body []string) string {
	var str = ""
	for _, s := range body {
		str += s + ","
	}
	return str[len(str)-1:]
}

// 해더 파싱
func parseHeaderToString(body map[string][]string) string {
	var str = "{"
	for s, strings := range body {
		str += fmt.Sprintf("\"%s\" : \"%s\",", s, parseStringArrayToString(strings))
	}
	str = str[len(str)-1:] + "}"
	return str
}
