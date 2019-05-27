package wallet

import (
	"fmt"
	"net/http"

	"gamestash.io/billing/api/common"
	"gamestash.io/billing/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Wallet = models.Wallet
type User = models.User
type JSON = common.JSON

type WithdrawlDepositBody struct {
	Id     uint     `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func getBalance(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	ownerId := c.Param("ownerId")

	var wallet Wallet

	if err := db.Set("gorm:auto_preload", true).Where("owner_i d = ?", ownerId).First(&wallet).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, wallet.Serialize())
}

func GetWalletById(c *gin.Context, id uint) (*Wallet, error) {
	db := c.MustGet("db").(*gorm.DB)

	var exists Wallet

	if err := db.Where("id = ?", id).First(&exists).Error; err == nil {
		return nil, fmt.Errorf("Could not find wallet with id %d", id)
	}

	return &exists, nil
}

func withdrawal(c *gin.Context) {
	var body WithdrawlDepositBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	wallet, err := GetWalletById(c, body.Id)

	if err != nil {
		c.JSON(http.StatusConflict, common.JSON{
			"message": "Could not find wallet with id " + string(body.Id),
		})
		return
	}

	c.JSON(http.StatusOK, common.JSON{
		"data": wallet.Serialize(),
	})
}

func deposit(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var body WithdrawlDepositBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	wallet, err := GetWalletById(c, body.Id)

	if err != nil {
		c.JSON(http.StatusConflict, common.JSON{
			"message": "Could not find wallet with id " + string(body.Id),
		})
		return
	}

	db.Model(&wallet).Update("amount", body.Id)

	updated, err := GetWalletById(c, wallet.ID)

	if err != nil {
		c.JSON(http.StatusConflict, common.JSON{
			"message": "Could not find wallet with id " + string(body.Id),
		})
		return
	}

	c.JSON(http.StatusOK, common.JSON{
		"data": updated.Serialize(),
	})
}
