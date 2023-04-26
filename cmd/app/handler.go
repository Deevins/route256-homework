package main

import (
	"gitlab.ozon.dev/daker255/homework-8/internal/pb"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
)

type UserImplementation struct {
	pb.UnimplementedUserServiceV1Server

	db database.DBops
}

func NewUserImplementation(db database.DBops) *UserImplementation {
	return &UserImplementation{db: db}
}

type OrderImplementation struct {
	pb.UnimplementedOrderServiceV1Server

	db database.DBops
}

func NewOrderImplementation(db database.DBops) *OrderImplementation {
	return &OrderImplementation{db: db}
}
