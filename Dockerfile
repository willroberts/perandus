FROM golang:1.9.2

COPY . /go/src/github.com/willroberts/perandus
WORKDIR /go/src/github.com/willroberts/perandus

ENTRYPOINT ["go", "run", "main.go"]
