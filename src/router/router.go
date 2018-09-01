// Package router is a demonsaw router implementation.
package router // import "github.com/kidoda/godemonsaw/router"

import (
	_ "bufio"
	_ "io"
	_ "net"
	"net/http"

	_ "github.com/gorilla/http"
)

type Router interface {
	// ServeHTTP(w http.ResponseWriter, r *http.Request)
	Listen()
}

type router struct {
	http.Server
}

func (r *router) Listen() {

}
