package handler

import (
	"net/http"
	"strconv"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
)

func (h *faqHandler) GetFAQs(ctx *gin.Context) {

	logger.Log().Info("FAQ Handler: GetFAQs API called")

	result, err := h.controller.GetFAQs()
	if err != nil {

		logger.Log().Error(err.Error())

		ctx.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Sugar().Infof("FAQ Handler: Total FAQs returned: %d", len(result.Data))

	ctx.JSON(
		http.StatusOK,
		network.SuccessResponse(result),
	)
}

func (h *faqHandler) GetFAQByID(ctx *gin.Context) {

	logger.Log().Info("FAQ Handler: GetFAQByID API called")

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {

		logger.Log().Error("Invalid FAQ ID: " + ctx.Param("id"))

		ctx.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("invalid faq id"),
			),
		)
		return
	}

	result, err := h.controller.GetFAQByID(id)
	if err != nil {

		logger.Log().Error(err.Error())

		ctx.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Info("FAQ Handler: GetFAQByID completed successfully")

	ctx.JSON(
		http.StatusOK,
		network.SuccessResponse(result),
	)
}

func (h *faqHandler) GetFAQsByCategory(ctx *gin.Context) {

	logger.Log().Info("FAQ Handler: GetFAQsByCategory API called")

	category := ctx.Param("category")

	if category == "" {

		logger.Log().Error("Category is required")

		ctx.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("category is required"),
			),
		)
		return
	}

	result, err := h.controller.GetFAQsByCategory(category)
	if err != nil {

		logger.Log().Error(err.Error())

		ctx.JSON(
			http.StatusInternalServerError,
			network.FailureResponse(
				network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
			),
		)
		return
	}

	logger.Log().Sugar().Infof("FAQ Handler: Category '%s' returned %d FAQs", category, len(result.Data))

	ctx.JSON(
		http.StatusOK,
		network.SuccessResponse(result),
	)
}