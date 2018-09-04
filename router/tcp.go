package router

import (
	"bufio"
	"log"
	"net"
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

func (c *TCPConn) Listen(readchan chan []byte, writechan chan []byte) {
serveloop:
	for {
		conn, err := c.Listener.Accept()
		if err != nil {
			log.Print(err)
			conn.Close()
			break serveloop
		}
		c.Conn = conn
		var readr = bufio.NewReader(c.Conn)
		var readBuff = make([]byte, len(readr.Size()))
		bufcp, _ := readr.ReadBytes(nil)
		for i := range bufcp {
			append(readBuff, i)
		}
		readchan <- readBuff

		var writeBuff = bufio.NewWriter(c.Conn)
		for range writechan {
			writeBuff.Write(<-writechan)
		}
	}
}

func (c *TCPConn) SetListener() error {
	c.Listener, err = net.ListenTCP(c.Addr.Network(), *c.Addr)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
