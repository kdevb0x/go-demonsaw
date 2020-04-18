package router

import (
	"net"
	"net/http"
)

type routerType int

const (
	None routerType = iota
	TransferRouter
	MessageRouter
	MainRouter
)

type msgRouter struct {
	ipaddr string
	port   int
	// network string
	httpServer *http.Server

	config *routerConfig
}

func newMsgRouter(host, port string) (*msgRouter, error) {
	var s = new(http.Server)
	s.Addr = host + ":" + port
	r := new(msgRouter)
	r.httpServer = s
	return r, nil
}

func (r msgRouter) Type() routerType {
	return MessageRouter
}

// [[router]]
type routerConfig struct {
	// path to toml file
	fpath string
	// enabled = true
	enabled bool `toml:"enabled"`

	// threads = 128
	threads int `toml:"threads"`

	// name = "Message Router #1"
	name string `toml:"name"`

	// address = "127.0.0.1"
	addr string `toml:"address"`
	// password = ""
	password []byte `toml:"password,omitempty"`
	// port = 8000
	port int `toml:"port"`

	options []routerOption

	rooms []room
	// transferRouters contains ALL transferRouters, both enabled, and
	// possibly dissabled; Use activeRouters(xferRouter)
	transferers []transferRouter
}

func (rc routerConfig) activeRouters(rt routerType) []RouterType {
	var ready = make([]RouterType, len(rc.transferers))
	for _, r := range rc.transferers {
		if r.Type() == rt {
			if r.enabled {
				ready[len(ready)] = r
			}
		}
	}
	return ready
}

type RouterType interface {
	Type() routerType
}

// 	[router.option]
type routerOptions struct {
	// buffer_size = 32
	buffer_size int `toml:"buffer_size"`
	// motd = "Welcome to Demonsaw 4"
	motd string `toml:"motd,omitempty"`
	// redirect = "https://www.demonsaw.com"
	redirect string `toml:"redirect,omitempty"`
}

// NOTE: (kdv) I debated between xferRouter and transferRouter, the latter was
// chosen for readability.
//
//
// called [[router.router]] in demonsaw.toml
//
// transferRouter accepts encrypted bytes and routes them to receivers.
type transferRouter struct {
	// enabled = false
	enabled bool `toml:"enabled"`
	// name = "Transfer Router #1"
	name string `toml:"name"`
	// address = "127.0.0.1"
	addr string `toml:"address"`
	// password = ""
	password []byte `toml:"password,omitempty"`
	// port = 80
	port int `toml"port"`
}

func newTransferRouter(parent *msgRouter, addr string, password ...[]byte) (*transferRouter, error) {
	ip, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &transferRouter{
		enabled: false,
		addr:    ip.String(),
		port:    ip.Port,
	}, nil
}
func (tr transferRouter) Type() routerType {
	return TransferRouter
}

func (tr *transferRouter) 
