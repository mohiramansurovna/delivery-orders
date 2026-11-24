package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mohiramansurovna/orders/internal/core/order"
)

var _ order.ItemsProvider = (*ItemsProvider)(nil)

type ItemsProvider struct {
	baseURL string
	client  *http.Client
}

func NewItemsProvider(baseURL string) *ItemsProvider {
	return &ItemsProvider{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

type itemDTO struct {
	ID        int64 `json:"id"`
	Stock     int64 `json:"stock"`
	UnitPrice int64 `json:"unit_price"`
}

func toItem(dto []itemDTO) []order.Item {
	var orderItems []order.Item
	for _, item := range dto {
		orderItems = append(orderItems, order.Item{
			ID:        item.ID,
			Stock:     item.Stock,
			UnitPrice: item.UnitPrice,
		})
	}
	return orderItems
}

func (p *ItemsProvider) GetByIDs(ctx context.Context, ids []int64) ([]order.Item, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no item ids in getByIDs")
	}

	stringIds := make([]string, len(ids))
	for _, id := range ids {
		stringIds = append(stringIds, strconv.FormatInt(id, 10))
	}
	url := fmt.Sprintf("%s?ids=%s", p.baseURL, strings.Join(stringIds, ","))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	res, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: call items service: %v", order.ErrItemsUnavailable, err)
	}
	defer res.Body.Close() //QUESTION: why do i need it?
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: items service returned status %d", order.ErrItemsUnavailable, res.StatusCode)
	}

	var dbItems []itemDTO
	if err := json.NewDecoder(res.Body).Decode(&dbItems); err != nil {
		return nil, fmt.Errorf("%w: decode items response: %v", order.ErrItemsUnavailable, err)
	}

	return toItem(dbItems), nil
}
