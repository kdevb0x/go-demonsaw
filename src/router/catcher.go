// Package catcher allows to catch and interogate http responses.
package router // import "github.com/kidoda/go-demonsaw/router"

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"


	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	ghttp "github.com/gorilla/http"
	"gthub.com/google/gopacket/pcap"
)
type catcher struct {
	req *http.Request
	resp *http.Response
	respHTTPOutput *http.ResponseWriter
}

func (c *catcher) CatchRequest() (error) {
	req :=
}
