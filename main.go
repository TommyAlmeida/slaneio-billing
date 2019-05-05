package main

import (
	"gamestash.io/billing/app"
	"gamestash.io/billing/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}