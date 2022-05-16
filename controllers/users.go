package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
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

// POST /login
// User login, return token
var APPLICATION_NAME = "Lesgo"
var LOGIN_EXPIRATION_DURATION = time.Duration(5) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("15 besar")

type M map[string]interface{}

func LoginUser(c *gin.Context) {
	var userInput models.UserInputLogin
	var databaseInput models.UserInputLogin
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.Db.QueryRow("SELECT id, username, password FROM users WHERE username=?", userInput.Username).Scan(&databaseInput.Id, &databaseInput.Username, &databaseInput.Password)

	if err != nil {
		c.JSON(http.StatusMovedPermanently, gin.H{"error": "No username"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databaseInput.Password), []byte(userInput.Password))
	if err != nil {
		c.JSON(http.StatusMovedPermanently, gin.H{"error": "Wrong Password"})
		return
	}

	claims := models.ClaimsJWT{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: userInput.Username,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	tokenString, _ := json.Marshal(M{"token": signedToken})
	fmt.Println(tokenString)
	c.JSON(http.StatusOK, gin.H{"userId": databaseInput.Id, "token": tokenString})
}

// GET /users
// get all users data
func FindUsers(c *gin.Context) {
	fmt.Println("GET /users")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	var stringData string = `{
		"users": [
		  {
			"firstName": "fname",
			"lastName": "lname",
			"ageOfBirth": "10-03-2000",
			"age": 21,
			"sex": "m/f",
			"address": "loremipsum...",
			"occupation": "student",
			"point": 10202,
			"totalRedeem": 100,
			"totalDistance": 2023.5,
			"totalEmissionReduced": 500,
			"badge": 0
		  }
		]
	  }`
	// c.JSON(http.StatusOK, gin.H(stringData))
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(stringData))

}

// GET /user/:userid
// get user data with userid
func FindUserById(c *gin.Context) {
	fmt.Println("GET /user/:userid")
	fmt.Println("User ID:", string(c.Param("userId")))
	var stringData string = `{"firstName":"fname","lastName":"lname","ageOfBirth":"10-03-2000","age":21,"sex":"m/f","address":"loremipsum...","occupation":"student","point":10202,"totalRedeem":100,"totalDistance":2023.5,"totalEmissionReduced":500,"badge":0}`
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(stringData))
}

// PUT /user/:userid
// update user data with userid
func UpdateUserById(c *gin.Context) {
	fmt.Println("GET /user/:userid")
	fmt.Println("User ID:", string(c.Param("userId")))
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	var stringData string = `
	{
		"firstName": "fname",
		"lastName": "lname",
		"ageOfBirth": "10-03-2000",
		"age": 21,
		"sex": "m/f",
		"address": "lorem ipsum...",
		"occupation": "student",
		"point": 10202,
		"totalRedeem" : 100,
		"totalDistance" : 2023.5,
		"totalEmissionReduced" : 500,
		"badge" : 0
	}
	`
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(stringData))
}
