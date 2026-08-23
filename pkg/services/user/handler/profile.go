package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/user/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUserDetails returns logged-in user's profile.
// If JWT middleware stores the email/user ID in context,
// prefer using c.GetString("email") instead of query params.
func (h *handler) GetUserDetails(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	email := c.Query("email")
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

	user, err := h.controller.GetUserDetails(c, email)
	if err != nil {
		logger.Log(c).Error("Failed to get user details", zap.Error(err))

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		c.Abort()
		return
	}

	if user == nil {
		c.JSON(
			http.StatusNotFound,
			network.FailureResponse(
				network.ApiErrors.NoDataFound.WithErrorDescription("user not found"),
			),
		)
		c.Abort()
		return
	}

	success(c, user)
}

// UpdateUserDetails updates the user's profile.
func (h *handler) UpdateUserDetails(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	response, err := h.controller.UpdateUserDetails(c, &request)
	if err != nil {
		logger.Log(c).Error("Failed to update user details", zap.Error(err))

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
