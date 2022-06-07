package controllers

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Forecast struct {
	Temp float32 `json:"temp"`
	Uv   float32 `json:"uv"`
	Aqi  float32 `json:"aqi"`
}

type ForecastAPIBody struct {
	destination string `json:"destination"`
	arrivedHour int    `json:"arriverHour"`
}

type Region int8

// Enums for Region
const (
	Pusat Region = iota
	Barat
	Timur
	Utara
	Selatan
)

// POST /forecast?destination=""&arrivedHour=""
// update user data with userid
func FindForecast(c *gin.Context) {
	// bind body
	body := ForecastAPIBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	// Get full address from Gmaps API

	// Check region
	// region := determineRegion("test123")
	// Check cache, if not found, run Query to Fill data from bigquery

	// Get data from cache

	// If somehow the value is still nil, randomize it
	temp := rand.Float32()*8 + 25
	aqi := rand.Float32()*15 + 50
	uv := rand.Float32()*2 + 6
	forecast := Forecast{temp, uv, aqi}
	c.JSON(http.StatusOK, gin.H{
		"error":    false,
		"forecast": forecast})
}

func determineRegion(address string) Region {
	stringLower := strings.ToLower(address)
	if strings.Contains(stringLower, "jakarta pusat") {
		return Pusat
	} else if strings.Contains(stringLower, "jakarta barat") {
		return Barat
	} else if strings.Contains(stringLower, "jakarta timur") {
		return Timur
	} else if strings.Contains(stringLower, "jakarta Utara") {
		return Utara
	} else {
		return Selatan
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.Contains(a, b) {
			return true
		}
	}
	return false
}
