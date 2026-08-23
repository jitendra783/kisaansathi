package handler

import (
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
)

func (h *homeHandler) GetHome(c *gin.Context) {

	resp, err := h.controller.GetHome(c)
	if err != nil {

		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}
func (h *homeHandler) GetDashboardData(c *gin.Context) {

	resp, err := h.controller.GetDashboard(c)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}
