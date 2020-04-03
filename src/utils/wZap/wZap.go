package wZap

import "go.uber.org/zap"

func initZap() zap.SugaredLogger {
	logger := zap.NewExample().Sugar()
	return *logger
}
