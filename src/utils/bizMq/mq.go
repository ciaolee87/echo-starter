package bizMq

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

type BizQueueFunc interface {
	Publish(msg string)
	Consume(func([]byte))
}

type BizQueue struct {
	connection *amqp.Connection
	queue      *amqp.Queue
	channel    *amqp.Channel
}

func (m *BizQueue) Publish(msg string) {
	err := m.channel.Publish(
		"",
		m.queue.Name,
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "utf8",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(msg),
		},
	)

	if err != nil {
		log.Fatal("큐 데이터 전송실패 ")
	}
}

func (m *BizQueue) Consume(callback func([]byte)) {
	msgs, err := m.channel.Consume(
		m.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("consumer 등록 실패")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			callback(d.Body)
		}
	}()

	<-forever
}

func NewBizQueue(server string, queueName string) *BizQueue {
	// 서버 연결
	conn, err := amqp.Dial(server)
	if err != nil {
		log.Fatal("레빗엠큐 로그인 실패", server, queueName)
	}

	// 체널 연결
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("레빗엠큐 체널 오픈 실패")
	}

	// 큐 생성하기
	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(queueName, " 큐 생성 실패")
	}

	return &BizQueue{
		connection: conn,
		queue:      &queue,
		channel:    ch,
	}
}
