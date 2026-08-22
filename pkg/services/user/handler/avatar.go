package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *handler) UpdateAvatar(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	email := c.PostForm("email")
	if email == "" {
		c.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("email is required"),
			),
		)
		c.Abort()
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		logger.Log(c).Error("Avatar file not found", zap.Error(err))

		c.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("avatar file is required"),
			),
		)
		c.Abort()
		return
	}

	response, err := h.controller.UpdateAvatar(c, email, file)
	if err != nil {
		logger.Log(c).Error("Failed to update avatar", zap.Error(err))

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		c.Abort()
		return
	}

	success(c, response)
}
