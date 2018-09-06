package router_test

import (
	"net/http/httptest"
	"testing"
)

const (
	host = ":8080"
)

var tableTests = []string{}

func TestNewTCPConn(t *testing.T) {
	tcpconn, err := NewTCPConn(host)
	if err != nil {
		t.Error(err)
	}
	if tcpconn.Addr.Network() != "tcp" {
		t.Fail()
	}
	if tcpconn.Addr.Port() != host {
		t.Fail()
	}
}
