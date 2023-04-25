package models

import (
	"database/sql"
	"strconv"
	"time"
)

// constants for calling repo methods from CLI
const (
	CreateUser     = "create"
	GetAllUsers    = "getAll"
	GetUserByID    = "getByID"
	UpdateUsername = "updateUsername"
	UpdateEmail    = "updateEmail"
	DeleteUser     = "delete"
)

type UserID uint64
type UserEmail string
type Username string

type User struct {
	ID           UserID       `db:"id" json:"id"`
	Username     Username     `db:"username" json:"username"`
	Email        UserEmail    `db:"email" json:"email"`
	RegisterDate time.Time    `db:"register_date" json:"-"`
	CreatedAt    time.Time    `db:"created_at" json:"-"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"-"`
}

func ParseValueToUserID(value string) (UserID, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return UserID(intValue), nil
}
