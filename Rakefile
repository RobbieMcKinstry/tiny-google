task build: [ :compass_build, :go_build ] do
end

task compass_build: [] do
    sh "compass compile frontend/static/sass/*"
end

task go_generate: [] do
    sh "goagen bootstrap -o src/ -d github.com/RobbieMcKinstry/neighborhoodwatch/design "
end

task go_build: [] do
    sh "GOOS=linux GOARCH=amd64 go build -o app src/main.go"
end

task clean: [] do
    sh "rm main"
end
