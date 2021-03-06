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

// POST /partner/login
// User login, return token

func LoginPartner(c *gin.Context) {
	var userInput models.PartnerLogin
	var databaseInput models.Partner
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	err := models.Db.Where("username = ?", userInput.Username).First(&databaseInput).Error
	if err != nil {
		c.JSON(http.StatusMovedPermanently, gin.H{
			"error": true,
			"msg":   "No username"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databaseInput.Password), []byte(userInput.Password))
	if err != nil {
		c.JSON(http.StatusMovedPermanently, gin.H{
			"error": true,
			"msg":   "Wrong Password"})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "bad request"})
		return
	}
	json.Marshal(M{"token": signedToken})
	c.JSON(http.StatusOK, gin.H{
		"error":     false,
		"msg":       "Succsess login",
		"partnerId": databaseInput.PartnerID, "token": signedToken})
}

// GET /partner/login/refresh
// give old token, return refreshed new token
func RefreshTokenPartner(c *gin.Context) {
	var input models.TokenRefresh
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}

	claims := &models.ClaimsJWT{}

	token, err := jwt.ParseWithClaims(input.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": true,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
		return
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	if time.Until(time.Unix(claims.ExpiresAt.Unix(), 0)) > models.LOGIN_EXPIRATION_DURATION {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "token expired"})
		return
	}

	expirationTime := time.Now().Add(models.LOGIN_EXPIRATION_DURATION)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token = jwt.NewWithClaims(models.JWT_SIGNING_METHOD, claims)
	tokenString, err := token.SignedString(models.JWT_SIGNATURE_KEY)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"token": tokenString})
}
