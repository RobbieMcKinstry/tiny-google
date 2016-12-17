package main

const (
	// StaticDir points to the static directory
	StaticDir = "/static/"

	// Port is the port number
	Port      = ":8080"
	indexPath = "src/InvertedIndex.json"
)

type (
	documentData struct {
		DocumentPath string
		DocumentName string
		Frequency    float64
	}

	invertedIndex struct {
		Time  string
		Links map[string][]documentData
	}
)
