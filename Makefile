#!/bin/bash

start-all:
	@make start-http

start-http:
	@go run ./cmd/server/http/main.go

#start-grpc:
#	@go build -o bin/user ./cmd/user
#	@bin/user
