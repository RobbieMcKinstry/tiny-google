task default: [ :build, :run, :clean ] do
end

task build: [ :scala_build, :go_build ] do
end

task go_build: [] do
    sh "go build -o main src/golang/*.go"
end

task compass_build: [] do
    sh "compass compile src/golang/static/sass/*"
end

task scala_build: [] do
    sh "sbt package"
end

task run: [ ] do
    sh "docker-compose build"
    sh "docker-compose up"
end

task clean: [] do
    sh "rm main"
end
