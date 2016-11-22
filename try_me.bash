# compile the jar (run from the host machine)
# sbt package

# upload the jar
curl --data-binary @/mounted/target/scala-2.11/simple-project_2.11-1.0.jar localhost:8090/jars/test

# create a new execution context
curl -d "" 'localhost:8090/contexts/test-context?num-cpu-cores=2&memory-per-node=1024m'

# execute the job
curl -d "" 'localhost:8090/jobs?appName=test&classPath=spark.jobserver.SimpleApp&context=test-context&sync=true'
