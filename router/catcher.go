// Package catcher allows to catch and interogate http responses.
package router // import "github.com/kidoda/go-demonsaw/router"

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	_ "github.com/google/gopacket/tcpassembly"
	_ "github.com/google/gopacket/tcpassembly/tcpreader"
	_ "github.com/gorilla/http"
	_ "gthub.com/google/gopacket/pcap"
)

type catcher struct {
	router
}

type Catcher interface {
	CatchRequest() (*http.Request, error)
	GetReqHeader() *http.Header
}

func (c *catcher) CatchRequest() (*http.Request, error) {
	req, err := c.listen()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return req, nil
}

func (c *catcher) GetReqHeader() *http.Header {
	req, _ := c.CatchRequest()
	return req.Header

}
