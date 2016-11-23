package main

import (
	"fmt"
	"net/http"
)

const (
	staticDir = "/static"
	port      = ":8080"
)

func main() {
	fmt.Println("This is not hello world.")
	fmt.Printf("Beginning the server on port %v \n", port)
	panic(http.ListenAndServe(port, http.FileServer(http.Dir(staticDir))))
	// First, we need to connect to the Yarn daemon.
	// Load the jar and create the context

	// Then, we add the HTTP handlers to a simple server.
}
