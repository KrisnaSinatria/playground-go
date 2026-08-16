package models

import "time"

type CategoryPackage struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Packages  []Package `json:"packages" gorm:"foreignKey:CategoryPackageID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
