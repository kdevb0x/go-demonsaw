package router

import (
	"log"
	"net"
)

type TCPConn struct {
	Conn net.Conn
	Addr net.TCPAddr
}

// NewTCPConn creates a new TCPConn and returns a pointer to it.
func NewTCPConn(host string) (*TCPConn, error) {
	tcpconn := &TCPConn{}
	tcpconn.Addr, err = net.ResolveTCPAddr("tcp", host)
	if err != nil {
		log.Printf("Error creating NewTCPConn: %s", err)
		return nil, err
	}
	tcpconn.Conn, err = net.ListenTCP(tcpconn.Addr.Network(), *tcpconn.Addr)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return tcpconn, nil

}
