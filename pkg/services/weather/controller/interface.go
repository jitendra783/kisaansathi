package controller

import (
	"context"
	"kisaanSathi/pkg/services/weather/models"

	"kisaanSathi/pkg/utils"
)

type WeatherController interface {
	CurrentWeather(ctx context.Context, lat, lng float64) (*models.CurrentWeatherResponse, error)
	Forecast(ctx context.Context, lat, lng float64) (*models.ForecastResponse, error)
	ForecastWeekly(ctx context.Context, lat, lng float64) (*models.WeeklyForecastResponse, error)
	HourlyForecast(ctx context.Context, lat, lng float64) (*models.HourlyResponse, error)
	Alerts(ctx context.Context, lat, lng float64) (*models.AlertResponse, error)
	AirQuality(ctx context.Context, lat, lng float64) (*models.AirQualityResponse, error)
	GetRainfall(ctx context.Context, lat, lng float64) (*models.RainfallResponse, error)
}

type weatherController struct {
	rest utils.RestCaller
}

func NewWeatherController(rest utils.RestCaller) WeatherController {
	return &weatherController{
		rest: rest,
	}
}
