package router // import "github.com/kidoda/godemonsaw/router"

import (
	"net"
	"testing"

	"github.com/gorilla/http"
)

func TestLog(logger *Logger, t *testing.T) {
	try := Logger.Emit(0)

}
