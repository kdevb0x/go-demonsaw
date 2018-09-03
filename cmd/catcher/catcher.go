package main // import "github.com/kidoda/go-demonsaw/cmd/catcher"

import (
	"fmt"
	"net/http"

	"github.com/kidoda/go-demonsaw/router"
)

func main() {
	var catcher router.Catcher
	header := catcher.GetReqHeader()
	for index, inmap := range header {
		fmt.Printf("index: %s, present: %b\n", index, inmap)
	}
}
