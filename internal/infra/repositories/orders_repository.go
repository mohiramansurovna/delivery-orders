package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mohiramansurovna/orders/internal/core/order"
)

type OrdersRepository struct {
	pool *pgxpool.Pool
}

func NewOrdersRepsitory(pool *pgxpool.Pool) *OrdersRepository {
	return &OrdersRepository{pool: pool}
}

var _ order.OrdersRepository = (*OrdersRepository)(nil)

func (r *OrdersRepository) Create(ctx context.Context, o order.Order) (order.Order, error) {
	const q = `
		INSERT INTO orders (customer_id, currency, items, status, total_amount, delivery_address)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, customer_id, currency, items, status, total_amount, delivery_address, created_at, updated_at`

	if o.Items == nil {
		o.Items = []order.OrderItem{}
	}

	var out order.Order
	err := r.pool.QueryRow(ctx, q,
		o.CustomerID, o.Currency, o.Items, o.Status, o.TotalAmount, o.DeliveryAddress,
	).Scan(
		&out.ID, &out.CustomerID, &out.Currency, &out.Items, &out.Status,
		&out.TotalAmount, &out.DeliveryAddress, &out.CreatedAt, &out.UpdatedAt,
	)
	if err != nil {
		return order.Order{}, err
	}
	return out, nil
}
