package catcher // import "github.com/kidoda/godemonsaw/catcher"

import (
	"net"
	"testing"

	"github.com/gorilla/http"
)

func TestLog(logger *Logger, t *testing.T) {
	try := Logger.Emit(0)

}
