package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

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
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    models.APPLICATION_NAME,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(models.LOGIN_EXPIRATION_DURATION)),
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

	if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > models.LOGIN_EXPIRATION_DURATION {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	expirationTime := time.Now().Add(models.LOGIN_EXPIRATION_DURATION)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token = jwt.NewWithClaims(models.JWT_SIGNING_METHOD, claims)
	tokenString, err := token.SignedString(models.JWT_SIGNATURE_KEY)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
