package controllers

import (
	"errors"
	"net/http"
	"net/mail"
	"unicode"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/bearbin/go-age"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func verifyPassword(s string) (sixOrMore, number, upper bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		// case unicode.IsPunct(c) || unicode.IsSymbol(c):
		//     special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sixOrMore = letters >= 6
	return
}
func checkRegisterInput(userInput models.User) (valid bool, msg string) {
	// check email
	if !isValidEmail(userInput.Email) {
		return false, "Email not valid"
	}
	sixOrMore, number, upper := verifyPassword(userInput.Password)
	if !sixOrMore {
		return false, "Password must have length six or more"
	}
	if !number {
		return false, "Password need consists of number"
	}

	if !upper {
		return false, "Password must have upper case"
	}
	return true, "Success"
}

// POST /user/register
// Register user
func RegisterUser(c *gin.Context) {
	var userInput models.User
	var databaseInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Data not complete",
		})
		return
	}
	if valid, msg := checkRegisterInput(userInput); !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   msg,
		})
		return
	}
	err := models.Db.Where("username = ? OR email = ?", userInput.Username, userInput.Email).First(&databaseInput).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Server error, unable to create your account"})
			return
		}
		// Update Hashed password
		databaseInput = userInput
		databaseInput.Password = string(hashedPassword)
		// Count Age
		databaseInput.Age = age.Age(databaseInput.BirthDate)
		result := models.Db.Session(&gorm.Session{SkipHooks: false}).Create(&databaseInput)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Server error, unable to create your account",
			})
			return
		}
	case !errors.Is(err, gorm.ErrRecordNotFound):
		if databaseInput.Username == userInput.Username && databaseInput.Email == userInput.Email {
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Username & email has been taken",
			})
		} else if databaseInput.Username == userInput.Username {
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Username has been taken",
			})
		} else {
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Email has been taken",
			})
		}

		return
	default:
		c.JSON(500, gin.H{
			"error": true,
			"msg":   "Server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"error":  false,
		"status": "Account has been created"})
}
