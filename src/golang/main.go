package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/RobbieMcKinstry/tiny-google/handlers"
	"github.com/gorilla/mux"
)

func main() {
	time.Sleep(9000 * time.Millisecond)

	fmt.Printf("Beginning the server on port %v \n", port)

	// Set up the router
	r := mux.NewRouter()
	s := http.StripPrefix(handlers.StaticDir, http.FileServer(http.Dir(handlers.StaticDir)))
	r.PathPrefix(handlers.StaticDir).Handler(s)

	r.Methods("GET").Path("/tiny-google").HandlerFunc(handlers.SearchHandler)
	r.HandleFunc("/", handlers.IndexHandler)

	http.Handle("/", r)

	panic(http.ListenAndServe(handlers.Port, nil))
}
