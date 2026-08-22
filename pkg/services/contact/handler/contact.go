package handler

import (
	"kisaanSathi/pkg/services/contact/models"

	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
)

func (h *contactHandler) ContactUs(ctx *gin.Context) {
	var req models.ContactRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	result, err := h.controller.ContactUs(&req)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}

// GetContactInfo
func (h *contactHandler) GetContactInfo(ctx *gin.Context) {

	result, err := h.controller.GetContactInfo()
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(result)
}
