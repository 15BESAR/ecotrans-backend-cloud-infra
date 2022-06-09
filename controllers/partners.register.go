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

func checkRegisterInputPartner(userInput models.Partner) (valid bool, msg string) {
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

// POST /partner/register
// Register partner
func RegisterPartner(c *gin.Context) {
	var userInput models.Partner
	var databaseInput models.Partner
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Data not complete"})
		return
	}
	if valid, err := checkRegisterInputPartner(userInput); !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err})
		return
	}
	err := models.Db.Where("username = ? OR email = ? OR partner_name = ?", userInput.Username, userInput.Email, userInput.PartnerName).First(&databaseInput).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		// user or email or partner name not found in db, continue to hash
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
		result := models.Db.Session(&gorm.Session{SkipHooks: false}).Create(&databaseInput)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Server error, unable to create your account"})
			return
		}
	case !errors.Is(err, gorm.ErrRecordNotFound):
		switch {
		case (databaseInput.Username == userInput.Username):
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Username has been taken"})
		case (databaseInput.Email == userInput.Email):
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Email has been taken"})
		case (databaseInput.PartnerName == userInput.PartnerName):
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Partner Name already Exist !"})
		default:
			c.JSON(500, gin.H{
				"error": true,
				"msg":   "Please try to register again!"})

		}
		return
	default:
		c.JSON(500, gin.H{
			"error": true,
			"msg":   "Server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"error":  false,
		"status": "Account has been created"})
}
