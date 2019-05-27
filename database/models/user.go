package models

import (
	"database/sql"
	"gamestash.io/billing/api/common"
	"github.com/jinzhu/gorm"
)

// User data model
type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string `sql:"unique"`
	PasswordHash string

	Wallet    Wallet
	WalletId  sql.NullInt64
}

func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":         u.ID,
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"wallet": u.Wallet,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.FirstName = m["first_name"].(string)
	u.LastName = m["last_name"].(string)
	u.Email = m["email"].(string)
}
