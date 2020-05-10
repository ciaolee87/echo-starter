package bizEcho

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type BizEcho struct {
	*echo.Echo
}

func NewEcho() *BizEcho {
	server := BizEcho{echo.New()}
	return &server
}

func (e *BizEcho) BizGET(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Echo.Add(http.MethodGet, path, getDefaultHandler(h), m...)
}

func (e *BizEcho) BizPOST(path string, h BizHandleFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return e.Echo.Add(http.MethodPost, path, getDefaultHandler(h), m...)
}

func (e *BizEcho) BizGroup(prefix string, m ...echo.MiddlewareFunc) *BizGroup {
	group := BizGroup{e.Echo.Group(prefix, m...)}
	return &group
}

func getDefaultHandler(h BizHandleFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		bizContext := NewBizContext(&context)

		// 기초 에러 처리
		defer func() {
			e := recover()
			if e == nil {
				return
			}

			// 사용자 정의 에러라면 콜백 실행
			if bizError, isBizError := e.(BizError); isBizError {
				// 로그생성
				bizContext.BizLogStack(fmt.Sprintf("Panic : %s", bizError.Body.Msg))
				bizContext.BizLogFlush(fmt.Sprintf("Exception : Log End"))

				// 에러 콜백 실행
				bizError.Callback(bizContext)
				bizContext.BizJson(bizError.Body)
			} else {
				// 로그 생성
				bizContext.BizLogStack(fmt.Sprintf("Panic : %s", e))
				bizContext.BizLogFlush(fmt.Sprintf("Exception : Log End"))
				bizContext.BizJson(NewErrorJSON())
			}
		}()

		// 핸들러 실행
		var err error = nil
		err = h(bizContext)
		return err
	}
}
