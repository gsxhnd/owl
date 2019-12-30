SHELL := /bin/bash
BASEDIR = $(shell pwd)
APPS = gecko
BuildDIR = build



windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gormt.exe main.go
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gormt main.go

linux:export GOOS=linux
linux:export GOARCH=amd64
linux:
	@go build -v -ldflags ${ldflags}  -o $@ ./apps/$*


all: $(APPS)

$(APPS): %:$(BuildDIR)/%

$(BuildDIR)/%:
	@go build -v -o $@ ./apps/$*