package main

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	pb "gitlab.ozon.dev/daker255/homework-8/internal/pb/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserClientImplementation struct {
	*pb.UnimplementedUserServiceV1Server

	userService *service.UserService
}

func NewUserClientImplementation(userService *service.UserService) *UserClientImplementation {
	return &UserClientImplementation{userService: userService}
}

func (o *UserClientImplementation) CreateUser(ctx context.Context, req *pb.CreateUserRequestV1) (*pb.CreateUserResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "Create User method called")
	defer span.End()

	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewUserServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.CreateUser(ctx, &pb.CreateUserRequestV1{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	// Forwarding the response from the server to the client
	return &pb.CreateUserResponseV1{
		UserId: resp.UserId,
	}, nil
}

func (o *UserClientImplementation) ListUser(ctx context.Context, _ *pb.ListUserRequestV1) (*pb.ListUserResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "List Users method called on client service")
	defer span.End()

	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewUserServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.ListUser(ctx, &pb.ListUserRequestV1{})
	if err != nil {
		return nil, err
	}

	// Forwarding the response from the server to the client
	return &pb.ListUserResponseV1{
		Users: resp.Users,
	}, nil
}

func (o *UserClientImplementation) GetUser(ctx context.Context, req *pb.GetUserRequestV1) (*pb.GetUserResponseV1, error) {
	id := models.UserID(req.UserId)

	user, err := o.userService.GetByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	result := &pb.UserDTO{
		UserId:   uint32(user.ID),
		Username: string(user.Username),
		Email:    string(user.Email),
	}

	return &pb.GetUserResponseV1{User: result}, nil
}

func (o *UserClientImplementation) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequestV1) (*pb.UpdateEmailResponseV1, error) {
	id := models.UserID(req.UserId)
	email := models.UserEmail(req.Email)

	isOk, err := o.userService.UpdateEmail(ctx, id, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.UpdateEmailResponseV1{IsOk: isOk}, nil
}

func (o *UserClientImplementation) DeleteUser(ctx context.Context, req *pb.DeleteUserRequestV1) (*pb.DeleteUserResponseV1, error) {
	id := models.UserID(req.UserId)

	isOk, err := o.userService.DeleteUser(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.DeleteUserResponseV1{IsOk: isOk}, nil
}

type OrderClientImplementation struct {
	*pb.UnimplementedOrderServiceV1Server

	orderService *service.OrderService
}

func NewOrderClientImplementation(orderService *service.OrderService) *OrderClientImplementation {
	return &OrderClientImplementation{orderService: orderService}
}

func (o *OrderClientImplementation) CreateOrder(ctx context.Context, req *pb.CreateOrderRequestV1) (*pb.CreateOrderResponseV1, error) {
	userID := models.UserID(req.UserId)
	productName := models.ProductName(req.GetProductName())
	quantity := models.Quantity(req.GetQuantity())

	orderID, err := o.orderService.CreateOrder(ctx, userID, productName, quantity)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	res := pb.CreateOrderResponseV1{OrderId: uint32(orderID)}

	return &res, nil
}

func (o *OrderClientImplementation) ListOrder(ctx context.Context, req *pb.ListOrderRequestV1) (*pb.ListOrderResponseV1, error) {

	orders, err := o.orderService.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	result := make([]*pb.Order, 0, len(orders))

	for _, m := range orders {
		orderDate := m.OrderDate
		timestamp := timestamppb.New(orderDate)
		result = append(result, &pb.Order{
			OrderId:     uint32(m.ID),
			UserId:      uint32(m.UserID),
			ProductName: string(m.ProductName),
			Status:      string(m.Status),
			Quantity:    uint32(m.Quantity),
			OrderDate:   timestamp,
		})
	}
	return &pb.ListOrderResponseV1{
		Orders: result,
	}, nil
}

func (o *OrderClientImplementation) GetOrder(ctx context.Context, req *pb.GetOrderRequestV1) (*pb.GetOrderResponseV1, error) {
	id := models.OrderID(req.OrderId)

	order, err := o.orderService.GetByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	orderDate := order.OrderDate
	timestamp := timestamppb.New(orderDate)
	result := &pb.Order{
		OrderId:     uint32(order.ID),
		UserId:      uint32(order.UserID),
		ProductName: string(order.ProductName),
		Status:      string(order.Status),
		Quantity:    uint32(order.Quantity),
		OrderDate:   timestamp,
	}

	return &pb.GetOrderResponseV1{Order: result}, nil
}

func (o *OrderClientImplementation) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequestV1) (*pb.UpdateOrderStatusResponseV1, error) {
	id := models.OrderID(req.OrderId)
	orderStatus := models.OrderStatus(req.Status)

	isOk, err := o.orderService.UpdateOrderStatus(ctx, id, orderStatus)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.UpdateOrderStatusResponseV1{IsOk: isOk}, nil
}

func (o *OrderClientImplementation) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequestV1) (*pb.DeleteOrderResponseV1, error) {
	id := models.OrderID(req.OrderId)

	isOk, err := o.orderService.DeleteOrder(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.DeleteOrderResponseV1{IsOk: isOk}, nil
}
