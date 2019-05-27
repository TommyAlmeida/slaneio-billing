package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Wallet{})
	db.Model(&User{}).AddForeignKey("wallet_id", "wallets(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")
}
