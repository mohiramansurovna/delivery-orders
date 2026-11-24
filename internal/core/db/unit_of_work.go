package db

import (
	"context"

	"github.com/jackc/pgx/v5" //TODO: remove this from core by creating tx for core
)

type UnitOfWork interface {
	Execute(ctx context.Context, callback func(pgx.Tx) error) error
}
