BINARY_NAME=urlshortener

DEFAULT_GOAL:=run

build:test
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux main.go

run: build,
	
	go run main.go

test:test_coverage
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

all:test build run