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
	QuantityType QuantityType `gorm:"foreignkey:QuantityType" json:"quantityType"`
	Status       string       `gorm:"type:ENUM( 'Not Payed', 'Declined', 'Disabled', 'Processing', 'On Hold', 'Complete');default:'0'" json:"status"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{}, &QuantityType{})
	db.Model(&Product{}).Related(&QuantityType{})

	return db
}
