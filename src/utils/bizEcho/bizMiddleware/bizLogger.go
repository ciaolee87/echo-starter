package bizMiddleware

import (
	"github.com/ciaolee87/echo-starter/src/utils/bizLogger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
	"strings"
)

func NewLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${id}|{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}`,
		Output: newCentralLogOut(),
	})
}

type centralLogOut struct {
	io.Writer
}

func (c *centralLogOut) Write(p []byte) (n int, err error) {
	logStr := string(p)
	splitIndex := strings.Index(logStr, "|")

	id := logStr[:splitIndex]
	logToSave := logStr[splitIndex+1:]

	// 로그를 추가하고 로그 서버로 출력한다.
	bizLogger.Log(id, "echoLog", logToSave)
	bizLogger.Flush(id)
	return len(logToSave), nil
}

func newCentralLogOut() *centralLogOut {
	logOut := centralLogOut{}
	return &logOut
}
