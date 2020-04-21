package bizLogger

import "time"

var logger map[string]*CentralLogger

type CentralLogger struct {
	CreatedAt *time.Time
	Logger    *StackLogger
}

func init() {
	logger = make(map[string]*CentralLogger)
}

func Log(requestId string, title string, contents string) {
	if stack, isExist := logger[requestId]; isExist {
		stack.Logger.Log(title, contents)
	} else {
		now := time.Now()
		newLogger := CentralLogger{
			CreatedAt: &now,
			Logger:    NewStackLogger(),
		}
		newLogger.Logger.Log(title, contents)
		logger[requestId] = &newLogger
	}
}

func Flush(requestId string) {
	if stack, isExist := logger[requestId]; isExist {
		stack.Flush()
		logger[requestId] = nil
	}
}

func newCentralLogger() *CentralLogger {
	cl := CentralLogger{
		CreatedAt: *time.Now(),
	}
}
