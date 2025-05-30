// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package market

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

// OperationClient is the client API for Operation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationClient interface {
	WithdrawCrypto(ctx context.Context, in *WithdrawCryptoRequest, opts ...grpc.CallOption) (*WithdrawCryptoResponse, error)
	GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...grpc.CallOption) (*GetUserBalanceResponse, error)
}

type operationClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationClient(cc grpc.ClientConnInterface) OperationClient {
	return &operationClient{cc}
}

func (c *operationClient) WithdrawCrypto(ctx context.Context, in *WithdrawCryptoRequest, opts ...grpc.CallOption) (*WithdrawCryptoResponse, error) {
	out := new(WithdrawCryptoResponse)
	err := c.cc.Invoke(ctx, "/market.operation/WithdrawCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...grpc.CallOption) (*GetUserBalanceResponse, error) {
	out := new(GetUserBalanceResponse)
	err := c.cc.Invoke(ctx, "/market.operation/GetUserBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServer is the server API for Operation service.
// All implementations must embed UnimplementedOperationServer
// for forward compatibility
type OperationServer interface {
	WithdrawCrypto(context.Context, *WithdrawCryptoRequest) (*WithdrawCryptoResponse, error)
	GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error)
	mustEmbedUnimplementedOperationServer()
}

// UnimplementedOperationServer must be embedded to have forward compatible implementations.
type UnimplementedOperationServer struct {
}

func (UnimplementedOperationServer) WithdrawCrypto(context.Context, *WithdrawCryptoRequest) (*WithdrawCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawCrypto not implemented")
}
func (UnimplementedOperationServer) GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}
func (UnimplementedOperationServer) mustEmbedUnimplementedOperationServer() {}

// UnsafeOperationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationServer will
// result in compilation errors.
type UnsafeOperationServer interface {
	mustEmbedUnimplementedOperationServer()
}

func RegisterOperationServer(s grpc.ServiceRegistrar, srv OperationServer) {
	s.RegisterService(&Operation_ServiceDesc, srv)
}

func _Operation_WithdrawCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).WithdrawCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.operation/WithdrawCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).WithdrawCrypto(ctx, req.(*WithdrawCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_GetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).GetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.operation/GetUserBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).GetUserBalance(ctx, req.(*GetUserBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operation_ServiceDesc is the grpc.ServiceDesc for Operation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.operation",
	HandlerType: (*OperationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WithdrawCrypto",
			Handler:    _Operation_WithdrawCrypto_Handler,
		},
		{
			MethodName: "GetUserBalance",
			Handler:    _Operation_GetUserBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market/operation.proto",
}
