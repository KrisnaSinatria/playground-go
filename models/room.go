package models

import "time"

type Room struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" form:"name" gorm:"not null" binding:"required"`
	No        string    `json:"no" form:"no" gorm:"unique;not null" binding:"required"`
	Img       string    `json:"img" form:"img"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
