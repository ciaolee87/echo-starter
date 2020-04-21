package bizEcho

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewLoggerMidware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{` +
			`"time" : "${time_rfc3339}"` +
			`"ip":"${remote_ip}",` +
			`"uri":"${uri}",` +
			`"host":"${host}",` +
			`"method":"${method}",` +
			`"path":"${path}"` +
			`}"`,
	})
}
