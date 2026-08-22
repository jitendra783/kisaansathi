package handler

import (
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/feedback/models"

	"github.com/gin-gonic/gin"
)

func (h *feedbackHandler) CreateFeedback(ctx *gin.Context) {
	var req models.FeedbackRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	result, err := h.controller.CreateFeedback(&req)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}

func (h *feedbackHandler) GetFeedbacks(ctx *gin.Context) {

	result, err := h.controller.GetFeedbacks()
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}
