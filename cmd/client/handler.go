package main

import (
	"context"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	"gitlab.ozon.dev/daker255/homework-8/internal/pb"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type UserClientImplementation struct {
	*pb.UnimplementedUserServiceV1Server

	userService *service.UserService
}

func NewUserClientImplementation(userService *service.UserService) *UserClientImplementation {
	return &UserClientImplementation{userService: userService}
}

func (o *UserClientImplementation) CreateUser(ctx context.Context, req *pb.CreateUserRequestV1) (*pb.CreateUserResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "CreateUser method called on Client")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	span.SetAttributes(attribute.Key("username").String(req.Username))
	span.SetAttributes(attribute.Key("email").String(req.Email))

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
	ctx, span := tracer.Tracer("client").Start(ctx, "ListUser method called on Client")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

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
	ctx, span := tracer.Tracer("client").Start(ctx, "GetUser method called on Client")

	span.SetAttributes(attribute.Key("user-id").Int(int(req.UserId)))
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewUserServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.GetUser(ctx, &pb.GetUserRequestV1{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	result := &pb.UserDTO{
		UserId:   resp.User.UserId,
		Username: resp.User.Username,
		Email:    resp.User.Email,
	}

	return &pb.GetUserResponseV1{User: result}, nil
}

func (o *UserClientImplementation) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequestV1) (*pb.UpdateEmailResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "UpdateEmail method called on Client")

	span.SetAttributes(attribute.Key("user-id").Int(int(req.UserId)))
	span.SetAttributes(attribute.Key("email-provided").String(req.GetEmail()))
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewUserServiceV1Client(conn)

	// Calling a method on the server side
	resp, _ := srv.UpdateEmail(ctx, &pb.UpdateEmailRequestV1{
		UserId: req.UserId,
		Email:  req.GetEmail(),
	})

	return &pb.UpdateEmailResponseV1{IsOk: resp.IsOk}, nil
}

func (o *UserClientImplementation) DeleteUser(ctx context.Context, req *pb.DeleteUserRequestV1) (*pb.DeleteUserResponseV1, error) {
	_, span := tracer.Tracer("client").Start(ctx, "DeleteUser method called on Client")

	span.SetAttributes(attribute.Key("user-id").Int(int(req.UserId)))
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewUserServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.DeleteUser(ctx, &pb.DeleteUserRequestV1{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponseV1{IsOk: resp.IsOk}, nil
}

type OrderClientImplementation struct {
	*pb.UnimplementedOrderServiceV1Server

	orderService *service.OrderService
}

func NewOrderClientImplementation(orderService *service.OrderService) *OrderClientImplementation {
	return &OrderClientImplementation{orderService: orderService}
}

func (o *OrderClientImplementation) CreateOrder(ctx context.Context, req *pb.CreateOrderRequestV1) (*pb.CreateOrderResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "CreateOrder method called on Client")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	span.SetAttributes(attribute.Key("username").String(req.ProductName))
	span.SetAttributes(attribute.Key("email").Int(int(req.UserId)))

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewOrderServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.CreateOrder(ctx, &pb.CreateOrderRequestV1{
		UserId:      req.UserId,
		ProductName: req.GetProductName(),
		Quantity:    req.GetQuantity(),
	})
	if err != nil {
		return nil, err
	}

	res := pb.CreateOrderResponseV1{OrderId: resp.OrderId}

	return &res, nil
}

func (o *OrderClientImplementation) ListOrder(ctx context.Context, _ *pb.ListOrderRequestV1) (*pb.ListOrderResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "ListOrder method called on Client")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewOrderServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.ListOrder(ctx, &pb.ListOrderRequestV1{})
	if err != nil {
		return nil, err
	}

	return &pb.ListOrderResponseV1{
		Orders: resp.Orders,
	}, nil
}

func (o *OrderClientImplementation) GetOrder(ctx context.Context, req *pb.GetOrderRequestV1) (*pb.GetOrderResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "GetOrder method called on Client")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	span.SetAttributes(attribute.Key("username").Int(int(req.OrderId)))

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewOrderServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.GetOrder(ctx, &pb.GetOrderRequestV1{
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponseV1{Order: resp.Order}, nil
}

func (o *OrderClientImplementation) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequestV1) (*pb.UpdateOrderStatusResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "UpdateOrderStatus method called on Client")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	span.SetAttributes(attribute.Key("username").Int(int(req.OrderId)))
	span.SetAttributes(attribute.Key("username").String(req.Status))

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewOrderServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.UpdateOrderStatus(ctx, &pb.UpdateOrderStatusRequestV1{
		OrderId: req.OrderId,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderStatusResponseV1{IsOk: resp.IsOk}, nil
}

func (o *OrderClientImplementation) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequestV1) (*pb.DeleteOrderResponseV1, error) {
	ctx, span := tracer.Tracer("client").Start(ctx, "DeleteOrder method called on Client")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceID)

	span.SetAttributes(attribute.Key("username").Int(int(req.OrderId)))

	// Create a connection to the server
	conn, err := grpc.DialContext(ctx, "localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a client object with connection
	srv := pb.NewOrderServiceV1Client(conn)

	// Calling a method on the server side
	resp, err := srv.DeleteOrder(ctx, &pb.DeleteOrderRequestV1{
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteOrderResponseV1{IsOk: resp.IsOk}, nil
}
