// Code generated by trpc-go/trpc-cmdline v1.0.7. DO NOT EDIT.
// source: query.proto

package query

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// QueryServiceService defines service.
type QueryServiceService interface {
	Query(ctx context.Context, req *QueryRequest) (*QueryResponse, error)
}

func QueryServiceService_Query_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(QueryServiceService).Query(ctx, reqbody.(*QueryRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// QueryServiceServer_ServiceDesc descriptor for server.RegisterService.
var QueryServiceServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "demo.query.QueryService",
	HandlerType: ((*QueryServiceService)(nil)),
	Methods: []server.Method{
		{
			Name: "/demo.query.QueryService/Query",
			Func: QueryServiceService_Query_Handler,
		},
	},
}

// RegisterQueryServiceService registers service.
func RegisterQueryServiceService(s server.Service, svr QueryServiceService) {
	if err := s.Register(&QueryServiceServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("QueryService register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedQueryService struct{}

func (s *UnimplementedQueryService) Query(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, errors.New("rpc Query of service QueryService is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// QueryServiceClientProxy defines service client proxy
type QueryServiceClientProxy interface {
	Query(ctx context.Context, req *QueryRequest, opts ...client.Option) (rsp *QueryResponse, err error)
}

type QueryServiceClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewQueryServiceClientProxy = func(opts ...client.Option) QueryServiceClientProxy {
	return &QueryServiceClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *QueryServiceClientProxyImpl) Query(ctx context.Context, req *QueryRequest, opts ...client.Option) (*QueryResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/demo.query.QueryService/Query")
	msg.WithCalleeServiceName(QueryServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("QueryService")
	msg.WithCalleeMethod("Query")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END