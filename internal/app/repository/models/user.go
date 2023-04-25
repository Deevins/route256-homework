package models_dto

import (
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
)

type UserDTO struct {
	ID       models.UserID    `db:"id" json:"id"`
	Username models.Username  `db:"username" json:"username"`
	Email    models.UserEmail `db:"email" json:"email"`
}
