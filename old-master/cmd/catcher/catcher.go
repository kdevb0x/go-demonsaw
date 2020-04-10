package main // import "github.com/kidoda/go-demonsaw/cmd/catcher"

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host = "localhost"
	port = ":8080"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		for key, val := range header {
			fmt.Printf("Key: %s, Value: %s", key, val)
		}
		fmt.Print(ioutil.ReadAll(r.Body))
		return
	})
	log.Fatal(http.ListenAndServe(port, router))
}

/*
func main() {
	catcher := router.NewCatcher()
	header, err := catcher.GetReqHeader()
	if err != nil {
		log.Printf("error getting request header: %s", err)
	}
	fmt.Print(header)
	err := router.CatcherListen(host, port)
}
*/
