package bizLogger

import (
	"strings"
	"time"
)

type StackLogger struct {
	logs []LogData
}

func NewStackLogger() *StackLogger {
	return &StackLogger{logs: make([]LogData, 0)}
}

// 로그 추가하기
func (s *StackLogger) Log(title string, contents string) {

	data := LogData{
		time:     time.Now(),
		title:    title,
		contents: strings.Replace(contents, "\"", "'", -1),
	}

	s.logs = append(s.logs, data)
}

func (s *StackLogger) Flush() {
	printLog(&s.logs)
}
