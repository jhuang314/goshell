### Makefile --- Makefile for gosh

## Author: Jesse Huang
## Version: 0.1

export GOPATH:=$(shell pwd)

BUILDTAGS=debug
default: all


fmt:
	go fmt ./...

main: lex
	go install -tags '$(BUILDTAGS)' gosh

lex:
	cd src/parser && golex -o scanner.go scanner.l 

all: fmt main

run:
	./bin/gosh

clean:
	go clean -i -r gosh/...

