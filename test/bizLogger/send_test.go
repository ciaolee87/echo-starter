package bizLogger

import (
	"github.com/ciaolee87/echo-starter/src/utils/bizRabbitMq"
	"github.com/ciaolee87/echo-starter/src/utils/bizRabbitMq/bizMqLogger"
	"github.com/hashicorp/go-uuid"
)

type LogEx struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ExampleSendLog() {
	conn := bizRabbitMq.NewConnection("amqp://root:flowd1234@13.125.105.222:5500")

	bizMqLogger.InitMqLogger(
		conn,
		"msgService",
		"front",
	)

	uid, _ := uuid.GenerateUUID()
	bizMqLogger.SendLog(uid, bizMqLogger.ORDER_STACK, "1")

	logEx := LogEx{
		Name: "Hello",
		Age:  10,
	}

	bizMqLogger.SendLog(uid, bizMqLogger.ORDER_STACK, logEx)
	bizMqLogger.SendLog(uid, bizMqLogger.ORDER_STACK, "3")
	bizMqLogger.SendLog(uid, bizMqLogger.ORDER_FLUSH, "4")

	// Output:

}
