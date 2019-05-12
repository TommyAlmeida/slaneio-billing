package models

import (
	"gamestash.io/billing/api/common"
	"github.com/jinzhu/gorm"
)

type Wallet struct {
	gorm.Model
	Amount float64
	Owner   User   `gorm:"foreignkey:UserID"`
	OwnerID uint
	//TODO: Add transactions
}

func (u *Wallet) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"amount":     u.Amount,
		"owner": u.Owner.Serialize(),
	}
}
