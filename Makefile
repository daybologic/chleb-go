#!/usr/bin/env make -f

all: hello

hello:
	go build

run : all
	./hello

clean:
	go clean

.PHONY: all run clean
