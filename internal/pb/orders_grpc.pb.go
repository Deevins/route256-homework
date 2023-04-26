// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: api/proto/orders.proto

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

const (
	OrderServiceV1_CreateOrder_FullMethodName       = "/order.v1.OrderServiceV1/CreateOrder"
	OrderServiceV1_ListOrder_FullMethodName         = "/order.v1.OrderServiceV1/ListOrder"
	OrderServiceV1_GetOrder_FullMethodName          = "/order.v1.OrderServiceV1/GetOrder"
	OrderServiceV1_UpdateOrderStatus_FullMethodName = "/order.v1.OrderServiceV1/UpdateOrderStatus"
	OrderServiceV1_DeleteOrder_FullMethodName       = "/order.v1.OrderServiceV1/DeleteOrder"
)

// OrderServiceV1Client is the client API for OrderServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceV1Client interface {
	// Create new Order
	CreateOrder(ctx context.Context, in *CreateOrderRequestV1, opts ...grpc.CallOption) (*CreateOrderResponseV1, error)
	// Get Order list
	ListOrder(ctx context.Context, in *ListOrderRequestV1, opts ...grpc.CallOption) (*ListOrderResponseV1, error)
	// Get Order by ID
	GetOrder(ctx context.Context, in *GetOrderRequestV1, opts ...grpc.CallOption) (*GetOrderResponseV1, error)
	// Update Order status by ID
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequestV1, opts ...grpc.CallOption) (*UpdateOrderStatusResponseV1, error)
	// Remove Order by ID
	DeleteOrder(ctx context.Context, in *DeleteOrderRequestV1, opts ...grpc.CallOption) (*DeleteOrderResponseV1, error)
}

type orderServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceV1Client(cc grpc.ClientConnInterface) OrderServiceV1Client {
	return &orderServiceV1Client{cc}
}

func (c *orderServiceV1Client) CreateOrder(ctx context.Context, in *CreateOrderRequestV1, opts ...grpc.CallOption) (*CreateOrderResponseV1, error) {
	out := new(CreateOrderResponseV1)
	err := c.cc.Invoke(ctx, OrderServiceV1_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceV1Client) ListOrder(ctx context.Context, in *ListOrderRequestV1, opts ...grpc.CallOption) (*ListOrderResponseV1, error) {
	out := new(ListOrderResponseV1)
	err := c.cc.Invoke(ctx, OrderServiceV1_ListOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceV1Client) GetOrder(ctx context.Context, in *GetOrderRequestV1, opts ...grpc.CallOption) (*GetOrderResponseV1, error) {
	out := new(GetOrderResponseV1)
	err := c.cc.Invoke(ctx, OrderServiceV1_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceV1Client) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequestV1, opts ...grpc.CallOption) (*UpdateOrderStatusResponseV1, error) {
	out := new(UpdateOrderStatusResponseV1)
	err := c.cc.Invoke(ctx, OrderServiceV1_UpdateOrderStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceV1Client) DeleteOrder(ctx context.Context, in *DeleteOrderRequestV1, opts ...grpc.CallOption) (*DeleteOrderResponseV1, error) {
	out := new(DeleteOrderResponseV1)
	err := c.cc.Invoke(ctx, OrderServiceV1_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceV1Server is the server API for OrderServiceV1 service.
// All implementations must embed UnimplementedOrderServiceV1Server
// for forward compatibility
type OrderServiceV1Server interface {
	// Create new Order
	CreateOrder(context.Context, *CreateOrderRequestV1) (*CreateOrderResponseV1, error)
	// Get Order list
	ListOrder(context.Context, *ListOrderRequestV1) (*ListOrderResponseV1, error)
	// Get Order by ID
	GetOrder(context.Context, *GetOrderRequestV1) (*GetOrderResponseV1, error)
	// Update Order status by ID
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequestV1) (*UpdateOrderStatusResponseV1, error)
	// Remove Order by ID
	DeleteOrder(context.Context, *DeleteOrderRequestV1) (*DeleteOrderResponseV1, error)
	mustEmbedUnimplementedOrderServiceV1Server()
}

// UnimplementedOrderServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceV1Server struct {
}

func (UnimplementedOrderServiceV1Server) CreateOrder(context.Context, *CreateOrderRequestV1) (*CreateOrderResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceV1Server) ListOrder(context.Context, *ListOrderRequestV1) (*ListOrderResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderServiceV1Server) GetOrder(context.Context, *GetOrderRequestV1) (*GetOrderResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceV1Server) UpdateOrderStatus(context.Context, *UpdateOrderStatusRequestV1) (*UpdateOrderStatusResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedOrderServiceV1Server) DeleteOrder(context.Context, *DeleteOrderRequestV1) (*DeleteOrderResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServiceV1Server) mustEmbedUnimplementedOrderServiceV1Server() {}

// UnsafeOrderServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceV1Server will
// result in compilation errors.
type UnsafeOrderServiceV1Server interface {
	mustEmbedUnimplementedOrderServiceV1Server()
}

func RegisterOrderServiceV1Server(s grpc.ServiceRegistrar, srv OrderServiceV1Server) {
	s.RegisterService(&OrderServiceV1_ServiceDesc, srv)
}

func _OrderServiceV1_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceV1Server).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServiceV1_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceV1Server).CreateOrder(ctx, req.(*CreateOrderRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServiceV1_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceV1Server).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServiceV1_ListOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceV1Server).ListOrder(ctx, req.(*ListOrderRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServiceV1_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceV1Server).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServiceV1_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceV1Server).GetOrder(ctx, req.(*GetOrderRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServiceV1_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceV1Server).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServiceV1_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceV1Server).UpdateOrderStatus(ctx, req.(*UpdateOrderStatusRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServiceV1_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceV1Server).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServiceV1_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceV1Server).DeleteOrder(ctx, req.(*DeleteOrderRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderServiceV1_ServiceDesc is the grpc.ServiceDesc for OrderServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.v1.OrderServiceV1",
	HandlerType: (*OrderServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderServiceV1_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _OrderServiceV1_ListOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderServiceV1_GetOrder_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _OrderServiceV1_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderServiceV1_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/orders.proto",
}

const ()

// OrderServiceV2Client is the client API for OrderServiceV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceV2Client interface {
}

type orderServiceV2Client struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceV2Client(cc grpc.ClientConnInterface) OrderServiceV2Client {
	return &orderServiceV2Client{cc}
}

// OrderServiceV2Server is the server API for OrderServiceV2 service.
// All implementations must embed UnimplementedOrderServiceV2Server
// for forward compatibility
type OrderServiceV2Server interface {
	mustEmbedUnimplementedOrderServiceV2Server()
}

// UnimplementedOrderServiceV2Server must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceV2Server struct {
}

func (UnimplementedOrderServiceV2Server) mustEmbedUnimplementedOrderServiceV2Server() {}

// UnsafeOrderServiceV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceV2Server will
// result in compilation errors.
type UnsafeOrderServiceV2Server interface {
	mustEmbedUnimplementedOrderServiceV2Server()
}

func RegisterOrderServiceV2Server(s grpc.ServiceRegistrar, srv OrderServiceV2Server) {
	s.RegisterService(&OrderServiceV2_ServiceDesc, srv)
}

// OrderServiceV2_ServiceDesc is the grpc.ServiceDesc for OrderServiceV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderServiceV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.v1.OrderServiceV2",
	HandlerType: (*OrderServiceV2Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/proto/orders.proto",
}
