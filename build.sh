#!/bin/bash

export SOURCE_CODE_PATH=$HOME/go_proto_files;
export GO_CODE_PATH="github.com/kutty-kumar/ho_oh_go";
export PWD=`pwd`;

build() {
  [ ! -d $SOURCE_CODE_PATH ] && mkdir -p $SOURCE_CODE_PATH;
  docker run --rm --user $(id -u):$(id -g) -v${PWD}:${PWD} -v${SOURCE_CODE_PATH}:${SOURCE_CODE_PATH} -w${SOURCE_CODE_PATH} thethingsindustries/protoc:latest --proto_path=${PWD} --go-grpc_out=${SOURCE_CODE_PATH} --grpc-gateway_out="logtostderr=true,allow_delete_body=true:." --swagger_out="allow_delete_body=true:." -I/home/kumard/go/src/ho_oh -I/usr/include/github.com/gogo/protobuf ${PWD}/core_v1/core.proto ${PWD}/pikachu_v1/pikachu.proto ${PWD}/bulbasur_v1/bulbasur.proto ${PWD}/ditto_v1/ditto.proto ${PWD}/snorlax_v1/snorlax.proto;
  cd $SOURCE_CODE_PATH/$GO_CODE_PATH;
  go mod init $GO_CODE_PATH;
  go mod tidy;
}

push() {
  cd $SOURCE_CODE_PATH/$GO_CODE_PATH;
  git init;
  git remote add origin $ORIGIN_PATH;
  git config user.name kutty-kumar;
  git config user.email kumar.varalakshmi@outlook.com;
  git add .;
  git commit -m "compiled protobuf code";
  git fetch --all;
  git branch --set-upstream-to=origin/master master;
  git pull --rebase;
  tag=$(git describe --always --tags || echo pre-commit);
  git tag $tag;
  git push --atomic origin HEAD $tag;
}

build_and_push(){
  build;
  push;
}
