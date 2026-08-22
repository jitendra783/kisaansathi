package handler

import (
	"strconv"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
)

func (h *videoHandler) GetVideos(ctx *gin.Context) {

	resp, err := h.controller.GetVideos()
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}

func (h *videoHandler) GetVideo(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	resp, err := h.controller.GetVideo(id)
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}

func (h *videoHandler) GetVideosByCategory(ctx *gin.Context) {

	category := ctx.Param("category")

	resp, err := h.controller.GetVideosByCategory(category)
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}