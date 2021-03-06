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
		if strings.Contains(a, b) {
			return true
		}
	}
	return false
}

var secretPath []string = []string{"/users", "/user", "/vouchers", "/voucher", "/journeys", "/journey", "/partners", "/partner", "/purchases", "/purchase"}

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Add CORS Header
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
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
				c.String(http.StatusBadRequest, "Signing method not valid")
				c.Abort()
				return
			}

			_, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				c.String(http.StatusBadRequest, "Signature not valid")
				c.Abort()
				return
			}
		}
		fmt.Println("PASS THROUGH MIDDLEWARE !")

		c.Next()
	}
}
