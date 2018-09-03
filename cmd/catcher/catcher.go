package main // import "github.com/kidoda/go-demonsaw/cmd/catcher"

import (
	"fmt"
	"log"

	"github.com/kidoda/go-demonsaw/router"
)

const (
	host = "localhost"
	port = ":8080"
)

func main() {
	catcher := router.NewCatcher()
	header, err := catcher.GetReqHeader()
	if err != nil {
		log.Printf("error getting request header: %s", err)
	}
	fmt.Print(header)
	err := router.CatcherListen(host, port)
}
