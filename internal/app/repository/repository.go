package repository

import (
	"context"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	models_dto "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/models"
	repository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username models.Username, email models.UserEmail) (models.UserID, error)
	GetAll(ctx context.Context) ([]*models_dto.UserDTO, error)
	GetByID(ctx context.Context, ID models.UserID) (*models_dto.UserDTO, error)
	UpdateUsername(ctx context.Context, ID models.UserID, username models.Username) (bool, error)
	UpdateEmail(ctx context.Context, ID models.UserID, email models.UserEmail) (bool, error)
	DeleteUser(ctx context.Context, ID models.UserID) (bool, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, userID models.UserID, productName models.ProductName, quantity models.Quantity) (models.OrderID, error)
	GetAll(ctx context.Context) ([]*models.Order, error)
	GetByID(ctx context.Context, ID models.OrderID) (*models.Order, error)
	UpdateOrderStatus(ctx context.Context, ID models.OrderID, state models.OrderStatus) (bool, error)
	DeleteOrder(ctx context.Context, ID models.OrderID) (bool, error)
}

type PostgresqlRepository struct {
	UserRepository
	OrderRepository
}

func NewPostgresqlRepository(db database.DBops) *PostgresqlRepository {
	return &PostgresqlRepository{
		UserRepository:  repository.NewPostgresqlUserRepo(db),
		OrderRepository: repository.NewPostgresqlOrderRepo(db),
	}
}
