package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kisaanSathi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (f *forecastHandler) GetForecast(c *gin.Context) {

	city := c.Query("city")

	geo, err := GetLatandLonFromCity(c, city)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lat := fmt.Sprintf("%f", geo.Results[0].Latitude)
	lon := fmt.Sprintf("%f", geo.Results[0].Longitude)
	if lat == "" || lon == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "latitude and longitude are required"})
		return
	}

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&daily=temperature_2m_max,temperature_2m_min,precipitation_sum&timezone=auto", lat, lon)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch forecast"})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	c.Data(http.StatusOK, "application/json", body)
}

func GetLatandLonFromCity(c *gin.Context, city string) (*models.GeoResponse, error) {
	var geo models.GeoResponse
	if city == "" {
		return nil, errors.New("either city or lat and lon are required")
	}

	// If city is provided, fetch lat and lon from geocoding API
	// If lat and lon are provided, use them directly

	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s", city)
	resp, err := http.Get(geoURL)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("failed to fetch location from city")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &geo); err != nil || len(geo.Results) == 0 {
		return nil, errors.New("invalid city or no coordinates found")
	}
	return &geo, nil
}
