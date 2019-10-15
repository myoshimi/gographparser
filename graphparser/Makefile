#!/bin/env make

GO=/usr/local/go/bin/go
# GOPATH=/home

test:
	$(GO) test -v -cover -coverprofile=./cover.out
	$(GO) tool cover -html=cover.out -o cover.html

clean:
	rm -f cover.html cover.out
