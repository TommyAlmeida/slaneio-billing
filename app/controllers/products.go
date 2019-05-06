package controllers

import (
	"gamestash.io/billing/app/model"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetAllProducts(db *gorm.DB, w http.ResponseWriter){
	products := model.Product{}
	db.Find(&products)

	respondJSON(w, http.StatusOK, products)
}
