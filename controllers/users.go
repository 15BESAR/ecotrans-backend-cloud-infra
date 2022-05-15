package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /register
// Register user
func RegisterUser(c *gin.Context) {
	fmt.Println("POST /Register")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))

	c.JSON(http.StatusOK, gin.H{"userId": "12d1d12d12", "token": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"})
}

// POST /login
// User login, return token
func LoginUser(c *gin.Context) {
	fmt.Println("POST /login")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))

	c.JSON(http.StatusOK, gin.H{"userId": "12d1d12d12", "token": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"})
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
