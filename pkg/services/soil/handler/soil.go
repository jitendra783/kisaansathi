package handler

import (
	"kisaanSathi/pkg/logger"

	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *soilHandler) GetSoilTypes(ctx *gin.Context) {

	logger.Log().Info("GetSoilTypes API called")

	response, err := h.controller.GetSoilTypes(ctx)
	if err != nil {
		logger.Log().Error("failed to fetch soil types : %v", zap.Error(err))

		network.FailureResponse(

			network.ApiErrors.InternalServerError.WithErrorDescription("Unable to fetch soil types"),
		)
		return
	}

	network.SuccessResponse(
		response,
	)
}

func (h *soilHandler) GetSoilReport(ctx *gin.Context) {

	logger.Log().Info("GetSoilReport API called")

	response, err := h.controller.GetSoilReport(ctx)
	if err != nil {

		//logger.Log().Errorf("GetSoilReport failed : %v", err)

		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription("Unable to fetch soil report"),
		)
		return
	}

	network.SuccessResponse(
		response,
	)
}

func (h *soilHandler) CreateSoilTest(ctx *gin.Context) {

	logger.Log().Info("CreateSoilTest API called")

	response, err := h.controller.CreateSoilTest(ctx)
	if err != nil {

		//logger.Log().Errorf("CreateSoilTest failed : %v", err)

		network.FailureResponse(
			//// "Unable to create soil test",
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(
		response,
	)
}

func (h *soilHandler) GetSoilRecommendation(ctx *gin.Context) {

	logger.Log().Info("GetSoilRecommendation API called")

	response, err := h.controller.GetSoilRecommendation(ctx)
	if err != nil {

		//logger.Log().Error("GetSoilRecommendation failed : %v", err.Error())

		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(
		//"Soil recommendations fetched successfully",
		response,
	)
}
