package forecast

import (
	"kisaanSathi/pkg/repo"

	"github.com/gin-gonic/gin"
)

type forecastHandler struct {
}
type ForecastHandler interface {
	GetForecast(c *gin.Context)
}

func NewForecastHandler(repo repo.DataObject) ForecastHandler {
	return &forecastHandler{}
}
