#!/bin/bash

set -xe

DIR="$( cd "$( dirname "$BASH_SOURCE[0]" )"/.. && pwd )"
cd ${DIR}/protos

grpc_tools_node_protoc --js_out=import_style=commonjs,binary:../node --grpc_out=../node *.proto