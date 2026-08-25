package dto

import (
	stockDTO "go-first/app/Domains/StockMovement/DTO"
	"time"
)

type CreateRepairRequest struct {
	TechnicianID uint    `json:"technician_id" binding:"required"`
	NameCustomer string  `json:"name_customer" binding:"required"`
	Description  string  `json:"description"`
	ServiceFee   float64 `json:"service_fee" binding:"required,gte=0"`
}

type UpdateRepairStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending in_progress completed cancelled"`
}

type RepairResponse struct {
	ID           uint      `json:"id"`
	TechnicianID uint      `json:"technician_id"`
	NameCustomer string    `json:"name_customer"`
	Description  string    `json:"description"`
	ServiceFee   float64   `json:"service_fee"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RepairDetailResponse struct {
	ID             uint                           `json:"id"`
	TechnicianID   uint                           `json:"technician_id"`
	TechnicianName string                         `json:"technician_name,omitempty"`
	NameCustomer   string                         `json:"name_customer"`
	Description    string                         `json:"description"`
	ServiceFee     float64                        `json:"service_fee"`
	Status         string                         `json:"status"`
	UsedItems      []stockDTO.StockMovementResponse `json:"used_items"`
	TotalCost      float64                        `json:"total_cost"`
	CreatedAt      time.Time                      `json:"created_at"`
	UpdatedAt      time.Time                      `json:"updated_at"`
}
