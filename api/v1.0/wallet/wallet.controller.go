package wallet

import (
	"gamestash.io/billing/api/common"
	"gamestash.io/billing/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Wallet = models.Wallet
type User = models.User
type JSON = common.JSON

func getBalance(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	ownerId := c.Param("ownerId")

	var wallet Wallet

	if err := db.Set("gorm:auto_preload", true).Where("ownerId = ?", ownerId).First(&wallet).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, wallet.Serialize())
}

func withdrawal(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type Body struct {
		Id int `json:"id" binding:"required"`
		Amount float64 `json:"amount" binding:"required"`
	}

	var body Body

	if err := c.BindJSON(&body); err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var exists Wallet

	if err := db.Where("id = ?", body.Id).First(&exists).Error; err == nil{
		c.JSON(http.StatusConflict, common.JSON{
			"message": "Could not find wallet with id " + string(body.Id),
		})

		return
	}
}

func deposit(c *gin.Context) {

}
