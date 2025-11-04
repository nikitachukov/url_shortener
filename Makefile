SHELL=/bin/bash
GOTOOLCHAIN=go1.24.5
SHORTENER_BIN=bin/shortener
SHORTENERTEST_BIN=tools/shortenertest
STATICTEST_BIN=tools/statictest
DATA_FILE=data.json
DSN="host=127.0.0.1 port=5432 user=user password=password dbname=db sslmode=disable"

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
up_migrations:
	migrate -path migrations/ -database "postgres://user:password@localhost:5432/db?sslmode=disable"  -verbose up
down_migrations:
	migrate -path migrations/ -database "postgres://user:password@localhost:5432/db?sslmode=disable"  -verbose down
TestIteration%: clean prep vet build myautotest
	$(SHORTENERTEST_BIN) -test.v -test.run=^$@$$ -binary-path=$(SHORTENER_BIN) -server-port=8888 -file-storage-path=$(DATA_FILE) -source-path="." -database-dsn=$(DSN)| tee >(richgo testfilter)
