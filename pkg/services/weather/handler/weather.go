package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *weatherHandler) GetCurrentWeather(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)

	if lat == 0.0 || lng == 0.0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "latitude and longitude are required",
		})
		return
	}

	data, err := h.controller.CurrentWeather(ctx, lat, lng)
	if err != nil {
		fmt.Printf("Forecast Error: %+v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Current weather fetched successfully",
		"data":    data,
	})
}

func (h *weatherHandler) GetForecast(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid latitude",
		})
		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid longitude",
		})
		return
	}

	data, err := h.controller.Forecast(ctx, lat, lng)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
func (c *weatherHandler) GetForecastWeekly(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid latitude",
		})
		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid longitude",
		})
		return
	}

	data, err := c.controller.ForecastWeekly(ctx.Request.Context(), lat, lng)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
func (h *weatherHandler) GetHourlyForecast(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid latitude",
		})
		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid longitude",
		})
		return
	}

	data, err := h.controller.HourlyForecast(ctx.Request.Context(), lat, lng)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *weatherHandler) GetWeatherAlerts(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid latitude",
		})
		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid longitude",
		})
		return
	}

	data, err := h.controller.CurrentWeather(ctx.Request.Context(), lat, lng)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *weatherHandler) GetAirQuality(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid latitude",
		})
		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid longitude",
		})
		return
	}

	data, err := h.controller.AirQuality(ctx.Request.Context(), lat, lng)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
func (c *weatherHandler) GetRainfall(ctx *gin.Context) {

	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid latitude",
		})
		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid longitude",
		})
		return
	}

	data, err := c.controller.GetRainfall(ctx.Request.Context(), lat, lng)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
