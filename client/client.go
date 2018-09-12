package client // import "github.com/kidoda/go-demonsaw/client"

import (
	"os"
	"time"
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
