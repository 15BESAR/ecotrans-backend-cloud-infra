package models

import "github.com/golang-jwt/jwt/v4"

type ClaimsJWT struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
