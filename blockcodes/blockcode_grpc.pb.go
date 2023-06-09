// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: blockcodes/blockcode.proto

package listing

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

// BlockCodeTransactionClient is the client API for BlockCodeTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockCodeTransactionClient interface {
	GetTransaction(ctx context.Context, in *Txn, opts ...grpc.CallOption) (*IdTxn, error)
}

type blockCodeTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockCodeTransactionClient(cc grpc.ClientConnInterface) BlockCodeTransactionClient {
	return &blockCodeTransactionClient{cc}
}

func (c *blockCodeTransactionClient) GetTransaction(ctx context.Context, in *Txn, opts ...grpc.CallOption) (*IdTxn, error) {
	out := new(IdTxn)
	err := c.cc.Invoke(ctx, "/blockcodes.BlockCodeTransaction/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockCodeTransactionServer is the server API for BlockCodeTransaction service.
// All implementations must embed UnimplementedBlockCodeTransactionServer
// for forward compatibility
type BlockCodeTransactionServer interface {
	GetTransaction(context.Context, *Txn) (*IdTxn, error)
	mustEmbedUnimplementedBlockCodeTransactionServer()
}

// UnimplementedBlockCodeTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedBlockCodeTransactionServer struct {
}

func (UnimplementedBlockCodeTransactionServer) GetTransaction(context.Context, *Txn) (*IdTxn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedBlockCodeTransactionServer) mustEmbedUnimplementedBlockCodeTransactionServer() {}

// UnsafeBlockCodeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockCodeTransactionServer will
// result in compilation errors.
type UnsafeBlockCodeTransactionServer interface {
	mustEmbedUnimplementedBlockCodeTransactionServer()
}

func RegisterBlockCodeTransactionServer(s grpc.ServiceRegistrar, srv BlockCodeTransactionServer) {
	s.RegisterService(&BlockCodeTransaction_ServiceDesc, srv)
}

func _BlockCodeTransaction_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Txn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockCodeTransactionServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockcodes.BlockCodeTransaction/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockCodeTransactionServer).GetTransaction(ctx, req.(*Txn))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockCodeTransaction_ServiceDesc is the grpc.ServiceDesc for BlockCodeTransaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockCodeTransaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blockcodes.BlockCodeTransaction",
	HandlerType: (*BlockCodeTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransaction",
			Handler:    _BlockCodeTransaction_GetTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockcodes/blockcode.proto",
}
