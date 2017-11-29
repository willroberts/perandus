FROM golang:1.9.2

COPY . /go/src/github.com/willroberts/perandus
WORKDIR /go/src/github.com/willroberts/perandus/server

EXPOSE 8000
ENTRYPOINT ["go", "run", "main.go"]
