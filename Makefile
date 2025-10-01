SHELL=/bin/bash
GOTOOLCHAIN=go1.24.5
SHORTENER_BIN=bin/shortener

all: clean prep build autotest

clean:
	rm -rf bin/*

prep:
	GOTOOLCHAIN=$(GOTOOLCHAIN) go mod tidy

build:
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go

autotest:
	go test -v ./cmd/*

at:
	 ./shortenertest -test.v -test.run=^TestIteration1$ -binary-path=$(SHORTENER_BIN)