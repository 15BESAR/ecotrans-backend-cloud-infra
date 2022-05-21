package controllers

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/bearbin/go-age"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kr/pretty"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func checkRegisterInput(userInput models.User) bool {
	// check user input sex -> move it to PUT /user to complete all user data
	// if !(userInput.Gender == "m" || userInput.Gender == "f") {
	// 	return false
	// }
	// check email
	if !isValidEmail(userInput.Email) {
		return false
	}

	return true
}

// POST /register
// Register user
func RegisterUser(c *gin.Context) {
	var userInput models.User
	var databaseInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not complete"})
		return
	}
	if !checkRegisterInput(userInput) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong format"})
		return
	}
	pretty.Println(userInput)
	err := models.Db.Where("username = ? OR email = ?", userInput.Username, userInput.Email).First(&databaseInput).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Server error, unable to create your account"})
			return
		}
		// Update Hashed password
		databaseInput = userInput
		databaseInput.Password = string(hashedPassword)
		// Count Age
		databaseInput.Age = age.Age(databaseInput.BirthDate)
		pretty.Println(databaseInput)
		result := models.Db.Session(&gorm.Session{SkipHooks: false}).Create(&databaseInput)

		if result.error != nil {
			c.JSON(500, gin.H{"error": "Server error, unable to create your account"})
			return
		}
	case !errors.Is(err, gorm.ErrRecordNotFound):
		if databaseInput.Username == userInput.Username && databaseInput.Email == userInput.Email {
			c.JSON(500, gin.H{"error": "Username & email has been taken"})
		} else if databaseInput.Username == userInput.Username {
			c.JSON(500, gin.H{"error": "Username has been taken"})
		} else {
			c.JSON(500, gin.H{"error": "Email has been taken"})
		}

		return
	default:
		c.JSON(500, gin.H{"error": "Server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "Account has been created"})
}
