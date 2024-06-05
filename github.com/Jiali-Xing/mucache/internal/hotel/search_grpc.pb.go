// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: internal/hotel/search.proto

package hotel

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	SearchHotels(ctx context.Context, in *SearchHotelsRequest, opts ...grpc.CallOption) (*SearchHotelsResponse, error)
	Nearby(ctx context.Context, in *NearbyRequest, opts ...grpc.CallOption) (*NearbyResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SearchHotels(ctx context.Context, in *SearchHotelsRequest, opts ...grpc.CallOption) (*SearchHotelsResponse, error) {
	out := new(SearchHotelsResponse)
	err := c.cc.Invoke(ctx, "/hotel.SearchService/SearchHotels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) Nearby(ctx context.Context, in *NearbyRequest, opts ...grpc.CallOption) (*NearbyResponse, error) {
	out := new(NearbyResponse)
	err := c.cc.Invoke(ctx, "/hotel.SearchService/Nearby", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	SearchHotels(context.Context, *SearchHotelsRequest) (*SearchHotelsResponse, error)
	Nearby(context.Context, *NearbyRequest) (*NearbyResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) SearchHotels(context.Context, *SearchHotelsRequest) (*SearchHotelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchHotels not implemented")
}
func (UnimplementedSearchServiceServer) Nearby(context.Context, *NearbyRequest) (*NearbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nearby not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_SearchHotels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchHotelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchHotels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel.SearchService/SearchHotels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchHotels(ctx, req.(*SearchHotelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_Nearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NearbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Nearby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel.SearchService/Nearby",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Nearby(ctx, req.(*NearbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hotel.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchHotels",
			Handler:    _SearchService_SearchHotels_Handler,
		},
		{
			MethodName: "Nearby",
			Handler:    _SearchService_Nearby_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/hotel/search.proto",
}
