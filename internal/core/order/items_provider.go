package order

import "context"

type Item struct {
	ID        int64
	Stock     int64
	UnitPrice int64
}

type ItemsProvider interface {
	GetByIDs(ctx context.Context, ids []int64) ([]Item, error)
}
