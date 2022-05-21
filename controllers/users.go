package controllers

import (
	"fmt"
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GET /users
// get all users data
func FindUsers(c *gin.Context) {
	fmt.Println("GET /users")
	var users []models.User
	models.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})

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
	// Bind body

	// check if data valid
	// Update to DB
	// Return User Information
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("test"))
}
