// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: message.proto

package message

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
	Message_MessageHandler_FullMethodName  = "/message.Message/MessageHandler"
	Message_ReadHandler_FullMethodName     = "/message.Message/ReadHandler"
	Message_WithdrawHandler_FullMethodName = "/message.Message/WithdrawHandler"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	MessageHandler(ctx context.Context, in *MessageHandlerRequest, opts ...grpc.CallOption) (*Empty, error)
	ReadHandler(ctx context.Context, in *ReadHandlerRequest, opts ...grpc.CallOption) (*Empty, error)
	WithdrawHandler(ctx context.Context, in *WithdrawHandlerRequest, opts ...grpc.CallOption) (*Empty, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) MessageHandler(ctx context.Context, in *MessageHandlerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Message_MessageHandler_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) ReadHandler(ctx context.Context, in *ReadHandlerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Message_ReadHandler_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) WithdrawHandler(ctx context.Context, in *WithdrawHandlerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Message_WithdrawHandler_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	MessageHandler(context.Context, *MessageHandlerRequest) (*Empty, error)
	ReadHandler(context.Context, *ReadHandlerRequest) (*Empty, error)
	WithdrawHandler(context.Context, *WithdrawHandlerRequest) (*Empty, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) MessageHandler(context.Context, *MessageHandlerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageHandler not implemented")
}
func (UnimplementedMessageServer) ReadHandler(context.Context, *ReadHandlerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadHandler not implemented")
}
func (UnimplementedMessageServer) WithdrawHandler(context.Context, *WithdrawHandlerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawHandler not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_MessageHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).MessageHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_MessageHandler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).MessageHandler(ctx, req.(*MessageHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_ReadHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).ReadHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_ReadHandler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).ReadHandler(ctx, req.(*ReadHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_WithdrawHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).WithdrawHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_WithdrawHandler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).WithdrawHandler(ctx, req.(*WithdrawHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageHandler",
			Handler:    _Message_MessageHandler_Handler,
		},
		{
			MethodName: "ReadHandler",
			Handler:    _Message_ReadHandler_Handler,
		},
		{
			MethodName: "WithdrawHandler",
			Handler:    _Message_WithdrawHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
