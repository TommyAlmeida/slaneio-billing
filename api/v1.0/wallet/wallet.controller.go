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

}

func deposit(c *gin.Context) {

}
