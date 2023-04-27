package repository

import (
	"context"
	"database/sql"
	"errors"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	models_dto "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/models"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
)

var (
	ErrEmptyEmailOrUsername = errors.New("empty email or username field")
)

type PostgresqlUserRepo struct {
	db database.DBops
}

func NewPostgresqlUserRepo(db database.DBops) *PostgresqlUserRepo {
	return &PostgresqlUserRepo{db: db}
}

func (r *PostgresqlUserRepo) CreateUser(ctx context.Context, username models.Username, email models.UserEmail) (models.UserID, error) {
	var id models.UserID
	if email == "" || username == "" {
		return models.UserID(0), ErrEmptyEmailOrUsername
	}
	err := r.db.ExecQueryRow(ctx, `INSERT INTO users(username,email) VALUES ($1, $2) RETURNING id`, username, email).Scan(&id)

	return id, err
}

func (r *PostgresqlUserRepo) GetByID(ctx context.Context, ID models.UserID) (*models_dto.UserDTO, error) {
	var u models_dto.UserDTO
	err := r.db.Get(ctx, &u, "SELECT id,username,email FROM users WHERE id=$1", ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, models.ErrObjectNotFound
	}
	return &u, err
}

func (r *PostgresqlUserRepo) GetAll(ctx context.Context) ([]*models_dto.UserDTO, error) {
	users := make([]*models_dto.UserDTO, 0)
	err := r.db.Select(ctx, &users, "SELECT id,username,email FROM users")

	return users, err
}

func (r *PostgresqlUserRepo) UpdateUsername(ctx context.Context, ID models.UserID, username models.Username) (bool, error) {
	result, _ := r.db.Exec(ctx,
		"UPDATE users SET username = $2 WHERE id = $1", ID, username)

	if result.RowsAffected() > 0 == false {
		return false, models.ErrObjectNotFound
	}
	return result.RowsAffected() > 0, nil
}

func (r *PostgresqlUserRepo) UpdateEmail(ctx context.Context, ID models.UserID, email models.UserEmail) (bool, error) {
	result, _ := r.db.Exec(ctx,
		"UPDATE users SET email = $2 WHERE id = $1", ID, email)
	if result.RowsAffected() > 0 == false {
		return false, models.ErrObjectNotFound
	}
	return result.RowsAffected() > 0, nil
}

func (r *PostgresqlUserRepo) DeleteUser(ctx context.Context, ID models.UserID) (bool, error) {
	result, err := r.db.Exec(ctx,
		"DELETE FROM users WHERE id = $1 RETURNING *", ID)
	return result.RowsAffected() > 0, err
}
