package handler

import (
	"kisaanSathi/pkg/services/crop/models"
	"strconv"

	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
)

func (h *cropHandler) GetCrops(ctx *gin.Context) {

	result, err := h.controller.GetCrops()
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}

func (h *cropHandler) GetCrop(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	result, err := h.controller.GetCrop(id)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}
func (h *cropHandler) CreateCrop(ctx *gin.Context) {

	var req models.CreateCropRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	if err := h.controller.CreateCrop(req); err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse("Crop created successfully")
}
func (h *cropHandler) UpdateCrop(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	var req models.UpdateCropRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	if err := h.controller.UpdateCrop(id, req); err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse("Crop updated successfully")
}
func (h *cropHandler) DeleteCrop(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	if err := h.controller.DeleteCrop(id); err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse("Crop deleted successfully")
}
func (h *cropHandler) GetCropSeason(ctx *gin.Context) {

	season := ctx.Query("season")

	result, err := h.controller.GetCropSeason(season)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}

func (h *cropHandler) GetRecommendedCrops(ctx *gin.Context) {

	soil := ctx.Query("soil")
	district := ctx.Query("district")

	result, err := h.controller.GetRecommendedCrops(soil, district)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}
