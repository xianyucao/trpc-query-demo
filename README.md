1：定义protobuf协议，定义查询服务的protobuf文件query.proto：

使用tRPC提供的工具生成对应的Go代码。

option go_package = "trpc-query-demo/proto/query";
//trpc create -f --protofile=query.proto --protocol=trpc --rpconly --mock=false

2: 业务逻辑实现
在service/query_service.go中

3：服务注册与启动
在main.go中注册tRPC服务并启动，tRPC_go.yaml服务配置

go mod tidy

//启动go run . -conf config/trpc_go.yaml
//连接本地mysql，redis缓存，实现通过输入关键字实现查询
//访问方式如：curl "http://localhost:8000/demo.query.QueryService/Query?queryText=安置"