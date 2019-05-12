package main

import (
	"fmt"
	"gamestash.io/billing/api"
	"gamestash.io/billing/api/middlewares"
	"gamestash.io/billing/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	err := godotenv.Load(pwd + "/.env")

	if err != nil {
		panic(err)
	}
	
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())

	if len(port) <= 0 || port == "" {
		fmt.Printf("Could not connect to port %s", port)
	}

	api.ApplyRoutes(app)
	app.Run(":" + port)
}
