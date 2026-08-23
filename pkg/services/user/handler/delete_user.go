package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DeleteUser deletes a user account.
func (h *handler) DeleteUser(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	// Prefer getting the user ID/email from JWT middleware.
	// Example:
	// userID := c.GetInt64("user_id")
	// email := c.GetString("email")

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

	response, err := h.controller.DeleteUser(c, email)
	if err != nil {
		logger.Log(c).Error("Failed to delete user", zap.Error(err))

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
