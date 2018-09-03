package main // import "github.com/kidoda/go-demonsaw/cmd/catcher"

import (
	"fmt"

	"github.com/kidoda/go-demonsaw/router"
)

func main() {
	catcher := router.NewCatcher()
	header, err := catcher.GetReqHeader()
	if err != nil {
		log.Printf("error getting request header: %s", err)
	}
	fmt.Print(header)
}
