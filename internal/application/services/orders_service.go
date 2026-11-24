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

func (this *OrdersService) Create(ctx context.Context, args CreateOrderArgs) (order.Order, error) {
	if len(args.Items) == 0 {
		return order.Order{}, order.ErrInvalidOrder
	}

	IDs := make([]string, 0, len(args.Items))
	for _, it := range args.Items {
		IDs = append(IDs, strconv.FormatInt(it.ItemID, 10))
	}

	items, err := this.itemsProvider.GetByIDs(ctx, IDs)
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

func (this *OrdersService) GetByID(ctx context.Context, ID int64) (order.Order, error) {
	return this.ordersRepository.GetByID(ctx, ID)
}
