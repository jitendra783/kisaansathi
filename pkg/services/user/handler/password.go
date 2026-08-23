package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/user/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ChangePassword updates the user's password.
func (h *handler) ChangePassword(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.ChangePasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	response, err := h.controller.ChangePassword(c, &request)
	if err != nil {
		logger.Log(c).Error("Failed to change password", zap.Error(err))

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

// ForgotPassword sends a reset password link or OTP.
func (h *handler) ForgotPassword(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	response, err := h.controller.ForgotPassword(c, &request)
	if err != nil {
		logger.Log(c).Error("Forgot password failed", zap.Error(err))

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

// ResetPassword resets the password using a token or OTP.
func (h *handler) ResetPassword(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.ResetPasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	response, err := h.controller.ResetPassword(c, &request)
	if err != nil {
		logger.Log(c).Error("Reset password failed", zap.Error(err))

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
