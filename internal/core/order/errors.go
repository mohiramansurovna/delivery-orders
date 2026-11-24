package order

import "errors"

var (
	ErrNotFound          = errors.New("order not found")
	ErrInvalidOrder      = errors.New("invalid order")
	ErrItemNotFound      = errors.New("order item not found")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrItemsUnavailable  = errors.New("items service unavailable")
)
