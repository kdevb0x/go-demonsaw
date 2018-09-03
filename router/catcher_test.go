package router // import "github.com/kidoda/go-demonsaw/router"

import (
	"net"
	"testing"

	"github.com/gorilla/http"
)

func TestLog(logger *Logger, t *testing.T) {
	try := Logger.Emit(0)

}
