package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/feedback/controller"
	"kisaanSathi/pkg/services/feedback/db"

	"github.com/gin-gonic/gin"
)

type feedbackHandler struct {
	controller controller.FeedbackController
}
type FeedbackHandler interface {
	CreateFeedback(ctx *gin.Context)
	GetFeedbacks(ctx *gin.Context)
}

func NewFeedbackHandler(controller controller.FeedbackController) FeedbackHandler {
	return &feedbackHandler{
		controller: controller,
	}
}

func NewFeedbackController(repo repo.DataObject) controller.FeedbackController {
	store := db.NewFeedbackStore(repo.Databases.PgDB)
	return controller.NewFeedbackController(store)
}
