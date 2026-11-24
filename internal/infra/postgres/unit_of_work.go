package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWork struct {
	pool *pgxpool.Pool
}

func NewUnitOfWork(pool *pgxpool.Pool) *UnitOfWork {
	return &UnitOfWork{pool: pool}
}

// should have returning, also, will it rollback after commit? shouln't pg throw in that case
func (uow *UnitOfWork) Execute(ctx context.Context, callback func(pgx.Tx) error) error {
	tx, err := uow.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if err := callback(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
