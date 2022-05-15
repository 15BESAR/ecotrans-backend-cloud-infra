package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /autocomplete
// update user data with userid
func AutocompleteLocation(c *gin.Context) {
	fmt.Println("GET /autocomplete")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	var someString string = `{
		"predictions" : [
			{"text" : "Jalan sudirman"},
			{"text" : "Jalan sutomo"},
			{"text" : "Jalan sukajadi"}
		]
	}`
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(someString))
}

// GET /routes
// GET routes based on origin and destination
func FindRoutes(c *gin.Context) {
	fmt.Println("GET /routes")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	c.JSON(http.StatusOK, DummyGmapsData)
}

// POST /finish
// add user journey
func AddJourney(c *gin.Context) {
	fmt.Println("POST /finish")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	c.JSON(http.StatusOK, gin.H{"msg": "data saved successfully"})
}
