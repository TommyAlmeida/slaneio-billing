package model

import (
	"github.com/jinzhu/gorm"
)

type QuantityType struct {
	gorm.Model
	Name  string
	Code  string
	Value float64 `gorm:"default:0"`
}

type Product struct {
	gorm.Model
	Title        string       `json:"status"`
	Active       string       `json:"active"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{}, &QuantityType{})

	return db
}
