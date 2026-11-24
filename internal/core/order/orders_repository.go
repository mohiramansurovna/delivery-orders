package order

import "context"

type OrdersRepository interface {
	Create(ctx context.Context, o Order) (Order, error)
	GetByID(ctx context.Context, ID int64) (Order, error)
}
