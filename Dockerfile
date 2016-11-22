FROM alpine:3.2
ADD app /go/bin/app
ENTRYPOINT /go/bin/app
