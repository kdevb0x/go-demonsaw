package router

import (
	"bufio"
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
	if err != nil {
		log.Printf("Error creating NewTCPConn: %s", err)
		return nil, err
	}
	tcpconn.Addr = addr
	return tcpconn, nil

}

func (c *TCPConn) Listen(writechan chan http.Response) chan http.Request {
	reqchan := make(chan http.Request, 4096) // Allocate an chan twice as big as we need to be safe
	var serveloop = func(reqChan chan http.Request) error {
		for {
			conn, err := c.Listener.Accept()
			if err != nil {
				log.Print(err)
				conn.Close()
				return err
			}
			c.Conn = conn
			defer c.Conn.Close()

			var readr = bufio.NewReader(c.Conn) // Reader for incoming requests.
			req, err := http.ReadRequest(readr) // When a request comes in, read it
			if err != nil {
				log.Printf("Error reading incoming request: %s", err)
				return err
			}
			reqChan <- *req
			/* TODO: Create a custom type that implements the
			   http.ResonseWriter interface; Header() Header  */

			var respWriter http.ResponseWriter

			var writeBuff = bufio.NewWriter(c.Conn)
			select {
			case res := <-writechan:
				writeBuff.Write([]byte(res))
			default:
				respWriter.WriteHeader(http.StatusOK)
			}

		}

	}
	go serveloop(reqchan)
	return reqchan
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
