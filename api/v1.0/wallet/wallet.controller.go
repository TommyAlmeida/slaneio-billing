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

type WithdrawlDepositBody struct {
	Id int `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

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

func GetWalletById(c *gin.Context, id int) (Wallet, error) {
	db := c.MustGet("db").(*gorm.DB)

	var exists Wallet

	if err := db.Where("id = ?", id).First(&exists).Error; err == nil{
		return nil, fmt.Errorf("Could not find wallet with id %d", id)
	}

	return exists, nil
}

func withdrawal(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var body WithdrawlDepositBody

	if err := c.BindJSON(&body); err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	wallet, err := GetWalletById(body.Id)

	if err != nil {
		c.JSON(http.StatusConflict, common.JSON{
			"message": "Could not find wallet with id " + string(body.Id),
		})
		return
	}

	fmt.Printf("Wallet %+T", wallet)
}

func deposit(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var body WithdrawlDepositBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var exists
}
