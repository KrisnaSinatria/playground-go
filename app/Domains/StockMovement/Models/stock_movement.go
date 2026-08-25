package models

import (
	"time"
)

const (
	MovementIn  = "IN"
	MovementOut = "OUT"
)

type StockMovement struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ItemID    uint      `gorm:"not null;index" json:"item_id"`
	RepairID  *uint     `gorm:"index" json:"repair_id"`
	Type      string    `gorm:"size:10;not null" json:"type"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	UnitPrice float64   `gorm:"type:decimal(10,2)" json:"unit_price"`
	Note      string    `gorm:"size:255" json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
