package handler

import (
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/user/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (f *handler) Register(c *gin.Context) {
	logger.Log(c).Debug("SERVICE-START")
	defer logger.Log(c).Debug("SERVICE-END")

	var request models.RegisterRequest

	if err := c.BindJSON(&request); err != nil {
		logger.Log(c).Error("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, network.BadRequestResponse(err, request))
		c.Abort()
		return
	}

	data, err := f.controller.Register(c, &request)

	logger.Log(c).Debug("data", zap.Any("data", data))

	if err != nil {
		logger.Log(c).Error("Something went wrong", zap.String("error", err.Error()))
		c.JSON(http.StatusNotFound, network.FailureResponse(network.ApiErrors.NoDataFound.WithErrorDescription(err.Error())))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, network.SuccessResponse(data))
}
func (f *handler) Login(c *gin.Context) {
	logger.Log(c).Debug("SERVICE-START")
	defer logger.Log(c).Debug("SERVICE-END")

	var request models.LoginRequest

	if err := c.BindJSON(&request); err != nil {
		logger.Log(c).Error("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, network.BadRequestResponse(err, request))
		c.Abort()
		return
	}

	data, err := f.controller.Login(c, &request)

	logger.Log(c).Debug("data", zap.Any("data", data))

	if err != nil {
		logger.Log(c).Error("Something went wrong", zap.String("error", err.Error()))
		c.JSON(http.StatusNotFound, network.FailureResponse(network.ApiErrors.NoDataFound.WithErrorDescription(err.Error())))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, network.SuccessResponse(data))
}

func (f *handler) Logout(c *gin.Context) {
	logger.Log(c).Debug("SERVICE-START")
	defer logger.Log(c).Debug("SERVICE-END")

	var (
		request models.LogoutRequest
	)

	if err := c.BindJSON(&request); err != nil {
		logger.Log(c).Error("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, network.BadRequestResponse(err, request))
		c.Abort()
		return
	}

	data, err := f.controller.Logout(c, &request)

	logger.Log(c).Debug("data", zap.Any("data", data))

	if err != nil {
		logger.Log(c).Error("Something went wrong", zap.String("error", err.Error()))
		c.JSON(http.StatusNotFound, network.FailureResponse(network.ApiErrors.NoDataFound.WithErrorDescription(err.Error())))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, network.SuccessResponse(data))
}
