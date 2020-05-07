include .env

.EXPORT_ALL_VARIABLES:

dev:
	go run *.go

build:
	CGO_ENABLED=0 go build -o ./bin/cda *.go

update:
	go get ./...
	go mod tidy

.PHONY: dev build update
