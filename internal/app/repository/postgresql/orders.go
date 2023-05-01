package repository

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
)

type PostgresqlOrderRepo struct {
	db database.DBops
}

func NewPostgresqlOrderRepo(db database.DBops) *PostgresqlOrderRepo {
	return &PostgresqlOrderRepo{db: db}
}

func (r *PostgresqlOrderRepo) CreateOrder(ctx context.Context,
	userID models.UserID,
	productName models.ProductName,
	quantity models.Quantity) (models.OrderID, error) {
	var id models.OrderID

	if productName == "" || quantity < 1 {
		return 0, errors.New("no product name or quantity provided")
	}
	if userID == 0 {
		return 0, errors.New("no user ID provided")
	}

	_ = r.db.ExecQueryRow(ctx, `INSERT INTO orders(user_id, product_name, quantity) VALUES ($1, $2, $3) RETURNING id`,
		userID, productName, quantity).Scan(&id)

	return id, nil
}

func (r *PostgresqlOrderRepo) GetAll(ctx context.Context) ([]*models.Order, error) {
	orders := make([]*models.Order, 1)
	err := r.db.Select(ctx, &orders, "SELECT id,user_id,product_name,quantity,status FROM orders")

	return orders, err
}

func (r *PostgresqlOrderRepo) GetByID(ctx context.Context, ID models.OrderID) (*models.Order, error) {
	var o models.Order
	err := r.db.Get(ctx, &o, "SELECT id, user_id, product_name, quantity, status, order_date FROM orders WHERE id=$1", ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, models.ErrObjectNotFound
	}
	return &o, err
}

func (r *PostgresqlOrderRepo) UpdateOrderStatus(ctx context.Context, ID models.OrderID, status models.OrderStatus) (bool, error) {

	result, err := r.db.Exec(ctx,
		"UPDATE orders SET status = $2 WHERE id = $1", ID, status)
	return result.RowsAffected() > 0, err

}

func (r *PostgresqlOrderRepo) DeleteOrder(ctx context.Context, ID models.OrderID) (bool, error) {
	result, _ := r.db.Exec(ctx,
		"DELETE FROM orders WHERE id = $1", ID)

	if !(result.RowsAffected() > 0) {
		return false, models.ErrObjectNotFound
	}

	return result.RowsAffected() > 0, nil
}
