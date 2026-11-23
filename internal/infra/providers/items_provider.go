package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func (p *ItemsProvider) GetByIDs(ctx context.Context, ids []string) ([]order.Item, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no item ids in getByIDs")
	}

	url := fmt.Sprintf("%s?ids=%s", p.baseURL, strings.Join(ids, ","))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	res, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call items service: %w", err)
	}
	defer res.Body.Close() //QUESTION: why do i need it?
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("items service returned status %d", res.StatusCode)
	}

	var dbItems []itemDTO
	if err := json.NewDecoder(res.Body).Decode(&dbItems); err != nil {
		return nil, fmt.Errorf("decode items response: %w", err)
	}

	return toItem(dbItems), nil
}
