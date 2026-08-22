package handler

import (
	"strconv"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/expert/controller"
	"kisaanSathi/pkg/services/expert/db"
	"kisaanSathi/pkg/services/expert/models"

	"github.com/gin-gonic/gin"
)

type expertHandler struct {
	controller controller.ExpertController
}

type ExpertHandler interface {
	GetExperts(ctx *gin.Context)
	GetExpert(ctx *gin.Context)
	GetExpertsByCategory(ctx *gin.Context)
	BookConsultation(ctx *gin.Context)
	GetConsultationHistory(ctx *gin.Context)
}

func NewExpertHandler(controller controller.ExpertController) ExpertHandler {
	return &expertHandler{
		controller: controller,
	}
}
func NewExpertController(repo repo.DataObject) controller.ExpertController {
	store := db.NewExpertStore(repo.Databases.PgDB)
	return controller.NewExpertController(store)
}

// GetExperts returns all experts.
func (h *expertHandler) GetExperts(ctx *gin.Context) {

	resp, err := h.controller.GetExperts(ctx.Request.Context())
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}

// GetExpert returns expert details by ID.
func (h *expertHandler) GetExpert(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription("invalid expert id"),
		)
		return
	}

	resp, err := h.controller.GetExpert(ctx.Request.Context(), id)
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}

// GetExpertsByCategory returns experts filtered by specialization.
func (h *expertHandler) GetExpertsByCategory(ctx *gin.Context) {

	category := ctx.Param("category")
	if category == "" {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription("category is required"),
		)
		return
	}

	resp, err := h.controller.GetExpertsByCategory(ctx.Request.Context(), category)
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}

// BookConsultation creates a consultation booking.
func (h *expertHandler) BookConsultation(ctx *gin.Context) {

	var req models.BookConsultationRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription(err.Error()),
		)
		return
	}

	resp, err := h.controller.BookConsultation(ctx.Request.Context(), &req)
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}

// GetConsultationHistory returns user's consultation history.
func (h *expertHandler) GetConsultationHistory(ctx *gin.Context) {

	userID, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		network.FailureResponse(
			network.ApiErrors.BadRequest.WithErrorDescription("invalid user id"),
		)
		return
	}

	resp, err := h.controller.GetConsultationHistory(ctx.Request.Context(), userID)
	if err != nil {
		logger.Log().Error(err.Error())
		network.FailureResponse(
			network.ApiErrors.InternalServerError.WithErrorDescription(err.Error()),
		)
		return
	}

	network.SuccessResponse(resp)
}
