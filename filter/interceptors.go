package interceptors

import (
	"context"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-go/filter"
)

// QueryServerFilter 是服务端耗时统计的拦截器
func QueryServerFilter(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error) {
	begin := time.Now() // 请求处理开始前打点

	rsp, err = next(ctx, req) // 执行下一个拦截器或请求处理

	cost := time.Since(begin)                             // 计算耗时
	log.Printf("Query service took %v to complete", cost) // 日志记录耗时信息

	return rsp, err
}

// QueryClientFilter 是客户端耗时统计的拦截器
func QueryClientFilter(ctx context.Context, req, rsp interface{}, next filter.ClientHandleFunc) error {
	begin := time.Now() // 发起请求前打点

	err := next(ctx, req, rsp) // 调用下一个拦截器或发起请求

	cost := time.Since(begin)                            // 计算耗时
	log.Printf("Client query took %v to complete", cost) // 日志记录耗时信息

	return err
}
