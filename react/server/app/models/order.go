package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TotalPrice  float64
	OrderItems  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID    uint
	ProductID  uint
	Product    Product
	Quantity   uint
}