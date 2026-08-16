package models

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password  string    `json:"-" gorm:"not null" binding:"required,min=6"` // json:"-" menyembunyikan password dari response JSON (seperti $hidden di Laravel)
	Role      string    `json:"role" gorm:"not null" binding:"required,oneof=admin staff customer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
