package controllers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Forecast struct {
	Temp float32 `json:"temp"`
	Uv   float32 `json:"uv"`
	Aqi  float32 `json:"aqi"`
}

// GET /forecast?destination=""&arrivedTime=""
// update user data with userid
func FindForecast(c *gin.Context) {
	// var journeys []models.Journey
	// models.Db.Find(&journeys)
	temp := rand.Float32()*8 + 25
	aqi := rand.Float32()*15 + 50
	uv := rand.Float32()*2 + 6
	forecast := Forecast{temp, uv, aqi}
	c.JSON(http.StatusOK, gin.H{
		"error":    false,
		"forecast": forecast})
}
