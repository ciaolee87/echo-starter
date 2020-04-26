package bizEcho

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
	log := string(p)
	splitIndex := strings.Index(log, "|")
	bizLogger.Log(log[:splitIndex], "echoLog", log[:splitIndex+1])
	return len(log), nil
}

func newCentralLogOut() *centralLogOut {
	logOut := centralLogOut{}
	return &logOut
}
