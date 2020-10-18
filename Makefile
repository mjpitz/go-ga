VERSION ?= local

build-deps:
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get -u oss.indeed.com/go/go-groups

deps:
	go mod download
	go mod verify

fmt:
	go-groups -w .
	gofmt -s -w .

test:
	go vet ./...
	golint -set_exit_status ./...
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

install:
	go install -ldflags="-X 'main.version=$(VERSION)'" ./cmd/beacon/

docker:
	docker build . \
		--build-arg VERSION=local \
		--tag mjpitz/beacon:latest \
		-f Dockerfile
