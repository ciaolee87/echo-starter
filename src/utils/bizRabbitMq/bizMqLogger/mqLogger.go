package bizMqLogger

import (
	"encoding/json"
	"fmt"
	"github.com/ciaolee87/echo-starter/src/utils/bizRabbitMq"
	"time"
)

type MqLogger struct {
	logId string `json:"logId"`
	id    string `json:"id"`
	log   string `json:"log"`
}

var (
	conn  *bizRabbitMq.Connection
	queue *bizRabbitMq.Queue
	svId  string
)

// 로거를 초기화 한다.
func InitMqLogger(
	connection *bizRabbitMq.Connection,
	queueName string,
	serverId string,
) {
	conn = connection
	queue = conn.NewBizQueue(queueName)
	svId = serverId
}

// 바로 로그 큐로 보낸다. -> id 파악할수 있게
func Log(logId string, value string) {
	logUnit := MqLogger{
		id:  fmt.Sprintf("%s : %s", time.Now().Format(time.RFC3339Nano), svId),
		log: value,
	}

	strValue, _ := json.Marshal(logUnit)
	queue.BizPublish(strValue)
}
