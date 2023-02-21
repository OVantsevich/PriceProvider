// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/model.proto

package __

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

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceServiceClient interface {
	GetCurrentPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error)
}

type priceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceServiceClient(cc grpc.ClientConnInterface) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) GetCurrentPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error) {
	out := new(GetPricesResponse)
	err := c.cc.Invoke(ctx, "/PriceService/GetCurrentPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceServiceServer is the server API for PriceService service.
// All implementations must embed UnimplementedPriceServiceServer
// for forward compatibility
type PriceServiceServer interface {
	GetCurrentPrices(context.Context, *GetPricesRequest) (*GetPricesResponse, error)
	mustEmbedUnimplementedPriceServiceServer()
}

// UnimplementedPriceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPriceServiceServer struct {
}

func (UnimplementedPriceServiceServer) GetCurrentPrices(context.Context, *GetPricesRequest) (*GetPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentPrices not implemented")
}
func (UnimplementedPriceServiceServer) mustEmbedUnimplementedPriceServiceServer() {}

// UnsafePriceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceServiceServer will
// result in compilation errors.
type UnsafePriceServiceServer interface {
	mustEmbedUnimplementedPriceServiceServer()
}

func RegisterPriceServiceServer(s grpc.ServiceRegistrar, srv PriceServiceServer) {
	s.RegisterService(&PriceService_ServiceDesc, srv)
}

func _PriceService_GetCurrentPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetCurrentPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PriceService/GetCurrentPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetCurrentPrices(ctx, req.(*GetPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceService_ServiceDesc is the grpc.ServiceDesc for PriceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentPrices",
			Handler:    _PriceService_GetCurrentPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/model.proto",
}
