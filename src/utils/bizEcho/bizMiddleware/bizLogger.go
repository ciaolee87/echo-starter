package bizMiddleware

import (
	"github.com/ciaolee87/echo-starter/src/utils/bizMq/bizMqLogger"
	"github.com/hashicorp/go-uuid"
	"github.com/labstack/echo"
)

type RequestInfo struct {
	Url    string `json:"url"`
	LogId  string `json:"logId"`
	Method string `json:"method"`
	Ip     string `json:"ip"`
}

// 요청자 정보를 로깅
// order
// "00" : STACK LogID 기준으로 메모리에 저장
// "01" : FLUSH 파일로 바로 저장되는 로그
func LogRequestInformation(order string) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			reqInfo := RequestInfo{
				Url:    context.Request().Host,
				LogId:  context.Request().Header.Get("LogId"),
				Method: context.Request().Method,
				Ip:     context.RealIP(),
			}

			if reqInfo.LogId == "" {
				uid, _ := uuid.GenerateUUID()
				reqInfo.LogId = uid
			}

			bizMqLogger.SendLog(reqInfo.LogId, order, reqInfo)
			return handlerFunc(context)
		}
	}
}
