package dtos

import (
	"github.com/mohiramansurovna/orders/internal/core/order"
)

type CreateOrderItemRequest struct {
        ItemID   int64 `json:"item_id"`
        Quantity int64 `json:"quantity"`
}

type CreateOrderRequest struct {
        CustomerID      int64                    `json:"customer_id"`
        Currency        string                   `json:"currency"`
        DeliveryAddress string                   `json:"delivery_address"`
        Items           []CreateOrderItemRequest `json:"items"`
}

func (r CreateOrderRequest) ToRequestedItems() []order.RequestedItem {
        items := make([]order.RequestedItem, 0, len(r.Items))
        for _, it := range r.Items {
                items = append(items, order.RequestedItem{ItemID: it.ItemID, Quantity: it.Quantity})
        }
        return items
}
