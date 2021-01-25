BINARY_NAME := 308reedpipes

.PHONY: all
all: build

build:
	go build -o ${BINARY_NAME}