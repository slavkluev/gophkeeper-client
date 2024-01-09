BUILD_VERSION := v1.0.0
BUILD_DATE := $(shell date +'%Y/%m/%d')
BUILD_COMMIT := $(shell git rev-parse HEAD)

build:
	go build -ldflags "-X client/config.BuildVersion=$(BUILD_VERSION) -X client/config.BuildDate=$(BUILD_DATE) -X client/config.BuildCommit=$(BUILD_COMMIT)" -o bin/gophkeeper-client main.go
