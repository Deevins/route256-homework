package database

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
)

// DBops interface for database
type DBops interface {
	Select(ctx context.Context, dest any, query string, args ...any) error
	Get(ctx context.Context, dest any, query string, args ...any) error
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...any) pgx.Row
	GetPool(ctx context.Context) *pgxpool.Pool
}

// PGX содержит все операции с базой данных, включая транзакцию
type PGX interface {
	DBops
	BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error)
}

// Tx - транзакция
type Tx interface {
	DBops
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type PostgresqlDatabase struct {
	cluster *pgxpool.Pool
}

func NewPostgresqlDatabase(cluster *pgxpool.Pool) *PostgresqlDatabase {
	return &PostgresqlDatabase{cluster: cluster}
}

func (db PostgresqlDatabase) GetPool(_ context.Context) *pgxpool.Pool {
	return db.cluster
}

func (db PostgresqlDatabase) Get(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
}

func (db PostgresqlDatabase) Select(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
}

func (db PostgresqlDatabase) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return db.cluster.Exec(ctx, query, args...)
}

func (db PostgresqlDatabase) ExecQueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return db.cluster.QueryRow(ctx, query, args...)
}

// NewDB return instance of Database
func NewDB(ctx context.Context, cfg models.Config) (*PostgresqlDatabase, error) {
	dsn := generateDsn(cfg)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return NewPostgresqlDatabase(pool), nil
}

func generateDsn(config models.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Dbname, config.SSLMode)
}
