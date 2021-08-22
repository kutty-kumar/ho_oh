PWD ?= $(shell pwd)
GO_CODE_PATH ?= "github.com/kutty-kumar/ho_oh"
SOURCE_CODE_PATH ?= ${HOME}/go_proto_files

.PHONY: compile_go
compile_go:
	@mkdir ${SOURCE_CODE_PATH}
	@docker run --rm -v${SOURCE_CODE_PATH}:${SOURCE_CODE_PATH} -w ${SOURCE_CODE_PATH} thethingsindustries/protoc -I${PWD} --go_out="plugins=grpc,paths=source_relative:${SOURCE_CODE_PATH}" --grpc-gateway_out="logtostderr=true,allow_delete_body=true,paths=source_relative:${SOURCE_CODE_PATH}" core_v1/core.proto pikachu_v1/pikachu.proto bulbasur_v1/bulbasur.proto snorlax_v1/snorlax.proto \
	@echo $(shell ls ${SOURCE_CODE_PATH}/github.com) \
	cd ${SOURCE_CODE_PATH}/github.com \
	go mod init ${GO_CODE_PATH} \
	go mod tidy
