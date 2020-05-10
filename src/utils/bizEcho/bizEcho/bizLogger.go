package bizEcho

import "github.com/ciaolee87/echo-starter/src/utils/bizMq/bizMqLogger"

type BizLogger struct {
	LogId string
}

func (bl *BizLogger) BizLogStack(value interface{}) {
	go bizMqLogger.SendLog(bl.LogId, bizMqLogger.ORDER_STACK, value)
}

func (bl *BizLogger) BizLogFlush(value interface{}) {
	go bizMqLogger.SendLog(bl.LogId, bizMqLogger.ORDER_FLUSH, value)
}
