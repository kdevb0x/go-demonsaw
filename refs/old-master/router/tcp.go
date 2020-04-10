package router

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"net/http"
)

type TCPConn struct {
	Conn     net.Conn
	Addr     *net.TCPAddr
	Listener *net.TCPListener
}

// NewTCPConn creates a new TCPConn and returns a pointer to it.
func NewTCPConn(host string) (*TCPConn, error) {
	tcpconn := &TCPConn{}
	addr, err := net.ResolveTCPAddr("tcp", host)
	tcpconn.Addr = addr
	if err != nil {
		log.Printf("Error creating NewTCPConn: %s", err)
		return nil, err
	}
	return tcpconn, nil

}

func (c *TCPConn) Listen(writechan chan []byte) chan http.Request {
serveloop:
	reqchan := make(chan http.Request, 4096) // Allocate an array twice as big as we need to be safe

	for {
		conn, err := c.Listener.Accept()
		if err != nil {
			log.Print(err)
			conn.Close()
			break serveloop
		}
		c.Conn = conn
		defer c.Conn.Close()

		var readr = bufio.NewReader(c.Conn) // Reader for incoming requests.
		req, err := http.ReadRequest(readr) // When a request comes in, read it
		if err != nil {
			log.Printf("Error reading incoming request: %s", err)
			break
		}

		/* TODO: Create a custom type that implements the
		   http.ResonseWriter interface; Header() Header  */

		var respWriter http.ResponseWriter

		var writeBuff = bufio.NewWriter(c.Conn)
		if len(writechan) > 0 {
			for rcv := range writechan {
				writeBuff.Write(rcv)
			}

		}

	}
}

func (c *TCPConn) SetListener() error {
	clisten, err := net.ListenTCP(c.Addr.Network(), c.Addr)
	if err != nil {
		log.Print(err)
		return err
	}
	c.Listener = clisten

	return nil
}
