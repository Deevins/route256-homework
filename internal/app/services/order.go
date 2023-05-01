package service

import (
	"context"

	models "gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	repository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
)

type OrderService struct {
	userRepo  *repository.PostgresqlUserRepo
	orderRepo *repository.PostgresqlOrderRepo
}

func NewOrderService(userRepo *repository.PostgresqlUserRepo, orderRepo *repository.PostgresqlOrderRepo) *OrderService {
	return &OrderService{
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, userID models.UserID, productName models.ProductName, quantity models.Quantity) (models.OrderID, error) {
	if _, err := o.userRepo.GetByID(ctx, userID); err != nil {
		return 0, err
	}

	return o.orderRepo.CreateOrder(ctx, userID, productName, quantity)
}

func (o *OrderService) GetAll(ctx context.Context) ([]*models.Order, error) {
	return o.orderRepo.GetAll(ctx)
}

func (o *OrderService) GetByID(ctx context.Context, ID models.OrderID) (*models.Order, error) {
	return o.orderRepo.GetByID(ctx, ID)
}

func (o *OrderService) UpdateOrderStatus(ctx context.Context, ID models.OrderID, status models.OrderStatus) (bool, error) {
	if _, err := o.orderRepo.GetByID(ctx, ID); err != nil {
		return false, err
	}
	return o.orderRepo.UpdateOrderStatus(ctx, ID, status)

}

func (o *OrderService) DeleteOrder(ctx context.Context, ID models.OrderID) (bool, error) {
	if _, err := o.orderRepo.GetByID(ctx, ID); err != nil {
		return false, err
	}
	return o.orderRepo.DeleteOrder(ctx, ID)
}
