package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
)

type Body struct {
	Input string `json:"input"`
}

type Locations struct {
	description string `json:"description"`
}

type ResponseAutoComplete struct {
	predictions []Locations `json:"predictions"`
}

// GET /autocomplete
// update user data with userid
func AutocompleteLocation(c *gin.Context) {
	fmt.Println("GET /autocomplete")
	body := Body{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r := &maps.PlaceAutocompleteRequest{
		Input:    body.Input,
		Language: "id",
	}
	resp, err := client.PlaceAutocomplete(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"predictions": resp.Predictions})
}

type RoutesWithCarbonCalculated struct{
	maps.Route,
	carbon float32
}
// GET /routes
// GET routes based on origin and destination
func FindRoutes(c *gin.Context) {
	fmt.Println("GET /routes")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	b, err := ioutil.ReadFile("./controllers/dummyJSON.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", b)
}

// POST /finish
// add user journey
func AddJourney(c *gin.Context) {
	fmt.Println("POST /finish")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	c.JSON(http.StatusOK, gin.H{"msg": "data saved successfully"})
}
