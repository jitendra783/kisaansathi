package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/user/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *handler) Logout(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.LogoutRequest

	// Validate Request
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	// Call Controller
	response, err := h.controller.Logout(c, &request)
	if err != nil {
		logger.Log(c).Error("Logout failed", zap.Error(err))

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		c.Abort()
		return
	}

	// Success Response
	success(c, response)
}
