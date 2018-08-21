#!/bin/bash

set -xe

DIR="$( cd "$( dirname "$BASH_SOURCE[0]" )"/.. && pwd )"
cd ${DIR}

mkdir -p go/math
protoc -I protos/ protos/math.proto --go_out=plugins=grpc:go/math
