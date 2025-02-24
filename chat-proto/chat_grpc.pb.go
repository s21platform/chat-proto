// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: chat.proto

package chat_proto

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
	ChatService_CreatePrivateChat_FullMethodName = "/ChatService/CreatePrivateChat"
	ChatService_GetChats_FullMethodName          = "/ChatService/GetChats"
	ChatService_GetRecentMessages_FullMethodName = "/ChatService/GetRecentMessages"
	ChatService_DeleteMessage_FullMethodName     = "/ChatService/DeleteMessage"
	ChatService_EditMessage_FullMethodName       = "/ChatService/EditMessage"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreatePrivateChat(ctx context.Context, in *CreatePrivateChatIn, opts ...grpc.CallOption) (*CreatePrivateChatOut, error)
	GetChats(ctx context.Context, in *ChatEmpty, opts ...grpc.CallOption) (*GetChatsOut, error)
	GetRecentMessages(ctx context.Context, in *ChatEmpty, opts ...grpc.CallOption) (*GetRecentMessagesOut, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageIn, opts ...grpc.CallOption) (*DeleteMessageOut, error)
	EditMessage(ctx context.Context, in *EditMessageIn, opts ...grpc.CallOption) (*EditMessageOut, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreatePrivateChat(ctx context.Context, in *CreatePrivateChatIn, opts ...grpc.CallOption) (*CreatePrivateChatOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePrivateChatOut)
	err := c.cc.Invoke(ctx, ChatService_CreatePrivateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChats(ctx context.Context, in *ChatEmpty, opts ...grpc.CallOption) (*GetChatsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatsOut)
	err := c.cc.Invoke(ctx, ChatService_GetChats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetRecentMessages(ctx context.Context, in *ChatEmpty, opts ...grpc.CallOption) (*GetRecentMessagesOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecentMessagesOut)
	err := c.cc.Invoke(ctx, ChatService_GetRecentMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageIn, opts ...grpc.CallOption) (*DeleteMessageOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMessageOut)
	err := c.cc.Invoke(ctx, ChatService_DeleteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EditMessage(ctx context.Context, in *EditMessageIn, opts ...grpc.CallOption) (*EditMessageOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditMessageOut)
	err := c.cc.Invoke(ctx, ChatService_EditMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
type ChatServiceServer interface {
	CreatePrivateChat(context.Context, *CreatePrivateChatIn) (*CreatePrivateChatOut, error)
	GetChats(context.Context, *ChatEmpty) (*GetChatsOut, error)
	GetRecentMessages(context.Context, *ChatEmpty) (*GetRecentMessagesOut, error)
	DeleteMessage(context.Context, *DeleteMessageIn) (*DeleteMessageOut, error)
	EditMessage(context.Context, *EditMessageIn) (*EditMessageOut, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) CreatePrivateChat(context.Context, *CreatePrivateChatIn) (*CreatePrivateChatOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivateChat not implemented")
}
func (UnimplementedChatServiceServer) GetChats(context.Context, *ChatEmpty) (*GetChatsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedChatServiceServer) GetRecentMessages(context.Context, *ChatEmpty) (*GetRecentMessagesOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentMessages not implemented")
}
func (UnimplementedChatServiceServer) DeleteMessage(context.Context, *DeleteMessageIn) (*DeleteMessageOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedChatServiceServer) EditMessage(context.Context, *EditMessageIn) (*EditMessageOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMessage not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreatePrivateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivateChatIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreatePrivateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreatePrivateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreatePrivateChat(ctx, req.(*CreatePrivateChatIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChats(ctx, req.(*ChatEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetRecentMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetRecentMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetRecentMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetRecentMessages(ctx, req.(*ChatEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteMessage(ctx, req.(*DeleteMessageIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMessageIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_EditMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EditMessage(ctx, req.(*EditMessageIn))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrivateChat",
			Handler:    _ChatService_CreatePrivateChat_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _ChatService_GetChats_Handler,
		},
		{
			MethodName: "GetRecentMessages",
			Handler:    _ChatService_GetRecentMessages_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _ChatService_DeleteMessage_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _ChatService_EditMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
