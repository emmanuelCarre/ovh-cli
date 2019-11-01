PROJECT_NAME 	:= ovh-cli
PKG 			:= github.com/emmanuelCarre/ovh-cli
OUTPUT_DIR 		:= output
GOARCH 			:= $(shell go env GOARCH)
GOOS 			:= $(shell go env GOOS)
GO_VERSION 		:= $(shell go version | awk '{print $$3}' )
APP_VERSION 	:= $(shell git describe --tag >/dev/null 2>&1; if [ $$? -ne 0 ]; then git rev-parse --short HEAD; else git describe --tag; fi)
GIT_COMMIT 		:= $(shell git rev-parse HEAD)
BUILD_DATE 		:= $(shell date '+%Y-%m-%d_%H:%M:%S' )

LDFLAGS = '-X ${PKG}/cmd.Version=${APP_VERSION} -X ${PKG}/cmd.GoVersion=${GO_VERSION} -X ${PKG}/cmd.OsArchi=${GOOS}/${GOARCH} -X ${PKG}/cmd.GitCommit=${GIT_COMMIT} -X ${PKG}/cmd.BuildDate=${BUILD_DATE}'

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Build ${PROJECT_NAME}"
	@go build -o ${OUTPUT_DIR}/${PROJECT_NAME} -ldflags ${LDFLAGS}

.PHONY: clean
clean:
	rm -rf ${OUTPUT_DIR}
