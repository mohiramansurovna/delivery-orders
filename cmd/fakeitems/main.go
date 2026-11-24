// Fake items service for manual testing. Dev tool only.
//
//	go run ./cmd/fakeitems
//	curl 'localhost:9000/items?ids=1,2,3'
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type item struct {
	ID        int64 `json:"id"`
	Stock     int64 `json:"stock"`
	UnitPrice int64 `json:"unit_price"`
}

// 10 items. Prices are in minor units (cents).
// Item 5 has stock 0 (test insufficient stock), item 6 has stock 1.
var items = []item{
	{ID: 1, Stock: 100, UnitPrice: 500},
	{ID: 2, Stock: 50, UnitPrice: 1250},
	{ID: 3, Stock: 20, UnitPrice: 300},
	{ID: 4, Stock: 10, UnitPrice: 4500},
	{ID: 5, Stock: 0, UnitPrice: 800},
	{ID: 6, Stock: 1, UnitPrice: 2000},
	{ID: 7, Stock: 75, UnitPrice: 150},
	{ID: 8, Stock: 30, UnitPrice: 990},
	{ID: 9, Stock: 5, UnitPrice: 7500},
	{ID: 10, Stock: 200, UnitPrice: 100},
}

func main() {
	byID := make(map[int64]item, len(items))
	for _, it := range items {
		byID[it.ID] = it
	}

	http.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		// Unknown ids are simply left out, like a real service would.
		found := []item{}
		for _, raw := range strings.Split(r.URL.Query().Get("ids"), ",") {
			id, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 64)
			if err != nil {
				continue
			}
			if it, ok := byID[id]; ok {
				found = append(found, it)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(found)
	})

	log.Println("fake items service on :9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
