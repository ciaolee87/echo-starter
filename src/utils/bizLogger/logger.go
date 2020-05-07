package bizLogger

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type LogData struct {
	time     time.Time
	title    string
	contents string
}

// 이곳에서 로그 프린트 설정정으로 한다
func printLog(data *[]LogData) {
	msg := makeMsgToJson(data)
	// 화면에 표시
	log.Print(msg)
	// 큐에 데이터를 보낸다
	printMQ(msg)
}

// 바로보내는 로그
func LineLogger(title string, content string) {
	logData := []LogData{{
		time:     time.Time{},
		title:    title,
		contents: content,
	}}
	printLog(&logData)
}

// 로그데이터를 JSON 스트링으로 만든다
// 시간:타이틀:컨텐츠 |  시간:타이틀:컨텐츠 | 시간:타이틀:컨텐츠 |  ... 형태
func makeMsgToJson(data *[]LogData) (msg string) {

	msg = fmt.Sprintf("{\"time\":\"%d\",", time.Now().Format())

	for i, logData := range *data {

		msg += fmt.Sprintf("\"%s\":\"%s\"", logData.title, *escapeRemove(&logData.contents))

		// 마지막 데이터가 아니면 , 를 붙혀준다
		if i != len(*data)-1 {
			msg += ","
		}
	}
	msg += "}"
	return
}

// 개행문자 삭제 하기
func escapeRemove(msg *string) *string {
	*msg = strings.Replace(*msg, "\n", " ", -1)
	*msg = strings.Replace(*msg, "\r", " ", -1)
	return msg
}
