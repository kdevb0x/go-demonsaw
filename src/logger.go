package godemonsaw // import "github.com/kidoda/godemonsaw"

import (
	"io"
	"os"
	"time"
)

type LoggerType uint

const (
	Void LoggerType = iota
	StdErr
	Panics
	Warnings
	Debug
)

type Logger interface {
	Emit(string, []string)
	Flush() bool
	Destroy(func(), time.Time)
}

type VoidLogger struct {
	Type    LoggerType
	Outchan chan []byte
	LogFile *os.File
}

// Create a VoidLogger, VoidLogger is a dummylogger.
func (l *VoidLogger) NewVoidLogger() *VoidLogger {
	logc := make(chan []byte, 100)
	logger := &VoidLogger{Void, logc, nil}
	return logger
}
