package services

import (
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/feeds"
	"kisaanSathi/pkg/services/forecast"
	"kisaanSathi/pkg/services/mandi"
	session "kisaanSathi/pkg/services/session/handler"
	reg "kisaanSathi/pkg/services/user/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceObject struct {
	session.SessionGroup
	reg.RegisterHandler
	feeds.FeedsHandler
	forecast.ForecastHandler
	mandi.MandiHandler
}

type ServiceLayer interface {
	GetMFHealth(c *gin.Context)
	session.SessionGroup
	reg.RegisterHandler
	feeds.FeedsHandler
	forecast.ForecastHandler
	mandi.MandiHandler
}

func NewServiceObject(repo repo.DataObject) ServiceLayer {
	return &serviceObject{
		session.NewSessionGroup(repo),
		reg.NewRegisterHandler(reg.RegisterController(repo)),
		feeds.NewFeedsHandler(repo),
		forecast.NewForecastHandler(repo),
		mandi.NewMandiHandler(repo),
	}
}

func (s *serviceObject) GetMFHealth(c *gin.Context) {
	c.JSON(http.StatusOK, network.SuccessResponse("I AM HEALTHY"))
}
