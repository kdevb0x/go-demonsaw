package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
)

type routerType int

const (
	None routerType = iota
	TransferRouter
	MessageRouter
	MainRouter
)

type router struct {
	ipaddr string
	port   int
	// network string
	httpServer *http.Server

	config *routerConfig
}

func newRouter(host, port string) (*router, error) {
	var s = new(http.Server)
	s.Addr = host + ":" + port
	r := new(router)
	r.httpServer = s
	return r, nil
}

func (r router) Type() routerType {
	return MainRouter
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
	address string
	// password = ""
	password []byte
	// port = 8000
	port int

	options []routerOption

	rooms []room
	// transferRouters contains ALL transferRouters, both enabled, and
	// possibly dissabled; Use activeRouters(xferRouter)
	transferers []transferRouter
}

func (rc routerConfig) activeRouters(rt routerType) []RouterTyper {
	var ready = make([]RouterTyper, len(rc.transferers))
	for _, r := range rc.transferers {
		if r.Type() == rt {
			if r.enabled {
				ready[len(ready)] = r
			}
		}
	}
	return ready
}

type RouterTyper interface {
	Type() routerType
}

// 	[router.option]
type routerOption struct {
	// buffer_size = 32
	buffer_size int
	// motd = "Welcome to Demonsaw 4"
	motd string
	// redirect = "https://www.demonsaw.com"
	redirect string
}

// 	[[router.room]]
type room struct {
	// enabled = true
	enabled bool

	// name = "Room #1"
	name string // maybe []byte?

	// color = "ff52c175"
	color [4]byte `toml:"color"`
}

//
// NOTE: (kdv) I debated between xferRouter and transferRouter, the latter was
// chosen for readability.
//
//
// called [[router.router]] in demonsaw.toml
//
// transferRouter accepts encrypted bytes and routes them to receivers.
type transferRouter struct {
	// enabled = false
	enabled bool
	// name = "Transfer Router #1"
	name string
	// address = "127.0.0.1"
	address string
	// password = ""
	password []byte
	// port = 80
	port int
}

func (tr transferRouter) Type() routerType {
	return TransferRouter
}

func (r *router) loadConfig(tomlpath string) error {
	var conf = new(routerConfig)
	m, err := toml.DecodeFile(tomlpath, conf)
	if err != nil {
		return err
	}

	conf.fpath = tomlpath
	r.config = *conf
	if len(m.Undecoded()) > 0 {
		return errors.New(fmt.Sprintf("error: unable to decode %d values from toml: %v\n", len(m.Undecoded()), m.Undecoded()))
	}
	return nil
}
