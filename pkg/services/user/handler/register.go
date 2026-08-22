package handler

import (
	"net/http"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/user/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *handler) Register(c *gin.Context) {
	serviceStart(c)
	defer serviceEnd(c)

	var request models.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err, request)
		return
	}

	if err := h.controller.Register(c, &request); err != nil {
		logger.Log(c).Error("Registration failed", zap.Error(err))

		c.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.NoDataFound.WithErrorDescription(err.Error()),
			),
		)
		c.Abort()
		return
	}

	success(c, "Registered Successfully")
}