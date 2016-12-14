package main

const (
	// StaticDir points to the static directory
	StaticDir = "/static/"

	// Port is the port number
	Port = ":8080"

	endpointBase       = "http://spark:8090"
	jarUploadExtension = "/jars/test"
	jarUploadEndpoint  = endpointBase + jarUploadExtension

	jobExtension = "/jobs?appName=test&classPath=spark.jobserver.SimpleApp&context=test-context&sync=true"
	jobEndpoint  = endpointBase + jobExtension

	contextExtension = "/contexts/test-context?num-cpu-cores=2&memory-per-node=1024m"
	contextEndpoint  = endpointBase + contextExtension

	executeURI  = endpointBase + "/jobs?appName=test&classPath=spark.jobserver.SimpleApp&context=test-context&sync=true"
	jarFilePath = "/mounted/target/scala-2.11/simple-project_2.11-1.0.jar"
)

type (
	documentData struct {
		DocumentPath string
		DocumentName string
		Frequency    float64
	}

	invertedIndex map[string]documentData
)
