#!/bin/bash

set -xe

DIR="$( cd "$( dirname "$BASH_SOURCE[0]" )"/.. && pwd )"
cd ${DIR}

python -m grpc_tools.protoc -I protos --python_out=python --grpc_python_out=python protos/*.proto
