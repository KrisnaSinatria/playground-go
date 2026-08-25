package dto

import "time"

type CreateTechnicianRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
}

type UpdateTechnicianRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
}

type TechnicianResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
