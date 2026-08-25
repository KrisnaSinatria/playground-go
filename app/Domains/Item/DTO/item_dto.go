package dto

import "time"

type CreateItemRequest struct {
	Name         string  `json:"name" binding:"required"`
	SKU          string  `json:"sku"`
	SellingPrice float64 `json:"selling_price" binding:"required,gte=0"`
}

type UpdateItemRequest struct {
	Name         string  `json:"name" binding:"required"`
	SKU          string  `json:"sku"`
	SellingPrice float64 `json:"selling_price" binding:"required,gte=0"`
}

type ItemResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SKU            string    `json:"sku"`
	QuantityOnHand int       `json:"quantity_on_hand"`
	SellingPrice   float64   `json:"selling_price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
