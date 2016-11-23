package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	staticDir = "/static"
	port      = ":8080"
)

func main() {
	fmt.Println("This is not hello world.")
	fmt.Printf("Beginning the server on port %v \n", port)

	r := mux.NewRouter()

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("/static/")))
	r.PathPrefix("/static/").Handler(s)

	r.Methods("GET").Path("/tiny-google").HandlerFunc(searchHandler)
	r.HandleFunc("/", indexHandler)

	http.Handle("/", r)

	panic(http.ListenAndServe(port, nil))
	// First, we need to connect to the Yarn daemon.
	// Load the jar and create the context

	// Then, we add the HTTP handlers to a simple server.
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/static/html", http.StatusFound)
}

func searchHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `{ "%v": { "%v": "%v" }}`, "spark", "runtime", "100 sec")
	fmt.Printf(`{ "%v": "%v" } \n`, "result", "foobar")
}
