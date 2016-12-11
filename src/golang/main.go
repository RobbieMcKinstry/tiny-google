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

	"github.com/gorilla/mux"
)

const (
	staticDir = "/static"
	port      = ":8080"
)

const (
	endpointBase       = "http://spark:8090"
	jarUploadExtension = "/jars/test"
	jarUploadEndpoint  = endpointBase + jarUploadExtension

	contextExtension = "/contexts/test-context?num-cpu-cores=2&memory-per-node=1024m"
	contextEndpoint  = endpointBase + contextExtension

	executeURI  = endpointBase + "/jobs?appName=test&classPath=spark.jobserver.SimpleApp&context=test-context&sync=true"
	jarFilePath = "/mounted/target/scala-2.11/simple-project_2.11-1.0.jar"
)

func main() {
	time.Sleep(8000 * time.Millisecond)

	fmt.Println("This is not hello world.")
	fmt.Printf("Beginning the server on port %v \n", port)

	// Initialize the context
	initContext()

	// Set up the router
	r := mux.NewRouter()
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("/static/")))
	r.PathPrefix("/static/").Handler(s)

	r.Methods("GET").Path("/tiny-google").HandlerFunc(searchHandler)
	r.HandleFunc("/", indexHandler)

	http.Handle("/", r)

	panic(http.ListenAndServe(port, nil))
}

func initContext() {
	if _, err := postFile(jarFilePath, jarUploadEndpoint); err != nil {
		log.Println(err)
	}
	if _, err := http.Post(contextEndpoint, "application/json", nil); err != nil {
		log.Println(err)
	}
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/static/html", http.StatusFound)
}

func searchHandler(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Post("http://spark:8090/jobs?appName=test&classPath=spark.jobserver.SimpleApp&context=test-context&sync=true", "application/json", nil)
	if err != nil {
		log.Println(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Here is the response body:")
	fmt.Println(string(content))
	fmt.Println(err)
	fmt.Println("Request completed.")

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

func postFile(filename string, targetURL string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("job.jar", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	fmt.Println(contentType)
	bodyWriter.Close()

	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return resp, nil
}
