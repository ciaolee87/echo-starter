package bizLogger

import (
	"io"
	"strings"
	"time"
)

type StackLogger struct {
	logs []LogData
	io.Writer
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

func (s *StackLogger) Write(body []byte) (int, error) {
	str := string(body)
	s.Log("io.writer", string(body))
	return len(str), nil
}

func (s *StackLogger) Flush() {
	printLog(&s.logs)
}
