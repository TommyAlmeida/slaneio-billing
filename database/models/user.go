package models

import (
	"gamestash.io/billing/api/common"
	"github.com/jinzhu/gorm"
)

// User data model
type User struct {
	gorm.Model
	FirstName     string
	LastName     string
	Email     string
	PasswordHash string
}

func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"email":     u.Email,
		"first_name": u.FirstName,
		"last_name": u.LastName,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.FirstName = m["first_name"].(string)
	u.LastName = m["last_name"].(string)
	u.Email = m["email"].(string)
}