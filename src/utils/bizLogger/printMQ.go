package bizLogger

import (
	"github.com/ciaolee87/echo-starter/src/utils/bizEnv"
	"github.com/ciaolee87/echo-starter/src/utils/bizMq"
)

var outRabbitMQ *bizMq.BizQueue

func printMQ(msg string) {
	if outRabbitMQ == nil {
		outRabbitMQ = bizMq.NewBizQueue(bizEnv.Get("LOG_MQ_SERVER"), bizEnv.Get("LOGGER_QUEUE"))
	}
	outRabbitMQ.Publish(msg)
}
