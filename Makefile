SHELL=/bin/bash
GOTOOLCHAIN=go1.24.5
SHORTENER_BIN=bin/shortener
SHORTENERTEST_BIN=tools/shortenertest
STATICTEST_BIN=tools/statictest

myautotest:
	go test -v ./cmd/*

clean:
	rm -rf bin/*
prep:
	GOTOOLCHAIN=$(GOTOOLCHAIN) go mod tidy
vet:
	go vet -vettool=$(STATICTEST_BIN) ./...
build:
	go build  -o $(SHORTENER_BIN) cmd/shortener/shortener.go

TestIteration1: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)

TestIteration2: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)

TestIteration3: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)

TestIteration4: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)

TestIteration5: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)

TestIteration6: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)

TestIteration7: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)

TestIteration8: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path="zzz" -source-path="." | tee >(richgo testfilter)
