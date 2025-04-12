package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
}

func PriceGreaterThan(min float64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price >= ?", min)
	}
}

func PriceSmallerThan(max float64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price <= ?", max)
	}
}

func FilterByCategory(name string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.name = ?", name)
	}
}