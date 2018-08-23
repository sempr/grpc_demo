### 从*.proto文件生成*pb.js和*grpc_pb.js文件

``` bash
npm install -g grpc-tools
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:../node --grpc_out=../node --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` math.proto
```

或直接运行

``` bash
bin/gen_node.sh
```

### 运行

#### 准备运行环境

```
npm i
```

### run server

```
node math_server.js
```

### run client

```
node math_client.js
```
