package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"kisaanSathi/pkg/config"
)

//
// OpenWeather Current Weather Response
//

type openWeatherAlertResponse struct {
	Alerts []struct {
		Event       string `json:"event"`
		Description string `json:"description"`
		Start       int64  `json:"start"`
		End         int64  `json:"end"`
		SenderName  string `json:"sender_name"`
	} `json:"alerts"`
}

type openWeatherCurrentResponse struct {
	Name string `json:"name"`

	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`

	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Visibility int `json:"visibility"`

	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
}

//
// OpenWeather Forecast Response
//

type openWeatherForecastResponse struct {
	City struct {
		Name string `json:"name"`

		Country string `json:"country"`

		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
	} `json:"city"`

	List []struct {
		DtTxt string `json:"dt_txt"`

		Main struct {
			Temp     float64 `json:"temp"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
			Humidity int     `json:"humidity"`
			Pressure int     `json:"pressure"`
		} `json:"main"`

		Pop float64 `json:"pop"`

		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`

		Rain struct {
			ThreeHour float64 `json:"3h"`
		} `json:"rain"`

		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"list"`
}

//
// OpenWeather Air Quality Response
//

type openWeatherAirQualityResponse struct {
	List []struct {
		Main struct {
			AQI int `json:"aqi"`
		} `json:"main"`

		Components struct {
			CO   float64 `json:"co"`
			NO   float64 `json:"no"`
			NO2  float64 `json:"no2"`
			O3   float64 `json:"o3"`
			SO2  float64 `json:"so2"`
			PM25 float64 `json:"pm2_5"`
			PM10 float64 `json:"pm10"`
			NH3  float64 `json:"nh3"`
		} `json:"components"`
	} `json:"list"`
}

//
// Generic OpenWeather Caller
//

func (c *weatherController) callOpenWeather(
	ctx context.Context,
	endpoint string,
	lat,
	lng float64,
	response interface{},
) error {

	baseURL := config.GetConfig().GetString("weather.base_url")

	apiKey := config.GetConfig().GetString("weather.api_key")
	fmt.Println(config.GetConfig().GetString("weather.api_key"))
	timeout := config.GetConfig().GetInt64("weather.timeout")

	query := map[string]string{
		"lat":   strconv.FormatFloat(lat, 'f', 6, 64),
		"lon":   strconv.FormatFloat(lng, 'f', 6, 64),
		"appid": apiKey,
		"units": "metric",
	}

	body, status, err := c.rest.InvokeHttp(
		ctx,
		http.MethodGet,
		baseURL+endpoint,
		nil,
		nil,
		timeout,
		nil,
		query,
		nil,
	)

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return fmt.Errorf("weather api returned status %d", status)
	}

	return json.Unmarshal(body, response)
}

//
// Current Weather
//

func (c *weatherController) fetchCurrentWeather(
	ctx context.Context,
	lat,
	lng float64,
) (*openWeatherCurrentResponse, error) {

	var response openWeatherCurrentResponse

	err := c.callOpenWeather(
		ctx,
		"/data/2.5/weather",
		lat,
		lng,
		&response,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

//
// Forecast
//

func (c *weatherController) fetchForecast(
	ctx context.Context,
	lat,
	lng float64,
) (*openWeatherForecastResponse, error) {

	var response openWeatherForecastResponse

	err := c.callOpenWeather(
		ctx,
		"/data/2.5/forecast",
		lat,
		lng,
		&response,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

//
// Air Quality
//

func (c *weatherController) fetchAirQuality(
	ctx context.Context,
	lat,
	lng float64,
) (*openWeatherAirQualityResponse, error) {

	var response openWeatherAirQualityResponse

	err := c.callOpenWeather(
		ctx,
		"/data/2.5/air_pollution",
		lat,
		lng,
		&response,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *weatherController) fetchAlerts(
	ctx context.Context,
	lat,
	lng float64,
) (*openWeatherAlertResponse, error) {

	var response openWeatherAlertResponse

	err := c.callOpenWeather(
		ctx,
		"/data/3.0/onecall?exclude=minutely,hourly,daily",
		lat,
		lng,
		&response,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
