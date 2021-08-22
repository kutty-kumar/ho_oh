PWD ?= $(shell pwd)
GO_CODE_PATH ?= "github.com/kutty-kumar/ho_oh"
SOURCE_CODE_PATH ?= ${HOME}/go_proto_files

.PHONY: compile_go
compile_go:
	$(shell mkdir ${SOURCE_CODE_PATH}) \
	docker run --rm --user $(id -u):$(id -g) -v${PWD}:${PWD} -v${SOURCE_CODE_PATH}:${SOURCE_CODE_PATH} -w${SOURCE_CODE_PATH} thethingsindustries/protoc:latest --proto_path=${PWD} --go-grpc_out=${SOURCE_CODE_PATH} --grpc-gateway_out="logtostderr=true,allow_delete_body=true:." -I/home/kumard/go/src/ho_oh -I/usr/include/github.com/gogo/protobuf ${PWD}/core_v1/core.proto ${PWD}/pikachu_v1/pikachu.proto ${PWD}/bulbasur_v1/bulbasur.proto ${PWD}/ditto_v1/ditto.proto ${PWD}/snorlax_v1/snorlax.proto
	$(shell cd ${SOURCE_CODE_PATH}/github.com/kutty-kumar/ho_oh) \
	go mod init ${GO_CODE_PATH} \
	go mod tidy
