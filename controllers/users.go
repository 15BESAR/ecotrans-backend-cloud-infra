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
			Issuer:    models.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(models.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: userInput.Username,
	}

	token := jwt.NewWithClaims(
		models.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(models.JWT_SIGNATURE_KEY)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	json.Marshal(M{"token": signedToken})
	c.JSON(http.StatusOK, gin.H{"userId": databaseInput.Id, "token": signedToken})
}

func RefreshToken(c *gin.Context) {
	var input models.TokenRefresh
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := &models.ClaimsJWT{}

	token, err := jwt.ParseWithClaims(input.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > models.LOGIN_EXPIRATION_DURATION {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	expirationTime := time.Now().Add(models.LOGIN_EXPIRATION_DURATION)
	claims.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(models.JWT_SIGNING_METHOD, claims)
	tokenString, err := token.SignedString(models.JWT_SIGNATURE_KEY)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
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
