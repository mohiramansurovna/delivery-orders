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
	ItemID   int64
	Quantity int64
}

func validCurrency(c string) bool {
	if len(c) != 3 {
		return false
	}
	for i := 0; i < 3; i++ {
		if c[i] < 'A' || c[i] > 'Z' {
			return false
		}
	}
	return true
}

func NewOrder(
	customerID int64,
	currency string,
	requested []RequestedItem,
	items []Item,
	deliveryAddress string,
) (Order, error) {
	if customerID <= 0 || !validCurrency(currency) || deliveryAddress == "" || len(requested) == 0 {
		return Order{}, ErrInvalidOrder
	}

	itemsMap := make(map[int64]Item, len(items))
	for _, item := range items {
		itemsMap[item.ID] = item
	}

	seenItems := make(map[int64]struct{}, len(requested))

	var total int64
	orderItems := make([]OrderItem, 0, len(requested))
	for _, req := range requested {

		//throw on dublicates
		if _, exists := seenItems[req.ItemID]; exists {
			return Order{}, ErrInvalidOrder
		}
		seenItems[req.ItemID] = struct{}{}

		item, ok := itemsMap[req.ItemID]
		if !ok {
			return Order{}, ErrItemNotFound
		}
		if req.Quantity <= 0 {
			return Order{}, ErrInvalidOrder
		}
		if item.Stock < req.Quantity { //QUESTION: will i somehow lock stock for concurrency?
			return Order{}, ErrInsufficientStock
		}
		total += item.UnitPrice * req.Quantity
		orderItems = append(orderItems, OrderItem{
			ItemID:    item.ID,
			Quantity:  req.Quantity,
			UnitPrice: item.UnitPrice,
		})
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
