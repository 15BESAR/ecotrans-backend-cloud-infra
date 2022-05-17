package controllers

import (
	"database/sql"
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func checkRegisterInput(username string, email string, pass string) bool {
	return true
}

// POST /register
// Register user
func RegisterUser(c *gin.Context) {
	var userInput models.UserInputRegis
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No binding"})
		return
	}
	if !checkRegisterInput(userInput.Username, userInput.Email, userInput.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong format"})
		return
	}

	var usertemp string
	var emailtemp string
	err := models.Db.QueryRow("SELECT username FROM users WHERE username=? OR email=?", userInput.Username, userInput.Email).Scan(&usertemp, &emailtemp)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Server error, unable to create your account"})
			return
		}

		_, err = models.Db.Exec("INSERT INTO users(username,email,password) VALUES(?,?,?)", userInput.Username, userInput.Email, hashedPassword)
		if err != nil {
			c.JSON(500, gin.H{"error": "Server error, unable to create your account"})
			return
		}
	case err != nil:
		c.JSON(500, gin.H{"error": "Username or email has been taken"})
		return
	default:
		c.JSON(500, gin.H{"error": "Server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Account has been created"})
}
