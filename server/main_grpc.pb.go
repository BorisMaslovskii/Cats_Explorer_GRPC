// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// CatsExplorerClient is the client API for CatsExplorer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatsExplorerClient interface {
	GetCat(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Cat, error)
	GetAllCats(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Cats, error)
	UpdateCat(ctx context.Context, in *Cat, opts ...grpc.CallOption) (*Info, error)
	CreateNewCat(ctx context.Context, in *Cat, opts ...grpc.CallOption) (*Info, error)
	DeleteCat(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Info, error)
}

type catsExplorerClient struct {
	cc grpc.ClientConnInterface
}

func NewCatsExplorerClient(cc grpc.ClientConnInterface) CatsExplorerClient {
	return &catsExplorerClient{cc}
}

func (c *catsExplorerClient) GetCat(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Cat, error) {
	out := new(Cat)
	err := c.cc.Invoke(ctx, "/main.CatsExplorer/GetCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsExplorerClient) GetAllCats(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Cats, error) {
	out := new(Cats)
	err := c.cc.Invoke(ctx, "/main.CatsExplorer/GetAllCats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsExplorerClient) UpdateCat(ctx context.Context, in *Cat, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/main.CatsExplorer/UpdateCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsExplorerClient) CreateNewCat(ctx context.Context, in *Cat, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/main.CatsExplorer/CreateNewCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsExplorerClient) DeleteCat(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/main.CatsExplorer/DeleteCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatsExplorerServer is the server API for CatsExplorer service.
// All implementations must embed UnimplementedCatsExplorerServer
// for forward compatibility
type CatsExplorerServer interface {
	GetCat(context.Context, *Id) (*Cat, error)
	GetAllCats(context.Context, *Nothing) (*Cats, error)
	UpdateCat(context.Context, *Cat) (*Info, error)
	CreateNewCat(context.Context, *Cat) (*Info, error)
	DeleteCat(context.Context, *Id) (*Info, error)
	mustEmbedUnimplementedCatsExplorerServer()
}

// UnimplementedCatsExplorerServer must be embedded to have forward compatible implementations.
type UnimplementedCatsExplorerServer struct {
}

func (UnimplementedCatsExplorerServer) GetCat(context.Context, *Id) (*Cat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCat not implemented")
}
func (UnimplementedCatsExplorerServer) GetAllCats(context.Context, *Nothing) (*Cats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCats not implemented")
}
func (UnimplementedCatsExplorerServer) UpdateCat(context.Context, *Cat) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCat not implemented")
}
func (UnimplementedCatsExplorerServer) CreateNewCat(context.Context, *Cat) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewCat not implemented")
}
func (UnimplementedCatsExplorerServer) DeleteCat(context.Context, *Id) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCat not implemented")
}
func (UnimplementedCatsExplorerServer) mustEmbedUnimplementedCatsExplorerServer() {}

// UnsafeCatsExplorerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatsExplorerServer will
// result in compilation errors.
type UnsafeCatsExplorerServer interface {
	mustEmbedUnimplementedCatsExplorerServer()
}

func RegisterCatsExplorerServer(s grpc.ServiceRegistrar, srv CatsExplorerServer) {
	s.RegisterService(&CatsExplorer_ServiceDesc, srv)
}

func _CatsExplorer_GetCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsExplorerServer).GetCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CatsExplorer/GetCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsExplorerServer).GetCat(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsExplorer_GetAllCats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsExplorerServer).GetAllCats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CatsExplorer/GetAllCats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsExplorerServer).GetAllCats(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsExplorer_UpdateCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsExplorerServer).UpdateCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CatsExplorer/UpdateCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsExplorerServer).UpdateCat(ctx, req.(*Cat))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsExplorer_CreateNewCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsExplorerServer).CreateNewCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CatsExplorer/CreateNewCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsExplorerServer).CreateNewCat(ctx, req.(*Cat))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsExplorer_DeleteCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsExplorerServer).DeleteCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CatsExplorer/DeleteCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsExplorerServer).DeleteCat(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// CatsExplorer_ServiceDesc is the grpc.ServiceDesc for CatsExplorer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatsExplorer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.CatsExplorer",
	HandlerType: (*CatsExplorerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCat",
			Handler:    _CatsExplorer_GetCat_Handler,
		},
		{
			MethodName: "GetAllCats",
			Handler:    _CatsExplorer_GetAllCats_Handler,
		},
		{
			MethodName: "UpdateCat",
			Handler:    _CatsExplorer_UpdateCat_Handler,
		},
		{
			MethodName: "CreateNewCat",
			Handler:    _CatsExplorer_CreateNewCat_Handler,
		},
		{
			MethodName: "DeleteCat",
			Handler:    _CatsExplorer_DeleteCat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}