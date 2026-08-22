package controller

import (
	"context"
	"kisaanSathi/pkg/services/weather/models"
	"time"
)

func (c *weatherController) CurrentWeather(
	ctx context.Context,
	lat,
	lng float64,
) (*models.CurrentWeatherResponse, error) {

	weather, err := c.fetchCurrentWeather(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	response := &models.CurrentWeatherResponse{
		Location: models.Location{
			City:      weather.Name,
			Country:   weather.Sys.Country,
			Latitude:  weather.Coord.Lat,
			Longitude: weather.Coord.Lon,
		},

		Temperature: weather.Main.Temp,
		FeelsLike:   weather.Main.FeelsLike,

		Humidity:   weather.Main.Humidity,
		Pressure:   weather.Main.Pressure,
		WindSpeed:  weather.Wind.Speed,
		Visibility: weather.Visibility,

		Clouds: weather.Clouds.All,

		Condition:   weather.Weather[0].Main,
		Description: weather.Weather[0].Description,
		Icon:        weather.Weather[0].Icon,

		Sunrise: weather.Sys.Sunrise,
		Sunset:  weather.Sys.Sunset,
	}

	return response, nil
}

func (c *weatherController) Forecast(
	ctx context.Context,
	lat,
	lng float64,
) (*models.ForecastResponse, error) {

	forecast, err := c.fetchForecast(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	response := &models.ForecastResponse{
		Location: models.Location{
			City:      forecast.City.Name,
			Country:   forecast.City.Country,
			Latitude:  forecast.City.Coord.Lat,
			Longitude: forecast.City.Coord.Lon,
		},
		Days: make([]models.ForecastDay, 0),
	}

	// OpenWeather returns 3-hour intervals.
	// We keep one forecast entry per day.

	processedDays := make(map[string]bool)

	for _, item := range forecast.List {

		t, err := time.Parse(
			"2006-01-02 15:04:05",
			item.DtTxt,
		)
		if err != nil {
			continue
		}

		date := t.Format("2006-01-02")

		if processedDays[date] {
			continue
		}

		processedDays[date] = true

		response.Days = append(
			response.Days,
			models.ForecastDay{
				Date: date,
				Day:  t.Weekday().String(),

				MinTemp: item.Main.TempMin,
				MaxTemp: item.Main.TempMax,

				Rain: item.Pop * 100,

				Condition: item.Weather[0].Main,

				Icon: item.Weather[0].Icon,
			},
		)

		// OpenWeather provides a maximum of 5 forecast days.
		if len(response.Days) == 5 {
			break
		}
	}

	return response, nil
}

func (c *weatherController) HourlyForecast(
	ctx context.Context,
	lat,
	lng float64,
) (*models.HourlyResponse, error) {

	forecast, err := c.fetchForecast(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	response := &models.HourlyResponse{
		Location: models.Location{
			City:      forecast.City.Name,
			Country:   forecast.City.Country,
			Latitude:  forecast.City.Coord.Lat,
			Longitude: forecast.City.Coord.Lon,
		},
		Hours: make([]models.HourlyWeather, 0),
	}

	// Return next 24 hours (8 entries × 3 hours)
	limit := 8
	if len(forecast.List) < limit {
		limit = len(forecast.List)
	}

	for i := 0; i < limit; i++ {

		item := forecast.List[i]

		t, err := time.Parse(
			"2006-01-02 15:04:05",
			item.DtTxt,
		)
		if err != nil {
			continue
		}

		hour := models.HourlyWeather{
			Time:        t.Format("03:04 PM"),
			Temperature: item.Main.Temp,
			Rain:        item.Pop * 100,
			WindSpeed:   item.Wind.Speed,
			Condition:   item.Weather[0].Main,
			Icon:        item.Weather[0].Icon,
		}

		response.Hours = append(response.Hours, hour)
	}

	return response, nil
}

func (c *weatherController) ForecastWeekly(
	ctx context.Context,
	lat,
	lng float64,
) (*models.WeeklyForecastResponse, error) {

	forecast, err := c.fetchForecast(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	response := &models.WeeklyForecastResponse{
		Location: models.Location{
			City:      forecast.City.Name,
			Country:   forecast.City.Country,
			Latitude:  forecast.City.Coord.Lat,
			Longitude: forecast.City.Coord.Lon,
		},
		Days: make([]models.ForecastDay, 0),
	}

	type daySummary struct {
		Date      string
		MinTemp   float64
		MaxTemp   float64
		Rain      float64
		Condition string
		Icon      string
	}

	dailyMap := make(map[string]*daySummary)
	order := make([]string, 0)

	for _, item := range forecast.List {

		t, err := time.Parse(
			"2006-01-02 15:04:05",
			item.DtTxt,
		)
		if err != nil {
			continue
		}

		date := t.Format("2006-01-02")

		if _, exists := dailyMap[date]; !exists {

			dailyMap[date] = &daySummary{
				Date:      date,
				MinTemp:   item.Main.TempMin,
				MaxTemp:   item.Main.TempMax,
				Rain:      item.Pop * 100,
				Condition: item.Weather[0].Main,
				Icon:      item.Weather[0].Icon,
			}

			order = append(order, date)
			continue
		}

		day := dailyMap[date]

		if item.Main.TempMin < day.MinTemp {
			day.MinTemp = item.Main.TempMin
		}

		if item.Main.TempMax > day.MaxTemp {
			day.MaxTemp = item.Main.TempMax
		}

		if item.Pop*100 > day.Rain {
			day.Rain = item.Pop * 100
			day.Condition = item.Weather[0].Main
			day.Icon = item.Weather[0].Icon
		}
	}

	for _, date := range order {

		day := dailyMap[date]

		response.Days = append(
			response.Days,
			models.ForecastDay{
				Date:      day.Date,
				MinTemp:   day.MinTemp,
				MaxTemp:   day.MaxTemp,
				Rain:      day.Rain,
				Condition: day.Condition,
				Icon:      day.Icon,
			},
		)
	}

	return response, nil
}

func (c *weatherController) Alerts(
	ctx context.Context,
	lat,
	lng float64,
) (*models.AlertResponse, error) {

	alerts, err := c.fetchAlerts(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	response := &models.AlertResponse{
		Alerts: make([]models.WeatherAlert, 0),
	}

	for _, alert := range alerts.Alerts {

		response.Alerts = append(
			response.Alerts,
			models.WeatherAlert{
				Title:       alert.Event,
				Description: alert.Description,
				Severity:    "Warning",
				StartTime:   alert.Start,
				EndTime:     alert.End,
				Source:      alert.SenderName,
			},
		)
	}

	return response, nil
}
func (c *weatherController) AirQuality(
	ctx context.Context,
	lat,
	lng float64,
) (*models.AirQualityResponse, error) {

	air, err := c.fetchAirQuality(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	response := &models.AirQualityResponse{}

	if len(air.List) == 0 {
		return response, nil
	}

	item := air.List[0]

	response.AQI = item.Main.AQI
	response.Quality = getAQILevel(item.Main.AQI)

	response.CO = item.Components.CO
	response.NO = item.Components.NO
	response.NO2 = item.Components.NO2
	response.O3 = item.Components.O3
	response.SO2 = item.Components.SO2
	response.PM25 = item.Components.PM25
	response.PM10 = item.Components.PM10
	response.NH3 = item.Components.NH3

	return response, nil
}
func getAQILevel(aqi int) string {
	switch aqi {
	case 1:
		return "Good"
	case 2:
		return "Fair"
	case 3:
		return "Moderate"
	case 4:
		return "Poor"
	case 5:
		return "Very Poor"
	default:
		return "Unknown"
	}
}

func (c *weatherController) GetRainfall(
	ctx context.Context,
	lat,
	lng float64,
) (*models.RainfallResponse, error) {

	forecast, err := c.fetchForecast(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	response := &models.RainfallResponse{
		Location: models.Location{
			City:      forecast.City.Name,
			Country:   forecast.City.Country,
			Latitude:  forecast.City.Coord.Lat,
			Longitude: forecast.City.Coord.Lon,
		},
		Forecast: make([]models.RainfallForecast, 0),
	}

	for _, item := range forecast.List {

		t, err := time.Parse(
			"2006-01-02 15:04:05",
			item.DtTxt,
		)
		if err != nil {
			continue
		}

		rainfall := 0.0
		if item.Rain.ThreeHour > 0 {
			rainfall = item.Rain.ThreeHour
		}

		response.Forecast = append(
			response.Forecast,
			models.RainfallForecast{
				Time: t.Format("02 Jan 03:04 PM"),

				RainProbability: item.Pop * 100,

				ExpectedRainMM: rainfall,
			},
		)
	}

	return response, nil
}
