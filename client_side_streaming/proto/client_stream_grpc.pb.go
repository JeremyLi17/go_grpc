// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: client_stream.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SimpleServer_RouteList_FullMethodName = "/SimpleServer/RouteList"
)

// SimpleServerClient is the client API for SimpleServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleServerClient interface {
	RouteList(ctx context.Context, opts ...grpc.CallOption) (SimpleServer_RouteListClient, error)
}

type simpleServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleServerClient(cc grpc.ClientConnInterface) SimpleServerClient {
	return &simpleServerClient{cc}
}

func (c *simpleServerClient) RouteList(ctx context.Context, opts ...grpc.CallOption) (SimpleServer_RouteListClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimpleServer_ServiceDesc.Streams[0], SimpleServer_RouteList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleServerRouteListClient{stream}
	return x, nil
}

type SimpleServer_RouteListClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type simpleServerRouteListClient struct {
	grpc.ClientStream
}

func (x *simpleServerRouteListClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleServerRouteListClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleServerServer is the server API for SimpleServer service.
// All implementations must embed UnimplementedSimpleServerServer
// for forward compatibility
type SimpleServerServer interface {
	RouteList(SimpleServer_RouteListServer) error
	mustEmbedUnimplementedSimpleServerServer()
}

// UnimplementedSimpleServerServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleServerServer struct {
}

func (UnimplementedSimpleServerServer) RouteList(SimpleServer_RouteListServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteList not implemented")
}
func (UnimplementedSimpleServerServer) mustEmbedUnimplementedSimpleServerServer() {}

// UnsafeSimpleServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServerServer will
// result in compilation errors.
type UnsafeSimpleServerServer interface {
	mustEmbedUnimplementedSimpleServerServer()
}

func RegisterSimpleServerServer(s grpc.ServiceRegistrar, srv SimpleServerServer) {
	s.RegisterService(&SimpleServer_ServiceDesc, srv)
}

func _SimpleServer_RouteList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleServerServer).RouteList(&simpleServerRouteListServer{stream})
}

type SimpleServer_RouteListServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type simpleServerRouteListServer struct {
	grpc.ServerStream
}

func (x *simpleServerRouteListServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleServerRouteListServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleServer_ServiceDesc is the grpc.ServiceDesc for SimpleServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SimpleServer",
	HandlerType: (*SimpleServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteList",
			Handler:       _SimpleServer_RouteList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client_stream.proto",
}
