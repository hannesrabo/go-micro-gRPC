#!/usr/bin/make

# SHELL := /bin/bash

.PHONY: build server protobuff

# default: all
# all: build

protobuff:
	./build/scripts/build-protobuff.sh

build: protobuff
	go build ./

server: protobuff
	go build -o ./server/server ./server
