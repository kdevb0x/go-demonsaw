// Package catcher allows to catch and interogate http responses.
package router // import "github.com/kidoda/go-demonsaw/router"

import (
	"bufio"
	"log"
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
	GetReqHeader() *http.Header
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

func (c *Cer) GetReqHeader() *http.Header {
	req, _ := c.CatchRequest()
	return &req.Header

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
