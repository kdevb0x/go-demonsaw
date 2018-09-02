// Package catcher allows to catch and interogate http responses.
package catcher

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

func CatchRequest() (*http.Request, error) {

}
