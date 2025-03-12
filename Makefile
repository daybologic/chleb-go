#!/usr/bin/env make -f

all: bible-votd

bible-votd:
	go build
	upx bible-votd

run : all
	./bible-votd

clean:
	go clean

.PHONY: all run clean
