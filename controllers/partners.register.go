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

func checkRegisterInputPartner(userInput models.Partner) bool {
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

// POST /partner/register
// Register partner
func RegisterPartner(c *gin.Context) {
	var userInput models.Partner
	var databaseInput models.Partner
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not complete"})
		return
	}
	if !checkRegisterInputPartner(userInput) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong format"})
		return
	}
	err := models.Db.Where("username = ? OR email = ? OR partner_name = ?", userInput.Username, userInput.Email, userInput.PartnerName).First(&databaseInput).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		// user or email or partner name not found in db, continue to hash
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Server error, unable to create your account"})
			return
		}
		// Update Hashed password
		databaseInput = userInput
		databaseInput.Password = string(hashedPassword)
		result := models.Db.Session(&gorm.Session{SkipHooks: false}).Create(&databaseInput)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Server error, unable to create your account"})
			return
		}
	case !errors.Is(err, gorm.ErrRecordNotFound):
		switch {
		case (databaseInput.Username == userInput.Username):
			c.JSON(500, gin.H{"error": "Username has been taken"})
		case (databaseInput.Email == userInput.Email):
			c.JSON(500, gin.H{"error": "Email has been taken"})
		case (databaseInput.PartnerName == userInput.PartnerName):
			c.JSON(500, gin.H{"error": "Partner Name already Exist !"})
		default:
			c.JSON(500, gin.H{"error": "Please try to register again!"})

		}
		return
	default:
		c.JSON(500, gin.H{"error": "Server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "Account has been created"})
}
