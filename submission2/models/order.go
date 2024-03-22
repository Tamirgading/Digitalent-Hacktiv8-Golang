package models

import "time"

type Item struct {
	ItemID      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}

type Order struct {
	OrderID      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}
