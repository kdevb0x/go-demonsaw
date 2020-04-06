package client // import "github.com/kidoda/go-demonsaw/client"
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	routerHost = "router.demonsaw.com"
	routerIP   = "104.238.179.22"
)

// RouterConn contains the state of a demonsaw router connection.
type RouterConn struct {
	URL, IP    string
	NetConn    net.Conn
	ConnReader bufio.Reader
	ConnBuff   []byte
}

func dialRouter(url string) (*RouterConn, error) {
	dialer, err := net.Dial("tcp", url)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	fmt.Fprintf(dialer, "POST / HTTP/1.1\r\n\r\n")
	var buff = bufio.NewReader(dialer)
	buff.ReadString('\n')
	return dialer, nil
}

// func main() {
//	routerConn := dialRouter(router)
//
// }
