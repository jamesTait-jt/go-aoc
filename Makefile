BINARY_NAME=main

build:
	go mod tidy
	gofmt -s -w ./..
	./scripts/plugins/build

clean:
	./scripts/plugins/clean