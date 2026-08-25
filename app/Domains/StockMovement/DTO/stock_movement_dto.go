package dto

import "time"

type StockInRequest struct {
	ItemID    uint    `json:"item_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" binding:"required,gte=0"`
	Note      string  `json:"note"`
}

type StockOutRequest struct {
	ItemID   uint   `json:"item_id" binding:"required"`
	RepairID uint   `json:"repair_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required,gt=0"`
	Note     string `json:"note"`
}

type StockMovementResponse struct {
	ID        uint      `json:"id"`
	ItemID    uint      `json:"item_id"`
	RepairID  *uint     `json:"repair_id"`
	Type      string    `json:"type"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
