package models

import (
	"time"
)

const (
	StatusPending    = "pending"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
	StatusCancelled  = "cancelled"
)

type Repair struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	TechnicianID uint      `gorm:"not null;index" json:"technician_id"`
	NameCustomer string    `gorm:"size:100;not null" json:"name_customer"`
	Description  string    `gorm:"type:text" json:"description"`
	ServiceFee   float64   `gorm:"type:decimal(10,2)" json:"service_fee"`
	Status       string    `gorm:"size:30;not null;default:'pending'" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
