PWD ?= $(shell pwd)
GO_CODE_PATH ?= "github.com/kutty-kumar/ho_oh"

.PHONY: compile_go
compile_go:
	docker run --rm -v${PWD}:${PWD} -w /tmp thethingsindustries/protoc -I${PWD} --go_out="plugins=grpc,paths=source_relative:." --grpc-gateway_out="logtostderr=true,allow_delete_body=true,paths=source_relative:." core_v1/core.proto pikachu_v1/pikachu.proto bulbasur_v1/bulbasur.proto snorlax_v1/snorlax.proto \
	cd ${GO_CODE_PATH} \
	go mod init ${GO_CODE_PATH} \
	go mod tidy
