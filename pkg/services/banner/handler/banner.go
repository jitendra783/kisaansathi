package handler

import (
	"log"
	"net/http"
	"strconv"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/services/banner/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBanners(c *gin.Context) {

	logger.Log().Info("Banner Handler: GetBanners API called")

	resp, err := h.controller.GetBanners(c.Request.Context())
	if err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	log.Println("resp:", resp)

	logger.Log().Sugar().Infof("Banner Handler: Total banners returned: %d", len(resp))

	for i, banner := range resp {
		logger.Log().Sugar().Infof("Banner %d: %+v", i+1, *banner)
	}

	c.JSON(http.StatusOK, network.SuccessResponse(resp))
}

func (h *handler) GetActiveBanners(c *gin.Context) {

	logger.Log().Info("Banner Handler: GetActiveBanners API called")

	resp, err := h.controller.GetActiveBanners(c.Request.Context())
	if err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Sugar().Infof("Banner Handler: Total active banners: %d", len(resp))

	c.JSON(http.StatusOK, network.SuccessResponse(resp))
}

func (h *handler) GetBannerByID(c *gin.Context) {

	logger.Log().Info("Banner Handler: GetBannerByID API called")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		logger.Log().Error("Invalid Banner ID: " + c.Param("id"))

		c.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("invalid banner id"),
			),
		)
		return
	}

	resp, err := h.controller.GetBannerByID(c.Request.Context(), id)
	if err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Info("Banner fetched successfully")

	c.JSON(http.StatusOK, network.SuccessResponse(resp))
}

func (h *handler) CreateBanner(c *gin.Context) {

	logger.Log().Info("Banner Handler: CreateBanner API called")

	var req models.CreateBannerRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	if err := h.controller.CreateBanner(c.Request.Context(), &req); err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Info("Banner created successfully")

	c.JSON(http.StatusOK, network.SuccessResponse("Banner created successfully"))
}

func (h *handler) UpdateBanner(c *gin.Context) {

	logger.Log().Info("Banner Handler: UpdateBanner API called")

	var req models.UpdateBannerRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	if err := h.controller.UpdateBanner(c.Request.Context(), &req); err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Info("Banner updated successfully")

	c.JSON(http.StatusOK, network.SuccessResponse("Banner updated successfully"))
}

func (h *handler) DeleteBanner(c *gin.Context) {

	logger.Log().Info("Banner Handler: DeleteBanner API called")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		logger.Log().Error("Invalid Banner ID: " + c.Param("id"))

		c.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("invalid banner id"),
			),
		)
		return
	}

	if err := h.controller.DeleteBanner(c.Request.Context(), id); err != nil {

		logger.Log().Error(err.Error())

		c.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Info("Banner deleted successfully")

	c.JSON(http.StatusOK, network.SuccessResponse("Banner deleted successfully"))
}