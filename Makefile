#!/usr/bin/env bash

test:
	go test -v ./...

clean:
	rm -rf ./bin/*

build: clean
	go build -o ./bin/alien-invasion ./cmd

run: clean build
	./bin/alien-invasion game --alien-count=2 --world-map=./data/world-map.txt
