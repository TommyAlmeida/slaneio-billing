package database

import (
	"fmt"
	"gamestash.io/billing/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.NewConfig())

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	//models.Migrate(database)
	return db, err
}
