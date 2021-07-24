FROM golang:1.15.2 as base
LABEL author=umeshkumhar

# Copy in the go src
WORKDIR /go/src/github.com/umeshkumhar/logging-golang-demo

# Copy the Go manifests
COPY logging.go logging.go
RUN go get github.com/sirupsen/logrus
RUN GOOS=linux GOARCH=amd64 go build -a -o /opt/app github.com/umeshkumhar/logging-golang-demo
RUN chmod -R g+rwX /opt/app
USER 1001
ENTRYPOINT ["/opt/app"]
