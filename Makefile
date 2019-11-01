PROJECT_NAME := ovh-cli
PKG := github.com/emmanuelCarre/ovh-cli
OUTPUT_DIR := output

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Build ${PROJECT_NAME}"
	@go build -o ${OUTPUT_DIR}/${PROJECT_NAME}


.PHONY: clean
clean:
	rm -rf ${OUTPUT_DIR}
