package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mohiramansurovna/orders/internal/core/order"
)

type OrdersService struct {
	ordersRepository order.OrdersRepository
	itemsProvider    order.ItemsProvider
}

func NewOrdersService(ordersRepository order.OrdersRepository, itemsProvider order.ItemsProvider) *OrdersService {
	return &OrdersService{ordersRepository: ordersRepository, itemsProvider: itemsProvider}
}

type CreateOrderArgs struct {
	CustomerID      int64
	Currency        string
	Items           []order.RequestedItem
	DeliveryAddress string
}

func (this *OrdersService) CreateOrder(ctx context.Context, args CreateOrderArgs) (order.Order, error) {
	if len(args.Items) == 0 {
		return order.Order{}, order.ErrInvalidOrder
	}

	ids := make([]string, 0, len(args.Items))
	for _, it := range args.Items {
		ids = append(ids, strconv.FormatInt(it.ItemID, 10))
	}

	items, err := this.itemsProvider.GetByIDs(ctx, ids)
	if err != nil {
		return order.Order{}, fmt.Errorf("get items: %w", err)
	}

	created, err := order.NewOrder(
		args.CustomerID,
		args.Currency,
		args.Items,
		items,
		args.DeliveryAddress,
	)
	if err != nil {
		return order.Order{}, fmt.Errorf("create order: %w", err)
	}

	persistend, err := this.ordersRepository.Create(ctx, created)
	if err != nil {
		return order.Order{}, fmt.Errorf("create order: %w", err)
	}
	return persistend, nil
}
