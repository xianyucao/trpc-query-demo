package main

//启动go run . -conf config/trpc_go.yaml
//连接本地mysql，redis缓存，实现通过输入关键字实现查询
//访问方式如：curl "http://localhost:8000/demo.query.QueryService/Query?queryText=安置"
import (
	"trpc-query-demo/proto/query"
	"trpc-query-demo/service"

	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	query.RegisterQueryServiceService(s, &service.QueryServiceImpl{})
	_ = s.Serve()
}
