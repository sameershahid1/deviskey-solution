package model

import (
	"gorm.io/gorm"
)

type VehiclePart struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}
