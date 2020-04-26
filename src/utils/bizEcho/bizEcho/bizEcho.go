package bizEcho

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type BizEcho struct {
	Echo *echo.Echo
}

func NewEcho() *BizEcho {
	server := BizEcho{Echo: echo.New()}
	return &server
}

func (e *BizEcho) BizGet(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Echo.Add(http.MethodGet, path, getDefaultHandler(h), m...)
}

func (e *BizEcho) BizPost(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Echo.Add(http.MethodPost, path, getDefaultHandler(h), m...)
}

func (e *BizEcho) BizGroup(prefix string, m ...echo.MiddlewareFunc) *BizGroup {
	group := BizGroup{e.Echo.Group(prefix, m...)}
	return &group
}

func getDefaultHandler(h BizHandleFunc) func(ctx echo.Context) error {
	return func(context echo.Context) error {
		bizContext := NewBizContext(&context)
		// 기초 에러 처리
		defer func() {
			recoverDefault(bizContext)
		}()
		var err error = nil
		err = h(bizContext)

		// 로그서버에 로그 전송
		bizContext.BizLogFlush()
		return err
	}
}

func recoverDefault(ctx *BizContext) {
	if e := recover(); e != nil {
		// 사용자 정의 에러라면 콜백 실행
		if bizError, isBizError := e.(BizError); isBizError {
			// 에러 콜백 실행
			bizError.Callback(ctx)
			ctx.BizLog("Panic", fmt.Sprintf("%s", e))
			ctx.BizLog("bizJSON", fmt.Sprintf("%s", bizError.Body))
			ctx.BizJson(bizError.Body)
			return
		}

		//
		ctx.BizLog("Panic", fmt.Sprintf("%s", e))
		ctx.BizJson(NewErrorJSON())
	}
}
