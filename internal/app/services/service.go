package service

import (
	"context"

	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/repository"
	models_dto "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type User interface {
	CreateUser(ctx context.Context, username models.Username, email models.UserEmail) (models.UserID, error)
	GetAll(ctx context.Context) ([]*models_dto.UserDTO, error)
	GetByID(ctx context.Context, ID models.UserID) (*models_dto.UserDTO, error)
	UpdateUsername(ctx context.Context, ID models.UserID, username models.Username) (bool, error)
	UpdateEmail(ctx context.Context, ID models.UserID, email models.UserEmail) (bool, error)
	DeleteUser(ctx context.Context, ID models.UserID) (bool, error)
}

type Order interface {
	CreateOrder(ctx context.Context, userID models.UserID, productName models.ProductName, quantity models.Quantity) (models.OrderID, error)
	GetAll(ctx context.Context) ([]*models.Order, error)
	GetByID(ctx context.Context, ID models.OrderID) (*models.Order, error)
	UpdateOrderStatus(ctx context.Context, ID models.OrderID, status models.OrderStatus) (bool, error)
	DeleteOrder(ctx context.Context, ID models.OrderID) (bool, error)
}

type CoreService struct {
	User  User
	Order Order
}

func NewCoreService(repos *repository.PostgresqlRepository) *CoreService {
	return &CoreService{
		User:  repos.UserRepository,
		Order: repos.OrderRepository,
	}
}
