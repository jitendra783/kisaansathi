package handler

import (
	"kisaanSathi/pkg/services/weather/controller"
	"kisaanSathi/pkg/utils"

	"github.com/gin-gonic/gin"
)

type WeatherHandler interface {
	GetCurrentWeather(*gin.Context)
	GetForecast(*gin.Context)
	GetHourlyForecast(*gin.Context)
	GetWeatherAlerts(*gin.Context)
	GetAirQuality(*gin.Context)
	GetRainfall(*gin.Context)
	GetForecastWeekly(*gin.Context)
}

type weatherHandler struct {
	controller controller.WeatherController
	rest       utils.RestCaller
}

func NewWeatherHandler(c controller.WeatherController) WeatherHandler {
	return &weatherHandler{
		controller: c,
	}
}
func NewWeatherController() controller.WeatherController {
	return controller.NewWeatherController(utils.GetRestCaller())
}
