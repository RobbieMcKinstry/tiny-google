package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Printf("Beginning the server on port %v \n", Port)

	// Set up the router
	r := mux.NewRouter()
	s := http.StripPrefix(StaticDir, http.FileServer(http.Dir(StaticDir)))
	r.PathPrefix(StaticDir).Handler(s)
	fmt.Println("Registered routes. Getting ready to launch")

	r.Methods("GET").Path("/tiny-google").HandlerFunc(SearchHandler)
	r.Methods("POST").Path("/upload").HandlerFunc(UploadHandler)
	r.HandleFunc("/", IndexHandler)

	http.Handle("/", r)

	panic(http.ListenAndServe(Port, nil))
}
