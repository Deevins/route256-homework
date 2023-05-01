package transaction

import (
	"context"

	repository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
	"gitlab.ozon.dev/daker255/homework-8/pkg/transaction"
)

type ServiceTxBuilder struct {
	db database.PGX
}

func NewServiceTxBuilder(db database.PGX) *ServiceTxBuilder {
	return &ServiceTxBuilder{db: db}
}

func (b *ServiceTxBuilder) ServiceTx(ctx context.Context) (*transaction.ServiceTx, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &transaction.ServiceTx{
		Tx:     tx,
		Users:  repository.NewPostgresqlUserRepo(tx),
		Orders: repository.NewPostgresqlOrderRepo(tx),
	}, nil
}
