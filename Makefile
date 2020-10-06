# The file name of the binary to output
BINARY_FILENAME := docker-credential-acrenv
# The output directory
OUT_DIR := bin

all: clean bin

deps:
	@go get -u -t ./...

bin: deps
	@go build -i -o ${OUT_DIR}/${BINARY_FILENAME} main.go
	@echo Binary created: ${OUT_DIR}/${BINARY_FILENAME}

clean:
	@rm -rf ${OUT_DIR}
	@go clean

vet:
	@go vet ./...

lint:
	@echo 'Running golint...'
	@$(foreach src,$(SRCS),golint $(src);)

criticism: clean vet lint

fmt:
	@gofmt -w -s .

fix:
	@go fix ./...
