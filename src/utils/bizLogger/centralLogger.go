package bizLogger

import (
	"time"
)

var CLogger = make(map[string]*CentralLogger)

type CentralLogger struct {
	CreatedAt *time.Time
	Logger    *StackLogger
}

func Log(requestId string, title string, contents string) {
	if stack, isExist := CLogger[requestId]; isExist {
		stack.Logger.Log(title, contents)
	} else {
		if requestId == "" {
			return
		}
		now := time.Now()
		newLogger := CentralLogger{
			CreatedAt: &now,
			Logger:    NewStackLogger(),
		}
		newLogger.Logger.Log(title, contents)
		CLogger[requestId] = &newLogger
	}
}

func Flush(requestId string) {
	if stack, isExist := CLogger[requestId]; isExist {
		stack.Logger.Flush()
		CLogger[requestId] = nil
	}
}
