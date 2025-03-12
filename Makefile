#!/usr/bin/env make -f

all: bible-votd

bible-votd:
	go build

run : all
	./bible-votd

clean:
	go clean

.PHONY: all run clean
