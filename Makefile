SHELL=/bin/bash
GOTOOLCHAIN=go1.24.5
SHORTENER_BIN=bin/shortener
SHORTENERTEST_BIN=tools/shortenertest
STATICTEST_BIN=tools/statictest
all: clean prep build autotest

clean:
	rm -rf bin/*

prep:
	GOTOOLCHAIN=$(GOTOOLCHAIN) go mod tidy

build:
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go

autotest:
	go test -v ./cmd/*


TestIteration1:
	go vet -vettool=$(STATICTEST_BIN) ./...
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	$(SHORTENERTEST_BIN) -test.v -test.run=^TestIteration1$ -binary-path=$(SHORTENER_BIN) -server-port=8888

TestIteration2:
	go vet -vettool=$(STATICTEST_BIN) ./...
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	$(SHORTENERTEST_BIN) -test.v -test.run=^TestIteration2$ -binary-path=$(SHORTENER_BIN) -server-port=8888

TestIteration3:
	go vet -vettool=$(STATICTEST_BIN) ./...
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	$(SHORTENERTEST_BIN) -test.v -test.run=^TestIteration3$ -binary-path=$(SHORTENER_BIN) -server-port=8888

TestIteration4:
	go vet -vettool=$(STATICTEST_BIN) ./...
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	$(SHORTENERTEST_BIN) -test.v -test.run=^TestIteration4$ -binary-path=$(SHORTENER_BIN) -server-port=8888

TestIteration5:
	go vet -vettool=$(STATICTEST_BIN) ./...
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	$(SHORTENERTEST_BIN) -test.v -test.run=^TestIteration5$$ -binary-path=$(SHORTENER_BIN) -server-port=8888

TestIteration6:
	go vet -vettool=$(STATICTEST_BIN) ./...
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	$(SHORTENERTEST_BIN) -test.v -test.run=^TestIteration6$$ -binary-path=$(SHORTENER_BIN) -server-port=8888

TestIteration7:
	go vet -vettool=$(STATICTEST_BIN) ./...
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go
	$(SHORTENERTEST_BIN) -test.v -test.run=^TestIteration7$$ -binary-path=$(SHORTENER_BIN) -server-port=8080 -file-storage-path="zzz" -source-path="."

