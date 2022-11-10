// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: helloworld.proto

package helloworld

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

// HWClient is the client API for HW service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HWClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	HelloWorldPrefix(ctx context.Context, in *HelloWorldPrefixRequest, opts ...grpc.CallOption) (*HelloWorldPrefixResponse, error)
}

type hWClient struct {
	cc grpc.ClientConnInterface
}

func NewHWClient(cc grpc.ClientConnInterface) HWClient {
	return &hWClient{cc}
}

func (c *hWClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/helloworld.HW/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hWClient) HelloWorldPrefix(ctx context.Context, in *HelloWorldPrefixRequest, opts ...grpc.CallOption) (*HelloWorldPrefixResponse, error) {
	out := new(HelloWorldPrefixResponse)
	err := c.cc.Invoke(ctx, "/helloworld.HW/HelloWorldPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HWServer is the server API for HW service.
// All implementations must embed UnimplementedHWServer
// for forward compatibility
type HWServer interface {
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	HelloWorldPrefix(context.Context, *HelloWorldPrefixRequest) (*HelloWorldPrefixResponse, error)
	mustEmbedUnimplementedHWServer()
}

// UnimplementedHWServer must be embedded to have forward compatible implementations.
type UnimplementedHWServer struct {
}

func (UnimplementedHWServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedHWServer) HelloWorldPrefix(context.Context, *HelloWorldPrefixRequest) (*HelloWorldPrefixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorldPrefix not implemented")
}
func (UnimplementedHWServer) mustEmbedUnimplementedHWServer() {}

// UnsafeHWServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HWServer will
// result in compilation errors.
type UnsafeHWServer interface {
	mustEmbedUnimplementedHWServer()
}

func RegisterHWServer(s grpc.ServiceRegistrar, srv HWServer) {
	s.RegisterService(&HW_ServiceDesc, srv)
}

func _HW_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HWServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HW/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HWServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HW_HelloWorldPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HWServer).HelloWorldPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HW/HelloWorldPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HWServer).HelloWorldPrefix(ctx, req.(*HelloWorldPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HW_ServiceDesc is the grpc.ServiceDesc for HW service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HW_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HW",
	HandlerType: (*HWServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _HW_HelloWorld_Handler,
		},
		{
			MethodName: "HelloWorldPrefix",
			Handler:    _HW_HelloWorldPrefix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
