package model

import (
	"github.com/jinzhu/gorm"
	"math/big"
)

type QuantityType struct {
	Name string
	Value big.Rat
}

type Product struct {
	Title string `json:"status"`
	Active string `json:"active"`
	QuantityType QuantityType `json:"active"`
	Status string `gorm:"type:ENUM( 'Not Payed', 'Declined', 'Disabled', 'Processing', 'On Hold', 'Complete');default:'0'" json:"status"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{})
	return db
}
