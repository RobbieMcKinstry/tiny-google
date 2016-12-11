# compile the jar (run from the host machine)
rake build

# launch the docker containers (run from the host machine)
rake run

# ssh into the spark box (run from the host machine)
docker exec -it finalproject_spark_1 /bin/bash

# now you're inside the docker container

# upload the jar
curl --data-binary @/mounted/target/scala-2.11/simple-project_2.11-1.0.jar localhost:8090/jars/test

# create a new execution context
curl -d "" 'localhost:8090/contexts/test-context?num-cpu-cores=2&memory-per-node=1024m'

# execute the job
curl -d "" 'localhost:8090/jobs?appName=test&classPath=spark.jobserver.SimpleApp&context=test-context&sync=true'
