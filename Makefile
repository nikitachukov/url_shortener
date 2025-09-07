SHELL=/bin/bash
GOTOOLCHAIN=go1.24.5
SHORTAINER_BIN=bin/shoratiner

all: clean prep build autotest

clean:
	rm -rf bin/*

prep:
	GOTOOLCHAIN=$(GOTOOLCHAIN) go mod tidy

build:
	go build  -o $(SHORTAINER_BIN)  cmd/shortener/shoratiner.go

autotest:
	go test -v ./cmd/*