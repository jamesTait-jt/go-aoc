BINARY_NAME=main

build:
	go mod tidy
	./scripts/plugins/build

clean:
	./scripts/plugins/clean