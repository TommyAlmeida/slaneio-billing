package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	fmt.Println("Auto Migration has beed processed")
}
