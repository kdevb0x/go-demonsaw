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

type logger struct {
	stringChanIn chan []string
	errChanIn    chan error
	logOut       io.Writer
}

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
func NewVoidLogger() *VoidLogger {
	logc := make(chan []byte, 100)
	logger := &VoidLogger{Void, logc, nil}
	return logger
}

// Creates a new logger. If out is nil, defaults to stderr.
func Newlogger(out io.Writer) *logger {
	logger := &logger{}
	logger.stringChanIn = make(chan []string, 10, 2000)
	logger.errChanIn = make(chan error, 10, 2000)
	if out == nil || ! > 0 {
		logger.logOut = os.StdErr
	}
}
