.PHONY: default run build test docs clean

APP_NAME=go_api

default: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@if exist $(APP_NAME).exe del /f /q $(APP_NAME).exe
	@if exist docs rmdir /s /q docs