package auth

import (
	"fmt"
	"gamestash.io/billing/api/common"
	"gamestash.io/billing/database/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type User = models.User

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(data common.JSON) (string, error) {
	date := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)

	if readErr != nil {
		return "", readErr
	}

	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		FirstName    string `json:"first_name" binding:"required"`
		LastName string `json:"last_name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	var body RequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// check existancy
	var exists User

	if err := db.Where("email = ?", body.Email).First(&exists).Error; err == nil {
		c.JSON(http.StatusConflict, common.JSON{
			"message": "Could not create user, email already exists!",
		})
		return
	}

	hash, hashErr := hash(body.Password)
	if hashErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// create user
	user := User{
		FirstName:     body.FirstName,
		LastName:  body.LastName,
		Email:  body.Email,
		PasswordHash: hash,
	}

	//Create a user wallet such is required
	wallet := models.Wallet{
		Amount: 0,
		Owner: user,
	}

	db.NewRecord(user)
	db.Create(&user)

	//Create new wallet on the database
	db.NewRecord(wallet)
	db.Create(&wallet)

	db.Model(&user).Update("wallet", &wallet)

	serialized := user.Serialize()
	token, _ := generateToken(serialized)

	const maxAge = 60 * 60 * 24 * 7 //7 days
	c.SetCookie("token", token, maxAge , "/", "", false, true)


	response := common.JSON{
		"token": token,
		"data":  common.JSON{
			"user": user.Serialize(),
			"wallet": wallet.Serialize(),
		},
	}

	c.JSON(http.StatusOK, response)
}

func login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// check existancy
	var user User
	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if !checkHash(body.Password, user.PasswordHash) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	serialized := user.Serialize()
	token, _ := generateToken(serialized)

	const maxAge = 60 * 60 * 24 * 7 //7 days
	c.SetCookie("token", token, maxAge, "/", "", false, true)

	c.JSON(http.StatusOK, common.JSON{
		"user":  user.Serialize(),
		"token": token,
	})
}

func check(c *gin.Context) {
	userRaw, ok := c.Get("user")

	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user := userRaw.(User)

	tokenExpire := int64(c.MustGet("token_expire").(float64))
	now := time.Now().Unix()
	diff := tokenExpire - now

	fmt.Println(diff)

	const diffDays = 60 * 60 * 24 * 3 //three days

	if diff < diffDays{
		// renew token
		token, _ := generateToken(user.Serialize())

		const maxAge = 60 * 60 * 24 * 7 //7 days
		c.SetCookie("token", token, maxAge, "/", "", false, true)
		c.JSON(http.StatusOK, common.JSON{
			"token": token,
			"user":  user.Serialize(),
		})
		return
	}

	c.JSON(http.StatusOK, common.JSON{
		"token": nil,
		"user":  user.Serialize(),
	})
}