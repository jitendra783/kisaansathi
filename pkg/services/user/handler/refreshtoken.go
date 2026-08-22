package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/user/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RefreshToken generates a new access token using a valid refresh token.
func (h *handler) RefreshToken(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.RefreshTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	response, err := h.controller.RefreshToken(c, &request)
	if err != nil {
		logger.Log(c).Error("Failed to refresh token", zap.Error(err))

		c.JSON(
			http.StatusUnauthorized,
			network.FailureResponse(
				network.ApiErrors.InvalidToken.WithErrorDescription(err.Error()),
			),
		)
		c.Abort()
		return
	}

	success(c, response)
}