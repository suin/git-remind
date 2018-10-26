APP_PACKAGE := github.com/suin/git-remind/app/cli
APP_NAME := git-remind
APP_DESCRIPTION := Never forget to git commit and push
APP_VERSION := $(shell git describe --tag)
BUILD_OPTION := -ldflags "\
    -X \"$(APP_PACKAGE).Name=$(APP_NAME)\"\
    -X \"$(APP_PACKAGE).Description=$(APP_DESCRIPTION)\"\
    -X \"$(APP_PACKAGE).Version=$(APP_VERSION)\""

run: fmt
	go run $(BUILD_OPTION) main.go ${args}

help: fmt
	go run $(BUILD_OPTION) main.go --help

fmt:
	go fmt ./...

build: fmt
	go build $(BUILD_OPTION) -o bin/$(APP_NAME) main.go

install: build
	cp bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)
