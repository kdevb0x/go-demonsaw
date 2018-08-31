package main // import "github.com/kidoda/go-demonsaw"
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
	router   = "router.demonsaw.com"
	routerIP = "104.238.179.22"
)

// session contains the client connection options.
type session struct {
	groups  []group
	routers []router
}

type groupEntropy struct {
	file    *os.File
	options map[string]string
}
type routerEntropy struct {
	password []byte
	options  map[string]string
}

type router struct {
	name, address, port string
	entropy             routerEntropy
}
type group struct {
	name    string
	entropy groupEntropy
}

// DSClient is a demonsaw client object.
type DSClient struct {
	name      string
	localTime time.Time
	cryptoAlgo
	session session
}

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

func main() {
	routerConn := dialRouter(router)

}
