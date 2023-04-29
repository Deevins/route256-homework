package main

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	pb "gitlab.ozon.dev/daker255/homework-8/internal/pb/server"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserImplementation struct {
	*pb.UnimplementedUserServiceV1Server

	userService *service.UserService
}

func NewUserImplementation(userService *service.UserService) *UserImplementation {
	return &UserImplementation{userService: userService}
}

// TODO: создать функции для клиента - вкладывать в контекст запроса x-trace-id,
// а для сервера - парсить из контекста запроса x-trace-id и вкладывать в trace-route

func (o *UserImplementation) CreateUser(ctx context.Context, req *pb.CreateUserRequestV1) (*pb.CreateUserResponseV1, error) {
	// Extract TraceID from header
	md, _ := metadata.FromIncomingContext(ctx)
	traceIdString := md["x-trace-id"][0]

	// Convert string to byte array
	traceId, err := trace.TraceIDFromHex(traceIdString)
	if err != nil {
		return nil, err
	}

	// Creating a span context with a predefined trace-id
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	})
	// Embedding span config into the context
	ctx = trace.ContextWithSpanContext(ctx, spanContext)

	ctx, span := tracer.Tracer("server").Start(ctx, "Create User method called on core-service")
	defer span.End()

	username := models.Username(req.Username)
	email := models.UserEmail(req.Email)

	id, err := o.userService.CreateUser(ctx, username, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	res := pb.CreateUserResponseV1{UserId: uint32(id)}
	return &res, nil
}

func (o *UserImplementation) ListUser(ctx context.Context, req *pb.ListUserRequestV1) (*pb.ListUserResponseV1, error) {
	// Extract TraceID from header
	md, _ := metadata.FromIncomingContext(ctx)
	traceIdString := md["x-trace-id"][0]

	// Convert string to byte array
	traceId, err := trace.TraceIDFromHex(traceIdString)
	if err != nil {
		return nil, err
	}

	// Creating a span context with a predefined trace-id
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	})
	// Embedding span config into the context
	ctx = trace.ContextWithSpanContext(ctx, spanContext)

	ctx, span := tracer.Tracer("server").Start(ctx, "List Users method called on core-service")
	defer span.End()

	users, err := o.userService.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	result := make([]*pb.UserDTO, 0, len(users))

	for _, m := range users {
		result = append(result, &pb.UserDTO{
			UserId:   uint32(m.ID),
			Username: string(m.Username),
			Email:    string(m.Email),
		})
	}
	return &pb.ListUserResponseV1{
		Users: result,
	}, nil
}

func (o *UserImplementation) GetUser(ctx context.Context, req *pb.GetUserRequestV1) (*pb.GetUserResponseV1, error) {
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

func (o *UserImplementation) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequestV1) (*pb.UpdateEmailResponseV1, error) {
	id := models.UserID(req.UserId)
	email := models.UserEmail(req.Email)

	isOk, err := o.userService.UpdateEmail(ctx, id, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.UpdateEmailResponseV1{IsOk: isOk}, nil
}

func (o *UserImplementation) DeleteUser(ctx context.Context, req *pb.DeleteUserRequestV1) (*pb.DeleteUserResponseV1, error) {
	id := models.UserID(req.UserId)

	isOk, err := o.userService.DeleteUser(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.DeleteUserResponseV1{IsOk: isOk}, nil
}

type OrderImplementation struct {
	*pb.UnimplementedOrderServiceV1Server

	orderService *service.OrderService
}

func NewOrderImplementation(orderService *service.OrderService) *OrderImplementation {
	return &OrderImplementation{orderService: orderService}
}

func (o *OrderImplementation) CreateOrder(ctx context.Context, req *pb.CreateOrderRequestV1) (*pb.CreateOrderResponseV1, error) {
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

func (o *OrderImplementation) ListOrder(ctx context.Context, req *pb.ListOrderRequestV1) (*pb.ListOrderResponseV1, error) {

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

func (o *OrderImplementation) GetOrder(ctx context.Context, req *pb.GetOrderRequestV1) (*pb.GetOrderResponseV1, error) {
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

func (o *OrderImplementation) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequestV1) (*pb.UpdateOrderStatusResponseV1, error) {
	id := models.OrderID(req.OrderId)
	orderStatus := models.OrderStatus(req.Status)

	isOk, err := o.orderService.UpdateOrderStatus(ctx, id, orderStatus)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.UpdateOrderStatusResponseV1{IsOk: isOk}, nil
}

func (o *OrderImplementation) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequestV1) (*pb.DeleteOrderResponseV1, error) {
	id := models.OrderID(req.OrderId)

	isOk, err := o.orderService.DeleteOrder(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %s", err))
	}

	return &pb.DeleteOrderResponseV1{IsOk: isOk}, nil
}
