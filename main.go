package main

import (
	"fmt"
	"gamestash.io/billing/api/middlewares"
	"gamestash.io/billing/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())

	if len(port) <= 0 || port == "" {
		fmt.Printf("Could not connect to port %s", port)
	}
	app.Run(":" + port)
}
