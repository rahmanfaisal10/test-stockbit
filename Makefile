#!/bin/bash

start-all:
	@make start-http
	@make start-grpc
	@make start-four

start-http:
	@go run ./cmd/server/http/main.go

start-grpc:
	@go run ./cmd/server/grpc/main.go

start-four:
	@go run ./cmd/server/number_4/main.go

