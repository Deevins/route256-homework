package service

import (
	"context"

	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	models_dto "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/models"
	repository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
)

type UserService struct {
	userRepo *repository.PostgresqlUserRepo
}

func NewUserService(userRepo *repository.PostgresqlUserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, username models.Username, email models.UserEmail) (models.UserID, error) {
	return u.userRepo.CreateUser(ctx, username, email)
}

func (u *UserService) GetAll(ctx context.Context) ([]*models_dto.UserDTO, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *UserService) GetByID(ctx context.Context, ID models.UserID) (*models_dto.UserDTO, error) {
	return u.userRepo.GetByID(ctx, ID)
}

func (u *UserService) UpdateUsername(ctx context.Context, ID models.UserID, username models.Username) (bool, error) {
	return u.userRepo.UpdateUsername(ctx, ID, username)
}

func (u *UserService) UpdateEmail(ctx context.Context, ID models.UserID, email models.UserEmail) (bool, error) {
	return u.userRepo.UpdateEmail(ctx, ID, email)
}

func (u *UserService) DeleteUser(ctx context.Context, ID models.UserID) (bool, error) {
	return u.userRepo.DeleteUser(ctx, ID)
}
