#!/bin/bash

set -xe

DIR="$( cd "$( dirname "$BASH_SOURCE[0]" )"/.. && pwd )"
cd ${DIR}

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc

mkdir -p ${GOPATH}/src/git.meideng.net/sempr
rm -f ${GOPATH}/src/git.meideng.net/sempr/grpc-talks
ln -sf ${DIR} ${GOPATH}/src/git.meideng.net/sempr/grpc-talks