package main

import "context"

type Order struct {
	ID           string      `json:"id,omitempty" `
	CustomerID   string      `json:"customer_id" binding:"required"`
	Status       string      `json:"status" binding:"required"`
	CreatedOn    int64       `json:"created_on,omitempty"`
	RestaurantId string      `json:"restaurant_id" binding:"required"`
	OrderItems   []OrderItem `json:"order_items,omitempty"`
}

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	Name        string  `json:"name"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Repository interface {
	CreateOrder(ctx context.Context, order Order) error
}
