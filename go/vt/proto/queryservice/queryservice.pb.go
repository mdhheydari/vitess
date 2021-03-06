// Code generated by protoc-gen-go.
// source: queryservice.proto
// DO NOT EDIT!

/*
Package queryservice is a generated protocol buffer package.

It is generated from these files:
	queryservice.proto

It has these top-level messages:
*/
package queryservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import query "github.com/youtube/vitess/go/vt/proto/query"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Query service

type QueryClient interface {
	// Execute executes the specified SQL query (might be in a
	// transaction context, if Query.transaction_id is set).
	Execute(ctx context.Context, in *query.ExecuteRequest, opts ...grpc.CallOption) (*query.ExecuteResponse, error)
	// ExecuteBatch executes a list of queries, and returns the result
	// for each query.
	ExecuteBatch(ctx context.Context, in *query.ExecuteBatchRequest, opts ...grpc.CallOption) (*query.ExecuteBatchResponse, error)
	// StreamExecute executes a streaming query. Use this method if the
	// query returns a large number of rows. The first QueryResult will
	// contain the Fields, subsequent QueryResult messages will contain
	// the rows.
	StreamExecute(ctx context.Context, in *query.StreamExecuteRequest, opts ...grpc.CallOption) (Query_StreamExecuteClient, error)
	// Begin a transaction.
	Begin(ctx context.Context, in *query.BeginRequest, opts ...grpc.CallOption) (*query.BeginResponse, error)
	// Commit a transaction.
	Commit(ctx context.Context, in *query.CommitRequest, opts ...grpc.CallOption) (*query.CommitResponse, error)
	// Rollback a transaction.
	Rollback(ctx context.Context, in *query.RollbackRequest, opts ...grpc.CallOption) (*query.RollbackResponse, error)
	// BeginExecute executes a begin and the specified SQL query.
	BeginExecute(ctx context.Context, in *query.BeginExecuteRequest, opts ...grpc.CallOption) (*query.BeginExecuteResponse, error)
	// BeginExecuteBatch executes a begin and a list of queries.
	BeginExecuteBatch(ctx context.Context, in *query.BeginExecuteBatchRequest, opts ...grpc.CallOption) (*query.BeginExecuteBatchResponse, error)
	// SplitQuery is the API to facilitate MapReduce-type iterations
	// over large data sets (like full table dumps).
	SplitQuery(ctx context.Context, in *query.SplitQueryRequest, opts ...grpc.CallOption) (*query.SplitQueryResponse, error)
	// StreamHealth runs a streaming RPC to the tablet, that returns the
	// current health of the tablet on a regular basis.
	StreamHealth(ctx context.Context, in *query.StreamHealthRequest, opts ...grpc.CallOption) (Query_StreamHealthClient, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Execute(ctx context.Context, in *query.ExecuteRequest, opts ...grpc.CallOption) (*query.ExecuteResponse, error) {
	out := new(query.ExecuteResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ExecuteBatch(ctx context.Context, in *query.ExecuteBatchRequest, opts ...grpc.CallOption) (*query.ExecuteBatchResponse, error) {
	out := new(query.ExecuteBatchResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/ExecuteBatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StreamExecute(ctx context.Context, in *query.StreamExecuteRequest, opts ...grpc.CallOption) (Query_StreamExecuteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Query_serviceDesc.Streams[0], c.cc, "/queryservice.Query/StreamExecute", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryStreamExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_StreamExecuteClient interface {
	Recv() (*query.StreamExecuteResponse, error)
	grpc.ClientStream
}

type queryStreamExecuteClient struct {
	grpc.ClientStream
}

func (x *queryStreamExecuteClient) Recv() (*query.StreamExecuteResponse, error) {
	m := new(query.StreamExecuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) Begin(ctx context.Context, in *query.BeginRequest, opts ...grpc.CallOption) (*query.BeginResponse, error) {
	out := new(query.BeginResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/Begin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Commit(ctx context.Context, in *query.CommitRequest, opts ...grpc.CallOption) (*query.CommitResponse, error) {
	out := new(query.CommitResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/Commit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Rollback(ctx context.Context, in *query.RollbackRequest, opts ...grpc.CallOption) (*query.RollbackResponse, error) {
	out := new(query.RollbackResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/Rollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BeginExecute(ctx context.Context, in *query.BeginExecuteRequest, opts ...grpc.CallOption) (*query.BeginExecuteResponse, error) {
	out := new(query.BeginExecuteResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/BeginExecute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BeginExecuteBatch(ctx context.Context, in *query.BeginExecuteBatchRequest, opts ...grpc.CallOption) (*query.BeginExecuteBatchResponse, error) {
	out := new(query.BeginExecuteBatchResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/BeginExecuteBatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SplitQuery(ctx context.Context, in *query.SplitQueryRequest, opts ...grpc.CallOption) (*query.SplitQueryResponse, error) {
	out := new(query.SplitQueryResponse)
	err := grpc.Invoke(ctx, "/queryservice.Query/SplitQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StreamHealth(ctx context.Context, in *query.StreamHealthRequest, opts ...grpc.CallOption) (Query_StreamHealthClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Query_serviceDesc.Streams[1], c.cc, "/queryservice.Query/StreamHealth", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryStreamHealthClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_StreamHealthClient interface {
	Recv() (*query.StreamHealthResponse, error)
	grpc.ClientStream
}

type queryStreamHealthClient struct {
	grpc.ClientStream
}

func (x *queryStreamHealthClient) Recv() (*query.StreamHealthResponse, error) {
	m := new(query.StreamHealthResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Query service

type QueryServer interface {
	// Execute executes the specified SQL query (might be in a
	// transaction context, if Query.transaction_id is set).
	Execute(context.Context, *query.ExecuteRequest) (*query.ExecuteResponse, error)
	// ExecuteBatch executes a list of queries, and returns the result
	// for each query.
	ExecuteBatch(context.Context, *query.ExecuteBatchRequest) (*query.ExecuteBatchResponse, error)
	// StreamExecute executes a streaming query. Use this method if the
	// query returns a large number of rows. The first QueryResult will
	// contain the Fields, subsequent QueryResult messages will contain
	// the rows.
	StreamExecute(*query.StreamExecuteRequest, Query_StreamExecuteServer) error
	// Begin a transaction.
	Begin(context.Context, *query.BeginRequest) (*query.BeginResponse, error)
	// Commit a transaction.
	Commit(context.Context, *query.CommitRequest) (*query.CommitResponse, error)
	// Rollback a transaction.
	Rollback(context.Context, *query.RollbackRequest) (*query.RollbackResponse, error)
	// BeginExecute executes a begin and the specified SQL query.
	BeginExecute(context.Context, *query.BeginExecuteRequest) (*query.BeginExecuteResponse, error)
	// BeginExecuteBatch executes a begin and a list of queries.
	BeginExecuteBatch(context.Context, *query.BeginExecuteBatchRequest) (*query.BeginExecuteBatchResponse, error)
	// SplitQuery is the API to facilitate MapReduce-type iterations
	// over large data sets (like full table dumps).
	SplitQuery(context.Context, *query.SplitQueryRequest) (*query.SplitQueryResponse, error)
	// StreamHealth runs a streaming RPC to the tablet, that returns the
	// current health of the tablet on a regular basis.
	StreamHealth(*query.StreamHealthRequest, Query_StreamHealthServer) error
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Execute(ctx, req.(*query.ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ExecuteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ExecuteBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ExecuteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/ExecuteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ExecuteBatch(ctx, req.(*query.ExecuteBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StreamExecute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(query.StreamExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).StreamExecute(m, &queryStreamExecuteServer{stream})
}

type Query_StreamExecuteServer interface {
	Send(*query.StreamExecuteResponse) error
	grpc.ServerStream
}

type queryStreamExecuteServer struct {
	grpc.ServerStream
}

func (x *queryStreamExecuteServer) Send(m *query.StreamExecuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_Begin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.BeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Begin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Begin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Begin(ctx, req.(*query.BeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Commit(ctx, req.(*query.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.RollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Rollback(ctx, req.(*query.RollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BeginExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.BeginExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BeginExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/BeginExecute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BeginExecute(ctx, req.(*query.BeginExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BeginExecuteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.BeginExecuteBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BeginExecuteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/BeginExecuteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BeginExecuteBatch(ctx, req.(*query.BeginExecuteBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SplitQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.SplitQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SplitQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/SplitQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SplitQuery(ctx, req.(*query.SplitQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StreamHealth_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(query.StreamHealthRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).StreamHealth(m, &queryStreamHealthServer{stream})
}

type Query_StreamHealthServer interface {
	Send(*query.StreamHealthResponse) error
	grpc.ServerStream
}

type queryStreamHealthServer struct {
	grpc.ServerStream
}

func (x *queryStreamHealthServer) Send(m *query.StreamHealthResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "queryservice.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Query_Execute_Handler,
		},
		{
			MethodName: "ExecuteBatch",
			Handler:    _Query_ExecuteBatch_Handler,
		},
		{
			MethodName: "Begin",
			Handler:    _Query_Begin_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Query_Commit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _Query_Rollback_Handler,
		},
		{
			MethodName: "BeginExecute",
			Handler:    _Query_BeginExecute_Handler,
		},
		{
			MethodName: "BeginExecuteBatch",
			Handler:    _Query_BeginExecuteBatch_Handler,
		},
		{
			MethodName: "SplitQuery",
			Handler:    _Query_SplitQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamExecute",
			Handler:       _Query_StreamExecute_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamHealth",
			Handler:       _Query_StreamHealth_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("queryservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xf5, 0x90, 0x2a, 0x63, 0x3c, 0x38, 0x5a, 0xff, 0xa4, 0x82, 0xe2, 0x03, 0x14, 0x51,
	0x41, 0x10, 0xbc, 0xb4, 0x08, 0x8a, 0x20, 0xd8, 0x5e, 0xbc, 0xa6, 0x61, 0xd0, 0x60, 0xd2, 0x8d,
	0xc9, 0x54, 0xf4, 0xf9, 0x7c, 0x31, 0x75, 0x37, 0xb3, 0xdd, 0x8d, 0xeb, 0x71, 0x7e, 0xdf, 0x37,
	0x1f, 0xb3, 0x3b, 0x03, 0xf8, 0xb6, 0xa0, 0xfa, 0xb3, 0xa1, 0xfa, 0x3d, 0xcf, 0x68, 0x58, 0xd5,
	0x8a, 0x15, 0xc6, 0x2e, 0x4b, 0x36, 0x74, 0x65, 0xa4, 0xb3, 0xaf, 0x08, 0xa2, 0xc7, 0xdf, 0x1a,
	0xaf, 0x60, 0xed, 0xe6, 0x83, 0xb2, 0x05, 0x13, 0xf6, 0x87, 0xc6, 0xd2, 0xd6, 0x13, 0xfa, 0x29,
	0x1b, 0x4e, 0x76, 0xbb, 0xb8, 0xa9, 0xd4, 0xbc, 0xa1, 0x93, 0x15, 0xbc, 0x83, 0xb8, 0x85, 0xa3,
	0x94, 0xb3, 0x17, 0x4c, 0x7c, 0xa7, 0x86, 0x92, 0x32, 0x08, 0x6a, 0x36, 0xea, 0x01, 0x36, 0xa7,
	0x5c, 0x53, 0x5a, 0xca, 0x30, 0xe2, 0xf7, 0xa8, 0x84, 0x1d, 0x86, 0x45, 0x49, 0x3b, 0x5d, 0xc5,
	0x0b, 0x88, 0x46, 0xf4, 0x9c, 0xcf, 0x71, 0xbb, 0xb5, 0xea, 0x4a, 0xfa, 0x77, 0x7c, 0x68, 0xa7,
	0xb8, 0x84, 0xde, 0x58, 0x95, 0x65, 0xce, 0x28, 0x0e, 0x53, 0x4a, 0x5f, 0xbf, 0x43, 0x6d, 0xe3,
	0x35, 0xac, 0x4f, 0x54, 0x51, 0xcc, 0xd2, 0xec, 0x15, 0xe5, 0xbf, 0x04, 0x48, 0xf3, 0xde, 0x1f,
	0xee, 0x7e, 0xa4, 0x1e, 0x45, 0x1e, 0x9f, 0xb8, 0xf3, 0x75, 0xde, 0x3e, 0x08, 0x6a, 0x36, 0xea,
	0x09, 0xb6, 0x5c, 0xc5, 0x2c, 0xe6, 0x28, 0xd0, 0xe3, 0x6d, 0xe7, 0xf8, 0x7f, 0x83, 0x4d, 0x1e,
	0x03, 0x4c, 0xab, 0x22, 0x67, 0x73, 0x37, 0xfb, 0xb2, 0x02, 0x8b, 0x24, 0xeb, 0x20, 0xa0, 0xd8,
	0x90, 0x7b, 0x88, 0xcd, 0xd2, 0x6e, 0x29, 0x2d, 0x78, 0x79, 0x32, 0x2e, 0xec, 0xbe, 0xd4, 0xd7,
	0x96, 0x4b, 0x9e, 0xf5, 0xf4, 0x31, 0x9f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xe2, 0xe9,
	0x0d, 0xfd, 0x02, 0x00, 0x00,
}
