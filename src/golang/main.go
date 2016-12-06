package main

import (
	"fmt"
	// "io/ioutil"
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
	// resp, _ := http.Post("http://spark:8090/jobs?appName=test&classPath=spark.jobserver.SimpleApp&context=test-context&sync=true", "application/json", nil)
	// content, _ := ioutil.ReadAll(resp.Body)

	// fmt.Fprintf(w, string(content))
	fmt.Fprintf(w,
		`{  
            "mr_runtime": 100,
            "spark_runtime": 200,
            "links": [{
                document_name: "Alice in Wonderland",
                document_link: "/alice_in_Wonderland",
                context:       "the mad hatte had a wondeful pair of shoes..."
            }, {
                document_name: "Hucklebery Finny",
                document_link: "/hucklebery.txt",
                context:       "Huck's shoes had a big ol hole in them"
            }]
        }`)
}
