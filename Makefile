BUILD_VERSION := v1.0.0
BUILD_DATE := $(shell date +'%Y/%m/%d')
BUILD_COMMIT := $(shell git rev-parse HEAD)

FLAGS := "-X client/config.BuildVersion=$(BUILD_VERSION) -X client/config.BuildDate=$(BUILD_DATE) -X client/config.BuildCommit=$(BUILD_COMMIT)"

build:
	GOOS=linux GOARCH=amd64 go build -ldflags $(FLAGS) -o bin/gophkeeper-client-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -ldflags $(FLAGS) -o bin/gophkeeper-client-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags $(FLAGS) -o bin/gophkeeper-client-darwin-amd64 main.go
