package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func serviceStart(c *gin.Context) {
	logger.Log(c).Debug("SERVICE-START")
}

func serviceEnd(c *gin.Context) {
	logger.Log(c).Debug("SERVICE-END")
}

func badRequest(c *gin.Context, err error, req interface{}) {
	logger.Log(c).Error("Invalid request payload", zap.Error(err))
	c.JSON(http.StatusBadRequest, network.BadRequestResponse(err, req))
	c.Abort()
}

func failure(c *gin.Context, err error) {
	logger.Log(c).Error("Request failed", zap.Error(err))
	c.JSON(
		http.StatusInternalServerError,
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		),
	)
	c.Abort()
}

func noData(c *gin.Context, err error) {
	logger.Log(c).Warn("No data found")
	c.JSON(
		http.StatusNotFound,
		network.FailureResponse(
			network.ApiErrors.NoDataFound.WithErrorDescription(err.Error()),
		),
	)
	c.Abort()
}

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, network.SuccessResponse(data))
}
