package handler

import (
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
)

func (h *medatahandler) GetMetadata(c *gin.Context) {
	logger.Log(c).Debug("SERVICE-START")
	defer logger.Log(c).Debug("SERVICE-END")
	response, err := h.controller.GetMetadata()
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}
	network.SuccessResponse(response)
}
