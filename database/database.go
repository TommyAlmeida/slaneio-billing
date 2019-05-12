package database

import (
	"fmt"
	"gamestash.io/billing/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
	"os"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("mysql", dbConfig)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
