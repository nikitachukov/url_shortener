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

TestIteration4:
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	./shortenertest -test.v -test.run=^TestIteration4$ -binary-path=$(SHORTENER_BIN) -server-port=8888


TestIteration5:
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	./shortenertest -test.v -test.run=^TestIteration5$$ -binary-path=$(SHORTENER_BIN) -server-port=8888



TestIteration6:
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	./shortenertest -test.v -test.run=^TestIteration6$$ -binary-path=$(SHORTENER_BIN) -server-port=8888



TestIteration7:
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	./shortenertest -test.v -test.run=^TestIteration7$$ -binary-path=$(SHORTENER_BIN) -server-port=8080 -file-storage-path="zzz" -source-path="."

