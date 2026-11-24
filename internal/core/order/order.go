package order

import (
	"time"
)

//QUESTION: i have customerId and 'internal'
// for pickup orders that cashier creates?,
// but I also wanna know which cashier it's
// and I dont know how many cashiers i will have later

type Order struct {
	ID              int64
	CustomerID      int64
	Currency        string
	Items           []OrderItem
	Status          Status
	TotalAmount     int64
	DeliveryAddress string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type RequestedItem struct {
	ItemID   int64 `json:"item_id"`
	Quantity int64 `json:"quantity"`
}

func NewOrder(
	customerID int64,
	currency string,
	requested []RequestedItem,
	items []Item,
	deliveryAddress string,
) (Order, error) {
	if customerID <= 0 || len(currency) != 3 || deliveryAddress == "" || len(requested) == 0 {
		return Order{}, ErrInvalidOrder
	}

	itemsMap := make(map[int64]Item, len(items))
	for _, item := range items {
		itemsMap[item.ID] = item
	}

	orderItems := make([]OrderItem, 0, len(requested))
	for _, req := range requested {
		item, ok := itemsMap[req.ItemID]
		if !ok {
			return Order{}, ErrItemNotFound
		}
		if item.Stock < req.Quantity {
			return Order{}, ErrInsufficientStock
		}
		orderItems = append(orderItems, OrderItem{
			ID:        item.ID,
			Quantity:  req.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}

	var total int64
	for _, it := range orderItems {
		if it.Quantity <= 0 || it.UnitPrice < 0 {
			return Order{}, ErrInvalidOrder
		}
		total += it.UnitPrice * it.Quantity
	}

	return Order{
		CustomerID:      customerID,
		Currency:        currency,
		Items:           orderItems,
		Status:          StatusPending,
		TotalAmount:     total,
		DeliveryAddress: deliveryAddress,
	}, nil
}
