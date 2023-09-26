package model

import "gorm.io/gorm"

type OrderShift struct {
	gorm.Model
	OrderID OrderID `gorm:"index"`
}
