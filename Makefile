.PHONY: clean check build.local build.linux build.osx

BINARY      ?= go-monit
VERSION     ?= $(shell git describe --tags --always --dirty)
SOURCES     = $(shell find . -name '*.go')
GOPKGS      = $(shell go list ./...)
BUILD_FLAGS ?= -v
LDFLAGS     ?= -X main.version=$(VERSION) -w -s

default: build.local

clean:
	rm -rf build
	rm -rf *.tar.gz

test:
	go test -v $(GOPKGS)

test-with-coverage:
	go test -coverprofile=coverage.out -v $(GOPKGS)

check:
	golint $(GOPKGS)
	go vet -v $(GOPKGS)

build.local: build/$(BINARY)
build.linux: build/linux/$(BINARY)

build/$(BINARY): $(SOURCES)
	go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" ./cmd/$(BINARY)

build/linux/$(BINARY): $(SOURC`ES)
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/linux/$(BINARY) -ldflags "$(LDFLAGS)" ./cmd/$(BINARY)

install: default
	go install