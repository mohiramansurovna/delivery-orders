package presenters

import (
	"time"

	"github.com/mohiramansurovna/orders/internal/core/order"
)

type OrderItemView struct {
	ItemID    int64 `json:"item_id"`
	Quantity  int64 `json:"quantity"`
	UnitPrice int64 `json:"unit_price"`
}

type OrderView struct {
	ID              int64           `json:"id"`
	CustomerID      int64           `json:"customer_id"`
	Currency        string          `json:"currency"`
	Items           []OrderItemView `json:"items"`
	Status          string          `json:"status"`
	TotalAmount     int64           `json:"total_amount"`
	DeliveryAddress string          `json:"delivery_address"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

func PresentOrder(o order.Order) OrderView {
	items := make([]OrderItemView, 0, len(o.Items))
	for _, it := range o.Items {
		items = append(items, OrderItemView{
			ItemID:    it.ID,
			Quantity:  it.Quantity,
			UnitPrice: it.UnitPrice,
		})
	}
	return OrderView{
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
