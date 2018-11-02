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

release:
	envchain goreleaser goreleaser release --rm-dist

demo: demo-prepare demo-exec

demo-prepare:
	-rm -rf /tmp/git-remind-demo
	mkdir -p /tmp/git-remind-demo
	cd /tmp/git-remind-demo && \
	git init myapp1 && \
	git init --bare myapp2-origin.git && \
	git clone /tmp/git-remind-demo/myapp2-origin.git myapp2 && \
	git init myapp3 && \
	cd /tmp/git-remind-demo/myapp1 && \
	touch unstaged-file
	cd /tmp/git-remind-demo/myapp2 && \
	git commit --allow-empty -m commit1 && \
	git push && \
	git commit --allow-empty -m commit2

demo-exec:
	@echo
	@echo "==== D E M O ===="
	@echo
	@set -eux && \
	export GIT_REMIND_PATHS='/tmp/git-remind-demo/*' && \
	git-remind paths && \
	git-remind status -a && \
	git-remind status-notification
