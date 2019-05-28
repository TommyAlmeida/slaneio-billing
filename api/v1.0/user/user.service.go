package services

import (
	"errors"
	"fmt"
	"gamestash.io/billing/api/structs"
	"gamestash.io/billing/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type User = models.User
type UserRequestBody = structs.UserRequestBody

//TODO: Change to interface in the future
type UserService struct {
	db *gorm.DB
}

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

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func query(db *gorm.DB, query string, value interface{}) (*User, error) {
	var user User

	if err := db.Where(query+" = ?", value).First(&user).Error; err == nil {
		return nil, fmt.Errorf("Couldn't find user with %s", query)
	}

	return &user, nil
}

func (us *UserService) Create(body UserRequestBody) (*User, error) {
	_, err := us.GetByEmail(body.Email)

	var user User

	if err != nil {
		return nil, err
	}

	hash, hashErr := hash(body.Password)

	if hashErr != nil {
		return nil, errors.New("Could not hash password")
	}

	user = User{
		FirstName:     body.FirstName,
		LastName:  body.LastName,
		Email:  body.Email,
		PasswordHash: hash,
	}

	us.db.NewRecord(user)
	us.db.Create(user)

	return &user, nil
}

func (us *UserService) GetById(id uint) (*User, error){
	user, err := query(us.db, "id", id)

	if err != nil {
		return nil, fmt.Errorf("Could not find user with id %d", id)
	}

	return user, nil
}

//FIX: Kids use unit tests to avoid this stupid errors
func (us *UserService) GetByEmail(email string) (*User, error) {
	user, err := query(us.db, "email", email)

	if err != nil {
		return nil, fmt.Errorf("Could not find user with email %s", email)
	}

	return user, nil
}
