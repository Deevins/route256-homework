package main

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	"gitlab.ozon.dev/daker255/homework-8/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserImplementation struct {
	*pb.UnimplementedUserServiceV1Server

	userService *service.UserService
}

//validate in grpc handlers ->
//service ->

func NewUserImplementation(userService *service.UserService) *UserImplementation {
	return &UserImplementation{userService: userService}
}

func (o *UserImplementation) CreateUser(ctx context.Context, req *pb.CreateUserRequestV1) (*pb.CreateUserResponseV1, error) {
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
	//TODO implement me
	panic("implement me")
}

func (o *OrderImplementation) ListOrder(ctx context.Context, req *pb.ListOrderRequestV1) (*pb.ListOrderResponseV1, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderImplementation) GetOrder(ctx context.Context, req *pb.GetOrderRequestV1) (*pb.GetOrderResponseV1, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderImplementation) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequestV1) (*pb.UpdateOrderStatusResponseV1, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderImplementation) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequestV1) (*pb.DeleteOrderResponseV1, error) {
	//TODO implement me
	panic("implement me")
}
