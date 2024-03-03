package model

import (
	"time"

	"gorm.io/gorm"
)

type VehiclePart struct {
	gorm.Model
	Id          uint `gorm:"primaryKey`
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time
}
