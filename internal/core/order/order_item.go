package order

type OrderItem struct {
	ID        int64 `json:"id"`
	Quantity  int64 `json:"quantity"`
	UnitPrice int64 `json:"unit_price"`
}