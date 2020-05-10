package bizMqLogger

import (
	"encoding/json"
	"fmt"
	"github.com/ciaolee87/echo-starter/src/utils/bizRabbitMq"
	"log"
	"time"
)

type LogData struct {
	Time     string      `json:"time"`
	ServerID string      `json:"serverID"`
	Log      interface{} `json:"log"`
}

type LogDataUnit struct {
	Log string `json:"log"`
}

var (
	conn        *bizRabbitMq.Connection
	queue       *bizRabbitMq.Queue
	svId        string
	ORDER_STACK = "00"
	ORDER_FLUSH = "01"
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
func SendLog(logId string, order string, value interface{}) {
	logData := LogData{
		Time:     time.Now().Format(time.RFC3339Nano),
		ServerID: svId,
		Log:      value,
	}
	byteString, err := json.Marshal(logData)
	if err != nil {
		log.Fatal("마샤링 실패")
	}
	madeLog := fmt.Sprintf("%s|%s|%s", logId, order, string(byteString))
	log.Println(madeLog)
	queue.BizPublish([]byte(madeLog))
}
