package models

// ===========================================
// Shared Models
// ===========================================

type Location struct {
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ===========================================
// Current Weather
// ===========================================

type CurrentWeatherResponse struct {
	Location Location `json:"location"`

	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feels_like"`

	Humidity   int     `json:"humidity"`
	Pressure   int     `json:"pressure"`
	WindSpeed  float64 `json:"wind_speed"`
	Visibility int     `json:"visibility"`

	Clouds  int     `json:"clouds"`
	UVIndex float64 `json:"uv_index,omitempty"`

	Condition   string `json:"condition"`
	Description string `json:"description"`
	Icon        string `json:"icon"`

	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}

// ===========================================
// Forecast
// ===========================================

type ForecastDay struct {
	Date      string  `json:"date"`
	Day       string  `json:"day"` // Monday, Tuesday...
	MinTemp   float64 `json:"min_temp"`
	MaxTemp   float64 `json:"max_temp"`
	Rain      float64 `json:"rain_probability"`
	Condition string  `json:"condition"`
	Icon      string  `json:"icon"`
}

type ForecastResponse struct {
	Location Location      `json:"location"`
	Days     []ForecastDay `json:"days"`
}

type WeeklyForecastResponse struct {
	Location Location      `json:"location"`
	Days     []ForecastDay `json:"days"`
}

// ===========================================
// Hourly Forecast
// ===========================================

type HourlyWeather struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature"`
	Rain        float64 `json:"rain_probability"`
	WindSpeed   float64 `json:"wind_speed"`
	Condition   string  `json:"condition"`
	Icon        string  `json:"icon"`
}

type HourlyResponse struct {
	Location Location        `json:"location"`
	Hours    []HourlyWeather `json:"hours"`
}

// ===========================================
// Weather Alerts
// ===========================================

type WeatherAlert struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	Severity string `json:"severity"`

	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`

	Source string `json:"source"`
}

type AlertResponse struct {
	Location Location       `json:"location"`
	Alerts   []WeatherAlert `json:"alerts"`
}

// ===========================================
// Air Quality
// ===========================================

type AirQualityResponse struct {
	Location Location `json:"location"`

	AQI     int    `json:"aqi"`
	Quality string `json:"quality"`

	CO   float64 `json:"co"`
	NO   float64 `json:"no"`
	NO2  float64 `json:"no2"`
	O3   float64 `json:"o3"`
	SO2  float64 `json:"so2"`
	PM25 float64 `json:"pm2_5"`
	PM10 float64 `json:"pm10"`
	NH3  float64 `json:"nh3"`
}

// ===========================================
// Rainfall
// ===========================================

type RainfallForecast struct {
	Time string `json:"time"`

	RainProbability float64 `json:"rain_probability"`

	ExpectedRainMM float64 `json:"expected_rain_mm"`
}

type RainfallResponse struct {
	Location Location           `json:"location"`
	Forecast []RainfallForecast `json:"forecast"`
}

// ===========================================
// Agriculture Advisory
// ===========================================

type CropAdvisory struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

type AdvisoryResponse struct {
	Location   Location       `json:"location"`
	Advisories []CropAdvisory `json:"advisories"`
}
