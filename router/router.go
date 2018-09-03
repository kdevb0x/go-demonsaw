// Package router is a demonsaw router implementation.
package router // import "github.com/kidoda/go-demonsaw/router"

import (
	_ "bufio"
	_ "io"
	"log"
	_ "net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type server struct {
	http.Server
	/*
		Addr string
			// TCP address to listen on, ":http" if empty
		Handler Handler
			// handler to invoke, http.DefaultServeMux if nil
		TLSConfig *tls.Config
			// optional TLS config, used by ListenAndServeTLS
			// ReadTimeout is the maximum duration for reading the entire
			// request, including the body.
			//
			// Because ReadTimeout does not let Handlers make per-request
			// decisions on each request body's acceptable deadline or
			// upload rate, most users will prefer to use
			// ReadHeaderTimeout. It is valid to use them both.
		ReadTimeout time.Duration
			// ReadHeaderTimeout is the amount of time allowed to read
			// request headers. The connection's read deadline is reset
			// after reading the headers and the Handler can decide what
			// is considered too slow for the body.
		ReadHeaderTimeout time.Duration
			// WriteTimeout is the maximum duration before timing out
			// writes of the response. It is reset whenever a new
			// request's header is read. Like ReadTimeout, it does not
			// let Handlers make decisions on a per-request basis.
		WriteTimeout time.Duration
			// IdleTimeout is the maximum amount of time to wait for the
			// next request when keep-alives are enabled. If IdleTimeout
			// is zero, the value of ReadTimeout is used. If both are
			// zero, there is no timeout.
		IdleTimeout time.Duration
			// MaxHeaderBytes controls the maximum number of bytes the
			// server will read parsing the request header's keys and
			// values, including the request line. It does not limit the
			// size of the request body.
			// If zero, DefaultMaxHeaderBytes is used.
		MaxHeaderBytes int
			// TLSNextProto optionally specifies a function to take over
			// ownership of the provided TLS connection when an NPN/ALPN
			// protocol upgrade has occurred. The map key is the protocol
			// name negotiated. The Handler argument should be used to
			// handle HTTP requests and will initialize the Request's TLS
			// and RemoteAddr if not already set. The connection is
			// automatically closed when the function returns.
			// If TLSNextProto is not nil, HTTP/2 support is not enabled automatically.
		TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
			// ConnState specifies an optional callback function that is
			// called when a client connection changes state. See the
			// ConnState type and associated constants for details.
		ConnState func(net.Conn, ConnState)
			// ErrorLog specifies an optional logger for errors accepting
			// connections and unexpected behavior from handlers.
			// If nil, logging goes to os.Stderr via the log package's
			// standard logger.
		ErrorLog *log.Logger
	*/
}

type Router interface {
	// ServeHTTP(w http.ResponseWriter, r *http.Request)
	Listen() error
}

type router struct {
	server
}

func (r *router) listen() (*http.Request, error) {
	handle := mux.NewRouter()
	route := handle.NewRoute()
	handler := route.GetHandler()
	handler.ServeHTTP(http.ResponseWriter, *http.Request)
}

func (lc logChan) LogToFile(filename string, data []byte) (int, error) {
	file, err := os.Create(filename)
	if err != nil {

		return nil, err
	}
}
