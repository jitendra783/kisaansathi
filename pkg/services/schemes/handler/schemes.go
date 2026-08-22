package handler

import (
	"net/http"
	"strconv"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"

	"github.com/gin-gonic/gin"
)

// GetSchemes returns paginated schemes.
func (h *schemeHandler) GetSchemes(ctx *gin.Context) {

	logger.Log().Info("Scheme Handler: GetSchemes API called")

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	logger.Log().Sugar().Infof(
		"Pagination -> Page: %d, Limit: %d",
		page,
		limit,
	)

	resp, err := h.controller.GetSchemes(page, limit)
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

	logger.Log().Sugar().Infof(
		"Total Schemes Returned: %d",
		len(resp.Schemes),
	)

	for i, scheme := range resp.Schemes {
		logger.Log().Sugar().Infof(
			"Scheme %d: %+v",
			i+1,
			scheme,
		)
	}

	ctx.JSON(
		http.StatusOK,
		network.SuccessResponse(resp),
	)
}

// GetScheme returns scheme details.
func (h *schemeHandler) GetScheme(ctx *gin.Context) {

	logger.Log().Info("Scheme Handler: GetScheme API called")

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {

		logger.Log().Error("Invalid scheme id")

		ctx.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("invalid scheme id"),
			),
		)
		return
	}

	resp, err := h.controller.GetScheme(id)
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

	if resp.Scheme != nil {
		logger.Log().Sugar().Infof(
			"Scheme Details: %+v",
			*resp.Scheme,
		)
	}

	ctx.JSON(
		http.StatusOK,
		network.SuccessResponse(resp),
	)
}

// GetEligibleSchemes returns eligible schemes.
func (h *schemeHandler) GetEligibleSchemes(ctx *gin.Context) {

	logger.Log().Info("Scheme Handler: GetEligibleSchemes API called")

	state := ctx.Query("state")
	category := ctx.Query("category")
	farmerType := ctx.Query("farmer_type")

	logger.Log().Sugar().Infof(
		"Filters -> State: %s, Category: %s, FarmerType: %s",
		state,
		category,
		farmerType,
	)

	resp, err := h.controller.GetEligibleSchemes(
		state,
		category,
		farmerType,
	)
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

	logger.Log().Sugar().Infof(
		"Eligible Schemes Returned: %d",
		len(resp.Schemes),
	)

	ctx.JSON(
		http.StatusOK,
		network.SuccessResponse(resp),
	)
}

// GetStateSchemes returns schemes for a state.
func (h *schemeHandler) GetStateSchemes(ctx *gin.Context) {

	logger.Log().Info("Scheme Handler: GetStateSchemes API called")

	state := ctx.Query("state")
	if state == "" {

		logger.Log().Error("State is required")

		ctx.JSON(
			http.StatusBadRequest,
			network.FailureResponse(
				network.ApiErrors.BadRequest.WithErrorDescription("state is required"),
			),
		)
		return
	}

	logger.Log().Sugar().Infof(
		"Requested State: %s",
		state,
	)

	resp, err := h.controller.GetStateSchemes(state)
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

	logger.Log().Sugar().Infof(
		"State Schemes Returned: %d",
		len(resp.Schemes),
	)

	ctx.JSON(
		http.StatusOK,
		network.SuccessResponse(resp),
	)
}
