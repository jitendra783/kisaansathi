package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/user/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login godoc
//
// @Summary User Login
// @Description Authenticate user and return JWT token
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login Request"
// @Router /user/login [post]
func (h *handler) Login(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.LoginRequest

	// Validate Request
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	// Call Controller
	response, err := h.controller.Login(c, &request)
	if err != nil {
		logger.Log(c).Error("Login failed", zap.Error(err))

		c.JSON(
			http.StatusUnauthorized,
			network.FailureResponse(
				network.ApiErrors.InvalidCredentials.WithErrorDescription(err.Error()),
			),
		)
		c.Abort()
		return
	}

	// Success Response
	network.SuccessResponse(response)
}
