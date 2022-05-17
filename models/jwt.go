package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME = "Lesgo"
var LOGIN_EXPIRATION_DURATION = 5 * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("15 besar")

type ClaimsJWT struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}
