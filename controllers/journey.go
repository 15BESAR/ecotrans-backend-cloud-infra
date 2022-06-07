package controllers

import (
	"context"
	"fmt"
	"math"
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
	fmt.Println("POST /autocomplete")
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
	Carbon float64 `json:"carbon"`
	Reward int     `json:"reward"`
	maps.Route
}

type BodyRoutes struct {
	UserId      string `json:"userId"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Preference  string `json:"preference"`
}

// POST /routes
// GET routes based on origin and destination
func FindRoutes(c *gin.Context) {
	// Hashmap contains multiplier of grams of co2 emission per km for each vehicle
	// Already adjusted with Co2 emission by mode of transport data [2018]
	multiplier := map[string]float64{
		"walking":   0,
		"bicycling": 0,
		"driving":   192,
		"transit":   41,
	}
	fmt.Println("POST /routes")
	body := BodyRoutes{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	var userDb models.User
	err := models.Db.Where("user_id = ?", body.UserId).First(&userDb).Error
	if err != nil {
		c.JSON(http.StatusMovedPermanently, gin.H{
			"error": true,
			"msg":   "No Id Found",
		})
		return
	}
	if _, ok := multiplier[userDb.Vehicle]; !ok {
		c.JSON(http.StatusMovedPermanently, gin.H{
			"error": true,
			"msg":   "Vehicle user wrong not Input",
		})
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
	carbonToRewardFactor := -0.1
	for i := range temp {
		temp[i].Route = routes[i]

		for _, item := range temp[i].Legs[0].Steps {
			sum += (multiplier[item.TravelMode] - multiplier[userDb.Vehicle]*float64(item.Distance.Meters))
		}
		temp[i].Carbon = sum
		temp[i].Reward = int(math.Max(0.0, sum*carbonToRewardFactor))
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
		"error":    false,
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
		"error":   false,
		"journey": journey,
	})
}
