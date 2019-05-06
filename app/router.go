package app

import (
	"gamestash.io/billing/app/controllers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

type BillingRouter struct {
	Router *mux.Router
	DB *gorm.DB
}
func (br *BillingRouter) registerRoutes(){
	br.Get("/products", br.GetProducts)
}

func (br *BillingRouter) GetProducts(w http.ResponseWriter, r *http.Request) {
	controllers.GetAllProducts(br.DB, w)
}

func (br *BillingRouter) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	br.Router.HandleFunc(path, f).Methods("GET")
}

func (br *BillingRouter) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	br.Router.HandleFunc(path, f).Methods("POST")
}

func (br *BillingRouter) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	br.Router.HandleFunc(path, f).Methods("PUT")
}

func (br *BillingRouter) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	br.Router.HandleFunc(path, f).Methods("DELETE")
}
