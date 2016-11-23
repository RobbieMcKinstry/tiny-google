FROM golang:1.7
COPY ./src/golang /go/src/app
ADD  ./src/golang/static /static
#RUN  ["go", "get",     "-d", "-v"]
RUN  ["go", "build",   "/go/src/app/main.go"]
RUN  ["go", "install", "-v", "app"]
CMD ["app"]
