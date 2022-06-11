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
	Time int     `json:"time"`
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

var arrayCity = [5]string{"jakarta_barat", "jakarta_timur", "jakarta_pusat", "jakarta_utara", "jakarta_selatan"}

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
	// Test Fetch

	// Check region
	region := determineRegion(resp.Candidates[0].FormattedAddress)
	// Check cache, if not found, run Query to Fill data from bigquery
	_, found := models.C.Get(region)
	if !found {
		fetchFromBigQuery()
	}
	// Get data from cache
	regionForecasts, f := models.C.Get(region)
	if !f || reflect.TypeOf(regionForecasts).Kind() != reflect.Slice {
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
	} else if strings.Contains(stringLower, "jakarta utara") {
		return "utara"
	} else {
		return "selatan"
	}
}

func fetchFromBigQuery() {
	fmt.Println("Get into fetchfrom bigquery")
	// yang sekarang isinya manual duls
	// Nanti ganti ke bigquery klo udah bisa
	m := GetForecastTemperatureFromBiqQuery()
	pretty.Println(m)
	// Kalo data bigquery format dan isinya udah sesuai, tuker random() ama data bq
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
		mySlice = append(mySlice, Forecast{i + 1, temp, uv, aqi})
	}
	return mySlice
}

func GetForecastTemperatureFromBiqQuery() map[string][]Forecast {
	var queryString = "SELECT uv.forecast_hour,tem.forecasted_temperature, uv.forecasted_uv, pm25.forecasted_pm25 ,'jakarta_barat' as city FROM `bangkit-352613.forecast.temperature_jakarta_barat` tem  JOIN `bangkit-352613.forecast.uv_index_jakarta_barat` as uv ON uv.forecast_hour = tem.forecast_hour JOIN `bangkit-352613.forecast.pm25_index_jakarta_barat` as pm25 ON pm25.forecast_hour = tem.forecast_hour union all SELECT uv.forecast_hour,tem.forecasted_temperature, uv.forecasted_uv, pm25.forecasted_pm25 ,'jakarta_timur' as city  FROM `bangkit-352613.forecast.temperature_jakarta_timur` tem JOIN `bangkit-352613.forecast.uv_index_jakarta_timur` as uv ON uv.forecast_hour = tem.forecast_hour JOIN `bangkit-352613.forecast.pm25_index_jakarta_timur` as pm25 ON pm25.forecast_hour = tem.forecast_hour union all SELECT uv.forecast_hour,tem.forecasted_temperature, uv.forecasted_uv, pm25.forecasted_pm25 ,'jakarta_utara' as city  FROM `bangkit-352613.forecast.temperature_jakarta_utara` tem JOIN `bangkit-352613.forecast.uv_index_jakarta_utara` as uv ON uv.forecast_hour = tem.forecast_hour JOIN `bangkit-352613.forecast.pm25_index_jakarta_utara` as pm25 ON pm25.forecast_hour = tem.forecast_hour union all SELECT uv.forecast_hour,tem.forecasted_temperature, uv.forecasted_uv, pm25.forecasted_pm25 ,'jakarta_selatan' as city  FROM `bangkit-352613.forecast.temperature_jakarta_selatan` tem JOIN `bangkit-352613.forecast.uv_index_jakarta_selatan` as uv ON uv.forecast_hour = tem.forecast_hour JOIN `bangkit-352613.forecast.pm25_index_jakarta_selatan` as pm25 ON pm25.forecast_hour = tem.forecast_hour union all SELECT uv.forecast_hour,tem.forecasted_temperature, uv.forecasted_uv, pm25.forecasted_pm25 ,'jakarta_pusat' as city  FROM `bangkit-352613.forecast.temperature_jakarta_pusat` tem JOIN `bangkit-352613.forecast.uv_index_jakarta_pusat` as uv ON uv.forecast_hour = tem.forecast_hour JOIN `bangkit-352613.forecast.pm25_index_jakarta_pusat` as pm25 ON pm25.forecast_hour = tem.forecast_hour;"
	var resultQuery = models.GetQuery(queryString)
	res := map[string][]Forecast{}
	for i := 0; i < len(resultQuery); i++ {
		var temps = resultQuery[i]
		if _, ok := res[formatCityString(temps.City)]; !ok {
			res[formatCityString(temps.City)] = make([]Forecast, 0)
		}
		var newDataTemp Forecast
		newDataTemp.Uv = temps.Forecasted_uv
		newDataTemp.Aqi = temps.Forecasted_pm25
		newDataTemp.Temp = temps.Forecasted_temperature
		newDataTemp.Time = temps.Forecast_hour
		res[formatCityString(temps.City)] = append(res[formatCityString(temps.City)], newDataTemp)
	}
	return res
}

func formatCityString(str string) string {
	stringLower := strings.ToLower(str)
	if strings.Contains(stringLower, "jakarta_pusat") {
		return "pusat"
	} else if strings.Contains(stringLower, "jakarta_barat") {
		return "barat"
	} else if strings.Contains(stringLower, "jakarta_timur") {
		return "timur"
	} else if strings.Contains(stringLower, "jakarta_utara") {
		return "utara"
	} else {
		return "selatan"
	}
}
