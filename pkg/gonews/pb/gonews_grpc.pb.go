// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/gonews/pb/gonews.proto

package pb

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

// GoNewsServiceClient is the client API for GoNewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoNewsServiceClient interface {
	Posts(ctx context.Context, in *PostsRequest, opts ...grpc.CallOption) (*PostsResponse, error)
	NewsFullDetailed(ctx context.Context, in *OneNewsRequest, opts ...grpc.CallOption) (*DetailedNewsResponse, error)
	NewsShortDetailed(ctx context.Context, in *OneNewsRequest, opts ...grpc.CallOption) (*OnePostResponse, error)
	FilterNews(ctx context.Context, in *FilterNewsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error)
	ListNews(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error)
}

type goNewsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoNewsServiceClient(cc grpc.ClientConnInterface) GoNewsServiceClient {
	return &goNewsServiceClient{cc}
}

func (c *goNewsServiceClient) Posts(ctx context.Context, in *PostsRequest, opts ...grpc.CallOption) (*PostsResponse, error) {
	out := new(PostsResponse)
	err := c.cc.Invoke(ctx, "/gonews.GoNewsService/Posts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goNewsServiceClient) NewsFullDetailed(ctx context.Context, in *OneNewsRequest, opts ...grpc.CallOption) (*DetailedNewsResponse, error) {
	out := new(DetailedNewsResponse)
	err := c.cc.Invoke(ctx, "/gonews.GoNewsService/NewsFullDetailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goNewsServiceClient) NewsShortDetailed(ctx context.Context, in *OneNewsRequest, opts ...grpc.CallOption) (*OnePostResponse, error) {
	out := new(OnePostResponse)
	err := c.cc.Invoke(ctx, "/gonews.GoNewsService/NewsShortDetailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goNewsServiceClient) FilterNews(ctx context.Context, in *FilterNewsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error) {
	out := new(ListPostsResponse)
	err := c.cc.Invoke(ctx, "/gonews.GoNewsService/FilterNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goNewsServiceClient) ListNews(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error) {
	out := new(ListPostsResponse)
	err := c.cc.Invoke(ctx, "/gonews.GoNewsService/ListNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoNewsServiceServer is the server API for GoNewsService service.
// All implementations must embed UnimplementedGoNewsServiceServer
// for forward compatibility
type GoNewsServiceServer interface {
	Posts(context.Context, *PostsRequest) (*PostsResponse, error)
	NewsFullDetailed(context.Context, *OneNewsRequest) (*DetailedNewsResponse, error)
	NewsShortDetailed(context.Context, *OneNewsRequest) (*OnePostResponse, error)
	FilterNews(context.Context, *FilterNewsRequest) (*ListPostsResponse, error)
	ListNews(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
	mustEmbedUnimplementedGoNewsServiceServer()
}

// UnimplementedGoNewsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoNewsServiceServer struct {
}

func (UnimplementedGoNewsServiceServer) Posts(context.Context, *PostsRequest) (*PostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Posts not implemented")
}
func (UnimplementedGoNewsServiceServer) NewsFullDetailed(context.Context, *OneNewsRequest) (*DetailedNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewsFullDetailed not implemented")
}
func (UnimplementedGoNewsServiceServer) NewsShortDetailed(context.Context, *OneNewsRequest) (*OnePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewsShortDetailed not implemented")
}
func (UnimplementedGoNewsServiceServer) FilterNews(context.Context, *FilterNewsRequest) (*ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterNews not implemented")
}
func (UnimplementedGoNewsServiceServer) ListNews(context.Context, *ListPostsRequest) (*ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNews not implemented")
}
func (UnimplementedGoNewsServiceServer) mustEmbedUnimplementedGoNewsServiceServer() {}

// UnsafeGoNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoNewsServiceServer will
// result in compilation errors.
type UnsafeGoNewsServiceServer interface {
	mustEmbedUnimplementedGoNewsServiceServer()
}

func RegisterGoNewsServiceServer(s grpc.ServiceRegistrar, srv GoNewsServiceServer) {
	s.RegisterService(&GoNewsService_ServiceDesc, srv)
}

func _GoNewsService_Posts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoNewsServiceServer).Posts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gonews.GoNewsService/Posts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoNewsServiceServer).Posts(ctx, req.(*PostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoNewsService_NewsFullDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoNewsServiceServer).NewsFullDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gonews.GoNewsService/NewsFullDetailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoNewsServiceServer).NewsFullDetailed(ctx, req.(*OneNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoNewsService_NewsShortDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoNewsServiceServer).NewsShortDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gonews.GoNewsService/NewsShortDetailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoNewsServiceServer).NewsShortDetailed(ctx, req.(*OneNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoNewsService_FilterNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoNewsServiceServer).FilterNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gonews.GoNewsService/FilterNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoNewsServiceServer).FilterNews(ctx, req.(*FilterNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoNewsService_ListNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoNewsServiceServer).ListNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gonews.GoNewsService/ListNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoNewsServiceServer).ListNews(ctx, req.(*ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoNewsService_ServiceDesc is the grpc.ServiceDesc for GoNewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoNewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gonews.GoNewsService",
	HandlerType: (*GoNewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Posts",
			Handler:    _GoNewsService_Posts_Handler,
		},
		{
			MethodName: "NewsFullDetailed",
			Handler:    _GoNewsService_NewsFullDetailed_Handler,
		},
		{
			MethodName: "NewsShortDetailed",
			Handler:    _GoNewsService_NewsShortDetailed_Handler,
		},
		{
			MethodName: "FilterNews",
			Handler:    _GoNewsService_FilterNews_Handler,
		},
		{
			MethodName: "ListNews",
			Handler:    _GoNewsService_ListNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/gonews/pb/gonews.proto",
}
