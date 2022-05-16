package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var secretPath []string = []string{"/user", "/users"}

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// token := c.Request.
		path := c.FullPath()

		if stringInSlice(path, secretPath) {
			auth := c.Request.Header.Get("Authorization")
			if auth == "" {
				c.String(http.StatusForbidden, "No header auth")
				c.Abort()
				return
			}
			tokenString := strings.TrimPrefix(auth, "Bearer ")
			if tokenString == auth {
				c.String(http.StatusForbidden, "No bearer token in auth")
				c.Abort()
				return
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("SIGNING METHOD INVALID")
				} else if method != models.JWT_SIGNING_METHOD {
					return nil, fmt.Errorf("SIGNING METHOD INVALID")
				}
				return models.JWT_SIGNATURE_KEY, nil
			})
			if err != nil {
				c.String(http.StatusBadRequest, "Parse vailed")
				c.Abort()
				return
			}

			_, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				c.String(http.StatusBadRequest, "not valid")
				c.Abort()
				return
			}
		}
		fmt.Println("PASS THROUGH MIDDLEWARE !")

		c.Next()
	}
}
