package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity  int
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}
