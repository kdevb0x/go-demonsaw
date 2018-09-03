// Package catcher allows to catch and interogate http responses.
package router // import "github.com/kidoda/go-demonsaw/router"

import (
	"log"
	_ "net"
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

func writeToFile(data []byte, filename string) (int, error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Unable to create file: %s, %s", filename, err)
		return nil, err
	}
	defer file.Close()
	size, err := file.Write(data)
	if err != nil {
		log.Printf("Error while writting file, only  %v bytes written. ERROR: %s", size, err)
		return size, errr
	}
	return size, nil
}
