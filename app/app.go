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
)

type App struct {
	BillingRouter BillingRouter
	DB            *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)

	println(dbURI)

	db, err := gorm.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatalf("Could not connect database, %s", err)
	}

	a.DB = model.DBMigrate(db)

	a.BillingRouter = BillingRouter{
		mux.NewRouter(),
		a.DB,
	}
	a.BillingRouter.registerRoutes()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.BillingRouter.Router))
}
