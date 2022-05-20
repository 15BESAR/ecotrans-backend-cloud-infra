package controllers

import (
	"errors"
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func checkRegisterInput(userInput models.User) bool {
	// check user input sex
	if !(userInput.Sex == "m" || userInput.Sex == "f") {
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
	err := models.Db.Where("username = ? OR email = ?", userInput.Username, userInput.Email).First(&databaseInput).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Server error, unable to create your account"})
			return
		}

		databaseInput = userInput
		databaseInput.Password = string(hashedPassword)
		models.Db.Create(&databaseInput)
		if err != nil {
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
