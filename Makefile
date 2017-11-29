dep:
	dep ensure
	dep prune

lint:
	-errcheck `go list ./... | grep -v vendor`
	-go vet `go list ./... | grep -v vendor`
	golint `go list ./... | grep -v vendor`

test:
	go test -race `go list ./... | grep -v vendor`

build:
	docker build -t willroberts/perandus .

run:
	docker run willroberts/perandus

shell:
	docker run -ti willroberts/perandus
