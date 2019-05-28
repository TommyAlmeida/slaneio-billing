package services

import (
	"fmt"
	"gamestash.io/billing/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"sync"
)

type UserService struct {
	db *gorm.DB
}

type User = models.User

var once sync.Once
var instance *UserService

func newUserService(c *gin.Context) *UserService {
	var us UserService
	us.db = c.MustGet("db").(*gorm.DB)

	return &us
}

func GetInstance(c *gin.Context) *UserService {
	once.Do(func() {
		instance = newUserService(c)
	})

	return instance
}

func (us *UserService) Query(query string, value string) (*User, error){
	var user User

	if err := us.db.Where(query + " = ?", value).First(&user).Error; err != nil {
		return nil, fmt.Errorf("Couldn't find user with %s", query)
	}

	return &user, nil
}


func (us *UserService) Create(body User) (*User, error){
	_, err := us.GetByEmail(body.Email)

	var user User

	if err == nil {
		user = User{
			Email: body.Email,
			FirstName: body.FirstName,
			LastName: body.LastName,
			PasswordHash: body.PasswordHash,
		}
	}

	return &user, nil
}

func (us *UserService) GetByEmail(email string) (*User, error){
	user, err := us.Query("email", email)

	if err != nil {
		return nil, fmt.Errorf("Could not find user with email %s", email)
	}

	return user, nil
}
