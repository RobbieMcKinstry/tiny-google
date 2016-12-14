package handlers

func init() {
	initContext()
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

// SearchHandler handles the search function.
func SearchHandler(w http.ResponseWriter, req *http.Request) {

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
}

func launchSparkJob() *http.Response {
	resp, err := http.Post(jobEndpoint, "application/json", nil)
	if err != nil {
		log.Println(err)
	}
	return resp
}
