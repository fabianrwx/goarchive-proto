// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: application/node.proto

package application

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NodeService_GetNode_FullMethodName    = "/NodeService/GetNode"
	NodeService_ListNodes_FullMethodName  = "/NodeService/ListNodes"
	NodeService_CreateNode_FullMethodName = "/NodeService/CreateNode"
	NodeService_UpdateNode_FullMethodName = "/NodeService/UpdateNode"
	NodeService_DeleteNode_FullMethodName = "/NodeService/DeleteNode"
)

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error)
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (NodeService_ListNodesClient, error)
	CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*Node, error)
	UpdateNode(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*Node, error)
	DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, NodeService_GetNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (NodeService_ListNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], NodeService_ListNodes_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceListNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_ListNodesClient interface {
	Recv() (*Node, error)
	grpc.ClientStream
}

type nodeServiceListNodesClient struct {
	grpc.ClientStream
}

func (x *nodeServiceListNodesClient) Recv() (*Node, error) {
	m := new(Node)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, NodeService_CreateNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) UpdateNode(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, NodeService_UpdateNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NodeService_DeleteNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	GetNode(context.Context, *GetNodeRequest) (*Node, error)
	ListNodes(*ListNodesRequest, NodeService_ListNodesServer) error
	CreateNode(context.Context, *CreateNodeRequest) (*Node, error)
	UpdateNode(context.Context, *UpdateNodeRequest) (*Node, error)
	DeleteNode(context.Context, *DeleteNodeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) GetNode(context.Context, *GetNodeRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedNodeServiceServer) ListNodes(*ListNodesRequest, NodeService_ListNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedNodeServiceServer) CreateNode(context.Context, *CreateNodeRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNode not implemented")
}
func (UnimplementedNodeServiceServer) UpdateNode(context.Context, *UpdateNodeRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNode not implemented")
}
func (UnimplementedNodeServiceServer) DeleteNode(context.Context, *DeleteNodeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ListNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListNodesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).ListNodes(m, &nodeServiceListNodesServer{stream})
}

type NodeService_ListNodesServer interface {
	Send(*Node) error
	grpc.ServerStream
}

type nodeServiceListNodesServer struct {
	grpc.ServerStream
}

func (x *nodeServiceListNodesServer) Send(m *Node) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_CreateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).CreateNode(ctx, req.(*CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_UpdateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).UpdateNode(ctx, req.(*UpdateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_DeleteNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).DeleteNode(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNode",
			Handler:    _NodeService_GetNode_Handler,
		},
		{
			MethodName: "CreateNode",
			Handler:    _NodeService_CreateNode_Handler,
		},
		{
			MethodName: "UpdateNode",
			Handler:    _NodeService_UpdateNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _NodeService_DeleteNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListNodes",
			Handler:       _NodeService_ListNodes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "application/node.proto",
}