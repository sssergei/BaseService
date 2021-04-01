// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// GoReleaseServiceClient is the client API for GoReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoReleaseServiceClient interface {
	GetReleaseInfo(ctx context.Context, in *GetReleaseInfoRequest, opts ...grpc.CallOption) (*ReleaseInfo, error)
	ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (*ListReleasesResponse, error)
}

type goReleaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoReleaseServiceClient(cc grpc.ClientConnInterface) GoReleaseServiceClient {
	return &goReleaseServiceClient{cc}
}

func (c *goReleaseServiceClient) GetReleaseInfo(ctx context.Context, in *GetReleaseInfoRequest, opts ...grpc.CallOption) (*ReleaseInfo, error) {
	out := new(ReleaseInfo)
	err := c.cc.Invoke(ctx, "/proto.GoReleaseService/GetReleaseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goReleaseServiceClient) ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (*ListReleasesResponse, error) {
	out := new(ListReleasesResponse)
	err := c.cc.Invoke(ctx, "/proto.GoReleaseService/ListReleases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoReleaseServiceServer is the server API for GoReleaseService service.
// All implementations must embed UnimplementedGoReleaseServiceServer
// for forward compatibility
type GoReleaseServiceServer interface {
	GetReleaseInfo(context.Context, *GetReleaseInfoRequest) (*ReleaseInfo, error)
	ListReleases(context.Context, *ListReleasesRequest) (*ListReleasesResponse, error)
	mustEmbedUnimplementedGoReleaseServiceServer()
}

// UnimplementedGoReleaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoReleaseServiceServer struct {
}

func (UnimplementedGoReleaseServiceServer) GetReleaseInfo(context.Context, *GetReleaseInfoRequest) (*ReleaseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReleaseInfo not implemented")
}
func (UnimplementedGoReleaseServiceServer) ListReleases(context.Context, *ListReleasesRequest) (*ListReleasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReleases not implemented")
}
func (UnimplementedGoReleaseServiceServer) mustEmbedUnimplementedGoReleaseServiceServer() {}

// UnsafeGoReleaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoReleaseServiceServer will
// result in compilation errors.
type UnsafeGoReleaseServiceServer interface {
	mustEmbedUnimplementedGoReleaseServiceServer()
}

func RegisterGoReleaseServiceServer(s grpc.ServiceRegistrar, srv GoReleaseServiceServer) {
	s.RegisterService(&GoReleaseService_ServiceDesc, srv)
}

func _GoReleaseService_GetReleaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleaseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoReleaseServiceServer).GetReleaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoReleaseService/GetReleaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoReleaseServiceServer).GetReleaseInfo(ctx, req.(*GetReleaseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoReleaseService_ListReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoReleaseServiceServer).ListReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoReleaseService/ListReleases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoReleaseServiceServer).ListReleases(ctx, req.(*ListReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoReleaseService_ServiceDesc is the grpc.ServiceDesc for GoReleaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoReleaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GoReleaseService",
	HandlerType: (*GoReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReleaseInfo",
			Handler:    _GoReleaseService_GetReleaseInfo_Handler,
		},
		{
			MethodName: "ListReleases",
			Handler:    _GoReleaseService_ListReleases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gofirst.proto",
}
