#!/usr/bin/env make -f

all: bible-votd

bible-votd:
	go build
	strip $@

bible-votd.exe:
	env GOOS=windows GOARCH=amd64 go build
	upx -q -9 $@

run: all
	./bible-votd

clean:
	go clean
	env GOOS=windows GOARCH=amd64 go clean

.PHONY: all run clean bible-votd
