package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	r := &maps.PlaceAutocompleteRequest{
		Input:    body.Input,
		Language: "id",
	}
	resp, err := client.PlaceAutocomplete(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":       false,
		"predictions": resp.Predictions})
}

type RoutesWithCarbonCalculated struct {
	Carbon float32 `json:"carbon"`
	maps.Route
}

type BodyRoutes struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Preference  string `json:"preference"`
}

// GET /routes
// GET routes based on origin and destination
func FindRoutes(c *gin.Context) {
	// Hashmap contains multiplier of grams of co2 emission per km for each vehicle
	multiplier := map[string]float64{
		"WALKING":   1.2,
		"BICYCLING": 21,
		"DRIVING":   192,
		"TRANSIT":   9.2,
	}
	fmt.Println("GET /routes")
	body := BodyRoutes{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	r := &maps.DirectionsRequest{
		Origin:       body.Origin,
		Destination:  body.Destination,
		Alternatives: true,
		Mode:         maps.Mode(body.Preference),
		Language:     "id",
	}
	routes, geos, err := client.Directions(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	temp := make([]RoutesWithCarbonCalculated, len(routes))
	var sum float64 = 0
	for i := range temp {
		temp[i].Route = routes[i]
		temp[i].Carbon = 99.48
		for _, item := range temp[i].Legs[0].Steps {
			sum += multiplier[item.TravelMode] * float64(item.Distance.Meters)
		}
		sum = 0
	}
	c.JSON(http.StatusOK, gin.H{
		"error":             false,
		"geocode_waypoints": geos,
		"routes":            temp,
	})
}

// POST /journey
// add user journey
func AddJourney(c *gin.Context) {
	var journey models.Journey
	// bind body
	if err := c.ShouldBindJSON(&journey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	// Check if user exist
	var user models.User
	if err := models.Db.Where("user_id = ?", journey.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User not found!"})
		return
	}
	// in the future maybe add checker to prevent double data
	// add journey
	models.Db.Create(&journey)
	c.JSON(http.StatusOK, journey)
}

// GET /journeys
// get All user journey
func FindAllJourneys(c *gin.Context) {
	var journeys []models.Journey
	models.Db.Find(&journeys)
	c.JSON(http.StatusOK, gin.H{
		"error":    true,
		"journeys": journeys})
}

// GET /journey/:journeyId
// get All user journey
func FindJourneyById(c *gin.Context) {
	var journey models.Journey
	if err := models.Db.Where("journey_id = ?", c.Param("journeyId")).First(&journey).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Journey not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   true,
		"journey": journey,
	})
}
