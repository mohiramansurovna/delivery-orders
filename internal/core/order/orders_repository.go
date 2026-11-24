package order

import "context"

type OrdersRepository interface {
	Create(ctx context.Context, o Order) (Order, error)
}
