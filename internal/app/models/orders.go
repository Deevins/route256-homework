package models

import (
	"strconv"
	"time"
)

// constants for order status processing
const (
	OrderStateCreated   = OrderStatus("created")
	OrderStateProcessed = OrderStatus("processed")
	OrderStateCompleted = OrderStatus("completed")
)

// constants for calling repo methods from CLI
const (
	CreateOrder       = "create"
	GetAllOrders      = "getAll"
	GetOrderByID      = "getByID"
	UpdateOrderStatus = "updateStatus"
	DeleteOrder       = "delete"
)

type OrderID uint64
type OrderStatus string
type ProductName string
type Quantity uint32

type Order struct {
	ID          OrderID     `db:"id" json:"id"`
	UserID      UserID      `db:"user_id" json:"user_id"`
	ProductName ProductName `db:"product_name" json:"product_name"`
	Status      OrderStatus `db:"status" json:"status,omitempty"`
	Quantity    Quantity    `db:"quantity" json:"quantity"`
	OrderDate   time.Time   `db:"order_date" json:"-,omitempty"`
}

func ParseValueToOrderID(value string) (OrderID, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return OrderID(intValue), nil
}
