package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"size:100;not null" json:"name"`
	SKU            string         `gorm:"size:50" json:"sku"`
	QuantityOnHand int            `gorm:"not null;default:0" json:"quantity_on_hand"`
	SellingPrice   float64        `gorm:"type:decimal(10,2)" json:"selling_price"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
