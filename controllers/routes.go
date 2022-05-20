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

type FinalResponse struct {
	routes            []RoutesWithCarbonCalculated `json:"routes"`
	geocode_waypoints []maps.GeocodedWaypoint      `json:"geocode_waypoints"`
}

type RoutesWithCarbonCalculated struct {
	carbon float32    `json:"carbon"`
	route  maps.Route `json:"route"`
}

// GET /routes
// GET routes based on origin and destination
func FindRoutes(c *gin.Context) {
	fmt.Println("GET /routes")
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r := &maps.DirectionsRequest{
		Origin:       "Summarecon Mal Bekasi",
		Destination:  "Grand Indonesia",
		Alternatives: true,
		Mode:         maps.TravelModeTransit,
		Language:     "id",
	}
	routes, geos, err := client.Directions(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	temp := make([]RoutesWithCarbonCalculated, len(routes))
	for i := range temp {
		// routesWithCarbonCalculated := RoutesWithCarbonCalculated{route: routes[i], carbon: 99.48}
		// pretty.Println(routesWithCarbonCalculated)
		temp[i].route = routes[i]
		temp[i].carbon = 99.48
		// pretty.Println(temp[i])
		// panic("test")

	}
	// respfinal := FinalResponse{routes: temp, geocode_waypoints: geos}
	// pretty.Println(temp)
	c.JSON(http.StatusOK, gin.H{
		"geocode_waypoints": geos,
		"routes":            routes,
	})
}

// POST /finish
// add user journey
func AddJourney(c *gin.Context) {
	fmt.Println("POST /finish")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	c.JSON(http.StatusOK, gin.H{"msg": "data saved successfully"})
}
