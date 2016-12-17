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
			fmt.Println("Found at least one file.")

			// open uploaded
			var infile multipart.File
			infile, err := hdr.Open()
			if err != nil {
				fmt.Println(err)
				return
			}
			// open destination
			var outfile *os.File
			path := "/static/html/" + hdr.Filename
			if outfile, err = os.Create(path); nil != err {
				fmt.Println(err)
				return
			}
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				fmt.Println(err)
				return
			}
			fmt.Println("Successfully uploaded file. Getting ready to run hadoop job.")
			fmt.Println(RunHadoop(path))
			fmt.Println(RunSpark(path))
			w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
			fmt.Println("Finished processing the jobs!")
		}
	}
}

// SearchHandler handles the search function.
func SearchHandler(w http.ResponseWriter, req *http.Request) {
	var i invertedIndex

	t := MeasureTime(func() {
		i = *loadInvertedIndex(indexPathSpark)
	})
	i.Time = t.String()
	js, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
