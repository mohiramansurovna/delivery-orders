package mappers

import (
	"time"

	"github.com/mohiramansurovna/orders/internal/core/order"
)

type OrderRecord struct {
	ID              int64             `db:"id"`
	CustomerID      int64             `db:"customer_id"`
	Currency        string            `db:"currency"`
	Items           []OrderItemRecord `db:"items"`
	Status          string            `db:"status"`
	TotalAmount     int64             `db:"total_amount"`
	DeliveryAddress string            `db:"delivery_address"`
	CreatedAt       time.Time         `db:"created_at"`
	UpdatedAt       time.Time         `db:"updated_at"`
}

type OrderItemRecord struct {
	ID        int64 `json:"id"`
	Quantity  int64 `json:"quantity"`
	UnitPrice int64 `json:"unit_price"`
}

func ToRecord(o order.Order) OrderRecord {
	items := make([]OrderItemRecord, 0, len(o.Items))
	for _, it := range o.Items {
		items = append(items, OrderItemRecord{ID: it.ItemID, Quantity: it.Quantity, UnitPrice: it.UnitPrice})
	}
	return OrderRecord{
		ID:              o.ID,
		CustomerID:      o.CustomerID,
		Currency:        o.Currency,
		Items:           items,
		Status:          string(o.Status),
		TotalAmount:     o.TotalAmount,
		DeliveryAddress: o.DeliveryAddress,
		CreatedAt:       o.CreatedAt,
		UpdatedAt:       o.UpdatedAt,
	}
}

func ToOrder(r OrderRecord) order.Order {
	items := make([]order.OrderItem, 0, len(r.Items))
	for _, it := range r.Items {
		items = append(items, order.OrderItem{ItemID: it.ID, Quantity: it.Quantity, UnitPrice: it.UnitPrice})
	}
	return order.Order{
		ID:              r.ID,
		CustomerID:      r.CustomerID,
		Currency:        r.Currency,
		Items:           items,
		Status:          order.Status(r.Status),
		TotalAmount:     r.TotalAmount,
		DeliveryAddress: r.DeliveryAddress,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}
}
