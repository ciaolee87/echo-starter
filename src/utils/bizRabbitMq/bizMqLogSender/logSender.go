package bizMqLogSender

import (
	"encoding/json"
	"errors"
	"github.com/ciaolee87/echo-starter/src/utils/bizRabbitMq"
	"github.com/hashicorp/go-uuid"
	"log"
	"time"
)

type BizLogUnit struct {
	id       string    `json:"id"`       // uuid
	time     time.Time `json:"time"`     // 로깅 시간
	serverID string    `json:"serverId"` // 서버 아이디
	order    string    `json:"order"`    // order : STACK, FLUSH
	body     []BizLog  `json:"body"`     // log 내용
}

type BizLog struct {
	key   string `json:"key"`
	value string `json:"value"`
}

var (
	STACK = "STACK"
	FLUSH = "FLUSH"
	conn  *bizRabbitMq.Connection
	queue *bizRabbitMq.Queue
	svId  = ""
)

// sever : mq 주소 (아이디, 비번 포함)
// serverId : 서버 아이디
func InitLogger(server string, queueName string, serverId string) {
	conn = bizRabbitMq.NewConnection(server)
	queue = conn.NewBizQueue(queueName)
	svId = serverId
}

func NewLogger(uid *string) *BizLogUnit {
	if uid != nil {
		// uuid 있을때
		return &BizLogUnit{
			id:       *uid,
			time:     time.Now(),
			serverID: svId,
		}
	}

	// uuid 없을떄
	id, err := uuid.GenerateUUID()
	if err != nil {
		log.Fatal("Logger uuid 생성실패!")
	}

	return &BizLogUnit{
		id:       id,
		time:     time.Now(),
		serverID: svId,
	}

}

func (lu *BizLogUnit) Log(key string, value string) {
	lu.body = append(lu.body, BizLog{
		key:   key,
		value: value,
	})
}

// 정상 출력
func (lu *BizLogUnit) SendStack() error {
	return lu.Send(STACK)

}

// 비정상 출력
func (lu *BizLogUnit) SendFlush() error {
	return lu.Send(FLUSH)
}

// 전송하기
func (lu *BizLogUnit) Send(order string) error {
	switch order {
	case STACK, FLUSH:
		lu.order = order
	default:
		return errors.New("Send order is not available! " + order)
	}

	// json 으로 마샬링
	byteBody, err := json.Marshal(lu)
	if err != nil {
		return err
	}

	queue.BizPublish(byteBody)
	return nil
}
