package logger

import (
	"context"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func Logger() echo.MiddlewareFunc {
	logger := logrus.New()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logEntry := logrus.NewEntry(logger)

			// 각 요청의 고유 requestId 가저오기
			id := c.Request().Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = c.Response().Header().Get(echo.HeaderXRequestID)
			}

			logEntry = logEntry.WithField("RequestID", id)
			req := c.Request()
			c.SetRequest(req.WithContext(
				context.WithValue(
					req.Context(),
					"LOG",
					logEntry,
				),
			))

			return next(c)
		}
	}
}
