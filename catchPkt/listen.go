package catchPkt

import (
	"bufio"
	"context"
	"net"
	"net/http"
	"os"
)

var ListenAddr string = os.Args[1]

type TcpListener struct {
	Addr string
	Port byte
}

func (t *TcpListener) Listen(ctx context.Context, addr string) (chan http,Request, net.Conn, error) {
	cc := make(chan net.Conn)
	var reqchan = make(chan http.Request)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return, err
		}
		ctx, cancelFunc := context.WithCancel(ctx)
		readr := bufio.NewReader(conn)
		req, err := http.ReadRequest(readr)
		if err != nil {
			return nil, err
		}

		reqChan <- req

	}

}

func (t *TcpListener) Accept() (net.Conn, error) {

}

func (t *TcpListener) Close() error {

}

func (t *TcpListener) Addr() net.Addr {
	return t.Addr
}
