package models

import "time"

type Package struct {
	ID                uint64           `json:"id" gorm:"primaryKey"`
	CategoryPackageID uint64           `json:"category_package_id" gorm:"not null" binding:"required"`
	CategoryPackage   *CategoryPackage `json:"category_package,omitempty" gorm:"foreignKey:CategoryPackageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Code              string           `json:"code" gorm:"unique;not null" binding:"required"`
	Name              string           `json:"name" gorm:"not null" binding:"required"`
	Price             float64          `json:"price" gorm:"not null" binding:"required"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}
