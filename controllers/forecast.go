package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	"github.com/patrickmn/go-cache"
	"googlemaps.github.io/maps"
)

type Forecast struct {
	Temp float32 `json:"temp"`
	Uv   float32 `json:"uv"`
	Aqi  float32 `json:"aqi"`
}

type ForecastAPIBody struct {
	Destination string `json:"destination"`
	ArrivedHour int    `json:"arrivedHour"`
}

type Forecasts struct {
	Temps []float32 `json:"temps"`
	UVs   []float32 `json:"uvs"`
	AQIs  []float32 `json:"aqis"`
}

// Struct for cache
// for each region, have temp,uv,aqi
// each is array of float32

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
	// check if hour 0<=x<=24
	if !(0 <= body.ArrivedHour && body.ArrivedHour <= 24) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Please input valid arrivedHour"})
		return
	}
	// Get full address from Gmaps API
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	pretty.Println(body.Destination)
	pretty.Println(body.ArrivedHour)
	r := &maps.FindPlaceFromTextRequest{
		Input:     body.Destination,
		InputType: "textquery",
		Fields:    []maps.PlaceSearchFieldMask{maps.PlaceSearchFieldMaskFormattedAddress},
	}
	resp, err := client.FindPlaceFromText(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error()})
		return
	}
	// if resp empty
	if len(resp.Candidates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Places not valid"})
		return
	}
	// Check region
	region := determineRegion(resp.Candidates[0].FormattedAddress)
	// Check cache, if not found, run Query to Fill data from bigquery
	_, found := models.C.Get(region)
	if !found {
		fetchFromBigQuery()
	}
	// Get data from cache
	regionForecasts, f := models.C.Get(region)
	if !f || reflect.TypeOf(regionForecasts).Kind() != reflect.Slice { // Still not found inside go-cache
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Problem with getting forecast data"})
		return
	}
	s := reflect.ValueOf(regionForecasts)
	forecast := s.Index(body.ArrivedHour - 1).Interface() // -1 since idx start at 0
	c.JSON(http.StatusOK, gin.H{
		"error":    false,
		"forecast": forecast})
}

func determineRegion(address string) string {
	stringLower := strings.ToLower(address)
	if strings.Contains(stringLower, "jakarta pusat") {
		return "pusat"
	} else if strings.Contains(stringLower, "jakarta barat") {
		return "barat"
	} else if strings.Contains(stringLower, "jakarta timur") {
		return "timur"
	} else if strings.Contains(stringLower, "jakarta Utara") {
		return "utara"
	} else {
		return "selatan"
	}
}

func fetchFromBigQuery() {
	fmt.Println("Get into fetchfrom bigquery")
	// yang sekarang isinya manual duls
	// Nanti ganti ke bigquery klo udah bisa
	models.C.Set("barat", createRandomForecast(), cache.DefaultExpiration)
	models.C.Set("timur", createRandomForecast(), cache.DefaultExpiration)
	models.C.Set("utara", createRandomForecast(), cache.DefaultExpiration)
	models.C.Set("selatan", createRandomForecast(), cache.DefaultExpiration)
	models.C.Set("pusat", createRandomForecast(), cache.DefaultExpiration)
}

func createRandomForecast() []Forecast {
	mySlice := make([]Forecast, 0, 24)
	var temp float32
	var uv float32
	var aqi float32
	for i := 0; i < 24; i++ {
		temp = rand.Float32()*8 + 25
		if (i+1) < 6 || (i+1) > 16 {
			uv = 0
		} else {
			uv = rand.Float32()*2 + 6
		}
		aqi = rand.Float32()*15 + 50
		mySlice = append(mySlice, Forecast{temp, uv, aqi})
	}
	return mySlice
}
