package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func init() {
	time.Sleep(9000 * time.Millisecond)
	go initContext()
}

func initContext() {
	if _, err := postFile(jarFilePath, jarUploadEndpoint); err != nil {
		log.Println(err)
	}
	if _, err := http.Post(contextEndpoint, "application/json", nil); err != nil {
		log.Println(err)
	}
}

// IndexHandler handles "/"
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/static/html", http.StatusFound)
}

// UploadHandle uploads the file
func UploadHandler(w http.ResponseWriter, req *http.Request) {
	// parse request
	const _24K = (1 << 13) * 24
	if err := req.ParseMultipartForm(_24K); nil != err {
		fmt.Println("Failed to parse.")
		return
	}

	for _, fheaders := range req.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File
			infile, err := hdr.Open()
			if err != nil {
				fmt.Println(err)
				return
			}
			// open destination
			var outfile *os.File
			if outfile, err = os.Create("/static/html/" + hdr.Filename); nil != err {
				fmt.Println(err)
				return
			}
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				fmt.Println(err)
				return
			}
			w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}
}

// SearchHandler handles the search function.
func SearchHandler(w http.ResponseWriter, req *http.Request) {
	var i invertedIndex

	t := MeasureTime(func() {
		i = *loadInvertedIndex(indexPath)
	})
	i.Time = t.String()
	js, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	/*
			resp := launchSparkJob()
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
	*/
}

func launchSparkJob() *http.Response {
	resp, err := http.Post(jobEndpoint, "application/json", nil)
	if err != nil {
		log.Println(err)
	}
	return resp
}

func loadInvertedIndex(path string) *invertedIndex {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	var index invertedIndex
	if err = json.Unmarshal(file, &index); err != nil {
		log.Println(err)
	}
	return &index
}

// MeasureTime returns how long the function took to run
func MeasureTime(f func()) time.Duration {
	t0 := time.Now()
	f()
	return time.Since(t0)
}
