package wallet

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/wallet")
	{
		auth.GET("/", getBalance)
		auth.POST("/deposit", deposit)
		auth.POST("/withdrawal", withdrawal)
	}
}
