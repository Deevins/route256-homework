package models_dto

import (
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
)

type OrderDTO struct {
	ID          models.OrderID     `db:"id" json:"id"`
	UserID      models.UserID      `db:"user_id" json:"user_id"`
	ProductName models.ProductName `db:"product_name" json:"product_name"`
	Status      models.OrderStatus `db:"status" json:"status"`
	Quantity    models.Quantity    `db:"quantity" json:"quantity"`
}
