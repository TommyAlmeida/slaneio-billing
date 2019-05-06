package app

import (
	"fmt"
	"gamestash.io/billing/app/model"
	"gamestash.io/billing/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
)

type App struct {
	BillingRouter *BillingRouter
	DB            *gorm.DB
	Logger        *log.Logger
}

func (a *App) Initialize(config *config.Config) {
	a.Logger = a.writeLogToFile()

	dbURI := fmt.Sprintf("%s:%s@/%s?parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)

	db, err := gorm.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatalf("Could not connect database, %s", err)
	}

	log.Print("Database successfully connected!")

	a.DB = model.DBMigrate(db)
	a.BillingRouter = &BillingRouter{mux.NewRouter(), a.DB}
	a.BillingRouter.registerRoutes()
	log.Print("Routes registered successfully!")
}

func (a *App) writeLogToFile() *log.Logger{
	f, err := os.OpenFile("billing.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(f, "Billing", 0)

	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	logger.SetOutput(f)

	return logger
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.BillingRouter.Router))
}
