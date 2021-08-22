PWD ?= $(shell pwd)
GO_CODE_PATH ?= "github.com/kutty-kumar/ho_oh"
SOURCE_CODE_PATH ?= "${HOME}/go_proto_files"

.PHONY: compile_go
compile_go:
	. ./build.sh; build

.PHONY: push_go
push_go:
	. ./build.sh; push
