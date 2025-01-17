DEFAULT_TARGET: build

.PHONY: vet test run build

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

build: vet
	@go build -o bin/converter.exe

run:
	@./bin/converter

test:
	@go test -v ./...
