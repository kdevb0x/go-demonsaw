// Package catcher allows to catch and interogate http responses.
package router // import "github.com/kidoda/go-demonsaw/router"

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	_ "github.com/google/gopacket/pcap"
	_ "github.com/google/gopacket/tcpassembly"
	_ "github.com/google/gopacket/tcpassembly/tcpreader"
	"github.com/gorilla/mux"
)

type Cer struct {
	Router *mux.Router
	Route  *mux.Route
}

func NewCatcher() *Cer {
	catcher := &Cer{}
	catcher.Router = mux.NewRouter()
	catcher.Route = catcher.Router.Methods("POST")
	return catcher
}

type Catcher interface {
	CatchRequest() (*http.Request, error)
	Listen(host, port string) error
}

func (c *Cer) CatchRequest() (*http.Request, error) {
	c.Route = c.Router.Handle("/", c.Router).Methods("POST")
	var buf bufio.Reader
	req, err := http.ReadRequest(&buf)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return req, nil
}

func (c *Cer) GetReqHeader() (http.Header, error) {
	req, err := c.CatchRequest()
	return req.Header, err

}

func writeToFile(data []byte, filename string) (int, error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Unable to create file: %s, %s", filename, err)
		return 0, err
	}
	defer file.Close()
	size, err := file.Write(data)
	if err != nil {
		log.Printf("Error while writting file, only  %v bytes written. ERROR: %s", size, err)
		return size, err
	}
	return size, nil
}

/*
func (c *Cer) Listen(host, port string) error {
	router := NewCatcher()
	conn, err := net.Listen("tcp", host+port)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		requests, err := conn.Accept()
		if err != nil {
			return err
		}
		header, err := router.GetReqHeader()
		if err != nil {
			return err
		}

		router.Router.ServeHTTP(requests, *router.CatchRequest())
	}
	return nil
}
*/
