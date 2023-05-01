package transaction

import (
	"context"

	"gitlab.ozon.dev/daker255/homework-8/internal/app/repository"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
)

type ServiceTxBuilder interface {
	ServiceTx(ctx context.Context) (*ServiceTx, error)
}

type ServiceTx struct {
	database.Tx
	Users  repository.UserRepository
	Orders repository.OrderRepository
}
