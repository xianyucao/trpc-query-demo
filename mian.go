package main

//启动go run . -conf config/trpc_go.yaml
//连接本地mysql，redis缓存，实现通过输入关键字实现查询

import (
    interceptors "trpc-query-demo/filter"
	"trpc-query-demo/proto/query"
	"trpc-query-demo/service"

	"trpc.group/trpc-go/trpc-go"
    "trpc.group/trpc-go/trpc-go/filter"
)

func main() {
    // 注册拦截器
	filter.Register("queryTiming", interceptors.QueryServerFilter, interceptors.QueryClientFilter)
	s := trpc.NewServer()
	query.RegisterQueryServiceService(s, &service.QueryServiceImpl{})
	_ = s.Serve()
}
