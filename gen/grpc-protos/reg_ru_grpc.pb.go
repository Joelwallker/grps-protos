// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: reg_ru/reg_ru.proto

package grpc_protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RegRuService_GetResources_FullMethodName = "/reg_ru.RegRuService/GetResources"
)

// RegRuServiceClient is the client API for RegRuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegRuServiceClient interface {
	GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error)
}

type regRuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegRuServiceClient(cc grpc.ClientConnInterface) RegRuServiceClient {
	return &regRuServiceClient{cc}
}

func (c *regRuServiceClient) GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResourcesResponse)
	err := c.cc.Invoke(ctx, RegRuService_GetResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegRuServiceServer is the server API for RegRuService service.
// All implementations must embed UnimplementedRegRuServiceServer
// for forward compatibility.
type RegRuServiceServer interface {
	GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error)
	mustEmbedUnimplementedRegRuServiceServer()
}

// UnimplementedRegRuServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRegRuServiceServer struct{}

func (UnimplementedRegRuServiceServer) GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}
func (UnimplementedRegRuServiceServer) mustEmbedUnimplementedRegRuServiceServer() {}
func (UnimplementedRegRuServiceServer) testEmbeddedByValue()                      {}

// UnsafeRegRuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegRuServiceServer will
// result in compilation errors.
type UnsafeRegRuServiceServer interface {
	mustEmbedUnimplementedRegRuServiceServer()
}

func RegisterRegRuServiceServer(s grpc.ServiceRegistrar, srv RegRuServiceServer) {
	// If the following call pancis, it indicates UnimplementedRegRuServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RegRuService_ServiceDesc, srv)
}

func _RegRuService_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegRuServiceServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegRuService_GetResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegRuServiceServer).GetResources(ctx, req.(*GetResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegRuService_ServiceDesc is the grpc.ServiceDesc for RegRuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegRuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reg_ru.RegRuService",
	HandlerType: (*RegRuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResources",
			Handler:    _RegRuService_GetResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reg_ru/reg_ru.proto",
}
