package bizMiddleware

import (
	"github.com/hashicorp/go-uuid"
	"github.com/labstack/echo"
)

var (
	LogHeaderKey = "LogId"
)

// 해더에 아이디값 존재하는지 확인하여 새로 부여한다.
func LogIdMiddleware() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			id := context.Request().Header.Get(LogHeaderKey)
			if id == "" {
				newId, _ := uuid.GenerateUUID()
				context.Request().Header.Set(LogHeaderKey, newId)
			}
			return handlerFunc(context)
		}
	}
}
