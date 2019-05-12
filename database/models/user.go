package models

import (
	"gamestash.io/billing/api"
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

func (u *User) Read(m api.JSON) {
	u.ID = uint(m["id"].(float64))
	u.FirstName = m["first_name"].(string)
	u.LastName = m["last_name"].(string)
	u.Email = m["email"].(string)
}