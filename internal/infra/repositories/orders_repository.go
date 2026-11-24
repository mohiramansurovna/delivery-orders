package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mohiramansurovna/orders/internal/core/order"
	"github.com/mohiramansurovna/orders/internal/infra/mappers"
)

type OrdersRepository struct {
	pool *pgxpool.Pool
}

func NewOrdersRepsitory(pool *pgxpool.Pool) *OrdersRepository {
	return &OrdersRepository{pool: pool}
}

var _ order.OrdersRepository = (*OrdersRepository)(nil)

func (this *OrdersRepository) Create(ctx context.Context, o order.Order) (order.Order, error) {
	const q = `
		INSERT INTO orders (customer_id, currency, items, status, total_amount, delivery_address)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, customer_id, currency, items, status, total_amount, delivery_address, created_at, updated_at`

	rec := mappers.ToRecord(o)
	var out mappers.OrderRecord
	err := this.pool.QueryRow(ctx, q,
		rec.CustomerID, rec.Currency, rec.Items, rec.Status, rec.TotalAmount, rec.DeliveryAddress,
	).Scan(
		&out.ID, &out.CustomerID, &out.Currency, &out.Items, &out.Status,
		&out.TotalAmount, &out.DeliveryAddress, &out.CreatedAt, &out.UpdatedAt,
	)
	if err != nil {
		return order.Order{}, err
	}
	return mappers.ToOrder(out), nil
}

func (this *OrdersRepository) GetByID(ctx context.Context, ID int64) (order.Order, error) {
	const q = `
        SELECT id, customer_id, currency, items, status, total_amount, delivery_address, created_at, updated_at
        FROM orders WHERE id = $1`

	var out mappers.OrderRecord
	//QUESTION: what if i forgot the order, will it find?
	err := this.pool.QueryRow(ctx, q, ID).Scan(
		&out.ID, &out.CustomerID, &out.Currency, &out.Items, &out.Status,
		&out.TotalAmount, &out.DeliveryAddress, &out.CreatedAt, &out.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return order.Order{}, order.ErrNotFound

	}
	if err != nil {
		return order.Order{}, fmt.Errorf("get order %d: %w", ID, err)
	}
	return mappers.ToOrder(out), nil

}
