package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// token := c.Request.

		fmt.Println("PASS THROUGH MIDDLEWARE !")

		c.Next()
	}
}
